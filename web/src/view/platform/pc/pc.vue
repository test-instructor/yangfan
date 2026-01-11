<template>
  <div class="pc-container">
    <div class="main-content">
      <!-- 使用 PythonCodeEditor 组件 -->
      <PythonCodeEditor
        ref="pythonEditorRef"
        v-model="pythonCode"
        height="calc(100vh - 120px)"
        title="代码编辑器"
        :show-function-list="true"
        :show-debug-button="true"
        :line-numbers="'off'"
        function-list-width="380px"
        @debug="showDebugDialog"
        @functions-change="handleFunctionsChange"
        @change="handleCodeChange"
      >
        <template #header-actions>
          <el-button type="primary" icon="DocumentChecked" @click="handleSave" :loading="saving">
            保存代码
          </el-button>
          <el-button icon="Clock" @click="showHistoryDialog = true">
            历史变更
          </el-button>
        </template>
      </PythonCodeEditor>
    </div>

    <HistoryDialog
      v-model:visible="showHistoryDialog"
      :type="pythonCodeType"
      :unique-key="uniqueKey"
      :project-id="projectId"
      :current-code="pythonCode"
    />

    <!-- 调试弹窗 -->
    <el-dialog
      v-model="showDebugDialogVisible"
      :title="`调试函数 - ${currentDebugFunction?.name || ''}`"
      width="700px"
      @close="handleDebugDialogClose"
      class="custom-dialog"
    >
      <div class="debug-dialog">
        <div class="section-title">函数签名</div>
        <div class="function-info">
          <el-input v-model="debugData.function" readonly>
            <template #prefix>
              <el-icon><Operation /></el-icon>
            </template>
          </el-input>
        </div>

        <div class="section-title">参数配置 (JSON)</div>
        <div class="parameters-section">
          <JsonEditor
            :key="JSON.stringify(debugData.parameters)"
            :heights="300"
            :jsons="debugData.parameters"
            @jsonData="handleJsonData"
          />
        </div>

        <div v-if="debugOptions.showCode && currentDebugFunction" class="function-code-preview">
          <div class="section-title">函数代码</div>
          <div class="code-content">
            <pre>{{ currentDebugFunction.fullCode || '无法获取函数完整代码' }}</pre>
          </div>
        </div>

        <div class="debug-result">
          <div class="section-title">调试结果</div>
          <div class="result-content">
            <pre>{{ debugResponse }}</pre>
          </div>
        </div>
      </div>

      <template #footer>
        <div class="dialog-footer">
          <el-button @click="showDebugDialogVisible = false">取消</el-button>
          <el-button
            type="primary"
            @click="handleDebugConfirm"
            :loading="debugging"
            icon="VideoPlay"
          >
            开始调试
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
  import { ref, onMounted, toRaw } from 'vue'
  import { findPythonCode, updatePythonCode } from '@/api/platform/pc.js'
  import { createPythonCodeDebug } from '@/api/platform/pythoncodedebug.js'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import JsonEditor from '@/components/platform/jsonEdit/index.vue'
  import { updatePythonCodeFunc } from '@/api/platform/pythoncodefunc.js'
  import PythonCodeEditor from '@/components/platform/pythonEditor/PythonCodeEditor.vue'
  import HistoryDialog from '@/components/platform/history/index.vue'
  import { useUserStore } from '@/pinia/modules/user'

  // 响应式引用
  const pythonEditorRef = ref(null)
  const pythonCode = ref('')
  const pythonCodeType = ref(1)
  const saving = ref(false)
  const showHistoryDialog = ref(false)

  const userStore = useUserStore()
  const projectId = userStore.userInfo?.projectId
  const uniqueKey = ''

  // 函数相关状态
  const functionList = ref([])
  const showDebugDialogVisible = ref(false)
  const currentDebugFunction = ref(null)
  const debugParameters = ref({})
  const debugging = ref(false)
  const debugOptions = ref({
    showCode: false,
    logExecution: true
  })
  const debugData = ref({})
  const debugType = ref(1)

  /**
   * 处理函数列表变化
   */
  const handleFunctionsChange = (functions) => {
    functionList.value = functions
  }

  /**
   * 处理代码变化
   */
  const handleCodeChange = (code) => {
    // 可以在这里添加额外的代码变化处理逻辑
  }

  /**
   * 格式化函数签名显示
   */
  const formatFunctionSignature = (func) => {
    if (!func) return ''
    const params = func.params.map(param => `$${param}`).join(', ')
    return `${func.name}(${params})`
  }

  /**
   * 显示调试弹窗
   */
  const showDebugDialog = (func) => {
    currentDebugFunction.value = { ...func }
    debugParameters.value = {}
    const parameters = ref({})
    // 初始化参数值为空字符串
    func.params.forEach(param => {
      parameters.value[param] = ''
    })

    debugData.value = {
      function: '${' + formatFunctionSignature(func) + '}',
      parameters: parameters.value,
      type: debugType.value
    }
    showDebugDialogVisible.value = true
  }

  /**
   * 处理调试确认
   */
  const debugResponse = ref()
  const handleDebugConfirm = async () => {
    if (!currentDebugFunction.value) return
    if (!isJsonValid.value) {
      ElMessage.error('json数据格式校验不通过')
      return
    }
    await handleSave()
    debugData.value.parameters = currentJson.value
    // debugging.value = true
    try {
      const res = await createPythonCodeDebug(debugData.value)
      if (res.code === 0) {
        debugResponse.value = res.data
      }
    } catch (error) {
      console.error('调试失败:', error)
      ElMessage.error('调试失败')
    } finally {
      // debugging.value = false
    }
  }

  /**
   * 处理调试弹窗关闭
   */
  const handleDebugDialogClose = () => {
    currentDebugFunction.value = null
    debugParameters.value = {}
    debugOptions.value = {
      showCode: false,
      logExecution: true
    }
    debugData.value = {
      function: '',
      parameters: {}
    }
  }

  onMounted(async () => {
    await getPythonCode()
  })

  const getPythonCode = async () => {
    try {
      const res = await findPythonCode({ type: pythonCodeType.value })
      if (res.code === 0) {
        pythonCode.value = res.data.code
      }
    } catch (error) {
      console.error('获取Python代码失败:', error)
      ElMessage.error('获取Python代码失败')
    }
  }

  // 保存代码
  const handleSave = async () => {
    if (!pythonEditorRef.value) {
      ElMessage.warning('编辑器未初始化')
      return
    }

    const currentCode = pythonEditorRef.value.getValue()

    // 刷新函数列表
    pythonEditorRef.value.refreshFunctionList()

    saving.value = true
    try {
      const res = await updatePythonCode({
        type: pythonCodeType.value,
        code: currentCode
      })

      if (res.code === 0) {
        ElMessage.success('代码保存成功')
        pythonCode.value = currentCode
        await updatePythonCodeFunc({ data: functionList.value })
      } else {
        ElMessage.error(res.msg || '代码保存失败')
      }
    } catch (error) {
      console.error('保存代码失败:', error)
      ElMessage.error('保存代码失败')
    } finally {
      saving.value = false
    }
  }

  // 存储编辑器当前的数据和状态
  const currentJson = ref({})
  const isJsonValid = ref(true)
  const errorMessage = ref('')
  const hasError = ref(false)
  const hasData = ref(false)

  // 处理编辑器返回的数据
  const handleJsonData = (result) => {
    hasData.value = true
    if (result.isValid) {
      // 格式正确
      isJsonValid.value = true
      hasError.value = false
      currentJson.value = result.data
    } else {

      // 格式错误，显示红色提示
      isJsonValid.value = false
      hasError.value = true
      errorMessage.value = result.error.message || '请检查JSON格式是否正确'
      currentJson.value = {}
    }
  }

