<template>
  <div style="display: flex">
    <div>
      <div class="left">
        <InterfaceTree
          menutype="1"
          @getTreeID="setTreeID"
          eventType="0"
        ></InterfaceTree>
      </div>
    </div>
    <div>
      <div class="container">
        <div class="left-table" id="left-table">
          <el-table
            border
            style="width: 620px"
            :data="apiData"
            ref="leftTable"
            row-key="ID"
            :cell-style="{ paddingTop: '5px', paddingBottom: '5px' }"
          >
            <el-table-column type="index" width="50"> </el-table-column>

            <el-table-column min-width="550" align="center" label="接口管理">
              <template #default="scope">
                <div
                  v-if="scope.row.request"
                  class="block"
                  :class="`block_${scope.row.request.method.toLowerCase()}`"
                >
                  <span
                    class="block-method block_method_color"
                    :class="`block_method_${scope.row.request.method.toLowerCase()}`"
                  >
                    {{ scope.row.request.method }}
                  </span>
                  <div class="block">
                    <span
                      class="block-method block_method_color block_method_options"
                      v-if="scope.row.creator === 'yapi'"
                      :title="'从YAPI导入的接口'"
                    >
                      YAPI
                    </span>
                  </div>
                  <span class="block-method block_url">{{
                    scope.row.request.url
                  }}</span>
                  <span class="block-summary-description">{{
                    scope.row.name
                  }}</span>
                </div>
                <div v-if="scope.row.gRPC" class="block" :class="`block_put`">
                  <span
                    class="block-method block_method_color"
                    :class="`block_method_put`"
                  >
                    {{ "gRPC" }}
                  </span>
                  <div class="block">
                    <span
                      class="block-method block_method_color block_method_options"
                      v-if="scope.row.creator === 'yapi'"
                      :title="'从YAPI导入的接口'"
                    >
                      YAPI
                    </span>
                  </div>

                  <span class="block-method block_url">{{
                    scope.row.gRPC.url
                  }}</span>
                  <span class="block-summary-description">{{
                    scope.row.name
                  }}</span>
                </div>
              </template>
            </el-table-column>
          </el-table>
        </div>
        <div class="right-table" id="right">
          <el-table
            border
            :key="tableKey"
            style="width: 620px"
            ref="rightTable"
            :data="caseData"
            row-key="ID"
            :cell-style="{ paddingTop: '5px', paddingBottom: '5px' }"
          >
            <el-table-column type="index" width="50"> </el-table-column>
            <el-table-column
              min-width="550"
              align="center"
              label="测试套件用例详情"
            >
              <template #default="scope">
                <div
                  v-if="scope.row.request"
                  class="block"
                  :class="`block_${scope.row.request.method.toLowerCase()}`"
                >
                  <span
                    class="block-method block_method_color"
                    :class="`block_method_${scope.row.request.method.toLowerCase()}`"
                  >
                    {{ scope.row.request.method }}
                  </span>
                  <div class="block">
                    <span
                      class="block-method block_method_color block_method_options"
                      v-if="scope.row.creator === 'yapi'"
                      :title="'从YAPI导入的接口'"
                    >
                      YAPI
                    </span>
                  </div>
                  <span class="block-method block_url">{{
                    scope.row.request.url
                  }}</span>
                  <span class="block-summary-description">{{
                    scope.row.name
                  }}</span>
                </div>
                <div v-if="scope.row.gRPC" class="block" :class="`block_put`">
                  <span
                    class="block-method block_method_color"
                    :class="`block_method_put`"
                  >
                    {{ "gRPC" }}
                  </span>
                  <div class="block">
                    <span
                      class="block-method block_method_color block_method_options"
                      v-if="scope.row.creator === 'yapi'"
                      :title="'从YAPI导入的接口'"
                    >
                      YAPI
                    </span>
                  </div>

                  <span class="block-method block_url">{{
                    scope.row.gRPC.url
                  }}</span>
                  <span class="block-summary-description">{{
                    scope.row.name
                  }}</span>
                </div>
              </template>
            </el-table-column>
          </el-table>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "testCaseAdd",
};
</script>

