<template>
  <div class="page">
    <a-card class="card" title="设置">
      <a-form :model="form" layout="vertical">
        <a-form-item field="baseURL" label="扬帆自动化测试平台域名（BaseURL）">
          <a-input v-model="form.baseURL" placeholder="https://xx.demo.com" />
        </a-form-item>

        <a-divider />
        <h3>日志配置</h3>
        <a-form-item field="logLevel" label="日志级别">
          <a-select v-model="form.logLevel" placeholder="请选择日志级别">
            <a-option value="debug">调试 (Debug)</a-option>
            <a-option value="info">信息 (Info)</a-option>
            <a-option value="warn">警告 (Warning)</a-option>
            <a-option value="error">错误 (Error)</a-option>
            <a-option value="fatal">致命 (Fatal)</a-option>
          </a-select>
        </a-form-item>
        <a-form-item field="logPrefix" label="日志前缀">
          <a-input v-model="form.logPrefix" placeholder="[ https://github.com/test-instructor/yangfan/ui ]" />
        </a-form-item>
        <a-form-item field="logRetention" label="日志留存时间（天）">
          <a-input-number v-model="form.logRetention" :min="1" :max="365" />
        </a-form-item>

        <a-space>
          <a-button type="primary" :loading="saving" @click="save">保存</a-button>
          <a-button @click="goLogin">去登录</a-button>
        </a-space>
      </a-form>
    </a-card>
  </div>
</template>

<script setup>
import { Modal, Message } from '@arco-design/web-vue'
import { reactive, ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { getBaseURL, setBaseURL, getLogConfig, setLogConfig } from '../services/appBridge'

const route = useRoute()
const router = useRouter()

const form = reactive({
  baseURL: '',
  logLevel: 'info',
  logPrefix: '[ https://github.com/test-instructor/yangfan/ui ]',
  logRetention: 30
})
const saving = ref(false)

const load = async () => {
  try {
    const { baseURL } = await getBaseURL()
    form.baseURL = baseURL
    
    const logCfg = await getLogConfig()
    if (logCfg) {
      form.logLevel = logCfg.logLevel || 'info'
      form.logPrefix = logCfg.logPrefix || '[ https://github.com/test-instructor/yangfan/ui ]'
      form.logRetention = logCfg.logRetention || 30
    }
  } catch (e) {
    console.error('Failed to load config', e)
  }
}

const save = async () => {
  saving.value = true
  try {
    await setBaseURL(form.baseURL)
    await setLogConfig({
      level: form.logLevel,
      prefix: form.logPrefix,
      retention: form.logRetention
    })
    Message.success('保存成功')
    await router.replace({ name: 'login' })
  } catch (e) {
    Message.error(e?.message || '保存失败')
  } finally {
    saving.value = false
  }
}

const goLogin = async () => {
  await router.push({ name: 'login' })
}

onMounted(async () => {
  await load()
  if (route.query?.missing === '1') {
    Modal.warning({
      title: '需要先设置域名',
      content: '请先设置扬帆自动化测试平台域名（BaseURL），保存后再登录。'
    })
  }
})
</script>

<style scoped>
.page {
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 24px;
}
.card {
  width: 520px;
  max-width: 100%;
}
</style>
