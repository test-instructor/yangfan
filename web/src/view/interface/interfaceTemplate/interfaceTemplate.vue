<template>
  <div class="parent">
    <div class="left">
      <InterfaceTree
        menutype="1"
        @getTreeID="setTreeID"
        eventType="1"
      ></InterfaceTree>
    </div>
    <div class="right">
      <div class="right2">
        <div class="gva-search-box" style="display: flex">
          <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
            <el-form-item label="接口名称">
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
              >新增(http)</el-button
            >
            <el-button
              size="mini"
              type="primary"
              icon="plus"
              @click="openDialogGrpc"
              >新增(grpc)</el-button
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
            <el-button type="primary" @click="setUserConfig"
              >调试配置</el-button
            >
            <user-config />
          </div>
          <el-table
            ref="multipleTable"
            style="width: 100%"
            :show-header="false"
            :data="tableData"
            row-key="ID"
            @selection-change="handleSelectionChange"
            :cell-style="{ paddingTop: '4px', paddingBottom: '4px' }"
          >
            <el-table-column type="selection" width="75" />
            <el-table-column min-width="600" align="center">
              <template #default="scope">
                <div
                  v-if="scope.row.request"
                  class="block"
                  :class="`block_${scope.row.request.method.toLowerCase()}`"
                >
                  <span
                    class="block-method block_method_color"
                    :class="`block_method_${scope.row.request.method.toLowerCase()}`"
                  >
                    {{ scope.row.request.method }}
                  </span>
                  <div class="block">
                    <span
                      class="block-method block_method_color block_method_options"
                      v-if="scope.row.creator === 'yapi'"
                      :title="'从YAPI导入的接口'"
                    >
                      YAPI
                    </span>
                  </div>

                  <span class="block-method block_url">{{
                    scope.row.request.url
                  }}</span>
                  <span class="block-summary-description">{{
                    scope.row.name
                  }}</span>
                </div>
                <div v-if="scope.row.gRPC" class="block" :class="`block_put`">
                  <span
                    class="block-method block_method_color"
                    :class="`block_method_put`"
                  >
                    {{ "gRPC" }}
                  </span>
                  <div class="block">
                    <span
                      class="block-method block_method_color block_method_options"
                      v-if="scope.row.creator === 'yapi'"
                      :title="'从YAPI导入的接口'"
                    >
                      YAPI
                    </span>
                  </div>

                  <span class="block-method block_url">{{
                    scope.row.gRPC.url
                  }}</span>
                  <span class="block-summary-description">{{
                    scope.row.name
                  }}</span>
                </div>
              </template>
            </el-table-column>

            <el-table-column label="按钮组" min-width="240">
              <template #default="scope">
                <el-button
                  type="text"
                  icon="debug"
                  size="small"
                  class="table-button"
                  @click="runInterfaceTemplateFunc(scope.row)"
                  >调试
                </el-button>
                <el-button
                  v-if="scope.row.request"
                  type="text"
                  icon="edit"
                  size="small"
                  class="table-button"
                  @click="updateInterfaceTemplateFunc(scope.row)"
                  >变更
                </el-button>
                <el-button
                  v-if="scope.row.gRPC"
                  type="text"
                  icon="edit"
                  size="small"
                  class="table-button"
                  @click="updateInterfaceTemplateFuncGrpc(scope.row)"
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
      v-model="interfaceTempleFormVisible"
      :before-close="closeDialog"
      :visible.sync="interfaceTempleFormVisible"
      :close-on-click-modal="false"
      :close-on-press-escape="false"
      :title="dialogTitle"
      width="1380px"
      top="30px"
    >
      <InterfaceTempleForm
        @close="closeDialog"
        v-if="interfaceTempleFormVisible"
        :heights="heightDiv"
        :eventType="type"
        :cid="configId"
        :apiType="apiTypes"
        :formData="formDatas"
        ref="menuRole"
      >
      </InterfaceTempleForm>
    </el-dialog>
    <el-dialog
      v-model="interfaceTempleFormVisibleGrpc"
      :before-close="closeDialogGrpc"
      :visible.sync="interfaceTempleFormVisibleGrpc"
      :close-on-click-modal="false"
      :close-on-press-escape="false"
      :title="dialogTitleGrpc"
      width="1380px"
      top="30px"
    >
      <InterfaceTempleGrpcForm
        @close="closeDialogGrpc"
        v-if="interfaceTempleFormVisibleGrpc"
        :heights="heightDiv"
        :eventType="typeGrpc"
        :cid="configId"
        :apiType="apiTypes"
        :formData="formDatasGrpc"
        ref="menuRole"
      >
      </InterfaceTempleGrpcForm>
    </el-dialog>
    <el-dialog
      title="设置运行配置"
      v-model="userConfigDialog"
      :before-close="closeDialogUserConfig"
      :close-on-click-modal="false"
      :close-on-press-escape="false"
    >
      <el-form>
        <EnvConfig
          v-if="userConfigDialog"
          @configId="configIdFun"
          @envId="envIdFun"
        ></EnvConfig>
      </el-form>
      <el-button type="primary" @click="saveUserConfig">保存</el-button>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: "InterfaceTemplate",
};
</script>

