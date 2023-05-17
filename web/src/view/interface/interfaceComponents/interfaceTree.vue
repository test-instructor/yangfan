<template>
  <div id="left">
    <el-input
      v-model="filterText"
      placeholder="输入关键字进行过滤"
      style="margin-bottom: 10px"
    />
    <el-button
      type="primary"
      round
      style="margin-bottom: 10px"
      @click="add_new_question"
      v-if="eventType === '1'"
      >添加一级节点
    </el-button>

    <div style="overflow-y: auto; height: 680px">
      <el-tree
        :data="trees"
        :render-after-expand="true"
        node-key="id"
        :default-expand-all="true"
        :expand-on-click-node="false"
        :filter-node-method="filterNode"
        ref="tree"
        @node-drag-start="handleDragStart"
        @node-drag-enter="handleDragEnter"
        @node-drag-leave="handleDragLeave"
        @node-drag-over="handleDragOver"
        @node-drag-end="handleDragEnd"
        @node-drop="handleDrop"
        @node-click="handleClick"
      >
        <template #default="{ node, data }" v-if="eventType === '1'">
          <span class="custom-tree-node">
            <span>{{ node.label }}</span>
            <span>
              <a @click="openappend(data)">
                <i class="iconfont icon-add"></i>
              </a>
              <a @click="openedit(node, data)">
                <i class="iconfont icon-edit"></i>
              </a>
              <a @click="open(node, data)">
                <i class="iconfont icon-delete"></i>
              </a>
            </span>
          </span>
        </template>
      </el-tree>
    </div>
  </div>
</template>

<script>
import { getTree, addTree, editTree, delTree } from "@/api/interfaceMenu";
let treeID = 0;

