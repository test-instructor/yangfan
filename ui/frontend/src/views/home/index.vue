<template>
  <div class="dashboard-page">
    <a-card class="dashboard-card">
      <template #title>{{ t('home.currentAuth') }}</template>
      <a-space direction="vertical" fill>
        <div>{{ t('home.user') }}：{{ userInfo?.userName || '-' }}</div>
        <div>{{ t('home.project') }}：{{ currentProject?.name || '-' }}</div>
        <div>{{ t('home.role') }}：{{ currentAuthority?.authorityName || '-' }}</div>
      </a-space>
    </a-card>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()
const props = defineProps({
  userInfo: Object
})

const projectList = computed(() => {
  const list = props.userInfo?.projectList
  return Array.isArray(list) ? list : []
})

const authorityList = computed(() => {
  const list = props.userInfo?.authorities
  return Array.isArray(list) ? list : []
})

const currentProject = computed(() => {
  return projectList.value.find((p) => p.id === props.userInfo?.projectId)
})

const currentAuthority = computed(() => {
  return authorityList.value.find((r) => r.authorityId === props.userInfo?.authorityId)
})
</script>

<style scoped>
.dashboard-page {
  height: 100%;
  width: 100%;
}
.dashboard-card {
  height: 100%;
  width: 100%;
}
</style>
