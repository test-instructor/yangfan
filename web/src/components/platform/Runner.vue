<template>
  <div>
    <el-button type="primary" link icon="video-play" @click="handleOpen">运行</el-button>
    <el-dialog
      v-model="visible"
      title="运行任务"
      width="520px"
      append-to-body
      destroy-on-close
    >
      <div class="runner-container">
        <el-form label-position="right" label-width="80px">
          <el-form-item label="运行模式:">
            <el-select v-model="form.run_mode" placeholder="请选择运行模式" style="width: 100%">
              <el-option
                v-for="item in runModes"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
          </el-form-item>

          <el-form-item label="运行配置:">
            <RunConfigSelector v-model="form.config_id" width="100%" />
          </el-form-item>

          <el-form-item label="运行环境:">
            <EnvSelector v-model="form.env_id" width="100%" />
          </el-form-item>

          <el-divider content-position="left">消息通知</el-divider>
          <el-form-item label="发送消息:">
            <el-switch v-model="form.notify_enabled" active-color="#13ce66" inactive-color="#ff4949" active-text="是"
                       inactive-text="否" clearable></el-switch>
          </el-form-item>
          <el-form-item label="发送规则:">
            <el-select v-model="form.notify_rule" placeholder="请选择发送规则" style="width: 100%"
                       :disabled="!form.notify_enabled">
              <el-option label="总是发送" value="always" />
              <el-option label="仅成功发送" value="success" />
              <el-option label="仅失败发送" value="fail" />
            </el-select>
          </el-form-item>
          <el-form-item label="通知通道:">
            <el-select v-model="form.notify_channel_ids" multiple collapse-tags collapse-tags-tooltip filterable clearable
                       style="width: 100%" placeholder="选择通道" :loading="notifyChannelLoading"
                       :disabled="!form.notify_enabled">
              <el-option
                v-for="item in notifyChannelOptions"
                :key="item.ID"
                :label="`${providerLabel(item.provider)}-${item.name}`"
                :value="item.ID"
              />
            </el-select>
          </el-form-item>

          <template v-if="runResult.report_id">
            <el-form-item label="报告:">
              <div class="flex items-center justify-between w-full">
                <div class="text-sm">
                  <span class="mr-2">report_id: {{ runResult.report_id }}</span>
                  <span v-if="runResult.task_id" class="text-gray-500">task_id: {{ runResult.task_id }}</span>
                </div>
                <el-button type="primary" link @click="openReport">打开报告</el-button>
              </div>
            </el-form-item>
            <el-form-item label="发送状态:">
              <el-select
                v-model="selectedChannelIds"
                multiple
                collapse-tags
                collapse-tags-tooltip
                style="width: 100%"
                :loading="notifyLoading"
                placeholder="等待通知结果"
              >
                <el-option
                  v-for="item in notifyItems"
                  :key="item.channelId"
                  :label="item.display"
                  :value="item.channelId"
                />
              </el-select>
            </el-form-item>
          </template>
        </el-form>
      </div>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="visible = false">取消</el-button>
          <el-button type="primary" :loading="loading" @click="handleRun">执行</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onUnmounted, watch } from 'vue'
import { ElMessage } from 'element-plus'
import EnvSelector from '@/components/platform/env.vue'
import RunConfigSelector from '@/components/platform/runConfig.vue'
import { runTask } from '@/api/run.js'
import { getAutoReportNotifyStatus } from '@/api/projectmgr/reportNotify'
import { getReportNotifyChannelList } from '@/api/projectmgr/reportNotify'
import { useRouter } from 'vue-router'

const props = defineProps({
  case_type: {
    type: String,
    required: true
  },
  case_id: {
    type: Number,
    required: true
  }
})

const visible = ref(false)
const loading = ref(false)
const notifyLoading = ref(false)
const notifyChannelLoading = ref(false)
const notifyChannelOptions = ref([])
const notifyItems = ref([])
const selectedChannelIds = ref([])
const runResult = reactive({
  task_id: '',
  report_id: 0,
})
const router = useRouter()
let notifyTimer = null

