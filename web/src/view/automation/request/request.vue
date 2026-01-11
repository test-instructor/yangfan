
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
        
            <el-table-column align="left" label="方法" prop="method" width="120" />

            <el-table-column align="left" label="URL" prop="url" width="120" />

            <el-table-column align="left" label="HTTP2" prop="http2" width="120">
    <template #default="scope">{{ formatBoolean(scope.row.http2) }}</template>
</el-table-column>
            <el-table-column align="left" label="参数" prop="params" width="120" />

            <el-table-column align="left" label="头" prop="headers" width="120" />

            <el-table-column align="left" label="Cookies" prop="cookies" width="120" />

            <el-table-column align="left" label="Body" prop="body" width="120" />

            <el-table-column align="left" label="Json" prop="json" width="120" />

            <el-table-column align="left" label="数据" prop="data" width="120" />

            <el-table-column align="left" label="超时" prop="timeout" width="120" />

            <el-table-column align="left" label="允许重定向" prop="allow_redirects" width="120">
    <template #default="scope">{{ formatBoolean(scope.row.allow_redirects) }}</template>
</el-table-column>
            <el-table-column align="left" label="验证" prop="verify" width="120">
    <template #default="scope">{{ formatBoolean(scope.row.verify) }}</template>
</el-table-column>
            <el-table-column align="left" label="上传" prop="upload" width="120" />

        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateRequestFunc(scope.row)">编辑</el-button>
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
            <el-form-item label="方法:" prop="method">
    <el-input v-model="formData.method" :clearable="false" placeholder="请输入方法" />
</el-form-item>
            <el-form-item label="URL:" prop="url">
    <el-input v-model="formData.url" :clearable="false" placeholder="请输入URL" />
</el-form-item>
            <el-form-item label="HTTP2:" prop="http2">
    <el-switch v-model="formData.http2" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
            <el-form-item label="参数:" prop="params">
</el-form-item>
            <el-form-item label="头:" prop="headers">
</el-form-item>
            <el-form-item label="Cookies:" prop="cookies">
</el-form-item>
            <el-form-item label="Body:" prop="body">
</el-form-item>
            <el-form-item label="Json:" prop="json">
</el-form-item>
            <el-form-item label="数据:" prop="data">
</el-form-item>
            <el-form-item label="超时:" prop="timeout">
    <el-input-number v-model="formData.timeout" style="width:100%" :precision="2" :clearable="false" />
</el-form-item>
            <el-form-item label="允许重定向:" prop="allow_redirects">
    <el-switch v-model="formData.allow_redirects" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
            <el-form-item label="验证:" prop="verify">
    <el-switch v-model="formData.verify" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
            <el-form-item label="上传:" prop="upload">
</el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="方法">
    {{ detailForm.method }}
</el-descriptions-item>
                    <el-descriptions-item label="URL">
    {{ detailForm.url }}
</el-descriptions-item>
                    <el-descriptions-item label="HTTP2">
    {{ detailForm.http2 }}
</el-descriptions-item>
                    <el-descriptions-item label="参数">
    {{ detailForm.params }}
</el-descriptions-item>
                    <el-descriptions-item label="头">
    {{ detailForm.headers }}
</el-descriptions-item>
                    <el-descriptions-item label="Cookies">
    {{ detailForm.cookies }}
</el-descriptions-item>
                    <el-descriptions-item label="Body">
    {{ detailForm.body }}
</el-descriptions-item>
                    <el-descriptions-item label="Json">
    {{ detailForm.json }}
</el-descriptions-item>
                    <el-descriptions-item label="数据">
    {{ detailForm.data }}
</el-descriptions-item>
                    <el-descriptions-item label="超时">
    {{ detailForm.timeout }}
</el-descriptions-item>
                    <el-descriptions-item label="允许重定向">
    {{ detailForm.allow_redirects }}
</el-descriptions-item>
                    <el-descriptions-item label="验证">
    {{ detailForm.verify }}
</el-descriptions-item>
                    <el-descriptions-item label="上传">
    {{ detailForm.upload }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createRequest,
  deleteRequest,
  deleteRequestByIds,
  updateRequest,
  findRequest,
  getRequestList
} from '@/api/automation/request'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useAppStore } from "@/pinia"




defineOptions({
    name: 'Request'
})

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
            method: '',
            url: '',
            http2: false,
            params: null,
            headers: null,
            cookies: null,
            body: null,
            json: null,
            data: null,
            timeout: 0,
            allow_redirects: false,
            verify: false,
            upload: null,
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
    if (searchInfo.value.http2 === ""){
        searchInfo.value.http2=null
    }
    if (searchInfo.value.allow_redirects === ""){
        searchInfo.value.allow_redirects=null
    }
    if (searchInfo.value.verify === ""){
        searchInfo.value.verify=null
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
  const table = await getRequestList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteRequestFunc(row)
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
      const res = await deleteRequestByIds({ IDs })
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
const updateRequestFunc = async(row) => {
    const res = await findRequest({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteRequestFunc = async (row) => {
    const res = await deleteRequest({ ID: row.ID })
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
        method: '',
        url: '',
        http2: false,
        params: null,
        headers: null,
        cookies: null,
        body: null,
        json: null,
        data: null,
        timeout: 0,
        allow_redirects: false,
        verify: false,
        upload: null,
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
                  res = await createRequest(formData.value)
                  break
                case 'update':
                  res = await updateRequest(formData.value)
                  break
                default:
                  res = await createRequest(formData.value)
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
  const res = await findRequest({ ID: row.ID })
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
