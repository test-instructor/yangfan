<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="任务名称">
          <el-input v-model="searchInfo.name" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item>
          <el-button icon="search" size="mini" type="primary" @click="onSubmit"
            >查询</el-button
          >
          <el-button icon="refresh" size="mini" @click="onReset"
            >重置</el-button
          >
        </el-form-item>
      </el-form>
    </div>

    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button icon="plus" size="mini" type="primary" @click="openDialog"
          >新增</el-button
        >
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px">
            <el-button size="mini" type="text" @click="deleteVisible = false"
              >取消</el-button
            >
            <!--            <el-button size="mini" type="primary" @click="onDelete">确定</el-button>-->
          </div>
          <template #reference>
            <el-button
              icon="delete"
              size="mini"
              style="margin-left: 10px"
              :disabled="!multipleSelection.length"
              >删除
            </el-button>
          </template>
        </el-popover>
      </div>
      <el-table
        ref="multipleTable"
        :data="tableData"
        row-key="ID"
        style="width: 100%"
        tooltip-effect="dark"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column
          align="left"
          label="任务名称"
          prop="name"
          width="240"
        />
        <el-table-column
          align="left"
          label="运行配置"
          prop="config.name"
          width="240"
        />
        <el-table-column
          align="left"
          label="运行环境"
          prop="api_env_name"
          width="240"
        />
        <el-table-column
          align="left"
          label="备注"
          prop="describe"
          width="280"
        />
        <el-table-column align="left" label="按钮组" min-width="460">
          <template #default="scope">
            <el-button
              type="text"
              icon="detail"
              size="small"
              class="table-button"
              @click="runCase(scope.row)"
              >启动压测</el-button
            >
            <el-button
              type="text"
              icon="detail"
              size="small"
              class="table-button"
              @click="runCaseDebug(scope.row, 6)"
              >调试运行</el-button
            >
            <el-button
              class="table-button"
              icon="detail"
              size="small"
              type="text"
              @click="detailFunc(scope.row)"
              >任务详情</el-button
            >
            <el-button
              class="table-button"
              icon="edit"
              size="small"
              type="text"
              @click="updateFunc(scope.row)"
              >变更
            </el-button>
            <el-button
              icon="delete"
              size="mini"
              type="text"
              @click="deleteRow(scope.row)"
              >删除</el-button
            >
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>

    <el-dialog
      v-model="dialogFormVisible"
      :before-close="closeDialog"
      :close-on-click-modal="false"
      :close-on-press-escape="false"
      :title="type === 'create' ? '新增性能任务' : '编辑性能任务'"
    >
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="任务名称:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>

        <el-form-item label="运行配置:">
          <el-select
            v-model="configID"
            placeholder="请选择"
            @change="configChange"
          >
            <el-option
              v-for="item in configData"
              :key="item.ID"
              :label="item.name"
              :value="item.ID"
            />
          </el-select>
        </el-form-item>

        <el-form-item label="环境变量:">
          <el-select
            v-model="apiEnvID"
            placeholder="请选择"
            @change="envChange"
          >
            <el-option
              v-for="item in apiEnvData"
              :key="item.ID"
              :label="item.name"
              :value="item.ID"
            />
          </el-select>
        </el-form-item>

        <el-form-item label="备注:">
          <el-input
            v-model="formData.describe"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog"
            >确 定</el-button
          >
        </div>
      </template>
    </el-dialog>

    <el-dialog
      v-model="dialogRunner"
      :before-close="closeRunner"
      :close-on-click-modal="false"
      :close-on-press-escape="false"
      width="400px"
      title="启动压测"
    >
      <el-form :model="runnerConfig" label-position="right" label-width="160px">
        <el-form-item label="并发用户数：">
          <el-input-number
            v-model="runnerConfig.spawnCount"
            :min="1"
            step-strictly
          />
        </el-form-item>
        <el-form-item label="初始每秒增加用户数：">
          <el-input-number
            v-model="runnerConfig.spawnRate"
            :min="1"
            step-strictly
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeRunner">取 消</el-button>
          <el-button size="small" type="primary" @click="runCaseConfig"
            >启动压测</el-button
          >
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: "index",
};
</script>

<script setup>
import { ref } from "vue";
import { getApiConfigList } from "@/api/apiConfig";
import { ElMessage, ElMessageBox } from "element-plus";
import {
  createPerformance,
  deletePerformance,
  findPerformance,
  getPerformanceList,
  updatePerformance,
} from "@/api/performance";
import {
  deleteTimerTask,
  findTimerTask,
  getTimerTaskList,
} from "@/api/timerTask";
import { useRouter } from "vue-router";
import { runBoomer, runBoomerDebug } from "@/api/runTestCase";
import { getEnvList } from "@/api/env";

const router = useRouter();
const deleteVisible = ref(false);
// 多选数据
const multipleSelection = ref([]);
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val;
};

const creatCron = ref(false);
const dialogFormVisible = ref(false);
const dialogRunner = ref(false);
let runnerConfig = {
  spawnCount: 1,
  spawnRate: 1,
};
const configID = ref();

// =========== 表格控制部分 ===========
const page = ref(1);
const total = ref(0);
const pageSize = ref(10);
const tableData = ref([]);
const searchInfo = ref({});

// 重置
const onReset = () => {
  searchInfo.value = {};
};

