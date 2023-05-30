<template>
  <div>
    <div>
      <user-config />
    </div>
    <div style="width: 800px">
      <el-input v-model="formLabelAlign.name" placeholder="请输入接口名称">
        <template #prepend>接口名称</template>
      </el-input>
    </div>
    <div>
      <el-input
        style="width: 800px; margin-top: 10px"
        placeholder="请输入接口请求地址"
        v-model="formLabelAlign.url"
        clearable
        class="input-with-select"
      >
        <template #prepend>
          <el-select
            style="width: 150px"
            v-model="formLabelAlign.method"
            placeholder="请选择请求方法"
          >
            <el-option
              v-for="item in httpOptions"
              :key="item.label"
              :label="item.label"
              :value="item.label"
            >
            </el-option>
          </el-select>
        </template>
      </el-input>
    </div>
    <div class="request">
      <el-tabs style="margin-left: 20px" v-model="activeTag">
        <el-tab-pane label="Header" name="Header">
          <headers
            @headerData="handleHeader"
            @exportHeader="handleExportHeader"
            @request="handleRequest"
            :header="reqData ? reqData.request.headers_json : []"
            :exportHeader="reqData ? reqData.export_header : []"
            :heights="heightDiv"
          >
          </headers>
        </el-tab-pane>

        <el-tab-pane label="Params" name="params">
          <params
            @requestParamsData="requestParams"
            @request="handleRequest"
            :params="reqData ? reqData.request.params_json : []"
            :heights="heightDiv"
          >
          </params>
        </el-tab-pane>

        <el-tab-pane label="Form" name="forms">
          <forms
            @requestFormData="requestForm"
            @request="handleRequest"
            :forms="reqData ? reqData.request.data_json : []"
            :heights="heightDiv"
          >
          </forms>
        </el-tab-pane>

        <el-tab-pane label="Jsons" name="jsons">
          <jsons
            @requestJsonData="requestJson"
            :heights="heightDiv"
            :jsons="reqData ? reqData.request.json : ''"
          >
          </jsons>
        </el-tab-pane>

        <el-tab-pane label="Extract" name="Extract">
          <extract
            @requestExtractData="requestExtract"
            @exportParameter="handleExportParameter"
            :exportParameter="reqData ? reqData.export_parameter : []"
            :extract="reqData ? reqData.extract_json : []"
            :heights="heightDiv"
          >
          </extract>
        </el-tab-pane>

        <el-tab-pane label="Validate" name="Validate">
          <validate
            @requestValidateData="requestValidateDate"
            @validates="handleValidate"
            :validate="reqData ? reqData.validate : []"
            :heights="heightDiv"
          >
          </validate>
        </el-tab-pane>

        <el-tab-pane label="Variables" name="Variables">
          <variables
            :heights="heightDiv"
            @requestVariablesData="requestVariables"
            @request="handleRequest"
            :variables="reqData ? reqData.variables_json : []"
          >
          </variables>
        </el-tab-pane>

        <el-tab-pane label="Hooks" name="Hooks">
          <hooks
            :heights="heightDiv"
            @teardownHooksData="teardownHooks"
            @setupHooksData="setupHooks"
            :setupHooks="reqData ? reqData.setup_hooks : []"
            :teardownHooks="reqData ? reqData.teardown_hooks : []"
          >
          </hooks>
        </el-tab-pane>
      </el-tabs>
    </div>
    <br />
    <el-button type="primary" @click="saveRun">保存并调试</el-button>
    <el-button type="primary" @click="saves">保存</el-button>
    <el-button type="info" @click="closeDialog">取消</el-button>
  </div>
</template>

<script setup>
import { reactive, ref, defineEmits } from "vue";

import Headers from "@/view/interface/interfaceComponents/Headers.vue";
import Forms from "@/view/interface/interfaceComponents/form.vue";
import Params from "@/view/interface/interfaceComponents/Params.vue";
import Jsons from "@/view/interface/interfaceComponents/Jsons.vue";
import Extract from "@/view/interface/interfaceComponents/Extract.vue";
import Validate from "@/view/interface/interfaceComponents/Validate.vue";
import Variables from "@/view/interface/interfaceComponents/Variables.vue";
import Hooks from "@/view/interface/interfaceComponents/Hooks.vue";
import { runApi } from "@/api/runTestCase";
import { useRouter } from "vue-router";
const router = useRouter();

