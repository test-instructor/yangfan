<template>
  <div class="autocase-step-container">
    <div class="left-panel">
      <el-tabs v-model="activeTab" type="border-card" class="step-tabs" @tab-change="handleTabChange">
        <el-tab-pane :label="typeStepLabel" name="typeSteps" />
        <el-tab-pane label="接口列表" name="interfaces" />
      </el-tabs>
      <div class="panel-content">
        <div class="menu-container">
          <ApiMenu
            :key="activeTab"
            :menutype="currentMenuType"
            @getTreeID="handleMenuClick"
            detail=true
          />
        </div>
        <div class="table-container">
          <div class="search-box">
            <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline"
                     @keyup.enter="onSubmit">
              <el-form-item prop="name" style="margin-bottom: 0; margin-right: 10px;">
                <el-input v-model="searchInfo.name" placeholder="请输入接口名称" clearable prefix-icon="Search" />
              </el-form-item>
              <el-form-item style="margin-bottom: 0;">
                <el-button type="primary" icon="search" @click="onSubmit" circle></el-button>
                <el-button icon="refresh" @click="onReset" circle></el-button>
              </el-form-item>
            </el-form>
          </div>
          <el-table
            ref="multipleTable"
            style="width: 100%"
            :show-header="true"
            :data="tableData"
            row-key="ID"
            :cell-style="{ paddingTop: '4px', paddingBottom: '4px' }"
            height="calc(100% - 50px)"
            stripe
          >
            <el-table-column sortable align="left" label="步骤名称" prop="name" show-overflow-tooltip />
            <el-table-column align="left" label="运行环境" prop="envName" width="100" show-overflow-tooltip />
          </el-table>
        </div>
      </div>
    </div>

    <div class="middle-actions">
      <div class="action-icon">
        <el-icon :size="20" color="#fff">
          <Right />
        </el-icon>
      </div>
      <div class="action-text">拖拽添加</div>
    </div>

    <div class="right-panel">
      <div class="panel-header">
        <span class="title">用例「{{ caseName }}」步骤</span>
        <span class="subtitle">（拖拽可排序）</span>
      </div>
      <div class="panel-content">
        <el-table
          v-if="rightDomKey"
          ref="rightTable"
          style="width: 100%"
          :show-header="true"
          :data="tableApiData"
          row-key="ID"
          :cell-style="{ paddingTop: '4px', paddingBottom: '4px' }"
          height="100%"
          stripe
        >
          <el-table-column type="index" label="序号" width="60" align="center" />
          <el-table-column sortable align="left" label="步骤名称" prop="name" show-overflow-tooltip />
          <el-table-column align="left" label="运行环境" prop="envName" width="100" show-overflow-tooltip />
          <el-table-column align="center" label="独立运行配置" width="100">
            <template #default="scope">
              <el-switch
                v-model="scope.row.isConfig"
                inline-prompt
                active-text="是"
                inactive-text="否"
                active-color="#13ce66"
                inactive-color="#ff4949"
                @change="handleConfigChange(scope.row, 'isConfig')"
              />
            </template>
          </el-table-column>
          <el-table-column align="center" label="加载StepConfig" width="120">
            <template #default="scope">
              <el-switch
                v-model="scope.row.isStepConfig"
                inline-prompt
                active-text="是"
                inactive-text="否"
                active-color="#13ce66"
                inactive-color="#ff4949"
                @change="handleConfigChange(scope.row, 'isStepConfig')"
              />
            </template>
          </el-table-column>
          <el-table-column label="操作" width="80" align="center">
            <template #default="scope">
              <el-button
                type="danger"
                link
                icon="delete"
                @click="deleteRow(scope.row)"
              ></el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </div>

    <el-dialog v-model="dialogFormVisible" :show-close="false" destroy-on-close
               :before-close="closeDialog"
               width="1200px"
               :close-on-press-escape="false"
               :close-on-click-modal="false"
               top="5vh"
               :title="type === 'create' ? '新增接口' :type === 'update' ? '编辑接口' :'复制接口'"
    >
      <stepForm
        menu=99999999
        :formData="formData"
        type="update"
        @close="closeDialog"
      />
    </el-dialog>
  </div>
