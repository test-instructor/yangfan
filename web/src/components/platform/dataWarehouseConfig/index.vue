<template>
  <div class="data-warehouse-config" :style="{ height: containerHeight }">
    <el-form :model="localConfig"  class="config-form">
      <!-- Type Selection -->
      <el-form-item label="数据类型" required>
        <el-select 
          v-model="localConfig.type" 
          @change="handleTypeChange" 
          placeholder="请选择数据类型" 
          style="width: 200px"
          filterable
          clearable
        >
          <el-option
            v-for="item in typeOptions"
            :key="item.type"
            :label="item.name ? `${item.name} (${item.type})` : item.type"
            :value="item.type"
          />
        </el-select>
      </el-form-item>

      <!-- Count -->
      <el-form-item label="获取数量" v-if="!hideCount">
        <el-input-number v-model="localConfig.count" :min="1" :max="1000" class="count-input" />
      </el-form-item>

      <!-- Filter Configuration -->
      <el-form-item class="filter-form-item">
        <template #label>
             <div class="label-with-btn">
                <span>筛选条件</span>
             </div>
        </template>
        <div class="filter-container">
            <!-- Global Logic -->
          <div class="logic-header">
             <div v-if="(localConfig.filter?.groups?.length || 0) > 1" class="logic-header">
                <el-tag type="info" size="small">组间逻辑</el-tag>
                <el-radio-group v-model="localConfig.filter.logic" size="small" class="logic-radio">
                  <el-radio-button label="AND">且 (AND)</el-radio-button>
                  <el-radio-button label="OR">或 (OR)</el-radio-button>
                </el-radio-group>
             </div>
            <el-button link type="primary" icon="Plus" @click="addGroup" size="small">添加</el-button>
          </div>

          <div class="filter-scroll-area">
            <transition-group name="list">
            <div v-for="(group, gIndex) in localConfig.filter?.groups" :key="gIndex" class="group-box">
                <div class="group-header">
                    <div class="group-title-area">
                        <span class="group-title">条件组 {{ gIndex + 1 }}</span>
                        <div v-if="(group.conditions?.length || 0) > 1" class="group-logic">
                            <span class="label">组内逻辑:</span>
                            <el-radio-group v-model="group.logic" size="small">
                                <el-radio-button label="AND">且</el-radio-button>
                                <el-radio-button label="OR">或</el-radio-button>
                            </el-radio-group>
                        </div>
                    </div>
                    <el-button v-if="(localConfig.filter?.groups?.length || 0) > 1" type="danger" link size="small" @click="removeGroup(gIndex)" icon="Delete" class="delete-group-btn">删除组</el-button>
                </div>
                
                <div class="conditions-list">
                    <div v-for="(cond, cIndex) in group.conditions" :key="cIndex" class="condition-row">
                        <!-- Field -->
                        <el-select 
                            v-model="cond.field" 
                            filterable 
                            allow-create 
                            default-first-option 
                            placeholder="字段" 
                            class="field-select" 
                            @change="handleFieldChange(cond)"
                        >
                            <el-option v-for="field in fieldOptions" :key="field" :label="field" :value="field" />
                        </el-select>

                        <!-- Operator -->
                        <el-select v-model="cond.operator" placeholder="操作符" class="operator-select">
                            <el-option v-for="op in operatorOptions" :key="op.value" :label="op.label" :value="op.value" />
                        </el-select>

                        <!-- Value Area -->
                        <div class="value-area">
                            <template v-if="!['IS_NULL', 'IS_NOT_NULL'].includes(cond.operator)">
                                <!-- Value Type Selector -->
                                <el-select v-model="cond.valueType" placeholder="类型" class="type-select" @change="handleValueTypeChange(cond)">
                                    <el-option label="文本" value="string" />
                                    <el-option label="数字" value="number" />
                                    <el-option label="布尔" value="boolean" />
                                </el-select>

                                <!-- Value Input based on Type -->
                                <template v-if="cond.valueType === 'boolean'">
                                    <el-switch
                                        v-model="cond.value"
                                        inline-prompt
                                        active-text="True"
                                        inactive-text="False"
                                    />
                                </template>
                                <template v-else-if="cond.valueType === 'number'">
                                    <el-input-number v-model="cond.value" controls-position="right" class="value-input" placeholder="值" />
                                    <span v-if="['BETWEEN', 'NOT_BETWEEN'].includes(cond.operator)" class="separator">-</span>
                                    <el-input-number 
                                        v-if="['BETWEEN', 'NOT_BETWEEN'].includes(cond.operator)"
                                        v-model="cond.value2" 
                                        controls-position="right" 
                                        class="value-input" 
                                        placeholder="结束值" 
                                    />
                                </template>
                                <template v-else>
                                    <el-input v-model="cond.value" placeholder="值" class="value-input" />
                                    <span v-if="['BETWEEN', 'NOT_BETWEEN'].includes(cond.operator)" class="separator">-</span>
                                    <el-input 
                                        v-if="['BETWEEN', 'NOT_BETWEEN'].includes(cond.operator)"
                                        v-model="cond.value2" 
                                        placeholder="结束值" 
                                        class="value-input" 
                                    />
                                </template>
                            </template>
                        </div>
                        
                        <el-button v-if="group.conditions.length > 1" icon="Delete" circle plain type="danger" size="small" @click="removeCondition(group, cIndex)" class="delete-btn" />
                    </div>
                </div>

                <div class="group-actions">
                    <el-button type="primary" link size="small" icon="Plus" @click="addCondition(group)">添加条件</el-button>
                </div>
            </div>
            </transition-group>
          </div>
            
            <!-- Add Group button moved to label -->
        </div>
      </el-form-item>
    </el-form>
  </div>
