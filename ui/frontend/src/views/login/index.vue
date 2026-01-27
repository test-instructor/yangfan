<template>
  <div class="page">
    <a-card class="card" title="登录">
      <a-form :model="form" layout="vertical">
        <a-form-item field="username" label="用户名">
          <a-input v-model="form.username" autocomplete="username" />
        </a-form-item>
        <a-form-item field="password" label="密码">
          <a-input-password v-model="form.password" autocomplete="current-password" />
        </a-form-item>
        <a-form-item v-if="form.openCaptcha" field="captcha" label="验证码">
          <div class="captcha-row">
            <a-input v-model="form.captcha" />
          </div>
          <div class="captcha-img" @click="refreshCaptcha" style="margin-left: 50px">
            <img v-if="form.picPath" :src="form.picPath" alt="captcha" />
          </div>
        </a-form-item>

        <a-space>
          <a-button type="primary" :loading="submitting" @click="submit">登录</a-button>
          <a-button @click="openSettings">设置域名</a-button>
        </a-space>
      </a-form>
    </a-card>

    <a-modal v-model:visible="showSettingsModal" title="设置域名" @ok="saveSettings" :ok-loading="savingSettings">
      <a-form :model="settingsForm" layout="vertical">
        <a-form-item field="baseURL" label="扬帆自动化测试平台域名（BaseURL）">
          <a-input v-model="settingsForm.baseURL" placeholder="https://xx.demo.com" />
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script setup>
import { Message } from '@arco-design/web-vue'
import { reactive, ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import {
  captcha as captchaApi,
  login as loginApi,
  getBaseURL,
  checkBaseURLConnectivity,
  setBaseURL,
  clearAuth
} from '../../services/appBridge'

const router = useRouter()
const submitting = ref(false)

const showSettingsModal = ref(false)
const settingsForm = reactive({
  baseURL: ''
})
const savingSettings = ref(false)
const originalBaseURL = ref('')

const form = reactive({
  username: 'admin',
  password: '',
  captcha: '',
  captchaId: '',
  picPath: '',
  openCaptcha: false
})

const refreshCaptcha = async () => {
  try {
    const res = (await captchaApi()) || {}

    form.captcha = ''
    form.captchaId = res.captchaId || ''
    form.picPath = res.picPath || ''
    form.openCaptcha = Boolean(res.openCaptcha)
  } catch (e) {
    const errMsg = e?.message || '获取验证码失败'
    Message.error(errMsg)
  }
}

const submit = async () => {
  submitting.value = true
  try {
    await loginApi({
      username: form.username,
      password: form.password,
      captcha: form.captcha,
      captchaId: form.captchaId
    })
    Message.success('登录成功')
    await router.replace({ name: 'home' })
  } catch (e) {
    Message.error(e?.message || '登录失败')
    await refreshCaptcha()
  } finally {
    submitting.value = false
  }
}

const openSettings = async () => {
  try {
    const { baseURL } = await getBaseURL()
    settingsForm.baseURL = baseURL
    originalBaseURL.value = baseURL
    showSettingsModal.value = true
  } catch (e) {
    Message.error('获取配置失败')
  }
}

const saveSettings = async () => {
  if (!settingsForm.baseURL) {
    Message.warning('请输入域名')
    return
  }

  savingSettings.value = true
  try {
    const normalizedInput = String(settingsForm.baseURL || '').trim().replace(/\/+$/, '')
    const isBaseURLChanged = normalizedInput !== originalBaseURL.value

    let normalized = normalizedInput
    if (isBaseURLChanged && normalizedInput) {
      const res = await checkBaseURLConnectivity(settingsForm.baseURL)
      normalized = res?.baseURL || normalizedInput
    }

    if (normalized) {
      await setBaseURL(normalized)
      settingsForm.baseURL = normalized
    } else {
      await setBaseURL(settingsForm.baseURL)
    }

    if (isBaseURLChanged) {
      await clearAuth()
    }

    Message.success('保存成功')
    showSettingsModal.value = false

    await refreshCaptcha()
  } catch (e) {
    Message.error(e?.message || '保存失败')
  } finally {
    savingSettings.value = false
  }
}

onMounted(async () => {
  const { baseURL } = await getBaseURL()
  if (!baseURL) {
    settingsForm.baseURL = ''
    originalBaseURL.value = ''
    showSettingsModal.value = true
    Message.info('请先配置服务域名')
  } else {
    await refreshCaptcha()
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
.captcha-row {
  display: flex;
  gap: 12px;
}
.captcha-btn {
  flex: none;
}
.captcha-img {
  margin-top: 12px;
  width: 160px;
  height: 60px;
  cursor: pointer;
}
.captcha-img img {
  width: 100%;
  height: 100%;
  object-fit: contain;
}
</style>

