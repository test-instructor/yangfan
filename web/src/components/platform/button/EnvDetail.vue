<template>
  <div>
    <el-button type="primary" @click="dialogVisible = true" style="margin-left: 15px;">
      环境变量
      <template v-if="isLoading">
        <el-icon style="margin-left: 5px;">
          <loading />
        </el-icon>
      </template>
    </el-button>

    <el-dialog
      title="环境变量列表"
      v-model="dialogVisible"
      width="800px"
    >
      <el-table
        :data="envList"
        border
        max-height="500px"
      >
        <el-table-column
          prop="key"
          label="变量名称"
          width="340"
        />
        <el-table-column
          prop="name"
          label="中文名称"
          width="340"
        />
        <el-table-column label="操作">
          <template #default="scope">
            <el-button
              type="text"
              @click="copyVariable(scope.row.key)"
            >
              复制
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-dialog>
  </div>
</template>

<script setup>
  import { ref, onMounted } from 'vue'
  import { getEnvDetailList } from '@/api/platform/envdetail.js'
  import { ElMessage, ElLoading } from 'element-plus'
  import { Loading } from '@element-plus/icons-vue'

  const envList = ref([])
  const dialogVisible = ref(false)
  const isLoading = ref(false) // 加载状态标识

  // 获取环境变量列表
  const getEnvList = async () => {
    isLoading.value = true
    try {
      const res = await getEnvDetailList({ page: 1, pageSize: 1000 })
      if (res.code === 0) {
        envList.value = res.data.list || []
      } else {
        ElMessage.error(`获取失败: ${res.msg || '未知错误'}`)
      }
    } catch (error) {
      console.error('接口请求异常:', error)
      ElMessage.error('网络异常，获取数据失败')
    } finally {
      isLoading.value = false
    }
  }

  // 组件挂载后立即执行 + 额外的安全调用
  onMounted(() => {
    getEnvList()
  })

  // 复制功能
  const copyVariable = (key) => {
    const textToCopy = `$${key}`
    navigator.clipboard.writeText(textToCopy)
      .then(() => ElMessage.success(`已复制: ${textToCopy}`))
      .catch(() => {
        const textarea = document.createElement('textarea')
        textarea.value = textToCopy
        document.body.appendChild(textarea)
        textarea.select()
        document.execCommand('copy')
        document.body.removeChild(textarea)
        ElMessage.success(`已复制: ${textToCopy}`)
      })
  }
</script>