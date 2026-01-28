
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="模型名称:" prop="name">
    <el-input v-model="formData.name" :clearable="false" placeholder="请输入模型名称" />
</el-form-item>
        <el-form-item label="请求模式:" prop="requestSchema">
    <el-select v-model="formData.requestSchema" placeholder="请选择请求模式" style="width:100%" filterable :clearable="false">
        <el-option v-for="(item,key) in RequestSchemaOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="模型标识:" prop="model">
    <el-select v-model="formData.model" placeholder="请选择模型标识" style="width:100%" filterable :clearable="false">
        <el-option v-for="(item,key) in modelOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="基础地址:" prop="baseURL">
    <el-input v-model="formData.baseURL" :clearable="false" placeholder="请输入基础地址" />
</el-form-item>
        <el-form-item label="密钥:" prop="apiKey">
    <el-input v-model="formData.apiKey" :clearable="false" placeholder="请输入密钥" />
</el-form-item>
        <el-form-item label="格式化输出:" prop="supportFormatOutput">
    <el-switch v-model="formData.supportFormatOutput" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="推理强度:" prop="reasoningEffort">
    <el-select v-model="formData.reasoningEffort" placeholder="请选择推理强度" style="width:100%" filterable :clearable="false">
        <el-option v-for="(item,key) in reasoningEffortOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="最大tokens:" prop="maxTokens">
    <el-input v-model.number="formData.maxTokens" :clearable="false" placeholder="请输入最大tokens" />
</el-form-item>
        <el-form-item label="随机性参数:" prop="temperature">
    <el-input-number v-model="formData.temperature" style="width:100%" :precision="2" :clearable="false" />
</el-form-item>
        <el-form-item label="核心采样:" prop="topP">
    <el-input-number v-model="formData.topP" style="width:100%" :precision="2" :clearable="false" />
</el-form-item>
        <el-form-item label="启用状态:" prop="enabled">
    <el-switch v-model="formData.enabled" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="超时时间:" prop="timeout">
    <el-input v-model.number="formData.timeout" :clearable="false" placeholder="请输入超时时间" />
</el-form-item>
        <el-form-item label="模型描述:" prop="description">
    <RichEdit v-model="formData.description"/>
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
  createLLMModelConfig,
  updateLLMModelConfig,
  findLLMModelConfig
} from '@/api/platform/llmModelService'

defineOptions({
    name: 'LLMModelConfigForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const RequestSchemaOptions = ref([])
const modelOptions = ref([])
const reasoningEffortOptions = ref([])
const formData = ref({
            name: '',
            requestSchema: '',
            model: '',
            baseURL: '',
            apiKey: '',
            supportFormatOutput: false,
            reasoningEffort: '',
            maxTokens: 0,
            temperature: 0,
            topP: 0,
            enabled: false,
            timeout: 0,
            description: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findLLMModelConfig({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    RequestSchemaOptions.value = await getDictFunc('RequestSchema')
    modelOptions.value = await getDictFunc('model')
    reasoningEffortOptions.value = await getDictFunc('reasoningEffort')
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createLLMModelConfig(formData.value)
               break
             case 'update':
               res = await updateLLMModelConfig(formData.value)
               break
             default:
               res = await createLLMModelConfig(formData.value)
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
