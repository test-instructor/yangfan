<template>
  <div>
    <div class="gva-form-box" style="width:90%">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="100px">
        <el-form-item label="步骤名称:" prop="name">
          <el-input v-model="formData.name" :clearable="false" placeholder="请输入步骤名称" />
        </el-form-item>

        <el-form-item>
          <div style="height: 560px;width: 1200px">
            <el-tabs v-model="activeTab">
              <el-tab-pane label="动作列表" name="actions">
                <div style="display: flex; align-items: center; gap: 10px; margin-bottom: 10px;">
                  <span style="width: 64px; color: #606266;">设备</span>
                  <el-select v-model="deviceSerial" :placeholder="devicePlaceholder" clearable filterable style="width: 420px" @change="handleDeviceChange">
                    <el-option v-for="item in deviceOptions" :key="item.ID" :label="getDeviceLabel(item)" :value="getDeviceValue(item)" />
                  </el-select>
                  <el-button type="primary" plain :disabled="!deviceSerial" @click="applyDeviceToAllActions">
                    应用到全部动作
                  </el-button>
                </div>
                <MobileActionList v-model="actions" />
              </el-tab-pane>

              <el-tab-pane label="跳过执行" name="skip">
                <Validate
                  :validate="formData?.skip_temp ?? []"
                  :heights="500"
                  idName="skip"
                  @jsonData="handleJsonDataSkip"
                />
              </el-tab-pane>
              <el-tab-pane label="变量" name="variables">
                <Variables
                  :variables="formData?.variables_temp ?? []"
                  :heights="500"
                  @variablesData="handleVariablesData"
                />
              </el-tab-pane>
              <el-tab-pane label="参数提取" name="extract">
                <Extractor
                  :extract="formData?.extract_temp ?? []"
                  :heights="500"
                  idName="extract"
                  @extractData="handleJsonDataExtract"
                />
              </el-tab-pane>
              <el-tab-pane label="勾子(hook)" name="hook">
                <Hook
                  :setupHooks="formData?.setup_hooks ?? []"
                  :teardownHooks="formData?.teardown_hooks ?? []"
                  :heights="500"
                  @setupHooksData="handleSetupHooksData"
                  @teardownHooksData="handleTeardownHooksData"
                />
              </el-tab-pane>
              <el-tab-pane label="断言" name="validate">
                <Validate
                  :validate="formData?.validators_temp ?? []"
                  :heights="500"
                  idName="validate"
                  @jsonData="handleJsonDataValidate"
                />
              </el-tab-pane>
              <el-tab-pane label="参数设置" name="parameters">
                <ParameterTable
                  :jsons="formData?.parameters ?? {}"
                  :parametersTemp="formData?.parameters_temp ?? {}"
                  @jsonData="handleJsonDataParameters"
                  @tempData="(val) => formData.parameters_temp = val"
                />
              </el-tab-pane>
              <el-tab-pane label="数据仓库" name="dataWarehouse">
                <DataWarehouseConfig
                  v-model="formData.data_warehouse_temp"
                  :hideCount="true"
                  :height="460"
                />
              </el-tab-pane>
            </el-tabs>
          </div>
        </el-form-item>

        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="cancel">取消</el-button>

          <EnvSelector />
          <PythonFuncSelector />
          <DataWarehouseFieldSelector :current-type="dataWarehouseCurrentType" />
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import { computed, onMounted, ref, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { createAutoStep, updateAutoStep } from '@/api/automation/autostep'
import MobileActionList from '@/components/platform/mobile/ActionList.vue'
import Variables from '@/components/platform/variables/index.vue'
import Extractor from '@/components/platform/extract/index.vue'
import Validate from '@/components/platform/validate/index.vue'
import Hook from '@/components/platform/hook/index.vue'
import ParameterTable from '@/components/platform/parameterTable/index.vue'
import DataWarehouseConfig from '@/components/platform/dataWarehouseConfig/index.vue'
import EnvSelector from '@/components/platform/button/EnvDetail.vue'
import PythonFuncSelector from '@/components/platform/button/PythonFuncSelector.vue'
import DataWarehouseFieldSelector from '@/components/platform/button/DataWarehouseFieldSelector.vue'
import { getAndroidDeviceOptionsList } from '@/api/platform/androidDeviceOptions'
import { getIOSDeviceOptionsList } from '@/api/platform/iosOptions'
import { getHarmonyDeviceOptionsList } from '@/api/platform/harmonyDeviceOption'
import { getBrowserDeviceOptionsList } from '@/api/platform/browserDeviceConfig'
import utils from '@/utils/dataTypeConverter.js'

defineOptions({
  name: 'AutoStepUIForm'
})

const props = defineProps({
  menu: {
    type: Number,
    default: 0
  },
  formData: {
    type: Object,
    default: () => ({})
  },
  stepType: {
    type: String,
    default: ''
  },
  platform: {
    type: String,
    default: 'android'
  }
})

const emit = defineEmits(['close'])

const btnLoading = ref(false)
const elFormRef = ref()
const activeTab = ref('actions')

const rule = {
  name: [{
    required: true,
    message: '',
    trigger: ['input', 'blur']
  }]
}

const formData = ref({
  name: '',
  loops: 0,
  retry: 0,
  variables: {},
  variables_temp: [],
  parameters: {},
  parameters_temp: {},
  data_warehouse: {},
  data_warehouse_temp: {},
  setup_hooks: [],
  teardown_hooks: [],
  extract: {},
  extract_temp: [],
  validate: [],
  validators_temp: [],
  export: [],
  skip: [],
  skip_temp: [],
  ignore_popup: false,
  menu: 0
})

const deviceOptions = ref([])
const deviceSerial = ref('')
const actions = ref([])
const variables = ref([])
const skipData = ref([])
const validateData = ref([])
const extractorData = ref([])
const setupHooksData = ref([])
const teardownHooksData = ref([])
const jsonErrorParameters = ref('')

const deepClone = (obj) => {
  try {
    return structuredClone(obj)
  } catch (e) {
    return JSON.parse(JSON.stringify(obj || {}))
  }
}

const devicePlaceholder = computed(() => {
  switch (props.platform) {
    case 'android':
      return '请选择安卓设备'
    case 'ios':
      return '请选择iOS设备'
    case 'harmony':
      return '请选择鸿蒙设备'
    case 'browser':
      return '请选择浏览器'
    default:
      return '请选择设备'
  }
})

const dataWarehouseCurrentType = computed(() => {
  return String(formData.value?.data_warehouse_temp?.type || formData.value?.data_warehouse?.type || '').trim()
})

const extractDeviceSerialFromActions = () => {
  const list = actions.value || []
  const serials = new Set(
    list
      .map(a => a?.options?.serial)
      .filter(v => typeof v === 'string' && v.trim() !== '')
  )
  if (serials.size === 1) {
    deviceSerial.value = Array.from(serials)[0]
  } else {
    deviceSerial.value = ''
  }
}

const fetchDeviceOptions = async () => {
  deviceOptions.value = []
  if (!['android', 'ios', 'harmony', 'browser'].includes(props.platform)) return
  let res
  switch (props.platform) {
    case 'android':
      res = await getAndroidDeviceOptionsList({ page: 1, pageSize: 999 })
      break
    case 'ios':
      res = await getIOSDeviceOptionsList({ page: 1, pageSize: 999 })
      break
    case 'harmony':
      res = await getHarmonyDeviceOptionsList({ page: 1, pageSize: 999 })
      break
    case 'browser':
      res = await getBrowserDeviceOptionsList({ page: 1, pageSize: 999 })
      break
  }
  if (res?.code === 0) {
    deviceOptions.value = res.data?.list || []
  }
}

const getDeviceLabel = (item) => {
  if (!item) return ''
  switch (props.platform) {
    case 'android':
      return item.name || item.serial || ''
    case 'ios':
      return item.name || item.udid || ''
    case 'harmony':
      return item.name || item.connectKey || ''
    case 'browser':
      return item.browserId || ''
    default:
      return item.name || ''
  }
}

const getDeviceValue = (item) => {
  if (!item) return ''
  switch (props.platform) {
    case 'android':
      return item.serial || ''
    case 'ios':
      return item.udid || ''
    case 'harmony':
      return item.connectKey || ''
    case 'browser':
      return item.browserId || ''
    default:
      return ''
  }
}

const applyDeviceToAllActions = () => {
  const serial = deviceSerial.value
  if (!serial) return
  const list = actions.value || []
  const next = list.map(a => {
    const options = { ...(a?.options || {}) }
    options.platform = props.platform
    options.serial = serial
    return { ...a, options }
  })
  actions.value = next
}

const handleDeviceChange = () => {
  if (!deviceSerial.value) return
  const list = actions.value || []
  const next = list.map(a => {
    const options = { ...(a?.options || {}) }
    if (!options.serial) {
      options.platform = props.platform
      options.serial = deviceSerial.value
    }
    return { ...a, options }
  })
  actions.value = next
}

const handleVariablesData = (tableData) => {
  variables.value = tableData
}
const handleJsonDataSkip = (tableData) => {
  skipData.value = tableData
}
const handleJsonDataValidate = (tableData) => {
  validateData.value = tableData
}
const handleJsonDataExtract = (tableData) => {
  extractorData.value = tableData
}
const handleSetupHooksData = (tableData) => {
  setupHooksData.value = tableData
}
const handleTeardownHooksData = (tableData) => {
  teardownHooksData.value = tableData
}
const handleJsonDataParameters = (result) => {
  if (!result.isValid) {
    jsonErrorParameters.value = result.error.message
  } else {
    jsonErrorParameters.value = ''
    formData.value.parameters = result.data
  }
}

const messageBox = (message) => {
  ElMessageBox.alert(
    message,
    '数据校验失败',
    {
      confirmButtonText: '确认',
      type: 'error',
      dangerouslyUseHTMLString: true,
      center: true
    }
  )
}

const dataValidation = async () => {
  const variablesData = utils.convertData(variables.value)
  if (!variablesData.success) {
    messageBox('数据处理失败字段:' + variablesData.errors.key + ',出现错误:' + variablesData.errors.error)
    activeTab.value = 'variables'
    return false
  }
  formData.value.variables_temp = variables.value
  formData.value.variables = variablesData.data

  const skipDataNew = utils.assertionData(skipData.value)
  if (!skipDataNew.success) {
    messageBox('数据处理失败字段:' + skipDataNew.errors.key + ',出现错误:' + skipDataNew.errors.error)
    activeTab.value = 'skip'
    return false
  }
  formData.value.skip = skipDataNew.data
  formData.value.skip_temp = skipData.value

  const validateDataNew = utils.assertionData(validateData.value)
  if (!validateDataNew.success) {
    messageBox('数据处理失败字段:' + validateDataNew.errors.key + ',出现错误:' + validateDataNew.errors.error)
    activeTab.value = 'validate'
    return false
  }
  formData.value.validate = validateDataNew.data
  formData.value.validators_temp = validateData.value

  const extractDataNew = utils.processData(extractorData.value)
  if (!extractDataNew.success) {
    messageBox('数据处理失败:' + extractDataNew.errors.join('、') + ',字段重复')
    activeTab.value = 'extract'
    return false
  }
  formData.value.extract = extractDataNew.data
  formData.value.extract_temp = extractorData.value

  if (jsonErrorParameters.value !== '') {
    messageBox('参数设置数据处理失败:' + jsonErrorParameters.value)
    activeTab.value = 'parameters'
    return false
  }

  formData.value.setup_hooks = setupHooksData.value
  formData.value.teardown_hooks = teardownHooksData.value

  if (formData.value.data_warehouse_temp && Object.keys(formData.value.data_warehouse_temp).length > 0) {
    formData.value.data_warehouse = formData.value.data_warehouse_temp
  } else {
    formData.value.data_warehouse = {}
  }

  const dwConfig = formData.value.data_warehouse_temp
  if (dwConfig && dwConfig.filter && dwConfig.filter.groups) {
    for (const group of dwConfig.filter.groups) {
      if (group.conditions) {
        for (const cond of group.conditions) {
          if (!cond.field || String(cond.field).trim() === '') {
            messageBox('数据仓库配置中有未选择字段的筛选条件')
            activeTab.value = 'dataWarehouse'
            return false
          }
        }
      }
    }
  }

  return true
}

const normalizePayload = () => {
  const payload = { ...formData.value }
  payload.menu = Number(props.menu)

  const knownPlatforms = ['android', 'ios', 'harmony', 'browser']
  for (const p of knownPlatforms) {
    if (p !== props.platform) {
      delete payload[p]
    }
  }
  payload.request = null
  payload.request_id = 0
  payload.type = 1
  payload[props.platform] = { actions: actions.value || [] }
  return payload
}

const save = async () => {
  const validation = await dataValidation()
  if (!validation) return
  btnLoading.value = true
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return btnLoading.value = false
    const payload = normalizePayload()
    let res
    switch (props.stepType) {
      case 'create':
        res = await createAutoStep(payload)
        break
      case 'update':
        res = await updateAutoStep(payload)
        break
      case 'copy':
        res = await createAutoStep(payload)
        break
      default:
        res = await createAutoStep(payload)
        break
    }
    btnLoading.value = false
    if (res?.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
      emit('close')
    }
  })
}

