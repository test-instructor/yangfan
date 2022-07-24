<template>
  <div>
    <div>
      <div style="display: flex;">
        <div style="width: 120px; margin-top: 5px; display: table; margin-left:20px; ">
          <span>verify: </span>
          <el-switch
              v-model="formLabelAlign.verify"
              class="mb-2"
              size="large"
              name="verify"
          />
        </div>
        <div style="width: 150px; margin-top: 5px; display: table;  ">
          <span style="margin-right:5px;">允许重定向:  </span>
          <el-switch
              v-model="formLabelAlign.allow_redirects"
              class="mb-2"
              size="large"
              name="allowRedirects"
          />
        </div>
        <div style="width: 150px; margin-top: 5px; display: table; margin-left:20px; ">
          <span style="margin-right:5px;">http2:  </span>
          <el-switch
              v-model="formLabelAlign.http2"
              class="mb-2"
              size="large"
              name="http2"
          />
        </div>

        <div>
          <span style="margin-right:5px;">超时时间:  </span>
          <el-input-number
              v-model="formLabelAlign.timeout"
              :precision="2"
              :step="0.1"
              :max="3600"
              size="small"
              controls-position="right"
              :min="0"
          />
        </div>

      </div>
    </div>
    <div style="width: 800px; margin-top: 10px">
      <el-input
          v-model="formLabelAlign.name"
          placeholder="请输入接口名称">
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
              style="width: 125px"
              v-model="formLabelAlign.method"
              placeholder="Select"
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
      <el-tabs
          style="margin-left: 20px;"
          v-model="activeTag"
      >
        <el-tab-pane label="Header" name="Header">
          <headers
              @headerData="handleHeader"
              @request="handleRequest"
              :header="reqData ? reqData.request.headers_json : []"
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
    <br/>
    <el-button type="primary" @click="saves">保存</el-button>
    <el-button type="info" @click="closeDialog">取消</el-button>
  </div>
</template>

<script setup>


import {reactive, ref, defineEmits} from 'vue'

import Headers from '@/view/interface/interfaceComponents/Headers.vue'
import Forms from '@/view/interface/interfaceComponents/form.vue'
import Params from '@/view/interface/interfaceComponents/Params.vue'
import Jsons from '@/view/interface/interfaceComponents/Jsons.vue'
import Extract from '@/view/interface/interfaceComponents/Extract.vue'
import Validate from '@/view/interface/interfaceComponents/Validate.vue'
import Variables from '@/view/interface/interfaceComponents/Variables.vue'
import Hooks from '@/view/interface/interfaceComponents/Hooks.vue'

import {
  createInterfaceTemplate,
  updateInterfaceTemplate,
  findInterfaceTemplate
} from '@/api/interfaceTemplate'
import {ElMessage} from "element-plus";
import {getDict} from "@/utils/dictionary";

const emit = defineEmits(["close"]);
const props = defineProps({
  heights: ref(),
  eventType: ref(),
  apiType: ref(),
  formData: ref({
    name: '',
    request: ref({
      agreement: '',
      method: '',
      url: '',
      params: '',
      headers: '',
      json: '',
      data: '',
    }),
    variables: '',
    extract: '',
    validate: '',
    hooks: '',
    apiMenuID: '',
  }),

})

const heightDiv = ref(false)
const eventType = ref('')
const formData = reactive({})
heightDiv.value = props.heights
eventType.value = props.eventType

let type = ref()
const formLabelAlign = reactive({
  name: '',
  region: '',
  type: '',
  url: '',
  method: '',
  timeout: 0,
  http2: false,
  allow_redirects: false,
  verify: false,
})
// eslint-disable-next-line no-unused-vars
// const httpOptions = await getDict('method')

const httpOptions = []

const httpOption = async () => {
  const res = await getDict('method')
  res && res.forEach(item => {
    httpOptions.push(item)
  })
}
httpOption()

let headers = []
let forms = []
let requestId = []
let validateId = []
let validate = []
let teardownHook = []
let setupHook = []
let requestFormData = []
let requestParamsData = []
let requestJsonData = {}
let requestFtData = []
let requestExtractData = []
let requestValidateData = []
let requestVariablesData = []
let activeTag = 'Header'
const eventMsg = ref()

