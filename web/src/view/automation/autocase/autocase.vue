<template>
  <div>
    <div style="display: flex; width: 100%; overflow: hidden;">
      <div style="width: 250px; flex-shrink: 0;">
        <ApiMenu
          :menutype="currentMenuType"
          @getTreeID="handleMenuClick"
        />
      </div>
      <div style="flex: 1; overflow-x: auto; min-width: 600px; padding-left: 16px;">
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

            <el-form-item label="用例名称" prop="caseName">
              <el-input v-model="searchInfo.caseName" placeholder="搜索条件" />
            </el-form-item>

            <el-form-item label="状态" prop="status">
              <el-input v-model="searchInfo.status" placeholder="搜索条件" />
            </el-form-item>

            <el-form-item label="运行配置" prop="configName">
              <el-input v-model="searchInfo.configName" placeholder="搜索条件" />
            </el-form-item>


            <template v-if="showAllQuery">
              <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
              <el-form-item label="运行环境" prop="envName">
                <el-input v-model="searchInfo.envName" placeholder="搜索条件" />
              </el-form-item>

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
            @sort-change="sortChange"
          >
            <el-table-column type="selection" width="55" />

            <el-table-column sortable align="left" label="日期" prop="CreatedAt" width="180">
              <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
            </el-table-column>

            <el-table-column sortable align="left" label="用例名称" prop="caseName" width="120" />

            <el-table-column align="left" label="类型" prop="type" width="120" />

            <el-table-column align="left" label="运行次数" prop="runNumber" width="120" />

            <el-table-column align="left" label="状态" prop="status" width="120" />

            <el-table-column align="left" label="运行环境" prop="envName" width="120" />

            <el-table-column align="left" label="运行配置" prop="configName" width="120" />

            <el-table-column align="left" label="操作" fixed="right" min-width="320">
              <template #default="scope">
                <div style="display: flex; align-items: center;">
                  <Runner case_type="case" :case_id="scope.row.ID" style="margin-right: 10px;" />
                  <el-button type="primary" link icon="step-api" class="table-button" @click="openStepDialog(scope.row)">
                    步骤管理
                  </el-button>
                  <el-button type="primary" link icon="edit" class="table-button" @click="updateAutoCaseFunc(scope.row)">
                    编辑
                  </el-button>
                  <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
                </div>
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
      </div>
    </div>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogFormVisible" :show-close="false"
               :before-close="closeDialog">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ type === 'create' ? '新增' : '编辑' }}</span>
          <div>
            <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
            <el-button @click="closeDialog">取 消</el-button>
          </div>
        </div>
      </template>

      <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item label="用例类型:" prop="type">
          <el-select v-model="formData.type" placeholder="请选择类型" style="width:100%" filterable :clearable="false">
             <el-option v-for="item in ['api','android','ios','harmony','browser']" :key="item" :label="item" :value="item" />
          </el-select>
        </el-form-item>
        <el-form-item label="用例名称:" prop="caseName">
          <el-input v-model="formData.caseName" :clearable="false" placeholder="请输入用例名称" />
        </el-form-item>
        <el-form-item label="运行次数:" prop="runNumber">
          <el-input v-model.number="formData.runNumber" :clearable="true" placeholder="请输入运行次数" />
        </el-form-item>
        <el-form-item label="状态:" prop="status">
          <el-select v-model="formData.status" placeholder="请选择状态" style="width:100%" filterable
                     :clearable="false">
            <el-option v-for="item in ['测试中','待评审','评审不通过','已发布','禁用','已废弃']" :key="item"
                       :label="item" :value="item" />
          </el-select>
        </el-form-item>
        <el-form-item label="运行配置:">
          <RunConfig v-model="formData.configID" @change="handleConfigSelect"/>
        </el-form-item>
        <el-form-item label="环境配置:">
          <Env v-model="formData.envID" @change="handleEnvChange"/>
        </el-form-item>
        <el-form-item label="描述:" prop="desc">
          <el-input v-model="formData.desc" :clearable="false" placeholder="请输入描述" />
        </el-form-item>
      </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true"
               :before-close="closeDetailShow" title="查看">
      <el-descriptions :column="1" border>
        <el-descriptions-item label="用例名称">
          {{ detailForm.caseName }}
        </el-descriptions-item>
        <el-descriptions-item label="运行次数">
          {{ detailForm.runNumber }}
        </el-descriptions-item>
        <el-descriptions-item label="状态">
          {{ detailForm.status }}
        </el-descriptions-item>
        <el-descriptions-item label="运行环境">
          {{ detailForm.envName }}
        </el-descriptions-item>
        <el-descriptions-item label="描述">
          {{ detailForm.desc }}
        </el-descriptions-item>
        <el-descriptions-item label="运行配置">
          {{ detailForm.configName }}
        </el-descriptions-item>
      </el-descriptions>
    </el-drawer>



    <el-drawer destroy-on-close size="90%" v-model="stepDialogVisible" :show-close="true"
               :before-close="closeStepDialog" title="步骤管理">
      <AutocaseStep :caseID="currentCaseID" :caseName="currentCaseName" />
    </el-drawer>

  </div>
