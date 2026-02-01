<template>
  <div class="task-case-container">
    <div class="left-panel card-panel">
      <div class="panel-header">
        <div class="header-title-group">
          <span class="title">用例菜单</span>
          <span class="subtitle">选择用例添加到任务</span>
        </div>
      </div>
      <div class="panel-content">
        <div class="left-content-wrapper" style="display: flex; flex-direction: column; width: 100%; height: 100%;">
          <el-tabs v-model="activeTab" type="border-card" class="left-tabs" style="height: 100%; display: flex; flex-direction: column; border: none; box-shadow: none;">
            <el-tab-pane label="类型步骤" name="type_step" class="custom-tab-pane">
              <div class="type-selector-bar" style="padding: 8px 12px; border-bottom: 1px solid var(--el-border-color-light);">
                <el-radio-group v-model="stepType" size="small">
                  <el-radio-button label="casestep_android">安卓</el-radio-button>
                  <el-radio-button label="casestep_ios">iOS</el-radio-button>
                  <el-radio-button label="casestep_harmony">鸿蒙</el-radio-button>
                  <el-radio-button label="casestep_browser">浏览器</el-radio-button>
                </el-radio-group>
              </div>
              <div class="pane-content-inner" style="display: flex; flex: 1; height: calc(100% - 40px);">
                <div class="menu-container custom-scrollbar">
                  <ApiMenu
                    :key="currentMenuType"
                    :menutype="currentMenuType"
                    eventType="1"
                    @getTreeID="handleMenuClick"
                    detail=true
                  />
                </div>
                <div class="table-container">
                  <div class="search-box">
                    <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="search-form" @keyup.enter="onSubmit">
                      <el-form-item prop="caseName" class="search-item">
                        <el-input 
                          v-model="searchInfo.caseName" 
                          placeholder="搜索用例名称" 
                          clearable 
                          prefix-icon="Search" 
                          class="search-input"
                        />
                      </el-form-item>
                      <el-form-item class="action-item">
                        <el-button type="primary" icon="search" @click="onSubmit" circle plain></el-button>
                        <el-button icon="refresh" @click="onReset" circle plain></el-button>
                      </el-form-item>
                    </el-form>
                  </div>
                  <div class="table-wrapper">
                    <el-table
                      ref="leftTableRef"
                      style="width: 100%"
                      :show-header="true"
                      :data="leftTableData"
                      row-key="ID"
                      :cell-style="{ paddingTop: '8px', paddingBottom: '8px' }"
                      :header-cell-style="{ background: 'var(--el-fill-color-light)', color: 'var(--el-text-color-secondary)', fontWeight: '600' }"
                      height="100%"
                      stripe
                      highlight-current-row
                      class="custom-table"
                    >
                      <el-table-column align="left" label="用例名称" prop="caseName" show-overflow-tooltip min-width="150" />
                      <el-table-column align="left" label="运行配置" prop="configName" width="120" show-overflow-tooltip />
                    </el-table>
                  </div>
                </div>
              </div>
            </el-tab-pane>
            <el-tab-pane label="接口列表" name="interface_list" class="custom-tab-pane">
              <div class="pane-content-inner" style="display: flex; flex: 1; height: 100%;">
                <div class="menu-container custom-scrollbar">
                  <ApiMenu
                    :key="'21'"
                    menutype="21"
                    eventType="1"
                    @getTreeID="handleMenuClick"
                    detail=true
                  />
                </div>
                <div class="table-container">
                  <div class="search-box">
                    <el-form ref="elSearchFormRef2" :inline="true" :model="searchInfo" class="search-form" @keyup.enter="onSubmit">
                      <el-form-item prop="caseName" class="search-item">
                        <el-input 
                          v-model="searchInfo.caseName" 
                          placeholder="搜索用例名称" 
                          clearable 
                          prefix-icon="Search" 
                          class="search-input"
                        />
                      </el-form-item>
                      <el-form-item class="action-item">
                        <el-button type="primary" icon="search" @click="onSubmit" circle plain></el-button>
                        <el-button icon="refresh" @click="onReset" circle plain></el-button>
                      </el-form-item>
                    </el-form>
                  </div>
                  <div class="table-wrapper">
                    <el-table
                      ref="leftTableRef2"
                      style="width: 100%"
                      :show-header="true"
                      :data="leftTableData"
                      row-key="ID"
                      :cell-style="{ paddingTop: '8px', paddingBottom: '8px' }"
                      :header-cell-style="{ background: 'var(--el-fill-color-light)', color: 'var(--el-text-color-secondary)', fontWeight: '600' }"
                      height="100%"
                      stripe
                      highlight-current-row
                      class="custom-table"
                    >
                      <el-table-column align="left" label="用例名称" prop="caseName" show-overflow-tooltip min-width="150" />
                      <el-table-column align="left" label="运行配置" prop="configName" width="120" show-overflow-tooltip />
                    </el-table>
                  </div>
                </div>
              </div>
            </el-tab-pane>
          </el-tabs>
        </div>
      </div>
    </div>

    <div class="middle-actions">
      <div class="transfer-arrow">
        <el-icon :size="20" color="#fff"><Right /></el-icon>
      </div>
      <div class="action-text">拖拽添加</div>
    </div>

    <div class="right-panel card-panel">
      <div class="panel-header">
        <div class="header-title-group">
          <span class="title">任务「{{ taskName }}」用例</span>
          <el-tag size="small" type="info" effect="plain" round class="subtitle-tag">拖拽可排序</el-tag>
        </div>
      </div>
      <div class="panel-content">
        <el-table
          v-if="rightDomKey"
          ref="rightTableRef"
          style="width: 100%"
          :show-header="true"
          :data="rightTableData"
          row-key="ID"
          :cell-style="{ paddingTop: '8px', paddingBottom: '8px' }"
          :header-cell-style="{ background: 'var(--el-fill-color-light)', color: 'var(--el-text-color-secondary)', fontWeight: '600' }"
          height="100%"
          stripe
          highlight-current-row
          class="custom-table"
        >
          <el-table-column type="index" label="序号" width="60" align="center" />
          <el-table-column align="left" label="用例名称" prop="caseName" show-overflow-tooltip min-width="150" />
          <el-table-column align="left" label="运行配置" prop="configName" width="120" show-overflow-tooltip />
          <el-table-column label="操作" width="80" align="center" fixed="right">
            <template #default="scope">
              <el-button type="danger" link icon="delete" @click="deleteRow(scope.row)" class="delete-btn"></el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </div>
  </div>
