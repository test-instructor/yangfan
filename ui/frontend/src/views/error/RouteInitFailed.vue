<template>
  <div class="page">
    <a-result status="error" :title="t('error.routeInitFailedTitle')" :subtitle="subtitle">
      <template #extra>
        <a-space>
          <a-button type="primary" :loading="retrying" @click="retry">{{ t('error.retry') }}</a-button>
          <a-button @click="reload">{{ t('error.reload') }}</a-button>
        </a-space>
      </template>
    </a-result>
  </div>
</template>

<script setup>
import { computed, ref } from 'vue'
import { useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { ensureUINodeDynamicRoutes, getFirstLeafMenuPath } from '../../router/uiNodeDynamicRoutes'
import { uiNodeMenus, uiNodeRouteInitError } from '../../router/uiNodeState'

const router = useRouter()
const { t } = useI18n()
const retrying = ref(false)

const subtitle = computed(() => {
  const msg = uiNodeRouteInitError.value?.message || ''
  return msg ? `${t('error.routeInitFailedSubtitle')} (${msg})` : t('error.routeInitFailedSubtitle')
})

const retry = async () => {
  retrying.value = true
  try {
    const ok = await ensureUINodeDynamicRoutes(router)
    if (ok) {
      const target = getFirstLeafMenuPath(uiNodeMenus.value)
      await router.replace({ path: target })
    }
  } finally {
    retrying.value = false
  }
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

