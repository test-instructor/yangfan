<template>
  <div v-if="!isReady" class="app-loading">
    <!-- Loading state if needed -->
  </div>
  <template v-else>
    <GlobalBaseURLModal :visible="showSetup" @success="onSetupSuccess" />
    <router-view v-if="!showSetup" />
  </template>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getBaseURL } from './services/appBridge'
import GlobalBaseURLModal from './components/GlobalBaseURLModal.vue'

const isReady = ref(false)
const showSetup = ref(false)

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
  justify-content: center;
  align-items: center;
}
</style>
