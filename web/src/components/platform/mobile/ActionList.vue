<template>
  <div class="mobile-action-list">
    <div class="toolbar">
      <el-button type="primary" icon="Plus" @click="addAction">添加动作</el-button>
      <el-button type="danger" icon="Delete" :disabled="selectedIndices.length === 0" @click="deleteSelected">删除选中</el-button>
      <el-button icon="Sort" @click="toggleSort">排序模式: {{ isSortMode ? '开启' : '关闭' }}</el-button>
    </div>

    <el-table
      :data="localActions"
      style="width: 100%"
      border
      row-key="id"
      @selection-change="handleSelectionChange"
      class="action-table"
    >
      <el-table-column type="selection" width="55" />
      <el-table-column label="序号" type="index" width="60" align="center" />
      
      <el-table-column label="动作类型" width="180">
        <template #default="{ row }">
          <el-select v-model="row.method" placeholder="选择动作" filterable @change="handleMethodChange(row)">
            <el-option-group label="常用动作">
              <el-option label="点击 (TapXY)" value="tap_xy" />
              <el-option label="输入 (Input)" value="input" />
              <el-option label="等待 (Sleep)" value="sleep" />
              <el-option label="截图 (ScreenShot)" value="screenshot" />
            </el-option-group>
            <el-option-group label="滑动操作">
              <el-option label="滑动 (Swipe)" value="swipe" />
              <el-option label="向上滑动 (SwipeUp)" value="swipe_up" />
              <el-option label="向下滑动 (SwipeDown)" value="swipe_down" />
              <el-option label="向左滑动 (SwipeLeft)" value="swipe_left" />
              <el-option label="向右滑动 (SwipeRight)" value="swipe_right" />
            </el-option-group>
            <el-option-group label="应用操作">
              <el-option label="启动应用 (AppLaunch)" value="app_launch" />
              <el-option label="终止应用 (AppTerminate)" value="app_terminate" />
              <el-option label="安装应用 (AppInstall)" value="install_app" />
              <el-option label="卸载应用 (AppUninstall)" value="uninstall_app" />
              <el-option label="回到桌面 (Home)" value="home" />
            </el-option-group>
            <el-option-group label="断言与AI">
              <el-option label="OCR点击 (TapByOCR)" value="tap_ocr" />
              <el-option label="图像点击 (TapByCV)" value="tap_cv" />
              <el-option label="AI操作 (AIAction)" value="ai_action" />
              <el-option label="AI断言 (AIAssert)" value="ai_assert" />
            </el-option-group>
          </el-select>
        </template>
      </el-table-column>

      <el-table-column label="参数配置" min-width="300">
        <template #default="{ row }">
          <div class="params-container">
            <!-- 动态渲染参数输入框 -->
            <template v-if="['tap_xy', 'tap_abs_xy', 'double_tap_xy'].includes(row.method)">
              <el-input-number v-model="row.options.x" :precision="2" :step="0.1" placeholder="X" controls-position="right" style="width: 100px" />
              <el-input-number v-model="row.options.y" :precision="2" :step="0.1" placeholder="Y" controls-position="right" style="width: 100px" />
            </template>
            
            <template v-else-if="['input', 'tap_ocr', 'tap_cv', 'app_launch', 'app_terminate', 'install_app', 'uninstall_app'].includes(row.method)">
              <el-input v-model="row.params" placeholder="请输入内容/包名/路径" style="flex: 1" />
            </template>

            <template v-else-if="['sleep'].includes(row.method)">
              <el-input-number v-model="row.params" :min="0" placeholder="秒" />
              <span class="unit">秒</span>
            </template>

            <template v-else-if="['swipe'].includes(row.method)">
              <el-input-number v-model="row.options.from_x" :precision="2" placeholder="From X" controls-position="right" style="width: 120px" />
              <el-input-number v-model="row.options.from_y" :precision="2" placeholder="From Y" controls-position="right" style="width: 120px" />
              <span class="arrow">→</span>
              <el-input-number v-model="row.options.to_x" :precision="2" placeholder="To X" controls-position="right" style="width: 120px" />
              <el-input-number v-model="row.options.to_y" :precision="2" placeholder="To Y" controls-position="right" style="width: 120px" />
            </template>
            
             <template v-else-if="['ai_action', 'ai_assert'].includes(row.method)">
              <el-input v-model="row.params" type="textarea" :rows="1" placeholder="请输入AI提示词/断言描述" style="flex: 1" />
            </template>
            
            <!-- 通用选项按钮 -->
            <el-popover placement="bottom" title="更多选项" :width="300" trigger="click">
              <template #reference>
                <el-button icon="Setting" circle size="small" class="setting-btn" />
              </template>
              <el-form label-width="80px" size="small">
                <el-form-item label="标识符">
                  <el-input v-model="row.options.identifier" placeholder="用于日志/报告" />
                </el-form-item>
                <el-form-item label="最大重试">
                  <el-input-number v-model="row.options.maxRetryTimes" :min="0" />
                </el-form-item>
                <el-form-item label="忽略错误">
                   <el-switch v-model="row.options.ignoreNotFoundError" />
                </el-form-item>
                <el-form-item label="超时(s)">
                   <el-input-number v-model="row.options.timeout" :min="0" />
                </el-form-item>
                 <el-form-item label="文本">
                   <el-input v-model="row.options.text" placeholder="辅助文本" />
                </el-form-item>
              </el-form>
            </el-popover>
          </div>
        </template>
      </el-table-column>

      <el-table-column label="操作" width="100" align="center">
        <template #default="{ $index }">
          <el-button type="danger" icon="Delete" circle size="small" @click="deleteAction($index)" />
           <div class="drag-handle" v-if="isSortMode" style="cursor: move; display: inline-block; margin-left: 5px;">
              <el-icon><Rank /></el-icon>
           </div>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script setup>
