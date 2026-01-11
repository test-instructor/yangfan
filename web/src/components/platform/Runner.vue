<template>
  <div>
    <el-button type="primary" link icon="video-play" @click="handleOpen">运行</el-button>
    <el-dialog
      v-model="visible"
      title="运行任务"
      width="400px"
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
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'
import EnvSelector from '@/components/platform/env.vue'
import RunConfigSelector from '@/components/platform/runConfig.vue'
import { runTask } from '@/api/run.js'

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

const form = reactive({
  run_mode: '调试模式', // 默认调试
  config_id: null,
  env_id: null
})

// 运行模式选项
const runModes = [
  { label: '调试', value: '调试模式' },
  { label: '后台运行', value: '后台运行' },
  { label: 'CI', value: 'CI' }
]

const handleOpen = () => {
  // 每次打开可以重置或保持状态，这里保持状态
  visible.value = true
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
      run_mode: form.run_mode
    }
    
    const res = await runTask(payload)
    if (res.code === 0) {
      ElMessage.success('运行任务已提交')
      visible.value = false
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
</script>

<style scoped>
.runner-container {
  padding: 10px;
}
.dialog-footer {
  text-align: right;
}
</style>
