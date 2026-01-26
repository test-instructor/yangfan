<template>
  <a-layout class="layout">
    <a-layout-sider collapsible breakpoint="xl">
      <div class="logo">扬帆</div>
      <a-menu :selected-keys="[activeKey]" @menu-item-click="onMenuClick">
        <a-menu-item key="home">
          <template #icon><IconHome /></template>
          Dashboard
        </a-menu-item>
        <a-menu-item key="settings">
          <template #icon><IconSettings /></template>
          Settings
        </a-menu-item>
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

          <a-dropdown @select="handleUserCommand">
            <div class="user-trigger">
              <a-avatar :size="32" style="margin-right: 8px; background-color: #3370ff">
                <img v-if="userInfo?.headerImg" :src="userInfo.headerImg" />
                <IconUser v-else />
              </a-avatar>
              <span class="username">{{ userInfo?.nickName || userInfo?.userName || 'User' }}</span>
              <IconDown />
            </div>
            <template #content>
              <a-doption value="person">
                <template #icon><IconUser /></template>
                个人信息
              </a-doption>
              <a-doption value="logout">
                <template #icon><IconPoweroff /></template>
                退出登录
              </a-doption>
            </template>
          </a-dropdown>
        </a-space>
      </a-layout-header>
      <a-layout-content class="content">
        <router-view :user-info="userInfo" :embedded="true" />
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>

<script setup>
import { Message } from '@arco-design/web-vue'
import { IconHome, IconSettings, IconUser, IconPoweroff, IconDown } from '@arco-design/web-vue/es/icon'
import { computed, onMounted, ref, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { getUserInfo as getUserInfoApi, setUserAuthority as setUserAuthorityApi, clearAuth } from '../services/appBridge'
import SettingsPage from './SettingsPage.vue'

const router = useRouter()
const route = useRoute()
const activeKey = ref('home')

// Sync menu with route changes
watch(() => route.name, (newVal) => {
  if (newVal === 'settings') {
    activeKey.value = 'settings'
  } else {
    activeKey.value = 'home'
  }
})

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
  if (key === 'home') {
    await router.push({ name: 'home' })
  } else if (key === 'settings') {
    await router.push({ name: 'settings' })
  }
}

const handleUserCommand = async (val) => {
  if (val === 'person') {
    await router.push({ name: 'person' })
  } else if (val === 'logout') {
    try {
      await clearAuth()
      Message.success('已退出登录')
      await router.replace({ name: 'login' })
    } catch (e) {
      Message.error('退出登录失败')
    }
  }
}

onMounted(async () => {
  await loadUserInfo()
  // Sync activeKey with current route
  if (route.name === 'settings') {
    activeKey.value = 'settings'
  } else {
    activeKey.value = 'home'
  }
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
  height: calc(100vh - 64px); /* Subtract header height */
  overflow: auto;
}
.user-trigger {
  display: flex;
  align-items: center;
  cursor: pointer;
  padding: 0 8px;
  height: 100%;
}
.user-trigger:hover {
  background-color: var(--color-fill-2);
  border-radius: 4px;
}
</style>
