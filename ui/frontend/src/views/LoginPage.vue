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
            <a-button class="captcha-btn" @click="refreshCaptcha">刷新</a-button>
          </div>
          <div class="captcha-img" @click="refreshCaptcha">
            <img v-if="form.picPath" :src="form.picPath" alt="captcha" />
          </div>
        </a-form-item>
        <a-space>
          <a-button type="primary" :loading="submitting" @click="submit">登录</a-button>
          <a-button @click="goSettings">设置域名</a-button>
        </a-space>
      </a-form>
    </a-card>
  </div>
</template>

<script setup>
import { Message } from '@arco-design/web-vue'
import { reactive, ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { captcha as captchaApi, login as loginApi } from '../services/appBridge'

const router = useRouter()
const submitting = ref(false)

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
    const res = await captchaApi()
    form.captcha = ''
    form.captchaId = res.captchaId || ''
    form.picPath = res.picPath || ''
    form.openCaptcha = Boolean(res.openCaptcha)
  } catch (e) {
    Message.error(e?.message || '获取验证码失败')
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

const goSettings = async () => {
  await router.push({ name: 'settings' })
}

onMounted(async () => {
  await refreshCaptcha()
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
