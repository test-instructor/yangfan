<template>

  <div>
    <div class="gva-search-box">
      <el-form :inline="true" class="demo-form-inline">
        <el-form-item>
          <div class="block" :class="`block_head`" style="height: 30px">
                  <span class="block-method block_method_color"
                        :class="`block_method_head`">
                    {{ "CASE" }}
                  </span>
            <div class="block">
            </div>
            <span class="block-method block_url">{{ taskName }}</span>
          </div>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="addApiCaseFunc" round>添加测试用例</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <el-table
          ref="multipleTable"
          style="width: 100%"
          :show-header="false"
          :data="tableData"
          row-key="id"
          :cell-style="{paddingTop: '4px', paddingBottom: '4px'}"
      >
        <el-table-column width="50" type="index" >
        </el-table-column>
        <el-table-column
            min-width="600"
            align="center"
        >
          <template #default="scope">
            <div class="block" :class="`block_get`">
                  <span class="block-method block_method_color"
                        :class="`block_method_get`">
                    {{ "CASE" }}
                  </span>
              <span class="block-method block_url">{{ scope.row.case.name }}</span>
            </div>
          </template>
        </el-table-column>

        <el-table-column align="right" label="按钮组" width="150px">
          <template #default="scope">
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
      <InterfaceTempleForm
          @close="closeDialog"
          v-if="interfaceTempleFormVisible"
          :heights="heightDiv"
          :eventType="type"
          :apiType="apiTypes"
          :formData="formDatas"
          ref="menuRole">
      </InterfaceTempleForm>

    </el-dialog>

    <el-dialog
        v-model="taskCaseVisible"
        :before-close="closeDialogAddCase"
        :visible.sync="taskCaseVisible"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        title="添加测试用例"
        width="1250px"
        top="30px"
    >
      <TaskCaseAdd
          @close="closeDialogAddCase"
          v-if="taskCaseVisible"
          @caseID="addTeseCase"
          ref="menuRole">
      </TaskCaseAdd>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="addTaskCases" >确 定</el-button>
        </div>
      </template>
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

import {
  sortTaskCase,
  findTaskTestCase,
  addTaskCase,
  delTaskCase
} from '@/api/timerTask'

import {useRoute} from "vue-router";
import {reactive, ref, onMounted, watch} from "vue";

import InterfaceTempleForm from '@/view/interface/interfaceTemplate/interfaceTemplateForm.vue'
import TaskCaseAdd from "@/view/interface/timerTask/taskAddCase.vue"
import {findInterfaceTemplate} from "@/api/interfaceTemplate";
import {ElMessage, ElMessageBox} from "element-plus";
import Sortable from 'sortablejs'

const route = useRoute()
const task_id = ref()
const tableData = ref([])
const apiTypes = 2
const interfaceTempleFormVisible = ref(false)
const taskCaseVisible = ref(false)
const dialogTitle = ref(false)
const type = ref('')
const heightDiv = ref()
const taskName = ref()
let caseID = []
let sortIdList = ""
heightDiv.value =  window.screen.height - 480
const formDatas = reactive({})
const sortData = ref([])

const  init = () => {
  if (route.params.id>0){
    task_id.value = Number(route.params.id)
  }else {
    task_id.value = 1
  }
  if (task_id.value) {
    getTaskCaseDetailFunc(task_id.value)
  }
}
const getTaskCaseDetailFunc = async(task_id) => {
  const res = await findTaskTestCase({ ID: task_id })
  if (res.code === 0) {
    tableData.value = res.data.reapicase.test_case
    taskName.value = res.data.reapicase.name
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
  taskCaseVisible.value = true
}


const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const res = await delTaskCase({ID: row.id})
    if (res.code === 0) {
      await getTaskCaseDetailFunc(task_id.value)
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
  sortData.value.ID = Number(task_id.value)
  const tbody = document.querySelector('.el-table__body-wrapper tbody')
  let taskCases = []

  Sortable.create(tbody, {
    animation:1000,
    async onEnd({newIndex, oldIndex}) {
      const currRow = tableData.value.splice(oldIndex, 1)[0]
      tableData.value.splice(newIndex, 0, currRow)
      tableData.value.forEach((item, index, arr) => {
        let tStep = {ID:item.id, sort:index+1}
        taskCases.push(tStep)
      })
      sortData.value = taskCases
      const res = await sortTaskCase(taskCases)
      if (res.code === 0) {
        taskCases = []
        await getTaskCaseDetailFunc(task_id.value)
        ElMessage({
          type: 'success',
          message: '已完成排序'
        })
      }
    }
  })

}

// 关闭弹窗
const closeDialog =  () => {
  taskCaseVisible.value = false
  getTaskCaseDetailFunc(task_id.value)
}

const addTaskCases = async () => {
  const res = await addTaskCase({task_id: task_id.value, case_id: caseID})
  if (res.code === 0) {
    closeDialogAddCase()
    ElMessage({
      type: 'success',
      message: '添加用例成功'
    })
  }

}

const closeDialogAddCase = () => {
  taskCaseVisible.value = false
  getTaskCaseDetailFunc(task_id.value)
}

const addTeseCase = (caseIDs) => {
  caseID = caseIDs
}

</script>

<style lang="scss" scoped>
@import 'src/style/apiList';
</style>
