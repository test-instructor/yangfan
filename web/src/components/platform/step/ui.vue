<template>
  <div>
    <div class="gva-form-box" style="width:90%">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="100px">
        <el-form-item label="步骤名称:" prop="name">
          <el-input v-model="formData.name" :clearable="false" placeholder="请输入步骤名称" />
        </el-form-item>

        <el-form-item label="设备:" class="form-item-width">
          <el-select v-model="deviceSerial" :placeholder="devicePlaceholder" clearable filterable style="width: 420px" @change="handleDeviceChange">
            <el-option v-for="item in deviceOptions" :key="item.ID" :label="getDeviceLabel(item)" :value="getDeviceValue(item)" />
          </el-select>
          <el-button type="primary" plain style="margin-left: 10px" :disabled="!deviceSerial" @click="applyDeviceToAllActions">
            应用到全部动作
          </el-button>
        </el-form-item>

        <el-form-item label="动作列表:">
          <MobileActionList v-model="currentActions" />
        </el-form-item>

        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="cancel">取消</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import { computed, onMounted, ref, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { createAutoStep, updateAutoStep } from '@/api/automation/autostep'
import MobileActionList from '@/components/platform/mobile/ActionList.vue'
import { getAndroidDeviceOptionsList } from '@/api/platform/androidDeviceOptions'
import { getIOSDeviceOptionsList } from '@/api/platform/iosOptions'
import { getHarmonyDeviceOptionsList } from '@/api/platform/harmonyDeviceOption'
import { getBrowserDeviceOptionsList } from '@/api/platform/browserDeviceConfig'

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

const ensurePlatformContainer = () => {
  if (!formData.value[props.platform]) {
    formData.value[props.platform] = { actions: [] }
  }
  if (!Array.isArray(formData.value[props.platform].actions)) {
    formData.value[props.platform].actions = []
  }
}

const currentActions = computed({
  get() {
    ensurePlatformContainer()
    return formData.value[props.platform].actions
  },
  set(val) {
    ensurePlatformContainer()
    formData.value[props.platform].actions = val || []
  }
})

const extractDeviceSerialFromActions = () => {
  const actions = currentActions.value || []
  const serials = new Set(
    actions
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
  const actions = currentActions.value || []
  const next = actions.map(a => {
    const options = { ...(a?.options || {}) }
    options.platform = props.platform
    options.serial = serial
    return { ...a, options }
  })
  currentActions.value = next
}

const handleDeviceChange = () => {
  if (!deviceSerial.value) return
  const actions = currentActions.value || []
  const next = actions.map(a => {
    const options = { ...(a?.options || {}) }
    if (!options.serial) {
      options.platform = props.platform
      options.serial = deviceSerial.value
    }
    return { ...a, options }
  })
  currentActions.value = next
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
  ensurePlatformContainer()
  payload[props.platform] = { actions: currentActions.value || [] }
  return payload
}

const save = async () => {
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

watch(() => props.formData, (newVal) => {
  if (newVal && Object.keys(newVal).length > 0) {
    formData.value = { ...formData.value, ...newVal }
    ensurePlatformContainer()
    extractDeviceSerialFromActions()
  }
}, { immediate: true, deep: true })

watch(() => props.platform, async () => {
  ensurePlatformContainer()
  await fetchDeviceOptions()
  extractDeviceSerialFromActions()
}, { immediate: true })

onMounted(async () => {
  ensurePlatformContainer()
  await fetchDeviceOptions()
  extractDeviceSerialFromActions()
})
</script>

<style scoped>
</style>
