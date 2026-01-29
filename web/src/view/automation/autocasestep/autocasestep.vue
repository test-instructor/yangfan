<template>
  <div>
    <div style="display: flex; width: 100%; overflow: hidden;">
      <div style="width: 250px; flex-shrink: 0;">
        <ApiMenu
          :menutype="currentMenuType"
          @getTreeID="handleMenuClick"
        />
      </div>
      <div style="flex: 1; overflow-x: auto; min-width: 600px; padding-left: 16px;">
        <div class="gva-search-box">
          <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline"
                   @keyup.enter="onSubmit">
            <el-form-item label="创建日期" prop="createdAtRange">
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

            <el-form-item label="步骤名称" prop="name">
              <el-input v-model="searchInfo.name" placeholder="搜索条件" />
            </el-form-item>


            <template v-if="showAllQuery">
              <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
            </template>

            <el-form-item>
              <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
              <el-button icon="refresh" @click="onReset">重置</el-button>
              <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开
              </el-button>
              <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
            </el-form-item>
          </el-form>
        </div>
        <div class="gva-table-box">
          <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog()">新增</el-button>
            <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">
              删除
            </el-button>

          </div>
          <el-table
            ref="multipleTable"
            style="width: 100%"
            tooltip-effect="dark"
            :data="tableData"
            row-key="ID"
            @selection-change="handleSelectionChange"
            @sort-change="sortChange"
          >
            <el-table-column type="selection" width="55" />

            <el-table-column sortable align="left" label="日期" prop="CreatedAt" width="180">
              <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
            </el-table-column>

            <el-table-column sortable align="left" label="步骤名称" prop="name" width="120" />

            <el-table-column align="left" label="运行环境" prop="envName" width="120" />

            <el-table-column align="left" label="运行配置" prop="configName" width="120" />

            <el-table-column align="left" label="循环次数" prop="loops" width="120" />

            <el-table-column align="left" label="操作" fixed="right" min-width="320">
              <template #default="scope">
<!--                <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">-->
<!--                  <el-icon style="margin-right: 5px">-->
<!--                    <InfoFilled />-->
<!--                  </el-icon>-->
<!--                  查看-->
<!--                </el-button>-->
                <div style="display: flex; align-items: center;">
                  <Runner case_type="step" :case_id="scope.row.ID" style="margin-right: 10px;" />
                  <el-button type="primary" link icon="step-api" class="table-button" @click="openApiDetail(scope.row)">
                    {{ manageTitle }}
                  </el-button>
                  <el-button type="primary" link icon="edit" class="table-button" @click="updateAutoCaseStepFunc(scope.row)">
                    编辑
                  </el-button>
                  <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
                </div>
              </template>
            </el-table-column>
          </el-table>
          <div class="gva-pagination">
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
    <el-drawer destroy-on-close size="1300px" v-model="dialogFormVisible" :show-close="false"
               :close-on-press-escape="false"
               :close-on-click-modal="false"
               :before-close="closeDialog">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ type === 'create' ? '新增' : '编辑' }}</span>
          <div class="flex gap-2">
            <EnvSelector />
            <PythonFuncSelector />
            <DataWarehouseFieldSelector :current-type="formData?.data_warehouse_temp?.type" />
            <el-button :loading="btnLoading" type="primary" @click="enterDialog">保存</el-button>
            <el-button type="primary" @click="closeDialog">取消</el-button>
          </div>
        </div>
      </template>

      <AutoCaseStepForm
        ref="autoCaseStepForm"
        :menu="menuId"
        :formData="formData"
        :type="type"
        @close="handleSavedClose"
      />
    </el-drawer>
    <el-drawer
      :close-on-press-escape="false"
      :close-on-click-modal="false"
      size="95%"
      v-model="dialogApiDetail"
      :show-close="true"
      :before-close="closeApiDetail"
      :title="manageTitle"
    >
      <Autocasesteprequest :stepID="stepId" :stepName="stepName" :type="caseType"/>
    </el-drawer>

  </div>
</template>

