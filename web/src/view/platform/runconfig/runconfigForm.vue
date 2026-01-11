<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="名称:" prop="name">
          <el-input v-model="formData.name" :clearable="false" placeholder="请输入名称" />
        </el-form-item>
        <el-form-item label="域名:" prop="base_url">
          <el-input v-model="formData.base_url" :clearable="false" placeholder="请输入域名" />
        </el-form-item>
        <el-form-item label="变量:" prop="variables">
          // 此字段为json结构，可以前端自行控制展示和数据绑定模式 需绑定json的key为 formData.variables 后端会按照json的类型进行存取
          {{ formData.variables }}
        </el-form-item>
        <el-form-item label="请求头:" prop="headers">
          // 此字段为json结构，可以前端自行控制展示和数据绑定模式 需绑定json的key为 formData.headers 后端会按照json的类型进行存取
          {{ formData.headers }}
        </el-form-item>
        <el-form-item label="参数:" prop="parameters">
          // 此字段为json结构，可以前端自行控制展示和数据绑定模式 需绑定json的key为 formData.parameters 后端会按照json的类型进行存取
          {{ formData.parameters }}
        </el-form-item>
        <el-form-item label="变量JSON:" prop="variables_json">
          // 此字段为json结构，可以前端自行控制展示和数据绑定模式 需绑定json的key为 formData.variables_json
          后端会按照json的类型进行存取
          {{ formData.variables_json }}
        </el-form-item>
        <el-form-item label="请求头JSON:" prop="headers_json">
          // 此字段为json结构，可以前端自行控制展示和数据绑定模式 需绑定json的key为 formData.headers_json
          后端会按照json的类型进行存取
          {{ formData.headers_json }}
        </el-form-item>
        <el-form-item label="权重:" prop="weight">
          <el-input v-model.number="formData.weight" :clearable="false" placeholder="请输入权重" />
        </el-form-item>
        <el-form-item label="超时:" prop="timeout">
          <el-input-number v-model="formData.timeout" style="width:100%" :precision="2" :clearable="false" />
        </el-form-item>
        <el-form-item label="允许重定向:" prop="allow_redirects">
          <el-switch v-model="formData.allow_redirects" active-color="#13ce66" inactive-color="#ff4949" active-text="是"
                     inactive-text="否" clearable></el-switch>
        </el-form-item>
        <el-form-item label="验证:" prop="verify">
          <el-switch v-model="formData.verify" active-color="#13ce66" inactive-color="#ff4949" active-text="是"
                     inactive-text="否" clearable></el-switch>
        </el-form-item>
        <el-form-item label="导出参数:" prop="export">
          <ArrayCtrl v-model="formData.export" editable />
        </el-form-item>
        <el-form-item label="前置步骤:" prop="preparatorySteps">
          <el-input v-model.number="formData.preparatorySteps" :clearable="false" placeholder="请输入前置步骤" />
        </el-form-item>
        <el-form-item label="前置步骤ID:" prop="setup_case_id">
        </el-form-item>
        <el-form-item label="报告ID:" prop="report_id">
          <el-input v-model.number="formData.report_id" :clearable="false" placeholder="请输入报告ID" />
        </el-form-item>
        <el-form-item label="重试次数:" prop="retry">
          <el-input v-model.number="formData.retry" :clearable="true" placeholder="请输入重试次数" />
        </el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
  import {
    createRunConfig,
    updateRunConfig,
    findRunConfig
  } from '@/api/platform/runconfig'

  defineOptions({
    name: 'RunConfigForm'
  })

  // 自动获取字典
  import { getDictFunc } from '@/utils/format'
  import { useRoute, useRouter } from 'vue-router'
  import { ElMessage } from 'element-plus'
  import { ref, reactive } from 'vue'
  // 数组控制组件
  import ArrayCtrl from '@/components/arrayCtrl/arrayCtrl.vue'


  const route = useRoute()
  const router = useRouter()

  // 提交按钮loading
  const btnLoading = ref(false)

  const type = ref('')
  const formData = ref({
    name: '',
    base_url: '',
    variables: {},
    headers: {},
    parameters: {},
    variables_json: {},
    headers_json: {},
    weight: undefined,
    timeout: 0,
    allow_redirects: false,
    verify: false,
    export: [],
    preparatorySteps: undefined,
    setup_case_id: null,
    report_id: undefined,
    retry: undefined
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
      const res = await findRunConfig({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
  }

  init()
  // 保存按钮
  const save = async () => {
    btnLoading.value = true
    elFormRef.value?.validate(async (valid) => {
      if (!valid) return btnLoading.value = false
      let res
      switch (type.value) {
        case 'create':
          res = await createRunConfig(formData.value)
          break
        case 'update':
          res = await updateRunConfig(formData.value)
          break
        default:
          res = await createRunConfig(formData.value)
          break
      }
      btnLoading.value = false
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '创建/更改成功'
        })
      }
    })
  }

  // 返回按钮
  const back = () => {
    router.go(-1)
  }

</script>

<style>
</style>
