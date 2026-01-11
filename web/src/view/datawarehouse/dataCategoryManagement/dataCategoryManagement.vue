<template>
  <div>
    <!-- 搜索区域 -->
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" @keyup.enter="onSubmit">
        <el-form-item label="创建日期" prop="createdAtRange">
          <el-date-picker
            v-model="searchInfo.createdAtRange"
            type="datetimerange"
            range-separator="至"
            start-placeholder="开始时间"
            end-placeholder="结束时间"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>

    <!-- 表格区域 -->
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog()">新增</el-button>
        <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
      </div>

      <el-table :data="tableData" row-key="ID" @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" />
        <el-table-column label="日期" prop="CreatedAt" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column label="数据分类名称" prop="name" width="150" />
        <el-table-column label="数据类型" prop="type" width="120" />
        <!-- 各环境总数量展示 -->
        <el-table-column label="各环境总数" min-width="200">
          <template #default="scope">
            <div class="env-counts">
              <span v-for="env in envList" :key="env.ID" class="env-count-chip">
                <span class="env-name">{{ env.name }}:</span>
                <span class="env-val">{{ getEnvNumber(scope.row.count, env.ID) }}</span>
              </span>
            </div>
          </template>
        </el-table-column>
        <!-- 各环境可用数展示 -->
        <el-table-column label="各环境可用数" min-width="200">
          <template #default="scope">
            <div class="env-counts">
              <span v-for="env in envList" :key="env.ID" class="env-count-chip">
                <span class="env-name">{{ env.name }}:</span>
                <span class="env-val">{{ getEnvNumber(scope.row.availableCount, env.ID) }}</span>
              </span>
            </div>
          </template>
        </el-table-column>
        <el-table-column label="创建数据方式" width="120">
          <template #default="scope">
            {{ scope.row.createCallType === 1 ? '测试步骤' : 'Python脚本' }}
          </template>
        </el-table-column>
        <el-table-column label="清洗数据方式" width="120">
          <template #default="scope">
            {{ formatCleanCallType(scope.row.cleanCallType) }}
          </template>
        </el-table-column>
        <el-table-column label="操作" fixed="right" width="250">
          <template #default="scope">
            <el-button type="primary" link @click="openCodeEditor(scope.row)">编辑代码</el-button>
            <el-button type="primary" link @click="getDetails(scope.row)">查看</el-button>
            <el-button type="primary" link @click="updateFunc(scope.row)">编辑</el-button>
            <el-button type="danger" link @click="deleteRow(scope.row)">删除</el-button>
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

    <!-- 新增/编辑抽屉 -->
    <el-drawer v-model="dialogFormVisible" :size="800" :show-close="false" destroy-on-close>
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ type === 'create' ? '新增数据分类' : '编辑数据分类' }}</span>
          <div>
            <el-button :loading="btnLoading" type="primary" @click="enterDialog">保存</el-button>
            <el-button @click="closeDialog">取消</el-button>
          </div>
        </div>
      </template>

      <el-form :model="formData" ref="elFormRef" :rules="rules" label-position="top">
        <!-- 基础信息 -->
        <el-divider content-position="left">基础信息</el-divider>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="数据分类名称" prop="name">
              <el-input v-model="formData.name" placeholder="请输入数据分类名称" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="数据类型" prop="type">
              <el-input v-model="formData.type" placeholder="请输入英文名称" :disabled="type === 'update'" />
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 创建数据配置 -->
        <el-divider content-position="left">创建数据配置</el-divider>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="调用类型" prop="createCallType">
              <el-radio-group v-model="formData.createCallType">
                <el-radio :label="1">测试步骤</el-radio>
                <el-radio :label="2">Python脚本</el-radio>
              </el-radio-group>
            </el-form-item>
          </el-col>
          <el-col :span="12" v-if="formData.createCallType === 1">
            <el-form-item label="运行配置" prop="createRunConfigId">
              <RunConfig v-model="formData.createRunConfigId" width="100%" />
            </el-form-item>
          </el-col>
          <el-col :span="12" v-if="formData.createCallType === 1">
            <el-form-item label="测试步骤" prop="createTestStepId">
              <AutoStepSelect v-model="formData.createTestStepId" width="100%" />
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 数量配置（各环境） -->
        <el-divider content-position="left">数量配置（各环境）</el-divider>
        <el-table :data="envList" style="width: 100%; margin-bottom: 16px;" size="small" v-if="envList.length">
          <el-table-column label="环境" width="200">
            <template #default="scope">{{ scope.row.name }}</template>
          </el-table-column>
          <el-table-column label="总数量">
            <template #default="scope">
              <el-input-number v-model="formData.count[scope.row.ID]" :min="0" :step="1" :controls="false" style="width: 160px" />
            </template>
          </el-table-column>
        </el-table>

        <!-- 清洗数据配置 -->
        <el-divider content-position="left">清洗数据配置</el-divider>
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="调用类型" prop="cleanCallType">
              <el-radio-group v-model="formData.cleanCallType">
                <el-radio :label="1">测试步骤</el-radio>
                <el-radio :label="2">Python脚本</el-radio>
                <el-radio :label="3">直接删除</el-radio>
              </el-radio-group>
            </el-form-item>
          </el-col>
          <el-col :span="12" v-if="formData.cleanCallType === 1">
             <el-form-item label="运行配置" prop="cleanRunConfigId">
              <RunConfig v-model="formData.cleanRunConfigId" width="100%" />
            </el-form-item>
          </el-col>
          <el-col :span="12" v-if="formData.cleanCallType === 1">
            <el-form-item label="测试步骤" prop="cleanTestStepId">
              <AutoStepSelect v-model="formData.cleanTestStepId" width="100%" />
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
    </el-drawer>

    <!-- 查看详情抽屉 -->
    <el-drawer v-model="detailShow" :size="600" title="查看详情" destroy-on-close>
      <el-descriptions :column="1" border>
        <el-descriptions-item label="数据分类名称">{{ detailForm.name }}</el-descriptions-item>
        <el-descriptions-item label="数据类型">{{ detailForm.type }}</el-descriptions-item>
        <el-descriptions-item label="创建数据方式">
          {{ detailForm.createCallType === 1 ? '测试步骤' : 'Python脚本' }}
        </el-descriptions-item>
        <el-descriptions-item label="创建数据测试步骤ID" v-if="detailForm.createCallType === 1">
          {{ detailForm.createTestStepId }}
        </el-descriptions-item>
        <el-descriptions-item label="清洗数据方式">
          {{ formatCleanCallType(detailForm.cleanCallType) }}
        </el-descriptions-item>
        <el-descriptions-item label="清洗数据测试步骤ID" v-if="detailForm.cleanCallType === 1">
          {{ detailForm.cleanTestStepId }}
        </el-descriptions-item>
      </el-descriptions>

      <!-- 数量查看 -->
      <div style="margin-top: 20px;">
        <el-divider content-position="left">数量（各环境）</el-divider>
        <div class="env-counts">
          <span v-for="env in detailForm.envList || envList" :key="env.ID" class="env-count-chip">
            <span class="env-name">{{ env.name }}:</span>
            <span class="env-val">{{ getEnvNumber(detailForm.count, env.ID) }}</span>
          </span>
        </div>
      </div>
      <!-- 数量查看 -->
      <div style="margin-top: 20px;">
        <el-divider content-position="left">可用数量（各环境）</el-divider>
        <div class="env-counts">
          <span v-for="env in detailForm.envList || envList" :key="env.ID" class="env-count-chip">
            <span class="env-name">{{ env.name }}:</span>
            <span class="env-val">{{ getEnvNumber(detailForm.availableCount, env.ID) }}</span>
          </span>
        </div>
      </div>

      <!-- 代码查看 -->
      <div v-if="detailForm.pythonCodes && Object.keys(detailForm.pythonCodes).length" style="margin-top: 20px;">
        <el-divider content-position="left">Python 代码</el-divider>
        <el-tabs type="border-card">
          <el-tab-pane
            v-for="env in detailForm.envList"
            :key="env.ID"
            :label="env.name"
          >
            <pre class="code-preview">
{{ formatPythonCode(detailForm.pythonCodes[env.ID]) }}
            </pre>
          </el-tab-pane>
        </el-tabs>
      </div>
    </el-drawer>
    <!-- 代码编辑弹窗 -->
    <el-dialog
      v-model="codeEditorVisible"
      title="编辑 Python 代码"
      width="90%"
      top="2vh"
      :close-on-click-modal="false"
      :close-on-press-escape="false"
      destroy-on-close
    >
      <el-alert
        v-if="!envList.length"
        type="warning"
        :closable="false"
        description="当前项目暂无环境配置，请先在环境管理中添加环境"
        style="margin-bottom: 15px;"
      />
      <el-tabs v-else v-model="activeCodeEnvTab" type="border-card">
        <el-tab-pane
          v-for="env in envList"
          :key="env.ID"
          :label="env.name"
          :name="String(env.ID)"
        >
          <PythonCodeEditor
            v-model="codeForm.pythonCodes[env.ID]"
            height="600px"
            :title="`${env.name} 环境代码`"
            :unique-key="`dcm_${codeForm.ID}_${env.ID}`"
            :show-history-button="true"
            @save="(code) => saveSingleCode(env.ID, code)"
            historyType=2
          >
            <template #header-actions>
              <el-button type="primary" size="small" icon="Document" @click="saveSingleCode(env.ID, codeForm.pythonCodes[env.ID])" :loading="codeBtnLoading">
                保存
              </el-button>
            </template>
          </PythonCodeEditor>
        </el-tab-pane>
      </el-tabs>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeCodeEditor">关 闭</el-button>
        </div>
      </template>
    </el-dialog>

  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { formatDate, formatBoolean } from '@/utils/format'
