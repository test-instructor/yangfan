<template>
  <div class="page">
    <RouteInitFailed v-if="hasError" />
    <a-spin v-else :tip="t('common.loading')" />
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'
import { uiNodeRouteInitError } from '../../router/uiNodeState'
import RouteInitFailed from '../error/RouteInitFailed.vue'

const { t } = useI18n()
const hasError = computed(() => Boolean(uiNodeRouteInitError.value))
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

