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
        <el-button  icon="setting" style="margin-left: 10px;" @click="openEnvManager">环境管理</el-button>
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

<!--        <el-table-column sortable align="left" label="日期" prop="CreatedAt" width="180">-->
<!--          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>-->
<!--        </el-table-column>-->

        <el-table-column align="left" label="变量名称" prop="key" width="240" fixed="left" />
        <el-table-column align="left" label="中文名" prop="name" width="120" fixed="left" />

        <!-- 动态生成环境列 -->
        <el-table-column
          v-for="env in envList"
          :key="env.ID"
          :label="env.name"
          align="left"
          width="180"
        >
          <template #default="scope">
            {{ scope.row.value[env.ID] }}
          </template>
        </el-table-column>

        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
          <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateEnvDetailFunc(scope.row)">编辑</el-button>
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
        <el-form-item label="变量名称:" prop="key">
          <el-input
            v-model="formData.key"
            :clearable="false"
            placeholder="请输入变量名称"
            :disabled="type === 'update'"
          >
            <template #prepend v-if="type === 'create'">ENV_</template>
          </el-input>
        </el-form-item>
        <el-form-item label="中文名:" prop="name">
          <el-input v-model="formData.name" :clearable="false" placeholder="请输入中文名" />
        </el-form-item>

        <!-- 动态生成环境输入框 -->
        <el-divider />
        <el-form-item
          v-for="env in envList"
          :key="env.ID"
          :label="env.name + ':'"
          :prop="`value.${env.ID}`"
        >
          <el-input
            v-model="formData.value[env.ID]"
            :clearable="false"
            :placeholder="`请输入${env.name}的值`"
          />
        </el-form-item>
      </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
      <el-descriptions :column="1" border>
        <el-descriptions-item label="变量名称">
          {{ detailForm.key }}
        </el-descriptions-item>
        <el-descriptions-item label="中文名">
          {{ detailForm.name }}
        </el-descriptions-item>

        <!-- 动态显示环境值 -->
        <el-descriptions-item
          v-for="env in envList"
          :key="env.ID"
          :label="env.name"
        >
          {{ detailForm.value && detailForm.value[env.ID] ? detailForm.value[env.ID] : '未设置' }}
        </el-descriptions-item>
      </el-descriptions>
    </el-drawer>

    <!-- 环境管理组件 -->
    <EnvManager v-model="envManagerVisible" @changed="onEnvChanged" />
  </div>
</template>

<script setup>
  import {
    createEnvDetail,
    deleteEnvDetail,
    deleteEnvDetailByIds,
    updateEnvDetail,
    findEnvDetail,
    getEnvDetailList
  } from '@/api/platform/envdetail'

  // 全量引入格式化工具 请按需保留
  import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { ref, reactive, onMounted } from 'vue'
  import { useAppStore } from "@/pinia"
  import { getEnvList } from '@/api/platform/env.js'
  import EnvManager from '@/components/platform/EnvManager.vue'

  defineOptions({
    name: 'EnvDetail'
  })

  // 提交按钮loading
  const btnLoading = ref(false)
  const appStore = useAppStore()

  // 控制更多查询条件显示/隐藏状态
  const showAllQuery = ref(false)

  // 环境列表
  const envList = ref([])

  // 表单数据
  const formData = ref({
    key: '',
    name: '',
    value: {}, // 使用对象存储，key为环境ID，value为对应值
  })

  // 验证规则
  const rule = reactive({
    key: [{
      required: true,
      message: '请输入变量名称',
      trigger: ['input','blur'],
    }, {
      whitespace: true,
      message: '不能只输入空格',
      trigger: ['input', 'blur'],
    }],
    name: [{
      required: true,
      message: '请输入中文名',
      trigger: ['input','blur'],
    }, {
      whitespace: true,
      message: '不能只输入空格',
      trigger: ['input', 'blur'],
    }],
  })

  const elFormRef = ref()
  const elSearchFormRef = ref()

  // 获取环境列表
  const getEnvListFun = async() => {
    const table = await getEnvList({ page: 1, pageSize: 99 })
    if (table.code === 0) {
      envList.value = table.data.list
    }
  }

  // 初始化环境值对象
  const initEnvValue = () => {
    const valueObj = {}
    envList.value.forEach(env => {
      valueObj[env.ID] = ''
    })
    return valueObj
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
    getTableData()
  }

  // 搜索
  const onSubmit = () => {
    elSearchFormRef.value?.validate(async(valid) => {
      if (!valid) return
      page.value = 1
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
    const table = await getEnvDetailList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
    if (table.code === 0) {
      // 确保每个数据项都有完整的value结构
      tableData.value = table.data.list.map(item => {
        if (!item.value || typeof item.value !== 'object') {
          item.value = initEnvValue()
        } else {
          // 确保包含所有环境的值
          envList.value.forEach(env => {
            if (item.value[env.ID] === undefined) {
              item.value[env.ID] = ''
            }
          })
        }
        return item
      })
      total.value = table.data.total
      page.value = table.data.page
      pageSize.value = table.data.pageSize
    }
  }

  // 表格中环境值变化处理
  const handleEnvValueChange = async (row, envId) => {
    try {
      const updateData = {
        ...row,
        value: { ...row.value }
      }
      const res = await updateEnvDetail(updateData)
      if (res.code === 0) {
        ElMessage.success('更新成功')
      }
    } catch (error) {
      console.error('更新环境值失败:', error)
      ElMessage.error('更新失败')
    }
  }

  // 初始化数据
  onMounted(async () => {
    await getEnvListFun()
    getTableData()
  })

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
      deleteEnvDetailFunc(row)
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
      const res = await deleteEnvDetailByIds({ IDs })
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
  const updateEnvDetailFunc = async(row) => {
    const res = await findEnvDetail({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
      formData.value = {
        ...res.data,
        value: res.data.value || initEnvValue()
      }
      dialogFormVisible.value = true
    }
  }

  // 删除行
  const deleteEnvDetailFunc = async (row) => {
    const res = await deleteEnvDetail({ ID: row.ID })
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
    formData.value = {
      key: '',
      name: '',
      value: initEnvValue(),
    }
    dialogFormVisible.value = true
  }

  // 关闭弹窗
  const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
      key: '',
      name: '',
      value: initEnvValue(),
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
          formData.value.key = 'ENV_' + formData.value.key
          res = await createEnvDetail(formData.value)
          break
        case 'update':
          res = await updateEnvDetail(formData.value)
          break
        default:
          res = await createEnvDetail(formData.value)
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
    const res = await findEnvDetail({ ID: row.ID })
    if (res.code === 0) {
      detailForm.value = {
        ...res.data,
        value: res.data.value || initEnvValue()
      }
      openDetailShow()
    }
  }

  // 关闭详情弹窗
  const closeDetailShow = () => {
    detailShow.value = false
    detailForm.value = {}
  }

  // 环境管理对话框控制
  const envManagerVisible = ref(false)

  // 打开环境管理对话框
  const openEnvManager = () => {
    envManagerVisible.value = true
  }

  // 环境变更后的回调，重新加载环境列表和表格数据
  const onEnvChanged = async () => {
    await getEnvListFun()
    getTableData()
  }
</script>

<style>
</style>