const handleRequest = (id) => {
  if (id > 0) {
    requestId.push(id)
  }
}

const handleValidate = (id) => {
  if (id > 0) {
    validateId.push(id)
  }
}

let reqData = reactive({
  ID: 0,
  name: '',
  type: 0,
  request: reactive({
    id: 0,
    agreement: '',
    method: '',
    url: '',
    params: {},
    params_json: '',
    headers: '',
    headers_json: '',
    json: '',
    data: '',
    data_json: '',
    allow_redirects: false,
    http2: false,
    verify: false,
    timeout: 0,
  }),
  variables: '',
  variables_json: [],
  extract: '',
  extract_json: '',
  validate: '',
  hooks: '',
  apiMenuID: '',
  setup_hooks: [],
  teardown_hooks: [],
})
if (eventType.value === "update") {
  reqData = props.formData.value
  formLabelAlign.method = reqData.request.method
  formLabelAlign.url = reqData.request.url
  formLabelAlign.name = reqData.name
  formLabelAlign.timeout = reqData.request.timeout
  formLabelAlign.verify = reqData.request.verify
  formLabelAlign.allow_redirects = reqData.request.allow_redirects
  formLabelAlign.http2 = reqData.request.http2
}


const closeDialog = () => {
  emit("close", true)
}

const params = reactive({
  menu: '',
})

const typeTransformation = (data) => {
  let dataJson = {}
  if (data.length > 0) {
    data.forEach((item, index, arr) => {
      dataJson[item.key] = item.value
    })
  }
  return dataJson
}

const saves = () => {
  if (formLabelAlign.name === "") {

  }
  if (formLabelAlign.method === "") {

  }
  reqData.type = props.apiType
  reqData.request.http2 = formLabelAlign.http2
  reqData.request.verify = formLabelAlign.verify
  reqData.request.allow_redirects = formLabelAlign.allow_redirects
  reqData.request.timeout = formLabelAlign.timeout
  reqData.request.url = formLabelAlign.url
  reqData.request.method = formLabelAlign.method
  reqData.request.headers = typeTransformation(headers)
  reqData.request.headers_json = headers
  reqData.request.data = typeTransformation(requestFormData)
  reqData.request.data_json = requestFormData
  reqData.request.params = typeTransformation(requestParamsData)
  reqData.request.params_json = requestParamsData
  reqData.request.json = requestJsonData
  reqData.name = formLabelAlign.name
  reqData.extract = typeTransformation(requestExtractData)
  reqData.extract_json = requestExtractData
  reqData.validate = requestValidateData
  reqData.variables = typeTransformation(requestVariablesData)
  reqData.variables_json = requestVariablesData
  reqData.teardown_hooks = teardownHook
  reqData.setup_hooks = setupHook
  params.menu = window.localStorage.getItem('menu')
  createInterface()
}

const createInterface = async () => {
  let res
  switch (eventType.value) {
    case 'create':
      res = await createInterfaceTemplate(reqData, params)
      eventMsg.value = "创建"
      break
    case 'update':
      res = await updateInterfaceTemplate(reqData, params)
      eventMsg.value = "修改"
      break
    default:
      res = await createInterfaceTemplate(reqData, params)
      eventMsg.value = "创建"
      break
  }
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: eventMsg.value + '成功'
    })
    emit("close")
  }
}

const handleHeader = (tableData) => {
  headers = tableData;
}
const requestForm = (requestForms) => {
  requestFormData = requestForms;
}
const requestJson = (requestJsons) => {
  requestJsonData = requestJsons;
}
const requestExtract = (requestExtract) => {
  requestExtractData = requestExtract;
}
const requestValidateDate = (requestValidate) => {
  requestValidateData = requestValidate;
}
const requestVariables = (requestVariables) => {
  requestVariablesData = requestVariables;
}
const requestParams = (requestParams) => {
  requestParamsData = requestParams;
}
const teardownHooks = (hooks) => {
  teardownHook = hooks
}

const setupHooks = (hooks) => {
  setupHook = hooks
}

</script>

<style>
.request {
  margin-top: 15px;
  border: 1px solid #ddd;
  height: 600px;
}
</style>
