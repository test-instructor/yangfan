<template>

  <div class="menu-wrapper">
    <el-input
      v-model="filterText"
      placeholder="输入关键字进行过滤"
      class="filter-input"
    />
    <el-button
      v-if="!props.detail"
      type="primary"
      @click="addNewQuestion"
      class="add-btn"
    >
      添加一级节点
    </el-button>

    <div class="tree-container">
      <el-tree
        ref="treeRef"
        :data="trees"
        :props="defaultProps"
        :filter-node-method="filterNode"
        :current-node-key="currentNodeKey"
        :default-checked-keys="[currentNodeKey]"
        node-key="id"
        default-expand-all
        highlight-current
        check-on-click-node
        :expand-on-click-node="false"
        :render-after-expand="true"
        @node-click="handleClick"
        @node-drag-start="handleDragStart"
        @node-drag-enter="handleDragEnter"
        @node-drag-leave="handleDragLeave"
        @node-drag-over="handleDragOver"
        @node-drag-end="handleDragEnd"
        @node-drop="handleDrop"
      >
        <template #default="{ node, data }">
          <span class="custom-tree-node">
            <span>{{ node.label }}</span>
            <span v-if="!props.detail" class="node-actions">
              <a v-if="getNodeLevel(node) < 4" @click.stop="openAppend(data)" class="node-action">
                <el-icon><add /></el-icon>
              </a>
              <a @click.stop="openEdit(node, data)" class="node-action">
                <el-icon><edit /></el-icon>
              </a>
              <a @click.stop="openDelete(node, data)" class="node-action">
                <el-icon><delete /></el-icon>
              </a>
            </span>
          </span>
        </template>
      </el-tree>
    </div>
  </div>

</template>