<script setup>
import {
  createInterfaceTemplate,
  deleteInterfaceTemplate,
  deleteInterfaceTemplateByIds,
  updateInterfaceTemplate,
  findInterfaceTemplate,
  getInterfaceTemplateList,
  createUserConfig,
  getUserConfig,
} from "@/api/interfaceTemplate";

import { runApi } from "@/api/runTestCase";

// 全量引入格式化工具 请按需保留
import { formatDate } from "@/utils/format";
import { ElMessage, ElMessageBox } from "element-plus";
import { ref } from "vue";
import InterfaceTree from "@/view/interface/interfaceComponents/interfaceTree.vue";
import InterfaceTempleForm from "@/view/interface/interfaceTemplate/interfaceTemplateForm.vue";
import InterfaceTempleGrpcForm from "@/view/interface/interfaceTemplate/interfaceTemplateGrpcForm.vue";
import EnvConfig from "@/view/interface/interfaceComponents/envConfig.vue";
import { reactive } from "vue";
import { useRouter } from "vue-router";
import UserConfig from "@/view/interface/interfaceComponents/userConfig.vue";
const router = useRouter();
const configId = ref("");
const apiEnvId = ref("");

// 自动化生成的字典（可能为空）以及字段
const formDatas = reactive({
  name: "",
  request: reactive({
    agreement: "",
    method: "",
    url: "",
    params: "",
    headers: "",
    json: "",
    data: "",
  }),
  variables: "",
  extract: "",
  validate: "",
  hooks: "",
  apiMenuID: "",
});
const formDatasGrpc = reactive({
  name: "",
  gRPC: reactive({
    Timeout: 0,
    url: "",
    headers: "",
    body: "",
    type: "",
  }),
  variables: "",
  extract: "",
  validate: "",
  hooks: "",
  apiMenuID: "",
});

// =========== 表格控制部分 ===========
const page = ref(1);
const total = ref(0);
const pageSize = ref(10);
const tableData = ref([]);
const searchInfo = ref({});
let treeID = 0;
// 行为控制标记（弹窗内部需要增还是改）
const type = ref("");
const typeGrpc = ref("");
const requestTYpe = ref("http");
const apiTypes = 1;

// 重置
const onReset = () => {
  searchInfo.value = {};
};

const setTreeID = (val) => {
  treeID = val;
  getTableData();
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
  const table = await getInterfaceTemplateList({
    type: apiTypes,
    menu: treeID,
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
    deleteInterfaceTemplateFunc(row);
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
  const res = await deleteInterfaceTemplateByIds({ ids });
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "删除成功",
    });
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--;
    }
    deleteVisible.value = false;
    getTableData();
  }
};

// 更新行
const updateInterfaceTemplateFunc = async (row) => {
  type.value = "update";
  dialogTitle.value = "编辑接口";
  const res = await findInterfaceTemplate({ ID: row.ID });
  if (res.code === 0) {
    formDatas.value = res.data.reapicase;
    interfaceTempleFormVisible.value = true;
  }
};

const updateInterfaceTemplateFuncGrpc = async (row) => {
  typeGrpc.value = "update";
  dialogTitleGrpc.value = "编辑接口grpc";
  const res = await findInterfaceTemplate({ ID: row.ID });
  if (res.code === 0) {
    formDatasGrpc.value = res.data.reapicase;
    interfaceTempleFormVisibleGrpc.value = true;
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

const configIdFun = (id) => {
  configId.value = id;
  console.log("============", configId.value);
};

const envIdFun = (id) => {
  apiEnvId.value = id;
  console.log("============", apiEnvId.value);
};

const runInterfaceTemplateFunc = async (row) => {
  if (
    !userConfigs.value ||
    !userConfigs.value.api_config_id ||
    userConfigs.value.api_config_id < 1
  ) {
    ElMessage({
      type: "error",
      message: "请设置配置后再运行",
    });
    return;
  }
  let data = {
    caseID: row.ID,
    configID: userConfigs.value.api_config_id,
    run_type: 5,
  };
  if (
    userConfigs.value &&
    userConfigs.value.api_env_id &&
    userConfigs.value.api_env_id > 0
  ) {
    console.log("============", userConfigs.value.api_env_id);
    data["env"] = userConfigs.value.api_env_id;
  }
  const res = await runApi(data);
  if (res.code === 0) {
    reportDetailFunc(res.data.id);
  }
};

// 删除行
const deleteInterfaceTemplateFunc = async (row) => {
  const res = await deleteInterfaceTemplate({ ID: row.ID });
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "删除成功",
    });
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--;
    }
    getTableData();
  }
};

// 弹窗控制标记
const interfaceTempleFormVisible = ref(false);
const interfaceTempleFormVisibleGrpc = ref(false);
const dialogTitle = ref("");
const dialogTitleGrpc = ref("");
const heightDiv = ref(0);
heightDiv.value =
  window.screen.height - 480 > 530 ? 530 : window.screen.height - 480;

