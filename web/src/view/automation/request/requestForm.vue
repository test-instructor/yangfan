
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="方法:" prop="method">
    <el-input v-model="formData.method" :clearable="false" placeholder="请输入方法" />
</el-form-item>
        <el-form-item label="URL:" prop="url">
    <el-input v-model="formData.url" :clearable="false" placeholder="请输入URL" />
</el-form-item>
        <el-form-item label="HTTP2:" prop="http2">
    <el-switch v-model="formData.http2" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="参数:" prop="params">
</el-form-item>
        <el-form-item label="头:" prop="headers">
</el-form-item>
        <el-form-item label="Cookies:" prop="cookies">
</el-form-item>
        <el-form-item label="Body:" prop="body">
</el-form-item>
        <el-form-item label="Json:" prop="json">
</el-form-item>
        <el-form-item label="数据:" prop="data">
</el-form-item>
        <el-form-item label="超时:" prop="timeout">
    <el-input-number v-model="formData.timeout" style="width:100%" :precision="2" :clearable="false" />
</el-form-item>
        <el-form-item label="允许重定向:" prop="allow_redirects">
    <el-switch v-model="formData.allow_redirects" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="验证:" prop="verify">
    <el-switch v-model="formData.verify" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="上传:" prop="upload">
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
  createRequest,
  updateRequest,
  findRequest
} from '@/api/automation/request'

defineOptions({
    name: 'RequestForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const formData = ref({
            method: '',
            url: '',
            http2: false,
            params: null,
            headers: null,
            cookies: null,
            body: null,
            json: null,
            data: null,
            timeout: 0,
            allow_redirects: false,
            verify: false,
            upload: null,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findRequest({ ID: route.query.id })
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
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createRequest(formData.value)
               break
             case 'update':
               res = await updateRequest(formData.value)
               break
             default:
               res = await createRequest(formData.value)
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
