<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="报告名称">
          <el-input v-model="searchInfo.name" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="用例类型">
          <el-select
            v-model="searchInfo.type"
            placeholder="请选择用例类型"
            clearable
          >
            <el-option
              v-for="item in case_type"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="执行类型">
          <el-select
            v-model="searchInfo.runType"
            placeholder="请选择执行类型"
            clearable
          >
            <el-option
              v-for="item in run_type"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="search" @click="onSubmit"
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
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px">
            <el-button size="mini" type="text" @click="deleteVisible = false"
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
              >删除
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
        <el-table-column align="left" label="执行时间" width="170">
          <template #default="scope">{{
            formatDate(scope.row.CreatedAt)
          }}</template>
        </el-table-column>
        <el-table-column align="left" label="用例类型" width="100">
          <template #default="scope">
            <el-tag
              :style="{
                background: caseType(scope.row.runType)[1],
                color: caseType(scope.row.runType)[2],
                border: 'none',
              }"
            >
              {{ caseType(scope.row.runType)[0] }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="执行类型" width="100">
          <template #default="scope">
            <el-tag
              :style="{
                background: runType(scope.row.type)[1],
                color: runType(scope.row.type)[2],
                border: 'none',
              }"
              effect="dark"
              >{{ runType(scope.row.type)[0] }}</el-tag
            >
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="报告名称"
          prop="name"
          width="220"
        />
        <el-table-column
          align="left"
          label="运行节点"
          prop="hostname"
          width="120"
        />
        <el-table-column
          align="left"
          label="运行环境"
          prop="api_env_name"
          width="190"
        />
        <el-table-column align="left" label="用例总数" width="100">
          <template #default="scope">
            <el-tag>{{
              scope.row.stat ? scope.row.stat.testcases.total : "-"
            }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="成功用例" width="100">
          <template #default="scope">
            <el-tag type="success">{{
              scope.row.stat ? scope.row.stat.testcases.success : "-"
            }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="失败用例" width="100">
          <template #default="scope">
            <el-tag type="danger">{{
              scope.row.stat ? scope.row.stat.testcases.fail : "-"
            }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="运行时长/秒" width="120">
          <template #default="scope">
            <!--            {{ scope.row.time.duration?Number(scope.row.time.duration).toFixed(2):0 }}-->
            {{ durations(scope.row) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="状态" prop="default" width="120">
          <template #default="scope">
            <el-tag :type="successType(scope.row)[0]" effect="dark">{{
              successType(scope.row)[1]
            }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="按钮组">
          <template #default="scope">
            <el-button
              type="text"
              size="small"
              class="table-button"
              @click="reportDetailFunc(scope.row)"
              >详情</el-button
            >
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
</template>

<script>
export default {
  name: "ApiReport",
};
</script>

<script setup>
import { deleteApiConfigByIds, findApiConfig } from "@/api/apiConfig";

import { delReport, getReportList } from "@/api/report";
import ReportDetail from "@/view/interface/Reports/reportDetail.vue";
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

const router = useRouter();

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  name: "",
  base_url: "",
  default: false,
});

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
  if (searchInfo.value.default === "") {
    searchInfo.value.default = null;
  }
  getTableData();
};

const successType = (row) => {
  if (row.status === 0) {
    return ["info", "运行中"];
  }
  if (row.status === 1) {
    if (row.success) {
      return ["success", "成功"];
    } else {
      return ["warning", "失败"];
    }
  }
  if (row.status === 2) {
    return ["danger", "错误"];
  }
  return ["info", "未知状态"];
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

const caseType = (t) => {
  if (t === 1) {
    return ["\u3000\xa0api\u3000\xa0", "#89E9E0", "#ffffff"];
  }
  if (t === 2) {
    return ["\xa0\xa0套\u3000件\xa0\xa0", "#3FD4CF", "#ffffff"];
  }
  if (t === 3) {
    return ["\xa0\xa0用\u3000例\xa0\xa0", "#1FA6AA", "#ffffff"];
  }
  if (t === 4) {
    return ["定时任务", "#4080FF", "#ffffff"];
  }
  if (t === 5) {
    return ["性能测试", "#D91AD9", "#ffffff"];
  }
  if (t === 7) {
    return ["任务标签", "#1664FF", "#ffffff"];
  }
  if (t === 8) {
    return ["CI运行", "#1664FF", "#ffffff"];
  }
  return ["定时任务"];
};

const case_type = ref([
  { value: 1, label: "api" },
  { value: 2, label: "套件" },
  { value: 3, label: "用例" },
  { value: 4, label: "定时任务" },
  { value: 5, label: "性能测试" },
  { value: 7, label: "任务标签" },
  { value: 8, label: "CI调用" },
]);

const run_type = ref([
  { value: 1, label: "调试运行" },
  { value: 2, label: "立即运行" },
  { value: 3, label: "后台运行" },
  { value: 4, label: "定时执行" },
  { value: 5, label: "保存调试" },
  { value: 8, label: "CI" },
]);

const runType = (t) => {
  if (t === 5) {
    return ["保存调试", "#FF9A2E"];
  }
  if (t === 1) {
    return ["调试运行", "#F76560"];
  }
  if (t === 2) {
    return ["立即运行", "#C1FFC1"];
  }
  if (t === 3) {
    return ["后台运行", "#5EDFD6"];
  }
  if (t === 4) {
    return ["定时执行", "#C9E968"];
  }
  if (t === 6) {
    return ["CI 调用", "#C9E968"];
  }
  return 7;
};

const durations = (row) => {
  const t = ref(0);
  if (row.time && row.time.duration) {
    let hours = Math.floor(row.time.duration / 3600);
    let minutes = Math.floor((row.time.duration % 3600) / 60);
    let seconds = Math.floor(row.time.duration % 60);

    let formattedTime = "";

    if (hours > 0) {
      formattedTime += hours + "小时";
    }

    if (minutes > 0 || (hours > 0 && seconds > 0)) {
      formattedTime += minutes + "分钟";
    }

    if (seconds <= 0 && formattedTime === "") {
      formattedTime = "1秒";
    } else if (seconds > 0) {
      formattedTime += seconds + "秒";
    }

    return formattedTime;
  }
  return t;
};

const reportDetailFunc = (row) => {
  if (row.status === 2) {
    ElMessageBox.alert(row.describe, "运行错误", {
      type: "error",
    });
  } else {
    router.push({
      name: "reportDetail",
      params: {
        report_id: row.ID,
      },
    });
  }
};

// 查询
const getTableData = async () => {
  const table = await getReportList({
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

getTableData();

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () => {};

// 获取需要的字典 可能为空 按需保留
setOptions();

const formDatas = reactive({
  name: "",
  base_url: "",
  headers: "",
  variables: "",
  extract: "",
  validate: "",
  hooks: "",
  apiMenuID: "",
  Parameters: "",
});

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
    deleteReportFunc(row);
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
  const res = await deleteApiConfigByIds({ ids });
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
const updateApiConfigFunc = async (row) => {
  const res = await findApiConfig({ ID: row.ID });
  type.value = "update";
  dialogTitle.value = "编辑配置";
  if (res.code === 0) {
    formData.value = res.data.reac;
    formDatas.value = res.data.reac;
    dialogFormVisible.value = true;
  }
};

// 删除行
const deleteReportFunc = async (row) => {
  const res = await delReport({ ID: row.ID });
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
const dialogTitle = ref(false);
const heightDiv = ref();
heightDiv.value = window.screen.height - 480;
</script>

<style></style>