import { getEnvList } from '@/api/platform/env'
import {
  createDataCategoryManagement,
  updateDataCategoryManagement,
  findDataCategoryManagement,
  getDataCategoryManagementList,
  deleteDataCategoryManagement,
  deleteDataCategoryManagementByIds
} from '@/api/datawarehouse/dataCategoryManagement'

import RunConfig from '@/components/platform/runConfig.vue'
import AutoStepSelect from '@/components/automation/AutoStepSelect.vue'
import PythonCodeEditor from '@/components/platform/pythonEditor/PythonCodeEditor.vue'

defineOptions({ name: 'DataCategoryManagement' })

// ========== 状态 ==========
const envList = ref([])
const activeEnvTab = ref('')
const dialogFormVisible = ref(false)
const btnLoading = ref(false)
const type = ref('create')

const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const tableData = ref([])
const searchInfo = ref({})

const elFormRef = ref()
const elSearchFormRef = ref()

// 表单数据
const formData = reactive({
  ID: 0,
  name: '',
  type: '',
  // 各环境数量
  count: {},
  createCallType: 1,
  createTestStepId: null,
  cleanCallType: 1,
  cleanTestStepId: null,
  directDelete: false,
  createRunConfigId: null,
  cleanRunConfigId: null,
  pythonCodes: {}
})