</template>

<script setup>

  import ApiMenu from '@/components/platform/menu/index.vue'
  import { ref, nextTick, onMounted, watch, computed } from 'vue'
  import { getAutoCaseStepList } from '@/api/automation/autocasestep.js'
  import {
    getAutoCaseSteps,
    addAutoCaseStep,
    sortAutoCaseStep,
    delAutoCaseStep,
    setStepConfig
  } from '@/api/automation/autocase.js'
  import Sortable from 'sortablejs'
  import { ElMessage, ElMessageBox } from 'element-plus'

  const tableData = ref([])
  const tableApiData = ref([])
  const searchInfo = ref({})
  const multipleTable = ref(null)
  const rightTable = ref(null)
  let leftSortable = null
  let rightSortable = null
  const rightDomKey = ref(true)

  const getTableData = async () => {
    const table = await getAutoCaseStepList({ page: 1, pageSize: 99999, menu: menuId.value, ...searchInfo.value })
    if (table.code === 0) {
      tableData.value = table.data.list
      nextTick(() => {
        initLeftDrag()
      })
    }
  }
  const menuId = ref(null)
  const handleMenuClick = (id) => {
    console.log('menuId', id)
    menuId.value = id
    getTableData()
  }
  const caseName = ref('')
  const onReset = () => {
    searchInfo.value = {}
    getTableData()
  }

  // 搜索
  const elSearchFormRef = ref()
  const onSubmit = () => {
    elSearchFormRef.value?.validate(async (valid) => {
      if (!valid) return
      if (searchInfo.value.ignore_popup === '') {
        searchInfo.value.ignore_popup = null
      }
      getTableData()
    })
  }

  const props = defineProps({
    caseID: {
      type: [Number, String],
      required: true,
      default: 1
    },
    caseName: {
      type: String,
      required: true,
      default: ''
    },
    caseType: {
      type: String,
      default: 'api'
    }
  })

  const activeTab = ref('typeSteps')
  const typeStepLabel = computed(() => {
    const map = {
      'api': '接口',
      'android': '安卓',
      'ios': 'iOS',
      'harmony': '鸿蒙',
      'browser': '浏览器'
    }
    return (map[props.caseType] || props.caseType) + '步骤'
  })

  const currentMenuType = computed(() => {
    if (activeTab.value === 'interfaces') {
      return '11'
    }
    return `casestep_${props.caseType}`
  })

  const handleTabChange = () => {
    menuId.value = null
    searchInfo.value = {}
    tableData.value = []
  }

  const getTableApiData = async () => {
    if (!props.caseID) return
    const table = await getAutoCaseSteps({ ID: props.caseID })
    if (table.code === 0) {
      tableApiData.value = table.data
      nextTick(() => {
        initRightSort()
      })
    }
  }

  const addAutoCaseStepFunc = (row, targetIndex) => {
    const newRow = JSON.parse(JSON.stringify(row))
    addAutoCaseStep({ caseId: Number(props.caseID), stepId: newRow.ID }).then(async res => {
      if (res.code === 0) {
        // 重新获取数据以确保ID正确
        await getTableApiData()
        // 移动数据
        const item = tableApiData.value.pop()
        tableApiData.value.splice(targetIndex, 0, item)
        // 排序
        sortAutoCaseStepFunc()
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

  const sortAutoCaseStepFunc = () => {
    const data = tableApiData.value.map((item, index) => ({
      id: item.parentId,
      sort: index + 1
    }))
    let req = {
      data: data
    }
    sortAutoCaseStep(req).then(res => {
      if (res.code === 0) {
        forceRefreshRightDom()
        getTableApiData()
      }
    })
  }

  const initLeftDrag = () => {
    const leftTable = multipleTable.value.$el.querySelector('.el-table__body-wrapper tbody')
    if (!leftTable) return

    if (leftSortable) leftSortable.destroy()

    leftSortable = Sortable.create(leftTable, {
      group: {
        name: 'shared',
        pull: 'clone',
        put: false
      },
      sort: false,
      draggable: '.el-table__row',
      onAdd: (evt) => {
        evt.preventDefault()
      },
      onEnd: (evt) => {
        if (evt.to !== evt.from) {
          const draggedRow = tableData.value[evt.oldIndex]
          if (draggedRow) {
            const targetIndex = evt.newIndex
            addAutoCaseStepFunc(draggedRow, targetIndex)
          }
        }
      }
    })
  }

  const initRightSort = () => {
    const tryGetTableBody = (retry = 3) => {
      if (retry <= 0) return
      const rightTableBody = rightTable.value?.$el?.querySelector('.el-table__body-wrapper tbody')
      if (rightTableBody) {
        if (rightSortable) {
          rightSortable.destroy()
        }
        rightSortable = Sortable.create(rightTableBody, {
          group: 'shared',
          sort: true,
          draggable: '.el-table__row',
          onEnd: (evt) => {
            const newIndex = evt.newIndex
            const oldIndex = evt.oldIndex
            if (newIndex !== oldIndex) {
              const movedItem = tableApiData.value.splice(oldIndex, 1)[0]
              tableApiData.value.splice(newIndex, 0, movedItem)
              sortAutoCaseStepFunc()
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
      console.log("=========111",row)
      const res = await delAutoCaseStep({ ID: row.parentId })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        getTableApiData()
      }
    })
  }

  const handleConfigChange = (row, field) => {
    setStepConfig({
      ID: row.parentId,
      isConfig: row.isConfig,
      isStepConfig: row.isStepConfig
    }).then(res => {
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '设置成功'
        })
      } else {
        // Revert change if failed
        if (field) {
            row[field] = !row[field]
        }
      }
    })
  }

  watch(
    () => props.caseID,
    (newVal) => {
      if (newVal) {
        getTableApiData()
      }
    },
    { immediate: true }
  )

  watch(
    () => props.caseName,
    (newVal) => {
      caseName.value = newVal
    },
    { immediate: true }
  )

  onMounted(() => {
    nextTick(() => {
      initLeftDrag()
      initRightSort()
    })
  })
</script>

<style scoped>
  .autocase-step-container {
    display: flex;
    width: 100%;
    height: 750px; /* Fixed height or calc(100vh - ...) depending on need */
    background-color: var(--el-bg-color-page);
    padding: 10px;
    box-sizing: border-box;
    gap: 10px;
  }

  .left-panel, .right-panel {
    flex: 1;
    display: flex;
    flex-direction: column;
    background-color: var(--el-bg-color);
    border-radius: 4px;
    box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
    overflow: hidden;
  }

  .panel-header {
    padding: 15px 20px;
    border-bottom: 1px solid var(--el-border-color-light);
    background-color: var(--el-fill-color-light);
    display: flex;
    align-items: center;
  }

  .title {
    font-size: 16px;
    font-weight: bold;
    color: var(--el-text-color-primary);
  }

  .subtitle {
    font-size: 12px;
    color: var(--el-text-color-secondary);
    margin-left: 8px;
  }

  .panel-content {
    flex: 1;
    display: flex;
    overflow: hidden;
    padding: 10px;
  }

  .menu-container {
    width: 220px;
    border-right: 1px solid var(--el-border-color-light);
    margin-right: 10px;
    overflow-y: auto;
  }

  .table-container {
    flex: 1;
    display: flex;
    flex-direction: column;
    overflow: hidden;
  }

  .search-box {
    margin-bottom: 10px;
    display: flex;
    justify-content: flex-end;
  }

  .middle-actions {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    width: 48px;
    gap: 8px;
  }

  .action-icon {
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

  .action-icon:hover {
    transform: scale(1.1);
  }

  .action-text {
    font-size: 12px;
    color: var(--el-text-color-secondary);
    writing-mode: vertical-rl;
    letter-spacing: 2px;
  }

  .step-tabs :deep(.el-tabs__content) {
    display: none;
  }
</style>
