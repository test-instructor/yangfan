<template>
  <div class="page">
    <a-result status="warning" :title="t('error.componentMissingTitle')" :subtitle="t('error.componentMissingSubtitle')">
      <template #extra>
        <a-space direction="vertical" fill>
          <a-alert type="warning" :title="t('error.componentKey')" :content="missingKeyText" />
          <a-space>
            <a-button type="primary" @click="goHome">{{ t('error.goHome') }}</a-button>
            <a-button @click="reload">{{ t('error.retry') }}</a-button>
          </a-space>
        </a-space>
      </template>
    </a-result>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'

const route = useRoute()
const router = useRouter()
const { t } = useI18n()

const missingKeyText = computed(() => {
  return String(route.meta?.missingComponentKey || route.meta?.componentKey || '')
})

const goHome = async () => {
  await router.replace({ path: '/route-init' })
}

const reload = () => {
  window.location.reload()
}
</script>

<style scoped>
.page {
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 24px;
}
</style>