// 搜索
const onSubmit = () => {
  page.value = 1;
  pageSize.value = 10;
  if (searchInfo.value.status === "") {
    searchInfo.value.status = null;
  }
  getTableData();
};

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val;
  getTableData();
};

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val;
  getTableData();
};

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm("确定要删除吗?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(() => {
    deletePerformanceFunc(row);
  });
};

const runCase = async (row) => {
  runnerConfig.caseID = Number(row.ID);
  dialogRunner.value = true;
};

const runCaseConfig = async () => {
  let data = {
    caseID: runnerConfig.caseID,
    run_type: 6,
    operation: {
      running: 1,
      spawnCount: runnerConfig.spawnCount,
      spawnRate: runnerConfig.spawnRate,
    },
  };
  const res = await runBoomer(data);
  if (res.code === 0) {
    closeRunner();
    ElMessage({
      type: "success",
      message: "运行成功",
    });
  }
};

const runCaseDebug = async (row, runType) => {
  let data = { caseID: Number(row.ID), run_type: runType };
  const res = await runBoomerDebug(data);
  if (res.code === 0) {
    reportDetailFunc(res.data.id);
    ElMessage({
      type: "success",
      message: "运行成功",
    });
  }
};

const reportDetailFunc = (ID) => {
  if (ID) {
    router.push({
      name: "reportDetail",
      params: {
        id: ID,
      },
    });
  } else {
    router.push({ name: "reportDetail" });
  }
};

const deletePerformanceFunc = async (row) => {
  const res = await deletePerformance({ ID: row.ID });
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "删除成功",
    });
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--;
    }
    await getTableData();
  }
};

const formData = ref({
  name: "",
  describe: "",
  TestCase: [],
  RunConfigID: 0,
  api_env_id: 0,
});

const configChange = (key) => {
  formData.value.RunConfigID = key;
};
// 行为控制标记（弹窗内部需要增还是改）
const type = ref("");

const openDialog = () => {
  getConfigData();
  getApiEnv();
  type.value = "create";
  creatCron.value = true;
  dialogFormVisible.value = true;
};

const configData = ref([]);
const getConfigData = async () => {
  const config = await getApiConfigList({ page: 1, pageSize: 99999 });
  if (config.code === 0) {
    configData.value = config.data.list;
  }
};
// 弹窗关闭
const closeDialog = () => {
  configID.value = "";
  dialogFormVisible.value = false;
};

const closeRunner = () => {
  runnerConfig = {
    spawnCount: 1,
    spawnRate: 1,
  };
  dialogRunner.value = false;
  // formData.value = {
  //   name: '',
  //   runTime: '',
  //   nextRunTime: new Date(),
  //   status: false,
  //   describe: '',
  //   runNumber: 0,
  //   config: {ID: 0},
  // }
  // creatCron.value = false
};

// 更新行
const updateFunc = async (row) => {
  const res = await findPerformance({ ID: row.ID });
  await getConfigData();
  await getApiEnv();

  type.value = "update";
  if (res.code === 0) {
    formData.value = res.data.reapicase;
    configID.value = formData.value.RunConfigID;
    dialogFormVisible.value = true;
    if (formData.value.api_env_id > 0) {
      apiEnvID.value = formData.value.api_env_id;
    }
  }
};

// 弹窗确定
const enterDialog = async () => {
  if (formData.value.name === "") {
    ElMessage({
      type: "error",
      message: "任务名称不能为空",
    });
    return;
  }
  if (formData.value.RunConfigID === 0) {
    ElMessage({
      type: "error",
      message: "请选择运行配置",
    });
    return;
  }
  if (formData.value.api_env_id < 1) {
    ElMessage({
      type: "error",
      message: "请选择环境变量",
    });
    return;
  }
  let res;
  apiEnvData.value.forEach((item, index, arr) => {
    if (item.ID === formData.value.api_env_id) {
      formData.value.api_env_name = item.name;
    }
  });
  formData.value.TestCase = [];
  switch (type.value) {
    case "create":
      res = await createPerformance(formData.value);
      break;
    case "update":
      res = await updatePerformance(formData.value);
      break;
    default:
      res = await createPerformance(formData.value);
      break;
  }
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "创建/更改成功",
    });
    closeDialog();
    await getTableData();
    creatCron.value = false;
    formData.value = {
      name: "",
      describe: "",
      TestCase: [],
      RunConfigID: 0,
      api_env_id: 0,
    };
  }
};

// 查询
const getTableData = async () => {
  const table = await getPerformanceList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value,
  });
  if (table.code === 0) {
    tableData.value = table.data.list;
    total.value = table.data.total;
    page.value = table.data.page;
    pageSize.value = table.data.pageSize;
  }
};

const init = () => {
  getTableData();
};
init();

const detailFunc = (row) => {
  if (row) {
    router.push({
      name: "performanceDetail",
      params: {
        id: row.ID,
      },
    });
  } else {
    router.push({ name: "performanceDetail" });
  }
};

const apiEnvID = ref();

const envChange = (key) => {
  formData.value.api_env_id = null;
  if (key && key > 0) {
    formData.value.api_env_id = key;
  }
};

const apiEnvData = ref([]);
const getApiEnv = async () => {
  const res = await getEnvList();
  if (res.code === 0) {
    apiEnvData.value = res.data.list;
    console.log("==========", apiEnvData.value);
  }
};
</script>

<style scoped></style>
