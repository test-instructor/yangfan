
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="设备名称:" prop="name">
    <el-input v-model="formData.name" :clearable="true" placeholder="请输入设备名称" />
</el-form-item>
        <el-form-item label="序列号:" prop="serial">
    <el-input v-model="formData.serial" :clearable="false" placeholder="请输入序列号" />
</el-form-item>
        <el-form-item label="日志开关:" prop="logOn">
    <el-switch v-model="formData.logOn" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="忽略弹窗:" prop="ignorePopup">
    <el-switch v-model="formData.ignorePopup" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="ADB主机:" prop="adbServerHost">
    <el-input v-model="formData.adbServerHost" :clearable="false" placeholder="请输入ADB主机" />
</el-form-item>
        <el-form-item label="ADB端口:" prop="adbServerPort">
    <el-input v-model.number="formData.adbServerPort" :clearable="false" placeholder="请输入ADB端口" />
</el-form-item>
        <el-form-item label="复合驱动:" prop="composite">
    <el-switch v-model="formData.composite" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="UIA2开关:" prop="uia2">
    <el-switch v-model="formData.uia2" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="UIA2地址:" prop="uia2Ip">
    <el-input v-model="formData.uia2Ip" :clearable="false" placeholder="请输入UIA2地址" />
</el-form-item>
        <el-form-item label="UIA2端口:" prop="uia2Port">
    <el-input v-model.number="formData.uia2Port" :clearable="false" placeholder="请输入UIA2端口" />
</el-form-item>
        <el-form-item label="UIA2服务包名:" prop="uia2ServerPackageName">
    <el-input v-model="formData.uia2ServerPackageName" :clearable="false" placeholder="请输入UIA2服务包名" />
</el-form-item>
        <el-form-item label="UIA2测试包名:" prop="uia2ServerTestPackageName">
    <el-input v-model="formData.uia2ServerTestPackageName" :clearable="false" placeholder="请输入UIA2测试包名" />
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
  createAndroidDeviceOptions,
  updateAndroidDeviceOptions,
  findAndroidDeviceOptions
} from '@/api/platform/androidDeviceOptions'

defineOptions({
    name: 'AndroidDeviceOptionsForm'
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
            serial: '',
            logOn: false,
            ignorePopup: false,
            adbServerHost: '',
            adbServerPort: 0,
            composite: false,
            uia2: false,
            uia2Ip: '',
            uia2Port: 0,
            uia2ServerPackageName: '',
            uia2ServerTestPackageName: '',
        })
// 验证规则
const rule = reactive({
               serial : [{
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
      const res = await findAndroidDeviceOptions({ ID: route.query.id })
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
               res = await createAndroidDeviceOptions(formData.value)
               break
             case 'update':
               res = await updateAndroidDeviceOptions(formData.value)
               break
             default:
               res = await createAndroidDeviceOptions(formData.value)
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
