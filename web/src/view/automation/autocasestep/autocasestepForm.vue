<template>
  <div>
    <div class="gva-form-box" style="width:100%">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="100px">
        <el-form-item label="步骤名称:" prop="name">
          <el-input v-model="formData.name" :clearable="false" placeholder="请输入步骤名称" />
        </el-form-item>
        <el-form-item label="请求配置">
          <div>
            <div style="display: flex; gap: 10px;">
              <span class="custom-prepend">失败重试次数 </span>
              <el-input-number
                v-model.number="formData.retry"
                :clearable="false"
                style="width:120px"
                :max="99"
              >
              </el-input-number>
              <span class="custom-prepend">循环次数 </span>
              <el-input-number
                v-model.number="formData.loops"
                :clearable="false"
                placeholder="请输入循环次数"
                style="width:120px"
                :max="99"
              >
              </el-input-number>
              <RunConfig v-model="formData.configID" @change="handleConfigSelect"/>
              <Env v-model="formData.envID" @change="handleEnvChange"/>
            </div>
          </div>
        </el-form-item>
        <el-form-item>
          <div style="height: 500px;width: 1200px">
            <el-tabs v-model="activeConfig" class="demo-tabs" @tab-click="handleClick">
              <el-tab-pane label="跳过执行" name="skip">
                <Validate
                  :validate="formData.skip_temp"
                  :heights="500"
                  idName="skip"
                  @jsonData="handleJsonDataSkip"
                />
              </el-tab-pane>
              <el-tab-pane label="变量" name="variables">
                <Variables
                  :variables="formData.variables_temp"
                  :heights="500"
                  @variablesData="handleVariablesData"
                />
              </el-tab-pane>
              <el-tab-pane label="参数提取" name="extract">
                <Extractor
                  :extract="formData.extract_temp"
                  :heights="500"
                  idName="extract"
                  @extractData="handleJsonDataExtract"
                />
              </el-tab-pane>
              <el-tab-pane label="勾子(hook)" name="hook">
                <Hook
                  :setupHooks="formData.setup_hooks"
                  :teardownHooks="formData.teardown_hooks"
                  :heights="500"
                  @setupHooksData="handleSetupHooksData"
                  @teardownHooksData="handleTeardownHooksData"
                />
              </el-tab-pane>
              <el-tab-pane label="断言" name="validate">
                <Validate
                  :validate="formData.validators_temp"
                  :heights="500"
                  idName="validate"
                  @jsonData="handleJsonDataValidate"
                />
              </el-tab-pane>
              <el-tab-pane label="参数设置" name="parameters">
                <ParameterTable
                  :jsons="formData.parameters"
                  :parametersTemp="formData.parameters_temp"
                  @jsonData="handleJsonDataParameters"
                  @tempData="(val) => formData.parameters_temp = val"
                />
              </el-tab-pane>
              <el-tab-pane label="数据仓库" name="datawarehouse">
                <DataWarehouseConfig
                  v-model="formData.data_warehouse_temp"
                  :height="500"
                />
              </el-tab-pane>
            </el-tabs>
          </div>
        </el-form-item>
        <el-form-item>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>


