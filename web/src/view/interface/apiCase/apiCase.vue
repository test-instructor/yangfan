<template>
  <div class="parent">

    <div class="right">
      <div class="left">
        <InterfaceTree
            menutype=3
            @getTreeID="setTreeID"
            eventType="1"
        ></InterfaceTree>
      </div>
      <div class="right2">
        <div class="gva-search-box">
          <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
            <el-form-item label="用例名称">
              <el-input v-model="searchInfo.name" placeholder="搜索条件"/>
            </el-form-item>
            <el-form-item>
              <el-button icon="search" size="mini" type="primary" @click="onSubmit">查询</el-button>
              <el-button icon="refresh" size="mini" @click="onReset">重置</el-button>
            </el-form-item>
          </el-form>
        </div>
        <div class="gva-table-box">
          <div class="gva-btn-list">
            <el-button icon="plus" size="mini" type="primary" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" placement="top" width="160">
              <p>确定要删除吗？</p>
              <div style="text-align: right; margin-top: 8px;">
                <el-button size="mini" type="text" @click="deleteVisible = false">取消</el-button>
                <el-button size="mini" type="primary" @click="onDelete">确定</el-button>
              </div>
              <template #reference>
                <el-button :disabled="!multipleSelection.length" icon="delete" size="mini" style="margin-left: 10px;">删除
                </el-button>
              </template>
            </el-popover>
          </div>
          <el-table
              ref="multipleTable"
              :data="tableData"
              row-key="ID"
              style="width: 100%"
              tooltip-effect="dark"
              @selection-change="handleSelectionChange"
          >
            <el-table-column type="selection" width="55"/>
    <!--        <el-table-column align="left" label="日期" width="180">-->
    <!--          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>-->
    <!--        </el-table-column>-->
            <el-table-column align="left" label="用例名称" prop="name" width="120"/>
            <el-table-column align="left" label="运行配置" prop="runConfig.name" width="180"/>
            <!--        <el-table-column align="left" label="测试用例集" min-width="80">-->
    <!--          <template #default="scope">-->
    <!--            <el-cascader-->
    <!--                v-model="scope.row.caseIds"-->
    <!--                :clearable="false"-->
    <!--                :options="caseOptions"-->
    <!--                :props="{ multiple:true,checkStrictly: true, label:'caseName',value:'caseId', disabled:'disabled', emitPath:false}"-->
    <!--                :show-all-levels="false"-->
    <!--                collapse-tags-->
    <!--                @visible-change="(flag)=>{changeCase(scope.row,flag)}"-->
    <!--                @remove-tag="()=>{changeCase(scope.row,false)}"-->
    <!--            />-->
    <!--          </template>-->
    <!--        </el-table-column>-->
            <el-table-column align="left" label="备注" prop="describe" width="120"/>
    <!--        <el-table-column align="left" label="状态" prop="status" width="120">-->
    <!--          <template #default="scope">-->
    <!--            <el-tag :type="scope.row.status ? 'success' : 'info'">{{ scope.row.status ? '启用' : '禁用' }}</el-tag>-->
    <!--          </template>-->
    <!--        </el-table-column>-->
            <el-table-column align="left" label="按钮组" width="360">
              <template #default="scope">
                <el-button class="table-button" icon="detail" size="small" type="text" @click="detailApisCaseFunc(scope.row)">用例详情</el-button>
                <el-button class="table-button" icon="detail" size="small" type="text" @click="runCase(scope.row)">后台运行</el-button>
                <el-button class="table-button" icon="edit" size="small" type="text"
                           @click="updateApiCaseFunc(scope.row)">变更
                </el-button>
                <el-button icon="delete" size="mini" type="text" @click="deleteRow(scope.row)">删除</el-button>
              </template>
            </el-table-column>
          </el-table>
          <div class="gva-pagination">
            <el-pagination
                :current-page="page"
                :page-size="pageSize"
                :page-sizes="[10, 30, 50, 100]"
                :total="total"
                layout="total, sizes, prev, pager, next, jumper"
                @current-change="handleCurrentChange"
                @size-change="handleSizeChange"
            />
          </div>
        </div>
      </div>
    </div>
    <el-dialog
        v-model="dialogFormVisible"
        :before-close="closeDialog"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="用例名称:">
          <el-input v-model="formData.name" clearable placeholder="请输入"/>
        </el-form-item>
<!--        <el-form-item label="状态:">-->
<!--          <el-switch v-model="formData.status" active-color="#13ce66" active-text="启用" clearable-->
<!--                     inactive-color="#ff4949" inactive-text="禁用"></el-switch>-->
<!--        </el-form-item>-->
        <el-form-item label="运行配置:">
          <el-select
              v-model="configID"
              placeholder="请选择"
              @change="configChange"
          >
            <el-option
                v-for="item in configData"
                :key="item.ID"
                :label="item.name"
                :value="item.ID"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="备注:">
          <el-input v-model="formData.describe" clearable placeholder="请输入"/>
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
  name: 'ApiCase'
}
</script>

