<template>
  <div style="display: flex;">
    <div>
      <div class="left">
        <interfaceTree
            menutype=2
            @getTreeID="setTreeID"
            eventType="0"
        ></interfaceTree>
      </div>
    </div>
    <div>
      <div class="container">
        <div class="left-table" id="left">
          <el-table
              border
              style="width: 620px"
              :data="caseLeftData"
              ref='leftTable'
              row-key="ID"
              :cell-style="{paddingTop: '5px', paddingBottom: '5px'}"
          >
            <el-table-column
                type="index"
                width="50"
            >
            </el-table-column>

            <el-table-column
                min-width="550"
                align="center"
            >
              <template #default="scope">
                <div class="block" :class="`block_get`">
                  <span class="block-method block_method_color"
                        :class="`block_method_get`">
                    CASE
                  </span>
                  <div class="block">
                    <span class="block-method block_method_color block_method_options"
                          v-if="scope.row.creator==='yapi'"
                          :title="'从YAPI导入的接口'">
                      YAPI
                    </span>
                  </div>
                  <span class="block-method block_url">{{ scope.row.name }}</span>
                  <span class="block-method block_url">{{ scope.row.ID }}</span>
                </div>
              </template>
            </el-table-column>
          </el-table>
        </div>
        <div class="right-table" id="right" >
          <el-table
              border
              :key="tableKey"
              style="width: 620px"
              ref='rightTable'
              :data="caseData"
              row-key='ID'
              :cell-style="{paddingTop: '5px', paddingBottom: '5px'}"
          >
            <el-table-column
                type="index"
                width="50"
            >
            </el-table-column>
            <el-table-column
                min-width="550"
                align="center"
            >
              <template #default="scope">
                <div class="block" :class="`block_post`">
                  <span class="block-method block_method_color"
                        :class="`block_method_post`">
                    TASK
                  </span>
                  <div class="block">
                    <span class="block-method block_method_color block_method_options"
                          v-if="scope.row.creator==='yapi'"
                          :title="'从YAPI导入的接口'">
                      YAPI
                    </span>
                  </div>
                  <span class="block-method block_url">{{ scope.row.ApiTestCase.name }}</span>
                </div>
              </template>
            </el-table-column>
          </el-table>
        </div>
      </div>
    </div>
  </div>

</template>

<script  setup>
import draggable from "vuedraggable";
import interfaceTree from '@/view/interface/interfaceComponents/interfaceTree.vue'
import {
  getInterfaceTemplateList
} from '@/api/interfaceTemplate'

import Sortable from 'sortablejs'
import {onMounted, ref} from "vue"
import {
  findTestCase,
  addTestCase, sortTestCase, getTestCaseList
} from "@/api/testCase";
import {
  findTaskTestCase,
  addTaskTestCase,
} from "@/api/timerTask";
import {ElMessage} from "element-plus";
const isRouterActive = ref(true)


let caseList = []
const listTypes = ref("")
const apiIds = ref(0)
const caseIds = ref(0)
const menuIds = ref(0)
let treeID = 0
let caseId = 0
let tableKey = 0
const caseLeftData = ref([])
const caseData = ref([])

listTypes.value = "api"
// listTypes.value = "case"
apiIds.value = 1
caseIds.value = 1
caseIds.menuIds = 47
const apiTypes = 1

const props = defineProps({
  caseId: ref(),
})


const getTestCaseDetailFunc = async(testCaseID) => {
  const res = await findTaskTestCase({ ID: testCaseID })
  if (res.code === 0) {
    caseData.value = res.data.reapicase
  }
}


const init = () => {
  // caseId = props.caseId
  caseId = 2
  getTestCaseDetailFunc(caseId)
}
init()

onMounted(()=>{
  const leftTable = document.getElementById("left").querySelector('tbody')
  const rightTable = document.getElementById("right").querySelector('tbody')
  leftDragHandler(leftTable)
  rightDragHandler(rightTable)
})

const setTreeID = (val) => {
  treeID = val
  getTableData()
}


const leftDragHandler = async(dom, target) => {



  Sortable.create(dom, {

    sort: false,
    group: { name: 'table-group', pull: 'clone', put: false },
    animation:1000,

    async onEnd(obj) {
      const { from, to, newIndex, oldIndex,index } = obj
      if (from === to) {
      } else if (from !== to) {
        let tSteps = []
        const sortData = ref({
          ID: 0,
          TStep: []
        })
        let currRow = caseLeftData.value[oldIndex]
        const res = await addTaskTestCase({ apiID: currRow.ID, caseID: Number(caseId) })
        if (res.code === 0) {
          ElMessage({
            type: 'success',
            message: '添加用例成功'
          })
          let currRowData = res.data
          caseData.value.splice(newIndex, 0, currRowData)
          const rightTable = document.getElementById("right").querySelector('tbody')
          rightTable.deleteRow(newIndex)
          caseData.value.forEach((item, index, arr) => {
            let tStep = {ID:item.ID, sort:index+1}
            tSteps.push(tStep)
          })
          sortData.value.TStep = tSteps
          const SortRes = await sortTestCase(sortData.value)
          if (SortRes.code === 0) {
            ElMessage({
              type: 'success',
              message: '已完成排序'
            })
          }
          tableKey++
        }
      }
    }

  })
}

const rightDragHandler = (dom, target) => {
  Sortable.create(dom, {
    sort: false,
    group: { name: 'table-group', pull: true, put: true },
    animation:1000,
  })
}


const log = (evt) => {
  // evt.added.element.id = 99
}

// 查询
const getTableData = async() => {
  let menu = treeID
  const table = await getTestCaseList({menu: treeID, page: 1, pageSize: 99999 })
  if (table.code === 0) {
    caseLeftData.value = table.data.list
  }
}

</script>


<style lang="scss" scoped>
@import 'src/style/apiList';
.container {
  margin: 10px;
  display: flex;

}
.right-table {
  margin-left: 30px;
  height: 750px;
  overflow: auto;
}

.left-table {
  margin-left: 30px;
  height: 750px;
  overflow: auto;
}
.custom-tree-node {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-size: 14px;
  padding-right: 8px;
}

//.parent {
//  width: 85%;
//  height: 85%;
//}

.left {
  margin:10px;
  width: 300px;
  height: 98%;
  padding:8px;

}

</style>
