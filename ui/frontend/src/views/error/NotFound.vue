<template>
  <div class="page">
    <a-result status="404" :title="t('error.notFoundTitle')" :subtitle="t('error.notFoundSubtitle')">
      <template #extra>
        <a-space>
          <a-button type="primary" @click="goHome">{{ t('error.goHome') }}</a-button>
          <a-button @click="goBack">{{ t('error.goBack') }}</a-button>
        </a-space>
      </template>
    </a-result>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'

const router = useRouter()
const { t } = useI18n()

const goHome = async () => {
  await router.replace({ path: '/route-init' })
}

const goBack = () => {
  router.back()
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

