<template>
  <div>
    <div style="margin-left: 200px;">
      <el-radio-group v-model="dataType">
        <el-radio
            v-for="item of dataOptions"
            :label="item.label"
            :key="item.value"
        >{{ item.value }}
        </el-radio>
      </el-radio-group>
    </div>
    <div style="margin-top: 5px">
      <el-table
          highlight-current-row
          :cell-style="{paddingTop: '4px', paddingBottom: '4px'}"
          strpe
          :height="height"
          :data="dataType === 'data' ? formData: paramsData"
          style="width: 100%;"
          @cell-mouse-enter="cellMouseEnter"
          @cell-mouse-leave="cellMouseLeave"
          v-show="dataType !== 'json' "
          @keyup="fromDatas"
      >
        <!--Request params-->
        <el-table-column
            label="请求Key"
            width="250">
          <template #default="scope">
            <el-input clearable v-model="scope.row.key" placeholder="Key"></el-input>
          </template>
        </el-table-column>
        <!--Request 表单-->
        <el-table-column
            v-if="dataType === 'data' "
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
            <!--            input 下加入 v-show="scope.row.type" 报错-->
            <el-input
                v-model="scope.row.value"
                placeholder="Value"
            ></el-input>

            <!--            <el-row v-show="scope.row.type === 5">-->
            <!--              <el-col :span="7">-->
            <!--                <el-upload-->
            <!--                    :show-file-list="false"-->
            <!--                    :action="uploadFile(scope.row)"-->
            <!--                    :limit="1"-->
            <!--                    type="small"-->
            <!--                    :file-list="fileList"-->
            <!--                    :on-error="uploadError"-->
            <!--                    :on-success="uploadSuccess"-->
            <!--                >-->
            <!--                  <el-button-->
            <!--                      size="small"-->
            <!--                      type="primary"-->
            <!--                      @click="currentIndex=scope.$index"-->
            <!--                  >选择文件-->
            <!--                  </el-button>-->
            <!--                </el-upload>-->
            <!--              </el-col>-->

            <!--              <el-col :span="12">-->
            <!--                <el-badge-->
            <!--                    :value="scope.row.size"-->
            <!--                    style="margin-top: 8px"-->
            <!--                >-->
            <!--                  <i class="el-icon-document" v-text="scope.row.value"></i>-->
            <!--                </el-badge>-->

            <!--              </el-col>-->
            <!--            </el-row>-->
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
            <el-row v-show="scope.row === currentRow">
              <el-button
                  icon="el-icon-circle-plus-outline"
                  size="mini"
                  type="info"
                  @click="handleEdit(scope.$index, scope.row)">
              </el-button>
              <el-button
                  icon="el-icon-document-copy"
                  size="mini"
                  type="info"
                  @click="handleCopy(scope.$index, scope.row)">
              </el-button>
              <el-button
                  icon="el-icon-delete"
                  size="mini"
                  type="danger"
                  v-show="scope.$index !== 0"
                  @click="handleDelete(scope.$index, scope.row)">
              </el-button>
            </el-row>

          </template>
        </el-table-column>


      </el-table>

      <!--Request json-->
      <div
          id="codeEditor"
          :style="contentStyleObj"
      ></div>

    </div>

  </div>


</template>

<script>
import 'jsoneditor/dist/jsoneditor.min.css'
import jsoneditor from 'jsoneditor';
import {ref} from "vue";

