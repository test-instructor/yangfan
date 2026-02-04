<template>
  <a-layout class="layout">
    <a-layout-sider collapsible breakpoint="xl" class="sider" v-model:collapsed="collapsed">
      <div class="logo">
        <img :src="logo" alt="logo" class="logo-img" />
        <div class="logo-text" v-if="!collapsed">{{ t('common.appName') }}</div>
      </div>
      <a-menu :selected-keys="[activeKey]" @menu-item-click="onMenuClick">
        <a-menu-item key="home">
          <template #icon><IconHome /></template>
          {{ t('common.dashboard') }}
        </a-menu-item>
        <a-menu-item key="settings">
          <template #icon><IconSettings /></template>
          {{ t('common.settings') }}
        </a-menu-item>
      </a-menu>
    </a-layout-sider>
    <a-layout>
      <a-layout-header class="header">
        <a-space>
          <a-breadcrumb :style="{ margin: '16px 0' }">
            <a-breadcrumb-item>
              <IconHome />
            </a-breadcrumb-item>
            <a-breadcrumb-item v-for="(item, index) in breadcrumbs" :key="index">
              {{ item.meta?.title ? t(item.meta.title) : item.name }}
            </a-breadcrumb-item>
          </a-breadcrumb>
        </a-space>
        <ul class="right-side">
          <li>
            <a-dropdown @select="handleProjectSwitch">
              <a-button class="nav-btn" type="text">
                <template #icon>
                  <IconApps />
                </template>
                {{ currentProjectName || t('common.selectProject') }}
                <IconDown class="icon-down" />
              </a-button>
              <template #content>
                <a-doption
                  v-for="p in projectList"
                  :key="p.id"
                  :value="p.id"
                  :active="p.id === selectedProjectId"
                >
                  {{ p.name }}
                </a-doption>
              </template>
            </a-dropdown>
          </li>
          <li>
            <a-dropdown @select="handleAuthoritySwitch">
              <a-button class="nav-btn" type="text">
                <template #icon>
                  <IconUserGroup />
                </template>
                {{ currentAuthorityName || t('common.selectRole') }}
                <IconDown class="icon-down" />
              </a-button>
              <template #content>
                <a-doption
                  v-for="r in authorityList"
                  :key="r.authorityId"
                  :value="r.authorityId"
                  :active="r.authorityId === selectedAuthorityId"
                >
                  {{ r.authorityName }}
                </a-doption>
              </template>
            </a-dropdown>
          </li>
          <li>
            <a-dropdown @select="handleLanguageSwitch">
              <a-button class="nav-btn" type="outline" :shape="'circle'">
                <template #icon>
                  <IconLanguage />
                </template>
              </a-button>
              <template #content>
                <a-doption value="zh">中文</a-doption>
                <a-doption value="en">English</a-doption>
              </template>
            </a-dropdown>
          </li>
          <li>
            <a-tooltip :content="theme === 'light' ? t('common.switchThemeDark') : t('common.switchThemeLight')">
              <a-button class="nav-btn" type="outline" :shape="'circle'" @click="toggleTheme">
                <template #icon>
                  <IconMoon v-if="theme === 'dark'" />
                  <IconSun v-else />
                </template>
              </a-button>
            </a-tooltip>
          </li>
          <li>
            <a-tooltip :content="t('common.notifications')">
              <div class="message-box-trigger">
                <a-badge :count="9" dot>
                  <a-button class="nav-btn" type="outline" :shape="'circle'">
                    <template #icon>
                      <IconNotification />
                    </template>
                  </a-button>
                </a-badge>
              </div>
            </a-tooltip>
          </li>
          <li>
            <a-dropdown @select="handleUserCommand">
              <a-avatar
                :size="32"
                :style="{ marginRight: '8px', cursor: 'pointer' }"
              >
                <img v-if="userInfo?.headerImg" :src="userInfo.headerImg" />
                <IconUser v-else />
              </a-avatar>
              <template #content>
                <a-doption value="person">
                  <template #icon><IconUser /></template>
                  {{ t('common.profile') }}
                </a-doption>
                <a-doption value="settings">
                  <template #icon><IconSettings /></template>
                  {{ t('common.userSettings') }}
                </a-doption>
                <a-doption value="logout">
                  <template #icon><IconPoweroff /></template>
                  {{ t('common.logout') }}
                </a-doption>
              </template>
            </a-dropdown>
          </li>
        </ul>
      </a-layout-header>
      <a-layout-content class="content">
        <router-view :user-info="userInfo" :embedded="true" />
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>