// 表单校验规则
const rules = {
  name: [{ required: true, message: '请输入数据分类名称', trigger: 'blur' }],
  type: [{ required: true, message: '请输入英文名称', trigger: 'blur' }]
}

// ========== 方法 ==========

// 获取环境列表
const fetchEnvList = async () => {
  const res = await getEnvList({ page: 1, pageSize: 1000 })
  if (res.code === 0) {
    envList.value = res.data.list || []
    // 初始化每个环境的代码/数量
    envList.value.forEach(env => {
      if (!(env.ID in formData.pythonCodes)) {
        formData.pythonCodes[env.ID] = ''
      }
      if (!(env.ID in formData.count)) {
        formData.count[env.ID] = 0
      }
    })
    // 默认选中第一个环境
    if (envList.value.length > 0 && !activeEnvTab.value) {
      activeEnvTab.value = String(envList.value[0].ID)
    }
  }
}

// 获取表格数据
const getTableData = async () => {
  const res = await getDataCategoryManagementList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value
  })
  if (res.code === 0) {
    tableData.value = res.data.list
    total.value = res.data.total
  }
}

// 搜索
const onSubmit = () => {
  page.value = 1
  getTableData()
}

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 多选数据
const multipleSelection = ref([])
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 重置表单
const resetForm = () => {
  formData.ID = 0
  formData.name = ''
  formData.type = ''
  formData.createCallType = 1
  formData.createTestStepId = null
  formData.cleanCallType = 1
  formData.cleanTestStepId = null
  formData.directDelete = false
  formData.createRunConfigId = null
  formData.cleanRunConfigId = null
  formData.pythonCodes = {}
  formData.count = {}
  // 重新初始化环境代码/数量
  envList.value.forEach(env => {
    formData.pythonCodes[env.ID] = ''
    formData.count[env.ID] = 0
  })
  if (envList.value.length > 0) {
    activeEnvTab.value = String(envList.value[0].ID)
  }
}

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  resetForm()
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  resetForm()
}