</template>

<script setup>
import ApiMenu from '@/components/platform/menu/index.vue'
import { ref, nextTick, onMounted, watch, computed } from 'vue'
import Sortable from 'sortablejs'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getAutoCaseList } from '@/api/automation/autocase.js'
import { getTimerTaskCases, addTimerTaskCase, sortTimerTaskCase, delTimerTaskCase } from '@/api/automation/timertask.js'

const props = defineProps({
  taskID: { type: [Number, String], required: true },
  taskName: { type: String, default: '' },
  taskType: { type: String, default: 'api' }
})

const activeTab = ref('type_step')
const stepType = ref('casestep_android')
const normalizePlatform = (raw) => String(raw ?? '').trim().toLowerCase()
const normalizedTaskType = computed(() => normalizePlatform(props.taskType) || 'api')

const stepTypeFromPlatform = (t) => {
  if (t === 'android') return 'casestep_android'
  if (t === 'ios') return 'casestep_ios'
  if (t === 'harmony') return 'casestep_harmony'
  if (t === 'browser') return 'casestep_browser'
  return ''
}

const platformFromStepType = (t) => {
  if (t === 'casestep_android') return 'android'
  if (t === 'casestep_ios') return 'ios'
  if (t === 'casestep_harmony') return 'harmony'
  if (t === 'casestep_browser') return 'browser'
  return ''
}

const effectiveCaseType = computed(() => {
  if (activeTab.value === 'interface_list') return 'api'
  return platformFromStepType(stepType.value) || normalizedTaskType.value
})

const initTab = () => {
  if (normalizedTaskType.value === 'api') {
    activeTab.value = 'interface_list'
    return
  }
  activeTab.value = 'type_step'
  const st = stepTypeFromPlatform(normalizedTaskType.value)
  if (st) stepType.value = st
}
const currentMenuType = computed(() => {
  return stepType.value
})

const leftTableData = ref([])
const rightTableData = ref([])
const searchInfo = ref({})
const leftTableRef = ref(null)
const leftTableRef2 = ref(null)
const rightTableRef = ref(null)
const elSearchFormRef = ref()
const elSearchFormRef2 = ref()
let leftSortable = null
let rightSortable = null
const rightDomKey = ref(true)