const form = reactive({
  run_mode: '调试模式', // 默认调试
  config_id: null,
  env_id: null,
  notify_enabled: false,
  notify_rule: 'always',
  notify_channel_ids: []
})

// 运行模式选项
const runModes = [
  { label: '调试', value: '调试模式' },
  { label: '后台运行', value: '后台运行' },
  { label: 'CI', value: 'CI' }
]

const handleOpen = () => {
  // 每次打开可以重置或保持状态，这里保持状态
  loadNotifyChannels()
  visible.value = true
}

const resetRunState = () => {
  runResult.task_id = ''
  runResult.report_id = 0
  notifyItems.value = []
  selectedChannelIds.value = []
  notifyLoading.value = false
  if (notifyTimer) {
    clearInterval(notifyTimer)
    notifyTimer = null
  }
}

const providerLabel = (p) => {
  if (p === 'feishu') return '飞书'
  if (p === 'dingtalk') return '钉钉'
  if (p === 'wecom') return '企业微信'
  return p || ''
}

const loadNotifyChannels = async () => {
  notifyChannelLoading.value = true
  try {
    const res = await getReportNotifyChannelList({ page: 1, pageSize: 1000 })
    if (res.code === 0) {
      notifyChannelOptions.value = res.data?.list || []
    }
  } finally {
    notifyChannelLoading.value = false
  }
}

const openReport = () => {
  if (!runResult.report_id) return
  router.push({ name: 'auto-report-detail', params: { id: runResult.report_id } })
}

const fetchNotifyStatus = async () => {
  if (!runResult.report_id) return
  notifyLoading.value = true
  try {
    const res = await getAutoReportNotifyStatus({ reportId: runResult.report_id })
    if (res.code === 0) {
      const items = res.data?.items || []
      notifyItems.value = items
      if (!selectedChannelIds.value.length) {
        selectedChannelIds.value = items.map((i) => i.channelId)
      }
      const reportStatus = res.data?.reportStatus
      const isFinished = reportStatus === 2 || reportStatus === 3
      const hasPending = items.some((i) => i.state === '待发送')
      if (isFinished && !hasPending && notifyTimer) {
        clearInterval(notifyTimer)
        notifyTimer = null
      }
    }
  } finally {
    notifyLoading.value = false
  }
}

const handleRun = async () => {
  if (!form.run_mode) {
    ElMessage.warning('请选择运行模式')
    return
  }
  
  // 某些模式下可能必须选择环境或配置，视具体业务而定，这里暂不做强制校验，交给后端或后续需求
  // 但通常调试模式需要环境
  if (!form.env_id) {
     ElMessage.warning('请选择运行环境')
     return
  }

  loading.value = true
  try {
    const payload = {
      case_type: props.case_type,
      case_id: props.case_id,
      config_id: form.config_id || 0,
      env_id: form.env_id || 0,
      run_mode: form.run_mode,
      notify_enabled: !!form.notify_enabled,
      notify_rule: form.notify_rule || '',
      notify_channel_ids: Array.isArray(form.notify_channel_ids) ? form.notify_channel_ids : []
    }
    
    const res = await runTask(payload)
    if (res.code === 0) {
      ElMessage.success('运行任务已提交')
      runResult.task_id = res.data?.task_id || ''
      runResult.report_id = res.data?.report_id || 0
      notifyItems.value = []
      selectedChannelIds.value = []
      await fetchNotifyStatus()
      if (notifyTimer) clearInterval(notifyTimer)
      notifyTimer = setInterval(fetchNotifyStatus, 3000)
    } else {
      ElMessage.error(res.msg || '运行失败')
    }
  } catch (error) {
    console.error(error)
    ElMessage.error('运行异常')
  } finally {
    loading.value = false
  }
}

watch(
  () => visible.value,
  (val) => {
    if (!val) resetRunState()
  }
)

onUnmounted(() => {
  resetRunState()
})
</script>

<style scoped>
.runner-container {
  padding: 10px;
}
.dialog-footer {
  text-align: right;
}
</style>
