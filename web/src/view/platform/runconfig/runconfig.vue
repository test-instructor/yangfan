<template>
  <div>
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

        <el-form-item label="名称" prop="name">
          <el-input v-model="searchInfo.name" placeholder="搜索条件" />
        </el-form-item>


        <template v-if="showAllQuery">
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
      >
        <el-table-column type="selection" width="55" />

        <el-table-column sortable align="left" label="日期" prop="CreatedAt" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>

        <el-table-column align="left" label="名称" prop="name" width="180" />

        <el-table-column align="left" label="域名" prop="base_url" width="240" />

        <el-table-column align="left" label="前置步骤" prop="preparatorySteps" width="120" />

        <el-table-column align="left" label="前置步骤ID" prop="setup_case_id" width="120" />

        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
          <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateRunConfigFunc(scope.row)">
              编辑
            </el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
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
    <el-drawer destroy-on-close
               v-model="dialogFormVisible"
               :show-close="false"
               :before-close="closeDialog"
               size="1200px"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ type === 'create' ? '新增' : '编辑' }}</span>
          <div class="flex gap-2">
            <EnvSelector />
            <PythonFuncSelector />
            <DataWarehouseFieldSelector :current-type="formData.data_warehouse_temp?.type" />
            <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
            <el-button @click="closeDialog">取 消</el-button>
          </div>
        </div>
      </template>
      <el-form
        :model="formData"
        label-position="left"
        ref="elFormRef"
        :rules="rule"
        label-width="100px"
        :inline="true"
        class="dialog-form"
      >
        <el-form-item prop="name" class="form-item-width">
          <el-input
            v-model="formData.name"
            :clearable="false"
            placeholder="请输入名称"
            style="width: 300px"
          >
            <template #prepend>名称</template>
          </el-input>
        </el-form-item>

        <el-form-item prop="base_url" class="form-item-width">
          <el-input
            v-model="formData.base_url"
            :clearable="false"
            placeholder="请输入域名"
            style="width: 100%"
          >
            <template #prepend>域名</template>
          </el-input>
        </el-form-item>

        <el-form-item label="超时:" prop="timeout" class="form-item-width">
          <el-input-number
            v-model="formData.timeout"
            :precision="2"
            :clearable="false"
            style="width: 100%"
          />
        </el-form-item>

        <el-form-item label="允许重定向:" prop="allow_redirects" class="form-item-width">
          <el-switch
            v-model="formData.allow_redirects"
            active-color="#13ce66"
            inactive-color="#ff4949"
            active-text="是"
            inactive-text="否"
            clearable
          />
        </el-form-item>

        <el-form-item label="启用证书验证:" prop="verify" class="form-item-width">
          <el-switch
            v-model="formData.verify"
            active-color="#13ce66"
            inactive-color="#ff4949"
            active-text="是"
            inactive-text="否"
            clearable
          />
        </el-form-item>

        <el-form-item label="前置步骤:" prop="preparatorySteps" class="form-item-width">
          <el-input
            v-model.number="formData.preparatorySteps"
            :clearable="false"
            placeholder="请输入前置步骤"
            style="width: 100%"
          />
        </el-form-item>

        <el-form-item label="重试次数:" prop="retry" class="form-item-width">
          <el-input
            v-model.number="formData.retry"
            :clearable="true"
            placeholder="请输入重试次数"
            style="width: 100%"
          />
        </el-form-item>
      </el-form>
      <div style="height: 400px;">
        <el-tabs v-model="activeConfig" class="demo-tabs" @tab-click="handleClick">
          <el-tab-pane label="请求头" name="header">
            <HeaderTable
              :header="formData.header_temp"
              :heights="400"
              @headerData="handleHeaderData"
            />
          </el-tab-pane>
          <el-tab-pane label="变量" name="variables">
            <Variables
              :variables="formData.variable_temp"
              :heights="400"
              @variablesData="handleVariablesData"
            />
          </el-tab-pane>
          <el-tab-pane label="参数设置" name="parameters">
            <ParameterTable
              :jsons="formData.parameters"
              :parametersTemp="formData.parameters_temp"
              @jsonData="handleJsonData"
              @tempData="(val) => formData.parameters_temp = val"
            />
          </el-tab-pane>
          <el-tab-pane label="数据仓库" name="datawarehouse">
            <DataWarehouseConfig
              v-model="formData.data_warehouse_temp"
              :height="400"
            />
          </el-tab-pane>
        </el-tabs>
      </div>
    </el-drawer>

  </div>
</template>

