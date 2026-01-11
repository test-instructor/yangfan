<template>
  <div ref="container" class="diff-editor-container"></div>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount, watch } from 'vue'
import * as monaco from 'monaco-editor'

const props = defineProps({
  original: {
    type: String,
    default: ''
  },
  modified: {
    type: String,
    default: ''
  },
  language: {
    type: String,
    default: 'python'
  },
  fontFamily: {
    type: String,
    default: 'inherit'
  }
})

const container = ref(null)
let diffEditor = null

onMounted(() => {
  if (container.value) {
    diffEditor = monaco.editor.createDiffEditor(container.value, {
      originalEditable: false, // 左侧不可编辑
      readOnly: true,         // 右侧也不可编辑（只是查看对比）
      theme: 'vs-dark',
      automaticLayout: true,
      renderSideBySide: true,
      fontFamily: props.fontFamily,
      minimap: {
        enabled: true
      }
    })
    
    updateModels()
  }
})

const updateModels = () => {
  if (!diffEditor) return

  // 清理旧模型
  const oldModel = diffEditor.getModel()
  if (oldModel) {
    if (oldModel.original) oldModel.original.dispose()
    if (oldModel.modified) oldModel.modified.dispose()
  }

  // 创建新模型
  // 左侧：历史记录 (original)
  // 右侧：最新记录 (modified)
  const originalModel = monaco.editor.createModel(props.original, props.language)
  const modifiedModel = monaco.editor.createModel(props.modified, props.language)

  diffEditor.setModel({
    original: originalModel,
    modified: modifiedModel
  })
}

watch(() => [props.original, props.modified, props.language], () => {
  updateModels()
})

onBeforeUnmount(() => {
  if (diffEditor) {
    // 清理模型
    const model = diffEditor.getModel()
    if (model) {
      if (model.original) model.original.dispose()
      if (model.modified) model.modified.dispose()
    }
    diffEditor.dispose()
  }
})
</script>

<style scoped>
.diff-editor-container {
  width: 100%;
  height: 100%;
  text-align: left;

  :deep(.monaco-editor) {
    text-align: left;
    
    .view-lines {
      text-align: left;
    }
  }
}
</style>
