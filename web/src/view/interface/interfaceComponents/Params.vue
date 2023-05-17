<template>
  <el-table
    highlight-current-row
    :cell-style="{ paddingTop: '4px', paddingBottom: '4px' }"
    strpe
    :height="height"
    :data="paramsData"
    style="width: 98%"
    @cell-mouse-enter="cellMouseEnter"
    @cell-mouse-leave="cellMouseLeave"
    @keyup="paramsDatas"
  >
    <!--Request params-->
    <el-table-column label="请求Key" width="250">
      <template #default="scope">
        <el-input
          clearable
          v-model="scope.row.key"
          placeholder="Key"
        ></el-input>
      </template>
    </el-table-column>
    <el-table-column label="请求Value" width="350">
      <template #default="scope">
        <el-input v-model="scope.row.value" placeholder="Value"></el-input>
      </template>
    </el-table-column>

    <el-table-column label="描述" width="400">
      <template #default="scope">
        <el-input
          clearable
          v-model="scope.row.desc"
          placeholder="参数简要描述"
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

export default {
  props: {
    save: Boolean,
    params: ref(),
    heights: ref(),
  },
  name: "Params",
  components: {},

  methods: {
    getHeight() {
      this.contentStyleObj.height = this.height + "px";
      this.contentStyleObj.width = "98%";
    },
    paramsDatas() {
      let emitdata = [];
      for (let i in this.paramsData) {
        if (
          this.paramsData[i].desc !== "" ||
          this.paramsData[i].key !== "" ||
          this.paramsData[i].value !== ""
        ) {
          emitdata.push(this.paramsData[i]);
        }
      }
      this.$emit("requestParamsData", emitdata);
    },

    cellMouseEnter(row) {
      this.currentRow = row;
    },

    cellMouseLeave(row) {
      this.currentRow = "";
    },

    handleEdit(index, row) {
      const data = this.paramsData;
      data.push({
        key: "",
        value: "",
        type: 1,
        desc: "",
      });
    },
    handleCopy(index, row) {
      const data = this.paramsData;
      data.splice(index + 1, 0, {
        key: row.key,
        value: row.value,
        type: row.type,
        desc: row.desc,
      });
      this.paramsDatas();
    },
    handleDelete(index, row) {
      const data = this.paramsData;
      this.$emit("request", data[index].ID);
      data.splice(index, 1);
      this.paramsDatas();
      if (data.length === 0) {
        this.handleEdit();
      }
    },

    // 表单格式化
    parseForm() {
      let form = {
        data: {},
        desc: {},
      };
      for (let content of this.formData) {
        if (value === "exception") {
          continue;
        }
        form.data[content["key"]] = value;
        form.desc[content["key"]] = content["desc"];
      }
      return form;
    },

    parseParams() {
      let params = {
        params: {},
        desc: {},
      };
      for (let content of this.paramsData) {
        if (content["key"] !== "") {
          params.params[content["key"]] = content["value"];
          params.desc[content["key"]] = content["desc"];
        }
      }
      return params;
    },
  },

  data() {
    return {
      contentStyleObj: {
        height: "",
        width: "",
      },
      currentRow: "",
      paramsData: [
        {
          key: "",
          value: "",
          type: "",
          desc: "",
        },
      ],
      timeStamp: "",
    };
  },
  computed: {
    height() {
      return this.heights - 70;
    },
  },
  created() {
    window.addEventListener("resize", this.getHeight);
    this.getHeight();

    if (this.params) {
      if (this.params.length !== 0) {
        this.paramsData = this.params;
        this.paramsDatas();
      }
    } else {
      this.paramsData = [
        {
          key: "",
          value: "",
          type: "",
          desc: "",
        },
      ];
    }
  },
};
</script>

<style scoped></style>
