<template>
  <div style="margin-bottom: 20px">
    <el-button type="primary"  @click="handleCode" round>
      点击保存
    </el-button>
  </div>
  <div class="wrapper">
    <el-row type="flex" >
      <el-col :span="14">
        <el-form :inline="true">
        </el-form>
        <div id="codeEditBox"></div>
      </el-col>

    </el-row>
  </div>

</template>

<script>
export default {
  name: "MonacoEditor"
}
</script>

<script setup>
import * as monaco from 'monaco-editor'
import { ref, toRaw } from 'vue'
import $ from 'jquery'
import {getDebugTalk, updateDebugTalk} from "@/api/interfaceTemplate";
import {ElMessage} from "element-plus";
import { language as pythonLanguage } from 'monaco-editor/esm/vs/basic-languages/python/python.js';
const editor = ref(null)
const editorTheme = ref("vs-dark")
const language = ref("python")
const content = ref("")

const props = defineProps({
  debugTalkType: ref(),
})


const initEditor = () => {

  // 初始化编辑器，确保dom已经渲染
  editor.value = monaco.editor.create(document.getElementById('codeEditBox'), {
    value: content.value, //编辑器初始显示文字
    language: language.value, //语言支持自行查阅demo
    theme: editorTheme.value, //官方自带三种主题vs, hc-black, or vs-dark
    selectOnLineNumbers: true,//显示行号
    roundedSelection: false,
    readOnly: false, // 只读
    cursorStyle: 'line', //光标样式
    automaticLayout: true, //自动布局
    glyphMargin: false, //字形边缘
    useTabStops: false,
    fontSize: 14, //字体大小
    autoIndent: true, //自动布局
    quickSuggestionsDelay: 100, //代码提示延时
  });

  // 监听值的变化
  // editor.value.onDidChangeModelContent((val) => {
  //   console.log("--------------",val.changes[0].text)
  // })

  // 创建代码提醒
  const pythonCompletion = monaco.languages.registerCompletionItemProvider('python', {
    provideCompletionItems: function () {
      let suggestions = [];
      pythonLanguage.keywords.forEach(item => {
        suggestions.push({
          label: item,
          kind: monaco.languages.CompletionItemKind.Keyword,
          insertText: item
        });
      })
      return {
        suggestions:suggestions
      };
    },
  });
}


const handleCode = async () => {
  // console.log("==========", toRaw(editor.value).getValue())
  let updateData = {file_type: props.debugTalkType, content: toRaw(editor.value).getValue()}
  const res = await updateDebugTalk(updateData)
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '更新成功'
    })

  }
}

const init = async () => {
  let data = {file_type:props.debugTalkType}
  const res = await getDebugTalk(data)
  if (res.code === 0) {
    content.value = res.data.reapicase.content
    $(document).ready(function () {
      initEditor()
    })
  }
}

init()

</script>

<style scoped>
.wrapper {
  margin: 20px auto;
}
#codeEditBox {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  height: 720px;
  width: 170%;
  padding: 0;
  overflow: hidden;
  margin-bottom: 15px;
}
</style>
