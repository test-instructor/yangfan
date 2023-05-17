<template>
  <el-table
    highlight-current-row
    ref="exportHeaderKey"
    :data="tableData"
    :height="height"
    style="width: 98%"
    :border="false"
    :cell-style="{ paddingTop: '4px', paddingBottom: '4px' }"
    @cell-mouse-enter="cellMouseEnter"
    @cell-mouse-leave="cellMouseLeave"
    @keyup="tableDatas"
    @selection-change="handleSelectionChange"
    row-key="rowKeyID"
  >
    <!--    <el-table-column width="10" type="index" >-->
    <el-table-column :reserve-selection="true" type="selection" width="55" />
    <el-table-column label="标签" width="300">
      <template #default="scope">
        <el-autocomplete
          v-model="scope.row.key"
          clearable
          :fetch-suggestions="querySearch"
          placeholder="头部标签"
        />
      </template>
    </el-table-column>

    <el-table-column label="内容" width="400">
      <template #default="scope">
        <el-input v-model="scope.row.value" clearable placeholder="头部内容" />
      </template>
    </el-table-column>

    <el-table-column label="描述" width="220">
      <template #default="scope">
        <el-input
          v-model="scope.row.desc"
          clearable
          placeholder="头部信息简要描述"
        />
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
import { ref } from "vue";
import { getDict } from "@/utils/dictionary";

export default {
  name: "Header",
  props: {
    save: Boolean,
    // eslint-disable-next-line vue/require-default-prop,vue/require-prop-types
    header: ref(),
    exportHeader: ref(),
    heights: ref(),
  },
  data() {
    return {
      rowKeyID: 0,
      multipleSelection: [],
      headerOptions: [],
      currentRow: "",
      tableData: [{ key: "", value: "", desc: "", rowKeyID: 0 }],
    };
  },
  created() {
    this.typeOption();
    if (this.header && this.header.length !== 0) {
      this.tableData = [];
      this.header.forEach((item) => {
        item["rowKeyID"] = this.rowKeyID++;
        this.tableData.push(item);
      });
    }
    if (this.exportHeader) {
      this.exportHeader.forEach((v) => {
        this.tableData.forEach((item) => {
          if (item.key === v) {
            this.multipleSelection.push(item);
          }
        });
      });

      this.$nextTick(() => {
        this.multipleSelection.forEach((row) => {
          this.$refs.exportHeaderKey.toggleRowSelection(row, true); // 回显
        });
      });
    }
  },
  computed: {
    height() {
      return this.heights - 70;
    },
  },
  watch: {
    tableData: function () {
      this.$emit("headerData", this.tableData);
    },

    header: function () {
      if (this.header && this.header.length !== 0) {
        this.tableData = this.header;
      }
    },
  },
  methods: {
    async typeOption() {
      const res = await getDict("requestHeader");
      res &&
        res.forEach((item) => {
          let header = {};
          header.value = item.label;
          this.headerOptions.push(header);
        });
    },
    handleSelectionChange(val) {
      this.multipleSelection = val;
      this.tableDatas();
    },
    tableDatas() {
      let emitdata = [];
      for (let i in this.tableData) {
        if (
          this.tableData[i].desc !== "" ||
          this.tableData[i].key !== "" ||
          this.tableData[i].value !== ""
        ) {
          emitdata.push(this.tableData[i]);
        }
      }
      this.$emit("headerData", emitdata);
      let exportHeader = [];
      this.multipleSelection &&
        this.multipleSelection.map((item) => {
          exportHeader.push(item.key);
        });
      this.$emit("exportHeader", exportHeader);
    },
    querySearch(queryString, cb) {
      const headerOptions = this.headerOptions;
      const results = queryString
        ? headerOptions.filter(this.createFilter(queryString))
        : headerOptions;
      console.log("results", results.value);
      cb(results);
    },

    createFilter(queryString) {
      return (headerOptions) => {
        return (
          headerOptions.value
            .toLowerCase()
            .indexOf(queryString.toLowerCase()) === 0
        );
      };
    },

    cellMouseEnter(row) {
      this.currentRow = row;
    },

    cellMouseLeave(row) {
      this.currentRow = "";
    },

    handleEdit(index, row) {
      this.tableData.push({
        key: "",
        value: "",
        desc: "",
        rowKeyID: this.rowKeyID++,
      });
      this.tableDatas();
    },

    handleDelete(index, row) {
      const data = this.tableData;
      this.$emit("request", data[index].ID);
      data.splice(index, 1);
      this.tableDatas();
      if (data.length === 0) {
        this.handleEdit();
      }
    },

    // 头部信息格式化
    parseHeader() {
      const header = {
        header: {},
        desc: {},
      };
      for (const content of this.tableData) {
        if (content["key"] !== "" && content["value"] !== "") {
          header.header[content["key"]] = content["value"];
          header.desc[content["key"]] = content["desc"];
        }
      }
      return header;
    },
  },
};
</script>

<style scoped></style>