<script setup>
  import {
    createAutoCaseStep,
    deleteAutoCaseStep,
    deleteAutoCaseStepByIds,
    updateAutoCaseStep,
    findAutoCaseStep,
    getAutoCaseStepList
  } from '@/api/automation/autocasestep'
  // 数组控制组件
  import ArrayCtrl from '@/components/arrayCtrl/arrayCtrl.vue'
  import AutocasestepForm  from '@/view/automation/autocasestep/autocasestepForm.vue'
  import AutocasestepRequest  from '@/view/automation/autocasestep/autocasesteprequest.vue'

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
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { ref, reactive, computed } from 'vue'
  import { useAppStore } from '@/pinia'
  import AutoCaseStepForm from '@/view/automation/autocasestep/autocasestepForm.vue'
  import ApiMenu from '@/components/platform/menu/index.vue'
  import Autocasesteprequest from '@/view/automation/autocasestep/autocasesteprequest.vue'

  import PythonFuncSelector from '@/components/platform/button/PythonFuncSelector.vue'
  import EnvSelector from '@/components/platform/button/EnvDetail.vue'
  import DataWarehouseFieldSelector from '@/components/platform/button/DataWarehouseFieldSelector.vue'
  import Runner from '@/components/platform/Runner.vue'

  import { useRoute } from 'vue-router'

  defineOptions({
    name: 'AutoCaseStep'
  })

  const route = useRoute()
  const caseType = computed(() => {
    return route.query.type || 'api'
  })

  const currentMenuType = computed(() => {
    if (caseType.value === 'api') {
      return '11'
    }
    return `casestep_${caseType.value}`
  })

  const manageTitle = computed(() => {
    return caseType.value === 'api' ? '接口管理' : '元素操作管理'
  })

  // 提交按钮loading
  const btnLoading = ref(false)
  const appStore = useAppStore()

  // 控制更多查询条件显示/隐藏状态
  const showAllQuery = ref(false)

  // 自动化生成的字典（可能为空）以及字段
  const formDataTemple = ref({
    "ID": 0,
    "name": "",
    "variables": {},
    "variables_temp": null,
    "parameters": {},
    "setup_hooks": [],
    "teardown_hooks": [],
    "extract": {},
    "extract_temp": [],
    "validate": [],
    "validators_temp": [],
    "export": null,
    "loops": 0,
    "ignore_popup": false,
    "sort": 0,
    "retry": 0,
    "request": {},
    "request_id": 0,
    "projectId": 0,
    "menu": 0,
    "parentId": 0
  })
  const formData = ref({})

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
    const table = await getAutoCaseStepList({ page: page.value, pageSize: pageSize.value,menu: menuId.value, ...searchInfo.value })
    if (table.code === 0) {
      tableData.value = table.data.list
      total.value = table.data.total
      page.value = table.data.page
      pageSize.value = table.data.pageSize
    }
  }

  // getTableData()

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
      deleteAutoCaseStepFunc(row)
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
      const res = await deleteAutoCaseStepByIds({ IDs })
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
  const updateAutoCaseStepFunc = async (row) => {
    const res = await findAutoCaseStep({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
      formData.value = res.data
      dialogFormVisible.value = true
    }
  }


  // 删除行
  const deleteAutoCaseStepFunc = async (row) => {
    const res = await deleteAutoCaseStep({ ID: row.ID })
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
  const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
  }

  // 关闭弹窗
  const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
      name: '',
      variables: null,
      parameters: null,
      setup_hooks: [],
      teardown_hooks: [],
      extract: null,
      validate: [],
      export: [],
      loops: 0,
      ignore_popup: false,
      data_warehouse: null,
      data_warehouse_temp: {}
    }
  }
  const handleSavedClose = () => {
    closeDialog()
    if (type.value === 'create') {
      page.value = 1
    }
    getTableData()
  }
  // 弹窗确定
  const autoCaseStepForm = ref(null)
  const enterDialog = async () => {
    btnLoading.value = true
    await autoCaseStepForm.value.save()
    btnLoading.value = false
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
    const res = await findAutoCaseStep({ ID: row.ID })
    if (res.code === 0) {
      detailForm.value = res.data
      openDetailShow()
    }
  }


  // 关闭详情弹窗
  const closeDetailShow = () => {
    detailShow.value = false
    detailForm.value = {}
    // formData.value = formDataTemple.value
  }

  const menuId = ref(null)
  const handleMenuClick = (id) => {
    menuId.value = id
    // 可以在这里添加根据菜单ID筛选表格数据的逻辑
    searchInfo.value.menu = id
    getTableData()
  }
  const stepId = ref(0)
  const stepName = ref('')
  const dialogApiDetail = ref(false)
  const openApiDetail = (row) => {
    console.log("==============",row)
    stepId.value = row.ID
    stepName.value = row.name
    console.log("==============123",stepName.value)
    dialogApiDetail.value = true
  }
  const closeApiDetail = () => {
    stepId.value = 0
    stepName.value = ''
    dialogApiDetail.value = false
  }

</script>

<style>

</style>
