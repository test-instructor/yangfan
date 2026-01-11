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
        <div class="menu-container custom-scrollbar">
          <ApiMenu
            menutype="21"
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
              :header-cell-style="{ background: '#f5f7fa', color: '#606266', fontWeight: '600' }"
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
          :header-cell-style="{ background: '#f5f7fa', color: '#606266', fontWeight: '600' }"
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
import { ref, nextTick, onMounted, watch } from 'vue'
import Sortable from 'sortablejs'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getAutoCaseList } from '@/api/automation/autocase.js'
import { getTimerTaskCases, addTimerTaskCase, sortTimerTaskCase, delTimerTaskCase } from '@/api/automation/timertask.js'

const props = defineProps({
  taskID: { type: [Number, String], required: true },
  taskName: { type: String, default: '' }
})

const leftTableData = ref([])
const rightTableData = ref([])
const searchInfo = ref({})
const leftTableRef = ref(null)
const rightTableRef = ref(null)
let leftSortable = null
let rightSortable = null
const rightDomKey = ref(true)

const menuId = ref(null)
const handleMenuClick = (id) => {
  menuId.value = id
  getLeftTableData()
}

const getLeftTableData = async () => {
  const table = await getAutoCaseList({ page: 1, pageSize: 99999, menu: menuId.value, ...searchInfo.value })
  if (table.code === 0) {
    leftTableData.value = table.data.list
    nextTick(() => { initLeftDrag() })
  }
}

const onReset = () => {
  searchInfo.value = {}
  getLeftTableData()
}

const elSearchFormRef = ref()
const onSubmit = () => {
  elSearchFormRef.value?.validate(async (valid) => {
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
  nextTick(() => { initLeftDrag(); initRightSort() })
})
</script>

<style scoped>
.task-case-container {
  display: flex;
  width: 100%;
  height: 100%;
  min-height: 600px;
  background-color: #f0f2f5;
  padding: 16px;
  box-sizing: border-box;
  gap: 16px;
}

.card-panel {
  display: flex;
  flex-direction: column;
  background-color: #fff;
  border-radius: 8px;
  box-shadow: 0 1px 2px -2px rgba(0, 0, 0, 0.16), 0 3px 6px 0 rgba(0, 0, 0, 0.12), 0 5px 12px 4px rgba(0, 0, 0, 0.09);
  overflow: hidden;
  transition: all 0.3s ease;
}

.card-panel:hover {
  box-shadow: 0 3px 6px -4px rgba(0, 0, 0, 0.12), 0 6px 16px 0 rgba(0, 0, 0, 0.08), 0 9px 28px 8px rgba(0, 0, 0, 0.05);
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
  border-bottom: 1px solid #f0f0f0;
  background-color: #fff;
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
  color: #303133;
  line-height: 1.5;
}

.subtitle {
  font-size: 12px;
  color: #909399;
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
  border-right: 1px solid #f0f0f0;
  background-color: #fafafa;
  overflow-y: auto;
}

.table-container {
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  padding: 12px;
  background-color: #fff;
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
  box-shadow: 0 0 0 1px #dcdfe6 inset;
  border-radius: 4px;
}

.search-input :deep(.el-input__wrapper:hover) {
  box-shadow: 0 0 0 1px #c0c4cc inset;
}

.search-input :deep(.el-input__wrapper.is-focus) {
  box-shadow: 0 0 0 1px #409eff inset;
}

.table-wrapper {
  flex: 1;
  overflow: hidden;
  border-radius: 4px;
  border: 1px solid #ebeef5;
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
  color: #909399;
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
  background: #dcdfe6;
  border-radius: 3px;
}

.custom-scrollbar::-webkit-scrollbar-track {
  background: transparent;
}

.delete-btn {
  transition: color 0.2s;
}

.delete-btn:hover {
  color: #f56c6c;
  background-color: #fef0f0;
  border-radius: 4px;
}

/* Table Tweaks */
:deep(.el-table__inner-wrapper::before) {
  display: none; /* Remove bottom border */
}
</style>