<script setup>
import { Message } from '@arco-design/web-vue'
import {
  IconHome,
  IconSettings,
  IconUser,
  IconPoweroff,
  IconDown,
  IconMoon,
  IconSun,
  IconSearch,
  IconLanguage,
  IconNotification,
  IconApps,
  IconUserGroup
} from '@arco-design/web-vue/es/icon'
import { computed, onMounted, ref, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { getUserInfo as getUserInfoApi, setUserAuthority as setUserAuthorityApi, clearAuth } from '../../services/appBridge'
import { getStoredTheme, setTheme, ThemeMode } from '../../utils/theme'
import logo from '../../assets/images/logo-universal.png'

const router = useRouter()
const route = useRoute()
const { t, locale } = useI18n()
const activeKey = ref('home')
const theme = ref(getStoredTheme())
const collapsed = ref(false)

const breadcrumbs = computed(() => {
  return route.matched.filter((item) => item.meta?.title || (item.name && item.name !== 'index'))
})

const toggleTheme = () => {
  theme.value = theme.value === ThemeMode.dark ? ThemeMode.light : ThemeMode.dark
  setTheme(theme.value)
}

const handleLanguageSwitch = (val) => {
  locale.value = val
  localStorage.setItem('locale', val)
}

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

const currentProjectName = computed(() => {
  if (!selectedProjectId.value) return ''
  const p = projectList.value.find(item => item.id === selectedProjectId.value)
  return p ? p.name : ''
})

const currentAuthorityName = computed(() => {
  if (!selectedAuthorityId.value) return ''
  const r = authorityList.value.find(item => item.authorityId === selectedAuthorityId.value)
  return r ? r.authorityName : ''
})

const loadUserInfo = async () => {
  try {
    userInfo.value = await getUserInfoApi()
    selectedProjectId.value = userInfo.value?.projectId
    selectedAuthorityId.value = userInfo.value?.authorityId
  } catch (e) {
    Message.error(e?.message || t('message.fetchUserInfoError'))
    await router.replace({ name: 'login' })
  }
}

const switchContext = async () => {
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
    Message.success(t('message.switchSuccess'))
  } catch (e) {
    Message.error(e?.message || t('message.switchError'))
    await loadUserInfo()
  } finally {
    switching.value = false
  }
}

const handleProjectSwitch = async (val) => {
  if (val === selectedProjectId.value) return
  selectedProjectId.value = val
  await switchContext()
}

const handleAuthoritySwitch = async (val) => {
  if (val === selectedAuthorityId.value) return
  selectedAuthorityId.value = val
  await switchContext()
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
      Message.success(t('message.logoutSuccess'))
      await router.replace({ name: 'login' })
    } catch (e) {
      Message.error(t('message.logoutError'))
    }
  } else if (val === 'settings') {
     await router.push({ name: 'settings' })
  }
}

onMounted(async () => {
  await loadUserInfo()
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
.sider {
  background: var(--color-bg-1);
  border-right: 1px solid var(--color-border);
}
.sider :deep(.arco-layout-sider-trigger) {
  border-right: 1px solid var(--color-border);
}
.logo {
  height: 32px;
  margin: 16px;
  background: var(--color-fill-2);
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--color-text-1);
  font-weight: bold;
  border-radius: 4px;
}
.logo-text {
  margin-left: 8px;
  color: var(--color-text-1);
  font-size: 16px;
  white-space: nowrap;
}
.logo-img {
  width: 32px;
  height: 32px;
}
.header {
  padding: 0 20px;
  background: var(--color-bg-1);
  border-bottom: 1px solid var(--color-border);
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 60px;
}
.right-side {
  display: flex;
  padding-left: 20px;
  list-style: none;
}
.right-side li {
  display: flex;
  align-items: center;
  padding: 0 10px;
}
.nav-btn {
  border-color: rgb(var(--gray-2));
  color: rgb(var(--gray-8));
  font-size: 16px;
}
.title {
  font-weight: 600;
}
.content {
  padding: 24px;
  height: calc(100vh - 64px);
  overflow: auto;
  background: var(--color-bg-2);
}
.icon-down {
  margin-left: 4px;
  font-size: 12px;
  color: var(--color-text-3);
}
</style>
