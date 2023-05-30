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
    <div style="margin-top: 10px">
      <el-autocomplete
        v-model="formGrpc.host"
        placeholder="请输入服务器地址"
        clearable
        :fetch-suggestions="querySearch"
        @select="getGrpcHost"
        @change="getGrpcHost"
        style="width: 400px; margin-right: 10px"
      >
        <template #prepend>服务器信息</template>
      </el-autocomplete>
      <el-select
        v-model="formGrpc.server"
        class="m-2"
        placeholder="Choose Service"
        size="small"
        style="margin-right: 10px"
      >
        <el-option
          v-for="item in serverOption"
          :key="item"
          :label="item"
          :value="item"
          @focus="getGrpc"
          @click.native="getGrpc"
        />
      </el-select>
      <el-select
        v-model="formGrpc.method"
        class="m-2"
        placeholder="Choose Method"
        size="small"
        style="margin-right: 10px"
      >
        <el-option
          v-for="item in methodOption"
          :key="item"
          :label="item"
          :value="item"
          @focus="getGrpc"
          @click.native="getGrpc"
        />
      </el-select>
      <el-button @click="openReqDetail">request 详情</el-button>
      <el-button @click="resetRef" icon="refresh" circle></el-button>
    </div>
    <div>
      <el-input
        style="width: 800px; margin-top: 10px"
        placeholder="请输入接口请求地址"
        v-model="formLabelAlign.url"
        clearable
        disabled
        class="input-with-select"
      >
        <template #prepend>url</template>
      </el-input>
    </div>
    <div class="request">
      <el-tabs style="margin-left: 20px" v-model="activeTag">
        <el-tab-pane label="Header" name="Header">
          <headers
            @headerData="handleHeader"
            @exportHeader="handleExportHeader"
            @request="handleRequest"
            :header="reqData ? reqData.gRPC.headers_json : []"
            :exportHeader="reqData ? reqData.export_header : []"
            :heights="heightDiv"
          >
          </headers>
        </el-tab-pane>

        <el-tab-pane label="Body" name="jsons">
          <div style="display: flex">
            <div class="lefts">
              <div style="margin-bottom: 5px">
                <span>请求参数</span>
              </div>
              <jsons
                @requestJsonData="requestJson"
                :heights="heightDiv - 30"
                :jsons="reqData ? reqData.gRPC.body : ''"
              >
              </jsons>
            </div>
            <div class="rights">
              <div style="margin-bottom: 5px">
                <span>服务器反射数据</span>
              </div>

              <jsonCompare
                @requestJsonData="requestJson"
                :heights="heightDiv - 30"
                :jsons="reqGrpc ? reqGrpc.body : ''"
                :key="reqGrpc.body"
              >
              </jsonCompare>
            </div>
          </div>
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
  <el-dialog
    v-model="reqDetail"
    :close-on-click-modal="false"
    :close-on-press-escape="false"
    title="request 详情"
    :before-close="closeReqDetail"
    width="90%"
    top="100px"
  >
    <el-tabs style="margin-left: 20px" v-model="reqGrpcType">
      <el-tab-pane v-if="reqGrpc.enum" label="enum" name="enum">
        <el-tabs style="margin-left: 20px" v-model="enumName">
          <el-tab-pane :label="e" :name="e" v-for="e in enumNames">
            <el-table :data="reqGrpc.enum[e]">
              <el-table-column prop="name" label="Name" width="280" />
              <el-table-column prop="num" label="Num" width="120" />
            </el-table>
          </el-tab-pane>
        </el-tabs>
      </el-tab-pane>

      <el-tab-pane :label="item" :name="item" v-for="item in grpcMessageReqKey">
        <el-table :data="reqGrpc.message[item]">
          <el-table-column prop="name" label="Name" min-width="50" />
          <el-table-column prop="protoName" label="ProtoName" min-width="50" />
          <el-table-column prop="type" label="Type" min-width="50" />
          <el-table-column label="Message" width="85">
            <template #default="scope">
              <span v-if="!scope.row.isMessage">{{ scope.row.isMessage }}</span>
              <el-button
                v-if="scope.row.isMessage"
                @click="clickMessage(scope.row)"
                link
                type="primary"
              >
                {{ "详情" }}
              </el-button>
            </template>
          </el-table-column>
          <el-table-column label="Enum" min-width="50">
            <template #default="scope">
              <span v-if="!scope.row.isEnum">{{ scope.row.isEnum }}</span>
              <el-button
                v-if="scope.row.isEnum"
                @click="clickEnum(scope.row)"
                link
                type="primary"
              >
                {{ "详情" }}
              </el-button>
            </template>
          </el-table-column>
          <el-table-column prop="isArray" label="Array" min-width="50" />
          <el-table-column prop="isMap" label="Map" min-width="50">
            <template #default="scope">
              <span v-if="!scope.row.isMap">{{ scope.row.isMap }}</span>
              <el-button
                v-if="scope.row.isMap"
                @click="clickMessage(scope.row)"
                link
                type="primary"
              >
                {{ "详情" }}
              </el-button>
            </template>
          </el-table-column>
          <el-table-column prop="isRequired" label="Required" min-width="50" />
          <el-table-column
            prop="defaultVal"
            label="DefaultVal"
            min-width="50"
          />
          <el-table-column
            prop="description"
            label="Description"
            min-width="50"
          />
        </el-table>
      </el-tab-pane>
    </el-tabs>
  </el-dialog>