const project_id = JSON.parse(window.localStorage.getItem("project")).ID;
export default {
  props: ["menutype", "eventType"],
  // name: interfaceTree,
  data() {
    return {
      params: { menutype: this.menutype },
      eventType: this.eventType,
      filterText: "",
      trees: [],
      defaultProps: {
        children: "children",
        label: "label",
      },
    };
  },
  watch: {
    filterText(val) {
      this.$refs.tree.filter(val);
    },
  },
  created() {
    this.getTrees();
  },
  mounted() {},
  methods: {
    add_new_question() {
      this.$prompt("请输入一级节点名称", "新增一级节点", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        inputPattern: /./,
        inputErrorMessage: "节点名称不能为空",
      })
        .then(({ value }) => {
          const newChild = { name: value, project: project_id, parent: "0" };
          this.addTrees(newChild);
          // this.trees.push()
          this.$message({
            type: "success",
            message: "新增节点: " + value,
          });
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "取消输入",
          });
        });
    },
    openappend(data) {
      this.$prompt("请输入节点名称", "新增子节点", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        inputPattern: /./,
        inputErrorMessage: "节点名称不能为空",
      })
        .then(({ value }) => {
          this.append(data, value);
          this.$message({
            type: "success",
            message: "新增节点: " + value,
          });
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "取消输入",
          });
        });
    },
    open(node, data) {
      this.$confirm(
        `节点${data.label}及子节点和对应的数据将被删除, 是否继续?`,
        "删除节点",
        {
          confirmButtonText: "删除",
          cancelButtonText: "取消",
          type: "error",
        }
      )
        .then(() => {
          this.remove(node, data);
          this.$message({
            type: "success",
            message: "删除成功!",
          });
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消删除",
          });
        });
    },
    openedit(node, data) {
      this.$prompt("请输入节点名称", "修改节点名称", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        inputValue: data.label,
        inputPattern: /./,
        inputErrorMessage: "节点名称不能为空",
      })
        .then(({ value }) => {
          this.editTree(data, value);
          this.$message({
            type: "success",
            // message: '你的邮箱是: ' + value
            message: "修改节点: " + value,
          });
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "取消输入",
          });
        });
    },
    getTrees() {
      // 获取树结构数据
      this.listLoading = true;
      getTree(this.params).then((response) => {
        if (response.code === 0) {
          this.trees = response.data.list;
        }
        setTimeout(() => {
          const dom = document.querySelector(".el-tree .el-tree-node__content");
          dom && dom.click();
        }, 20);
      });
    },
    filterNode(value, data) {
      if (!value) return true;
      return data.label.indexOf(value) !== -1;
    },

    editTrees(editChild) {
      // 添加树节点
      var trees = {
        id: editChild.id,
        name: editChild.label,
        parent: editChild.parent,
        project: project_id,
      };
      editTree(trees, this.params).then((response) => {
        if (response.code === 0) {
          this.trees = response.data.list;
        }
        setTimeout(() => {
          this.listLoading = false;
        }, 1.5 * 1000);
      });
    },

    delTrees(delChild) {
      // 添加树节点
      var trees = delChild;
      delTree(trees, this.params).then((response) => {
        if (response.code === 0) {
          this.trees = response.data.list;
        }
        setTimeout(() => {
          this.listLoading = false;
        }, 1.5 * 1000);
      });
    },

    addTrees(newChild) {
      // 添加树节点
      var trees = newChild;
      addTree(trees, this.params).then((response) => {
        if (response.code === 0) {
          this.trees = response.data.list;
        }
        setTimeout(() => {
          this.listLoading = false;
        }, 1.5 * 1000);
      });
    },

    append(data, value) {
      const newChild = { parent: data.id, name: value };
      if (!data.children) {
        data.children = [];
      }
      // data.children.push(newChild)
      this.addTrees(newChild);
    },

    editTree(data, value) {
      data.label = value;
      this.editTrees(data);
    },

    remove(node, data) {
      var trees = { id: data.id, isdel: 1 };
      this.delTrees(trees);
    },
    renderContent(h, { node, data, store }) {
      return h(
        "span",
        {
          class: "custom-tree-node",
        },
        h("span", null, node.label),
        h(
          "span",
          null,
          h(
            "a",
            {
              onClick: () => this.append(data),
            },
            "Append "
          ),
          h(
            "a",
            {
              onClick: () => this.remove(node, data),
            },
            "Delete"
          )
        )
      );
    },
    handleDragStart(node, ev) {},
    handleDragEnter(draggingNode, dropNode, ev) {},
    handleDragLeave(draggingNode, dropNode, ev) {},
    handleDragOver(draggingNode, dropNode, ev) {},
    handleDragEnd(draggingNode, dropNode, dropType, ev) {},
    handleDrop(draggingNode, dropNode, dropType, ev) {},
    handleClick(data) {
      this.$emit("getTreeID", data.id);
      treeID = data.id;
      if (Number(this.params.menutype) === 1 && Number(this.eventType) === 0) {
        window.localStorage.setItem("menuAddCase", data.id);
      }
      if (Number(this.params.menutype) === 1 && Number(this.eventType) === 1) {
        window.localStorage.setItem("menu", data.id);
      }
      if (Number(this.params.menutype) === 2 && Number(this.eventType) === 0) {
        window.localStorage.setItem("menuCaseAddStep", data.id);
      }
      if (Number(this.params.menutype) === 2 && Number(this.eventType) === 1) {
        window.localStorage.setItem("menuStep", data.id);
      }
      if (Number(this.params.menutype) === 3 && Number(this.eventType) === 0) {
        window.localStorage.setItem("menuTaskAddStep", data.id);
      }
      if (Number(this.params.menutype) === 3 && Number(this.eventType) === 1) {
        window.localStorage.setItem("menuCase", data.id);
      }
    },
  },
};
</script>

<script setup></script>

<style>
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

/*高亮当前选中的树节点*/
.el-tree-node.is-current > .el-tree-node__content {
  background-color: #66b1ff87;
}

/*树节点点击时的样式设置*/
.el-tree-node:focus > .el-tree-node__content {
  background-color: #66b1ff87;
}

/*树节点hover时的样式设置*/
.el-tree-node__content:hover {
  background-color: #26b1ff87;
}

.custom-tree-node {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-size: 16px;
}
</style>