<script setup>
  import { ref, watch, onMounted, nextTick } from 'vue'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import {
    createCategoryMenu,
    deleteCategoryMenu,
    deleteCategoryMenuByIds,
    updateCategoryMenu,
    findCategoryMenu,
    getCategoryMenuList
  } from '@/api/platform/categorymenu.js'
  import { useUserStore } from '@/pinia/modules/user'

  const userStore = useUserStore()
  const projectId = userStore.userInfo.projectId
  // 定义组件属性
  const props = defineProps({
    menutype: {
      type: String,
      default: '99999'
    },
    eventType: {
      type: String,
      default: '0'
    },
    detail:{
      type: Boolean,
      default: false
    }
  })

  // 定义事件
  const emit = defineEmits(['getTreeID'])

  // 响应式数据
  const filterText = ref('')
  const trees = ref([])
  const treeRef = ref(null)
  const currentNodeKey = ref(0)

  // 常量定义
  const defaultProps = {
    children: 'children',
    label: 'label'
  }


  // 存储键映射
  const storageKeyMap = {
    '1-0': 'menuAddCase',
    '1-1': 'menu',
    '2-0': 'menuCaseAddStep',
    '2-1': 'menuStep',
    '3-0': 'menuTaskAddStep',
    '3-1': 'menuCase'
  }

  // 监听器
  watch(filterText, (val) => {
    treeRef.value.filter(val)
  })

  // 生命周期
  onMounted(() => {
    getTrees()
    currentNodeKey.value = 1
  })

  // 方法定义
  const getStorageKey = () => {
    return storageKeyMap[`${props.menutype}-${props.eventType}`]
  }
  const getTrees = async () => {
    try {
      const params = { menuType: props.menutype }
      const response = await getCategoryMenuList(params)
      if (response.code === 0) {
        trees.value = response.data.list || []
        setNodeStatus()
      }
    } catch (error) {
      ElMessage.error('获取菜单树失败')
    }
  }

  watch(props, (newVal) => {
    getTrees()
  }, { deep: false })

  const addNewQuestion = async () => {
    try {
      const { value } = await ElMessageBox.prompt('请输入一级节点名称', '新增一级节点', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        inputPattern: /.+/,
        inputErrorMessage: '节点名称不能为空'
      })

      const newChild = {
        name: value,
        project: projectId,
        parent: 0,
        menuType: props.menutype
      }

      await createCategoryMenu(newChild)
      ElMessage.success(`新增节点: ${value}`)
      await getTrees()
    } catch (error) {
      if (error !== 'cancel') {
        ElMessage.info('取消输入')
      }
    }
  }

  const openAppend = async (data) => {
    try {
      const { value } = await ElMessageBox.prompt('请输入节点名称', '新增子节点', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        inputPattern: /.+/,
        inputErrorMessage: '节点名称不能为空'
      })

      await appendNode(data, value)
      ElMessage.success(`新增节点: ${value}`)
    } catch (error) {
      if (error !== 'cancel') {
        ElMessage.info('取消输入')
      }
    }
  }

  const openEdit = async (node, data) => {
    try {
      const { value } = await ElMessageBox.prompt('请输入节点名称', '修改节点名称', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        inputValue: data.label,
        inputPattern: /.+/,
        inputErrorMessage: '节点名称不能为空'
      })

      await editTreeNode(data, value)
      ElMessage.success(`修改节点: ${value}`)
    } catch (error) {
      if (error !== 'cancel') {
        ElMessage.info('取消输入')
      }
    }
  }

  const openDelete = async (node, data) => {
    try {
      await ElMessageBox.confirm(
        `节点${data.label}及子节点和对应的数据将被删除, 是否继续?`,
        '删除节点',
        {
          confirmButtonText: '删除',
          cancelButtonText: '取消',
          type: 'error'
        }
      )

      await removeNode(node, data)
      ElMessage.success('删除成功!')
    } catch (error) {
      if (error !== 'cancel') {
        ElMessage.info('已取消删除')
      }
    }
  }

  const appendNode = async (data, value) => {
    const newChild = {
      parent: data.id,
      name: value,
      project: projectId,
      menuType: props.menutype
    }

    await createCategoryMenu(newChild)
    await getTrees()
  }

  const editTreeNode = async (data, value) => {
    const updateData = {
      id: data.id,
      name: value,
      parent: data.parent,
      project: projectId,
      menuType: props.menutype
    }

    await updateCategoryMenu(updateData)
    await getTrees()
  }

  const removeNode = async (node, data) => {
    const deleteData = {
      id: data.id,
      isdel: 1
    }

    await deleteCategoryMenu(deleteData)
    await getTrees()
    emit('getTreeID', 0)
  }

  const filterNode = (value, data) => {
    if (!value) return true
    return data.label.indexOf(value) !== -1
  }

  const findTreeNodeById = (list, id) => {
    for (const element of list) {
      if (element.id === id) {
        return element
      }
      if (element.children && element.children.length > 0) {
        const foundElement = findTreeNodeById(element.children, id)
        if (foundElement) {
          return foundElement
        }
      }
    }
    return null
  }

  const getDefaultTreeNode = (data) => {
    const storageKey = getStorageKey()
    if (!storageKey) return data

    let getTreeID = localStorage.getItem('getTreeID')
    let treeData
    try {
      // 尝试解析为对象，如果解析失败（比如是数字、字符串），则初始化为空对象
      treeData = JSON.parse(getTreeID) || {}
      // 确保解析后是对象（如果原始值是数组/其他类型，也转为对象）
      if (typeof treeData !== 'object' || treeData === null) {
        treeData = {}
      }
    } catch (e) {
      // 解析出错（比如原始值不是JSON格式），直接初始化为空对象
      treeData = {}
    }
    const storedId = treeData[props.menutype]

    if (storedId && Number(storedId) > 0) {
      const foundNode = findTreeNodeById(trees.value, Number(storedId))
      if (foundNode) {
        return foundNode
      }
    }
    return data
  }

  const setNodeStatus = () => {
    nextTick(() => {
      if (trees.value && trees.value.length > 0) {

        let targetNode = null
        
        if (!props.defaultFirst) {
          let getTreeID = localStorage.getItem('getTreeID')
          let treeData
          try {
            // 尝试解析为对象，如果解析失败（比如是数字、字符串），则初始化为空对象
            treeData = JSON.parse(getTreeID) || {}
            // 确保解析后是对象（如果原始值是数组/其他类型，也转为对象）
            if (typeof treeData !== 'object' || treeData === null) {
              treeData = {}
            }
          } catch (e) {
            // 解析出错（比如原始值不是JSON格式），直接初始化为空对象
            treeData = {}
          }
          const storedTreeId = treeData[props.menutype]

          let targetId = storedTreeId ? Number(storedTreeId) : null

          if (targetId) {
            targetNode = findTreeNodeById(trees.value, targetId)
          }
        }

        if (!targetNode) {
          let defaultData = trees.value[0]
          if (!props.defaultFirst) {
            defaultData = getDefaultTreeNode(defaultData)
          }
          targetNode = defaultData
        }

        if (treeRef.value && targetNode) {
          treeRef.value.setCurrentKey(targetNode.id)
          currentNodeKey.value = targetNode.id
          handleClick(targetNode)
        }
      }
    })
  }

  const getNodeLevel = (node) => {
    let level = 0
    let currentNode = node
    while (currentNode.parent) {
      currentNode = currentNode.parent
      level++
    }
    return level
  }

  const handleClick = (data) => {
    emit('getTreeID', data.id)

    // 1. 获取localStorage中的值，强制转为对象（无论原始值是什么类型）
    let getTreeID = localStorage.getItem('getTreeID')
    let treeData
    try {
      // 尝试解析为对象，如果解析失败（比如是数字、字符串），则初始化为空对象
      treeData = JSON.parse(getTreeID) || {}
      // 确保解析后是对象（如果原始值是数组/其他类型，也转为对象）
      if (typeof treeData !== 'object' || treeData === null) {
        treeData = {}
      }
    } catch (e) {
      // 解析出错（比如原始值不是JSON格式），直接初始化为空对象
      treeData = {}
    }

    // 2. 处理props.menutype可能为undefined的情况（添加默认键名，避免报错）
    const menuTypeKey = props.menutype || 'default_menutype' // 若为undefined，用默认键

    // 3. 设置属性并存储
    treeData[menuTypeKey] = data.id
    localStorage.setItem('getTreeID', JSON.stringify(treeData))
  }

  // 拖拽相关方法（保持原有功能，可根据需要实现）
  const handleDragStart = (node, ev) => {
  }
  const handleDragEnter = (draggingNode, dropNode, ev) => {
  }
  const handleDragLeave = (draggingNode, dropNode, ev) => {
  }
  const handleDragOver = (draggingNode, dropNode, ev) => {
  }
  const handleDragEnd = (draggingNode, dropNode, dropType, ev) => {
  }
  const handleDrop = (draggingNode, dropNode, dropType, ev) => {
  }
