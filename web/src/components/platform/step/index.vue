<template>
  <div>
    <div class="gva-form-box" style="width:90%">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="100px">
        <el-form-item label="步骤名称:" prop="name">
          <el-input v-model="formData.name" :clearable="false" placeholder="请输入步骤名称" />
        </el-form-item>
        <el-form-item label="请求配置">
          <div>
            <div style="display: flex; gap: 10px;">
              <el-select
                v-model="formData.request.method"
                placeholder="选择请求方法"
                clearable
                style="width: 150px;"
              >
                <el-option label="GET" value="GET"></el-option>
                <el-option label="POST" value="POST"></el-option>
                <el-option label="PUT" value="PUT"></el-option>
                <el-option label="DELETE" value="DELETE"></el-option>
                <el-option label="PATCH" value="PATCH"></el-option>
                <el-option label="HEAD" value="HEAD"></el-option>
                <el-option label="OPTIONS" value="OPTIONS"></el-option>
                <template #prepend>URL</template>
              </el-select>
              <el-input
                v-model="formData.request.url"
                placeholder="请输入URL或请求地址"
                style="width: 420px;"
                clearable
              >
                <template #prepend>URL</template>
              </el-input>
              <span class="custom-prepend">失败重试次数 </span>
              <el-input-number
                v-model.number="formData.retry"
                :clearable="false"
                placeholder="请输入循环次数"
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
            </div>
          </div>
        </el-form-item>
        <el-form-item>
          <div style="height: 500px;width: 1200px">
            <el-tabs v-model="activeConfig" class="demo-tabs" @tab-click="handleClick">

              <el-tab-pane label="请求头" name="header">
                <HeaderTable
                  :header="formData?.request?.header_temp ?? []"
                  :heights="500"
                  @headerData="handleHeaderData"
                />
              </el-tab-pane>
              <el-tab-pane label="查询参数(params)" name="params">
                <Params
                  :params="formData?.request?.param_temp ?? []"
                  :heights="500"
                  @paramsData="handleParamsData"
                />
              </el-tab-pane>
              <el-tab-pane label="表单(form)" name="form">
                <Form
                  :forms="formData?.request?.data_temp ?? []"
                  :heights="500"
                  @requestFormData="handleFormData"
                />
              </el-tab-pane>
              <el-tab-pane label="请求体(json)" name="json">
                <JsonEditor
                  :heights="500"
                  :jsons="formData?.request?.json ?? {}"
                  @jsonData="handleJsonData"
                />
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
                  v-model="formData.request.data_warehouse"
                  :hideCount="hideCount"
                  :height="460"
                />
              </el-tab-pane>
            </el-tabs>
          </div>
        </el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">取消</el-button>

          <EnvSelector />
          <PythonFuncSelector />
          <DataWarehouseFieldSelector :current-type="formData.request?.data_warehouse?.type" />
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
  import {
    createAutoStep,
    updateAutoStep,
    findAutoStep
  } from '@/api/automation/autostep'

  defineOptions({
    name: 'AutoStepForm'
  })

  // 自动获取字典
  import { getDictFunc } from '@/utils/format'
  // import { useRoute, useRouter } from 'vue-router'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { ref, reactive, watch, onMounted } from 'vue'
  // 数组控制组件
  import ArrayCtrl from '@/components/arrayCtrl/arrayCtrl.vue'
  import ParameterTable from '@/components/platform/parameterTable/index.vue'
  import JsonEditor from '@/components/platform/jsonEdit/index.vue'
  import Variables from '@/components/platform/variables/index.vue'
  import HeaderTable from '@/components/platform/header/index.vue'
  import Params from '@/components/platform/params/index.vue'
  import Form from '@/components/platform/form/index.vue'
  import Extractor from '@/components/platform/extract/index.vue'
  import Validate from '@/components/platform/validate/index.vue'
  import Hook from '@/components/platform/hook/index.vue'
  import utils from '@/utils/dataTypeConverter.js'
  import EnvSelector from '@/components/platform/button/EnvDetail.vue'
  import PythonFuncSelector from '@/components/platform/button/PythonFuncSelector.vue'
  import DataWarehouseFieldSelector from '@/components/platform/button/DataWarehouseFieldSelector.vue'
  import DataWarehouseConfig from '@/components/platform/dataWarehouseConfig/index.vue'

  // const route = useRoute()

  // 提交按钮loading
  const btnLoading = ref(false)

  const props = defineProps({
    menu: {
      type: Number,
      default: 0
    },
    eventType: {
      type: String,
      default: '0'
    },
    formData: {
      type: Object,
      default: () => ({})
    },
    stepType: {
      type: String,
      default: ''
    },
    hideCount: {
      type: Boolean,
      default: false
    }
  })
  const emit = defineEmits(['close'])

  const stepType = ref('')
  const formData = ref({
    name: '',
    loops: 0,
    retry: 0,
    request: {
      method: '',
      url: '',
      header_temp: [],
      headers: {},
      json: {},
      data_warehouse: {}
    },
    parameters: {},
    parameters_temp: {},
    menu: 0
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
  // const init = async () => {
  //   // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  //   if (route.query.id) {
  //     const res = await findAutoStep({ ID: route.query.id })
  //     if (res.code === 0) {
  //       formData.value = res.data
  //       type.value = 'update'
  //     }
  //   } else {
  //     type.value = 'create'
  //   }
  // }
  //
  // init()

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
    if (formData.value.request.method === '') {
      messageBox('请求方法不能为空')
      return false
    }
    const variablesData = utils.convertData(variables.value)
    if (!variablesData.success) {
      messageBox('数据处理失败字段:' + variablesData.errors.key + ',出现错误:' + variablesData.errors.error)
      activeConfig.value = 'variables'
      return false
    }
    formData.value.variables_temp = variables.value
    formData.value.variables = variablesData.data

    const headerDataNew = utils.processData(headers.value)
    if (!headerDataNew.success) {
      messageBox('数据处理失败:' + headerDataNew.errors.join('、') + ',字段重复')
      activeConfig.value = 'header'
      return false
    }
    formData.value.request.header_temp = headers.value
    formData.value.request.headers = headerDataNew.data

    const paramDataNew = utils.processData(paramsData.value)
    if (!paramDataNew.success) {

      messageBox('数据处理失败:' + paramDataNew.errors.join('、') + ',字段重复')
      activeConfig.value = 'params'
      return false
    }
    formData.value.request.param_temp = paramsData.value
    formData.value.request.params = paramDataNew.data

    if (jsonError.value != '') {
      messageBox('参数设置数据处理失败:' + jsonError.value)
      activeConfig.value = 'json'
      return false
    }

    if (jsonErrorParameters.value != '') {
      messageBox('参数设置数据处理失败:' + jsonErrorParameters.value)
      activeConfig.value = 'parameters'
      return false
    }


    const formDataNew = utils.convertData(formDataJson.value)
    if (!formDataNew.success) {

      messageBox('数据处理失败字段:' + formDataNew.errors.key + ',出现错误:' + formDataNew.errors.error)
      activeConfig.value = 'form'
      return false
    }
    formData.value.request.data_temp = formDataJson.value
    formData.value.request.data = formDataNew.data

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

    // Data Warehouse Validation
    const dwConfig = formData.value.request.data_warehouse
    if (dwConfig && dwConfig.filter && dwConfig.filter.groups) {
      for (const group of dwConfig.filter.groups) {
        if (group.conditions) {
          for (const cond of group.conditions) {
            if (!cond.field || cond.field.trim() === '') {
              messageBox('数据仓库配置中有未选择字段的筛选条件')
              activeConfig.value = 'dataWarehouse'
              return false
            }
          }
        }
      }
    }

    return true
  }

  // 保存按钮
  const save = async () => {
    const validation = await dataValidation()

    if (!validation) return
    btnLoading.value = true
    elFormRef.value?.validate(async (valid) => {
      if (!valid) return btnLoading.value = false
      let res
      switch (stepType.value) {
        case 'create':
          formData.value.menu = Number(props.menu)
          res = await createAutoStep(formData.value)
          break
        case 'update':
          res = await updateAutoStep(formData.value)
          break
        case 'copy':
          res = await createAutoStep(formData.value)
          break
        default:
          formData.value.menu = Number(props.menu)
          res = await createAutoStep(formData.value)
          break
      }
      btnLoading.value = false
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '创建/更改成功'
        })
        emit('close')
      }
    })
  }

  // 返回按钮
  const back = () => {
    // router.go(-1)
    const close = ref(true)
    if (stepType.value == 'update'){
      close.value = false
    }else {
      close.value = true
    }

    emit('close', {value: close})

  }
  const jsonError = ref('')
  const handleJsonData = (result) => {
    if (!result.isValid) {
      jsonError.value = result.error.message // 展示错误信息
    } else {
      jsonError.value = '' // 清空错误信息
      formData.value.request.json = result.data
    }
  }
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
  const activeConfig = ref('header')
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

      if (!formData.value.request) {
        formData.value.request = {};
      }

      // 判断 url 是否存在，不存在则设为空字符串
      if (formData.value.request.url === undefined || formData.value.request.url === null) {
        formData.value.request.url = '';
      }

      // 判断 method 是否存在，不存在则设为空字符串
      if (formData.value.request.method === undefined || formData.value.request.method === null) {
        formData.value.request.method = '';
      }

      if (!formData.value.request.data_warehouse) {
        formData.value.request.data_warehouse = {};
      }

    }
  }, { deep: true })

  onMounted(() => {
    // 组件挂载完成（打开时）执行判断
    if (Object.keys(props.formData).length) {
      formData.value = props.formData
      if (!formData.value.request) {
        formData.value.request = {};
      }

      // 判断 url 是否存在，不存在则设为空字符串
      if (formData.value.request.url === undefined || formData.value.request.url === null) {
        formData.value.request.url = '';
      }

      // 判断 method 是否存在，不存在则设为空字符串
      if (formData.value.request.method === undefined || formData.value.request.method === null) {
        formData.value.request.method = '';
      }
      
      if (!formData.value.request.data_warehouse) {
        formData.value.request.data_warehouse = {};
      }
    }
    console.log('props========', props.stepType)
    stepType.value = props.stepType
  });
  watch(
    () => props, // 监听整个 props 对象
    (newVal, oldVal) => {
      console.log('props', newVal, oldVal)
    },
    {
      deep: true,
      immediate: true,
    }
  )

</script>

<style>
</style>
