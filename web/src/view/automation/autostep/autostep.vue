<template>
  <div class="autostep-container">
    <div class="main-panel card-panel">

      <div class="panel-content">
        <div class="menu-container custom-scrollbar">
          <ApiMenu
            :menutype="currentMenuType"
            @getTreeID="handleMenuClick"
          />
        </div>
        <div class="table-container">
          <div class="search-box">
            <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="search-form"
                     @keyup.enter="onSubmit">
              <el-form-item label="创建日期" prop="createdAtRange" class="search-item">
                <template #label>
                  <span>
                    创建日期
                    <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                      <el-icon><QuestionFilled /></el-icon>
                    </el-tooltip>
                  </span>
                </template>
                <el-date-picker
                  v-model="searchInfo.createdAtRange"
                  class="!w-380px"
                  type="datetimerange"
                  range-separator="至"
                  start-placeholder="开始时间"
                  end-placeholder="结束时间"
                />
              </el-form-item>

              <el-form-item label="接口名称" prop="name" class="search-item">
                <el-input v-model="searchInfo.name" placeholder="搜索条件" clearable prefix-icon="Search" />
              </el-form-item>

              <template v-if="showAllQuery">
                <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
              </template>

              <el-form-item class="action-item">
                <el-button type="primary" icon="search" @click="onSubmit" circle plain></el-button>
                <el-button icon="refresh" @click="onReset" circle plain></el-button>
              </el-form-item>
            </el-form>
          </div>
          <div class="action-buttons">
            <el-button type="primary" icon="plus" @click="openDialog('create')">新增</el-button>
            <el-button type="primary" icon="upload" @click="openCurlImport">导入CURL</el-button>
            <el-button icon="delete" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
          </div>
          <div class="table-wrapper">
            <el-table
              ref="multipleTable"
              style="width: 100%"
              :show-header="false"
              :data="tableData"
              row-key="ID"
              @selection-change="handleSelectionChange"
              :cell-style="{ paddingTop: '4px', paddingBottom: '4px' }"
              height="100%"
              stripe
            >
              <!-- 选择列 -->
              <el-table-column type="selection" />

              <!-- 接口信息列（HTTP/Grpc 区分展示） -->
              <el-table-column min-width="600" align="center">
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
                    <span class="block-method block_url">{{ scope.row.request.url }}</span>
                    <span class="block-summary-description">{{ scope.row.name }}</span>
                  </div>

                  <!-- UI 动作展示 -->
                  <div v-else-if="scope.row[currentPlatform]" class="block block_post">
                      <span class="block-method block_method_color block_method_post">
                          {{ currentPlatform.toUpperCase() }}
                      </span>
                      <span class="block-summary-description" style="margin-left: 10px;">{{ scope.row.name }}</span>
                      <span class="block-method block_url" style="margin-left: 10px; color: #999;">
                          {{ (scope.row[currentPlatform].actions || []).length }} Actions
                      </span>
                  </div>

                  <!-- Grpc 接口展示 -->
                  <div v-else-if="scope.row.gRPC" class="block" :class="`block_put`">
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
                    <span class="block-method block_url">{{ scope.row.gRPC.url }}</span>
                    <span class="block-summary-description">{{ scope.row.name }}</span>
                  </div>
                </template>
              </el-table-column>

              <!-- 操作按钮列 -->
              <el-table-column label="按钮组" min-width="240">
                <template #default="scope">
                  <div style="display: flex; align-items: center;">
                    <Runner case_type="api" :case_id="scope.row.ID" style="margin-right: 10px;"/>
                    <el-button
                      type="text"
                      icon="edit"
                      class="table-button"
                      @click="updateAutoStepFunc(scope.row)"
                    >变更</el-button>
                    <el-button
                      type="text"
                      icon="copy"
                      class="table-button"
                      @click="copyAutoStepFunc(scope.row)"
                    >复制</el-button>
                    <el-button
                      type="text"
                      icon="delete"
                      @click="deleteRow(scope.row)"
                    >删除</el-button>
                  </div>
                </template>
              </el-table-column>
            </el-table>
          </div>
          <div class="pagination-box">
            <el-pagination
              layout="total, sizes, prev, pager, next, jumper"
              :current-page="page"
              :page-size="pageSize"
              :page-sizes="[10, 30, 50, 100]"
              :total="total"
              @current-change="handleCurrentChange"
              @size-change="handleSizeChange"
            />
          </div>
        </div>
      </div>
    </div>
    <!--    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogFormVisible" :show-close="false"-->
    <!--               :before-close="closeDialog">-->
    <!--      <template #header>-->
    <!--        <div class="flex justify-between items-center">-->
    <!--          <span class="text-lg">{{ type === 'create' ? '新增' : '编辑' }}</span>-->
    <!--          <div>-->
    <!--            <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>-->
    <!--            <el-button @click="closeDialog">取 消</el-button>-->
    <!--          </div>-->
    <!--        </div>-->
    <!--      </template>-->


    <!--    </el-drawer>-->
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
        :menu="menuId"
        :formData="formData"
        :stepType="type"
        @close="closeDialog"
      />
    </el-dialog>
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true"
               :before-close="closeDetailShow" title="查看">
      <el-descriptions :column="1" border>
        <el-descriptions-item label="接口名称">
          {{ detailForm.name }}
        </el-descriptions-item>
        <el-descriptions-item label="变量">
          {{ detailForm.variables }}
        </el-descriptions-item>
        <el-descriptions-item label="参数">
          {{ detailForm.parameters }}
        </el-descriptions-item>
        <el-descriptions-item label="设置钩子">
          <ArrayCtrl v-model="detailForm.setup_hooks" />
        </el-descriptions-item>
        <el-descriptions-item label="清理钩子">
          <ArrayCtrl v-model="detailForm.teardown_hooks" />
        </el-descriptions-item>
        <el-descriptions-item label="提取">
          {{ detailForm.extract }}
        </el-descriptions-item>
        <el-descriptions-item label="验证器">
          <ArrayCtrl v-model="detailForm.validate" />
        </el-descriptions-item>
        <el-descriptions-item label="步骤导出">
          <ArrayCtrl v-model="detailForm.export" />
        </el-descriptions-item>
        <el-descriptions-item label="循环次数">
          {{ detailForm.loops }}
        </el-descriptions-item>
        <el-descriptions-item label="忽略弹出窗口">
          {{ detailForm.ignore_popup }}
        </el-descriptions-item>
      </el-descriptions>
    </el-drawer>
    <CurlImport ref="curlImportRef" @success="handleCurlSuccess" />
  </div>
