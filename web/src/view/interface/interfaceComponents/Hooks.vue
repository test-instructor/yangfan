<template>
  <div style="display: flex;">
    <div>
      <el-table
          highlight-current-row
          :cell-style="{paddingTop: '4px', paddingBottom: '4px'}"
          strpe
          :height="height"
          :data="setupData"
          style="width: 98%;"
          @cell-mouse-enter="cellMouseEnter"
          @cell-mouse-leave="cellMouseLeave"
          @keyup="setupHooksDatas"
      >
        <el-table-column
            label="测试之前执行的方法"
            width="500">
          <template #default="scope">
            <el-input clearable
                      v-model="scope.row.setup"
                      placeholder="${ setup_hooks function($request, *args, **kwargs) }"
            >
            </el-input>
          </template>
        </el-table-column>

        <el-table-column width="150">
          <template #default="scope">
            <el-row>
              <el-button
                  size="mini"
                  type="info"
                  @click="setupHandleEdit(scope.$index, scope.row)">
                <i size="mini" class="iconfont icon-add"></i>
              </el-button>
              <el-button
                  size="mini"
                  type="danger"
                  v-show="setupLenData > 1"
                  @click="setupHandleDelete(scope.$index, scope.row)">
                <i size="mini" class="iconfont icon-delete"></i>
              </el-button>
            </el-row>

          </template>
        </el-table-column>

      </el-table>
    </div>
    <div>
      <el-table
          highlight-current-row
          :cell-style="{paddingTop: '4px', paddingBottom: '4px'}"
          strpe
          :height="height"
          :data="teardownData"
          style="width: 98%;"
          @cell-mouse-enter="cellMouseEnter"
          @cell-mouse-leave="cellMouseLeave"
          @keyup="teardownHooksDatas"
      >
        <el-table-column
            label="测试之后执行的方法"
            width="500">
          <template #default="scope">
            <el-input clearable
                      v-model="scope.row.teardown"
                      placeholder="${ teardown_hooks function(response, *args, **kwargs) }"
            >
            </el-input>
          </template>
        </el-table-column>


        <el-table-column width="150">
          <template #default="scope">
            <el-row>
              <el-button
                  size="mini"
                  type="info"
                  @click="teardownHandleEdit(scope.$index, scope.row)">
                <i size="mini" class="iconfont icon-add"></i>
              </el-button>
              <el-button
                  size="mini"
                  type="danger"
                  v-show="teardownLenData > 1"
                  @click="teardownHandleDelete(scope.$index, scope.row)">
                <i size="mini" class="iconfont icon-delete"></i>
              </el-button>
            </el-row>

          </template>
        </el-table-column>

      </el-table>
    </div>
  </div>
</template>

<script>
import {ref} from "vue";

export default {
  props: {
    save: Boolean,
    hooks: {
      require: false
    },
    heights: ref(),
    setupHooks: [],
    teardownHooks: [],
  },
  computed: {
    height() {
      return this.heights - 70
    }
  },
  created() {
    if (this.setupHooks&&this.setupHooks.length>0){
      let setups = []
      for (let content of this.setupHooks) {
        let setup = {setup: content}
        setups.push(setup)
      }
      this.setupData = setups
      this.setupHooksDatas()
    }

    if (this.teardownHooks&&this.teardownHooks.length>=1){
      let teardowns = []
      for (let content of this.teardownHooks) {
        let teardown = {teardown: content}
        teardowns.push(teardown)
      }
      this.teardownData = teardowns
      this.teardownHooksDatas()
    }


  },

  methods: {
    teardownHooksDatas() {
      this.teardownLenData = this.teardownData.length
      let teardownData = []
      this.teardownData.forEach((item, index, arr) => {
        teardownData.push(item.teardown)
      })
      this.$emit('teardownHooksData', teardownData)
    },
    setupHooksDatas() {
      let setupDatas = []
      this.setupLenData = this.setupData.length
      this.setupData.forEach((item, index, arr) => {
        setupDatas.push(item.setup)
      })
      this.$emit('setupHooksData', setupDatas)
    },
    cellMouseEnter(row) {
      this.currentRow = row;
    },

    cellMouseLeave(row) {
      this.currentRow = '';
    },


    teardownHandleEdit(index, row, flag) {
      this.teardownData.push({
        teardown: ''
      });
      this.teardownHooksDatas();
    },

    teardownHandleDelete(index, row) {
      this.teardownData.splice(index, 1);
      this.teardownHooksDatas();
    },


    setupHandleEdit(index, row, flag) {
      this.setupData.push({
        setup: ''
      });
      this.setupHooksDatas();
    },


    setupHandleDelete(index, row) {
      this.setupData.splice(index, 1);
      this.setupHooksDatas();
    },

    parse_hooks() {
      let hooks = {
        setup_hooks: [],
        teardown_hooks: []
      };
      for (let content of this.tableData) {
        if (content.setup !== '') {
          hooks.setup_hooks.push(content.setup);
        }
        if (content.teardown !== '') {
          hooks.teardown_hooks.push(content.teardown);
        }
      }
      return hooks;
    }
  },
  data() {
    return {
      lenData: 0,
      setupLenData: 0,
      teardownLenData: 0,
      currentRow: '',
      setupData:[{
        setup: ''
      }],
      teardownData:[{
        teardown: ''
      }]
    }
  },
  name: "Hooks",


}
</script>

<style scoped>
</style>
