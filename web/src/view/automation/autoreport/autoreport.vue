<template>
  <div class="auto-report-container">
    <div class="gva-search-box">
      <el-card shadow="hover" class="search-card">
        <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
          <el-form-item label="创建日期" prop="createdAtRange">
            <template #label>
              <span class="flex items-center">
                创建日期
                <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                  <el-icon class="ml-1"><QuestionFilled /></el-icon>
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
          
          <template v-if="showAllQuery">
            <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
          </template>

          <el-form-item>
            <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
            <el-button icon="refresh" @click="onReset">重置</el-button>
            <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
            <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
          </el-form-item>
        </el-form>
      </el-card>
    </div>

    <div class="gva-table-box mt-4">
      <el-card shadow="hover" class="table-card">
        <div class="gva-btn-list mb-4">
            <el-button type="primary" icon="plus" @click="openDialog()">新增</el-button>
            <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
        </div>
        <el-table
          ref="multipleTable"
          style="width: 100%"
          tooltip-effect="dark"
          :data="tableData"
          row-key="ID"
          @selection-change="handleSelectionChange"
          header-cell-class-name="table-header"
        >
          <el-table-column type="selection" width="55" />
          
          <el-table-column sortable align="left" label="日期" prop="CreatedAt" width="180">
              <template #default="scope">
                <div class="flex items-center">
                  <el-icon class="mr-1 text-gray-400"><Clock /></el-icon>
                  {{ formatDate(scope.row.CreatedAt) }}
                </div>
              </template>
          </el-table-column>
          
          <el-table-column align="left" label="名称" prop="name" min-width="150" show-overflow-tooltip>
            <template #default="scope">
              <span class="font-medium text-blue-600 cursor-pointer hover:underline" @click="getDetails(scope.row)">{{ scope.row.name }}</span>
            </template>
          </el-table-column>


          <el-table-column label="测试进度" width="220">
              <template #default="scope">
                  <div v-if="scope.row.progress" class="stats-container">
                      <el-progress 
                        :percentage="scope.row.progress.total_apis > 0 ? Math.floor((scope.row.progress.executed_apis / scope.row.progress.total_apis) * 100) : 0" 
                        :status="scope.row.status === 2 ? 'exception' : (scope.row.status === 3 ? 'success' : '')"
                      />
                  </div>
                  <div v-else class="text-gray-400">-</div>
              </template>
          </el-table-column>

          <el-table-column label="任务类型" prop="case_type" width="100">
             <template #default="scope">
               <el-tag type="info" effect="plain">{{ scope.row.case_type }}</el-tag>
             </template>
          </el-table-column>
          
          <el-table-column label="运行模式" prop="run_mode" width="100">
             <template #default="scope">
               <el-tag type="warning" effect="plain">{{ scope.row.run_mode }}</el-tag>
             </template>
          </el-table-column>

          <el-table-column label="节点" prop="node_name" width="140">
             <template #default="scope">
               <el-tag v-if="scope.row.node_name" type="info" effect="plain">{{ scope.row.node_name }}</el-tag>
               <span v-else class="text-gray-400">-</span>
             </template>
          </el-table-column>

          <el-table-column label="状态" prop="status" width="100">
              <template #default="scope">
                  <el-tag v-if="scope.row.status===1" effect="dark">运行中</el-tag>
                  <el-tag v-else-if="scope.row.status===2" type="danger" effect="dark">失败</el-tag>
                  <el-tag v-else-if="scope.row.status===3" type="success" effect="dark">成功</el-tag>
                  <el-tag v-else type="info" effect="dark">未知</el-tag>
              </template>
          </el-table-column>


          <el-table-column align="left" label="操作" fixed="right" width="200"  header-align="center" >
              <template #default="scope" >
                <el-button type="primary" link size="small" @click="getDetails(scope.row)">
                  <el-icon class="mr-1" style="padding-left: 10px"><DataAnalysis /></el-icon>详情
                </el-button>
                <el-button type="primary" link size="small" icon="edit" @click="updateAutoReportFunc(scope.row)">编辑</el-button>
                <el-button type="danger" link size="small" icon="delete" @click="deleteRow(scope.row)">删除</el-button>
              </template>
          </el-table-column>
        </el-table>
        <div class="gva-pagination mt-4 flex justify-end">
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
      </el-card>
    </div>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center border-b pb-3">
                <span class="text-lg font-bold text-gray-800">{{type==='create'?'新增报告':'编辑报告'}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px" class="p-4">
            <el-row :gutter="20">
              <el-col :span="12">
                <el-form-item label="名称:" prop="name">
                    <el-input v-model="formData.name" :clearable="false" placeholder="请输入名称" />
                </el-form-item>
              </el-col>
              <el-col :span="12">
                 <el-form-item label="成功:" prop="success">
                    <el-switch v-model="formData.success" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
                </el-form-item>
              </el-col>
            </el-row>
            
            <el-row :gutter="20">
               <el-col :span="12">
                  <el-form-item label="统计ID:" prop="stat_id">
                      <el-input v-model.number="formData.stat_id" :clearable="false" placeholder="请输入统计ID" />
                  </el-form-item>
               </el-col>
               <el-col :span="12">
                  <el-form-item label="时间ID:" prop="time_id">
                      <el-input v-model.number="formData.time_id" :clearable="false" placeholder="请输入时间ID" />
                  </el-form-item>
               </el-col>
            </el-row>

            <el-form-item label="平台:" prop="platform">
               <div class="text-gray-500 text-sm mb-2">JSON结构数据</div>
               <el-input type="textarea" :rows="3" v-model="formData.platform" placeholder="请输入平台JSON数据" disabled />
            </el-form-item>

            <el-row :gutter="20">
               <el-col :span="12">
                  <el-form-item label="状态:" prop="status">
                      <el-input v-model.number="formData.status" :clearable="false" placeholder="请输入状态" />
                  </el-form-item>
               </el-col>
               <el-col :span="12">
                  <el-form-item label="设置案例:" prop="setup_case">
                      <el-switch v-model="formData.setup_case" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
                  </el-form-item>
               </el-col>
            </el-row>

            <el-form-item label="描述:" prop="describe">
                <el-input type="textarea" v-model="formData.describe" :clearable="false" placeholder="请输入描述" />
            </el-form-item>

            <el-row :gutter="20">
               <el-col :span="12">
                  <el-form-item label="API环境名称:" prop="api_env_name">
                      <el-input v-model="formData.api_env_name" :clearable="false" placeholder="请输入API环境名称" />
                  </el-form-item>
               </el-col>
               <el-col :span="12">
                  <el-form-item label="API环境ID:" prop="api_env_id">
                      <el-input v-model.number="formData.api_env_id" :clearable="false" placeholder="请输入API环境ID" />
                  </el-form-item>
               </el-col>
            </el-row>
            
            <el-form-item label="运行节点:" prop="node_name">
                <el-input v-model="formData.node_name" :clearable="false" placeholder="请输入运行节点" disabled />
            </el-form-item>
          </el-form>
    </el-drawer>

    <!-- 详情弹窗 (保留但优化样式) -->
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看详情">
            <el-descriptions :column="1" border class="m-4">
                    <el-descriptions-item label="名称">
                        {{ detailForm.name }}
                    </el-descriptions-item>
                    <el-descriptions-item label="成功">
                        <el-tag :type="detailForm.success ? 'success' : 'danger'">{{ detailForm.success ? '是' : '否' }}</el-tag>
                    </el-descriptions-item>
                    <el-descriptions-item label="统计ID">
                        {{ detailForm.stat_id }}
                    </el-descriptions-item>
                    <el-descriptions-item label="时间ID">
                        {{ detailForm.time_id }}
                    </el-descriptions-item>
                    <el-descriptions-item label="平台">
                        {{ detailForm.platform }}
                    </el-descriptions-item>
                    <el-descriptions-item label="状态">
                         <el-tag v-if="detailForm.status===1">运行中</el-tag>
                         <el-tag v-else-if="detailForm.status===2" type="danger">失败</el-tag>
                         <el-tag v-else-if="detailForm.status===3" type="success">成功</el-tag>
                         <el-tag v-else type="info">未知</el-tag>
                    </el-descriptions-item>
                    <el-descriptions-item label="设置案例">
                        {{ detailForm.setup_case }}
                    </el-descriptions-item>
                    <el-descriptions-item label="描述">
                        {{ detailForm.describe }}
                    </el-descriptions-item>
                    <el-descriptions-item label="API环境名称">
                        {{ detailForm.api_env_name }}
                    </el-descriptions-item>
                    <el-descriptions-item label="API环境ID">
                        {{ detailForm.api_env_id }}
                    </el-descriptions-item>
                    <el-descriptions-item label="运行节点">
                        {{ detailForm.node_name || '-' }}
                    </el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createAutoReport,
  deleteAutoReport,
  deleteAutoReportByIds,
  updateAutoReport,
  findAutoReport,
  getAutoReportList
} from '@/api/automation/autoreport'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { useAppStore } from "@/pinia"
import { QuestionFilled, Search, Refresh, Plus, Delete, Edit, DataAnalysis, Clock, ArrowDown, ArrowUp } from '@element-plus/icons-vue'

