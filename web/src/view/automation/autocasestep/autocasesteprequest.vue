<template>
  <div class="autocase-step-request-container">
    <div class="left-panel">
      <div class="panel-header">
        <span class="title">接口列表</span>
      </div>
      <div class="panel-content">
        <div class="menu-container">
          <ApiMenu
            menutype="1"
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
            :show-header="false"
            :data="tableData"
            row-key="ID"
            :cell-style="{ paddingTop: '4px', paddingBottom: '4px' }"
            height="calc(100% - 50px)"
            stripe
          >
            <!-- 接口信息列（HTTP/Grpc 区分展示） -->
            <el-table-column min-width="300" align="center">
              <template #default="scope">
                <!-- HTTP 接口展示 -->
                <div
                  v-if="scope.row.request"
                  class="block"
                  :class="`block_${scope.row.request.method.toLowerCase()}`"
                >
          <span
            class="block-method block_method_color"
            :class="`block_method_${scope.row.request.method.toLowerCase()}`"
          >
            {{ scope.row.request.method }}
          </span>
                  <div class="block">
            <span
              class="block-method block_method_color block_method_options"
              v-if="scope.row.creator === 'yapi'"
              :title="'从YAPI导入的接口'"
            >
              YAPI
            </span>
                  </div>
                  <span class="block-method block_url">{{
                      scope.row.request.url
                    }}</span>
                  <span class="block-summary-description">{{
                      scope.row.name
                    }}</span>
                </div>

                <!-- Grpc 接口展示 -->
                <div v-if="scope.row.gRPC" class="block" :class="`block_put`">
          <span
            class="block-method block_method_color"
            :class="`block_method_put`"
          >
            {{ 'gRPC' }}
          </span>
                  <div class="block">
            <span
              class="block-method block_method_color block_method_options"
              v-if="scope.row.creator === 'yapi'"
              :title="'从YAPI导入的接口'"
            >
              YAPI
            </span>
                  </div>
                  <span class="block-method block_url">{{
                      scope.row.gRPC.url
                    }}</span>
                  <span class="block-summary-description">{{
                      scope.row.name
                    }}</span>
                </div>
              </template>
            </el-table-column>
          </el-table>
        </div>
      </div>
    </div>

    <div class="middle-actions">
      <div class="action-icon">
        <el-icon :size="20" color="#fff"><Right /></el-icon>
      </div>
      <div class="action-text">拖拽添加</div>
    </div>

    <div class="right-panel">
      <div class="panel-header">
        <span class="title">步骤「{{ stepName }}」接口</span>
        <span class="subtitle">（拖拽可排序）</span>
      </div>
      <div class="panel-content">
        <el-table
          v-if="rightDomKey"
          ref="rightTable"
          style="width: 100%"
          :show-header="false"
          :data="tableApiData"
          row-key="ID"
          :cell-style="{ paddingTop: '4px', paddingBottom: '4px' }"
          height="100%"
          stripe
        >
          <!-- 接口信息列（HTTP/Grpc 区分展示） -->
          <el-table-column min-width="300" align="center">
            <template #default="scope">
              <!-- HTTP 接口展示 -->
              <div
                v-if="scope.row.request"
                class="block"
                :class="`block_${scope.row.request.method.toLowerCase()}`"
              >
          <span
            class="block-method block_method_color"
            :class="`block_method_${scope.row.request.method.toLowerCase()}`"
          >
            {{ scope.row.request.method }}
          </span>
                <div class="block">
            <span
              class="block-method block_method_color block_method_options"
              v-if="scope.row.creator === 'yapi'"
              :title="'从YAPI导入的接口'"
            >
              YAPI
            </span>
                </div>
                <span class="block-method block_url">{{
                    scope.row.request.url
                  }}</span>
                <span class="block-summary-description">{{
                    scope.row.name
                  }}</span>
              </div>

              <!-- Grpc 接口展示 -->
              <div v-if="scope.row.gRPC" class="block" :class="`block_put`">
          <span
            class="block-method block_method_color"
            :class="`block_method_put`"
          >
            {{ 'gRPC' }}
          </span>
                <div class="block">
            <span
              class="block-method block_method_color block_method_options"
              v-if="scope.row.creator === 'yapi'"
              :title="'从YAPI导入的接口'"
            >
              YAPI
            </span>
                </div>
                <span class="block-method block_url">{{
                    scope.row.gRPC.url
                  }}</span>
                <span class="block-summary-description">{{
                    scope.row.name
                  }}</span>
              </div>

            </template>
          </el-table-column>
          <!-- 操作按钮列 -->
          <el-table-column label="操作" width="100" align="center">
            <template #default="scope">
              <el-button
                type="primary"
                link
                icon="edit"
                @click="updateAutoStepFunc(scope.row)"
              ></el-button>
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

    <el-dialog v-model="dialogFormVisible"
               :show-close="false"
               destroy-on-close
               :before-close="closeDialog"
               style="width:1400px;"
               :close-on-press-escape="false"
               :close-on-click-modal="false"
               top="0"
               :title="type === 'create' ? '新增接口' :type === 'update' ? '编辑接口' :'复制接口'"
    >
      <stepForm
        :menu="99999999"
        :formData="formData"
        :stepType="type"
        @close="closeDialog"
      />
    </el-dialog>
  </div>
