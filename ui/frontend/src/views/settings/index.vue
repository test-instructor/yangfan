<template>
  <div :class="['settings-container', { page: !embedded }]">
    <a-card class="card" :bordered="false" :title="embedded ? '' : t('settings.title')">
      <a-tabs default-active-key="general">
        <a-tab-pane key="general" :title="t('settings.general')">
          <a-form :model="form" layout="vertical" class="tab-content">
            <a-form-item field="baseURL" :label="t('settings.baseURL')">
              <a-input v-model="form.baseURL" placeholder="https://xx.demo.com" />
            </a-form-item>
            <a-form-item field="theme" :label="t('settings.theme')">
              <a-radio-group v-model="form.theme" type="button" @change="onThemeChange">
                <a-radio value="light">{{ t('settings.themeLight') }}</a-radio>
                <a-radio value="dark">{{ t('settings.themeDark') }}</a-radio>
              </a-radio-group>
            </a-form-item>
          </a-form>
        </a-tab-pane>

        <a-tab-pane key="logs" :title="t('settings.logConfig')">
          <a-form :model="form" layout="vertical" class="tab-content">
            <a-form-item field="logLevel" :label="t('settings.logLevel')">
              <a-select v-model="form.logLevel" :placeholder="t('settings.selectLogLevel')">
                <a-option value="debug">{{ t('settings.levels.debug') }}</a-option>
                <a-option value="info">{{ t('settings.levels.info') }}</a-option>
                <a-option value="warn">{{ t('settings.levels.warn') }}</a-option>
                <a-option value="error">{{ t('settings.levels.error') }}</a-option>
                <a-option value="fatal">{{ t('settings.levels.fatal') }}</a-option>
              </a-select>
            </a-form-item>
            <a-form-item field="logPrefix" :label="t('settings.logPrefix')">
              <a-input v-model="form.logPrefix" placeholder="[ https://github.com/test-instructor/yangfan/ui ]" />
            </a-form-item>
            <a-form-item field="logRetention" :label="t('settings.logRetention')">
              <a-input-number v-model="form.logRetention" :min="1" :max="365" />
            </a-form-item>
          </a-form>
        </a-tab-pane>
      </a-tabs>

      <div class="actions">
        <a-space>
          <a-button type="primary" :loading="saving" @click="save">{{ t('settings.save') }}</a-button>
          <a-button v-if="!embedded" @click="goLogin">{{ t('settings.goLogin') }}</a-button>
        </a-space>
      </div>
    </a-card>
  </div>
</template>

<script setup>
import { Modal, Message } from '@arco-design/web-vue'
import { reactive, ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { getStoredTheme, setTheme } from '../../utils/theme'
import {
  getBaseURL,
  checkBaseURLConnectivity,
  setBaseURL,
  getLogConfig,
  setLogConfig,
  clearAuth
} from '../../services/appBridge'

const props = defineProps({
  embedded: {
    type: Boolean,
    default: false
  }
})

const route = useRoute()
const router = useRouter()
const { t } = useI18n()

const form = reactive({
  baseURL: '',
  theme: 'light',
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
    form.theme = getStoredTheme()

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

const onThemeChange = (val) => {
  const next = typeof val === 'string' ? val : form.theme
  setTheme(next)
}

const save = async () => {
  saving.value = true
  try {
    const normalizedInput = String(form.baseURL || '').trim().replace(/\/+$/, '')
    const isBaseURLChanged = normalizedInput !== originalBaseURL.value

    let normalized = normalizedInput
    if (isBaseURLChanged && normalizedInput) {
      const res = await checkBaseURLConnectivity(form.baseURL)
      normalized = res?.baseURL || normalizedInput
    }

    if (normalized) {
      await setBaseURL(normalized)
      form.baseURL = normalized
    } else {
      await setBaseURL(form.baseURL)
    }
    await setLogConfig({
      level: form.logLevel,
      prefix: form.logPrefix,
      retention: form.logRetention
    })

    if (isBaseURLChanged) {
      await clearAuth()
      Message.success(t('settings.saveSuccessRelogin'))
      if (props.embedded) {
        await router.replace({ name: 'login' })
      } else {
        await router.replace({ name: 'login' })
      }
    } else {
      Message.success(t('settings.saveSuccess'))
      if (!props.embedded) {
        await router.replace({ name: 'login' })
      }
    }

    originalBaseURL.value = form.baseURL
  } catch (e) {
    Message.error(e?.message || t('settings.saveError'))
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
      title: t('settings.needSetDomain'),
      content: t('settings.setDomainTip')
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
