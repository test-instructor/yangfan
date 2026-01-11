
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="步骤名称:" prop="name">
    <el-input v-model="formData.name" :clearable="false" placeholder="请输入步骤名称" />
</el-form-item>
        <el-form-item label="变量:" prop="variables">
</el-form-item>
        <el-form-item label="参数:" prop="parameters">
</el-form-item>
        <el-form-item label="设置钩子:" prop="setup_hooks">
    <ArrayCtrl v-model="formData.setup_hooks" editable/>
</el-form-item>
        <el-form-item label="清理钩子:" prop="teardown_hooks">
    <ArrayCtrl v-model="formData.teardown_hooks" editable/>
</el-form-item>
        <el-form-item label="提取:" prop="extract">
</el-form-item>
        <el-form-item label="验证器:" prop="validate">
    <ArrayCtrl v-model="formData.validate" editable/>
</el-form-item>
        <el-form-item label="步骤导出:" prop="export">
    <ArrayCtrl v-model="formData.export" editable/>
</el-form-item>
        <el-form-item label="循环次数:" prop="loops">
    <el-input v-model.number="formData.loops" :clearable="false" placeholder="请输入循环次数" />
</el-form-item>
        <el-form-item label="忽略弹出窗口:" prop="ignore_popup">
    <el-switch v-model="formData.ignore_popup" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
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
  createAutoStep,
  updateAutoStep,
  findAutoStep
} from '@/api/automation/autostep'

defineOptions({
    name: 'AutoStepForm'
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
            variables: null,
            parameters: null,
            setup_hooks: [],
            teardown_hooks: [],
            extract: null,
            validate: [],
            export: [],
            loops: 0,
            ignore_popup: false,
        })
// 验证规则
const rule = reactive({
               name : [{
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
      const res = await findAutoStep({ ID: route.query.id })
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
               res = await createAutoStep(formData.value)
               break
             case 'update':
               res = await updateAutoStep(formData.value)
               break
             default:
               res = await createAutoStep(formData.value)
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
