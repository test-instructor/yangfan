<template>
  <div class="page">
    <a-card class="card" title="设置扬帆自动化测试平台域名">
      <a-form :model="form" layout="vertical">
        <a-form-item field="baseURL" label="扬帆自动化测试平台域名（BaseURL）">
          <a-input v-model="form.baseURL" placeholder="https://xx.demo.com" />
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
import { getBaseURL, setBaseURL } from '../services/appBridge'

const route = useRoute()
const router = useRouter()

const form = reactive({
  baseURL: ''
})
const saving = ref(false)

const load = async () => {
  const { baseURL } = await getBaseURL()
  form.baseURL = baseURL
}

const save = async () => {
  saving.value = true
  try {
    await setBaseURL(form.baseURL)
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
