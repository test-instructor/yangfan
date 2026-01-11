<template>
  <div>
    <el-button type="primary" @click="openPythonFuncDialog" style="margin-left: 15px;margin-right: 15px;">Python函数</el-button>

    <el-dialog
      title="选择Python函数"
      v-model="showPythonFuncDialog"
      width="800px"
      destroy-on-close
    >
      <el-table
        :data="pythonFuncList"
        border
        size="small"
        row-key="ID"
        :expand-row-keys="expandedRowKeys"
        @row-click="handleRowClick"
        style="width: 100%"
        height="500px"
      >
        <el-table-column label="函数格式" align="left" width="flex:1">
          <template #default="scope">
      <span class="font-medium text-primary">
        ${{ `{${scope.row.name}(${scope.row.params.map(param => '$' + param).join(', ')})}` }}
      </span>
          </template>
        </el-table-column>

        <el-table-column label="操作" align="center" width="120">
          <template #default="scope">
            <el-button
              icon="copy-document"
              type="text"
              @click="copyFuncName(scope.row)"
              size="small"
              @click.stop
            >
              复制
            </el-button>
          </template>
        </el-table-column>

        <el-table-column type="expand">
          <template #default="scope">
            <div class="p-2 bg-gray-50 rounded text-sm">
              <div class="text-gray-600 mb-1 font-medium">代码内容：</div>
              <pre class="text-gray-800 overflow-x-auto whitespace-pre-wrap">{{ scope.row.fullCode }}</pre>
            </div>
          </template>
        </el-table-column>
      </el-table>

      <div v-if="pythonFuncList.length === 0" class="text-center text-gray-500 py-4">
        {{ "暂无Python函数数据" }}
      </div>
    </el-dialog>
  </div>
</template>

<script setup>
  import { onMounted, ref } from 'vue'
  import { ElMessage } from 'element-plus'
  import { getPythonCodeFuncList } from '@/api/platform/pythoncodefunc.js'

  // 弹窗控制
  const showPythonFuncDialog = ref(false)
  const openPythonFuncDialog = () => {
    showPythonFuncDialog.value = true
  }

  // Python函数列表
  const pythonFuncList = ref([])
  // 拉取函数列表：组件内部自动调用，无需父组件处理
  const getPythonCodeFunc = async () => {
    const res = await getPythonCodeFuncList({ page: 1, pageSize: 1000 })
    if (res.code === 0) {
// 处理空参数，避免显示异常
      pythonFuncList.value = res.data.list.map(func => ({
        ...func,
        params: func.params && Array.isArray(func.params) ? func.params : []
      }))
    } else {
      ElMessage({ type: 'error', message: '获取Python函数列表失败' })
    }
  }

  // 表格展开控制
  const expandedRowKeys = ref([])
  const handleRowClick = (row) => {
    const isExpanded = expandedRowKeys.value.includes(row.ID)
    expandedRowKeys.value = isExpanded
      ? expandedRowKeys.value.filter(id => id !== row.ID)
      : [row.ID] // 单次只展开一行，如需多开可改为 push
  }

  // 复制函数格式
  const copyFuncName = (func) => {
    const paramStr = func.params.map(param => `$${param}`).join(', ')
    const funcText = `\${${func.name}(${paramStr})}`

    navigator.clipboard.writeText(funcText)
      .then(() => ElMessage({ type: 'success', message: '函数格式已复制' }))
      .catch(() => ElMessage({ type: 'error', message: '复制失败，请手动复制' }))
  }

  onMounted(() => {
    getPythonCodeFunc();
  });

</script>

<style scoped>
  /* 组件内部样式隔离 */
  pre {
    white-space: pre-wrap;
    word-break: break-all;
  }
</style>