<template>
  <el-dialog v-model="visible" title="标签管理" width="700px">
    <div class="gva-btn-list" style="margin-bottom: 10px;">
      <el-button type="primary" @click="startCreateRow">新增</el-button>
    </div>
    <el-table :data="table" style="width: 100%">
      <el-table-column label="名称" prop="name">
        <template #default="scope">
          <template v-if="scope.row._editing || scope.row._isNewRow">
            <el-input v-model="scope.row._nameEdit" placeholder="请输入名称" />
          </template>
          <template v-else>
            {{ scope.row.name }}
          </template>
        </template>
      </el-table-column>
      <el-table-column label="备注信息" prop="description">
        <template #default="scope">
          <template v-if="scope.row._editing || scope.row._isNewRow">
            <el-input v-model="scope.row._descEdit" placeholder="请输入备注信息" />
          </template>
          <template v-else>
            {{ scope.row.description }}
          </template>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="200">
        <template #default="scope">
          <template v-if="scope.row._isNewRow">
            <el-button type="primary" link @click="saveNewRow(scope.row)">保存</el-button>
            <el-button type="primary" link @click="cancelNewRow">取消</el-button>
          </template>
          <template v-else>
            <template v-if="scope.row._editing">
              <el-button type="primary" link @click="saveRow(scope.row)">保存</el-button>
              <el-button type="primary" link @click="cancelEdit(scope.row)">取消</el-button>
            </template>
            <template v-else>
              <el-button type="primary" link @click="editRow(scope.row)">编辑</el-button>
              <el-button type="primary" link @click="deleteRow(scope.row)">删除</el-button>
            </template>
          </template>
        </template>
      </el-table-column>
    </el-table>
    <div class="gva-pagination" style="margin-top: 10px;">
      <el-pagination
        layout="total, sizes, prev, pager, next, jumper"
        :current-page="page"
        :page-size="pageSize"
        :page-sizes="[10, 30, 50, 100]"
        :total="total"
        @current-change="handlePageChange"
        @size-change="handleSizeChange"
      />
    </div>
  </el-dialog>
</template>

<script setup>
import { computed, ref, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { createTag, updateTag, deleteTag, getTagList } from '@/api/automation/tag'

const props = defineProps({ modelValue: Boolean })
const emit = defineEmits(['update:modelValue', 'changed'])

const visible = computed({
  get: () => props.modelValue,
  set: v => emit('update:modelValue', v)
})

const table = ref([])
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)

const loadTable = async () => {
  const res = await getTagList({ page: page.value, pageSize: pageSize.value })
  if (res.code === 0) {
    const list = res.data.list || []
    table.value = list.map(item => ({
      ...item,
      _editing: false,
      _isNewRow: false,
      _nameEdit: item.name,
      _descEdit: item.description
    }))
    total.value = res.data.total
    page.value = res.data.page
    pageSize.value = res.data.pageSize
  }
}

watch(visible, (v) => { if (v) loadTable() })

const handlePageChange = (val) => { page.value = val; loadTable() }
const handleSizeChange = (val) => { pageSize.value = val; loadTable() }

const editRow = (row) => { row._editing = true }
const cancelEdit = (row) => { row._editing = false; row._nameEdit = row.name; row._descEdit = row.description }
const saveRow = async (row) => {
  const payload = { ID: row.ID, name: row._nameEdit, description: row._descEdit }
  if (!payload.name || payload.name.trim() === '') { ElMessage({ type: 'warning', message: '名称为必填' }); return }
  const res = await updateTag(payload)
  if (res.code === 0) {
    row.name = row._nameEdit
    row.description = row._descEdit
    row._editing = false
    ElMessage({ type: 'success', message: '保存成功' })
    emit('changed')
    loadTable()
  }
}

const deleteRow = async (row) => {
  ElMessageBox.confirm('确定要删除该标签吗?', '提示', { confirmButtonText: '确定', cancelButtonText: '取消', type: 'warning' })
    .then(async () => {
      const res = await deleteTag({ ID: row.ID })
      if (res.code === 0) {
        ElMessage({ type: 'success', message: '删除成功' })
        if (table.value.length === 1 && page.value > 1) page.value--
        emit('changed')
        loadTable()
      }
    })
}

const startCreateRow = () => {
  const hasNew = table.value.some(r => r._isNewRow)
  if (hasNew) return
  table.value.push({ _isNewRow: true, _nameEdit: '', _descEdit: '', _editing: true })
}
const cancelNewRow = () => { table.value = table.value.filter(r => !r._isNewRow) }
const saveNewRow = async (row) => {
  const payload = { name: row._nameEdit, description: row._descEdit }
  if (!payload.name || payload.name.trim() === '') { ElMessage({ type: 'warning', message: '名称为必填' }); return }
  const res = await createTag(payload)
  if (res.code === 0) {
    ElMessage({ type: 'success', message: '创建成功' })
    emit('changed')
    await loadTable()
  }
}
</script>

<style>
</style>