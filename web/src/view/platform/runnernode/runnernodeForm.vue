
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="节点ID:" prop="nodeId">
    <el-input v-model="formData.nodeId" :clearable="false" placeholder="请输入节点ID" />
</el-form-item>
        <el-form-item label="别名:" prop="alias">
    <el-input v-model="formData.alias" :clearable="false" placeholder="请输入别名" />
</el-form-item>
        <el-form-item label="IP:" prop="ip">
    <el-input v-model="formData.ip" :clearable="false" placeholder="请输入IP" />
</el-form-item>
        <el-form-item label="端口:" prop="port">
    <el-input v-model.number="formData.port" :clearable="false" placeholder="请输入端口" />
</el-form-item>
        <el-form-item label="状态:" prop="status">
    <el-input v-model.number="formData.status" :clearable="false" placeholder="请输入状态" />
</el-form-item>
        <el-form-item label="最后心跳:" prop="lastHeartbeat">
    <el-date-picker v-model="formData.lastHeartbeat" type="date" style="width:100%" placeholder="选择日期" :clearable="false" />
</el-form-item>
        <el-form-item label="注册时间:" prop="createTime">
    <el-date-picker v-model="formData.createTime" type="date" style="width:100%" placeholder="选择日期" :clearable="false" />
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
  createRunnerNode,
  updateRunnerNode,
  findRunnerNode
} from '@/api/platform/runnernode'

defineOptions({
    name: 'RunnerNodeForm'
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
            nodeId: '',
            alias: '',
            ip: '',
            port: 0,
            status: 0,
            lastHeartbeat: new Date(),
            createTime: new Date(),
        })
// 验证规则
const rule = reactive({
               nodeId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               alias : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               ip : [{
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
      const res = await findRunnerNode({ ID: route.query.id })
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
               res = await createRunnerNode(formData.value)
               break
             case 'update':
               res = await updateRunnerNode(formData.value)
               break
             default:
               res = await createRunnerNode(formData.value)
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
