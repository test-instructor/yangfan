<template>
  <div style="width: 300px">

    <div>
      <a-input-search
          style="margin-bottom: 8px; max-width: 130px"
          v-model="searchKey"
          placeholder="搜索节点"
      />
      <a-button type="outline" style="margin-left: 10px" @click="addParentNode" >新增一级节点</a-button>
    </div>
    <a-tree
        :defaultExpandAll="true"
        :blockNode="true"
        :data="treeData"
        @select="getTreeID"
        :default-selected-keys="defaultKeys"
    >
      <template #title="nodeData">
        <IconPlus
            style="position: absolute; right: 48px; font-size: 12px; top: 10px; color: #3370ff;"
            @click="() => addNode(nodeData)"
        />
        <IconEdit
            style="position: absolute; right: 28px; font-size: 12px; top: 10px; color: #3370ff;"
            @click="() => updateNode(nodeData)"
        />
        <IconDelete
            style="position: absolute; right: 8px; font-size: 12px; top: 10px; color: #3370ff;"
            @click="() => delNode(nodeData)"
        />
        <template v-if="index = getMatchIndex(nodeData?.title), index < 0">{{ nodeData?.title }}</template>
        <span v-else>
          {{ nodeData?.title?.substr(0, index) }}
          <span style="color: var(--color-primary-light-4);">
            {{ nodeData?.title?.substr(index, searchKey.length) }}
          </span>{{ nodeData?.title?.substr(index + searchKey.length) }}
        </span>
      </template>
    </a-tree>
  </div>
  <a-modal :visible="addVisible" :title="title" @cancel="handleCancel" @ok="handleBeforeOk" >
    <a-input  v-model="node.name" :placeholder="placeholder" allow-clear/>
  </a-modal>
  <a-modal :visible="delVisible" @ok="handleDel" @cancel="handleCancelDel">
    <template #title>
      删除节点
    </template>
    <div>{{ delNodeMessage }}</div>
  </a-modal>
</template>

<script setup>
import {ref, computed, reactive, onBeforeMount} from 'vue';
import { IconPlus, IconEdit, IconDelete } from '@arco-design/web-vue/es/icon';
// import {getTree} from "@/api/interfaceMenu";
import {getTree, addTree, editTree, delTree} from '@/api/interfaceMenu'
import {Message} from "@arco-design/web-vue";

const originTreeData = ref();
let params = {'menutype': 1}
const addVisible =ref(false)
const delNodeMessage =ref("")
const delVisible =ref(false)
const delName = ref("")
const defaultKeys = ref("0")

const project_id = JSON.parse(window.localStorage.getItem('project')).ID
var  trees = ref({'id': 0, 'isdel': 1})
const node = reactive({
  name: "",
  project: project_id,
  parent: 0,
  id: 0,
})
const title = ref("")
const placeholder = ref("")
const searchKey = ref('');
const treeData = computed(() => {
  if (!searchKey.value) return originTreeData;
  return searchData(searchKey.value);
})
const getTreeID = (nodeData) => {
}
const nodeType = ref(1)
const searchData = (keyword) => {
  const loop = (data) => {
    const result = [];
    data.forEach(item => {
      if (item.title.toLowerCase().indexOf(keyword.toLowerCase()) > -1) {
        result.push({...item});
      } else if (item.children) {
        const filterData = loop(item.children);
        if (filterData.length) {
          result.push({
            ...item,
            children: filterData
          })
        }
      }
    })
    return result;
  }

  return loop(originTreeData);
}
const getMatchIndex = (title) => {
  if (!searchKey.value) return -1;
  return title.toLowerCase().indexOf(searchKey.value.toLowerCase());
}

onBeforeMount(() => {
  // 调接口
})

const addParentNode = () => {
  nodeType.value = 1
  title.value = "新增一级节点"
  placeholder.value = "请输入节点名称"
  addVisible.value = true
}

const addNode = (nodeData) => {
  nodeType.value = 1
  title.value = `【${nodeData.label}】新增子节点`
  placeholder.value = "请输入节点名称"
  node.parent = nodeData.id
  addVisible.value = true
}

const updateNode = (nodeData) => {
  nodeType.value = 2
  title.value = `修改节点【${nodeData.label}】名称`
  node.id = nodeData.key
  node.name = nodeData.label
  addVisible.value = true
}

const delNode = (nodeData) => {
  delNodeMessage.value = `节点【${nodeData.label}】及子节点和对应的数据将被删除, 删除后数据无法恢复，是否继续删除?`
  delVisible.value = true
  delName.value = nodeData.title
  trees.value.id = nodeData.key
}

const handleDel = () => {
  delTree(trees.value, params).then((response) => {
    originTreeData.value = response.data.list
    delVisible.value = false
    Message.success(`节点【${delName.value}】创建成功`)
  })
}

const handleCancel = () => {
  addVisible.value = false
  node.name = ""
  node.parent = 0
}

const handleCancelDel = () => {
  delVisible.value = false
}

const handleBeforeOk = () => {
  addTrees(node)
}



const getTrees = () => {
  getTree(params).then((response) => {

    if (response.data.list && response.data.list.length>0){

      if (defaultKeys.value === "0"){
        defaultKeys.value = "1"
      }
      originTreeData.value = response.data.list
      // dom && dom.click();
    }
  })
}
getTrees()
defaultKeys.value = "6"
// defaultKeys.value = originTreeData.value[0].key.toString()
// defaultKeys.value = "1"
const addTrees = (newChild) => {
  if (nodeType.value===1){
    addTree(newChild, params).then((response) => {
      originTreeData.value = response.data.list
      Message.success(`节点【${newChild.name}】创建成功`)
      addVisible.value = false
      node.name = ""
      node.parent = 0
    })
  }
  if (nodeType.value===2){
    editTree(newChild, params).then((response) => {
      originTreeData.value = response.data.list
      addVisible.value = false
      Message.success(`节点【${newChild.name}】修改成功`)
      node.name = ""
      node.parent = 0
      node.id = 0
    })
  }
}

</script>


