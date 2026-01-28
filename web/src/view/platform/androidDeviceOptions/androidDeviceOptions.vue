
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
      
            <el-form-item label="设备名称" prop="name">
  <el-input v-model="searchInfo.name" placeholder="搜索条件" />
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
        @sort-change="sortChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column sortable align="left" label="日期" prop="CreatedAt" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
            <el-table-column align="left" label="设备名称" prop="name" width="120" />

            <el-table-column sortable align="left" label="序列号" prop="serial" width="120" />

            <el-table-column align="left" label="日志开关" prop="logOn" width="120">
    <template #default="scope">{{ formatBoolean(scope.row.logOn) }}</template>
</el-table-column>
            <el-table-column align="left" label="忽略弹窗" prop="ignorePopup" width="120">
    <template #default="scope">{{ formatBoolean(scope.row.ignorePopup) }}</template>
</el-table-column>
            <el-table-column align="left" label="ADB主机" prop="adbServerHost" width="120" />

            <el-table-column align="left" label="ADB端口" prop="adbServerPort" width="120" />

            <el-table-column align="left" label="复合驱动" prop="composite" width="120">
    <template #default="scope">{{ formatBoolean(scope.row.composite) }}</template>
</el-table-column>
            <el-table-column align="left" label="UIA2开关" prop="uia2" width="120">
    <template #default="scope">{{ formatBoolean(scope.row.uia2) }}</template>
</el-table-column>
        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateAndroidDeviceOptionsFunc(scope.row)">编辑</el-button>
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
            <el-form-item label="设备名称:" prop="name">
    <el-input v-model="formData.name" :clearable="true" placeholder="请输入设备名称" />
</el-form-item>
            <el-form-item label="序列号:" prop="serial">
    <el-input v-model="formData.serial" :clearable="false" placeholder="请输入序列号" />
</el-form-item>
            <el-form-item label="日志开关:" prop="logOn">
    <el-switch v-model="formData.logOn" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
            <el-form-item label="忽略弹窗:" prop="ignorePopup">
    <el-switch v-model="formData.ignorePopup" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
            <el-form-item label="ADB主机:" prop="adbServerHost">
    <el-input v-model="formData.adbServerHost" :clearable="false" placeholder="请输入ADB主机" />
</el-form-item>
            <el-form-item label="ADB端口:" prop="adbServerPort">
    <el-input v-model.number="formData.adbServerPort" :clearable="false" placeholder="请输入ADB端口" />
</el-form-item>
            <el-form-item label="复合驱动:" prop="composite">
    <el-switch v-model="formData.composite" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
            <el-form-item label="UIA2开关:" prop="uia2">
    <el-switch v-model="formData.uia2" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
            <el-form-item label="UIA2地址:" prop="uia2Ip">
    <el-input v-model="formData.uia2Ip" :clearable="false" placeholder="请输入UIA2地址" />
</el-form-item>
            <el-form-item label="UIA2端口:" prop="uia2Port">
    <el-input v-model.number="formData.uia2Port" :clearable="false" placeholder="请输入UIA2端口" />
</el-form-item>
            <el-form-item label="UIA2服务包名:" prop="uia2ServerPackageName">
    <el-input v-model="formData.uia2ServerPackageName" :clearable="false" placeholder="请输入UIA2服务包名" />
</el-form-item>
            <el-form-item label="UIA2测试包名:" prop="uia2ServerTestPackageName">
    <el-input v-model="formData.uia2ServerTestPackageName" :clearable="false" placeholder="请输入UIA2测试包名" />
</el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="设备名称">
    {{ detailForm.name }}
</el-descriptions-item>
                    <el-descriptions-item label="序列号">
    {{ detailForm.serial }}
</el-descriptions-item>
                    <el-descriptions-item label="日志开关">
    {{ detailForm.logOn }}
</el-descriptions-item>
                    <el-descriptions-item label="忽略弹窗">
    {{ detailForm.ignorePopup }}
</el-descriptions-item>
                    <el-descriptions-item label="ADB主机">
    {{ detailForm.adbServerHost }}
</el-descriptions-item>
                    <el-descriptions-item label="ADB端口">
    {{ detailForm.adbServerPort }}
</el-descriptions-item>
                    <el-descriptions-item label="复合驱动">
    {{ detailForm.composite }}
</el-descriptions-item>
                    <el-descriptions-item label="UIA2开关">
    {{ detailForm.uia2 }}
</el-descriptions-item>
                    <el-descriptions-item label="UIA2地址">
    {{ detailForm.uia2Ip }}
</el-descriptions-item>
                    <el-descriptions-item label="UIA2端口">
    {{ detailForm.uia2Port }}
</el-descriptions-item>
                    <el-descriptions-item label="UIA2服务包名">
    {{ detailForm.uia2ServerPackageName }}
</el-descriptions-item>
                    <el-descriptions-item label="UIA2测试包名">
    {{ detailForm.uia2ServerTestPackageName }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createAndroidDeviceOptions,
  deleteAndroidDeviceOptions,
  deleteAndroidDeviceOptionsByIds,
  updateAndroidDeviceOptions,
  findAndroidDeviceOptions,
  getAndroidDeviceOptionsList
} from '@/api/platform/androidDeviceOptions'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useAppStore } from "@/pinia"




defineOptions({
    name: 'AndroidDeviceOptions'
})

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
            name: '',
            serial: '',
            logOn: false,
            ignorePopup: false,
            adbServerHost: '',
            adbServerPort: 0,
            composite: false,
            uia2: false,
            uia2Ip: '',
            uia2Port: 0,
            uia2ServerPackageName: '',
            uia2ServerTestPackageName: '',
        })



// 验证规则
const rule = reactive({
               serial : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
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
    CreatedAt:"created_at",
    ID:"id",
            serial: 'serial_number',
  }

  let sort = sortMap[prop]
  if(!sort){
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
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    if (searchInfo.value.logOn === ""){
        searchInfo.value.logOn=null
    }
    if (searchInfo.value.ignorePopup === ""){
        searchInfo.value.ignorePopup=null
    }
    if (searchInfo.value.composite === ""){
        searchInfo.value.composite=null
    }
    if (searchInfo.value.uia2 === ""){
        searchInfo.value.uia2=null
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
  const table = await getAndroidDeviceOptionsList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteAndroidDeviceOptionsFunc(row)
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
      const res = await deleteAndroidDeviceOptionsByIds({ IDs })
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
const updateAndroidDeviceOptionsFunc = async(row) => {
    const res = await findAndroidDeviceOptions({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteAndroidDeviceOptionsFunc = async (row) => {
    const res = await deleteAndroidDeviceOptions({ ID: row.ID })
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
        serial: '',
        logOn: false,
        ignorePopup: false,
        adbServerHost: '',
        adbServerPort: 0,
        composite: false,
        uia2: false,
        uia2Ip: '',
        uia2Port: 0,
        uia2ServerPackageName: '',
        uia2ServerTestPackageName: '',
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
                  res = await createAndroidDeviceOptions(formData.value)
                  break
                case 'update':
                  res = await updateAndroidDeviceOptions(formData.value)
                  break
                default:
                  res = await createAndroidDeviceOptions(formData.value)
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
  const res = await findAndroidDeviceOptions({ ID: row.ID })
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
