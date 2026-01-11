<template>
  <el-dialog
    v-model="dialogVisible"
    title="历史变更记录"
    width="90%"
    top="5vh"
    @open="handleDialogOpen"
    class="custom-dialog"
    append-to-body
  >
    <div class="history-container">
      <!-- 左侧历史记录列表 -->
      <div class="history-list">
        <div
          v-for="item in pythonCodeList"
          :key="item.ID"
          class="history-item"
          :class="{ active: selectedHistory?.ID === item.ID }"
          @click="selectHistory(item)"
        >
          <div class="user-info">
            <div class="nickname">{{ item.updateByNickname }}</div>
            <div class="time">{{ formatTime(item.UpdatedAt) }}</div>
          </div>
          <el-icon v-if="selectedHistory?.ID === item.ID" class="check-icon"><Check /></el-icon>
        </div>
      </div>

      <!-- 右侧代码显示区域 -->
      <div class="code-preview">
        <div v-if="selectedHistory" class="preview-content">
          <div class="preview-header">
            <div class="preview-user">
              <el-icon><User /></el-icon>
              {{ selectedHistory.updateByNickname }}
              <span class="preview-time">{{ formatTime(selectedHistory.UpdatedAt) }}</span>
            </div>
            <el-button type="primary" size="small" @click="handleRestore">恢复此版本</el-button>
          </div>
          <div class="preview-code" style="padding: 0; overflow: hidden;">
            <CodeDiff
              :original="selectedHistory.code"
              :modified="currentCode"
              language="python"
            />
          </div>
        </div>
        <div v-else class="empty-preview">
          <el-empty description="请选择一条历史记录查看代码" />
        </div>
      </div>
    </div>
  </el-dialog>
</template>

<script setup>
import { ref, computed } from 'vue'
import { findPythonCode, getPythonCodeList } from '@/api/platform/pc.js'
import { getUserList } from '@/api/user.js'
import { ElMessage, ElMessageBox } from 'element-plus'
import CodeDiff from '@/components/platform/codeDiff/index.vue'

const props = defineProps({
  // 对话框显示控制
  visible: {
    type: Boolean,
    default: false
  },
  // 业务唯一标识，可以为空
  uniqueKey: {
    type: String,
    default: ''
  },
  // 项目 ID
  projectId: {
    type: [String, Number],
    default: null
  },
  // 历史记录类型（当前 pc 页面传入 pythonCodeType）
  type: {
    type: [Number, String],
    required: true
  },
  // 当前最新代码内容，用于做 diff
  currentCode: {
    type: String,
    default: ''
  }
})

const emit = defineEmits(['update:visible', 'restore'])

const dialogVisible = computed({
  get: () => props.visible,
  set: (val) => emit('update:visible', val)
})

const pythonCodeList = ref([])
const selectedHistory = ref(null)
const userJSON = ref({})
const userList = ref([])

const buildBaseParams = () => {
  const params = {
    type: props.type
  }
  if (props.uniqueKey) {
    params.uniqueKey = props.uniqueKey
  }
  if (props.projectId) {
    params.projectId = props.projectId
  }
  return params
}

const getUserListFunc = async () => {
  try {
    const res = await getUserList({
      page: 1,
      pageSize: 99999999
    })
    if (res.code === 0) {
      userList.value = res.data.list
      userList.value.forEach(user => {
        user.nickName = `${user.nickName}（${user.userName}）`
      })

      userJSON.value = userList.value.reduce((acc, user) => {
        acc[user.ID] = user.nickName
        return acc
      }, {})
    }
  } catch (error) {
    console.error('获取用户列表失败:', error)
    ElMessage.error('获取用户列表失败')
  }
}

const updateCodeForAllItems = async () => {
  if (!pythonCodeList.value || pythonCodeList.value.length === 0) {
    return
  }

  try {
    const promises = pythonCodeList.value.map(async (item) => {
      try {
        if (item.updateBy) {
          item.updateByNickname = userJSON.value[item.updateBy] || '--'
        } else {
          item.updateByNickname = '--'
        }
      } catch (error) {
        console.error(`处理 ID 为 ${item.ID} 的历史代码失败:`, error)
        item.code = ''
        item.updateByNickname = '--'
      }
    })

    await Promise.all(promises)
  } catch (error) {
    console.error('更新代码列表时发生错误:', error)
    ElMessage.error('更新代码列表失败')
  }
}