</script>

<style lang="scss" scoped>
.pc-container {
  height: calc(100vh - 100px);
  padding: 10px;
  box-sizing: border-box;
  background-color: #f5f7fa;
}

.main-content {
  height: 100%;
}

/* Debug Dialog Styles */
.debug-dialog {
  padding: 0 10px;

  .section-title {
    font-size: 14px;
    font-weight: 600;
    color: #303133;
    margin-bottom: 10px;
    margin-top: 15px;
    display: flex;
    align-items: center;

    &::before {
      content: '';
      display: inline-block;
      width: 3px;
      height: 14px;
      background-color: #409eff;
      margin-right: 8px;
      border-radius: 2px;
    }

    &:first-child {
      margin-top: 0;
    }
  }

  .function-info {
    margin-bottom: 20px;
  }

  .parameters-section {
    margin-bottom: 20px;
    border: 1px solid #dcdfe6;
    border-radius: 4px;
    overflow: hidden;
  }

  .function-code-preview, .debug-result {
    margin-top: 20px;

    .code-content, .result-content {
      background-color: #f5f7fa;
      padding: 15px;
      border-radius: 4px;
      border: 1px solid #ebeef5;
      max-height: 200px;
      overflow: auto;
      font-family: 'Monaco', 'Menlo', 'Consolas', monospace;
      font-size: 12px;
      color: #606266;

      pre {
        margin: 0;
        white-space: pre-wrap;
      }
    }

    .result-content {
      background-color: #ecf5ff;
      color: #333;
    }
  }
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
}
</style>

