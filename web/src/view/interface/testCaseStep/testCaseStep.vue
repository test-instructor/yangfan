<template>
  <div class="parent">
    <div class="left">
      <InterfaceTree
        menutype="2"
        @getTreeID="setTreeID"
        eventType="1"
      ></InterfaceTree>
    </div>
    <div class="right">
      <div class="right2">
        <div class="gva-search-box" style="display: flex">
          <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
            <el-form-item label="步骤名称">
              <el-input v-model="searchInfo.name" placeholder="搜索条件" />
            </el-form-item>
            <el-form-item>
              <el-button
                size="mini"
                type="primary"
                icon="search"
                @click="onSubmit"
                >查询</el-button
              >
              <el-button size="mini" icon="refresh" @click="onReset"
                >重置</el-button
              >
            </el-form-item>
          </el-form>
        </div>
        <div class="gva-table-box">
          <div class="gva-btn-list">
            <el-button
              size="mini"
              type="primary"
              icon="plus"
              @click="openDialog"
              >新增</el-button
            >
            <el-popover
              v-model:visible="deleteVisible"
              placement="top"
              width="160"
            >
              <p>确定要删除吗？</p>
              <div style="text-align: right; margin-top: 8px">
                <el-button
                  size="mini"
                  type="text"
                  @click="deleteVisible = false"
                  >取消</el-button
                >
                <el-button size="mini" type="primary" @click="onDelete"
                  >确定</el-button
                >
              </div>
              <template #reference>
                <el-button
                  icon="delete"
                  size="mini"
                  style="margin-left: 10px"
                  :disabled="!multipleSelection.length"
                >
                  删除
                </el-button>
              </template>
            </el-popover>
          </div>
          <el-table
            ref="multipleTable"
            style="width: 100%"
            tooltip-effect="dark"
            :data="tableData"
            row-key="ID"
            @selection-change="handleSelectionChange"
          >
            <el-table-column type="selection" width="55" />
            <el-table-column align="left" label="日期" width="180">
              <template #default="scope">{{
                formatDate(scope.row.CreatedAt)
              }}</template>
            </el-table-column>
            <el-table-column
              align="left"
              label="步骤名称"
              prop="name"
              width="200"
            />
            <el-table-column
              align="left"
              label="调试运行配置"
              prop="RunConfigName"
              width="200"
            />
            <el-table-column
              align="left"
              label="运行环境"
              prop="api_env_name"
              width="200"
            />
            <el-table-column align="left" label="按钮组" min-width="300">
              <template #default="scope">
                <el-button
                  type="text"
                  icon="detail"
                  size="small"
                  class="table-button"
                  @click="runCase(scope.row, 1)"
                  >调试运行
                </el-button>
                <el-button
                  type="text"
                  icon="detail"
                  size="small"
                  class="table-button"
                  @click="runCase(scope.row, 3)"
                  >后台运行
                </el-button>
                <el-button
                  type="text"
                  icon="detail"
                  size="small"
                  class="table-button"
                  @click="detailTestCaseFunc(scope.row)"
                  >详情
                </el-button>
                <el-button
                  type="text"
                  icon="edit"
                  size="small"
                  class="table-button"
                  @click="updateTestCaseFunc(scope.row)"
                  >变更
                </el-button>
                <el-button
                  type="text"
                  icon="delete"
                  size="mini"
                  @click="deleteRow(scope.row)"
                  >删除</el-button
                >
              </template>
            </el-table-column>
          </el-table>
          <div class="gva-pagination">
            <el-pagination
              layout="total, sizes, prev, pager, next, jumper"
              :current-page="page"
              :page-size="pageSize"
              :page-sizes="[10, 30, 50, 100]"
              :total="total"
              @current-change="handleCurrentChange"
              @size-change="handleSizeChange"
            />
          </div>
        </div>
      </div>
    </div>
    <el-dialog
      v-model="dialogFormVisible"
      :before-close="closeDialog"
      title="弹窗操作"
    >
      <el-form :model="formData" label-position="right" label-width="120px">
        <el-form-item label="步骤名称:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="调试运行配置:">
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
        <el-form-item label="前置步骤:">
          <el-switch v-model="formData.front_case" />
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
  </div>
</template>

<script>
export default {
  name: "TestCase",
};
</script>

<script setup>
import {
  createTestCase,
  deleteTestCase,
  deleteTestCaseByIds,
  updateTestCase,
  findTestCase,
  getTestCaseList,
} from "@/api/testCase";

// 全量引入格式化工具 请按需保留
import {
  getDictFunc,
  formatDate,
  formatBoolean,
  filterDict,
} from "@/utils/format";
import { ElMessage, ElMessageBox } from "element-plus";
import { ref, reactive } from "vue";
import { useRouter } from "vue-router";
import { getApiConfigList } from "@/api/apiConfig";
import { runTestCaseStep } from "@/api/runTestCase";
import InterfaceTree from "@/view/interface/interfaceComponents/interfaceTree.vue";
import envConfig from "@/view/interface/interfaceComponents/envConfig.vue";
import { getEnvList } from "@/api/env";

const router = useRouter();

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  name: "",
  type: 2,
  front_case: false,
  RunConfigID: 0,
  api_env_id: 0,
});

// =========== 表格控制部分 ===========
const page = ref(1);
const total = ref(0);
const pageSize = ref(10);
const tableData = ref([]);
const searchInfo = ref({});
const configID = ref();
const apiEnvID = ref();
const configChange = (key) => {
  formData.value.RunConfigID = key;
};
const envChange = (key) => {
  formData.value.api_env_id = null;
  if (key && key > 0) {
    formData.value.api_env_id = key;
  }
};
const params = reactive({
  menu: "",
});