const menuId = ref(null)
const handleMenuClick = (id) => {
  menuId.value = id
  getLeftTableData()
}

watch(activeTab, () => {
  menuId.value = null
  leftTableData.value = []
})

const getLeftTableData = async () => {
  const table = await getAutoCaseList({ page: 1, pageSize: 99999, menu: menuId.value, type: effectiveCaseType.value, ...searchInfo.value })
  if (table.code === 0) {
    leftTableData.value = table.data.list
    nextTick(() => { initLeftDrag() })
  }
}

const onReset = () => {
  searchInfo.value = {}
  getLeftTableData()
}

const onSubmit = () => {
  const ref = activeTab.value === 'type_step' ? elSearchFormRef : elSearchFormRef2
  ref.value?.validate(async (valid) => {
    if (!valid) return
    getLeftTableData()
  })
}

const getRightTableData = async () => {
  if (!props.taskID) return
  const res = await getTimerTaskCases({ ID: props.taskID })
  if (res.code === 0) {
    rightTableData.value = res.data
    nextTick(() => { initRightSort() })
  }
}

const addTaskCaseFunc = (row, targetIndex) => {
  const newRow = JSON.parse(JSON.stringify(row))
  addTimerTaskCase({ taskId: Number(props.taskID), caseId: newRow.ID }).then(async (res) => {
    if (res.code === 0) {
      await getRightTableData()
      const item = rightTableData.value.pop()
      rightTableData.value.splice(targetIndex, 0, item)
      sortTaskCaseFunc()
    }
  })
}

const forceRefreshRightDom = async () => {
  rightDomKey.value = false
  await nextTick()
  rightDomKey.value = true
  await nextTick()
  initRightSort()
}

const sortTaskCaseFunc = () => {
  const data = rightTableData.value.map((item, index) => ({ id: item.parentId, sort: index + 1 }))
  sortTimerTaskCase({ data }).then(res => {
    if (res.code === 0) {
      forceRefreshRightDom()
      getRightTableData()
    }
  })
}

const initLeftDrag = () => {
  const leftTable = leftTableRef.value.$el.querySelector('.el-table__body-wrapper tbody')
  if (!leftTable) return
  if (leftSortable) leftSortable.destroy()
  leftSortable = Sortable.create(leftTable, {
    group: { name: 'shared', pull: 'clone', put: false },
    sort: false,
    draggable: '.el-table__row',
    onAdd: (evt) => { evt.preventDefault() },
    onEnd: (evt) => {
      if (evt.to !== evt.from) {
        const draggedRow = leftTableData.value[evt.oldIndex]
        if (draggedRow) {
          const targetIndex = evt.newIndex
          addTaskCaseFunc(draggedRow, targetIndex)
        }
      }
    }
  })
}

const initRightSort = () => {
  const tryGetTableBody = (retry = 3) => {
    if (retry <= 0) return
    const rightTableBody = rightTableRef.value?.$el?.querySelector('.el-table__body-wrapper tbody')
    if (rightTableBody) {
      if (rightSortable) { rightSortable.destroy() }
      rightSortable = Sortable.create(rightTableBody, {
        group: 'shared',
        sort: true,
        draggable: '.el-table__row',
        onEnd: (evt) => {
          const newIndex = evt.newIndex
          const oldIndex = evt.oldIndex
          if (newIndex !== oldIndex) {
            const movedItem = rightTableData.value.splice(oldIndex, 1)[0]
            rightTableData.value.splice(newIndex, 0, movedItem)
            sortTaskCaseFunc()
          }
        }
      })
    } else {
      setTimeout(() => tryGetTableBody(retry - 1), 100)
    }
  }
  tryGetTableBody()
}

const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const res = await delTimerTaskCase({ ID: row.parentId })
    if (res.code === 0) {
      ElMessage({ type: 'success', message: '删除成功' })
      getRightTableData()
    }
  })
}

watch(() => props.taskID, (newVal) => { if (newVal) getRightTableData() }, { immediate: true })
watch(() => props.taskName, (newVal) => { /* no-op: title reactive */ }, { immediate: true })

onMounted(() => {
  initTab()
  nextTick(() => { initLeftDrag(); initRightSort() })
})
</script>

