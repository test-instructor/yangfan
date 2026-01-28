
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="浏览器标识:" prop="browserId">
    <el-input v-model="formData.browserId" :clearable="false" placeholder="请输入浏览器标识" />
</el-form-item>
        <el-form-item label="登录状态:" prop="logOn">
    <el-switch v-model="formData.logOn" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="忽略弹窗:" prop="ignorePopup">
    <el-switch v-model="formData.ignorePopup" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="宽度:" prop="width">
    <el-input v-model.number="formData.width" :clearable="false" placeholder="请输入宽度" />
</el-form-item>
        <el-form-item label="高度:" prop="height">
    <el-input v-model.number="formData.height" :clearable="false" placeholder="请输入高度" />
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
  createBrowserDeviceOptions,
  updateBrowserDeviceOptions,
  findBrowserDeviceOptions
} from '@/api/platform/browserDeviceConfig'

defineOptions({
    name: 'BrowserDeviceOptionsForm'
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
            browserId: '',
            logOn: false,
            ignorePopup: false,
            width: 0,
            height: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findBrowserDeviceOptions({ ID: route.query.id })
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
               res = await createBrowserDeviceOptions(formData.value)
               break
             case 'update':
               res = await updateBrowserDeviceOptions(formData.value)
               break
             default:
               res = await createBrowserDeviceOptions(formData.value)
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
