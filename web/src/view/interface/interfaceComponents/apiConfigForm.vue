<template>
  <div>
    <div style="display: flex">
      <div>
        <el-form-item label="前置步骤：">
          <el-select
            v-model="reqData.setup_case_id"
            class="m-2"
            placeholder="请选择前置用例"
            size="small"
            clearable
          >
            <el-option
              v-for="item in CaseOptions"
              :label="item.name"
              :value="item.ID"
              :key="item.ID"
            />
          </el-select>
        </el-form-item>
      </div>

      <div
        style="width: 200px; margin-top: 5px; display: flex; margin-left: 20px"
      >
        <el-form-item label="重试次数：">
          <el-input-number
            v-model="reqData.retry"
            placeholder=""
            size="small"
            :min="0"
            :max="50"
          >
          </el-input-number>
        </el-form-item>
      </div>
      <div
        style="width: 150px; margin-top: 5px; display: table; margin-left: 20px"
      >
        <span>verify: </span>
        <el-switch
          v-model="reqData.verify"
          class="mb-2"
          size="large"
          name="verify"
        />
      </div>
      <div
        style="width: 150px; margin-top: 5px; display: table; margin-left: 20px"
      >
        <env-copy></env-copy>
      </div>
    </div>
    <div style="width: 1000px; margin-top: 10px; display: flex">
      <el-input v-model="reqData.name" placeholder="请输入配置名称">
        <template #prepend>配置名称</template>
      </el-input>
    </div>
    <div style="width: 1000px; margin-top: 10px; display: flex">
      <el-input v-model="reqData.base_url" placeholder="请输入域名">
        <template #prepend>域名</template>
      </el-input>
    </div>

    <div class="request">
      <el-tabs
        style="margin-left: 20px"
        v-model="activeTag"
        :heights="heightDiv"
      >
        <el-tab-pane label="Header" name="Header">
          <headers
            @headerData="handleHeader"
            @request="handleRequest"
            :header="reqData ? reqData.headers_json : []"
            :heights="heightDiv"
          >
          </headers>
        </el-tab-pane>

        <el-tab-pane label="变量" name="Variables">
          <variables
            :heights="heightDiv"
            @requestVariablesData="variablesVariables"
            @request="handleRequest"
            :variables="reqData ? reqData.variables_json : []"
          >
          </variables>
        </el-tab-pane>

        <el-tab-pane label="Parameters" name="params">
          <!--          <params-->
          <!--              @requestParamsData="requestParams"-->
          <!--              @request="handleRequest"-->
          <!--              :params="reqData ? reqData.parameters_json : []"-->
          <!--              :heights="heightDiv"-->
          <!--          >-->
          <!--          </params>-->
          <jsons
            @requestJsonData="requestParams"
            :heights="heightDiv"
            :jsons="reqData ? reqData.parameters : []"
          >
          </jsons>
        </el-tab-pane>
      </el-tabs>
    </div>
    <br />
    <el-button type="primary" @click="saves">保存</el-button>
    <el-button type="info" @click="closeDialog">取消</el-button>
  </div>
</template>

<script>
export default {
  name: "apiConfigForm",
};
</script>

<script setup>
import Headers from "@/view/interface/interfaceComponents/Headers.vue";
import Variables from "@/view/interface/interfaceComponents/Variables.vue";
import Jsons from "@/view/interface/interfaceComponents/Jsons.vue";
import { getTestCaseList } from "@/api/testCase";

import {
  createApiConfig,
  updateApiConfig,
  findApiConfig,
} from "@/api/apiConfig";

// 自动获取字典
import { getDictFunc } from "@/utils/format";
import { useRoute, useRouter } from "vue-router";
import { ElMessage } from "element-plus";
import { ref, reactive } from "vue";
import { getApiCaseList } from "@/api/apiCase";
import EnvCopy from "@/view/interface/interfaceComponents/envCopy.vue";
const emit = defineEmits(["close"]);
const route = useRoute();
const router = useRouter();
const type = ref("");
const formData = ref({
  name: "",
  base_url: "",
  default: false,
});
const CaseOptions = ref([]);
const props = defineProps({
  heights: ref(),
  eventType: ref(),
  formData: ref({
    ID: 0,
    parameters: [],
    name: "",
    base_url: "",
    headers: [],
    variables: "",
    extract: "",
    validate: "",
    hooks: [],
    apiMenuID: "",
    verify: false,
    default: false,
  }),
});

