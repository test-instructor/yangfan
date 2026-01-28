
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="设备名称:" prop="name">
    <el-input v-model="formData.name" :clearable="true" placeholder="请输入设备名称" />
</el-form-item>
        <el-form-item label="设备标识:" prop="udid">
    <el-input v-model="formData.udid" :clearable="false" placeholder="请输入设备标识" />
</el-form-item>
        <el-form-item label="无线连接:" prop="wireless">
    <el-switch v-model="formData.wireless" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="WDA端口:" prop="port">
    <el-input v-model.number="formData.port" :clearable="false" placeholder="请输入WDA端口" />
</el-form-item>
        <el-form-item label="WDA MJPEG 端口:" prop="mjpeg_port">
    <el-input v-model.number="formData.mjpeg_port" :clearable="false" placeholder="请输入WDA MJPEG 端口" />
</el-form-item>
        <el-form-item label="开启日志:" prop="log_on">
    <el-switch v-model="formData.log_on" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="忽略弹窗:" prop="ignore_popup">
    <el-switch v-model="formData.ignore_popup" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="重置主页:" prop="reset_home_on_startup">
    <el-switch v-model="formData.reset_home_on_startup" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="快照深度:" prop="snapshot_max_depth">
    <el-input v-model.number="formData.snapshot_max_depth" :clearable="false" placeholder="请输入快照深度" />
</el-form-item>
        <el-form-item label="接受按钮选择器:" prop="accept_alert_button_selector">
    <el-input v-model="formData.accept_alert_button_selector" :clearable="false" placeholder="请输入接受按钮选择器" />
</el-form-item>
        <el-form-item label="取消按钮选择器:" prop="dismiss_alert_button_selector">
    <el-input v-model="formData.dismiss_alert_button_selector" :clearable="false" placeholder="请输入取消按钮选择器" />
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
  createIOSDeviceOptions,
  updateIOSDeviceOptions,
  findIOSDeviceOptions
} from '@/api/platform/iosOptions'

defineOptions({
    name: 'IOSDeviceOptionsForm'
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
            udid: '',
            wireless: false,
            port: 0,
            mjpeg_port: 0,
            log_on: false,
            ignore_popup: false,
            reset_home_on_startup: false,
            snapshot_max_depth: 0,
            accept_alert_button_selector: '',
            dismiss_alert_button_selector: '',
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
      const res = await findIOSDeviceOptions({ ID: route.query.id })
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
               res = await createIOSDeviceOptions(formData.value)
               break
             case 'update':
               res = await updateIOSDeviceOptions(formData.value)
               break
             default:
               res = await createIOSDeviceOptions(formData.value)
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
