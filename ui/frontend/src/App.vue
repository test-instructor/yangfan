<template>
  <a-layout class="layout-demo">
    <a-layout-sider collapsible breakpoint="xl">
      <div class="logo">Yangfan UI</div>
      <a-menu
        :default-selected-keys="['1']"
        :style="{ width: '100%' }"
      >
        <a-menu-item key="1">
          Dashboard
        </a-menu-item>
        <a-menu-item key="2">
          Test Execution
        </a-menu-item>
        <a-menu-item key="3">
          Settings
        </a-menu-item>
      </a-menu>
    </a-layout-sider>
    <a-layout>
      <a-layout-header style="padding-left: 20px; background: var(--color-bg-2); border-bottom: 1px solid var(--color-border);">
        <a-space>
          <h3>Yangfan Automation Client</h3>
        </a-space>
      </a-layout-header>
      <a-layout-content style="padding: 24px;">
        <HelloWorld />
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>

<script setup>
import HelloWorld from './components/HelloWorld.vue'
</script>

<style>
body {
  margin: 0;
}
.layout-demo {
  height: 100vh;
  width: 100vw;
}
.layout-demo .arco-layout-sider {
  background: var(--color-bg-2);
  border-right: 1px solid var(--color-border);
}
.layout-demo .arco-layout-content {
  background: var(--color-bg-1);
  display: flex;
  flex-direction: column;
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
</style>
