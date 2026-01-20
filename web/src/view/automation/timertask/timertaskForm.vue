
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="任务名称:" prop="name">
    <el-input v-model="formData.name" :clearable="false" placeholder="请输入任务名称" />
</el-form-item>
        <el-form-item label="运行时间:" prop="runTime">
    <el-input v-model="formData.runTime" :clearable="false" placeholder="请输入运行时间" />
</el-form-item>
        <el-form-item label="下次运行时间:" prop="nextRunTime">
    <el-date-picker v-model="formData.nextRunTime" type="date" style="width:100%" placeholder="选择日期" :clearable="false" />
</el-form-item>
        <el-form-item label="运行状态:" prop="status">
    <el-switch v-model="formData.status" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="运行次数:" prop="runNumber">
    <el-input v-model.number="formData.runNumber" :clearable="false" placeholder="请输入运行次数" />
</el-form-item>
        <el-form-item label="失败停止:" prop="failfast">
    <el-switch v-model="formData.failfast" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="运行配置:" prop="configName">
    <el-input v-model="formData.configName" :clearable="true" placeholder="请输入运行配置" />
</el-form-item>
        <el-form-item label="标签:" prop="tag">
    <el-select v-model="formData.tag" multiple filterable placeholder="选择标签">
      <el-option
        v-for="opt in tagOptions"
        :key="opt.ID || opt.id"
        :label="opt.name"
        :value="opt.ID || opt.id"
      />
    </el-select>
</el-form-item>
        <el-form-item label="环境名称:" prop="envName">
    <el-input v-model="formData.envName" :clearable="false" placeholder="请输入环境名称" />
</el-form-item>
        <el-form-item label="消息名称:" prop="messageName">
    <el-input v-model="formData.messageName" :clearable="false" placeholder="请输入消息名称" />
</el-form-item>
        <el-form-item label="备注:" prop="describe">
    <el-input v-model="formData.describe" :clearable="false" placeholder="请输入备注" />
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
  createTimerTask,
  updateTimerTask,
  findTimerTask
} from '@/api/automation/timertask'
import { getTagList } from '@/api/automation/tag'

defineOptions({
    name: 'TimerTaskForm'
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
            runTime: '',
            nextRunTime: new Date(),
            status: false,
            failfast: false,
            runNumber: 0,
            configName: '',
            tag: [],
            envName: '',
            messageName: '',
            describe: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findTimerTask({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        formData.value.tag = normalizeTagArray(res.data.tag)
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
const tagOptions = ref([])
const tagNameMap = ref({})

const loadTagOptions = async () => {
  const res = await getTagList({ page: 1, pageSize: 1000 })
  if (res.code === 0) {
    const list = res.data.list || []
    tagOptions.value = list
    const map = {}
    list.forEach(item => { map[item.ID || item.id] = item.name })
    tagNameMap.value = map
  }
}

loadTagOptions()

const normalizeTagArray = (raw) => {
  if (!raw) return []
  if (Array.isArray(raw)) {
    const ids = []
    raw.forEach(it => {
      if (typeof it === 'number') ids.push(it)
      else if (typeof it === 'string') {
        const entry = Object.entries(tagNameMap.value).find(([, name]) => name === it)
        if (entry) ids.push(Number(entry[0]))
      } else if (it && typeof it === 'object') {
        if (it.ID) ids.push(it.ID)
        else if (it.id) ids.push(it.id)
      }
    })
    return ids
  }
  return []
}
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createTimerTask(formData.value)
               break
             case 'update':
               res = await updateTimerTask(formData.value)
               break
             default:
               res = await createTimerTask(formData.value)
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