// 编辑
const updateFunc = async (row) => {
  const res = await findDataCategoryManagement({ ID: row.ID })
    if (res.code === 0) {
      type.value = 'update'
      const data = res.data

      formData.ID = data.ID
      formData.name = data.name || ''
      formData.type = data.type || ''
      formData.createCallType = data.createCallType || 1
      formData.createTestStepId = data.createTestStepId || null
      formData.cleanCallType = data.cleanCallType || 1
      formData.cleanTestStepId = data.cleanTestStepId || null
      formData.directDelete = data.directDelete || false
      formData.createRunConfigId = data.createRunConfigId || null
      formData.cleanRunConfigId = data.cleanRunConfigId || null

      // 初始化 count
      formData.count = {}
      envList.value.forEach(env => {
        formData.count[env.ID] = 0
      })

      // pythonCodes: key 需要转为数字，值兼容字符串和对象结构
      formData.pythonCodes = {}
      envList.value.forEach(env => {
        formData.pythonCodes[env.ID] = ''
      })
      // 填充 count: 兼容后端 JSONMap 类型（可能值为 number 或字符串）
      if (data.count) {
        for (const [key, value] of Object.entries(data.count)) {
          const num = Number(value || 0)
          formData.count[Number(key)] = Number.isFinite(num) ? num : 0
        }
      }

      if (data.pythonCodes) {
        for (const [key, value] of Object.entries(data.pythonCodes)) {
          if (typeof value === 'string') {
            formData.pythonCodes[Number(key)] = value
          } else if (value && typeof value === 'object') {
            formData.pythonCodes[Number(key)] = value.code || ''
          }
        }
      }

      // 设置默认 Tab
      if (envList.value.length > 0) {
        activeEnvTab.value = String(envList.value[0].ID)
      }

      dialogFormVisible.value = true
    }
  }

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除该数据分类吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const res = await deleteDataCategoryManagement({ ID: row.ID })
    if (res.code === 0) {
      ElMessage.success('删除成功')
      getTableData()
    }
  })
}

// 多选删除
const onDelete = () => {
  if (multipleSelection.value.length === 0) {
    ElMessage.warning('请选择要删除的数据')
    return
  }
  ElMessageBox.confirm('确定要删除选中的数据吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const IDs = multipleSelection.value.map(item => item.ID)
    const res = await deleteDataCategoryManagementByIds({ IDs })
    if (res.code === 0) {
      ElMessage.success('删除成功')
      getTableData()
    }
  })
}

// 保存
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return

    btnLoading.value = true

    // 构建请求数据，pythonCodes 的 key 转为字符串
    const submitData = {
      ID: formData.ID,
      name: formData.name,
      type: formData.type,
      // 各环境数量
      count: {},
      createCallType: formData.createCallType,
      createTestStepId: formData.createTestStepId,
      cleanCallType: formData.cleanCallType,
      cleanTestStepId: formData.cleanTestStepId,
      directDelete: formData.directDelete,
      createRunConfigId: formData.createRunConfigId,
      cleanRunConfigId: formData.cleanRunConfigId,
      pythonCodes: {}
    }
    // count: key 为环境 ID，值为 number
    for (const [key, value] of Object.entries(formData.count)) {
      submitData.count[String(key)] = Number(value || 0)
    }
    // pythonCodes: key 为环境 ID，值为包含 code / update_at 的对象，这里仅提交 code，update_at 由后端维护
    for (const [key, value] of Object.entries(formData.pythonCodes)) {
      submitData.pythonCodes[String(key)] = { code: value }
    }

    let res
    if (type.value === 'create') {
      res = await createDataCategoryManagement(submitData)
    } else {
      res = await updateDataCategoryManagement(submitData)
    }

    btnLoading.value = false

    if (res.code === 0) {
      ElMessage.success(type.value === 'create' ? '创建成功' : '更新成功')
      closeDialog()
      getTableData()
    }
  })
}

// 代码编辑相关
const codeEditorVisible = ref(false)
const codeBtnLoading = ref(false)
const activeCodeEnvTab = ref('')
const codeForm = reactive({
  ID: 0,
  pythonCodes: {}
})

