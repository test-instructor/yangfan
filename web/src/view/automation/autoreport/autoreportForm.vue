
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="名称:" prop="name">
    <el-input v-model="formData.name" :clearable="false" placeholder="请输入名称" />
</el-form-item>
        <el-form-item label="成功:" prop="success">
    <el-switch v-model="formData.success" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="统计ID:" prop="stat_id">
    <el-input v-model.number="formData.stat_id" :clearable="false" placeholder="请输入统计ID" />
</el-form-item>
        <el-form-item label="时间ID:" prop="time_id">
    <el-input v-model.number="formData.time_id" :clearable="false" placeholder="请输入时间ID" />
</el-form-item>
        <el-form-item label="平台:" prop="platform">
    // 此字段为json结构，可以前端自行控制展示和数据绑定模式 需绑定json的key为 formData.platform 后端会按照json的类型进行存取
    {{ formData.platform }}
</el-form-item>
        <el-form-item label="状态:" prop="status">
    <el-input v-model.number="formData.status" :clearable="false" placeholder="请输入状态" />
</el-form-item>
        <el-form-item label="设置案例:" prop="setup_case">
    <el-switch v-model="formData.setup_case" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="描述:" prop="describe">
    <el-input v-model="formData.describe" :clearable="false" placeholder="请输入描述" />
</el-form-item>
        <el-form-item label="API环境名称:" prop="api_env_name">
    <el-input v-model="formData.api_env_name" :clearable="false" placeholder="请输入API环境名称" />
</el-form-item>
        <el-form-item label="API环境ID:" prop="api_env_id">
    <el-input v-model.number="formData.api_env_id" :clearable="false" placeholder="请输入API环境ID" />
</el-form-item>
        <el-form-item label="运行节点:" prop="node_name">
    <el-input v-model="formData.node_name" :clearable="false" placeholder="请输入运行节点" disabled />
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
  createAutoReport,
  updateAutoReport,
  findAutoReport
} from '@/api/automation/autoreport'

defineOptions({
    name: 'AutoReportForm'
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
            name: '',
            success: false,
            stat_id: 0,
            time_id: 0,
            platform: {},
            status: 0,
            setup_case: false,
            describe: '',
            api_env_name: '',
            api_env_id: 0,
            node_name: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findAutoReport({ ID: route.query.id })
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
               res = await createAutoReport(formData.value)
               break
             case 'update':
               res = await updateAutoReport(formData.value)
               break
             default:
               res = await createAutoReport(formData.value)
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