// 打开弹窗
const openDialog = () => {
  type.value = "create";
  dialogTitle.value = "新增接口";
  requestTYpe.value = "http";
  interfaceTempleFormVisible.value = true;
};

const openDialogGrpc = () => {
  console.log("=======grpc");
  typeGrpc.value = "create";
  dialogTitleGrpc.value = "新增接口Grpc";
  interfaceTempleFormVisibleGrpc.value = true;
};

// 关闭弹窗
const closeDialog = () => {
  interfaceTempleFormVisible.value = false;
  formDatas.value = reactive({
    name: "",
    request: reactive({
      agreement: "",
      method: "",
      url: "",
      params: "",
      headers: "",
      json: "",
      data: "",
    }),
    variables: "",
    extract: "",
    validate: "",
    hooks: "",
    apiMenuID: "",
  });
};

const closeDialogGrpc = () => {
  interfaceTempleFormVisibleGrpc.value = false;
  formDatasGrpc.value = reactive({
    name: "",
    gRPC: reactive({
      Timeout: 0,
      url: "",
      headers: "",
      body: "",
      type: "",
    }),
    variables: "",
    extract: "",
    validate: "",
    hooks: "",
    apiMenuID: "",
  });
};

const userConfigDialog = ref(false);
const setUserConfig = () => {
  console.log("userconfig", userConfigs.value);
  userConfigDialog.value = true;
};

const closeDialogUserConfig = () => {
  userConfigDialog.value = false;
};

const saveUserConfig = async () => {
  if (configId.value === "" || apiEnvId.value === "") {
    ElMessage({
      type: "error",
      message: "运行标签和环境变量不能为空，请设置后再运行",
    });
    return;
  }
  const res = await createUserConfig({
    api_config_id: configId.value,
    api_env_id: apiEnvId.value,
  });
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "保存成功",
    });
    userConfigDialog.value = false;
  }
};
const userConfigs = ref({
  api_config_id: "",
  api_env_id: "",
});
const getUserConfigs = async () => {
  let res = await getUserConfig();
  if (res.code === 0 && res.data) {
    if (res.data.api_env_id > 0) {
      userConfigs.value.api_env_id = res.data.api_env_id;
    }
    if (res.data.api_config_id > 0) {
      userConfigs.value.api_config_id = res.data.api_config_id;
    }
  }
};
getUserConfigs();
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

.block_post {
  border: 1px solid #49cc90;
  background-color: rgba(73, 204, 144, 0.1);
}

.block_method_post {
  background-color: #49cc90;
}

.block_put {
  border: 1px solid #fca130;
  background-color: rgba(252, 161, 48, 0.1);
}

.block_method_put {
  background-color: #fca130;
}

.block_get {
  border: 1px solid #61affe;
  background-color: rgba(97, 175, 254, 0.1);
}

.block_method_get {
  background-color: #61affe;
}

.block_delete {
  border: 1px solid #f93e3e;
  background-color: rgba(249, 62, 62, 0.1);
}

.block_method_delete {
  background-color: #f93e3e;
}

.block_patch {
  border: 1px solid #50e3c2;
  background-color: rgba(80, 227, 194, 0.1);
}

.block_method_patch {
  background-color: #50e3c2;
}

.block_head {
  border: 1px solid #e6a23c;
  background-color: rgba(230, 162, 60, 0.1);
}

.block_method_head {
  background-color: #e6a23c;
}

.block_options {
  border: 1px solid #409eff;
  background-color: rgba(64, 158, 255, 0.1);
}

.block_method_options {
  background-color: #409eff;
}

.block {
  position: relative;
  border-radius: 4px;
  height: 48px;
  overflow: hidden;
  padding: 5px;
  display: flex;
  align-items: center;
}

.block_url {
  word-break: normal;
  width: auto;
  display: block;
  white-space: pre-wrap;
  word-wrap: break-word;
  overflow: hidden;
  -webkit-box-flex: 1;
  -ms-flex: 1;
  font-family: Open Sans, sans-serif;
  color: #3b4151;
}

.block_name {
  padding-left: 5px;
  word-break: normal;
  width: auto;
  display: block;
  white-space: pre-wrap;
  word-wrap: break-word;
  overflow: hidden;
  -webkit-box-flex: 1;
  -ms-flex: 1;
  font-family: Open Sans, sans-serif;
  color: #3b4151;
}

.block_method_color {
  cursor: pointer;
  color: #fff;
}

.block-method {
  font-size: 14px;
  font-weight: 600;
  min-width: 50px;
  padding: 0px 10px;
  text-align: center;
  border-radius: 5px;
  text-shadow: 0 1px 0 rgba(0, 0, 0, 0.1);
  font-family: Titillium Web, sans-serif;
}

.block-summary-description {
  word-break: normal;
  width: auto;
  display: block;
  white-space: pre-wrap;
  word-wrap: break-word;
  overflow: hidden;
  -webkit-box-flex: 1;
  -ms-flex: 1;
  font-family: Open Sans, sans-serif;
  color: #3b4151;
}
</style>
