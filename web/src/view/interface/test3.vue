<template>
  <div>

    <el-table
        style="width: 1000px"
        ref="apiData"
        id="apiData"
        :data="tableData"
        :show-header="false"
    >
      <el-table-column
          width="100"
      >
        <template #default="scope" >
          <!--            <div id="success">-->
          <!--              <el-result-->
          <!--                  :icon="scope.row.success?'success':'warning'"-->
          <!--              >-->
          <!--              </el-result>-->
          <!--            </div>-->

          <el-button
              style="width:90px"
              :type="scope.row.success?'success':'danger'"
              plain
          >
            {{scope.row.success?'success':'fail'}}
          </el-button>
        </template>
      </el-table-column>
      <el-table-column
          min-width="550"
          align="center"
      >
        <template #default="scope">
          <div class="block" :class="`block_${scope.row.data.req_resps.request.method.toLowerCase()}`">
                <span class="block-method block_method_color"
                      :class="`block_method_${scope.row.data.req_resps.request.method.toLowerCase()}`">
                  {{ scope.row.data.req_resps.request.method }}
                </span>
            <span class="block-method block_url">{{ scope.row.data.req_resps.request.url }}</span>
            <span class="block-summary-description">{{ scope.row.name }}</span>
          </div>
        </template>
      </el-table-column>
      <el-table-column width="70">
        <template #default="scope">
          <el-button type="text" @click="toggleExpand(scope.row)">
            <span>{{scope.row.ID===currentIndex?'收起':'展开'}}</span>
          </el-button>
        </template>
      </el-table-column>
      <el-table-column type="expand" width="1">
        <template #default="scope">
          <div>
            <div>request</div>
            <div>response</div>
            <div>{{scope.row}}</div>
          </div>
        </template>
      </el-table-column>
    </el-table>



  </div>

</template>
<script setup>

import {getCurrentInstance, onMounted, ref} from "vue"


let tableData = [{"ID":1,"parntID":6,"name":"Header1","step_type":"request","success":false,"elapsed_ms":5,"httpstat":{"Connect":0,"ContentTransfer":0,"DNSLookup":0,"NameLookup":0,"Pretransfer":0,"ServerProcessing":4,"StartTransfer":4,"TCPConnection":0,"TLSHandshake":0,"Total":5},"data":{"success":false,"req_resps":{"request":{"body":{},"data":{"Form":"Form"},"headers":{"Content-Type":"application/json;charset=utf-8","Header":"Header"},"method":"GET","params":{"Params":"Params"},"url":"Header"},"response":{"body":"\"\"","cookies":{},"headers":{"Access-Control-Allow-Origin":"*","Connection":"keep-alive","Content-Length":"0","Date":"Wed,01Jun202210:31:07GMT","Keep-Alive":"timeout=5"},"proto":"HTTP/1.1","status_code":404}}},"content_size":0},{"ID":2,"parntID":7,"name":"Header2","step_type":"request","success":true,"elapsed_ms":3,"httpstat":{"Connect":0,"ContentTransfer":0,"DNSLookup":0,"NameLookup":0,"Pretransfer":0,"ServerProcessing":3,"StartTransfer":3,"TCPConnection":0,"TLSHandshake":0,"Total":3},"data":{"success":true,"req_resps":{"request":{"body":{},"data":{"Form":"Form"},"headers":{"Content-Type":"application/json;charset=utf-8","Header":"Header"},"method":"GET","params":{"Params":"Params"},"url":"Header"},"response":{"body":"\"\"","cookies":{},"headers":{"Access-Control-Allow-Origin":"*","Connection":"keep-alive","Content-Length":"0","Date":"Wed,01Jun202210:31:07GMT","Keep-Alive":"timeout=5"},"proto":"HTTP/1.1","status_code":404}}},"content_size":0},{"ID":4,"parntID":5,"name":"Header","step_type":"request","success":true,"elapsed_ms":4,"httpstat":{"Connect":0,"ContentTransfer":0,"DNSLookup":0,"NameLookup":0,"Pretransfer":0,"ServerProcessing":3,"StartTransfer":3,"TCPConnection":0,"TLSHandshake":0,"Total":4},"data":{"success":true,"req_resps":{"request":{"body":{},"data":{"Form":"Form"},"headers":{"Content-Type":"application/json;charset=utf-8","Header":"Header"},"method":"GET","params":{"Params":"Params"},"url":"Header"},"response":{"body":"\"\"","cookies":{},"headers":{"Access-Control-Allow-Origin":"*","Connection":"keep-alive","Content-Length":"0","Date":"Wed,01Jun202210:31:07GMT","Keep-Alive":"timeout=5"},"proto":"HTTP/1.1","status_code":404}}},"content_size":0,"export_vars":{"Extract":null}}]
let method = tableData[0].data.req_resps.request.url

let currentInstance
const currentIndex=ref(0);
onMounted(() => {
  currentInstance = getCurrentInstance()
})



const toggleExpand=(row) => {
  let table =  currentInstance.ctx.$refs.apiData
  tableData.map((item) => {
    if (row.ID !== item.ID) {
      table.toggleRowExpansion(item, false)
    }
  })
  table.toggleRowExpansion(row)
  if (currentIndex.value === row.ID){
    currentIndex.value = 0
  }else {
    currentIndex.value = row.ID
  }

}


</script>
<style lang="scss" scoped>
@import 'src/style/apiList';
.container {
  margin: 100px;
  display: flex;

}
.apiData-table {
  margin-left: 30px;
}
.success {
  height: 10px;
  display: flex;

}
</style>
