
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="变量名称:" prop="key">
    <el-input v-model="formData.key" :clearable="false" placeholder="请输入变量名称" />
</el-form-item>
        <el-form-item label="中文名:" prop="name">
    <el-input v-model="formData.name" :clearable="false" placeholder="请输入中文名" />
</el-form-item>
        <el-form-item label="变量值:" prop="value">
    // 此字段为json结构，可以前端自行控制展示和数据绑定模式 需绑定json的key为 formData.value 后端会按照json的类型进行存取
    {{ formData.value }}
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
  createEnvDetail,
  updateEnvDetail,
  findEnvDetail
} from '@/api/platform/envdetail'

defineOptions({
    name: 'EnvDetailForm'
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
            key: '',
            name: '',
            value: {},
        })
// 验证规则
const rule = reactive({
               key : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               value : [{
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
      const res = await findEnvDetail({ ID: route.query.id })
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
               res = await createEnvDetail(formData.value)
               break
             case 'update':
               res = await updateEnvDetail(formData.value)
               break
             default:
               res = await createEnvDetail(formData.value)
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
