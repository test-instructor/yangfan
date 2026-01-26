<template>
  <div :class="['settings-container', { page: !embedded }]">
    <a-card class="card" :bordered="false" :title="embedded ? '' : '设置'">
      <a-tabs default-active-key="general">
        <a-tab-pane key="general" title="常规设置">
          <a-form :model="form" layout="vertical" class="tab-content">
            <a-form-item field="baseURL" label="扬帆自动化测试平台域名（BaseURL）">
              <a-input v-model="form.baseURL" placeholder="https://xx.demo.com" />
            </a-form-item>
          </a-form>
        </a-tab-pane>

        <a-tab-pane key="logs" title="日志配置">
          <a-form :model="form" layout="vertical" class="tab-content">
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
          </a-form>
        </a-tab-pane>
      </a-tabs>

      <div class="actions">
        <a-space>
          <a-button type="primary" :loading="saving" @click="save">保存</a-button>
          <a-button v-if="!embedded" @click="goLogin">去登录</a-button>
        </a-space>
      </div>
    </a-card>
  </div>
</template>

<script setup>
import { Modal, Message } from '@arco-design/web-vue'
import { reactive, ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { getBaseURL, setBaseURL, getLogConfig, setLogConfig, clearAuth } from '../services/appBridge'

const props = defineProps({
  embedded: {
    type: Boolean,
    default: false
  }
})

const route = useRoute()
const router = useRouter()

const form = reactive({
  baseURL: '',
  logLevel: 'info',
  logPrefix: '[ https://github.com/test-instructor/yangfan/ui ]',
  logRetention: 30
})
const saving = ref(false)
const originalBaseURL = ref('')

const load = async () => {
  try {
    const { baseURL } = await getBaseURL()
    form.baseURL = baseURL
    originalBaseURL.value = baseURL
    
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
    let url = form.baseURL
    if (url.endsWith('/')) {
      url = url.slice(0, -1)
    }
    
    // Only check health if BaseURL changed
    const isBaseURLChanged = form.baseURL !== originalBaseURL.value
    
    if (isBaseURLChanged) {
        try {
        const healthRes = await fetch(`${url}/api/health`)
        const healthText = await healthRes.text()
        if (!healthRes.ok || (healthText !== 'ok' && !healthText.includes('ok'))) {
            throw new Error('Health check failed')
        }
        } catch (err) {
        throw new Error('域名连通性检查失败，请检查域名是否正确')
        }
    }

    await setBaseURL(form.baseURL)
    await setLogConfig({
      level: form.logLevel,
      prefix: form.logPrefix,
      retention: form.logRetention
    })

    if (isBaseURLChanged) {
        await clearAuth()
        Message.success('保存成功，域名已修改，请重新登录')
        if (props.embedded) {
             await router.replace({ name: 'login' })
        } else {
            // Standalone mode usually wants to go to login anyway
             await router.replace({ name: 'login' })
        }
    } else {
        Message.success('保存成功')
        if (!props.embedded) {
            // If standalone (e.g. from login page), go back to login
             await router.replace({ name: 'login' })
        }
    }
    
    // Update originalBaseURL
    originalBaseURL.value = form.baseURL

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
.settings-container {
  height: 100%;
  width: 100%;
}
.page {
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 24px;
}
.card {
  width: 100%;
  height: 100%;
}
.page .card {
  width: 520px;
  height: auto;
  max-width: 100%;
}
.tab-content {
  padding-top: 16px;
}
.actions {
  margin-top: 24px;
  text-align: right;
}
</style>
