<template>
  <div v-if="!isReady" class="app-loading">
    <a-spin dot />
    <div class="loading-text">扬帆自动化测试平台</div>
  </div>
  <template v-else>
    <GlobalBaseURLModal :visible="showSetup" @success="onSetupSuccess" />
    <router-view v-if="!showSetup" />
  </template>
  
  <!-- 全局错误提示 -->
  <div id="global-error-toast" v-if="showError" class="error-toast">
    <a-alert :type="errorType" closable @close="hideError">
      {{ errorMessage }}
    </a-alert>
  </div>
</template>

<script setup>
import { ref, onMounted, provide } from 'vue'
import { getBaseURL } from './services/appBridge'
import GlobalBaseURLModal from './components/GlobalBaseURLModal.vue'

const isReady = ref(false)
const showSetup = ref(false)
const showError = ref(false)
const errorMessage = ref('')
const errorType = ref('error')

// 全局错误处理函数
const showErrorMessage = (message, type = 'error') => {
  errorMessage.value = message
  errorType.value = type
  showError.value = true
  
  // 5秒后自动隐藏
  setTimeout(() => {
    hideError()
  }, 5000)
}

const hideError = () => {
  showError.value = false
  errorMessage.value = ''
}

// 提供给子组件使用
provide('showError', showErrorMessage)

const onSetupSuccess = () => {
  showSetup.value = false
}

onMounted(async () => {
  try {
    const { ok } = await getBaseURL()
    if (!ok) {
      showSetup.value = true
    }
  } catch (e) {
    console.error('Failed to check base URL', e)
    showErrorMessage('初始化失败，请检查配置')
  } finally {
    isReady.value = true
  }
})
</script>

<style>
html, body, #app {
  height: 100%;
  margin: 0;
  padding: 0;
}

.app-loading {
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
}

.loading-text {
  margin-top: 20px;
  font-size: 18px;
  font-weight: 500;
}

.error-toast {
  position: fixed;
  top: 20px;
  right: 20px;
  z-index: 9999;
  max-width: 400px;
  animation: slideIn 0.3s ease-out;
}

@keyframes slideIn {
  from {
    transform: translateX(100%);
    opacity: 0;
  }
  to {
    transform: translateX(0);
    opacity: 1;
  }
}
</style>