const cancel = () => {
  const close = ref(true)
  if (props.stepType === 'update') close.value = false
  emit('close', { value: close })
}

const syncFromProps = async () => {
  const source = deepClone(props.formData || {})
  formData.value = {
    ...formData.value,
    ...source
  }

  if (!formData.value.data_warehouse) formData.value.data_warehouse = {}
  if (!formData.value.data_warehouse_temp || Object.keys(formData.value.data_warehouse_temp || {}).length === 0) {
    formData.value.data_warehouse_temp = formData.value.data_warehouse || {}
  }

  if (!formData.value[props.platform]) formData.value[props.platform] = { actions: [] }
  actions.value = Array.isArray(formData.value?.[props.platform]?.actions) ? deepClone(formData.value[props.platform].actions) : []

  variables.value = formData.value.variables_temp || []
  skipData.value = formData.value.skip_temp || []
  validateData.value = formData.value.validators_temp || []
  extractorData.value = formData.value.extract_temp || []
  setupHooksData.value = formData.value.setup_hooks || []
  teardownHooksData.value = formData.value.teardown_hooks || []

  await fetchDeviceOptions()
  extractDeviceSerialFromActions()
}

watch(() => [props.formData, props.platform], () => {
  syncFromProps()
}, { immediate: true })

watch(actions, (val) => {
  if (!formData.value[props.platform]) formData.value[props.platform] = { actions: [] }
  formData.value[props.platform].actions = val || []
}, { deep: true })

onMounted(async () => {
  await fetchDeviceOptions()
  extractDeviceSerialFromActions()
})
</script>

<style scoped>
</style>
