<template>
  <a-modal
    :visible="visible"
    :closable="false"
    :mask-closable="false"
    :footer="false"
    title="设置服务器地址"
    width="500px"
  >
    <div style="margin-bottom: 20px;">
      <a-alert type="warning" style="margin-bottom: 16px;">
        检测到未配置服务器地址（BaseURL），请先配置扬帆自动化测试平台域名。
      </a-alert>
      <a-form :model="form" layout="vertical">
        <a-form-item field="baseURL" label="扬帆自动化测试平台域名（BaseURL）" required>
          <a-input 
            v-model="form.baseURL" 
            placeholder="例如: https://yangfan.demo.com" 
            allow-clear
            @press-enter="save"
          />
        </a-form-item>
      </a-form>
      <div style="text-align: right; margin-top: 24px;">
        <a-button type="primary" :loading="saving" @click="save" long>
          保存并进入
        </a-button>
      </div>
    </div>
  </a-modal>
</template>

<script setup>
import { reactive, ref } from 'vue'
import { Message } from '@arco-design/web-vue'
import { setBaseURL, clearAuth } from '../services/appBridge'

const props = defineProps({
  visible: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['success'])

const form = reactive({
  baseURL: ''
})
const saving = ref(false)

const save = async () => {
  if (!form.baseURL) {
    Message.warning('请输入域名')
    return
  }

  saving.value = true
  try {
    let url = form.baseURL.trim()
    if (url.endsWith('/')) {
      url = url.slice(0, -1)
    }

    // Health check
    try {
      const healthRes = await fetch(`${url}/api/health`)
      const healthText = await healthRes.text()
      if (!healthRes.ok || (healthText !== 'ok' && !healthText.includes('ok'))) {
        throw new Error('Health check failed')
      }
    } catch (err) {
      throw new Error('域名连通性检查失败，请检查域名是否正确')
    }

    await setBaseURL(url)
    // Clear auth just in case, though this is fresh setup
    await clearAuth()
    
    Message.success('配置成功')
    emit('success')
  } catch (e) {
    Message.error(e?.message || '保存失败')
  } finally {
    saving.value = false
  }
}
</script>