</template>

<script setup>
import { ref, watch, onMounted, computed, nextTick } from 'vue'
import { getDataCategoryTypeList } from '@/api/datawarehouse/dataCategoryManagement'
import { ElMessage } from 'element-plus'

const props = defineProps({
  modelValue: {
    type: Object,
    default: () => ({})
  },
  hideCount: {
    type: Boolean,
    default: false
  },
  height: {
    type: [Number, String],
    default: '100%'
  }
})

const emit = defineEmits(['update:modelValue'])

// Initial structure
const localConfig = ref({
    type: '',
    count: 0,
    filter: {
        logic: 'AND',
        groups: []
    }
})

const typeOptions = ref([])
const currentTypeData = ref(null)

const fieldOptions = computed(() => {
    if (!currentTypeData.value) return []
    return Object.keys(currentTypeData.value)
})

const containerHeight = computed(() => {
    if (typeof props.height === 'number') {
        return props.height + 'px'
    }
    return props.height
})

const operatorOptions = [
  { label: '等于', value: 'EQ' },
  { label: '不等于', value: 'NE' },
  { label: '大于', value: 'GT' },
  { label: '小于', value: 'LT' },
  { label: '大于等于', value: 'GE' },
  { label: '小于等于', value: 'LE' },
  { label: '区间', value: 'BETWEEN' },
  { label: '不在区间', value: 'NOT_BETWEEN' },
  { label: '为空', value: 'IS_NULL' },
  { label: '不为空', value: 'IS_NOT_NULL' },
  { label: '模糊匹配', value: 'LIKE' },
  { label: '反向模糊匹配', value: 'NOT_LIKE' }
]

// Initialize
onMounted(async () => {
    try {
        const res = await getDataCategoryTypeList()
        if (res.code === 0) {
            typeOptions.value = res.data || []
            // If type is already selected, try to find data
            if (localConfig.value.type) {
                const found = typeOptions.value.find(item => item.type === localConfig.value.type)
                if (found) {
                     currentTypeData.value = found.value
                }
            }
        }
    } catch (e) {
        console.error(e)
    }
})

// Sync from parent
watch(() => props.modelValue, (val) => {
    if (!val) return
    
    // Check if external update is different from local to avoid infinite loop
    const currentLocal = JSON.stringify({
        type: localConfig.value.type,
        count: localConfig.value.count,
        filter: localConfig.value.filter
    })
    
    // Construct what the new local would be
    const newCount = val.count !== undefined ? val.count : (props.hideCount ? 1 : 0)
    const newFilter = val.filter ? val.filter : { logic: 'AND', groups: [] }
    
    const incoming = JSON.stringify({
        type: val.type || '',
        count: newCount,
        filter: newFilter
    })

    if (currentLocal === incoming) return

    // Merge props into localConfig
    localConfig.value.type = val.type || ''
    localConfig.value.count = newCount
    
    // Deep merge filter or default
    if (val.filter) {
        localConfig.value.filter = JSON.parse(JSON.stringify(val.filter))
    } else {
        localConfig.value.filter = { logic: 'AND', groups: [] }
    }
    
    // Re-process conditions to add valueType for UI if missing
    if (localConfig.value.filter.groups) {
        localConfig.value.filter.groups.forEach(group => {
            if (group.conditions) {
                group.conditions.forEach(cond => {
                    if (!cond.valueType) {
                        cond.valueType = typeof cond.value === 'number' ? 'number' : (typeof cond.value === 'boolean' ? 'boolean' : 'string')
                    }
                })
            }
        })
    }
    
}, { immediate: true, deep: true })

// Sync to parent
watch(localConfig, (val) => {
    // Clean up data before emitting? 
    // We need to ensure types are correct.
    // JSON.stringify will handle numbers/booleans if they are actual numbers/booleans in JS.
    if (!val.type) {
        emit('update:modelValue', {})
        return
    }
    emit('update:modelValue', val)
}, { deep: true })

const handleTypeChange = (type) => {
    const found = typeOptions.value.find(item => item.type === type)
    if (found) {
        currentTypeData.value = found.value
        // 自动添加一个条件组
        if (!localConfig.value.filter.groups || localConfig.value.filter.groups.length === 0) {
             addGroup()
        }
    } else {
        currentTypeData.value = null
    }
}

const handleFieldChange = (cond) => {
    // Try to infer type from sample data
    if (currentTypeData.value && cond.field) {
        const sampleVal = currentTypeData.value[cond.field]
        if (typeof sampleVal === 'number') cond.valueType = 'number'
        else if (typeof sampleVal === 'boolean') cond.valueType = 'boolean'
        else cond.valueType = 'string'
    } else {
        // Default to string if no sample data found
        cond.valueType = 'string'
    }
    cond.value = '' // Reset value
    cond.value2 = ''
}