</template>

<script setup>
  import {
    createAutoStep,
    deleteAutoStep,
    deleteAutoStepByIds,
    updateAutoStep,
    findAutoStep,
    getAutoStepList
  } from '@/api/automation/autostep'
  // 数组控制组件
  import ArrayCtrl from '@/components/arrayCtrl/arrayCtrl.vue'
  import ApiMenu from '@/components/platform/menu/index.vue'
  import stepForm from '@/components/platform/step/index.vue'
  // 全量引入格式化工具 请按需保留
  import {
    getDictFunc,
    formatDate,
    formatBoolean,
    filterDict,
    filterDataSource,
    returnArrImg,
    onDownloadFile
  } from '@/utils/format'
  import Runner from '@/components/platform/Runner.vue'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { ref, reactive, computed } from 'vue'
  import { useAppStore } from '@/pinia'
  import CurlImport from '@/components/curlImport/index.vue'
  import { useRoute } from 'vue-router'

  const route = useRoute()
  const appStore = useAppStore()

  // 计算当前平台类型（默认为 api）
  const currentPlatform = computed(() => route.meta?.type || 'api')
  
  // 计算当前 MenuType（默认为 1）
  const currentMenuType = computed(() => {
      // 优先使用 meta 定义的 menuType
      if (route.meta?.menuType) return String(route.meta.menuType)
      // 否则根据 type 推断
      const typeMap = {
          'api': '1',
          'android': '100',
          'ios': '200',
          'harmony': '300',
          'browser': '400'
      }
      return typeMap[currentPlatform.value] || '1'
  })

  defineOptions({
    name: 'AutoStep'
  })

  // 提交按钮loading
  const btnLoading = ref(false)
  // appStore 已在上方定义，此处移除重复声明

  // CURL 导入相关
  const curlImportRef = ref(null)
  const openCurlImport = () => {
    if (curlImportRef.value) {
      curlImportRef.value.open()
    } else {
      ElMessage.error('CURL 导入组件未加载完成，请稍后重试')
    }
  }
  
  const handleCurlSuccess = (parsedData) => {
    // 重置 formData
    formData.value = {
      name: '',
      loops: 0,
      retry: 0,
      request: {
        method: parsedData.method,
        url: parsedData.url,
        header_temp: parsedData.headers,
        headers: {}, // 实际保存时会根据 header_temp 生成
        json: parsedData.json,
        param_temp: parsedData.params,
        params: {}, // 实际保存时会根据 param_temp 生成
        data_warehouse: {}
      },
      parameters: {},
      parameters_temp: {},
      menu: 0
    }
    
    // 打开新增弹窗
    type.value = 'create'
    dialogFormVisible.value = true
  }

  // 控制更多查询条件显示/隐藏状态
  const showAllQuery = ref(false)

  // 自动化生成的字典（可能为空）以及字段
  const formData = ref({
    name: '',
    // variables: null,
    // parameters: null,
    // setup_hooks: [],
    // teardown_hooks: [],
    // extract: null,
    // validate: [],
    // export: [],
    loops: 0
    // ignore_popup: false
  })


  // 验证规则
  const rule = reactive({
    name: [{
      required: true,
      message: '',
      trigger: ['input', 'blur']
    },
      {
        whitespace: true,
        message: '不能只输入空格',
        trigger: ['input', 'blur']
      }
    ]
  })

  const elFormRef = ref()
  const elSearchFormRef = ref()

  // =========== 表格控制部分 ===========
  const page = ref(1)
  const total = ref(0)
  const pageSize = ref(10)
  const tableData = ref([])
  const searchInfo = ref({})
  // 排序
  const sortChange = ({ prop, order }) => {
    const sortMap = {
      CreatedAt: 'created_at',
      ID: 'id',
      name: 'step_name'
    }

    let sort = sortMap[prop]
    if (!sort) {
      sort = prop.replace(/[A-Z]/g, match => `_${match.toLowerCase()}`)
    }

    searchInfo.value.sort = sort
    searchInfo.value.order = order
    getTableData()
  }
  // 重置
  const onReset = () => {
    searchInfo.value = {}
    getTableData()
  }

  // 搜索
  const onSubmit = () => {
    elSearchFormRef.value?.validate(async (valid) => {
      if (!valid) return
      page.value = 1
      if (searchInfo.value.ignore_popup === '') {
        searchInfo.value.ignore_popup = null
      }
      getTableData()
    })
  }

  // 分页
  const handleSizeChange = (val) => {
    pageSize.value = val
    getTableData()
  }

  // 修改页面容量
  const handleCurrentChange = (val) => {
    page.value = val
    getTableData()
  }

  // 查询
  const getTableData = async () => {
    const table = await getAutoStepList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
    if (table.code === 0) {
      tableData.value = table.data.list
      total.value = table.data.total
      page.value = table.data.page
      pageSize.value = table.data.pageSize
    }
  }

  getTableData()

  // ============== 表格控制部分结束 ===============

  // 获取需要的字典 可能为空 按需保留
  const setOptions = async () => {
  }

  // 获取需要的字典 可能为空 按需保留
  setOptions()


  // 多选数据
  const multipleSelection = ref([])
  // 多选
  const handleSelectionChange = (val) => {
    multipleSelection.value = val
  }

  // 删除行
  const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(() => {
      deleteAutoStepFunc(row)
    })
  }

  // 多选删除
  const onDelete = async () => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(async () => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
      multipleSelection.value.map(item => {
        IDs.push(item.ID)
      })
      const res = await deleteAutoStepByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
    })
  }

  // 行为控制标记（弹窗内部需要增还是改）
  const type = ref('')

  // 更新行
  const updateAutoStepFunc = async (row) => {
    const res = await findAutoStep({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
      formData.value = res.data
      dialogFormVisible.value = true
    }
  }

  const copyAutoStepFunc = async (row) => {
    const res = await findAutoStep({ ID: row.ID })
    type.value = 'copy'
    if (res.code === 0) {
      formData.value = res.data
      formData.value.parentId = formData.value.ID
      formData.value.ID = 0
      formData.value.request.ID = 0
      dialogFormVisible.value = true
    }
  }


  // 删除行
  const deleteAutoStepFunc = async (row) => {
    const res = await deleteAutoStep({ ID: row.ID })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      if (tableData.value.length === 1 && page.value > 1) {
        page.value--
      }
      getTableData()
    }
  }

  // 弹窗控制标记
  const dialogFormVisible = ref(false)

  // 打开弹窗
  const openDialog = (v) => {
    type.value = v
    dialogFormVisible.value = true
  }

  // 关闭弹窗
  const closeDialog = (close) => {
    getTableData()
    formData.value = {}
    dialogFormVisible.value = false

  }
  // 弹窗确定
  const enterDialog = async () => {
    btnLoading.value = true
    formData.value.menu = menuId.value
    elFormRef.value?.validate(async (valid) => {
      if (!valid) return btnLoading.value = false
      let res
      switch (type.value) {
        case 'create':
          res = await createAutoStep(formData.value)
          break
        case 'update':
          res = await updateAutoStep(formData.value)
          break
        default:
          res = await createAutoStep(formData.value)
          break
      }
      btnLoading.value = false
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '创建/更改成功'
        })
        closeDialog()
        getTableData()
      }
    })
  }

  const detailForm = ref({})

  // 查看详情控制标记
  const detailShow = ref(false)


  // 打开详情弹窗
  const openDetailShow = () => {
    detailShow.value = true
  }


  // 打开详情
  const getDetails = async (row) => {
    // 打开弹窗
    const res = await findAutoStep({ ID: row.ID })
    if (res.code === 0) {
      detailForm.value = res.data
      openDetailShow()
    }
  }


  // 关闭详情弹窗
  const closeDetailShow = () => {
    detailShow.value = false
    detailForm.value = {}
  }

  const menuId = ref(null)
  const handleMenuClick = (id) => {
    menuId.value = id
    // 可以在这里添加根据菜单ID筛选表格数据的逻辑
    searchInfo.value.menu = id
    getTableData()
  }