const getPythonCodeListFunc = async () => {
  try {
    const res = await getPythonCodeList({
      ...buildBaseParams()
    })
    if (res.code === 0) {
      pythonCodeList.value = res.data.list
      await updateCodeForAllItems()
      if (pythonCodeList.value.length > 0) {
        // 默认选中第一条历史记录
        await selectHistory(pythonCodeList.value[0])
      }
    }
  } catch (error) {
    console.error('获取历史代码列表失败:', error)
    ElMessage.error('获取历史代码列表失败')
  }
}

// 选择历史记录
const selectHistory = async (item) => {
  const res = await findPythonCode({
    id: item.ID,
    ...buildBaseParams()
  })
  if (res.code === 0) {
    item.code = res.data.code
    selectedHistory.value = item
  }
}

// 弹窗打开时的处理
const handleDialogOpen = async () => {
  await getUserListFunc()
  await getPythonCodeListFunc()
}

// 格式化时间显示
const formatTime = (timeString) => {
  if (!timeString) return '--'
  try {
    const date = new Date(timeString)
    return date.toLocaleString('zh-CN', {
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
      hour: '2-digit',
      minute: '2-digit',
      second: '2-digit'
    })
  } catch (error) {
    console.error('时间格式化错误:', error)
    return timeString
  }
}

// 恢复版本
const handleRestore = () => {
  if (!selectedHistory.value || !selectedHistory.value.code) return

  ElMessageBox.confirm('确定要恢复到此版本吗？当前未保存的内容将丢失。', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    emit('restore', selectedHistory.value.code)
    dialogVisible.value = false
    ElMessage.success('已恢复到历史版本')
  })
}
</script>

<style scoped lang="scss">
.history-container {
  display: flex;
  height: 80vh;
  border: 1px solid #dcdfe6;
  border-radius: 4px;
  overflow: hidden;
}

.history-list {
  width: 280px;
  border-right: 1px solid #dcdfe6;
  background-color: #f5f7fa;
  overflow-y: auto;

  .history-item {
    padding: 15px;
    cursor: pointer;
    border-bottom: 1px solid #ebeef5;
    transition: all 0.2s;
    display: flex;
    justify-content: space-between;
    align-items: center;

    &:hover {
      background-color: #e6f1fc;
    }

    &.active {
      background-color: #ecf5ff;
      border-left: 3px solid #409eff;

      .nickname {
        color: #409eff;
        font-weight: 600;
      }
    }

    .user-info {
      flex: 1;
    }

    .nickname {
      font-size: 14px;
      color: #303133;
      margin-bottom: 4px;
    }

    .time {
      font-size: 12px;
      color: #909399;
    }

    .check-icon {
      color: #409eff;
    }
  }
}

.code-preview {
  flex: 1;
  background-color: #fafafa;
  display: flex;
  flex-direction: column;
  overflow: hidden;

  .preview-content {
    display: flex;
    flex-direction: column;
    height: 100%;
  }

  .preview-header {
    padding: 10px 20px;
    background-color: #fff;
    border-bottom: 1px solid #ebeef5;
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .preview-user {
    display: flex;
    align-items: center;
    gap: 8px;
    font-weight: 600;
    color: #303133;
  }

  .preview-time {
    font-weight: normal;
    color: #909399;
    font-size: 12px;
    margin-left: 10px;
  }

  .preview-code {
    flex: 1;
    padding: 20px;
    overflow: auto;
    background-color: #282c34;
    color: #abb2bf;
    font-family: 'Monaco', 'Menlo', 'Consolas', monospace;
    font-size: 13px;
    line-height: 1.5;

    pre {
      margin: 0;
      white-space: pre-wrap;
      word-wrap: break-word;
    }
  }

  .empty-preview {
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: center;
    color: #909399;
  }
}
</style>
