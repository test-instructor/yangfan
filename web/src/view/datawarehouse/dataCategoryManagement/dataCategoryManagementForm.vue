
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="数据分类名称:" prop="name">
    <el-input v-model="formData.name" :clearable="true" placeholder="请输入数据分类名称" />
</el-form-item>
        <el-form-item label="数据所在环境:" prop="envName">
    <el-input v-model="formData.envName" :clearable="true" placeholder="请输入数据所在环境" />
</el-form-item>
        <el-form-item label="数据总数量:" prop="count">
    <el-input v-model.number="formData.count" :clearable="true" placeholder="请输入数据总数量" />
</el-form-item>
        <el-form-item label="可用数据数量:" prop="availableCount">
    <el-input v-model.number="formData.availableCount" :clearable="true" placeholder="请输入可用数据数量" />
</el-form-item>
        <el-form-item label="创建数据的调用类型:" prop="createCallType">
    <el-input v-model.number="formData.createCallType" :clearable="true" placeholder="请输入创建数据的调用类型" />
</el-form-item>
        <el-form-item label="创建数据的测试步骤ID:" prop="createTestStepId">
    <el-input v-model="formData.createTestStepId" :clearable="true" placeholder="请输入创建数据的测试步骤ID" />
</el-form-item>
        <el-form-item label="清洗数据的调用类型:" prop="cleanCallType">
    <el-input v-model.number="formData.cleanCallType" :clearable="true" placeholder="请输入清洗数据的调用类型" />
</el-form-item>
        <el-form-item label="清洗数据的测试步骤ID:" prop="cleanTestStepId">
    <el-input v-model="formData.cleanTestStepId" :clearable="true" placeholder="请输入清洗数据的测试步骤ID" />
</el-form-item>
        <el-form-item label="是否直接删除数据:" prop="directDelete">
    <el-switch v-model="formData.directDelete" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="Python执行代码:" prop="pythonCode">
    <el-input v-model="formData.pythonCode" :clearable="true" placeholder="请输入Python执行代码" />
</el-form-item>
        <el-form-item label="projectId字段:" prop="projectId">
    <el-input v-model.number="formData.projectId" :clearable="true" placeholder="请输入projectId字段" />
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
  createDataCategoryManagement,
  updateDataCategoryManagement,
  findDataCategoryManagement
} from '@/api/datawarehouse/dataCategoryManagement'

defineOptions({
    name: 'DataCategoryManagementForm'
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
            envName: '',
            count: undefined,
            availableCount: undefined,
            createCallType: undefined,
            createTestStepId: '',
            cleanCallType: undefined,
            cleanTestStepId: '',
            directDelete: false,
            pythonCode: '',
            projectId: undefined,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findDataCategoryManagement({ ID: route.query.id })
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
               res = await createDataCategoryManagement(formData.value)
               break
             case 'update':
               res = await updateDataCategoryManagement(formData.value)
               break
             default:
               res = await createDataCategoryManagement(formData.value)
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