const handleValueTypeChange = (cond) => {
    // Reset value when type changes to avoid type mismatch
    cond.value = cond.valueType === 'boolean' ? false : (cond.valueType === 'number' ? 0 : '')
    cond.value2 = cond.valueType === 'boolean' ? false : (cond.valueType === 'number' ? 0 : '')
}

const addGroup = () => {
    if (!localConfig.value.filter) localConfig.value.filter = { logic: 'AND', groups: [] }
    if (!localConfig.value.filter.groups) localConfig.value.filter.groups = []
    localConfig.value.filter.groups.push({
        logic: 'AND',
        conditions: [{ field: '', operator: 'EQ', value: '', valueType: 'string' }]
    })
}

const removeGroup = (index) => {
    localConfig.value.filter.groups.splice(index, 1)
}

const addCondition = (group) => {
    if (!group.conditions) group.conditions = []
    group.conditions.push({ field: '', operator: 'EQ', value: '', valueType: 'string' })
}

const removeCondition = (group, index) => {
    group.conditions.splice(index, 1)
}

</script>

<style scoped>
.data-warehouse-config {
    height: 100%;
    overflow: hidden;
    padding-right: 12px;
}

.config-form {
    height: 100%;
    display: flex;
    flex-direction: column;
}

/* 表单项之间增加 5px 间距 */
.config-form :deep(.el-form-item) {
    margin-bottom: 25px;
}

/* 让筛选条件这一项占满剩余高度 */
.filter-form-item {
    flex: 1;
    display: flex;
    align-items: stretch;
    min-height: 0; /* 防止内容撑开容器 */
    overflow: hidden;
}

.filter-form-item :deep(.el-form-item__content) {

    min-height: 0; /* 防止内容撑开容器 */
    overflow: hidden;
}

.label-with-btn {
    display: flex;
    align-items: center;
    gap: 4px;
    justify-content: flex-start;
    width: 100%;
}


.filter-container {
    display: flex;
    flex-direction: column;
    gap: 10px;
    height: 100%;
    min-height: 0;
    overflow: hidden;
}

.filter-scroll-area {
    flex: 1;
    overflow-y: auto;
    display: flex;
    flex-direction: column;
    gap: 17px;
    padding-right: 8px;
}

.logic-header {
    display: flex;
    align-items: center;
    gap: 10px;
    padding-bottom: 4px;
}

.group-box {
    border: 1px solid #e4e7ed;
    padding: 12px;
    border-radius: 6px;
    background: #f8f9fb;
    transition: all 0.3s;
    width: 1000px  ;
}

.group-box:hover {
    box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.05);
    border-color: #dcdfe6;
}

.group-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 12px;
    padding-bottom: 8px;
    border-bottom: 1px solid #ebeef5;
}

.group-title-area {
    display: flex;
    align-items: center;
    gap: 16px;
}

.group-title {
    font-weight: 600;
    font-size: 14px;
    color: #303133;
    display: flex;
    align-items: center;
}

.group-title::before {
    content: '';
    display: inline-block;
    width: 4px;
    height: 14px;
    background: #409eff;
    margin-right: 8px;
    border-radius: 2px;
}

.group-logic {
    display: flex;
    align-items: center;
    gap: 8px;
}

.group-logic .label {
    font-size: 12px;
    color: #909399;
}

.conditions-list {
    display: flex;
    flex-direction: column;
    gap: 13px; /* 条件行之间增加 5px */
}

.condition-row {
    display: flex;
    align-items: center;
    gap: 8px;
    background: #fff;
    padding: 4px 10px;
    border-radius: 4px;
    border: 1px solid transparent;
}

.condition-row:hover {
    border-color: #ebeef5;
}

.field-select {
    width: 200px;
    flex-shrink: 0;
}

.operator-select {
    width: 150px;
    flex-shrink: 0;
}

.value-area {
    display: flex;
    align-items: center;
    gap: 8px;
}

.type-select {
    width: 120px;
    flex-shrink: 0;
}

.value-input {
    width: 220px;
    min-width: 220px;
    flex: 0 0 auto; /* 宽度固定，不再拉伸 */
}

.separator {
    color: #909399;
    padding: 0 4px;
}

.delete-btn {
    opacity: 0;
    transition: opacity 0.2s;
    margin-left: auto;
}

.condition-row:hover .delete-btn {
    opacity: 1;
}

.group-actions {
    margin-top: 15px; /* 比原来多 5px */
    padding-left: 4px;
}

.add-group-wrapper {
    margin-top: 4px;
}

.add-group-btn {
    width: 100%;
    border-style: dashed;
}

/* Animations */
.list-move,
.list-enter-active,
.list-leave-active {
  transition: all 0.3s ease;
}

.list-enter-from,
.list-leave-to {
  opacity: 0;
  transform: translateY(-10px);
}

.list-leave-active {
  position: absolute;
  width: 100%;
}
</style>