const value = ref("");
let treeID = 0;

const setTreeID = (val) => {
  treeID = val;
  getTableData();
};

// 重置
const onReset = () => {
  searchInfo.value = {};
};

// 搜索
const onSubmit = () => {
  page.value = 1;
  pageSize.value = 10;
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

// 查询
const getTableData = async () => {
  let menu = treeID;
  const table = await getTestCaseList({
    menu,
    page: page.value,
    pageSize: pageSize.value,
    type: 2,
    ...searchInfo.value,
  });
  if (table.code === 0) {
    tableData.value = table.data.list;
    total.value = table.data.total;
    page.value = table.data.page;
    pageSize.value = table.data.pageSize;
  }
};
getTableData();
const configData = ref([]);
const getConfigData = async () => {
  const config = await getApiConfigList({ page: 1, pageSize: 777777 });
  if (config.code === 0) {
    configData.value = config.data.list;
  }
};

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () => {};

// 获取需要的字典 可能为空 按需保留
setOptions();

// 多选数据
const multipleSelection = ref([]);
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val;
};

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm("确定要删除吗?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(() => {
    deleteTestCaseFunc(row);
  });
};

// 批量删除控制标记
const deleteVisible = ref(false);

// 多选删除
const onDelete = async () => {
  const ids = [];
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: "warning",
      message: "请选择要删除的数据",
    });
    return;
  }
  multipleSelection.value &&
    multipleSelection.value.map((item) => {
      ids.push(item.ID);
    });
  const res = await deleteTestCaseByIds({ ids });
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "删除成功",
    });
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--;
    }
    deleteVisible.value = false;
    await getTableData();
  }
};

// 行为控制标记（弹窗内部需要增还是改）
const type = ref("");

// 更新行
const updateTestCaseFunc = async (row) => {
  await getConfigData();
  await getApiEnv();
  const res = await findTestCase({ ID: row.ID });
  type.value = "update";
  if (res.code === 0) {
    formData.value = res.data.reapicase;
    configID.value = formData.value.RunConfigID;
    if (formData.value.api_env_id > 0) {
      apiEnvID.value = formData.value.api_env_id;
    }
    dialogFormVisible.value = true;
  }
};

const detailTestCaseFunc = (row) => {
  if (row) {
    router.push({
      name: "testCaseStepDetail",
      params: {
        id: row.ID,
      },
    });
  } else {
    router.push({ name: "testCaseStepDetail" });
  }
};

const reportDetailFunc = (report_id) => {
  if (report_id) {
    router.push({
      name: "reportDetail",
      params: {
        report_id: report_id,
      },
    });
  } else {
    router.push({ name: "reportDetail" });
  }
};

const runCase = async (row, runType) => {
  let data = { caseID: Number(row.ID), run_type: runType };
  const res = await runTestCaseStep(data);
  if (res.code === 0) {
    if (runType === 1) {
      reportDetailFunc(res.data.id);
    } else {
      ElMessage({
        type: "success",
        message: "运行成功",
      });
    }
  }
};

// 删除行
const deleteTestCaseFunc = async (row) => {
  const res = await deleteTestCase({ ID: row.ID });
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

// 弹窗控制标记
const dialogFormVisible = ref(false);

// 打开弹窗
const openDialog = () => {
  getConfigData();
  getApiEnv();
  type.value = "create";
  dialogFormVisible.value = true;
};

// 关闭弹窗
const closeDialog = () => {
  configID.value = "";
  apiEnvID.value = "";
  dialogFormVisible.value = false;
  formData.value = {
    name: "",
    type: 2,
    front_case: false,
    RunConfigID: 0,
    runConfig: {
      ID: 0,
    },
  };
};
// 弹窗确定
const enterDialog = async () => {
  if (formData.value.name === "") {
    ElMessage({
      type: "error",
      message: "步骤名称不能为空",
    });
    return;
  }
  if (formData.value.RunConfigID === 0) {
    ElMessage({
      type: "error",
      message: "请选择调试运行配置",
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
  params.menu = treeID;
  let res;
  configData.value.forEach((item, index, arr) => {
    if (item.ID === formData.value.RunConfigID) {
      formData.value.RunConfigName = item.name;
    }
  });
  apiEnvData.value.forEach((item, index, arr) => {
    if (item.ID === formData.value.api_env_id) {
      formData.value.api_env_name = item.name;
    }
  });
  switch (type.value) {
    case "create":
      res = await createTestCase(formData.value, params);
      break;
    case "update":
      console.log("update", formData.value);
      res = await updateTestCase(formData.value, params);
      break;
    default:
      res = await createTestCase(formData.value, params);
      break;
  }
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "创建/更改成功",
    });
    closeDialog();
    await getTableData();
  }
};

const apiEnvData = ref([]);
const getApiEnv = async () => {
  const res = await getEnvList();
  if (res.code === 0) {
    apiEnvData.value = res.data.list;
  }
};
</script>

<style>
.parent {
  width: 100%;
  height: 85%;
}

.left {
  margin-top: 10px;
  width: 300px;
  height: 98%;
  padding: 8px;
  background: #ffffff;
}

.right {
  display: flex;
  flex: 1;
  height: 100%;
  margin-right: 10px;
}

.right2 {
  flex: 1;
  height: 100%;
  padding: 10px;
}
</style>