// 打开代码编辑弹窗
const openCodeEditor = async (row) => {
  const res = await findDataCategoryManagement({ ID: row.ID })
  if (res.code === 0) {
    const data = res.data
    codeForm.ID = data.ID
    codeForm.pythonCodes = {}
    
    envList.value.forEach(env => {
      codeForm.pythonCodes[env.ID] = ''
    })
    
    if (data.pythonCodes) {
      for (const [key, value] of Object.entries(data.pythonCodes)) {
        if (typeof value === 'string') {
          codeForm.pythonCodes[Number(key)] = value
        } else if (value && typeof value === 'object') {
          codeForm.pythonCodes[Number(key)] = value.code || ''
        }
      }
    }

    if (envList.value.length > 0) {
      activeCodeEnvTab.value = String(envList.value[0].ID)
    }
    
    codeEditorVisible.value = true
  }
}

// 关闭代码编辑弹窗
const closeCodeEditor = () => {
  codeEditorVisible.value = false
  codeForm.ID = 0
  codeForm.pythonCodes = {}
}

// 保存单个环境代码
const saveSingleCode = async (envId, code) => {
  codeBtnLoading.value = true
  try {
    const resDetail = await findDataCategoryManagement({ ID: codeForm.ID })
    if (resDetail.code !== 0) {
      ElMessage.error('获取数据失败，无法保存')
      codeBtnLoading.value = false
      return
    }
    
    const currentData = resDetail.data
    const submitData = {
      ID: currentData.ID,
      name: currentData.name,
      createCallType: currentData.createCallType,
      createTestStepId: currentData.createTestStepId,
      cleanCallType: currentData.cleanCallType,
      cleanTestStepId: currentData.cleanTestStepId,
      directDelete: currentData.directDelete,
      createRunConfigId: currentData.createRunConfigId,
      cleanRunConfigId: currentData.cleanRunConfigId,
      pythonCodes: {}
    }
    
    // 只提交当前环境的代码
    submitData.pythonCodes[String(envId)] = { code }
    
    const res = await updateDataCategoryManagement(submitData)
    if (res.code === 0) {
      ElMessage.success('代码保存成功')
      // 不需要关闭弹窗
      // 也不需要刷新整个表格，但需要刷新代码状态吗？
      // 前端状态已经是新的了，无需刷新
    }
  } catch (e) {
    console.error(e)
    ElMessage.error('保存失败')
  } finally {
    codeBtnLoading.value = false
  }
}

// 查看详情
const detailForm = ref({})

// 兼容字符串 / 对象两种结构的 pythonCodes 展示
const formatPythonCode = (val) => {
  if (!val) return '未配置'
  if (typeof val === 'string') return val || '未配置'
  if (typeof val === 'object') return val.code || '未配置'
  return '未配置'
}

// 获取某环境下的数字（不存在返回0）
const getEnvNumber = (obj, envId) => {
  if (!obj) return 0
  const val = obj[String(envId)] ?? obj[envId]
  const num = Number(val || 0)
  return Number.isFinite(num) ? num : 0
}

// 格式化清洗数据方式
const formatCleanCallType = (type) => {
  switch (type) {
    case 1: return '测试步骤'
    case 2: return 'Python脚本'
    case 3: return '直接删除'
    default: return '未配置'
  }
}

const detailShow = ref(false)

const getDetails = async (row) => {
  const res = await findDataCategoryManagement({ ID: row.ID })
  if (res.code === 0) {
    detailForm.value = res.data
    detailShow.value = true
  }
}

// ========== 初始化 ==========
onMounted(async () => {
  await fetchEnvList()
  getTableData()
})
</script>

<style scoped>
.code-textarea :deep(.el-textarea__inner) {
  font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', 'Consolas', monospace;
  font-size: 13px;
  line-height: 1.5;
}
.code-preview {
  background-color: #f5f7fa;
  padding: 15px;
  border-radius: 4px;
  font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', 'Consolas', monospace;
  font-size: 13px;
  line-height: 1.5;
  white-space: pre-wrap;
  word-wrap: break-word;
  max-height: 400px;
  overflow-y: auto;
}
</style>
<style scoped>
.env-counts { display: flex; flex-wrap: wrap; gap: 8px 12px; }
.env-count-chip { background: #f5f7fa; border-radius: 4px; padding: 2px 8px; }
.env-name { color: #606266; margin-right: 4px; }
.env-val { font-weight: 600; }
</style>
