<template>
  <div>
    <div>
      <div class="gva-search-box">
        <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
          <el-form-item label="变量key">
            <el-input v-model="searchInfo.key" placeholder="搜索条件" />
          </el-form-item>
          <el-form-item label="变量名称">
            <el-input v-model="searchInfo.name" placeholder="搜索条件" />
          </el-form-item>
          <el-form-item>
            <el-button
              icon="search"
              size="mini"
              type="primary"
              @click="onSubmit()"
              >查询</el-button
            >
            <el-button icon="refresh" size="mini" @click="onReset()"
              >重置</el-button
            >
          </el-form-item>
          <el-form-item>
            <el-button
              icon="plus"
              size="mini"
              type="primary"
              @click="openDialogEnv()"
              >环境管理</el-button
            >
          </el-form-item>
        </el-form>
      </div>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button
          icon="plus"
          size="mini"
          type="primary"
          @click="openDialogEnv()"
          >环境管理</el-button
        >
        <el-button
          icon="plus"
          size="mini"
          type="primary"
          @click="openDialogEnvVar()"
          >添加变量</el-button
        >
      </div>
      <el-table
        ref="multipleTable"
        row-key="ID"
        :data="envVarData"
        style="width: 100%"
        tooltip-effect="dark"
        @selection-change="handleSelectionChange"
      >
        <el-table-column label="变量key" prop="key" width="150" />

        <el-table-column label="变量名" prop="name" width="150" />
        <el-table-column
          v-for="title in envTableData"
          :label="title.name"
          align="center"
          :key="title.ID"
          :index="title.index"
        >
          <template #default="scope">
            {{ scope.row.value[title.ID] }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="按钮组" width="160">
          <template #default="scope">
            <el-button
              class="table-button"
              icon="edit"
              size="small"
              type="text"
              @click="updateRow(scope.row)"
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
    <el-dialog
      v-model="envDialog"
      :before-close="closeDialogEnv"
      :close-on-click-modal="false"
      :close-on-press-escape="false"
      title="环境管理"
    >
      <env v-if="envDialog"></env>
    </el-dialog>
    <el-dialog
      v-model="envVarDialog"
      :before-close="closeDialogEnvVar"
      :close-on-click-modal="false"
      :close-on-press-escape="false"
      :title="envVarType"
    >
      <el-form :model="varForm" label-width="120px">
        <el-form-item label="变量key">
          <el-input v-model="varForm.key" :disabled="envVarType === '更新变量'">
            <template v-if="envVarType != '更新变量'" #prepend>env_</template>
          </el-input>
        </el-form-item>
        <el-form-item label="变量名称">
          <el-input v-model="varForm.name" />
        </el-form-item>
        <el-divider />
        <el-form-item v-for="(item, index) of envTableData" :label="item.name">
          <el-input v-model="varForm.value[String(item.ID)]" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="addEnvVarFunc" round
            >保存</el-button
          >
          <el-button type="primary" @click="closeDialogEnvVar" round
            >取消</el-button
          >
        </el-form-item>
      </el-form>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: "environment",
};
</script>

<script setup>
import { ref } from "vue";
import Env from "@/view/interface/environment/env.vue";
import {
  getEnvList,
  createEnvVariable,
  getEnvVariableList,
  deleteEnvVariable,
  findEnvVariable,
} from "@/api/env";
import { ElMessage, ElMessageBox } from "element-plus";

// =========== 表格控制部分 ===========
const page = ref(1);
const total = ref(0);
const pageSize = ref(10);
const tableData = ref([]);
const searchInfo = ref({});

const envTableData = ref([]);
const init = async () => {
  const tables = await getEnvList();
  envTableData.value = tables.data.list;
  console.log("envTableData", envTableData.value);
  const res = await getTableData();
};
init();

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

const getTableData = async () => {
  const table = await getEnvVariableList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value,
  });
  if (table.code === 0) {
    envVarData.value = table.data.list;
    total.value = table.data.total;
    page.value = table.data.page;
    pageSize.value = table.data.pageSize;
  }
};

const getVar = (val) => {
  console.log("=======val", val);
};

const onReset = () => {
  searchInfo.value = {};
};

const onSubmit = () => {
  page.value = 1;
  pageSize.value = 10;
  getTableData();
};

const envDialog = ref(false);
const envVarDialog = ref(false);
const envVarType = ref("");
const envVarData = ref([]);
const varForm = ref({
  key: "",
  name: "",
  value: {},
});
const varFormValue = ref({});
const openDialogEnv = () => {
  envDialog.value = true;
};

const openDialogEnvVar = () => {
  envVarDialog.value = true;
  envVarType.value = "添加变量";
};

const closeDialogEnv = () => {
  envDialog.value = false;
  init();
};
const closeDialogEnvVar = () => {
  envVarDialog.value = false;
  varForm.value = {
    key: "",
    name: "",
    value: {},
  };
  getTableData();
};
// 多选数据
const multipleSelection = ref([]);
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val;
};

const addEnvVarFunc = async () => {
  console.log("==============", varForm.value);

  if (varForm.value.key === "") {
    ElMessage({
      type: "warning",
      message: "变量key不能为空",
    });
    return;
  }
  if (varForm.value.name === "") {
    ElMessage({
      type: "warning",
      message: "变量名不能为空",
    });
    return;
  }
  let newVarForm = varForm.value;
  if (envVarType.value === "添加变量") {
    newVarForm.key = "env_" + varForm.value.key;
  }
  let res = await createEnvVariable(varForm.value);
  let message = "变量操作成功";
  switch (envVarType.value) {
    case "添加变量":
      message = "变量创建成功";
      break;
    case "更新变量":
      message = "变量更新成功";
      break;
    default:
      break;
  }
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: message,
    });
    closeDialogEnvVar();
  }
};

const deleteRow = (row) => {
  ElMessageBox.confirm(
    "删除变量会导致被应用的测试套件、测试用例、定时任务、性能测试任务无法正常运行，确定要删除吗?",
    "提示",
    {
      confirmButtonText: "确定",
      cancelButtonText: "取消",
      type: "error",
    }
  ).then(() => {
    deleteEnvVarRow(row);
  });
};
const deleteEnvVarRow = async (row) => {
  let res = await deleteEnvVariable({ ID: row.ID });
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "变量删除成功",
    });
    await getTableData();
  }
};

const updateRow = async (row) => {
  let res = await findEnvVariable({ ID: row.ID });
  if (res.code === 0) {
    envVarDialog.value = true;
    envVarType.value = "更新变量";
    varForm.value = res.data.env;
    console.log("varForm", varForm.value);
  }
};
</script>

<style scoped></style>