<template>
  <div class="python-editor-container" :style="{ height: height }">
    <!-- 编辑器区域 -->
    <div class="editor-section" :class="{ 'with-function-list': showFunctionList }">
      <div class="editor-header" v-if="showHeader">
        <div class="header-title">{{ title }}</div>
        <div class="header-actions">
          <slot name="header-actions"></slot>
          <el-button
            v-if="showHistoryButton"
            type="info"
            link
            icon="Clock"
            @click="toggleHistory"
          >
            历史记录
          </el-button>
        </div>
      </div>
      <div class="editor-wrapper">
        <div ref="codeContainer" class="monaco-editor-container"></div>
      </div>
    </div>

    <!-- 函数列表区域 -->
    <div class="function-list-section" v-if="showFunctionList">
      <div class="function-header">
        <div class="header-title">函数列表</div>
        <div class="header-actions">
          <el-button
            type="primary"
            link
            icon="Refresh"
            @click="refreshFunctionList"
            :loading="refreshing"
          >
            刷新
          </el-button>
        </div>
      </div>
      <div class="function-list-container">
        <div class="function-list">
          <div
            v-for="(func, index) in functionList"
            :key="index"
            class="function-item"
            @click="handleFunctionClick(func)"
          >
            <div class="function-info">
              <div class="function-name">{{ func.name }}</div>
              <div class="function-signature" :title="formatFunctionSignature(func)">
                {{ formatFunctionSignature(func) }}
              </div>
            </div>
            <slot name="function-action" :func="func">
              <el-button
                v-if="showDebugButton"
                type="primary"
                size="small"
                icon="VideoPlay"
                @click.stop="handleDebug(func)"
                class="debug-btn"
                plain
              >
                调试
              </el-button>
            </slot>
          </div>

          <div v-if="functionList.length === 0" class="empty-function">
            <el-empty description="未检测到函数定义" :image-size="60" />
          </div>
        </div>
      </div>
    </div>
    <!-- 历史变更弹窗 -->
    <HistoryDialog
      v-model:visible="showHistoryDialog"
      :type="historyType"
      :unique-key="uniqueKey"
      :project-id="projectId"
      :current-code="currentCode"
      @restore="handleRestore"
    />
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, watch, toRaw, nextTick } from 'vue'
import * as monaco from 'monaco-editor'
import HistoryDialog from '@/components/platform/history/index.vue'

defineOptions({
  name: 'PythonCodeEditor'
})

// Props 定义
const props = defineProps({
  // v-model 绑定的代码内容
  modelValue: {
    type: String,
    default: ''
  },
  // 编辑器高度
  height: {
    type: String,
    default: '500px'
  },
  // 标题
  title: {
    type: String,
    default: '代码编辑器'
  },
  // 是否显示头部
  showHeader: {
    type: Boolean,
    default: true
  },
  // 是否显示函数列表
  showFunctionList: {
    type: Boolean,
    default: false
  },
  // 是否显示调试按钮
  showDebugButton: {
    type: Boolean,
    default: true
  },
  // 是否显示历史记录按钮
  showHistoryButton: {
    type: Boolean,
    default: false
  },
  // 是否只读
  readOnly: {
    type: Boolean,
    default: false
  },
  // 编辑器主题
  theme: {
    type: String,
    default: 'vs-dark'
  },
  // 字体大小
  fontSize: {
    type: Number,
    default: 14
  },
  // 字体
  fontFamily: {
    type: String,
    default: 'inherit'
  },
  // 是否显示行号
  lineNumbers: {
    type: String,
    default: 'on'
  },
  // 是否显示小地图
  minimap: {
    type: Boolean,
    default: true
  },
  // 函数列表宽度
  functionListWidth: {
    type: String,
    default: '320px'
  },
  // 唯一标识
  uniqueKey: {
    type: String,
    default: ''
  },
  // 历史记录类型
  historyType: {
    type: [Number, String],
    default: 1
  },
  // 项目ID
  projectId: {
    type: [Number, String],
    default: 0
  }
})