const heightDiv = ref(false);
const eventType = ref("");
heightDiv.value =
  window.screen.height - 480 > 530 ? 530 : window.screen.height - 480;
eventType.value = props.eventType;

const getTableData = async () => {
  const table = await getTestCaseList({
    page: 1,
    pageSize: 99999,
    front_case: true,
    type: 2,
  });
  if (table.code === 0) {
    CaseOptions.value = table.data.list;
  }
};

// 初始化方法
const init = async () => {
  await getTableData();
  if (route.query.id) {
    const res = await findApiConfig({ ID: route.query.id });
    if (res.code === 0) {
      formData.value = res.data.reac;
      type.value = "update";
    }
  } else {
    type.value = "create";
  }
};
let reqData = reactive({
  ID: null,
  parameters: {},
  name: "",
  case_id: 0,
  base_url: "",
  headers: [],
  variables: "",
  extract: "",
  validate: "",
  hooks: [],
  apiMenuID: "",
  verify: false,
  default: false,
  setup_hooks: [],
  teardown_hooks: [],
  retry: 0,
});

init();
const configLabel = reactive({
  name: "",
  base_url: "",
});
// 保存按钮
const save = async () => {
  let res;
  switch (type.value) {
    case "create":
      res = await createApiConfig(formData.value);
      break;
    case "update":
      res = await updateApiConfig(formData.value);
      break;
    default:
      res = await createApiConfig(formData.value);
      break;
  }
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "创建/更改成功",
    });
  }
};

// 数据

if (eventType.value === "update") {
  reqData = props.formData.value;
  if (reqData.setup_case_id === 0) {
    reqData.setup_case_id = null;
  }
  configLabel.base_url = reqData.base_url;
  configLabel.name = reqData.name;
}
let headers = [];
let validate = [];
let requestId = [];
let requestFormData = [];
let requestParamsData = {};
let requestJsonData = {};
let requestExtractData = [];
let requestValidateData = [];
let requestVariablesData = [];
let activeTag = "Header";
let teardownHook = [];
let setupHook = [];
let parametersData = [];
const eventMsg = ref();

const handleHeader = (tableData) => {
  headers = tableData;
};

const variablesVariables = (requestVariables) => {
  requestVariablesData = requestVariables;
};

const configParameters = (ParametersJson) => {
  parametersData = ParametersJson;
};

const closeDialog = () => {
  emit("close", true);
};

const handleRequest = (id) => {
  if (id > 0) {
    requestId.push(id);
  }
};

const requestParams = (requestParams) => {
  requestParamsData = requestParams;
};

const typeTransformation = (data) => {
  let dataJson = {};
  if (data.length > 0) {
    data.forEach((item, index, arr) => {
      dataJson[item.key] = item.value;
    });
  }
  return dataJson;
};

const saves = () => {
  if (reqData.base_url === "") {
    ElMessage({
      type: "error",
      message: "域名不能为空",
    });
    return;
  }
  if (reqData.name === "") {
    ElMessage({
      type: "error",
      message: "配置名称不能为空",
    });
    return;
  }
  reqData.headers = typeTransformation(headers);
  reqData.headers_json = headers;
  reqData.setup_hooks = setupHook;
  reqData.teardown_hooks = teardownHook;
  reqData.parameters = requestParamsData;
  reqData.variables = typeTransformation(requestVariablesData);
  reqData.variables_json = requestVariablesData;
  createInterface();
};

const createInterface = async () => {
  let res;
  switch (eventType.value) {
    case "create":
      res = await createApiConfig(reqData);
      eventMsg.value = "创建";
      break;
    case "update":
      reqData.setup_case = null;
      res = await updateApiConfig(reqData);
      eventMsg.value = "修改";
      break;
    default:
      res = await createApiConfig(reqData);
      eventMsg.value = "创建";
      break;
  }
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: eventMsg.value + "成功",
    });
    emit("close");
  }
};

// 返回按钮
const back = () => {
  router.go(-1);
};
</script>

<style>
.request {
  margin-top: 15px;
  border: 1px solid #ddd;
}
</style>
