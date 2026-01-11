<template>
  <div style="display: flex; gap: 20px;">
    <!-- 测试前执行方法表格 -->
    <div style="flex: 1;">
      <el-table
        highlight-current-row
        :cell-style="{ paddingTop: '4px', paddingBottom: '4px' }"
        stripe
        :height="height"
        :data="setupData"
        style="width: 100%"
        @cell-mouse-enter="handleCellMouseEnter"
        @cell-mouse-leave="handleCellMouseLeave"
      >
        <el-table-column label="请求之前执行的方法" width="460">
          <template #default="scope">
            <el-input
              clearable
              v-model="scope.row.setup"
              placeholder="${ setup_hooks function($request, *args, **kwargs) }"
              @input="handleSetupHooksData"
            />
          </template>
        </el-table-column>
        <el-table-column width="120">
          <template #default="scope">
            <el-row :gutter="8">
              <el-button
                size="small"
                type="info"
                @click="handleSetupAdd"
                icon="add"
              />
              <el-button
                size="small"
                type="danger"
                v-show="setupData.length > 1"
                @click="handleSetupDelete(scope.$index)"
                icon="delete"
              />
            </el-row>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <!-- 测试后执行方法表格 -->
    <div style="flex: 1;">
      <el-table
        highlight-current-row
        :cell-style="{ paddingTop: '4px', paddingBottom: '4px' }"
        stripe
        :height="height"
        :data="teardownData"
        style="width: 100%"
        @cell-mouse-enter="handleCellMouseEnter"
        @cell-mouse-leave="handleCellMouseLeave"
      >
        <el-table-column label="测试之后执行的方法" width="460">
          <template #default="scope">
            <el-input
              clearable
              v-model="scope.row.teardown"
              placeholder="${ teardown_hooks function(response, *args, **kwargs) }"
              @input="handleTeardownHooksData"
            />
          </template>
        </el-table-column>
        <el-table-column width="120">
          <template #default="scope">
            <el-row :gutter="8">
              <el-button
                size="small"
                type="info"
                @click="handleTeardownAdd"
                icon="add"
              />
              <el-button
                size="small"
                type="danger"
                v-show="teardownData.length > 1"
                @click="handleTeardownDelete(scope.$index)"
                icon="delete"
              />
            </el-row>
          </template>
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>

<script setup>
  import { ref, computed, onMounted,watch } from 'vue'

  const props = defineProps({
    save: {
      type: Boolean,
      default: false
    },
    hooks: {
      type: Object,
      default: () => ({}) // 修复原代码require语法错误，补充默认值
    },
    heights: {
      type: Number,
      default: 400 // 避免无高度时表格异常
    },
    setupHooks: {
      type: Array,
      default: () => []
    },
    teardownHooks: {
      type: Array,
      default: () => []
    }
  })

  const emit = defineEmits(['setupHooksData', 'teardownHooksData'])

  const setupData = ref([{ setup: '' }]) // 测试前方法数据
  const teardownData = ref([{ teardown: '' }]) // 测试后方法数据
  const currentRow = ref('') // 鼠标悬浮行

  const height = computed(() => props.heights - 70)

  const initHooksData = () => {
    // 初始化测试前方法数据
    if (props.setupHooks?.length) {
      setupData.value = props.setupHooks.map(content => ({ setup: content }))
    }
    // 初始化测试后方法数据
    if (props.teardownHooks?.length) {
      teardownData.value = props.teardownHooks.map(content => ({ teardown: content }))
    }
    // 初始化后触发数据提交
    handleSetupHooksData()
    handleTeardownHooksData()
  }
  watch(
    () => [props.teardownHooks, props.setupHooks],
    ([newTeardown, newSetup], [oldTeardown, oldSetup]) => {
      // 只有当两个值中至少有一个发生变化时才执行
      if (newTeardown !== oldTeardown || newSetup !== oldSetup) {
        initHooksData();
      }
    }
    // 基本类型不需要 deep: true
  );

  const handleSetupHooksData = () => {
    // 提取非空的方法内容，避免空值提交
    const setupDatas = setupData.value
      .map(item => item.setup)
      .filter(content => content.trim() !== '')
    emit('setupHooksData', setupDatas)
  }

  const handleTeardownHooksData = () => {
    // 提取非空的方法内容，避免空值提交
    const teardownDatas = teardownData.value
      .map(item => item.teardown)
      .filter(content => content.trim() !== '')
    emit('teardownHooksData', teardownDatas)
  }

  const handleCellMouseEnter = (row) => {
    currentRow.value = row
  }
  const handleCellMouseLeave = () => {
    currentRow.value = ''
  }

  const handleSetupAdd = () => {
    setupData.value.push({ setup: '' })
    handleSetupHooksData() // 添加后实时提交
  }
  const handleSetupDelete = (index) => {
    setupData.value.splice(index, 1)
    handleSetupHooksData() // 删除后实时提交
  }

  const handleTeardownAdd = () => {
    teardownData.value.push({ teardown: '' })
    handleTeardownHooksData() // 添加后实时提交
  }
  const handleTeardownDelete = (index) => {
    teardownData.value.splice(index, 1)
    handleTeardownHooksData() // 删除后实时提交
  }

  const parseHooks = () => {
    return {
      setup_hooks: setupData.value
        .map(item => item.setup)
        .filter(content => content.trim() !== ''),
      teardown_hooks: teardownData.value
        .map(item => item.teardown)
        .filter(content => content.trim() !== '')
    }
  }

  onMounted(() => {
    initHooksData()
  })

  defineExpose({ parseHooks })
</script>

<style scoped></style>