// Emits 定义
const emit = defineEmits([
  'update:modelValue',
  'change',
  'function-click',
  'debug',
  'functions-change',
  'editor-ready',
  'save'
])

// 响应式引用
const codeContainer = ref(null)
const monacoEditor = ref(null)
const functionList = ref([])
const refreshing = ref(false)

/**
 * 解析 Python 函数
 */
const extractPythonFunctions = (code) => {
  if (!code || typeof code !== 'string') {
    return []
  }

  const functions = []
  // 优化正则以支持更多格式，包括跨行参数
  const functionRegex = /(?:async\s+)?def\s+([a-zA-Z_][a-zA-Z0-9_]*)\s*\(([\s\S]*?)\)\s*(?:->\s*[^:]+)?\s*:/g

  let match
  while ((match = functionRegex.exec(code)) !== null) {
    const functionName = match[1]
    const paramsString = match[2].trim()
    const params = []

    if (paramsString) {
      // 简单处理换行，将换行替换为空格
      const cleanParamsString = paramsString.replace(/\n/g, ' ').replace(/\s+/g, ' ')
      const paramSegments = splitParams(cleanParamsString)

      paramSegments.forEach(segment => {
        const param = segment.trim()
        if (!param) return

        if (param.startsWith('*')) {
          const starParamMatch = param.match(/^\*\*?[a-zA-Z_][a-zA-Z0-9_]*/)
          if (starParamMatch) {
            params.push(starParamMatch[0])
          }
          return
        }

        const paramName = param
          .split(':')[0]
          .split('=')[0]
          .trim()

        if (paramName && !['self', 'cls'].includes(paramName)) {
          params.push(paramName)
        }
      })
    }

    const fullCode = extractFunctionCode(code, match.index, functionName)
    functions.push({
      name: functionName,
      params: params,
      fullCode: fullCode,
      startIndex: match.index
    })
  }

  return functions
}

/**
 * 分割参数字符串
 */
const splitParams = (paramsString) => {
  const segments = []
  let currentSegment = ''
  let bracketDepth = 0

  for (const char of paramsString) {
    if (char === ',' && bracketDepth === 0) {
      segments.push(currentSegment.trim())
      currentSegment = ''
    } else {
      currentSegment += char
      if (['(', '[', '{'].includes(char)) bracketDepth++
      if ([')', ']', '}'].includes(char)) bracketDepth--
    }
  }

  if (currentSegment.trim()) {
    segments.push(currentSegment.trim())
  }

  return segments
}

/**
 * 提取函数完整代码
 */
const extractFunctionCode = (code, startIndex, functionName) => {
  if (startIndex < 0 || startIndex >= code.length) return ''

  const functionDefEnd = code.indexOf(':', startIndex)
  if (functionDefEnd === -1) return code.substring(startIndex)

  let currentIndex = functionDefEnd + 1
  let inString = false
  let stringChar = ''
  let lastChar = ''

  while (currentIndex < code.length) {
    const char = code[currentIndex]

    if (!inString && (char === '"' || char === "'")) {
      inString = true
      stringChar = char
    } else if (inString && char === stringChar && lastChar !== '\\') {
      inString = false
    }

    if (!inString) {
      if (char === '\n') {
        const nextLineStart = currentIndex + 1
        if (nextLineStart < code.length) {
          const nextLineMatch = code.substring(nextLineStart).match(/^(\s*)/)
          const nextLineIndent = nextLineMatch ? nextLineMatch[1].length : 0

          const functionLine = code.substring(startIndex, code.indexOf('\n', startIndex))
          const functionIndentMatch = functionLine.match(/^(\s*)/)
          const functionIndent = functionIndentMatch ? functionIndentMatch[1].length : 0

          if (nextLineIndent <= functionIndent) {
            const nextLineContent = code.substring(nextLineStart).split('\n')[0].trim()
            if (nextLineContent && !nextLineContent.startsWith('#')) {
              break
            }
          }
        }
      }
    }

    lastChar = char
    currentIndex++
  }

  return code.substring(startIndex, currentIndex).trim()
}

