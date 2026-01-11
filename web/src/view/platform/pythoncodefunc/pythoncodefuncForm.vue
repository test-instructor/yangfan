
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="名称:" prop="name">
    <el-input v-model="formData.name" :clearable="false" placeholder="请输入名称" />
</el-form-item>
        <el-form-item label="参数:" prop="params">
    <ArrayCtrl v-model="formData.params" editable/>
</el-form-item>
        <el-form-item label="完整代码:" prop="fullCode">
    <el-input v-model="formData.fullCode" :clearable="false" placeholder="请输入完整代码" />
</el-form-item>
        <el-form-item label="起始索引:" prop="startIndex">
    <el-input v-model.number="formData.startIndex" :clearable="false" placeholder="请输入起始索引" />
</el-form-item>
        <el-form-item label="项目信息:" prop="projectId">
    <el-input v-model.number="formData.projectId" :clearable="true" placeholder="请输入项目信息" />
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
  createPythonCodeFunc,
  updatePythonCodeFunc,
  findPythonCodeFunc
} from '@/api/platform/pythoncodefunc'

defineOptions({
    name: 'PythonCodeFuncForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
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
            params: [],
            fullCode: '',
            startIndex: 0,
            projectId: undefined,
        })
// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               fullCode : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               startIndex : [{
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
      const res = await findPythonCodeFunc({ ID: route.query.id })
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
               res = await createPythonCodeFunc(formData.value)
               break
             case 'update':
               res = await updatePythonCodeFunc(formData.value)
               break
             default:
               res = await createPythonCodeFunc(formData.value)
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