<script setup>
  import {
    createRunConfig,
    deleteRunConfig,
    deleteRunConfigByIds,
    updateRunConfig,
    findRunConfig,
    getRunConfigList
  } from '@/api/platform/runconfig'
  import ParameterTable from '@/components/platform/parameterTable/index.vue'
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
  import { ref, reactive } from 'vue'
  import { useAppStore } from '@/pinia'
  // 页面组件
  import PythonFuncSelector from '@/components/platform/button/PythonFuncSelector.vue'
  import EnvSelector from '@/components/platform/button/EnvDetail.vue'
  import DataWarehouseFieldSelector from '@/components/platform/button/DataWarehouseFieldSelector.vue'
  import HeaderTable from '@/components/platform/header/index.vue'
  import Variables from '@/components/platform/variables/index.vue'
  import JsonEditor from '@/components/platform/jsonEdit/index.vue'
  import DataWarehouseConfig from '@/components/platform/dataWarehouseConfig/index.vue'

  import utils from '@/utils/dataTypeConverter.js'

  defineOptions({
    name: 'RunConfig'
  })

  // 提交按钮loading
  const btnLoading = ref(false)
  const appStore = useAppStore()

  // 控制更多查询条件显示/隐藏状态
  const showAllQuery = ref(false)

  // 自动化生成的字典（可能为空）以及字段
  const formData = ref({
    name: '',
    base_url: '',
    variables: {},
    headers: {},
    parameters: {},
    parameters_temp: {},
    data_warehouse: {},
    data_warehouse_temp: {},
    variable_temp: {},
    header_temp: {},
    weight: undefined,
    timeout: 0,
    allow_redirects: false,
    verify: false,
    preparatorySteps: undefined,
    setup_case_id: null,
    report_id: undefined,
    retry: undefined
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
      if (searchInfo.value.allow_redirects === '') {
        searchInfo.value.allow_redirects = null
      }
      if (searchInfo.value.verify === '') {
        searchInfo.value.verify = null
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
    const table = await getRunConfigList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
      deleteRunConfigFunc(row)
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
      const res = await deleteRunConfigByIds({ IDs })
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
  const updateRunConfigFunc = async (row) => {
    const res = await findRunConfig({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
      formData.value = res.data
      dialogFormVisible.value = true
    }
  }


  // 删除行
  const deleteRunConfigFunc = async (row) => {
    const res = await deleteRunConfig({ ID: row.ID })
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
      base_url: '',
      variables: {},
      headers: {},
      parameters: {},
      parameters_temp: {},
      data_warehouse: {},
      data_warehouse_temp: {},
      variable_temp: {},
      header_temp: {},
      weight: undefined,
      timeout: 0,
      allow_redirects: false,
      verify: false,
      export: [],
      preparatorySteps: undefined,
      setup_case_id: null,
      report_id: undefined,
      retry: undefined
    }
  }
  // 弹窗确定
  const enterDialog = async () => {
    const variablesData = utils.convertData(variables.value)
    if (!variablesData.success) {
      ElMessageBox.alert(
        '数据处理失败字段:' + variablesData.errors.key + ',出现错误:' + variablesData.errors.error,
        '变量数据处理失败',
        {
          confirmButtonText: '确认',
          type: 'error',
          dangerouslyUseHTMLString: true,
          center: true
        }
      )
      return
    }
    formData.value.variable_temp = variables.value
    formData.value.variables = variablesData.data

    const headersData = utils.processData(headers.value)
    if (!headersData.success) {
      ElMessageBox.alert(
        '数据处理失败:' + headersData.errors.join('、') + ',字段重复',
        '请求头数据处理失败',
        {
          confirmButtonText: '确认',
          type: 'error',
          dangerouslyUseHTMLString: true,
          center: true
        }
      )
      return
    }
    formData.value.header_temp = headers.value
    formData.value.headers = headersData.data

    if (jsonError.value != '') {
      ElMessageBox.alert(
        '参数设置数据处理失败:' + jsonError.value,
        '参数设置数据处理失败',
        {
          confirmButtonText: '确认',
          type: 'error',
          dangerouslyUseHTMLString: true,
          center: true
        }
      )
      return
    }

    // 同步 data_warehouse_temp 到 data_warehouse
    if (formData.value.data_warehouse_temp && Object.keys(formData.value.data_warehouse_temp).length > 0) {
      formData.value.data_warehouse = formData.value.data_warehouse_temp
    } else {
      formData.value.data_warehouse = {}
    }

    btnLoading.value = true
    elFormRef.value?.validate(async (valid) => {
      if (!valid) return btnLoading.value = false
      let res
      switch (type.value) {
        case 'create':
          res = await createRunConfig(formData.value)
          break
        case 'update':
          res = await updateRunConfig(formData.value)
          break
        default:
          res = await createRunConfig(formData.value)
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

  const activeConfig = ref('header')
  // 请求头数据接收
  const headers = ref([])
  const handleHeaderData = (tableData) => {
    headers.value = tableData
  }
  // 变量数据接收
  const variables = ref([])
  const handleVariablesData = (tableData) => {
    variables.value = tableData
  }
  // 参数设置数据接收
  const jsonError = ref('')
  const handleJsonData = (result) => {
    if (!result.isValid) {
      jsonError.value = result.error.message // 展示错误信息
    } else {
      jsonError.value = '' // 清空错误信息
      formData.value.parameters = result.data
    }
  }
</script>

<style>
  /* 弹窗表单整体样式：控制内边距，避免内容贴边 */
  .dialog-form {
    padding: 5px;
    display: flex;
    flex-wrap: wrap;
    gap: 5px;
  }

  /* 单个表单项宽度：根据弹窗宽度调整，建议2-3个一行 */
  .form-item-width {
    width: calc(33.33% - 12px);
    min-width: 320px;
  }

  /* 适配小屏幕：屏幕小于768px时，2个表单项一行 */
  @media (max-width: 768px) {
    .form-item-width {
      width: calc(50% - 8px);
    }
  }

  /* 适配超小屏幕：屏幕小于480px时，1个表单项一行 */
  @media (max-width: 480px) {
    .form-item-width {
      width: 100%;
    }
  }
</style>

