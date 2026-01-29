<template>
  <a-modal
    :visible="visible"
    :closable="false"
    :mask-closable="false"
    :footer="false"
    :title="t('globalBaseURL.title')"
    width="500px"
  >
    <div style="margin-bottom: 20px;">
      <a-alert type="warning" style="margin-bottom: 16px;">
        {{ t('globalBaseURL.alert') }}
      </a-alert>
      <a-form :model="form" layout="vertical">
        <a-form-item field="baseURL" :label="t('globalBaseURL.label')" required>
          <a-input 
            v-model="form.baseURL" 
            :placeholder="t('globalBaseURL.placeholder')" 
            allow-clear
            @press-enter="save"
          />
        </a-form-item>
      </a-form>
      <div style="text-align: right; margin-top: 24px;">
        <a-button type="primary" :loading="saving" @click="save" long>
          {{ t('globalBaseURL.saveAndEnter') }}
        </a-button>
      </div>
    </div>
  </a-modal>
</template>

<script setup>
import { reactive, ref } from 'vue'
import { Message } from '@arco-design/web-vue'
import { useI18n } from 'vue-i18n'
import { checkBaseURLConnectivity, setBaseURL, clearAuth } from '../services/appBridge'

const props = defineProps({
  visible: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['success'])
const { t } = useI18n()

const form = reactive({
  baseURL: ''
})
const saving = ref(false)

const save = async () => {
  if (!form.baseURL) {
    Message.warning(t('globalBaseURL.enterDomain'))
    return
  }

  saving.value = true
  try {
    const res = await checkBaseURLConnectivity(form.baseURL)
    const normalized = res?.baseURL || ''
    await setBaseURL(normalized || form.baseURL)
    // Clear auth just in case, though this is fresh setup
    await clearAuth()
    if (normalized) {
      form.baseURL = normalized
    }
    
    Message.success(t('globalBaseURL.configSuccess'))
    emit('success')
  } catch (e) {
    Message.error(e?.message || t('globalBaseURL.saveError'))
  } finally {
    saving.value = false
  }
}
</script>
