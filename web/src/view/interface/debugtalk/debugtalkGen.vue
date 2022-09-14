<template>
  <div class="wrapper">
    <el-row type="flex" justify="center">
      <el-col :span="14">
        <el-form :inline="true">
        </el-form>
        <div id="codeEditBox"></div>
        <el-button @click="handleCode">点击运行</el-button>
      </el-col>

    </el-row>
  </div>

</template>

<script>

import * as monaco from 'monaco-editor'
import { ref, toRaw } from 'vue'
import $ from 'jquery'
import { pythonCompletion } from "@/utils/completion";

export default{
  setup() {
    const editor = ref(null)
    const editorTheme = ref("vs-dark")
    const language = ref("java")
    const initEditor = () => {
      // 初始化编辑器，确保dom已经渲染
      editor.value = monaco.editor.create(document.getElementById('codeEditBox'), {
        value: '', //编辑器初始显示文字
        language: language.value, //语言支持自行查阅demo
        theme: editorTheme.value, //官方自带三种主题vs, hc-black, or vs-dark
        selectOnLineNumbers: true,//显示行号
        roundedSelection: false,
        readOnly: false, // 只读
        cursorStyle: 'line', //光标样式
        automaticLayout: true, //自动布局
        glyphMargin: true, //字形边缘
        useTabStops: false,
        fontSize: 15, //字体大小
        autoIndent: true, //自动布局
        quickSuggestionsDelay: 100, //代码提示延时
      });

      // 监听值的变化
      editor.value.onDidChangeModelContent((val) => {
        console.log("--------------",val.changes[0].text)
      })

      // 创建代码提醒
      pythonCompletion
    }
    $(document).ready(function () {
      initEditor()
    })
    const handleCode = () => {
      // console.log("==========", toRaw(editor.value).getValue())
    }

    return {
      editor,
      editorTheme,
      language,
      handleCode
    }
  },
}
</script>

<style scoped>
.wrapper {
  margin: 20px auto;
}
#codeEditBox {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  height: 400px;
  padding: 0;
  overflow: hidden;
  margin-bottom: 15px;
}
</style>