</script>

<style scoped>
.autostep-container {
  display: flex;
  width: 100%;
  height: 100%;
  min-height: 600px;
  background-color: var(--el-bg-color-page);
  padding: 16px;
  box-sizing: border-box;
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

.main-panel {
  flex: 1;
  width: 100%;
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

.panel-content {
  flex: 1;
  -webkit-box-flex: 1;
  -ms-flex: 1;
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

.table-container {
  flex: 1;
  -webkit-box-flex: 1;
  -ms-flex: 1;
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
  flex-wrap: wrap;
  width: 100%;
}

.search-item {
  margin-bottom: 8px !important;
  margin-right: 12px !important;
}

.action-item {
  margin-bottom: 8px !important;
  margin-right: 0 !important;
}

.action-buttons {
  margin-bottom: 12px;
  display: flex;
  align-items: center;
  gap: 8px;
}

.table-wrapper {
  flex: 1;
  -webkit-box-flex: 1;
  -ms-flex: 1;
  overflow: hidden;
  border-radius: 4px;
  border: 1px solid var(--el-border-color-light);
}

.pagination-box {
  margin-top: 12px;
  padding: 0 4px;
  text-align: right;
}

/* Table Row Block Styles - Keep unchanged */
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

.table-button {
  margin-right: 8px;
}

:deep(.el-table__inner-wrapper::before) {
  display: none;
}
</style>
