
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
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
      
            <el-form-item label="模型名称" prop="name">
  <el-input v-model="searchInfo.name" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="请求模式" prop="requestSchema">
  <el-select v-model="searchInfo.requestSchema" clearable filterable placeholder="请选择" @clear="()=>{searchInfo.requestSchema=undefined}">
    <el-option v-for="(item,key) in RequestSchemaOptions" :key="key" :label="item.label" :value="item.value" />
  </el-select>
</el-form-item>
            
            <el-form-item label="模型标识" prop="model">
  <el-select v-model="searchInfo.model" clearable filterable placeholder="请选择" @clear="()=>{searchInfo.model=undefined}">
    <el-option v-for="(item,key) in modelOptions" :key="key" :label="item.label" :value="item.value" />
  </el-select>
</el-form-item>
            
            <el-form-item label="推理强度" prop="reasoningEffort">
  <el-select v-model="searchInfo.reasoningEffort" clearable filterable placeholder="请选择" @clear="()=>{searchInfo.reasoningEffort=undefined}">
    <el-option v-for="(item,key) in reasoningEffortOptions" :key="key" :label="item.label" :value="item.value" />
  </el-select>
</el-form-item>
            
            <el-form-item label="启用状态" prop="enabled">
  <el-select v-model="searchInfo.enabled" clearable placeholder="请选择">
    <el-option key="true" label="是" value="true"></el-option>
    <el-option key="false" label="否" value="false"></el-option>
  </el-select>
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
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button  type="primary" icon="plus" @click="openDialog()">新增</el-button>
            <el-button  icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            
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
        
            <el-table-column align="left" label="模型名称" prop="name" width="120" />

            <el-table-column align="left" label="请求模式" prop="requestSchema" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.requestSchema,RequestSchemaOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="模型标识" prop="model" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.model,modelOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="推理强度" prop="reasoningEffort" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.reasoningEffort,reasoningEffortOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="最大tokens" prop="maxTokens" width="120" />

            <el-table-column align="left" label="启用状态" prop="enabled" width="120">
    <template #default="scope">{{ formatBoolean(scope.row.enabled) }}</template>
</el-table-column>
            <el-table-column align="left" label="超时时间" prop="timeout" width="120" />

        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateLLMModelConfigFunc(scope.row)">编辑</el-button>
            <el-button   type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
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
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'新增':'编辑'}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="模型名称:" prop="name">
    <el-input v-model="formData.name" :clearable="false" placeholder="请输入模型名称" />
</el-form-item>
            <el-form-item label="请求模式:" prop="requestSchema">
    <el-select v-model="formData.requestSchema" placeholder="请选择请求模式" style="width:100%" filterable :clearable="false">
        <el-option v-for="(item,key) in RequestSchemaOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
            <el-form-item label="模型标识:" prop="model">
    <el-select v-model="formData.model" placeholder="请选择模型标识" style="width:100%" filterable :clearable="false">
        <el-option v-for="(item,key) in modelOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
            <el-form-item label="基础地址:" prop="baseURL">
    <el-input v-model="formData.baseURL" :clearable="false" placeholder="请输入基础地址" />
</el-form-item>
            <el-form-item label="密钥:" prop="apiKey">
    <el-input v-model="formData.apiKey" :clearable="false" placeholder="请输入密钥" />
</el-form-item>
            <el-form-item label="格式化输出:" prop="supportFormatOutput">
    <el-switch v-model="formData.supportFormatOutput" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
            <el-form-item label="推理强度:" prop="reasoningEffort">
    <el-select v-model="formData.reasoningEffort" placeholder="请选择推理强度" style="width:100%" filterable :clearable="false">
        <el-option v-for="(item,key) in reasoningEffortOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
            <el-form-item label="最大tokens:" prop="maxTokens">
    <el-input v-model.number="formData.maxTokens" :clearable="false" placeholder="请输入最大tokens" />
</el-form-item>
            <el-form-item label="随机性参数:" prop="temperature">
    <el-input-number v-model="formData.temperature" style="width:100%" :precision="2" :clearable="false" />