<script setup>
import draggable from "vuedraggable";
import InterfaceTree from "@/view/interface/interfaceComponents/interfaceTree.vue";
import { getInterfaceTemplateList } from "@/api/interfaceTemplate";

import Sortable from "sortablejs";
import { onMounted, ref } from "vue";
import { findTestCase, addTestCase, sortTestCase } from "@/api/testCase";
import { ElMessage } from "element-plus";
const isRouterActive = ref(true);

const props = defineProps({
  caseId: ref(),
});

let caseList = [];
const listTypes = ref("");
const apiIds = ref(0);
const caseIds = ref(0);
const menuIds = ref(0);
let treeID = 0;
let caseId = 0;
let tableKey = 0;
const apiData = ref([]);
const caseData = ref([]);

listTypes.value = "api";
// listTypes.value = "case"
apiIds.value = 1;
caseIds.value = 1;
caseIds.menuIds = 47;
const apiTypes = 1;

const getTestCaseDetailFunc = async (testCaseID) => {
  const res = await findTestCase({ ID: testCaseID });
  if (res.code === 0) {
    caseData.value = res.data.reapicase.TStep;
    // dialogFormVisible.value = true
  }
};

const addTestCaseFunc = async () => {};

const init = () => {
  caseId = props.caseId;
  getTestCaseDetailFunc(caseId);
};
init();

onMounted(() => {
  const leftTable = document
    .getElementById("left-table")
    .querySelector("tbody");
  const rightTable = document.getElementById("right").querySelector("tbody");
  leftDragHandler(leftTable);
  rightDragHandler(rightTable);
});

const setTreeID = (val) => {
  treeID = val;
  getTableData();
};

const leftDragHandler = async (dom, target) => {
  Sortable.create(dom, {
    sort: false,
    group: { name: "table-group", pull: "clone", put: false },
    animation: 1000,

    async onEnd(obj) {
      const { from, to, newIndex, oldIndex, index } = obj;
      if (from === to) {
      } else if (from !== to) {
        let tSteps = [];
        const sortData = ref({
          ID: 0,
          TStep: [],
        });
        let currRow = apiData.value[oldIndex];
        // 复制接口到用例下
        const res = await addTestCase({
          apiID: currRow.ID,
          caseID: Number(caseId),
        });
        if (res.code === 0) {
          ElMessage({
            type: "success",
            message: "添加套件成功",
          });
          let currRowData = res.data;
          caseData.value.splice(newIndex, 0, currRowData);
          const rightTable = document
            .getElementById("right")
            .querySelector("tbody");
          rightTable.deleteRow(newIndex);
          caseData.value.forEach((item, index, arr) => {
            let tStep = { ID: item.ID, sort: index + 1 };
            tSteps.push(tStep);
          });
          sortData.value.TStep = tSteps;
          const SortRes = await sortTestCase(sortData.value);
          if (SortRes.code === 0) {
            ElMessage({
              type: "success",
              message: "已完成排序",
            });
          }
          tableKey++;
        }
      }
    },
  });
};

const rightDragHandler = (dom, target) => {
  Sortable.create(dom, {
    sort: false,
    group: { name: "table-group", pull: true, put: true },
    animation: 1000,
  });
};

const log = (evt) => {
  // evt.added.element.id = 99
};

// 查询
const getTableData = async () => {
  const table = await getInterfaceTemplateList({
    type: apiTypes,
    menu: treeID,
    page: 1,
    pageSize: 99999,
  });
  if (table.code === 0) {
    apiData.value = table.data.list;
  }
};
</script>

<style lang="scss" scoped>
@import "src/style/apiList";
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

//.parent {
//  width: 85%;
//  height: 85%;
//}

.left {
  margin: 10px;
  width: 300px;
  height: 98%;
  padding: 8px;
  background: #ffffff;
}
</style>
