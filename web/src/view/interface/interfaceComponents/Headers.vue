<template>
  <el-table
      highlight-current-row
      :data="tableData"
      :height="height"
      style="width: 98%;"
      :border="false"
      :cell-style="{paddingTop: '4px', paddingBottom: '4px'}"
      @cell-mouse-enter="cellMouseEnter"
      @cell-mouse-leave="cellMouseLeave"
      @keyup="tableDatas"
  >
    <el-table-column
        label="标签"
        width="300"
    >
      <template #default="scope">
        <el-autocomplete
            v-model="scope.row.key"
            clearable
            :fetch-suggestions="querySearch"
            placeholder="头部标签"
        />
      </template>
    </el-table-column>

    <el-table-column
        label="内容"
        width="400"
    >
      <template #default="scope">
        <el-input v-model="scope.row.value" clearable placeholder="头部内容"/>
      </template>
    </el-table-column>

    <el-table-column
        label="描述"
        width="220"
    >
      <template #default="scope">
        <el-input v-model="scope.row.desc" clearable placeholder="头部信息简要描述"/>

      </template>
    </el-table-column>

    <el-table-column>
      <template #default="scope">
        <el-row size="mini">
          <el-button
              size="mini"
              type="info"
              @click="handleEdit(scope.$index, scope.row)"
          >
            <i size="mini" class="iconfont icon-add"></i>
          </el-button>

          <el-button
              size="mini"
              type="danger"
              @click="handleDelete(scope.$index, scope.row)"
          >
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
  name: 'Header',
  props: {
    save: Boolean,
    // eslint-disable-next-line vue/require-default-prop,vue/require-prop-types
    header: ref(),
    heights: ref()
  },
  data() {
    return {
      headerOptions: [{
        value: 'Accept'
      }, {
        value: 'Accept-Charset'
      }, {
        value: 'Accept-Language'
      }, {
        value: 'Accept-Datetime'
      }, {
        value: 'Authorization'
      }, {
        value: 'Cache-Control'
      }, {
        value: 'Connection'
      }, {
        value: 'Cookie'
      }, {
        value: 'Content-Length'
      }, {
        value: 'Content-MD5'
      }, {
        value: 'Content-Type'
      }, {
        value: 'Expect'
      }, {
        value: 'Date'
      }, {
        value: 'From'
      }, {
        value: 'Host'
      }, {
        value: 'If-Match'
      }, {
        value: 'If-Modified-Since'
      }, {
        value: 'If-None-Match'
      }, {
        value: 'If-Range'
      }, {
        value: 'If-Unmodified-Since'
      }, {
        value: 'Max-Forwards'
      }, {
        value: 'Origin'
      }, {
        value: 'Pragma'
      }, {
        value: 'Proxy-Authorization'
      }, {
        value: 'Range'
      }, {
        value: 'Referer'
      }, {
        value: 'TE'
      }, {
        value: 'User-Agent'
      }, {
        value: 'Upgrade'
      }, {
        value: 'Via'
      }, {
        value: 'Warning'
      }],
      currentRow: '',
      tableData: [{key: '', value: '', desc: ''}]
    }
  },
  created() {
    if (this.header && this.header.length !== 0) {
      this.tableData = this.header;
    }
  },
  computed: {
    height() {
      return this.heights - 70
    }
  },
  watch: {
    tableData: function () {
      this.$emit('headerData', this.tableData)
    },

    header: function () {
      if (this.header && this.header.length !== 0) {
        this.tableData = this.header
      }
    }
  },
  methods: {
    tableDatas() {
      let emitdata = []
      for (let i in this.tableData) {
        if (this.tableData[i].desc !== "" || this.tableData[i].key !== "" || this.tableData[i].value !== "") {
          emitdata.push(this.tableData[i])
        }
      }
      this.$emit('headerData', emitdata)
    },
    querySearch(queryString, cb) {
      const headerOptions = this.headerOptions
      const results = queryString ? headerOptions.filter(this.createFilter(queryString)) : headerOptions
      cb(results)
    },

    createFilter(queryString) {
      return (headerOptions) => {
        return (headerOptions.value.toLowerCase().indexOf(queryString.toLowerCase()) === 0)
      }
    },

    cellMouseEnter(row) {
      this.currentRow = row
    },

    cellMouseLeave(row) {
      this.currentRow = ''
    },

    addLine() {
      this.tableData.push({
        key: '',
        value: '',
        desc: ''
      })
    },

    handleEdit(index, row) {
      this.tableData.push({
        key: '',
        value: '',
        desc: ''
      });
      this.tableDatas()
    },

    handleDelete(index, row) {
      const data = this.tableData;
      this.$emit("request", data[index].ID)
      data.splice(index, 1);
      this.tableDatas()
      if (data.length === 0) {
        this.handleEdit()
      }
    },

    // 头部信息格式化
    parseHeader() {
      const header = {
        header: {},
        desc: {}
      }
      for (const content of this.tableData) {
        if (content['key'] !== '' && content['value'] !== '') {
          header.header[content['key']] = content['value']
          header.desc[content['key']] = content['desc']
        }
      }
      return header
    }
  }
}
</script>

<style scoped>
</style>
