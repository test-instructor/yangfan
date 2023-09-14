<template>
  <el-table
    highlight-current-row
    :cell-style="{ paddingTop: '4px', paddingBottom: '4px' }"
    strpe
    :height="height"
    :data="tableData"
    style="width: 98%"
    @cell-mouse-enter="cellMouseEnter"
    @cell-mouse-leave="cellMouseLeave"
    @keyup="variablesDatas"
  >
    <el-table-column label="变量名" width="220">
      <template #default="scope">
        <el-input
          clearable
          v-model="scope.row.key"
          placeholder="Key"
        ></el-input>
      </template>
    </el-table-column>

    <el-table-column label="类型" width="120">
      <template #default="scope">
        <el-select v-model="scope.row.type" @change="handleTypeChange()">
          <el-option
            v-for="item in dataTypeOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          >
          </el-option>
        </el-select>
      </template>
    </el-table-column>

    <el-table-column label="变量值" width="375">
      <template #default="scope">
        <el-input
          clearable
          v-model="scope.row.valueTemp"
          placeholder="Value"
        ></el-input>
      </template>
    </el-table-column>

    <el-table-column label="内容" width="375">
      <template #default="scope">
        <el-input
          clearable
          v-model="scope.row.desc"
          placeholder="变量简要描述"
        ></el-input>
      </template>
    </el-table-column>

    <el-table-column>
      <template #default="scope">
        <el-row>
          <el-button
            size="mini"
            type="info"
            @click="handleEdit(scope.$index, scope.row)"
          >
            <i size="mini" class="iconfont icon-add"></i>
          </el-button>
          <el-button
            size="mini"
            type="info"
            @click="handleCopy(scope.$index, scope.row)"
          >
            <i size="mini" class="iconfont icon-copy"></i>
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
  name: "Variables",

  props: {
    save: Boolean,
    variables: ref(),
    heights: ref(),
  },
  computed: {
    height() {
      return this.heights - 70;
    },
  },

  methods: {
    async typeOption() {
      const res = await getDict("variablesType");
      res &&
        res.forEach((item) => {
          let validate = {};
          validate.label = item.label;
          validate.value = item.value;
          this.dataTypeOptions.push(validate);
        });
    },
    handleTypeChange() {
      this.variablesDatas();
    },
    variablesDatas() {
      let emitdata = [];
      for (let i in this.tableData) {
        let emitdataDict = {
          type: this.tableData[i].type,
          desc: this.tableData[i].desc,
          key: this.tableData[i].key,
          valueTemp: this.tableData[i].valueTemp,
          value: this.tableData[i].valueTemp,
        };
        if (
          this.tableData[i].desc !== "" ||
          this.tableData[i].key !== "" ||
          this.tableData[i].valueTemp !== ""
        ) {
          if (this.tableData[i].valueTemp !== "") {
            let type = this.tableData[i].type;
            let tempValue = this.tableData[i].value;
            let value = this.tableData[i].valueTemp;
            switch (type) {
              case 1:
                tempValue = value;
                break;
              case 2:
                // 包含$是引用类型,可以任意类型
                if (value.indexOf("$") !== -1) {
                  tempValue = value;
                } else {
                  tempValue = parseInt(value);
                }
                break;
              case 3:
                tempValue = parseFloat(value);
                break;
              case 4:
                if (value === "False" || value === "True") {
                  let bool = {
                    True: true,
                    False: false,
                  };
                  tempValue = bool[value];
                } else {
                  // this.$notify.error({
                  //   title: '类型转换错误',
                  //   message: msg,
                  //   duration: 2000
                  // });
                  // return 'exception'
                }
                break;
              case 5:
              case 6:
                try {
                  tempValue = JSON.parse(value);
                } catch (err) {
                  // 包含$是引用类型,可以任意类型
                  if (value.indexOf("$") != -1) {
                    tempValue = value;
                  } else {
                    tempValue = false;
                  }
                }
                break;
            }
            emitdataDict.value = tempValue;
            emitdata.push(emitdataDict);
          }
        }
      }
      this.$emit("requestVariablesData", emitdata);
    },

    cellMouseEnter(row) {
      this.currentRow = row;
    },

    cellMouseLeave(row) {
      this.currentRow = "";
    },

    handleEdit(index, row) {
      this.tableData.splice(index + 1, 0, {
        key: "",
        value: "",
        type: 1,
        desc: "",
      });
      this.variablesDatas();
    },
    handleCopy(index, row) {
      this.tableData.splice(index + 1, 0, {
        key: row.key,
        value: row.value,
        type: row.type,
        desc: row.desc,
      });
      this.variablesDatas();
    },
    handleDelete(index, row) {
      const data = this.tableData;
      this.$emit("request", data[index].ID);
      data.splice(index, 1);
      this.variablesDatas();
      if (data.length === 0) {
        this.handleEdit();
      }
    },

    // 类型转换
  },
  data() {
    return {
      currentRow: "",
      tableData: [
        {
          key: "",
          value: "",
          valueTemp: "",
          type: 1,
          desc: "",
        },
      ],

      dataTypeOptions: [],
      dataType: "data",
    };
  },
  created() {
    this.typeOption();
    if (this.variables && this.variables.length !== 0) {
      this.tableData = this.variables;
      this.variablesDatas();
    }
  },
};
</script>

<style scoped></style>