import { ref, watch, onMounted } from 'vue'
import { Sortable } from 'sortablejs'

const props = defineProps({
  modelValue: {
    type: Array,
    default: () => []
  }
})

const emit = defineEmits(['update:modelValue'])

const localActions = ref([])
const selectedIndices = ref([])
const isSortMode = ref(false)
let syncingFromParent = false

const safeStringify = (val) => {
  try {
    return JSON.stringify(val ?? null)
  } catch (e) {
    return ''
  }
}

// 初始化
watch(() => props.modelValue, (newVal) => {
  // 简单深拷贝，避免直接修改 props，并确保每个 action 有唯一 id 用于 key
  syncingFromParent = true
  localActions.value = (newVal || []).map((item, index) => ({
    ...item,
    id: item.id || Date.now() + index, // 确保有唯一 ID
    options: item.options || {}
  }))
  Promise.resolve().then(() => {
    syncingFromParent = false
  })
}, { immediate: true, deep: true })

// 监听本地变化同步回父组件
watch(localActions, (newVal) => {
  if (syncingFromParent) return
  // 清理临时 ID
  const cleanData = newVal.map(({ id, ...rest }) => rest)
  if (safeStringify(cleanData) === safeStringify(props.modelValue || [])) return
  emit('update:modelValue', cleanData)
}, { deep: true })

const addAction = () => {
  localActions.value.push({
    id: Date.now(),
    method: 'tap_xy',
    params: null,
    options: {
        x: 0.5,
        y: 0.5
    }
  })
}

const deleteAction = (index) => {
  localActions.value.splice(index, 1)
}

const handleSelectionChange = (selection) => {
  // 记录选中项的 ID
  const ids = selection.map(item => item.id)
  selectedIndices.value = ids
}

const deleteSelected = () => {
  localActions.value = localActions.value.filter(item => !selectedIndices.value.includes(item.id))
  selectedIndices.value = []
}

const handleMethodChange = (row) => {
    // 重置 params 和 options 的默认值
    row.params = null
    row.options = {}
    
    if (['tap_xy', 'tap_abs_xy'].includes(row.method)) {
        row.options.x = 0.5
        row.options.y = 0.5
    } else if (row.method === 'sleep') {
        row.params = 1
    }
}

const toggleSort = () => {
    isSortMode.value = !isSortMode.value
    initSortable()
}

let sortableInstance = null
const initSortable = () => {
    const el = document.querySelector('.action-table .el-table__body-wrapper tbody')
    if (!el) return

    if (isSortMode.value) {
        if (!sortableInstance) {
            sortableInstance = Sortable.create(el, {
                handle: '.drag-handle',
                animation: 150,
                onEnd: ({ newIndex, oldIndex }) => {
                    const targetRow = localActions.value.splice(oldIndex, 1)[0]
                    localActions.value.splice(newIndex, 0, targetRow)
                }
            })
        }
    } else {
        if (sortableInstance) {
            sortableInstance.destroy()
            sortableInstance = null
        }
    }
}

</script>

<style scoped>
.mobile-action-list {
  padding: 20px 10px 10px 10px;
}
.toolbar {
  margin-bottom: 10px;
  display: flex;
  gap: 10px;
}
.params-container {
  display: flex;
  align-items: center;
  gap: 8px;
  width: 100%;
}
.unit {
  color: #666;
  font-size: 12px;
}
.arrow {
  color: #999;
  font-weight: bold;
}
.setting-btn {
  margin-left: auto;
}
</style>
