<template>
  <div>
    <div class="parent">
      <div id="g1">
        <el-table
            ref="multipleTable"
            style="width: 100%"
            :show-header="false"
            :data="tableData"
            row-key="ID"
            :cell-style="{paddingTop: '4px', paddingBottom: '4px'}"
            @keyup="rowDrop"
        >
          <el-table-column width="35" type="index">
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
                <!--                  <div>-->
                <!--                    <span class="el-icon-s-flag"-->
                <!--                          v-if="scope.row.cases.length > 0 "-->
                <!--                          :title="'API已经被用例引用,共计: '+scope.row.cases.length + '次'">-->
                <!--                    </span>-->
                <!--                  </div>-->
              </div>
            </template>
          </el-table-column>

          <!--          <el-table-column align="right" label="按钮组">-->
          <!--            <template #default="scope">-->
          <!--              <el-button type="text" icon="edit" size="small" class="table-button" @click="updateInterfaceTemplateFunc(scope.row)">变更</el-button>-->
          <!--              <el-button type="text" icon="delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>-->
          <!--            </template>-->
          <!--          </el-table-column>-->
        </el-table>
      </div>
      <!--        <div>-->
      <!--          <el-table-->
      <!--            ref="multipleTable"-->
      <!--            style="width: 100%"-->
      <!--            :show-header="false"-->
      <!--            :data="tableData"-->
      <!--            row-key="ID"-->
      <!--            :cell-style="{paddingTop: '4px', paddingBottom: '4px'}"-->
      <!--            @keyup="rowDrop"-->
      <!--        >-->
      <!--          <el-table-column width="35" type="index" >-->
      <!--          </el-table-column>-->
      <!--          <el-table-column-->
      <!--              min-width="600"-->
      <!--              align="center"-->
      <!--          >-->
      <!--            <template #default="scope">-->
      <!--              <div class="block" :class="`block_${scope.row.request.method.toLowerCase()}`">-->
      <!--                  <span class="block-method block_method_color"-->
      <!--                        :class="`block_method_${scope.row.request.method.toLowerCase()}`">-->
      <!--                    {{ scope.row.request.method }}-->
      <!--                  </span>-->
      <!--                <div class="block">-->
      <!--                    <span class="block-method block_method_color block_method_options"-->
      <!--                          v-if="scope.row.creator==='yapi'"-->
      <!--                          :title="'从YAPI导入的接口'">-->
      <!--                      YAPI-->
      <!--                    </span>-->
      <!--                </div>-->
      <!--                <span class="block-method block_url">{{ scope.row.request.url }}</span>-->
      <!--                <span class="block-summary-description">{{ scope.row.name }}</span>-->
      <!--                &lt;!&ndash;                  <div>&ndash;&gt;-->
      <!--                &lt;!&ndash;                    <span class="el-icon-s-flag"&ndash;&gt;-->
      <!--                &lt;!&ndash;                          v-if="scope.row.cases.length > 0 "&ndash;&gt;-->
      <!--                &lt;!&ndash;                          :title="'API已经被用例引用,共计: '+scope.row.cases.length + '次'">&ndash;&gt;-->
      <!--                &lt;!&ndash;                    </span>&ndash;&gt;-->
      <!--                &lt;!&ndash;                  </div>&ndash;&gt;-->
      <!--              </div>-->
      <!--            </template>-->
      <!--          </el-table-column>-->

      <!--          &lt;!&ndash;          <el-table-column align="right" label="按钮组">&ndash;&gt;-->
      <!--          &lt;!&ndash;            <template #default="scope">&ndash;&gt;-->
      <!--          &lt;!&ndash;              <el-button type="text" icon="edit" size="small" class="table-button" @click="updateInterfaceTemplateFunc(scope.row)">变更</el-button>&ndash;&gt;-->
      <!--          &lt;!&ndash;              <el-button type="text" icon="delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>&ndash;&gt;-->
      <!--          &lt;!&ndash;            </template>&ndash;&gt;-->
      <!--          &lt;!&ndash;          </el-table-column>&ndash;&gt;-->
      <!--        </el-table>-->
      <!--        </div>-->
    </div>
  </div>
</template>

<script>
export default {
  name: 'testCaseList'
}
</script>

<script setup>
import {
  createTestCase,
  deleteTestCase,
  deleteTestCaseByIds,
  updateTestCase,
  findTestCase,
  getTestCaseList,
  sortTestCase
} from '@/api/testCase'
import {useRoute} from "vue-router";
import {ref, watch} from "vue";
import {onMounted} from "vue"

import interfaceTempleForm from '@/view/interface/interfaceTemplate/interfaceTemplateForm.vue'
import {findInterfaceTemplate} from "@/api/interfaceTemplate";
import {reactive} from "vue";
import {ElMessageBox} from "element-plus";
import Sortable from 'sortablejs'

const props = defineProps({
  listType: ref(""),
  apiId: ref(0),
  caseId: ref(0),
  menuId: ref(0),
  data: ref({}),
})
let listType = ""
const route = useRoute()
let testCaseID = 1
const tableData = ref([])
const apiTypes = 2
const interfaceTempleFormVisible = ref(false)
const dialogTitle = ref(false)
const type = ref('')
const heightDiv = ref()
const caseName = ref()
let sortIdList = ""
heightDiv.value = window.screen.height - 480
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
  apiMenuID: '',
})
const sortData = ref({
  ID: 0,
  TStep: []
})

watch(() => props.data, (newValue, oldValue) => {
  tableData.value = newValue
  // console.log("props.data ==> ", newValue, oldValue);
});


const getTestCaseDetailFunc = async (testCaseID) => {
  const res = await findTestCase({ID: testCaseID})
  if (res.code === 0) {
    tableData.value = res.data.reapicase.TStep
    caseName.value = res.data.reapicase.name
    // dialogFormVisible.value = true
  }
}


const init = () => {
  tableData.value = props.data
  if (props.listType === "api") {

  } else if (props.listType === "case") {
    testCaseID = props.caseId
    if (testCaseID) {
      getTestCaseDetailFunc(testCaseID)
    }
  }
}


const updateInterfaceTemplateFunc = async (row) => {
  const res = await findInterfaceTemplate({ID: row.ID})
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
  }).then(() => {
    // deleteInterfaceTemplateFunc(row)
  })
}
onMounted(() => {
  rowDrop()
})

//行拖拽
const rowDrop = async () => {
  sortData.value.ID = Number(testCaseID)
  const tbody = document.getElementById('g1');

  Sortable.create(tbody, {
    async onEnd(evt) {
      // console.log("拖动了行", "当前序号: " + newIndex)
      // const { to, from, pullMode } = evt;
      // const toContext = window.bridge.get(to)
      // const fromContext = window.bridge.get(from)
      // let { newIndex, oldIndex, item } = evt;
      // const currRow = tableData.value.splice(oldIndex, 1)[0]
      // tableData.value.splice(newIndex, 0, currRow)
      // tableData.value.forEach((item, index, arr) => {
      //   let tStep = {ID:item.ID, sort:index+1}
      //   // tStep.ID = item.ID
      //   tSteps.push(tStep)
      // })
      // sortData.value.TStep = tSteps
      // const res = await sortTestCase(sortData.value)
    }

  })

}

// 关闭弹窗
const closeDialog = () => {
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
    apiMenuID: '',
  })
}

</script>

<style lang="scss" scoped>
@import 'src/style/apiList';
</style>