/**
 * 格式化函数签名
 */
const formatFunctionSignature = (func) => {
  if (!func) return ''
  const params = func.params.map(param => `$${param}`).join(', ')
  return `${func.name}(${params})`
}

/**
 * 刷新函数列表
 */
const refreshFunctionList = () => {
  if (!monacoEditor.value) {
    return
  }

  refreshing.value = true
  try {
    const currentCode = toRaw(monacoEditor.value).getValue()
    functionList.value = extractPythonFunctions(currentCode)
    emit('functions-change', functionList.value)
  } catch (error) {
    console.error('刷新函数列表失败:', error)
  } finally {
    refreshing.value = false
  }
}

/**
 * 处理函数点击
 */
const handleFunctionClick = (func) => {
  emit('function-click', func)
}

/**
 * 处理调试
 */
const handleDebug = (func) => {
  emit('debug', func)
}

/**
 * 获取编辑器内容
 */
const getValue = () => {
  if (!monacoEditor.value) return ''
  return toRaw(monacoEditor.value).getValue()
}

/**
 * 设置编辑器内容
 */
const setValue = (value) => {
  if (!monacoEditor.value) return
  toRaw(monacoEditor.value).setValue(value)
}

/**
 * 获取函数列表
 */
const getFunctionList = () => {
  return functionList.value
}

/**
 * 获取 Monaco 编辑器实例
 */
const getEditor = () => {
  return monacoEditor.value
}

/**
 * 初始化编辑器
 */
const initEditor = () => {
  if (codeContainer.value && !monacoEditor.value) {
    monacoEditor.value = monaco.editor.create(codeContainer.value, {
      value: props.modelValue,
      language: 'python',
      automaticLayout: true,
      theme: props.theme,
      minimap: {
        enabled: props.minimap
      },
      fontSize: props.fontSize,
      fontFamily: props.fontFamily,
      lineNumbers: props.lineNumbers,
      readOnly: props.readOnly,
      scrollBeyondLastLine: false,
      wordWrap: 'on',
      contextmenu: true,
      scrollbar: {
        vertical: 'visible',
        horizontal: 'visible'
      }
    })

    // 添加 Ctrl+S 保存快捷键
    monacoEditor.value.addCommand(monaco.KeyMod.CtrlCmd | monaco.KeyCode.KeyS, () => {
      const value = toRaw(monacoEditor.value).getValue()
      emit('save', value)
    })

    // 监听内容变化
    monacoEditor.value.onDidChangeModelContent(() => {
      const value = toRaw(monacoEditor.value).getValue()
      emit('update:modelValue', value)
      emit('change', value)
    })

    // 初始化函数列表
    if (props.showFunctionList) {
      nextTick(() => {
        refreshFunctionList()
      })
    }

    emit('editor-ready', monacoEditor.value)
  }
}

// 监听 modelValue 变化
watch(() => props.modelValue, (newVal) => {
  if (monacoEditor.value) {
    const currentValue = toRaw(monacoEditor.value).getValue()
    // 只有当内容真正改变时才设置，避免光标跳动
    if (newVal !== currentValue && (newVal || newVal === '')) {
      toRaw(monacoEditor.value).setValue(newVal)
      if (props.showFunctionList) {
        refreshFunctionList()
      }
    }
  }
})

// 监听只读属性变化
watch(() => props.readOnly, (newVal) => {
  if (monacoEditor.value) {
    toRaw(monacoEditor.value).updateOptions({ readOnly: newVal })
  }
})

// 监听主题变化
watch(() => props.theme, (newVal) => {
  if (monacoEditor.value) {
    monaco.editor.setTheme(newVal)
  }
})

