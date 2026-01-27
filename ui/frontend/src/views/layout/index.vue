<template>
  <a-layout class="layout">
    <a-layout-sider collapsible breakpoint="xl" class="sider">
      <div class="logo">
        <div class="logo-text">扬帆测试平台</div>
      </div>
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


        </a-space>
        <ul class="right-side">
          <li>
            <a-tooltip content="语言">
              <a-button class="nav-btn" type="outline" :shape="'circle'">
                <template #icon>
                  <IconLanguage />
                </template>
              </a-button>
            </a-tooltip>
          </li>
          <li>
            <a-tooltip :content="theme === 'light' ? '点击切换为暗黑模式' : '点击切换为亮色模式'">
              <a-button class="nav-btn" type="outline" :shape="'circle'" @click="toggleTheme">
                <template #icon>
                  <IconMoon v-if="theme === 'dark'" />
                  <IconSun v-else />
                </template>
              </a-button>
            </a-tooltip>
          </li>
          <li>
            <a-tooltip content="消息通知">
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
                  个人信息
                </a-doption>
                <a-doption value="settings">
                  <template #icon><IconSettings /></template>
                  用户设置
                </a-doption>
                <a-doption value="logout">
                  <template #icon><IconPoweroff /></template>
                  退出登录
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
} from '@arco-design/web-vue/es/icon'
import { computed, onMounted, ref, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { getUserInfo as getUserInfoApi, setUserAuthority as setUserAuthorityApi, clearAuth } from '../../services/appBridge'
import { getStoredTheme, setTheme, ThemeMode } from '../../utils/theme'

const router = useRouter()
const route = useRoute()
const activeKey = ref('home')
const theme = ref(getStoredTheme())

const toggleTheme = () => {
  theme.value = theme.value === ThemeMode.dark ? ThemeMode.light : ThemeMode.dark
  setTheme(theme.value)
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
</style>