</template>

<script setup>
  import ApiMenu from '@/components/platform/menu/index.vue'
  import { findAutoStep, getAutoStepList } from '@/api/automation/autostep.js'
  import {
    findAutoCaseStepApi,
    addAutoCaseStepApi,
    sortAutoCaseStepApi,
    deleteAutoCaseStep, deleteAutoCaseStepApi
  } from '@/api/automation/autocasestep.js'
  import { ref, onMounted, nextTick } from 'vue'
  import Sortable from 'sortablejs'
  import { watch } from 'vue'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import stepForm from '@/components/platform/step/index.vue'

  const tableData = ref([])
  const tableApiData = ref([])
  const multipleTable = ref(null)
  const rightTable = ref(null)
  let leftSortable = null
  let rightSortable = null
  const stepID = ref()
  const rightDomKey = ref(true)
  const stepName = ref('')
  const searchInfo = ref({})
  const props = defineProps({
    stepID: {
      type: [Number, String],
      required: true,
      default: 0
    },
    stepName: {
      type: String,
      required: true,
      default: ''
    }
  })

  const getTableData = async () => {
    const table = await getAutoStepList({ page: 1, pageSize: 9999, meun: Number(menuId.value), ...searchInfo.value })
    if (table.code === 0) {
      tableData.value = table.data.list
      // 重新初始化左侧拖动
      nextTick(() => {
        initLeftDrag()
      })
    }
  }
  // 生命周期
  onMounted(() => {
    console.log('============stepID11111', stepID.value)
    getTableApiData()
  })
  const getTableApiData = async () => {
    console.log('============stepID', stepID.value)
    const table = await findAutoCaseStepApi({ ID: stepID.value })
    if (table.code === 0) {
      const prevLen = tableApiData.value.length
      tableApiData.value = []
      tableApiData.value = table.data || []
      const newLen = tableApiData.value.length
      console.log('============table', table.data)
      // 当右侧列表从空变为有数据时，强制刷新一次表格
      if (prevLen === 0 && newLen > 0) {
        await forceRefreshRightDom()
      } else {
        // 否则只需要重新初始化右侧拖动排序
        nextTick(() => {
          initRightSort()
        })
      }
    }
  }
  watch(
    () => props.stepID,
    (newVal) => {
      console.log('stepID 变化:', newVal)
      stepID.value = newVal
      getTableApiData()
    },
    { immediate: true }
  )
  watch(
    () => props.stepName,
    (newVal) => {
      stepName.value = newVal
    },
    { immediate: true }
  )
  const menuId = ref(null)
  const handleMenuClick = (id) => {
    menuId.value = id
    getTableData()
  }

  const addAutoCaseStepApiFunc = (row, targetIndex) => {
    // 深拷贝原始数据避免修改源数据
    const newRow = JSON.parse(JSON.stringify(row))
    addAutoCaseStepApi({ id: stepID.value, api_id: newRow.ID }).then(res => {
      if (res.code === 0) {
        // 更新新增行的ID为接口返回的ID
        newRow.ID = res.data.id

        // 插入到目标位置，若目标位置不存在则追加到末尾
        if (typeof targetIndex === 'number') {
          console.log('============targetIndex1', tableApiData.value)
          console.log('============targetIndex', targetIndex)
          console.log('============targetIndexnewRow', newRow)
          tableApiData.value.splice(targetIndex, 0, newRow)
        } else {
          tableApiData.value.push(newRow)
        }

        // 更新排序（会触发重新从后端加载数据）
        sortAutoCaseStepApiFunc()
      }
    })

  }
  // 新增强制刷新函数：通过切换v-if状态销毁/重建DOM
  const forceRefreshRightDom = async () => {
    rightDomKey.value = false // 销毁DOM
    await nextTick() // 等待Vue完成DOM卸载
    rightDomKey.value = true // 重建DOM
    await nextTick() // 等待DOM重新挂载
    initRightSort() // 重新初始化排序（必须，否则拖动失效）
  }

  /**
   * 排序自动化用例步骤API
   * 根据tableApiData的数据，遍历tableApiData，将每个元素中的id和排序序号组装成数据
   * 然后调用排序API
   */
  const sortAutoCaseStepApiFunc = () => {
    // 遍历tableApiData，组装排序数据
    const data = tableApiData.value.map((item, index) => ({
      id: item.ID, // 使用接口的ID
      sort: index + 1 // 排序序号从1开始
    }))
    let req = {
      data: data
    }
    sortAutoCaseStepApi(req).then(res => {
      if (res.code === 0) {
        // 成功后重新从后端获取最新数据
        getTableApiData()
      }
    })
  }

  // 初始化左侧拖动到右侧功能（仅作为拖拽源，不在这里更新数据）
  const initLeftDrag = () => {
    const leftTable = multipleTable.value?.$el?.querySelector('.el-table__body-wrapper tbody')
    if (!leftTable) return

    // 销毁旧的排序实例，避免重复绑定事件导致接口被添加多次
    if (leftSortable) {
      leftSortable.destroy()
    }

    leftSortable = Sortable.create(leftTable, {
      group: {
        name: 'shared',
        pull: 'clone', // 从左侧拖出的是克隆
        put: false // 左侧不接收任何拖入
      },
      sort: false,
      draggable: '.el-table__row'
      // 不在左侧的事件里更新右侧数据，统一在右侧 onAdd 中处理
    })
  }


  // 初始化右侧拖动排序功能
  const initRightSort = () => {
    console.log('初始化右侧排序，数据:', tableApiData.value)

    // 尝试获取tbody，最多重试3次（每次间隔100ms）
    const tryGetTableBody = (retry = 3) => {
      if (retry <= 0) {
        console.error('无法获取右侧表格DOM，初始化失败')
        return
      }
      const rightTableBody = rightTable.value?.$el?.querySelector('.el-table__body-wrapper tbody')
      if (rightTableBody) {
        console.log('成功获取右侧表格DOM，初始化排序')
        // 销毁旧的排序实例（避免重复初始化）
        if (rightSortable) {
          rightSortable.destroy()
        }
        // 创建新的排序实例
        rightSortable = Sortable.create(rightTableBody, {
          group: {
            name: 'shared',
            pull: true, // 右侧可以把自己的元素拖出来（用于右侧内部排序）
            put: true   // 右侧可以接收来自左侧的元素
          },
          sort: true,
          draggable: '.el-table__row',
          // 从左侧拖拽到右侧时，在这里统一处理新增逻辑
          onAdd: (evt) => {
            // 只处理跨列表拖拽
            if (evt.from === evt.to) return

            const draggedRow = tableData.value[evt.oldIndex]
            if (draggedRow) {
              const targetIndex = evt.newIndex
              addAutoCaseStepApiFunc(draggedRow, targetIndex)
            }

            // 删除 Sortable 自动插入的 DOM 行，交给 Vue 根据 tableApiData 渲染
            if (evt.item && evt.item.parentNode) {
              evt.item.parentNode.removeChild(evt.item)
            }
          },
          // 仅处理右侧内部排序
          onEnd: (evt) => {
            // 只在同一个列表内拖动时才更新顺序，避免跨列表时重复处理
            if (evt.from !== evt.to) return

            console.log('右侧排序结束，更新顺序')
            const newIndex = evt.newIndex
            const oldIndex = evt.oldIndex
            if (newIndex !== oldIndex) {
              const movedItem = tableApiData.value.splice(oldIndex, 1)[0]
              tableApiData.value.splice(newIndex, 0, movedItem)
              sortAutoCaseStepApiFunc()
            }
          }
        })
      } else {
        // 未获取到DOM，重试
        setTimeout(() => tryGetTableBody(retry - 1), 100)
      }
    }

    tryGetTableBody() // 启动重试逻辑
  }

  onMounted(() => {
    getTableApiData()
    nextTick(() => {
      initLeftDrag()
      initRightSort()
    })
  })

  const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(() => {
      deleteAutoCaseStepFunc(row)
    })
  }
  const formData = ref()
  const type = ref('')
  const dialogFormVisible = ref(false)
  const updateAutoStepFunc = async (row) => {
    const res = await findAutoStep({ ID: row.ID })
    if (res.code === 0) {
      formData.value = res.data
      type.value = 'update'
      dialogFormVisible.value = true
    }
  }
  const closeDialog = (close) => {
    getTableApiData()
    dialogFormVisible.value = false
  }

  const deleteAutoCaseStepFunc = async (row) => {
    const res = await deleteAutoCaseStepApi({ ID: row.ID })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      if (tableData.value.length === 1 && page.value > 1) {
        page.value--
      }
      getTableApiData()
    }
  }
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
</script>