<style scoped>
.task-case-container {
  display: flex;
  width: 100%;
  height: 100%;
  min-height: 600px;
  background-color: var(--el-bg-color-page);
  padding: 16px;
  box-sizing: border-box;
  gap: 16px;
}

.card-panel {
  display: flex;
  flex-direction: column;
  background-color: var(--el-bg-color);
  border-radius: 8px;
  box-shadow: 0 1px 2px -2px rgba(0, 0, 0, 0.16), 0 3px 6px 0 rgba(0, 0, 0, 0.12), 0 5px 12px 4px rgba(0, 0, 0, 0.09);
  overflow: hidden;
  transition: all 0.3s ease;
}

.card-panel:hover {
  box-shadow: 0 3px 6px -4px rgba(0, 0, 0, 0.12), 0 6px 16px 0 rgba(0, 0, 0, 0.08), 0 9px 28px 8px rgba(0, 0, 0, 0.05);
}

.left-tabs :deep(.el-tabs__content) {
  flex: 1;
  display: flex;
  flex-direction: column;
  padding: 0 !important;
}

.left-tabs :deep(.el-tab-pane) {
  height: 100%;
  display: flex;
  flex-direction: column;
}

.custom-tab-pane {
  display: flex;
  flex-direction: column;
  height: 100%;
}

.left-panel {
  flex: 0 0 600px;
  max-width: 600px;
}

.right-panel {
  flex: 1;
}

.panel-header {
  padding: 16px 24px;
  border-bottom: 1px solid var(--el-border-color-light);
  background-color: var(--el-bg-color);
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.header-title-group {
  display: flex;
  align-items: center;
  gap: 8px;
}

.title {
  font-size: 16px;
  font-weight: 600;
  color: var(--el-text-color-primary);
  line-height: 1.5;
}

.subtitle {
  font-size: 12px;
  color: var(--el-text-color-secondary);
  margin-left: 4px;
}

.subtitle-tag {
  margin-left: 8px;
  font-weight: normal;
}

.panel-content {
  flex: 1;
  display: flex;
  overflow: hidden;
  padding: 0;
}

.menu-container {
  width: 240px;
  border-right: 1px solid var(--el-border-color-light);
  background-color: var(--el-fill-color-light);
  overflow-y: auto;
}

.table-container {
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  padding: 12px;
  background-color: var(--el-bg-color);
}

.search-box {
  margin-bottom: 12px;
  padding: 0 4px;
}

.search-form {
  display: flex;
  align-items: center;
  width: 100%;
}

.search-item {
  flex: 1;
  margin-bottom: 0 !important;
  margin-right: 12px !important;
}

.action-item {
  margin-bottom: 0 !important;
  margin-right: 0 !important;
}

.search-input :deep(.el-input__wrapper) {
  box-shadow: 0 0 0 1px var(--el-border-color) inset;
  border-radius: 4px;
}

.search-input :deep(.el-input__wrapper:hover) {
  box-shadow: 0 0 0 1px var(--el-border-color) inset;
}

.search-input :deep(.el-input__wrapper.is-focus) {
  box-shadow: 0 0 0 1px var(--el-color-primary) inset;
}

.table-wrapper {
  flex: 1;
  overflow: hidden;
  border-radius: 4px;
  border: 1px solid var(--el-border-color-light);
}

.middle-actions {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  width: 48px;
  gap: 8px;
}

.transfer-arrow {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  background: linear-gradient(135deg, #409eff, #3a8ee6);
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 2px 8px rgba(64, 158, 255, 0.3);
  transition: transform 0.3s ease;
}

.transfer-arrow:hover {
  transform: scale(1.1);
}

.action-text {
  font-size: 12px;
  color: var(--el-text-color-secondary);
  writing-mode: vertical-rl;
  letter-spacing: 4px;
  font-weight: 500;
}

/* Custom Scrollbar */
.custom-scrollbar::-webkit-scrollbar {
  width: 6px;
  height: 6px;
}

.custom-scrollbar::-webkit-scrollbar-thumb {
  background: var(--el-border-color);
  border-radius: 3px;
}

.custom-scrollbar::-webkit-scrollbar-track {
  background: transparent;
}

.delete-btn {
  transition: color 0.2s;
}

.delete-btn:hover {
  color: var(--el-color-danger);
  background-color: var(--el-color-danger-light-9);
  border-radius: 4px;
}

/* Table Tweaks */
:deep(.el-table__inner-wrapper::before) {
  display: none;
}
</style>
