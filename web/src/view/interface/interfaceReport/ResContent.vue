<template>
  <div class="wrap" id="resContent">
    <el-table
        :data="list"
        row-key="path"
        border
        default-expand-all
        :tree-props="{
        children: 'children',
        hasChildren: 'hasChildren'
      }"
    >
      <el-table-column
          v-for="(col, index) in cols"
          :key="index"
          :prop="col.key"
          :label="col.label"
      >
        <template #default="scope">
          <template v-if="col.key === 'path'">
            {{ scope.row.path }}
            <el-button
                type="success"
                size="mini"
                icon="el-icon-document-copy"
                circle
                @click="handleCopy(scope.row.path)"
            >
            </el-button>
          </template>

          <template v-else>
            {{ scope.row[col.key] }}
          </template>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
export default {
  name: 'ResContent',
  props: {
    // 数据源
    data: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
      cols: [
        {
          key: 'label',
          label: '字段名',
        },
        {
          key: 'value',
          label: '字段值',
        },
        {
          key: 'path',
          label: '取值路径',
        },
      ],
      list: []
    }
  },
  computed: {
    // String数据源转JSON数据源
    originJSON() {
      return this.data ? JSON.parse(this.data) : {};
    }
  },
  mounted() {
    const {originJSON} = this;

    this.list = this.escapeJsonToArray({
      originJSON,
      ancestorPath: 'content'
    });
  },
  methods: {
    /**
     * 转义JSON为数组
     */
    escapeJsonToArray({originJSON = {}, ancestorPath = ''}) {
      let escapeArray = [];

      Object.keys(originJSON).forEach(key => {
        let item = this.escapeDataToArr({
          originData: originJSON[key],
          ancestorPath,
          parentPath: key
        });

        escapeArray.push({
          key,
          label: key,
          value: originJSON[key],
          path: ancestorPath ? `${ancestorPath}.${key}` : key,
          children: [...item]
        });
      });

      return escapeArray;
    },

    /**
     * 递归转义
     */
    escapeDataToArr({originData = {}, ancestorPath = '', parentPath = ''}) {
      let item = [];

      if (originData !== null && originData !== undefined && originData.constructor === Object) {
        Object.keys(originData).forEach(key => {
          let subItem = this.escapeDataToArr({
            originData: originData[key],
            ancestorPath: parentPath || parentPath === 0 ? ancestorPath + '.' + parentPath : ancestorPath,
            parentPath: key
          });

          item.push({
            key,
            label: key,
            value: originData[key],
            path: `${ancestorPath}.${parentPath}.${key}`,
            children: [...subItem]
          });
        });
      }

      if (Array.isArray(originData)) {
        for (let index = 0; index < originData.length; index += 1) {
          let subItem = this.escapeDataToArr({
            originData: originData[index],
            ancestorPath: parentPath || parentPath === 0 ? ancestorPath + '.' + parentPath : ancestorPath,
            parentPath: index
          });

          item.push({
            key: index,
            label: index,
            value: JSON.stringify(originData[index]),
            path: `${ancestorPath}.${parentPath}.${index}`,
            children: [...subItem]
          });
        }
      }

      return item;
    },

    /**
     * 复制取值路径
     */
    handleCopy(data) {
      const oInput = document.createElement('input');

      oInput.value = data;
      document.body.appendChild(oInput);
      // 选择对象;
      oInput.select();
      // 执行浏览器复制命令
      document.execCommand("Copy");

      this.$message({
        message: '复制成功',
        type: 'success'
      });

      oInput.remove();
    },
  }
}
</script>