</template>

<script setup>
  import {
    createAutoCase,
    deleteAutoCase,
    deleteAutoCaseByIds,
    updateAutoCase,
    findAutoCase,
    getAutoCaseList
  } from '@/api/automation/autocase'

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
  import { ref, reactive, computed, watch } from 'vue'
  import { useAppStore } from '@/pinia'
  import ApiMenu from '@/components/platform/menu/index.vue'

  import RunConfig from '@/components/platform/runConfig.vue'

  import Env from '@/components/platform/env.vue'
  import AutocaseStep from '@/view/automation/autocase/autocaseStep.vue'
  import Runner from '@/components/platform/Runner.vue'
  import { useRoute } from 'vue-router'

  defineOptions({
    name: 'AutoCase'
  })

  const route = useRoute()
  const normalizePlatform = (val) => {
    const raw = String(val || '').trim().toLowerCase()
    if (!raw) return 'api'
    if (raw === 'web') return 'browser'
    return raw
  }

  // 计算当前平台类型（优先 query.type，其次 meta.type，默认为 api）
  const currentPlatform = computed(() => {
    const hasQueryType = route.query?.type !== undefined && route.query?.type !== null && String(route.query.type).trim() !== ''
    if (hasQueryType) return normalizePlatform(route.query.type)
    return normalizePlatform(route.meta?.type) || 'api'
  })

  // 计算当前 MenuType（默认为 21）
  const currentMenuType = computed(() => {
    // 假设非 api 类型使用 autocase_{type} 作为菜单类型
    if (currentPlatform.value === 'api') {
      return '21'
    }
    return `autocase_${currentPlatform.value}`
  })

  // 提交按钮loading
  const btnLoading = ref(false)
  const appStore = useAppStore()

  // 控制更多查询条件显示/隐藏状态
  const showAllQuery = ref(false)

  // 自动化生成的字典（可能为空）以及字段
  const formData = ref({
    caseName: '',
    type: currentPlatform.value,
    runNumber: undefined,
    status: null,
    envName: '',
    desc: '',
    configName: '',
    configID: 0,
    envID: 0,
    menu: 0
  })


  // 验证规则
  const rule = reactive({
    caseName: [{
      required: true,
      message: '',
      trigger: ['input', 'blur']
    },
      {
        whitespace: true,
        message: '不能只输入空格',
        trigger: ['input', 'blur']
      }
    ],
    status: [{
      required: true,
      message: '',
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
  const searchInfo = ref({ type: currentPlatform.value })
  // 排序
  const sortChange = ({ prop, order }) => {
    const sortMap = {
      CreatedAt: 'created_at',
      ID: 'id',
      caseName: 'case_name'
    }

    let sort = sortMap[prop]
    if (!sort) {
      sort = prop.replace(/[A-Z]/g, match => `_${match.toLowerCase()}`)
    }

    searchInfo.value.sort = sort
    searchInfo.value.order = order
    getTableData()
  }
  // 重置
  const onReset = () => {
    const next = { type: currentPlatform.value }
    if (menuId.value !== null && menuId.value !== undefined) {
      next.menu = menuId.value
    }
    searchInfo.value = next
    getTableData()
  }

  // 搜索
  const onSubmit = () => {
    elSearchFormRef.value?.validate(async (valid) => {
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
  const getTableData = async () => {
    // 自动带上当前的 type 过滤
    const params = { page: page.value, pageSize: pageSize.value, ...searchInfo.value }
    if (!params.type) {
        params.type = currentPlatform.value
    }
    const table = await getAutoCaseList(params)
    if (table.code === 0) {
      tableData.value = table.data.list
      total.value = table.data.total
      page.value = table.data.page
      pageSize.value = table.data.pageSize
    }
  }

  // 监听路由变化，重置搜索条件并刷新列表
  watch(() => currentPlatform.value, (newVal) => {
      searchInfo.value = { type: newVal }
      formData.value.type = newVal
      menuId.value = null
      page.value = 1
      getTableData()
  })

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
      deleteAutoCaseFunc(row)
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
      const res = await deleteAutoCaseByIds({ IDs })
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
  const updateAutoCaseFunc = async (row) => {
    const res = await findAutoCase({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
      formData.value = res.data
      dialogFormVisible.value = true
    }
  }


  // 删除行
  const deleteAutoCaseFunc = async (row) => {
    const res = await deleteAutoCase({ ID: row.ID })
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
      caseName: '',
      type: currentPlatform.value,
      runNumber: undefined,
      status: null,
      envName: '',
      desc: '',
      configName: '',
      configID: 0,
      envID: 0,
      menu: menuId.value ?? 0
    }
    dialogFormVisible.value = true
  }

  // 关闭弹窗
  const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
      caseName: '',
      type: currentPlatform.value,
      runNumber: undefined,
      status: null,
      envName: '',
      desc: '',
      configName: '',
      configID: 0,
      envID: 0,
      menu: menuId.value ?? 0
    }
  }
  // 弹窗确定
  const enterDialog = async () => {
    btnLoading.value = true
    elFormRef.value?.validate(async (valid) => {
      if (!valid) return btnLoading.value = false
      let res
      if (!formData.value.type) {
        formData.value.type = currentPlatform.value
      }
      if ((formData.value.menu === null || formData.value.menu === undefined || formData.value.menu === 0) && menuId.value) {
        formData.value.menu = menuId.value
      }
      switch (type.value) {
        case 'create':
          res = await createAutoCase(formData.value)
          break
        case 'update':
          res = await updateAutoCase(formData.value)
          break
        default:
          res = await createAutoCase(formData.value)
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
    const res = await findAutoCase({ ID: row.ID })
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

  const menuId = ref(null)
  const handleMenuClick = (id) => {
    menuId.value = id
    // 可以在这里添加根据菜单ID筛选表格数据的逻辑
    searchInfo.value.menu = id
    getTableData()
  }

  const handleConfigSelect = (config) => {
    if (config) {
      formData.value.configName = config.name;
    } else {
      formData.value.configName = null;
    }
  }

  // 处理选择结果
  const handleEnvChange = (env) => {
    if (env) {
      console.log('选中的环境：', env.ID, env.name)
      formData.value.envName = env.name
    } else {
      formData.value.envName = null
    }
  }


  const stepDialogVisible = ref(false)
  const currentCaseID = ref(0)
  const currentCaseName = ref('')

  const openStepDialog = (row) => {
    console.log('打开步骤弹窗', row)
    currentCaseID.value = row.ID
    currentCaseName.value = row.caseName
    stepDialogVisible.value = true
  }

  const closeStepDialog = () => {
    stepDialogVisible.value = false
    currentCaseID.value = 0
    currentCaseName.value = ''
  }

</script>

<style>

</style>
