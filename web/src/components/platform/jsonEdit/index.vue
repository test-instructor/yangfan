<template>
  <div>
    <div
      :id="props.jsonType"
      :style="contentStyleObj"
      @keyup="jsonDatas"
      :height="height"
    ></div>
  </div>
</template>
<script>
  export default {
    name: 'JsonEditor'
  }
</script>
<script setup>
  import 'jsoneditor/dist/jsoneditor.min.css'
  import jsoneditor from 'jsoneditor'
  import { ref, onMounted, onBeforeMount, computed, onUnmounted,watch } from 'vue'
  import ace from 'ace-builds'
  import 'ace-builds/src-noconflict/mode-json'
  import 'ace-builds/src-noconflict/theme-github'
  import 'ace-builds/src-noconflict/theme-twilight'
  import 'ace-builds/src-noconflict/ext-language_tools'

  // 定义组件props
  const props = defineProps({
    request: {
      required: false
    },
    heights: {
      type: Number,
      required: true
    },
    jsons: {
      type: [Object, String],
      default: ''
    },
    jsonType: {
      type: String,
      default: 'codeEditorJson'
    }
  })

  // 定义组件 emits，添加验证状态
  const emit = defineEmits(['jsonData'])

  // 响应式变量
  const codeEditor = ref(null)
  const contentStyleObj = ref({
    height: '',
    width: ''
  })
  const timeStamp = ref('')

  // 计算属性
  const height = computed(() => {
    return props.heights - 50
  })

  // 方法定义
  const getHeight = () => {
    contentStyleObj.value.height = height.value + 'px'
    contentStyleObj.value.width = '98%'
  }

  // 验证JSON格式的方法
  const isValidJson = (data) => {
    try {
      // 如果是对象类型，尝试字符串化再解析来验证
      if (typeof data === 'object') {
        JSON.parse(JSON.stringify(data))
      } else {
        JSON.parse(data)
      }
      return true
    } catch (e) {
      console.error('JSON格式错误:', e)
      return false
    }
  }

  // 修改jsonDatas方法，包含格式验证结果
  const jsonDatas = () => {
    if (codeEditor.value) {
      try {
        const data = codeEditor.value.get()
        const isValid = isValidJson(data)

        // 传递一个包含数据和验证状态的对象
        emit('jsonData', {
          data: data,
          isValid: isValid,
          error: isValid ? null : new Error('Invalid JSON format')
        })
      } catch (error) {
        // 捕获编辑器获取数据时可能出现的错误
        emit('jsonData', {
          data: null,
          isValid: false,
          error: error
        })
      }
    }
  }

  // 监听主题变化
  const observer = ref(null)
  
  const updateTheme = () => {
    if (!codeEditor.value) return
    const isDark = document.documentElement.classList.contains('dark')
    console.log('JsonEditor updateTheme:', isDark, codeEditor.value.aceEditor)
    if (codeEditor.value.aceEditor) {
      codeEditor.value.aceEditor.setTheme(isDark ? 'ace/theme/twilight' : 'ace/theme/github')
    }
  }

  // 生命周期钩子
  onBeforeMount(() => {
    window.addEventListener('resize', getHeight)
    getHeight()
  })

  onMounted(() => {
    const codeOptions = {
      ace: ace,
      mode: 'code',
      modes: ['code', 'tree'],
      onModeChange: (newMode, oldMode) => {
        // 切换模式时重新应用主题
        if (newMode === 'code') {
          // 需要等待编辑器初始化完成
          setTimeout(updateTheme, 100)
        }
      }
    }
    const codeEditorElement = document.getElementById(props.jsonType)
    let json = {}

    if (props.jsons !== '') {
      json = typeof props.jsons === 'string' ? JSON.parse(props.jsons) : props.jsons
    }

    codeEditor.value = new jsoneditor(codeEditorElement, codeOptions, json)
    jsonDatas()
    
    // 初始化主题并监听变化
    updateTheme()
    observer.value = new MutationObserver(() => {
      updateTheme()
    })
    observer.value.observe(document.documentElement, { attributes: true, attributeFilter: ['class'] })
  })
  watch(
    () => props.jsons,
    (newVal, oldVal) => {
      if (newVal !== oldVal) {
        if (codeEditor.value) {
          codeEditor.value.update(newVal)
        }
      }
    },
    { deep: true }
  )
  // 清理工作
  onUnmounted(() => {
    window.removeEventListener('resize', getHeight)
    if (codeEditor.value) {
      codeEditor.value.destroy()
    }
    if (observer.value) {
      observer.value.disconnect()
    }
  })