</template>

<script setup>
import { reactive, ref, defineEmits } from "vue";

import Headers from "@/view/interface/interfaceComponents/Headers.vue";
import jsonCompare from "@/view/interface/interfaceComponents/jsonCompare.vue";
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
  getGrpcFunc,
  getUserConfig,
} from "@/api/interfaceTemplate";
import { ElMessage } from "element-plus";
import { getDict } from "@/utils/dictionary";
import Jsons from "@/view/interface/interfaceComponents/Jsons.vue";
import UserConfig from "@/view/interface/interfaceComponents/userConfig.vue";

const emit = defineEmits(["close"]);
const props = defineProps({
  cid: ref(),
  heights: ref(),
  eventType: ref(),
  apiType: ref(),
  formData: ref({
    name: "",
    variables: "",
    extract: "",
    validate: "",
    hooks: "",
    apiMenuID: "",
  }),
});

const serverOption = ref([]);
const methodOption = ref([]);
const heightDiv = ref(0);
const eventType = ref("");
const formData = reactive({});
let configId;
heightDiv.value = props.heights;
eventType.value = props.eventType;
configId = props.cid;
let type = ref();
const reqDetail = ref(false);
const formLabelAlign = reactive({
  url: "",
  headers: "",
  headers_json: "",
  body: "",
  timeout: 0,
  detail: {},
});

const formGrpc = ref({
  host: "",
  server: "",
  method: "",
  ref: false,
});

const openReqDetail = () => {
  reqDetail.value = true;
  showJson.value = true;
};

const resetRef = async () => {
  if (formGrpc.value.host === "") {
    ElMessage({
      type: "error",
      message: "请输入要重置的服务",
    });
    return;
  }
  let reRef = JSON.parse(JSON.stringify(formGrpc.value));
  reRef.ref = true;
  let res = await getGrpcFunc(reRef);
  if (res.code === 0) {
    ElMessage.success("服务" + formGrpc.value.host + "重置成功");
    await getGrpc();
  }
};

const closeReqDetail = () => {
  reqDetail.value = false;
};

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

const getGrpcHost = () => {
  formGrpc.value.server = "";
  formGrpc.value.method = "";
  getGrpc();
};

const getGrpc = async () => {
  setHosts(formGrpc.value.host);
  if (formGrpc.value.host === "") {
    return;
  }
  if (formGrpc.value.method !== "" && formGrpc.value.server !== "") {
    formLabelAlign.url = formGrpc.value.server + "." + formGrpc.value.method;
  }
  let res = await getGrpcFunc(formGrpc.value);
  if (res.code === 0) {
    serverOption.value = res.data.servers;
    if (serverOption.value.indexOf(formGrpc.value.server) < 0) {
      formGrpc.value.server = "";
    }
    methodOption.value = res.data.methods;
    if (
      formGrpc.value.method !== "" &&
      methodOption.value.indexOf(formGrpc.value.method) < 0
    ) {
      formGrpc.value.method = "";
    }
    if (res.data.request) {
      getReqBody(res.data.request);
    }
  }
};

