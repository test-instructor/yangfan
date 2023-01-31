<template>
  <div>
    <!--
    <div class="gva-search-box">
        <el-form-item>
          <el-form-item label="包名称:" prop="name" style="margin-right: 10px;">
            <el-input v-model="search.name" :clearable="true" placeholder="请输入"/>
          </el-form-item>
          <el-form-item >
          <el-button size="small" type="primary" icon="search" @click="onSubmit" >查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
        </el-form-item>
    </div>
    -->
    <div class="gva-table-box">
      <div class="gva-btn-list" style="float:right;">
        <el-button size="small" type="primary" icon="plus" @click="openDialog" style="margin-right: 100px;">安装
        </el-button>
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要卸载吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button size="small" type="primary" link @click="deleteVisible = false">取消</el-button>
            <el-button size="small" type="primary" @click="onDelete">确定</el-button>
          </div>
        </el-popover>
      </div>
      <el-table
          v-fit-columns
          ref="multipleTable"
          style="width: 100%"
          tooltip-effect="dark"
          :data="tableData"
          row-key="ID"
          @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55"/>
        <!-- <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column> -->
        <el-table-column align="center" label="包名称" prop="name"/>
        <el-table-column align="center" label="包版本" prop="version"/>
        <el-table-column align="center" label="操作">
          <template #default="scope">
            <el-button type="primary" link icon="icon-upgrade" size="small" @click="updateRow(scope.row)">升级
            </el-button>
            <el-button v-if="scope.row.isUninstall" type="primary" link icon="delete" size="small"
                       @click="deleteRow(scope.row)">卸载
            </el-button>
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="版本更新">
      <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item label="包名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入"/>
        </el-form-item>
        <el-form-item label="包版本:" prop="version">
          <el-input v-model="formData.version" :clearable="true" placeholder="请输入"/>
        </el-form-item>
        <el-form-item label="是否可以卸载:" prop="isUninstall" label-width="120px">
<!--      formData.isUninstall为否则入参0    -->

          <el-switch v-model="formData.isUninstall" active-color="#13ce66" inactive-color="#ff4949" active-text="是"
                     inactive-text="否" clearable></el-switch>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'HrpPyPkg'
}
</script>

<script setup>
import {
  installHrpPyPkg,
  uninstallHrpPyPkg,
  uninstallHrpPyPkgByIds,
  updateHrpPyPkg,
  searchHrpPyPkg,
  getHrpPyPkgList
} from '@/api/py_pkg'

// 全量引入格式化工具 请按需保留
import {getDictFunc, formatDate, formatBoolean, filterDict} from '@/utils/format'
import {ElMessage, ElMessageBox} from 'element-plus'
import {ref, reactive} from 'vue'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  name: '',
  version: '',
  isUninstall: true,
})


const search = ref({
  name: '',
})

// 验证规则
const rule = reactive({
  name: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],

})

const elFormRef = ref()


// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})


// 搜索
const onSubmit = async () => {
  const table = await searchHrpPyPkg({page: page.value, pageSize: pageSize.value, search})
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
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
  const table = await getHrpPyPkgList({page: page.value, pageSize: pageSize.value, ...searchInfo.value})
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
    uninstallHrpPyPkgFunc(row)
  })
}

// 更新行
const updateRow = (row) => {
  ElMessageBox.confirm('更新前请确认版本兼容性，确定要更新吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'info'
  }).then(() => {
    updateHrpPyPkgFunc(row)
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
  const res = await uninstallHrpPyPkgByIds({ids})
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateHrpPyPkgFunc = async (row) => {
  const res = await updateHrpPyPkg({name: row.name})
  // console.log(res)
  type.value = 'update'
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '升级成功'
    })
  }
  getTableData()
}

// 卸载python包
const uninstallHrpPyPkgFunc = async (row) => {
  const res = await uninstallHrpPyPkg({name: row.name})
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '卸载成功'
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
    version: '',
    isUninstall: true,
  }
}
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    res = await installHrpPyPkg(formData.value)
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '安装成功'
      })
      closeDialog()
      getTableData()
    }
  })
}
</script>

<style>
</style>
