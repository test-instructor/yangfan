<template>
  <el-card class="box-card mt-4" shadow="hover">
    <template #header>
      <div class="flex justify-between items-center">
        <span class="text-lg font-bold">层级列表</span>
        <div class="flex gap-2">
          <el-input 
            v-model="searchText" 
            placeholder="搜索用例/接口名称..." 
            style="width: 200px;" 
            prefix-icon="Search" 
            clearable
          />
          <el-button type="primary" @click="$emit('refresh')" icon="Refresh">刷新</el-button>
        </div>
      </div>
    </template>
    
    <el-table
      :data="filteredDetails"
      row-key="ID"
      style="width: 100%"
      :expand-row-keys="expandedRowKeys"
      @expand-change="handleExpandChange"
      @row-click="handleRowClick"
    >
      <el-table-column type="expand">
        <template #default="props">
          <div class="p-4 bg-gray-50">
            <el-collapse>
              <el-collapse-item
                v-for="group in props.row.groupedRecords"
                :key="group.stepName"
                :name="group.stepName"
                :title="group.stepName"
              >
                <el-table :data="group.records" style="width: 100%" size="small" border>
                  <el-table-column prop="index" label="序号" width="60">
                    <template #default="scope">{{ scope.$index + 1 }}</template>
                  </el-table-column>

                  <el-table-column label="接口名称" min-width="260">
                    <template #default="scope">
                      <div class="flex flex-col">
                        <span class="truncate" :title="getRequestUrl(scope.row)">
                          {{ getRequestUrl(scope.row) }}
                        </span>
                        <span
                          v-if="scope.row._stepDesc"
                          class="text-xs text-gray-500 mt-0.5 truncate"
                          :title="scope.row._stepDesc"
                        >
                          {{ scope.row._stepDesc }}
                        </span>
                      </div>
                    </template>
                  </el-table-column>

                  <el-table-column label="类型" width="100">
                    <template #default="scope">
                      <el-tag size="small">
                        {{ getRequestMethod(scope.row) }}
                      </el-tag>
                    </template>
                  </el-table-column>

                  <el-table-column prop="success" label="状态" width="80">
                    <template #default="scope">
                      <el-tag :type="scope.row.success ? 'success' : 'danger'" size="small">
                        {{ scope.row.success ? '成功' : '失败' }}
                      </el-tag>
                    </template>
                  </el-table-column>

                  <el-table-column prop="elapsed_ms" label="耗时(ms)" width="100" />

                  <el-table-column label="操作" width="100">
                    <template #default="scope">
                      <el-button
                        v-if="['request', 'api'].includes(scope.row.step_type) || getRequestMethod(scope.row)"
                        type="primary"
                        link
                        size="small"
                        @click.stop="$emit('show-detail', scope.row)"
                      >
                        查看详情
                      </el-button>
                    </template>
                  </el-table-column>
                </el-table>
              </el-collapse-item>
            </el-collapse>
          </div>
        </template>
      </el-table-column>

      <el-table-column label="用例名称" prop="name" show-overflow-tooltip />
      <el-table-column label="状态" prop="success" width="100">
        <template #default="scope">
          <el-tag :type="scope.row.success ? 'success' : 'danger'">
            {{ scope.row.success ? '成功' : '失败' }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column label="执行时间" width="180">
        <template #default="scope">
          {{ formatDate(scope.row.time?.start_at) }}
        </template>
      </el-table-column>
      <el-table-column label="耗时(s)" width="100">
        <template #default="scope">
          {{ scope.row.time?.duration ? scope.row.time.duration.toFixed(2) : '-' }}
        </template>
      </el-table-column>
    </el-table>
  </el-card>
</template>

<script setup>
import { ref, computed } from 'vue'
import { formatDate } from '@/utils/format'

const props = defineProps({
  details: {
    type: Array,
    default: () => []
  }
})

const emit = defineEmits(['refresh', 'show-detail'])

const searchText = ref('')
const expandedRowKeys = ref([])

const filteredDetails = computed(() => {
  if (!props.details) return []
  let list = props.details
  if (searchText.value) {
    const lower = searchText.value.toLowerCase()
    list = list.filter(item => {
      const matchCase = item.name.toLowerCase().includes(lower)
      if (matchCase) return true
      if (item.records && item.records.some(r => r.name.toLowerCase().includes(lower))) {
        return true
      }
      return false
    })
  }
  return list
})

const getRequest = (record) => record?.data?.req_resps?.request || {}
const getRequestMethod = (record) => getRequest(record).method || ''
const getRequestUrl = (record) => getRequest(record).url || ''

const handleExpandChange = (row, expandedRows) => {
  expandedRowKeys.value = expandedRows.map(r => r.ID)
}

const handleRowClick = (row) => {
  if (!row.records || !row.records.length) return
  const idx = expandedRowKeys.value.indexOf(row.ID)
  if (idx > -1) {
    expandedRowKeys.value.splice(idx, 1)
  } else {
    expandedRowKeys.value.push(row.ID)
  }
}
</script>