const reqGrpc = ref({
  body: {},
  message: {},
  type: "",
  enum: {},
});
const enumNames = ref([]);
const enumName = ref([]);
const reqGrpcType = ref("");
const showJson = ref(true);
const grpcMessageReqKey = ref([]);
const getReqBody = (req) => {
  grpcMessageReqKey.value = [];
  enumNames.value = [];
  reqGrpc.value.body = JSON.parse(req.body);
  grpcMessageReqKey.value.push(req.type);
  for (let m in req.message) {
    if (m !== req.type) {
      grpcMessageReqKey.value.push(m);
    }
  }
  for (let e in req.enum) {
    enumNames.value.push(e);
    enumName.value = enumNames.value[0];
  }
  reqGrpc.value.message = req.message;
  reqGrpc.value.type = req.type;
  reqGrpc.value.enum = req.enum;

  reqGrpcType.value = req.type;

  // reqData.gRPC.body = reqGrpc.value.body
};

const clickEnum = (row) => {
  enumName.value = row.type;
  reqGrpcType.value = "enum";
};

const clickMessage = (row) => {
  reqGrpcType.value = row.type;
};

const hostOptions = ref([]);
const querySearch = (queryString, cb) => {
  let hostOption = hostOptions.value;
  const results = queryString
    ? hostOption.filter(createFilter(queryString))
    : hostOption;
  cb(results);
};

const createFilter = (queryString) => {
  return (hostOptions) => {
    return (
      hostOptions.value.toLowerCase().indexOf(queryString.toLowerCase()) === 0
    );
  };
};

const getHosts = () => {
  let hosts = JSON.parse(window.localStorage.getItem("hosts"));
  if (hosts) {
    hosts.reverse();
  }
  hostOptions.value = hosts;
  hostOptions.value = [];
  hosts &&
    hosts.forEach((item) => {
      let hostOption = {};
      hostOption.value = item;
      hostOptions.value.push(hostOption);
    });
};
getHosts();

const setHosts = (host) => {
  let hosts = JSON.parse(window.localStorage.getItem("hosts"));
  if (!hosts) {
    hosts = [];
  }

  if (hosts.indexOf(host) >= 0) {
    hosts.splice(hosts.indexOf(host), 1);
  }
  hosts.push(host);
  window.localStorage.setItem("hosts", JSON.stringify(hosts));
  hosts.reverse();
  hostOptions.value = hosts;
  hostOptions.value = [];
  hosts &&
    hosts.forEach((item) => {
      let hostOption = {};
      hostOption.value = item;
      hostOptions.value.push(hostOption);
    });
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
  gRPC: reactive({
    id: 0,
    url: "",
    headers: "",
    headers_json: "",
    body: "",
    timeout: 0,
    type: "Simple",
    detail: {},
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
  formLabelAlign.url = reqData.gRPC.url;
  formLabelAlign.name = reqData.name;
  formLabelAlign.timeout = reqData.gRPC.timeout;
  if (reqData.gRPC.detail) {
    formGrpc.value = reqData.gRPC.detail;
  }
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

  setReqData();
  if (formLabelAlign.name === "") {
    ElMessage({
      type: "error",
      message: "接口名称不能为空",
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
  reqData.gRPC.type = "Simple";
  reqData.gRPC.body = requestJsonData;
  reqData.gRPC.url = formLabelAlign.url;
  reqData.gRPC.headers = typeTransformation(headers);
  reqData.gRPC.headers_json = headers;
  reqData.gRPC.timeout = 0;
  reqData.gRPC.detail = formGrpc.value;
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
  height: 520px;
}
.lefts {
  float: left;
  width: 48%;
}
.rights {
  float: right;
  width: 49%;
  margin-left: 2%;
}
</style>
