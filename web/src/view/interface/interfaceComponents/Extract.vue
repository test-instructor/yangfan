<template>
  <el-table
      highlight-current-row
      strpe
      :height="height"
      :data="tableData"
      style="width: 98%;"
      @cell-mouse-enter="cellMouseEnter"
      @cell-mouse-leave="cellMouseLeave"
      :cell-style="{paddingTop: '4px', paddingBottom: '4px'}"
      @keyup="extractData"
  >
    <el-table-column
        label="变量名"
        width="300">
      <template #default="scope">
        <el-input clearable v-model="scope.row.key" placeholder="接收抽取值后的变量名"></el-input>
      </template>
    </el-table-column>
    <el-table-column
        label="抽取表达式"
        width="420">
      <template #default="scope">
        <el-input clearable v-model="scope.row.value" placeholder="抽取表达式"></el-input>

      </template>
    </el-table-column>

    <el-table-column
        label="描述"
        width="375">
      <template #default="scope">
        <el-input clearable v-model="scope.row.desc" placeholder="抽取值简要描述"></el-input>
      </template>
    </el-table-column>

    <el-table-column>
      <template #default="scope">
        <el-row>
          <el-button
              size="mini"
              type="info"
              @click="handleEdit(scope.$index, scope.row)">
            <i size="mini" class="iconfont icon-add"></i>
          </el-button>
          <el-button
              size="mini"
              type="info"
              @click="handleCopy(scope.$index, scope.row)">
            <i size="mini" class="iconfont icon-copy"></i>
          </el-button>
          <el-button
              size="mini"
              type="danger"
              v-show="lenData > 1"
              @click="handleDelete(scope.$index, scope.row)">
            <i size="mini" class="iconfont icon-delete"></i>
          </el-button>
        </el-row>

      </template>
    </el-table-column>
  </el-table>

</template>

<script>
import {ref} from "vue";

export default {
  props: {
    save: Boolean,
    extract: ref(),
    heights: ref(),
  },
  computed: {
    height() {
      return this.heights - 70
    },
  },
  watch: {
    tableData: function () {
      // this.$emit('requestExtractData', this.tableData);
      this.extractData()
    },
  },
  created() {
    this.lenData = this.extract.length
    if (this.extract && this.extract.length !== 0) {
      this.tableData = this.extract;
    }
  },

  methods: {
    extractData() {
      this.lenData = this.tableData.length
      this.$emit('requestExtractData', this.tableData)
    },
    cellMouseEnter(row) {
      this.currentRow = row;
    },

    cellMouseLeave(row) {
      this.currentRow = '';
    },

    handleEdit(index, row) {
      this.tableData.push({
        key: '',
        value: '',
        desc: ''
      });
      this.extractData()
    },
    handleCopy(index, row) {
      this.tableData.splice(index + 1, 0, {
        key: row.key,
        value: row.value,
        desc: row.desc
      });
      this.extractData()
    },

    handleDelete(index, row) {
      this.tableData.splice(index, 1);
      this.extractData()
    },
    // 抽取格式化
    parseExtract() {
      let extract = {
        extract: [],
        desc: {}
      };
      for (let content of this.tableData) {
        const key = content['key'];
        const value = content['value'];
        if (key !== '' && value !== '') {
          let obj = {};
          obj[key] = value;
          extract.extract.push(obj);
          extract.desc[key] = content['desc'];
        }
      }
      return extract;
    }
  },

  data() {
    return {
      lenData: 0,
      currentRow: '',
      tableData: [{
        key: '',
        value: '',
        desc: ''
      }]
    }
  },
  name: "Extract"
}
</script>

<style scoped>
</style>