<style scoped>
  /* 保持你原有的样式不变 */
  .gva-btn-list {
    margin-bottom: 10px;
    display: flex;
    align-items: center;
    flex-wrap: wrap;
  }

  .block_post {
    border: 1px solid #49cc90;
    background-color: rgba(73, 204, 144, 0.1);
  }

  .block_method_post {
    background-color: #49cc90;
  }

  .block_put {
    border: 1px solid #fca130;
    background-color: rgba(252, 161, 48, 0.1);
  }

  .block_method_put {
    background-color: #fca130;
  }

  .block_get {
    border: 1px solid #61affe;
    background-color: rgba(97, 175, 254, 0.1);
  }

  .block_method_get {
    background-color: #61affe;
  }

  .block_delete {
    border: 1px solid #f93e3e;
    background-color: rgba(249, 62, 62, 0.1);
  }

  .block_method_delete {
    background-color: #f93e3e;
  }

  .block_patch {
    border: 1px solid #50e3c2;
    background-color: rgba(80, 227, 194, 0.1);
  }

  .block_method_patch {
    background-color: #50e3c2;
  }

  .block_head {
    border: 1px solid #e6a23c;
    background-color: rgba(230, 162, 60, 0.1);
  }

  .block_method_head {
    background-color: #e6a23c;
  }

  .block_options {
    border: 1px solid #409eff;
    background-color: rgba(64, 158, 255, 0.1);
  }

  .block_method_options {
    background-color: #409eff;
  }

  .block {
    position: relative;
    border-radius: 4px;
    height: 48px;
    overflow: hidden;
    padding: 5px;
    display: flex;
    align-items: center;
  }

  .block_url {
    word-break: normal;
    width: auto;
    display: block;
    white-space: pre-wrap;
    word-wrap: break-word;
    overflow: hidden;
    -webkit-box-flex: 1;
    -ms-flex: 1;
    flex: 1;
    font-family: Open Sans, sans-serif;
    color: var(--el-text-color-regular);
  }

  .block_method_color {
    cursor: pointer;
    color: #fff;
  }

  .block-method {
    font-size: 14px;
    font-weight: 600;
    min-width: 50px;
    padding: 0px 10px;
    text-align: center;
    border-radius: 5px;
    text-shadow: 0 1px 0 rgba(0, 0, 0, 0.1);
    font-family: Titillium Web, sans-serif;
    margin-right: 8px;
  }

  .block-summary-description {
    word-break: normal;
    width: auto;
    display: block;
    white-space: pre-wrap;
    word-wrap: break-word;
    overflow: hidden;
    -webkit-box-flex: 1;
    -ms-flex: 1;
    flex: 1;
    font-family: Open Sans, sans-serif;
    color: var(--el-text-color-regular);
  }

  .gva-pagination {
    margin-top: 16px;
    text-align: right;
  }

  .table-button {
    margin-right: 8px;
  }

  /* New Layout Styles */
  .autocase-step-request-container {
    display: flex;
    width: 100%;
    height: 750px;
    background-color: var(--el-bg-color-page);
    padding: 16px;
    box-sizing: border-box;
    gap: 16px;
  }

  .left-panel, .right-panel {
    display: flex;
    flex-direction: column;
    background-color: var(--el-bg-color);
    border-radius: 8px;
    box-shadow: 0 1px 2px -2px rgba(0, 0, 0, 0.16), 0 3px 6px 0 rgba(0, 0, 0, 0.12), 0 5px 12px 4px rgba(0, 0, 0, 0.09);
    overflow: hidden;
    transition: all 0.3s ease;
  }

  .left-panel:hover, .right-panel:hover {
    box-shadow: 0 3px 6px -4px rgba(0, 0, 0, 0.12), 0 6px 16px 0 rgba(0, 0, 0, 0.08), 0 9px 28px 8px rgba(0, 0, 0, 0.05);
  }

  .left-panel {
    flex: 1;
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

  /* Custom Scrollbar */
  .menu-container::-webkit-scrollbar {
    width: 6px;
    height: 6px;
  }

  .menu-container::-webkit-scrollbar-thumb {
    background: var(--el-border-color);
    border-radius: 3px;
  }

  .menu-container::-webkit-scrollbar-track {
    background: transparent;
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
    letter-spacing: 4px;
    font-weight: 500;
  }
</style>
