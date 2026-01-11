
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用例名称:" prop="caseName">
    <el-input v-model="formData.caseName" :clearable="false" placeholder="请输入用例名称" />
</el-form-item>
        <el-form-item label="运行次数:" prop="runNumber">
    <el-input v-model.number="formData.runNumber" :clearable="true" placeholder="请输入运行次数" />
</el-form-item>
        <el-form-item label="状态:" prop="status">
    <el-select v-model="formData.status" placeholder="请选择状态" style="width:100%" filterable :clearable="false">
       <el-option v-for="item in ['测试中 ',' 待评审 ',' 评审不通过 ',' 已发布 ',' 禁用 ',' 已废弃 ']" :key="item" :label="item" :value="item" />
    </el-select>
</el-form-item>
        <el-form-item label="运行环境:" prop="envName">
    <el-input v-model="formData.envName" :clearable="true" placeholder="请输入运行环境" />
</el-form-item>
        <el-form-item label="描述:" prop="desc">
    <el-input v-model="formData.desc" :clearable="false" placeholder="请输入描述" />
</el-form-item>
        <el-form-item label="配置名称:" prop="configName">
    <el-input v-model="formData.configName" :clearable="true" placeholder="请输入配置名称" />
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
  createAutoCase,
  updateAutoCase,
  findAutoCase
} from '@/api/automation/autocase'

defineOptions({
    name: 'AutoCaseForm'
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
            caseName: '',
            runNumber: undefined,
            status: null,
            envName: '',
            desc: '',
            configName: '',
        })
// 验证规则
const rule = reactive({
               caseName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               desc : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findAutoCase({ ID: route.query.id })
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
               res = await createAutoCase(formData.value)
               break
             case 'update':
               res = await updateAutoCase(formData.value)
               break
             default:
               res = await createAutoCase(formData.value)
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