</script>

<style scoped>
  .menu-wrapper {
    background-color: var(--el-bg-color);
    margin: 8px;
    border-radius: 8px;
    padding: 12px 10px;
    border: 1px solid var(--el-border-color-light);
  }

  .filter-input {
    margin: 8px 0 10px;
  }

  .add-btn {
    margin: 2px 10px;
  }

  .tree-container {
    overflow-y: auto;
    height: 560px;
    margin: 2px 10px;
  }

  .mask::-webkit-scrollbar {
    width: 0;
  }

  .parent {
    display: flex;
    padding: 0px;
    height: 90%;
  }

  .left {
    width: 300px;
    height: 90%;
  }

  /* 高亮当前选中的树节点 */
  :deep(.el-tree-node.is-current > .el-tree-node__content) {
    background-color: var(--el-color-primary-light-9);
  }

  /* 树节点点击时的样式设置 */
  :deep(.el-tree-node:focus > .el-tree-node__content) {
    background-color: var(--el-color-primary-light-9);
  }

  /* 树节点hover时的样式设置 */
  :deep(.el-tree-node__content:hover) {
    background-color: var(--el-color-primary-light-8);
  }

  .custom-tree-node {
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: space-between;
    font-size: 16px;
  }

  .node-actions {
    display: inline-flex;
    align-items: center;
  }

  .node-action {
    margin: 0 10px;
    color: var(--el-text-color-regular);
  }

  .node-action:hover {
    color: var(--el-color-primary);
  }
</style>
