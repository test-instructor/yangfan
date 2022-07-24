<template>
  <el-table
      highlight-current-row
      :cell-style="{paddingTop: '4px', paddingBottom: '4px'}"
      strpe
      :data="formData"
      :height="height"
      style="width: 98%;"
      @cell-mouse-enter="cellMouseEnter"
      @cell-mouse-leave="cellMouseLeave"
      @keyup="fromDatas"
  >
    <!--Request params-->
    <el-table-column
        label="请求Key"
        width="220">
      <template #default="scope">
        <el-input clearable v-model="scope.row.key" placeholder="Key"></el-input>
      </template>
    </el-table-column>
    <!--Request 表单-->
    <el-table-column
        label="类型"
        width="120"
    >
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

    <el-table-column
        label="请求Value"
        width="350">
      <template #default="scope">
        <el-input
            v-model="scope.row.value"
            placeholder="Value"
        ></el-input>
      </template>
    </el-table-column>

    <el-table-column
        label="描述"
        width="400">
      <template #default="scope">
        <el-input clearable v-model="scope.row.desc" placeholder="参数简要描述"></el-input>
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
    forms: ref(),
    heights: ref(),
  },
  name: "Forms",

  methods: {
    getHeight() {
      this.contentStyleObj.height = this.height + 'px';
      this.contentStyleObj.width = '98%';
    },
    fromDatas() {
      let emitdata = []
      for (let i in this.formData) {
        if (this.formData[i].desc !== "" || this.formData[i].key !== "" || this.formData[i].value !== "") {
          emitdata.push(this.formData[i])
        }
      }
      this.$emit('requestFormData', emitdata)
    },
    cellMouseEnter(row) {
      this.currentRow = row;
    },

    cellMouseLeave(row) {
      this.currentRow = '';
    },

    handleEdit() {
      const data = this.formData
      data.push({
        key: '',
        value: '',
        type: 1,
        desc: ''
      });
    },
    handleCopy(index, row) {
      const data = this.formData;
      data.splice(index + 1, 0, {
        key: row.key,
        value: row.value,
        type: row.type,
        desc: row.desc
      });
      this.fromDatas();
    },
    handleDelete(index, row) {
      const data = this.formData;
      this.$emit("request", data[index].ID)
      data.splice(index, 1);
      this.fromDatas();
      if (data.length === 0) {
        this.handleEdit()
      }
    },

    // 表单格式化
    parseJson() {
      let json = {};
      if (this.jsonData !== '') {
        try {
          json = JSON.parse(this.jsonData);
        } catch (err) {
          this.$notify.error({
            title: 'json错误',
            message: '不是标准的json数据格式',
            duration: 2000
          });
        }
      }
      return json;
    },


    // 类型转换
    parseType(type, value) {
      let tempValue;
      const msg = value + ' => ' + this.dataTypeOptions[type - 1].label + ' 转换异常, 该数据自动剔除';
      switch (type) {
        case 1:
          tempValue = value;
          break;
        case 2:
          tempValue = parseInt(value);
          break;
        case 3:
          tempValue = parseFloat(value);
          break;
        case 4:
          if (value === 'False' || value === 'True') {
            let bool = {
              'True': true,
              'False': false
            };
            tempValue = bool[value];
          } else {
            this.$notify.error({
              title: '类型转换错误',
              message: msg,
              duration: 2000
            });
            return 'exception'
          }
          break;
        case 5:
        case 6:
          try {
            tempValue = JSON.parse(value);
          } catch (err) {
            // 包含$是引用类型,可以任意类型
            if (value.indexOf("$") != -1) {
              tempValue = value
            } else {
              tempValue = false
            }
          }
          break;
      }
      if (tempValue !== 0 && !tempValue && type !== 4 && type !== 1) {
        this.$notify.error({
          title: '类型转换错误',
          message: msg,
          duration: 2000
        });
        return 'exception'
      }
      return tempValue;
    },
  },


  data() {
    return {
      contentStyleObj: {
        height: '',
        width: ''
      },
      fileList: [],
      currentIndex: 0,
      currentRow: '',
      formData: [{
        key: '',
        value: '',
        type: '',
        desc: ''
      }],
      dataTypeOptions: [{
        label: 'String',
        value: 1
      }, {
        label: 'Integer',
        value: 2
      }, {
        label: 'Float',
        value: 3
      }, {
        label: 'Boolean',
        value: 4
      },
      ],
      timeStamp: "",
    }
  },
  computed: {
    height() {
      return this.heights - 70

    }
  },

  created() {
    window.addEventListener('resize', this.getHeight);
    this.getHeight()
    if (this.forms && this.forms.length !== 0) {
      this.formData = this.forms;

      this.fromDatas()
    }
  },
}
</script>

<style scoped>

</style>
