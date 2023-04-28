<template>
  <div>
    <el-form model="addTestCaseTag">
      <div style="padding: 5px">
        <el-tag
          style="margin: 3px 5px 3px 5px"
          v-for="item in addTestCaseTag"
          :key="item.key"
          class="mx-1"
          closable
          :disable-transitions="false"
          @close="handleClose(item)"
        >
          {{ item.name }}
        </el-tag>
      </div>
    </el-form>
  </div>
  <div style="display: flex">
    <div>
      <div class="left">
        <InterfaceTree
          menutype="3"
          @getTreeID="setTreeID"
          eventType="0"
        ></InterfaceTree>
      </div>
    </div>
    <div>
      <div class="container">
        <div class="left-table" id="left">
          <el-table
            border
            style="width: 820px"
            :data="apiData"
            ref="leftTable"
            row-key="ID"
            :cell-style="{ paddingTop: '5px', paddingBottom: '5px' }"
          >
            <el-table-column type="index" width="50"> </el-table-column>

            <el-table-column min-width="550" align="center">
              <template #default="scope">
                <div class="block" :class="`block_get`">
                  <span
                    class="block-method block_method_color"
                    :class="`block_method_get`"
                  >
                    {{ "CASE" }}
                  </span>
                  <div class="block"></div>
                  <span class="block-method block_url">{{
                    scope.row.name
                  }}</span>
                  <!--                  <span class="block-summary-description">{{ scope.row.name }}</span>-->
                </div>
              </template>
            </el-table-column>
            <el-table-column width="130">
              <template #default="scope">
                <el-button @click="addCase(scope.row)">
                  <span> 添加到任务 </span>
                </el-button>
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
  name: "TaskCaseAdd",
};
</script>

<script setup>
import draggable from "vuedraggable";
import InterfaceTree from "@/view/interface/interfaceComponents/interfaceTree.vue";
import { getInterfaceTemplateList } from "@/api/interfaceTemplate";

import Sortable from "sortablejs";
import { defineEmits, ref } from "vue";
import { findTestCase, addTestCase, sortTestCase } from "@/api/testCase";
import { ElMessage } from "element-plus";
import { getApiConfigList } from "@/api/apiConfig";

import { getApiCaseList } from "@/api/apiCase";

const isRouterActive = ref(true);
const emit = defineEmits(["close"]);

let caseList = [];
const listTypes = ref("");
const apiIds = ref(0);
const caseIds = ref(0);
const menuIds = ref(0);
let treeID = 0;
let tableKey = 0;
const apiData = ref([]);
const addTestCaseTag = ref([]);
let tagKey = 0;

listTypes.value = "api";
// listTypes.value = "case"
apiIds.value = 1;
caseIds.value = 1;
caseIds.menuIds = 47;
const apiTypes = 1;

const setTreeID = (val) => {
  treeID = val;
  getTableData();
};

const pushCase = () => {
  let caseID = [];
  addTestCaseTag.value.forEach((item, index) => {
    caseID.push(item.ID);
  });
  emit("caseID", caseID);
};

const addCase = (row) => {
  tagKey++;
  let rows = JSON.parse(JSON.stringify(row));
  rows.key = tagKey;
  addTestCaseTag.value.push(rows);
  pushCase();
};

const handleClose = (itemOld) => {
  let indexOld = 0;
  addTestCaseTag.value.forEach((item, index) => {
    if (itemOld.key === item.key) {
      indexOld = index;
    }
  });
  addTestCaseTag.value.splice(indexOld, 1);
  pushCase();
};

const log = (evt) => {
  // evt.added.element.id = 99
};

// 查询
const getTableData = async () => {
  const table = await getApiCaseList({
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

.left-table {
  margin-left: 30px;
  height: 580px;
  overflow: auto;
}

//.parent {
//  width: 85%;
//  height: 85%;
//}

.left {
  margin: 10px;
  width: 300px;
  height: 580px;
  padding: 8px;
  background: #ffffff;
}
</style>