</el-form-item>
            <el-form-item label="核心采样:" prop="topP">
    <el-input-number v-model="formData.topP" style="width:100%" :precision="2" :clearable="false" />
</el-form-item>
            <el-form-item label="启用状态:" prop="enabled">
    <el-switch v-model="formData.enabled" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
            <el-form-item label="超时时间:" prop="timeout">
    <el-input v-model.number="formData.timeout" :clearable="false" placeholder="请输入超时时间" />
</el-form-item>
            <el-form-item label="模型描述:" prop="description">
    <RichEdit v-model="formData.description"/>
</el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="模型名称">
    {{ detailForm.name }}
</el-descriptions-item>
                    <el-descriptions-item label="请求模式">
    {{ detailForm.requestSchema }}
</el-descriptions-item>
                    <el-descriptions-item label="模型标识">
    {{ detailForm.model }}
</el-descriptions-item>
                    <el-descriptions-item label="基础地址">
    {{ detailForm.baseURL }}
</el-descriptions-item>
                    <el-descriptions-item label="格式化输出">
    {{ detailForm.supportFormatOutput }}
</el-descriptions-item>
                    <el-descriptions-item label="推理强度">
    {{ detailForm.reasoningEffort }}
</el-descriptions-item>
                    <el-descriptions-item label="最大tokens">
    {{ detailForm.maxTokens }}
</el-descriptions-item>
                    <el-descriptions-item label="随机性参数">
    {{ detailForm.temperature }}
</el-descriptions-item>
                    <el-descriptions-item label="核心采样">
    {{ detailForm.topP }}
</el-descriptions-item>
                    <el-descriptions-item label="启用状态">
    {{ detailForm.enabled }}
</el-descriptions-item>
                    <el-descriptions-item label="超时时间">
    {{ detailForm.timeout }}
</el-descriptions-item>
                    <el-descriptions-item label="模型描述">
    <RichView v-model="detailForm.description" />
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createLLMModelConfig,
  deleteLLMModelConfig,
  deleteLLMModelConfigByIds,
  updateLLMModelConfig,
  findLLMModelConfig,
  getLLMModelConfigList
} from '@/api/platform/llmModelService'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'
import RichView from '@/components/richtext/rich-view.vue'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useAppStore } from "@/pinia"




defineOptions({
    name: 'LLMModelConfig'
})

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const RequestSchemaOptions = ref([])
const modelOptions = ref([])
const reasoningEffortOptions = ref([])
const formData = ref({
            name: '',
            requestSchema: '',
            model: '',
            baseURL: '',
            apiKey: '',
            supportFormatOutput: false,
            reasoningEffort: '',
            maxTokens: 0,
            temperature: 0,
            topP: 0,
            enabled: false,
            timeout: 0,
            description: '',
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
    if (searchInfo.value.supportFormatOutput === ""){
        searchInfo.value.supportFormatOutput=null
    }
    if (searchInfo.value.enabled === ""){
        searchInfo.value.enabled=null
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
  const table = await getLLMModelConfigList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    RequestSchemaOptions.value = await getDictFunc('RequestSchema')
    modelOptions.value = await getDictFunc('model')
    reasoningEffortOptions.value = await getDictFunc('reasoningEffort')
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
            deleteLLMModelConfigFunc(row)
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
      const res = await deleteLLMModelConfigByIds({ IDs })
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
const updateLLMModelConfigFunc = async(row) => {
    const res = await findLLMModelConfig({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteLLMModelConfigFunc = async (row) => {
    const res = await deleteLLMModelConfig({ ID: row.ID })
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
        requestSchema: '',
        model: '',
        baseURL: '',
        apiKey: '',
        supportFormatOutput: false,
        reasoningEffort: '',
        maxTokens: 0,
        temperature: 0,
        topP: 0,
        enabled: false,
        timeout: 0,
        description: '',
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
                  res = await createLLMModelConfig(formData.value)
                  break
                case 'update':
                  res = await updateLLMModelConfig(formData.value)
                  break
                default:
                  res = await createLLMModelConfig(formData.value)
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
  const res = await findLLMModelConfig({ ID: row.ID })
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


</script>

<style>

</style>
