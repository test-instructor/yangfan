<template>
  <a-layout class="layout">
    <a-layout-sider collapsible breakpoint="xl">
      <div class="logo">Yangfan UI</div>
      <a-menu :selected-keys="[activeKey]" @menu-item-click="onMenuClick">
        <a-menu-item key="home">Dashboard</a-menu-item>
        <a-menu-item key="settings">Settings</a-menu-item>
      </a-menu>
    </a-layout-sider>
    <a-layout>
      <a-layout-header class="header">
        <a-space>
          <div class="title">Yangfan Automation Client</div>
        </a-space>
        <a-space>
          <a-select
            v-model="selectedProjectId"
            placeholder="选择项目"
            style="width: 200px"
            @change="onSwitch"
          >
            <a-option v-for="p in projectList" :key="p.id" :value="p.id">
              {{ p.name }}
            </a-option>
          </a-select>
          <a-select
            v-model="selectedAuthorityId"
            placeholder="选择角色"
            style="width: 200px"
            @change="onSwitch"
          >
            <a-option v-for="r in authorityList" :key="r.authorityId" :value="r.authorityId">
              {{ r.authorityName }}
            </a-option>
          </a-select>
        </a-space>
      </a-layout-header>
      <a-layout-content class="content">
        <a-card>
          <template #title>当前授权</template>
          <a-space direction="vertical" fill>
            <div>用户：{{ userInfo?.userName || '-' }}</div>
            <div>项目：{{ currentProject?.name || '-' }}</div>
            <div>角色：{{ currentAuthority?.authorityName || '-' }}</div>
          </a-space>
        </a-card>
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>

<script setup>
import { Message } from '@arco-design/web-vue'
import { computed, onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { getUserInfo as getUserInfoApi, setUserAuthority as setUserAuthorityApi } from '../services/appBridge'

const router = useRouter()
const activeKey = ref('home')

const userInfo = ref(null)
const selectedProjectId = ref(undefined)
const selectedAuthorityId = ref(undefined)
const switching = ref(false)

const projectList = computed(() => {
  const list = userInfo.value?.projectList
  return Array.isArray(list) ? list : []
})

const authorityList = computed(() => {
  const list = userInfo.value?.authorities
  return Array.isArray(list) ? list : []
})

const currentProject = computed(() => {
  return projectList.value.find((p) => p.id === selectedProjectId.value)
})

const currentAuthority = computed(() => {
  return authorityList.value.find((r) => r.authorityId === selectedAuthorityId.value)
})

const loadUserInfo = async () => {
  try {
    userInfo.value = await getUserInfoApi()
    selectedProjectId.value = userInfo.value?.projectId
    selectedAuthorityId.value = userInfo.value?.authorityId
  } catch (e) {
    Message.error(e?.message || '获取用户信息失败')
    await router.replace({ name: 'login' })
  }
}

const onSwitch = async () => {
  if (switching.value) return
  if (!selectedAuthorityId.value || !selectedProjectId.value) return
  switching.value = true
  try {
    userInfo.value = await setUserAuthorityApi({
      authorityId: selectedAuthorityId.value,
      projectId: selectedProjectId.value
    })
    selectedProjectId.value = userInfo.value?.projectId
    selectedAuthorityId.value = userInfo.value?.authorityId
    Message.success('切换成功')
  } catch (e) {
    Message.error(e?.message || '切换失败')
    await loadUserInfo()
  } finally {
    switching.value = false
  }
}

const onMenuClick = async (key) => {
  activeKey.value = key
  if (key === 'settings') {
    await router.push({ name: 'settings' })
  }
}

onMounted(async () => {
  await loadUserInfo()
})
</script>

<style scoped>
.layout {
  height: 100vh;
  width: 100vw;
}
.logo {
  height: 32px;
  margin: 16px;
  background: rgba(255, 255, 255, 0.2);
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--color-text-1);
  font-weight: bold;
}
.header {
  padding: 0 16px;
  background: var(--color-bg-2);
  border-bottom: 1px solid var(--color-border);
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
}
.title {
  font-weight: 600;
}
.content {
  padding: 24px;
}
</style>