</script>

<style>
  .ace_editor,
  .ace_editor * {
    font-family: "Monaco", "Menlo", "Ubuntu Mono", "Droid Sans Mono", "Consolas",
    monospace !important;
    font-size: 14px !important;
    font-weight: 400 !important;
    letter-spacing: 0 !important;
  }

  /* Dark mode overrides for jsoneditor */
  html.dark .jsoneditor {
    border-color: #4b5563;
    background-color: #111827;
  }
  html.dark .jsoneditor-menu {
    background-color: #1f2937;
    border-bottom: 1px solid #4b5563;
  }
  html.dark .jsoneditor-statusbar {
    background-color: #1f2937;
    border-top: 1px solid #4b5563;
    color: #9ca3af;
  }
  html.dark .jsoneditor-statusbar a {
    color: #60a5fa;
  }
  html.dark .jsoneditor-outer {
    background-color: #111827;
  }
  html.dark .jsoneditor-tree, 
  html.dark textarea.jsoneditor-text {
    background-color: #111827;
    color: #e5e7eb;
  }
  html.dark .jsoneditor-field,
  html.dark .jsoneditor-value {
    color: #e5e7eb !important;
  }
  /* Tree view value colors for dark mode */
  html.dark .jsoneditor-value.jsoneditor-string {
    color: #a5d6a7 !important;
  }
  html.dark .jsoneditor-value.jsoneditor-number {
    color: #ef9a9a !important;
  }
  html.dark .jsoneditor-value.jsoneditor-boolean {
    color: #ffcc80 !important;
  }
  html.dark .jsoneditor-value.jsoneditor-null {
    color: #90caf9 !important;
  }
  html.dark .jsoneditor-value.jsoneditor-object,
  html.dark .jsoneditor-value.jsoneditor-array {
    color: #d1d5db !important;
  }

  /* Context Menu */
  html.dark .jsoneditor-contextmenu .jsoneditor-menu {
    background-color: #1f2937 !important;
    border: 1px solid #4b5563 !important;
  }
  html.dark .jsoneditor-contextmenu .jsoneditor-menu button {
    color: #e5e7eb !important;
  }
  html.dark .jsoneditor-contextmenu .jsoneditor-menu button:hover,
  html.dark .jsoneditor-contextmenu .jsoneditor-menu button:focus {
    background-color: #374151 !important;
    color: #ffffff !important;
  }
  html.dark .jsoneditor-contextmenu .jsoneditor-menu li button.jsoneditor-selected {
    background-color: #374151 !important;
    color: #ffffff !important;
  }

  /* Search */
  html.dark .jsoneditor-search input {
    background-color: #374151 !important;
    color: #e5e7eb !important;
    border: 1px solid #4b5563 !important;
  }
  html.dark .jsoneditor-search .results {
    color: #9ca3af !important;
  }
  
  /* Icons */
  html.dark .jsoneditor-menu > button,
  html.dark .jsoneditor-tree button.jsoneditor-button {
    filter: invert(0.8) !important;
  }

  html.dark .jsoneditor-modes select {
    background-color: #1f2937 !important;
    color: #e5e7eb !important;
    border-color: #4b5563 !important;
  }

  /* ACE Editor Dark Mode Fixes - Removed aggressive overrides to let ace theme work */
  /* Only override background if theme fails or for specific container adjustments */
  html.dark .ace-jsoneditor.ace_editor {
    background-color: #141414;
    color: #f8f8f2;
  }
  
  /* Ensure gutter matches twilight theme */
  html.dark .ace-jsoneditor .ace_gutter {
    background-color: #232323;
    color: #E2E2E2;
  }
  
  /* Ensure correct font family for jsoneditor elements to match ACE */
  .jsoneditor-field,
  .jsoneditor-value, 
  textarea.jsoneditor-text {
    font-family: "Monaco", "Menlo", "Ubuntu Mono", "Droid Sans Mono", "Consolas", monospace !important;
  }
</style>
