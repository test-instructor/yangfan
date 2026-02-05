<template>
  <template v-for="m in visibleMenus" :key="m.name">
    <a-sub-menu v-if="hasChildren(m)" :key="String(m.name)">
      <template #title>
        <span class="menu-title">
          <component v-if="resolveIcon(m)" :is="resolveIcon(m)" />
          <span class="menu-text">{{ titleText(m) }}</span>
        </span>
      </template>
      <SideMenu :menus="m.children" />
    </a-sub-menu>
    <a-menu-item v-else :key="String(m.name)">
      <template #icon>
        <component v-if="resolveIcon(m)" :is="resolveIcon(m)" />
      </template>
      {{ titleText(m) }}
    </a-menu-item>
  </template>
</template>

<script setup>
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'
import {
  IconHome,
  IconSettings,
  IconUser,
  IconApps,
  IconFile,
  IconClockCircle
} from '@arco-design/web-vue/es/icon'

defineOptions({ name: 'SideMenu' })

const props = defineProps({
  menus: {
    type: Array,
    default: () => []
  }
})

const { t, te } = useI18n()

const visibleMenus = computed(() => {
  const list = Array.isArray(props.menus) ? props.menus : []
  return list.filter((m) => m && !m.hidden).slice().sort((a, b) => (a.sort || 0) - (b.sort || 0))
})

const hasChildren = (m) => {
  const children = Array.isArray(m?.children) ? m.children : []
  return children.some((c) => c && !c.hidden)
}

const titleText = (m) => {
  const key = m?.meta?.title || ''
  if (key && te(key)) return t(key)
  return key || String(m?.name || '')
}

const iconMap = {
  home: IconHome,
  settings: IconSettings,
  user: IconUser,
  android: IconApps,
  action: IconFile,
  'case-step': IconFile,
  testcase: IconFile,
  'time-task': IconClockCircle,
  'bxs-report': IconFile
}

const resolveIcon = (m) => {
  const icon = m?.meta?.icon || ''
  return iconMap[icon] || null
}
</script>

<style scoped>
.menu-title {
  display: inline-flex;
  align-items: center;
  gap: 8px;
}
.menu-text {
  display: inline-block;
}
</style>

