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
    @keyup="validateDatas"
  >
    <el-table-column fixed label="断言字段" width="200">
      <template #default="scope">
        <el-input
          clearable
          v-model="scope.row.check"
          placeholder="断言字段"
        ></el-input>
      </template>
    </el-table-column>

    <el-table-column label="断言类型" width="250">
      <template #default="scope">
        <el-select
          filterable
          v-model="scope.row.assert"
          placeholder="请选择断言类型"
        >
          <el-option
            v-for="item in validateOptions"
            :key="item.value"
            :label="item.value"
            :value="item.value"
          >
          </el-option>
        </el-select>
      </template>
    </el-table-column>

    <el-table-column label="期望类型" width="110">
      <template #default="scope">
        <el-select v-model="scope.row.type">
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

    <el-table-column label="期望返回值" width="300">
      <template #default="scope">
        <el-input
          clearable
          v-model="scope.row.expectTemp"
          placeholder="期望返回值"
        >
        </el-input>
      </template>
    </el-table-column>

    <el-table-column label="描述" width="240">
      <template #default="scope">
        <el-input
          clearable
          v-model="scope.row.msg"
          placeholder="描述"
        ></el-input>
      </template>
    </el-table-column>
    <el-table-column width="300">
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
  props: {
    save: Boolean,
    validate: ref(),
    heights: ref(),
  },
  computed: {
    height() {
      return this.heights - 70;
    },
  },
  created() {
    if (this.validate && this.validate.length !== 0) {
      this.tableData = this.validate;
    }
    this.Option();
    this.typeOption();
  },
  watch: {
    tableData: function () {
      this.validateDatas();
    },
  },
  methods: {
    async Option() {
      const res = await getDict("assert");
      res &&
        res.forEach((item) => {
          let validate = {};
          validate.value = item.label;
          this.validateOptions.push(validate);
        });
    },
    async typeOption() {
      const res = await getDict("assertType");
      res &&
        res.forEach((item) => {
          let validate = {};
          validate.label = item.label;
          validate.value = item.value;
          this.dataTypeOptions.push(validate);
        });
    },
    validateDatas() {
      let emitdata = [];
      for (let i in this.tableData) {
        let emitdataDict = {
          expect: "",
          check: "",
          msg: "",
          assert: "",
          expectTemp: "",
          type: 1,
        };
        if (this.tableData[i].check !== "") {
          emitdataDict.check = this.tableData[i].check;
          emitdataDict.msg = this.tableData[i].msg;
          emitdataDict.assert = this.tableData[i].assert;
          emitdataDict.expectTemp = this.tableData[i].expectTemp;
          emitdataDict.type = this.tableData[i].type;
          let value = this.tableData[i].expectTemp;
          let tempValue;
          switch (this.tableData[i].type) {
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
              if (
                value.toLowerCase() === "false" ||
                value.toLowerCase() === "true"
              ) {
                let bool = {
                  true: true,
                  false: false,
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
                if (value.indexOf("$") !== -1) {
                  tempValue = value;
                } else {
                  tempValue = false;
                }
              }
              break;
            case 7:
              // None 转 null
              if (value === "None") {
                tempValue = null;
              } else if (value.indexOf("$") !== -1) {
                tempValue = value;
              } else {
                tempValue = null;
              }
              break;
          }

          emitdataDict.expect = tempValue;
          emitdata.push(emitdataDict);
        }
      }
      this.$emit("requestValidateData", emitdata);
    },
    querySearch(queryString, cb) {
      let validateOptions = this.validateOptions;
      let results = queryString
        ? validateOptions.filter(this.createFilter(queryString))
        : validateOptions;
      cb(results);
    },

    createFilter(queryString) {
      return (validateOptions) => {
        return (
          validateOptions.value
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
        expect: "",
        check: "",
        assert: "equals",
        msg: "",
        expectTemp: "",
        type: 1,
      });
      this.validateDatas();
    },
    handleCopy(index, row) {
      this.tableData.splice(index + 1, 0, {
        expect: row.expect,
        check: row.check,
        assert: row.assert,
        msg: row.msg,
        expectTemp: row.expectTemp,
        type: row.type,
      });
      this.validateDatas();
    },
    handleDelete(index, row) {
      const data = this.tableData;
      data.splice(index, 1);
      this.validateDatas();
      if (data.length === 0) {
        this.handleEdit();
      }
    },

    assertInt(assert) {
      let number = "text";
      this.assertIntKey.forEach((item) => {
        if (assert === item) {
          number = "number";
        }
      });
      return number;
    },

    parseValidate() {
      let validate = {
        validate: [],
      };
      for (let content of this.tableData) {
        if (content["check"] !== "") {
          let obj = {};
          const expect = this.parseType(content["msg"], content["expect"]);
          if (expect === "exception") {
            continue;
          }
          obj[content["assert"]] = [content["check"], expect];
          validate.validate.push(obj);
        }
      }
      return validate;
    },
  },
  data() {
    return {
      currentValidate: "",
      currentRow: "",
      tableData: [
        {
          expect: "",
          check: "",
          assert: "equals",
          msg: "",
          expectTemp: "",
          type: 1,
        },
      ],
      assertIntKey: [
        "equals",
        "less_than",
        "less_or_equals",
        "greater_than",
        "greater_or_equals",
        "not_equal",
        "length_equals",
        "length_greater_than",
        "length_greater_or_equals",
        "length_less_than",
        "length_less_or_equals",
        "contains",
      ],
      // dataTypeOptions: [{
      //   label: 'String',
      //   value: 1
      // }, {
      //   label: 'Integer',
      //   value: 2
      // }, {
      //   label: 'Float',
      //   value: 3
      // }, {
      //   label: 'Boolean',
      //   value: 4
      // }, {
      //   label: 'List',
      //   value: 5
      // }, {
      //   label: 'Dict',
      //   value: 6
      // }],

      validateOptions: [],
      dataTypeOptions: [],
    };
  },
  name: "Validate",
};
</script>

<style scoped></style>
