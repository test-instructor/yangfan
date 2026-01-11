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
  const editorInit = () => {
    require('brace/ext/language_tools')
    require('brace/mode/json')
    require('brace/theme/github')
    require('brace/snippets/json')
  }

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

  // 生命周期钩子
  onBeforeMount(() => {
    window.addEventListener('resize', getHeight)
    getHeight()
  })

  onMounted(() => {
    const codeOptions = {
      mode: 'code',
      modes: ['code', 'tree']
    }
    const codeEditorElement = document.getElementById(props.jsonType)
    let json = {}

    if (props.jsons !== '') {
      json = typeof props.jsons === 'string' ? JSON.parse(props.jsons) : props.jsons
    }

    codeEditor.value = new jsoneditor(codeEditorElement, codeOptions, json)
    jsonDatas()
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
</style>