onMounted(() => {
  initEditor()
})

onUnmounted(() => {
  if (monacoEditor.value) {
    toRaw(monacoEditor.value).dispose()
    monacoEditor.value = null
  }
})

// 历史记录相关
const showHistoryDialog = ref(false)
const currentCode = ref('')

const toggleHistory = () => {
  currentCode.value = getValue()
  showHistoryDialog.value = true
}

const handleRestore = (code) => {
  emit('update:modelValue', code)
  emit('change', code)
  if (monacoEditor.value) {
    toRaw(monacoEditor.value).setValue(code)
  }
}

// 暴露方法给父组件
defineExpose({
  getValue,
  setValue,
  getFunctionList,
  getEditor,
  refreshFunctionList
})
</script>

<style lang="scss" scoped>
.python-editor-container {
  display: flex;
  gap: 15px;
  background-color: #f5f7fa;
  border-radius: 8px;
  overflow: hidden;
  width: 100%;
}

.editor-section {
  flex: 1;
  min-width: 0;
  background: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.05);
  display: flex;
  flex-direction: column;
  overflow: hidden;
  position: relative;

  &.with-function-list {
    // 当有函数列表时的样式调整
  }
}

.editor-header, .function-header {
  padding: 12px 16px;
  border-bottom: 1px solid #ebeef5;
  background-color: #fafafa;
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-shrink: 0;

  .header-title {
    font-size: 14px;
    font-weight: 600;
    color: #303133;
    display: flex;
    align-items: center;

    &::before {
      content: '';
      display: block;
      width: 4px;
      height: 14px;
      background-color: #409eff;
      margin-right: 8px;
      border-radius: 2px;
    }
  }

  .header-actions {
    display: flex;
    gap: 8px;
  }
}

.editor-wrapper {
  flex: 1;
  position: relative;
  overflow: hidden;
}

.monaco-editor-container {
  width: 100%;
  height: 100%;
  text-align: left;

  :deep(.monaco-editor) {
    text-align: left;
    
    .view-lines {
      text-align: left;
    }
    
    .monaco-scrollable-element {
      width: 100% !important;
    }
  }
}

.function-list-section {
  width: v-bind(functionListWidth);
  flex-shrink: 0;
  background: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.05);
  display: flex;
  flex-direction: column;
  overflow: hidden;
  transition: width 0.3s ease;
}

.function-list-container {
  flex: 1;
  overflow: hidden;
  padding: 10px;
  background-color: #fff;
}

.function-list {
  height: 100%;
  overflow-y: auto;
  padding-right: 5px;

  &::-webkit-scrollbar {
    width: 6px;
  }
  &::-webkit-scrollbar-thumb {
    background-color: #dcdfe6;
    border-radius: 3px;
  }
}

.function-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px 12px;
  margin-bottom: 8px;
  background: #fff;
  border: 1px solid #ebeef5;
  border-radius: 6px;
  transition: all 0.3s ease;
  box-shadow: 0 1px 4px rgba(0,0,0,0.02);
  cursor: pointer;

  &:hover {
    background: #ecf5ff;
    border-color: #c6e2ff;
    transform: translateY(-1px);
    box-shadow: 0 2px 8px rgba(0,0,0,0.05);
  }

  .function-info {
    flex: 1;
    min-width: 0;
    margin-right: 8px;
  }

  .function-name {
    font-size: 13px;
    font-weight: 600;
    color: #303133;
    margin-bottom: 2px;
  }

  .function-signature {
    font-family: 'Monaco', 'Menlo', 'Consolas', monospace;
    font-size: 11px;
    color: #909399;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  .debug-btn {
    flex-shrink: 0;
    opacity: 0.8;
    &:hover {
      opacity: 1;
    }
  }
}

.empty-function {
  padding: 30px 0;
  display: flex;
  justify-content: center;
}

/* History Dialog Styles */
/* 历史记录相关样式已移除，使用外部组件 */
</style>