<script setup>
import {
  createApiCase,
  deleteApiCase,
  deleteApiCaseByIds,
  updateApiCase,
  findApiCase,
  getApiCaseList,
  setApisCase
} from '@/api/apiCase'

// 全量引入格式化工具 请按需保留
import {getDictFunc, formatDate, formatBoolean, filterDict} from '@/utils/format'
import {ElMessage, ElMessageBox} from 'element-plus'
import {ref, watch, nextTick} from 'vue'
import InterfaceTree from '@/view/interface/interfaceComponents/interfaceTree.vue'

// import apiCaseCron from '@/view/interface/apiCase/apiCaseCron.vue'
import {getApiConfigList} from "@/api/apiConfig";
import {runApiCase} from "@/api/runTestCase";
import {useRouter} from "vue-router";

const router = useRouter()

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  name: '',
  runTime: '',
  nextRunTime: new Date(),
  status: false,
  describe: '',
  runNumber: 0,
  case: [],
  runConfig: {
    ID: 0
  },
})
const cronVisible = ref(false)
const cronFun = () => {
  cronVisible.value = true
}
const creatCron = ref(false)
const runTimeCron = (cron) => {
  formData.value.runTime = cron
}
const closeRunTimeCron = (isClose) => {
  cronVisible.value = isClose

}
const configID = ref()
const configChange = (key) => {
  formData.value.runConfig.ID = key
}
let treeID = 0

const setTreeID = (val) => {
  treeID = val
  getTableData()
}

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
  if (searchInfo.value.status === "") {
    searchInfo.value.status = null
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
const configData = ref([])
const getConfigData = async () => {
  const config = await getApiConfigList({page: 1, pageSize: 99999})
  if (config.code === 0) {
    configData.value = config.data.list
  }
}

// 查询
const getTableData = async () => {
  let menu = treeID
  const table = await getApiCaseList({menu,page: page.value, pageSize: pageSize.value, ...searchInfo.value})
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}


watch(tableData, () => {
  setCaseIds()
})

const setCaseIds = () => {
  tableData.value && tableData.value.forEach((testCase) => {
    const caseIds = testCase.case && testCase.case.map(i => {
      return i.ID
    })
    testCase.caseIds = caseIds
  })
}

const changeCase = async (row, flag) => {
  if (flag) {
    return
  }
  await nextTick()
  const res = await setApisCase({
    ID: row.ID,
    caseIds: row.caseIds
  })
  if (res.code === 0) {
    ElMessage({type: 'success', message: '用例设置成功'})
  } else {
    ElMessage({type: 'success', message: '用例设置失败'})
  }
}


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
    deleteApiCaseFunc(row)
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
  const res = await deleteApiCaseByIds({ids})
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

const detailApisCaseFunc = (row) => {
  if (row) {
    router.push({
      name: 'apisCaseDetail', params: {
        id: row.ID
      }
    })
  } else {
    router.push({name: 'apisCaseDetail'})
  }
}

const runCase = async (row) => {
  let data = {caseID: Number(row.ID)}
  const res = await runApiCase(data)
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '运行成功'
    })
  }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateApiCaseFunc = async (row) => {
  const res = await findApiCase({ID: row.ID})
  await getConfigData()
  creatCron.value = true
  // configID.value = configData.value.runConfig.ID
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.retestCase
    configID.value = formData.value.runConfig.ID
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteApiCaseFunc = async (row) => {
  const res = await deleteApiCase({ID: row.ID})
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
  getConfigData()
  type.value = 'create'
  creatCron.value = true
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  configID.value = ''
  dialogFormVisible.value = false
  formData.value = {
    name: '',
    runTime: '',
    nextRunTime: new Date(),
    status: false,
    describe: '',
    runNumber: 0,
    runConfig: {ID: 0},
  }
  creatCron.value = false
}
// 弹窗确定
const enterDialog = async () => {
  if (formData.value.name===''){
    ElMessage({
      type: 'error',
      message: '任务名称不能为空'
    })
    return
  }
  // params.menu = treeID
  if (formData.value.runConfig.ID===0){
    ElMessage({
      type: 'error',
      message: '请选择运行配置'
    })
    return
  }
  if (formData.value.status && formData.value.runTime ===""){
    ElMessage({
      type: 'error',
      message: '开启定时任务时时间配置不能为空'
    })
    return
  }
  let res
  let params = {menu: treeID}
  formData.value.case = []
  switch (type.value) {
    case 'create':
      res = await createApiCase(formData.value, params)
      break
    case 'update':
      res = await updateApiCase(formData.value, params)
      break
    default:
      res = await createApiCase(formData.value, params)
      break
  }
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '创建/更改成功'
    })
    closeDialog()
    getTableData()
    creatCron.value = false
  }
}
</script>

<style>
.parent {
  width: 100%;
  height: 85%;
}

.left {
  margin-top: 10px;
  width: 300px;
  height: 98%;
  padding: 8px;
  background: #ffffff;
}

.right {
  display: flex;
  flex: 1;
  height: 100%;
  margin-right: 10px;
}

.right2 {
  flex: 1;
  height: 100%;
  padding: 10px;
}
</style>
