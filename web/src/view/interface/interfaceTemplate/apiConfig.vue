<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="配置名称">
          <el-input v-model="searchInfo.name" placeholder="搜索条件"/>
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="mini" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="mini" type="primary" icon="plus" @click="openDialog">新增</el-button>
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button size="mini" type="text" @click="deleteVisible = false">取消</el-button>
            <el-button size="mini" type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button icon="delete" size="mini" style="margin-left: 10px;" :disabled="!multipleSelection.length">删除
            </el-button>
          </template>
        </el-popover>
      </div>
      <el-table
          ref="multipleTable"
          style="width: 100%"
          tooltip-effect="dark"
          :data="tableData"
          row-key="ID"
          @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55"/>
        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="配置名称" prop="name" width="120"/>
        <el-table-column align="left" label="域名" prop="base_url" width="360"/>
        <el-table-column align="left" label="默认配置" prop="default" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.default) }}</template>
        </el-table-column>
        <el-table-column align="left" label="按钮组">
          <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button"
                       @click="updateApiConfigFunc(scope.row)">变更
            </el-button>
            <el-button type="text" icon="delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
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
    <el-dialog
        v-model="dialogFormVisible"
        :before-close="closeDialog"
        :visible.sync="dialogFormVisible"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        :title="dialogTitle"
        width="1380px"
        top="30px"
    >
      <apiConfigForm
          @close="closeDialog"
          v-if="dialogFormVisible"
          :heights="heightDiv"
          :eventType="type"
          :formData="formDatas"
          ref="menuRole">
      </apiConfigForm>

    </el-dialog>

  </div>
</template>

<script>
export default {
  name: 'ApiConfig'
}
</script>

<script setup>
import {
  createApiConfig,
  deleteApiConfig,
  deleteApiConfigByIds,
  updateApiConfig,
  findApiConfig,
  getApiConfigList
} from '@/api/apiConfig'

// 全量引入格式化工具 请按需保留
import {getDictFunc, formatDate, formatBoolean, filterDict} from '@/utils/format'
import {ElMessage, ElMessageBox} from 'element-plus'
import {ref, reactive} from 'vue'
import apiConfigForm from '@/view/interface/interfaceComponents/apiConfigForm.vue'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  name: '',
  base_url: '',
  default: false,
})

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  if (searchInfo.value.default === "") {
    searchInfo.value.default = null
  }
  getTableData()
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
  const table = await getApiConfigList({page: page.value, pageSize: pageSize.value, ...searchInfo.value})
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

const formDatas = reactive({
  name: '',
  base_url: '',
  headers: '',
  variables: '',
  extract: '',
  validate: '',
  hooks: '',
  apiMenuID: '',
  Parameters: '',
})

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
    deleteApiConfigFunc(row)
  })
}


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async () => {
  const ids = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要删除的数据'
    })
    return
  }
  multipleSelection.value &&
  multipleSelection.value.map(item => {
    ids.push(item.ID)
  })
  const res = await deleteApiConfigByIds({ids})
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    await getTableData()
  }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateApiConfigFunc = async (row) => {
  const res = await findApiConfig({ID: row.ID})
  type.value = 'update'
  dialogTitle.value = '编辑配置'
  if (res.code === 0) {
    formData.value = res.data.reac
    formDatas.value = res.data.reac
    dialogFormVisible.value = true
  }
}


// 删除行
const deleteApiConfigFunc = async (row) => {
  const res = await deleteApiConfig({ID: row.ID})
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    await getTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)
const dialogTitle = ref(false)
const heightDiv = ref()
heightDiv.value = window.screen.height - 480

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogTitle.value = '新增配置'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    name: '',
    base_url: '',
    default: false,
  }
}
// 弹窗确定
const enterDialog = async () => {
  let res
  switch (type.value) {
    case 'create':
      res = await createApiConfig(formData.value)
      break
    case 'update':
      res = await updateApiConfig(formData.value)
      break
    default:
      res = await createApiConfig(formData.value)
      break
  }
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '创建/更改成功'
    })
    closeDialog()
    getTableData()
  }
}
</script>

<style>
</style>