import {
  createInterfaceTemplate,
  updateInterfaceTemplate,
  findInterfaceTemplate,
  getUserConfig,
} from "@/api/interfaceTemplate";
import { ElMessage } from "element-plus";
import { getDict } from "@/utils/dictionary";
import UserConfig from "@/view/interface/interfaceComponents/userConfig.vue";

const emit = defineEmits(["close"]);
const props = defineProps({
  cid: ref(),
  heights: ref(0),
  eventType: ref(),
  apiType: ref(),
  formData: ref({
    name: "",
    request: ref({
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
  }),
});

const heightDiv = ref(0);
const eventType = ref("");
const formData = reactive({});
let configId;
heightDiv.value = props.heights;
eventType.value = props.eventType;
configId = props.cid;
let type = ref();
const formLabelAlign = reactive({
  name: "",
  region: "",
  type: "",
  url: "",
  method: "",
  timeout: 0,
  http2: false,
  allow_redirects: false,
  verify: false,
});
// eslint-disable-next-line no-unused-vars
// const httpOptions = await getDict('method')

const httpOptions = ref([]);

const httpOption = async () => {
  const res = await getDict("method");
  res &&
    res.forEach((item) => {
      httpOptions.value.push(item);
    });
};
httpOption();

let headers = [];
let export_header = [];
let export_parameter = [];
let requestId = [];
let validateId = [];
let validate = [];
let teardownHook = [];
let setupHook = [];
let requestFormData = [];
let requestParamsData = [];
let requestJsonData = {};
let requestFtData = [];
let requestExtractData = [];
let requestValidateData = [];
let requestVariablesData = [];
let activeTag = "Header";
const eventMsg = ref();
const handleRequest = (id) => {
  if (id > 0) {
    requestId.push(id);
  }
};

const handleValidate = (id) => {
  if (id > 0) {
    validateId.push(id);
  }
};

let reqData = reactive({
  ID: 0,
  name: "",
  type: 0,
  export_header: [],
  export_parameter: [],
  request: reactive({
    id: 0,
    agreement: "",
    method: "",
    url: "",
    params: {},
    params_json: "",
    headers: "",
    headers_json: "",
    json: "",
    data: "",
    data_json: "",
    allow_redirects: false,
    http2: false,
    verify: false,
    timeout: 0,
  }),
  variables: "",
  variables_json: [],
  extract: "",
  extract_json: "",
  validate: "",
  hooks: "",
  apiMenuID: "",
  setup_hooks: [],
  teardown_hooks: [],
});
if (eventType.value === "update") {
  reqData = props.formData.value;
  formLabelAlign.method = reqData.request.method;
  formLabelAlign.url = reqData.request.url;
  formLabelAlign.name = reqData.name;
  formLabelAlign.timeout = reqData.request.timeout;
  formLabelAlign.verify = reqData.request.verify;
  formLabelAlign.allow_redirects = reqData.request.allow_redirects;
  formLabelAlign.http2 = reqData.request.http2;
}

const closeDialog = () => {
  emit("close", true);
};

const params = reactive({
  menu: "",
});

const typeTransformation = (data) => {
  let dataJson = {};
  if (data.length > 0) {
    data.forEach((item, index, arr) => {
      dataJson[item.key] = item.value;
    });
  }
  return dataJson;
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

const saveRun = async () => {
  let res;
  let res1;
  setReqData();
  if (formLabelAlign.name === "") {
    ElMessage({
      type: "error",
      message: "接口名称不能为空",
    });
    return;
  }
  if (formLabelAlign.method === "") {
    ElMessage({
      type: "error",
      message: "请求方法不能为空",
    });
    return;
  }
  res = await createInterface(false);
  if (res.code === 0) {
    await runInterfaceTemplateFunc(res.data.id);
  }
};

const runInterfaceTemplateFunc = async (id) => {
  if (
    !userConfigs.value ||
    !userConfigs.value.api_config_id ||
    userConfigs.value.api_config_id < 1
  ) {
    ElMessage({
      type: "error",
      message: "请选择配置后再运行",
    });
    return;
  }
  let data = {
    caseID: id,
    configID: userConfigs.value.api_config_id,
    run_type: 5,
  };
  if (
    userConfigs.value &&
    userConfigs.value.api_env_id &&
    userConfigs.value.api_env_id > 0
  ) {
    data["env"] = userConfigs.value.api_env_id;
  }
  const res = await runApi(data);
  if (res.code === 0) {
    reportDetailFunc(res.data.id);
  }
};

const setReqData = () => {
  reqData.type = props.apiType;
  reqData.request.http2 = formLabelAlign.http2;
  reqData.request.verify = formLabelAlign.verify;
  reqData.request.allow_redirects = formLabelAlign.allow_redirects;
  reqData.request.timeout = formLabelAlign.timeout;
  reqData.request.url = formLabelAlign.url;
  reqData.request.method = formLabelAlign.method;
  reqData.request.headers = typeTransformation(headers);
  reqData.request.headers_json = headers;
  reqData.request.data = typeTransformation(requestFormData);
  reqData.request.data_json = requestFormData;
  reqData.request.params = typeTransformation(requestParamsData);
  reqData.request.params_json = requestParamsData;
  reqData.request.json = requestJsonData;
  reqData.name = formLabelAlign.name;
  reqData.extract = typeTransformation(requestExtractData);
  reqData.extract_json = requestExtractData;
  reqData.validate = requestValidateData;
  reqData.variables = typeTransformation(requestVariablesData);
  reqData.variables_json = requestVariablesData;
  reqData.teardown_hooks = teardownHook;
  reqData.setup_hooks = setupHook;
  params.menu = window.localStorage.getItem("menu");
  reqData.export_header = export_header;
  reqData.export_parameter = export_parameter;
};

const saves = () => {
  setReqData();
  createInterface(true);
};

const createInterface = async (close) => {
  if (formLabelAlign.name === "") {
    ElMessage({
      type: "error",
      message: "接口名称不能为空",
    });
    return;
  }
  if (formLabelAlign.method === "") {
    ElMessage({
      type: "error",
      message: "请求方法不能为空",
    });
    return;
  }
  let res;
  switch (eventType.value) {
    case "create":
      res = await createInterfaceTemplate(reqData, params);
      eventMsg.value = "创建";
      break;
    case "update":
      res = await updateInterfaceTemplate(reqData, params);
      eventMsg.value = "修改";
      break;
    default:
      res = await createInterfaceTemplate(reqData, params);
      eventMsg.value = "创建";
      break;
  }
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: eventMsg.value + "成功",
    });
    if (close) {
      emit("close");
    }
    return res;
  }
};

const handleHeader = (tableData) => {
  headers = tableData;
};
const handleExportHeader = (tableData) => {
  export_header = tableData;
};
const handleExportParameter = (tableData) => {
  export_parameter = tableData;
};
const requestForm = (requestForms) => {
  requestFormData = requestForms;
};
const requestJson = (requestJsons) => {
  requestJsonData = requestJsons;
};
const requestExtract = (requestExtract) => {
  requestExtractData = requestExtract;
};
const requestValidateDate = (requestValidate) => {
  requestValidateData = requestValidate;
};
const requestVariables = (requestVariables) => {
  requestVariablesData = requestVariables;
};
const requestParams = (requestParams) => {
  requestParamsData = requestParams;
};
const teardownHooks = (hooks) => {
  teardownHook = hooks;
};

const setupHooks = (hooks) => {
  setupHook = hooks;
};

const userConfigs = ref({});
const getUserConfigs = async () => {
  let res = await getUserConfig();
  if (res.code === 0 && res.data) {
    userConfigs.value = res.data;
  }
};
getUserConfigs();
</script>

<style>
.request {
  margin-top: 15px;
  border: 1px solid #ddd;
  height: 600px;
}
</style>