defineOptions({
    name: 'AutoReport'
})

// 提交按钮loading
const router = useRouter()
const btnLoading = ref(false)
const appStore = useAppStore()

// 确保 AutoReportDetail 路由已注册（用于查看运行报告详情）

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
            name: '',
            success: false,
            stat_id: 0,
            time_id: 0,
            platform: {},
            status: 0,
            setup_case: false,
            describe: '',
            api_env_name: '',
            api_env_id: 0,
            node_name: '',
        })



// 验证规则
const rule = reactive({
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
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    if (searchInfo.value.success === ""){
        searchInfo.value.success=null
    }
    if (searchInfo.value.setup_case === ""){
        searchInfo.value.setup_case=null
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
const getTableData = async() => {
  const table = await getAutoReportList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
const setOptions = async () =>{
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
            deleteAutoReportFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
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
      const res = await deleteAutoReportByIds({ IDs })
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
const updateAutoReportFunc = async(row) => {
    const res = await findAutoReport({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteAutoReportFunc = async (row) => {
    const res = await deleteAutoReport({ ID: row.ID })
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
        success: false,
        stat_id: 0,
        time_id: 0,
        platform: {},
        status: 0,
        setup_case: false,
        describe: '',
        api_env_name: '',
        api_env_id: 0,
        node_name: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     btnLoading.value = true
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return btnLoading.value = false
              let res
              switch (type.value) {
                case 'create':
                  res = await createAutoReport(formData.value)
                  break
                case 'update':
                  res = await updateAutoReport(formData.value)
                  break
                default:
                  res = await createAutoReport(formData.value)
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
  if (row.status === 2 && row.progress && row.progress.total_cases === 0) {
    const res = await findAutoReport({ ID: row.ID })
    if (res.code === 0) {
      const report = res.data
      if (report.details && report.details.length > 0 && report.details[0].records && report.details[0].records.length > 0) {
        const record = report.details[0].records[0]
         if (['task', 'case', 'step'].includes(record.step_type) && record.attachments) {
           ElMessageBox.alert(record.attachments, '错误信息', {
             confirmButtonText: '确定',
             type: 'error'
           })
           return
         }
      }
    }
  }
  router.push({
    name: 'auto-report-detail',
    params: {
      id: row.ID
    }
  })
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailForm.value = {}
}


</script>

<style scoped>
.auto-report-container {
  padding: 20px;
  background-color: var(--el-bg-color-page);
  min-height: 100vh;
}

.search-card {
  margin-bottom: 20px;
  border-radius: 8px;
}

.table-card {
  border-radius: 8px;
}

.table-header {
  background-color: var(--el-fill-color-light) !important;
  color: var(--el-text-color-secondary);
  font-weight: 600;
}

.stats-container {
  width: 100%;
}

:deep(.el-table .cell) {
  padding: 8px 0;
}
</style>