export default {
  props: {
    save: Boolean,
    request: {
      require: false
    },
    heights: ref(),
  },
  // computed: {
  //   height() {
  //     return window.screen.height - 464
  //   },
  // },
  name: "Request",
  components: {
    jsoneditor
  },


  watch: {
    save: function () {
      this.$emit('request', {
        form: this.parseForm(),
        json: this.parseJson(),
        params: this.parseParams(),
        // files: this.parseFile()
      }, {
        data: this.formData,
        params: this.paramsData,
        json_data: this.jsonData
      });
    },

    request: function () {
      if (this.request && this.request.length !== 0) {
        this.formData = this.request.data;
        this.jsonData = this.request.json_data;
        this.paramsData = this.request.params;
      }
      this.timeStamp = (new Date()).getTime()
    }
  },

  methods: {
    editorInit() {
      require('brace/ext/language_tools');
      require('brace/mode/json');
      require('brace/theme/github');
      require('brace/snippets/json');
    },
    getHeight() {
      this.contentStyleObj.height = this.height + 'px';
      this.contentStyleObj.width = '98%';
    },

    // uploadSuccess(response, file, fileList) {
    //   let size = file.size;
    //   if (size >= 1048576) {
    //     size = (size / 1048576).toFixed(2).toString() + 'MB';
    //   } else if (size >= 1024) {
    //     size = (size / 1024).toFixed(2).toString() + 'KB';
    //   } else {
    //     size = size.toString() + 'Byte'
    //   }
    //   this.formData[this.currentIndex]['value'] = file.name;
    //   this.formData[this.currentIndex]['size'] = size;
    //   this.fileList = [];
    //   if (!response.success) {
    //     this.$message.error(file.name + response.msg);
    //   }
    //
    // },

    // uploadFile(row) {
    //     return this.$api.uploadFile();
    // },

    // uploadError(error) {
    //     if (error.status === 401) {
    //         this.$router.replace({
    //             name: 'Login'
    //         })
    //     } else {
    //         this.$message.error('Sorry，文件上传失败啦, 请重试！')
    //     }
    // },

    cellMouseEnter(row) {
      this.currentRow = row;
    },

    cellMouseLeave(row) {
      this.currentRow = '';
    },

    handleEdit(index, row) {
      const data = this.dataType === 'data' ? this.formData : this.paramsData;
      data.push({
        key: '',
        value: '',
        type: 1,
        desc: ''
      });
    },
    handleCopy(index, row) {
      const data = this.dataType === 'data' ? this.formData : this.paramsData;
      data.splice(index + 1, 0, {
        key: row.key,
        value: row.value,
        type: row.type,
        desc: row.desc
      });
    },
    handleDelete(index, row) {
      const data = this.dataType === 'data' ? this.formData : this.paramsData;
      data.splice(index, 1);
    },

    // 文件格式化
    // parseFile() {
    //   let files = {
    //     files: {},
    //     desc: {}
    //   };
    //
    //   for (let content of this.formData) {
    //     // 是文件
    //     if (content['key'] !== '' && content['type'] === 5) {
    //       files.files[content['key']] = content['value'];
    //       files.desc[content['key']] = content['desc'];
    //     }
    //   }
    //   return files
    // },

    // 表单格式化
    parseForm() {
      let form = {
        data: {},
        desc: {}
      };
      for (let content of this.formData) {
        // file 不处理
        // if (content['key'] !== '' && content['type'] !== 5) {
        //   const value = this.parseType(content['type'], content['value']);
        //
        //   if (value === 'exception') {
        //     continue;
        //   }
        //
        //   form.data[content['key']] = value;
        //   form.desc[content['key']] = content['desc'];
        // }

        if (value === 'exception') {
          continue;
        }
        form.data[content['key']] = value;
        form.desc[content['key']] = content['desc'];
      }
      return form;
    },

    parseParams() {
      let params = {
        params: {},
        desc: {}
      };
      for (let content of this.paramsData) {
        if (content['key'] !== '') {
          params.params[content['key']] = content['value'];
          params.desc[content['key']] = content['desc'];
        }
      }
      return params;
    },

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
      codeEditor: null,
      treeEditor: null,
      syncData: true,
      contentStyleObj: {
        height: '',
        width: ''
      },
      fileList: [],
      currentIndex: 0,
      currentRow: '',
      jsonData: '',
      formData: [{
        key: '',
        value: '',
        type: 1,
        desc: ''
      }],
      paramsData: [{
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
        //   {
        //   label: 'File',
        //   value: 5
        // }
      ],

      dataOptions: [{
        label: 'data',
        value: '表单',
      }, {
        label: 'json',
        value: 'json',
      }, {
        label: 'params',
        value: 'params'
      }],
      dataType: 'json',
      timeStamp: "",
    }
  },
  computed: {
    height() {
      return this.heights - 105

    }
  },
  mounted: function () {
    // let treeEditorElement = document.getElementById('treeEditor')
    let codeEditorElement = document.getElementById('codeEditor')
    let json = {
      "Array": [1, 2, 3],
      "Boolean": true,
      "Null": null,
      "Number": 123,
      "Object": {"a": "b", "c": "d"},
      "String": "Hello World",
    };
    let codeOptions = {
      mode: 'code',
      modes: ['code', 'tree'],
    }
    // this.treeEditor = new jsoneditor(treeEditorElement, treeOptions, json);
    this.codeEditor = new jsoneditor(codeEditorElement, codeOptions, json);
  },
  created() {
    window.addEventListener('resize', this.getHeight);
    this.getHeight()
  },
}
</script>

<style scoped>
.ace_editor, .ace_editor * {
  font-family: "Monaco", "Menlo", "Ubuntu Mono", "Droid Sans Mono", "Consolas", monospace !important;
  font-size: 14px !important;
  font-weight: 400 !important;
  letter-spacing: 0 !important;
}
</style>

<!--<style scoped>-->
<!--.ace_editor, .ace_editor *{-->
<!--  font-family: "Monaco", "Menlo", "Ubuntu Mono", "Droid Sans Mono", "Consolas", monospace !important;-->
<!--  font-size: 14px !important;-->
<!--  font-weight: 400 !important;-->
<!--  letter-spacing: 0 !important;-->
<!--}-->
<!--</style>-->
