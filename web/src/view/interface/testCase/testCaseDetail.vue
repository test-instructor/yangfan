<template>

  <div>
    <div class="gva-search-box">

      <el-form :inline="true" class="demo-form-inline">
        <el-form-item label="用例名称">
          <el-input v-model="caseName" placeholder="用例名称" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="addApiCaseFunc" round>添加步骤</el-button>
        </el-form-item>

<!--        <el-form-item>-->
<!--          <el-button size="mini" type="primary" icon="search" @click="onSubmit">查询</el-button>-->
<!--          <el-button size="mini" icon="refresh" @click="onReset">重置</el-button>-->
<!--        </el-form-item>-->
      </el-form>
    </div>
    <div class="gva-table-box">
      <el-table
          ref="multipleTable"
          style="width: 100%"
          :show-header="false"
          :data="tableData"
          row-key="ID"
          :cell-style="{paddingTop: '4px', paddingBottom: '4px'}"
      >
        <el-table-column width="50" type="index" >
        </el-table-column>
        <el-table-column
            min-width="600"
            align="center"
        >
          <template #default="scope">
            <div class="block" :class="`block_${scope.row.request.method.toLowerCase()}`">
                  <span class="block-method block_method_color"
                        :class="`block_method_${scope.row.request.method.toLowerCase()}`">
                    {{ scope.row.request.method }}
                  </span>
              <div class="block">
                    <span class="block-method block_method_color block_method_options"
                          v-if="scope.row.creator==='yapi'"
                          :title="'从YAPI导入的接口'">
                      YAPI
                    </span>
              </div>
              <span class="block-method block_url">{{ scope.row.request.url }}</span>
              <span class="block-summary-description">{{ scope.row.name }}</span>
            </div>
          </template>
        </el-table-column>

        <el-table-column align="right" label="按钮组" width="150px">
          <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updateInterfaceTemplateFunc(scope.row)">变更</el-button>
            <el-button type="text" icon="delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <el-dialog
        v-model="interfaceTempleFormVisible"
        :before-close="closeDialog"
        :visible.sync="interfaceTempleFormVisible"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        :title="dialogTitle"
        width="1380px"
        top="30px"
    >
      <interfaceTempleForm
          @close="closeDialog"
          v-if="interfaceTempleFormVisible"
          :heights="heightDiv"
          :eventType="type"
          :apiType="apiTypes"
          :formData="formDatas"
          ref="menuRole">
      </interfaceTempleForm>

    </el-dialog>

    <el-dialog
        v-model="apiCaseVisible"
        :before-close="closeDialogAddCase"
        :visible.sync="apiCaseVisible"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        title="添加步骤"
        width="1700px"
        top="30px"
    >
      <testCaseAdd
          @close="closeDialogAddCase"
          v-if="apiCaseVisible"
          :caseId="testCaseID"
          ref="menuRole">
      </testCaseAdd>
    </el-dialog>

  </div>

</template>

<script setup>
import {
    createTestCase,
    deleteTestCase,
    deleteTestCaseByIds,
    updateTestCase,
    findTestCase,
    getTestCaseList,
    sortTestCase,
    addTestCase,
    delTestCase
} from '@/api/testCase'
import {useRoute} from "vue-router";
import {reactive, ref, onMounted, watch} from "vue";

import interfaceTempleForm from '@/view/interface/interfaceTemplate/interfaceTemplateForm.vue'
import testCaseAdd from "@/view/interface/testCase/testCaseAdd.vue"
import {findInterfaceTemplate} from "@/api/interfaceTemplate";
import {ElMessage, ElMessageBox} from "element-plus";
import Sortable from 'sortablejs'

const route = useRoute()
const testCaseID = ref()
const tableData = ref([])
const apiTypes = 2
const interfaceTempleFormVisible = ref(false)
const apiCaseVisible = ref(false)
const dialogTitle = ref(false)
const type = ref('')
const heightDiv = ref()
const caseName = ref()
let sortIdList = ""
heightDiv.value =  window.screen.height - 480
const formDatas = reactive({
  name: '',
  request: reactive({
    agreement: '',
    method: '',
    url: '',
    params: '',
    headers: '',
    json: '',
    data: '',
  }),
  variables: '',
  extract: '',
  validate: '',
  hooks: '',
  apiMenuID:'',
})
const sortData = ref({
  ID: 0,
  TStep: []
})

const  init = () => {
  // getDbFunc()
  // setFdMap()
  if (route.params.id>0){
    testCaseID.value = route.params.id
  }else {
    testCaseID.value = 1
  }


  if (testCaseID.value) {
    getTestCaseDetailFunc(testCaseID.value)
  }
}
const getTestCaseDetailFunc = async(testCaseID) => {
  const res = await findTestCase({ ID: testCaseID })
  if (res.code === 0) {
    tableData.value = res.data.reapicase.TStep
    caseName.value = res.data.reapicase.name
    // dialogFormVisible.value = true
  }
}
init()
watch(() => route.params.id, () => {
  if (route.params.id){
    init()
  }
})

watch(() => route.params.id, () => {
  if (route.params.id){
    init()
  }
})

const addApiCaseFunc = async() => {
  apiCaseVisible.value = true

}

const updateInterfaceTemplateFunc = async(row) => {
  const res = await findInterfaceTemplate({ ID: row.ID })
  type.value = 'update'
  dialogTitle.value = '编辑用例步骤'
  if (res.code === 0) {
    formDatas.value = res.data.reapicase
    interfaceTempleFormVisible.value = true
  }
}

const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const res = await delTestCase({apiID: row.ID, caseID: Number(testCaseID.value)})
    if (res.code === 0) {
      await getTestCaseDetailFunc(testCaseID.value)
      ElMessage({
        type: 'success',
        message: '用例删除成功'
      })
    }
  })
}
onMounted(()=>{
  rowDrop()
})

//行拖拽
const rowDrop= async() => {
  sortData.value.ID = Number(testCaseID.value)
  const tbody = document.querySelector('.el-table__body-wrapper tbody')
  let tSteps = []

  Sortable.create(tbody, {
    animation:1000,
    async onEnd({newIndex, oldIndex}) {
      const currRow = tableData.value.splice(oldIndex, 1)[0]
      tableData.value.splice(newIndex, 0, currRow)
      tableData.value.forEach((item, index, arr) => {
        let tStep = {ID:item.ID, sort:index+1}
        tSteps.push(tStep)
      })
      sortData.value.TStep = tSteps
      const res = await sortTestCase(sortData.value)
      if (res.code === 0) {
        await getTestCaseDetailFunc(testCaseID.value)
        ElMessage({
          type: 'success',
          message: '已完成排序'
        })
      }
    }

  })

}

// 关闭弹窗
const closeDialog = () => {
  getTestCaseDetailFunc(testCaseID.value)
  interfaceTempleFormVisible.value = false
  formDatas.value = reactive({
    name: '',
    request: reactive({
      agreement: '',
      method: '',
      url: '',
      params: '',
      headers: '',
      json: '',
      data: '',
    }),
    variables: '',
    extract: '',
    validate: '',
    hooks: '',
    apiMenuID:'',
  })
}

const closeDialogAddCase = () => {
  apiCaseVisible.value = false
  getTestCaseDetailFunc(testCaseID.value)
}

</script>

<style lang="scss" scoped>
@import 'src/style/apiList';
</style>
