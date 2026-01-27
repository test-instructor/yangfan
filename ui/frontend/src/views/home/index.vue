<template>
  <div class="dashboard-page">
    <a-card class="dashboard-card">
      <template #title>当前授权</template>
      <a-space direction="vertical" fill>
        <div>用户：{{ userInfo?.userName || '-' }}</div>
        <div>项目：{{ currentProject?.name || '-' }}</div>
        <div>角色：{{ currentAuthority?.authorityName || '-' }}</div>
      </a-space>
    </a-card>
  </div>
</template>

<script setup>
import { computed } from 'vue'

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