<script setup>
  import {
    createAutoCaseStep,
    updateAutoCaseStep,
    findAutoCaseStep
  } from '@/api/automation/autocasestep'

  defineOptions({
    name: 'AutoCaseStepForm'
  })

  // 自动获取字典
  import { getDictFunc } from '@/utils/format'
  import { useRoute, useRouter } from 'vue-router'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { ref, reactive, watch, onMounted } from 'vue'
  // 数组控制组件
  import ArrayCtrl from '@/components/arrayCtrl/arrayCtrl.vue'
  import HeaderTable from '@/components/platform/header/index.vue'
  import Variables from '@/components/platform/variables/index.vue'
  import Extractor from '@/components/platform/extract/index.vue'
  import Form from '@/components/platform/form/index.vue'
  import ParameterTable from '@/components/platform/parameterTable/index.vue'
  import JsonEditor from '@/components/platform/jsonEdit/index.vue'
  import Hook from '@/components/platform/hook/index.vue'
  import PythonFuncSelector from '@/components/platform/button/PythonFuncSelector.vue'
  import Validate from '@/components/platform/validate/index.vue'
  import EnvSelector from '@/components/platform/button/EnvDetail.vue'
  import Params from '@/components/platform/params/index.vue'
  import utils from '@/utils/dataTypeConverter.js'
  import Env from '@/components/platform/env.vue'
  import RunConfig from '@/components/platform/runConfig.vue'
  import DataWarehouseConfig from '@/components/platform/dataWarehouseConfig/index.vue'


  const route = useRoute()
  const router = useRouter()

  // 提交按钮loading
  const btnLoading = ref(false)

  const props = defineProps({
    menu: {
      type: String,
      default: '99999'
    },
    formData: {
      type: Object,
      default: () => ({})
    },
    type: {
      type: String,
      default: 'create'
    }
  })
  const emit = defineEmits(['close'])

  const type = ref('')
  const formData = ref({
    name: '',
    variables: null,
    parameters: null,
    parameters_temp: {},
    data_warehouse: null,
    data_warehouse_temp: {},
    setup_hooks: [],
    teardown_hooks: [],
    extract: null,
    validate: [],
    export: [],
    loops: 0,
    ignore_popup: false
  })
  // 验证规则
  const rule = reactive({
    name: [{
      required: true,
      message: '',
      trigger: ['input', 'blur']
    }]
  })

  const elFormRef = ref()

  // 初始化方法
  const init = async () => {
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findAutoCaseStep({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
  }

  init()

  // 数据校验
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
      activeConfig.value = 'variables'
      return false
    }
    formData.value.variables_temp = variables.value
    formData.value.variables = variablesData.data

    if (jsonErrorParameters.value != '') {
      messageBox('参数设置数据处理失败:' + jsonErrorParameters.value)
      activeConfig.value = 'parameters'
      return false
    }

    // If using ParameterTable, conversion is already done via handleJsonDataParameters
    // But if we need to ensure consistency, we can double check
    // Actually handleJsonDataParameters updates formData.parameters directly

    // 同步 data_warehouse_temp 到 data_warehouse
    if (formData.value.data_warehouse_temp && Object.keys(formData.value.data_warehouse_temp).length > 0) {
      formData.value.data_warehouse = formData.value.data_warehouse_temp
    } else {
      formData.value.data_warehouse = null
    }

    const skipDataNew = utils.assertionData(skipData.value)
    if (!skipDataNew.success) {
      messageBox('数据处理失败字段:' + skipDataNew.errors.key + ',出现错误:' + skipDataNew.errors.error)
      activeConfig.value = 'skip'
      return false
    }
    formData.value.skip = skipDataNew.data
    formData.value.skip_temp = skipData.value

    const validateDataNew = utils.assertionData(validateData.value)
    if (!validateDataNew.success) {

      messageBox('数据处理失败字段:' + validateDataNew.errors.key + ',出现错误:' + validateDataNew.errors.error)
      activeConfig.value = 'validate'
      return false
    }
    formData.value.validate = validateDataNew.data
    formData.value.validators_temp = validateData.value
    const extractDataNew = utils.processData(extractorData.value)
    if (!extractDataNew.success) {

      messageBox('数据处理失败:' + extractDataNew.errors.join('、') + ',字段重复')
      activeConfig.value = 'extract'
      return false
    }
    formData.value.extract = extractDataNew.data
    formData.value.extract_temp = extractorData.value

    formData.value.setup_hooks = setupHooksData.value
    formData.value.teardown_hooks = teardownHooksData.value

    return true
  }


  // 保存按钮
  const save = async () => {
    const validation = await dataValidation()
    if (!validation) return

    try {
      const valid = await elFormRef.value?.validate()
      if (!valid) return
    } catch (e) {
      return
    }

    let res
    switch (type.value) {
      case 'create':
        formData.value.menu = props.menu
        res = await createAutoCaseStep(formData.value)
        break
      case 'update':
        res = await updateAutoCaseStep(formData.value)
        break
      default:
        res = await createAutoCaseStep(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
      const close = ref(true)
      emit('close', {value: close})
    }
  }

  defineExpose({
    save,
    formData
  })

  // handleJsonDataParameters
  const jsonErrorParameters = ref('')
  const handleJsonDataParameters = (result) => {
    if (!result.isValid) {
      jsonErrorParameters.value = result.error.message // 展示错误信息
    } else {
      jsonErrorParameters.value = '' // 清空错误信息
      formData.value.parameters = result.data
    }
  }
  // el-tabs 默认选中 header
  const activeConfig = ref('skip')
  // 请求头数据接收
  const headers = ref([])
  const handleHeaderData = (tableData) => {
    headers.value = tableData
  }
  // 变量数据接收
  const variables = ref([])
  const handleVariablesData = (tableData) => {
    variables.value = tableData
  }
  const paramsData = ref([])
  const handleParamsData = (tableData) => {
    paramsData.value = tableData
  }
  const formDataJson = ref([])
  const handleFormData = (tableData) => {
    formDataJson.value = tableData
  }

  const skipData = ref([])
  const handleJsonDataSkip = (tableData) => {
    skipData.value = tableData
  }
  const extractorData = ref([])
  const handleJsonDataExtract = (tableData) => {
    extractorData.value = tableData
  }
  const setupHooksData = ref([])
  const handleSetupHooksData = (tableData) => {
    // formData.value.setup_hooks = tableData
    setupHooksData.value = tableData
  }
  const teardownHooksData = ref([])
  const handleTeardownHooksData = (tableData) => {
    // formData.value.teardown_hooks = tableData
    teardownHooksData.value = tableData
  }
  // handleJsonDataValidate
  const validateData = ref([])
  const handleJsonDataValidate = (tableData) => {
    validateData.value = tableData
  }

  watch(() => props.formData, (newVal) => {
    if (newVal?.length) {
      formData.value = newVal
    }
  }, { deep: true })

  onMounted(() => {
    // 组件挂载完成（打开时）执行判断
    if (Object.keys(props.formData).length) {
      formData.value = props.formData
    }
    type.value = props.type
  });

  const handleConfigSelect = (config) => {
    if (config) {
      formData.value.configName = config.name;
    } else {
      formData.value.configName = null;
    }
  }

  // 处理选择结果
  const handleEnvChange = (env) => {
    if (env) {
      console.log('选中的环境：', env.ID, env.name)
      formData.value.envName = env.name
    } else {
      formData.value.envName = null
    }
  }


</script>

<style>
</style>
