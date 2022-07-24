<template>
  <div style="margin-bottom: 20px">
    <el-button type="primary"  @click="handleDebugTalk" round>
      点击保存
    </el-button>
  </div>
  <codemirror
      v-model="code"
      placeholder="Code goes here..."
      style="height:700px"
      :autofocus="true"
      :indent-with-tab="true"
      tab-size=4
      :extensions="extensions"
      @change="debugTalkChange($event)"
  />
<!--  @ready="log('ready', $event)"-->
<!--  @change="log('change', $event)"-->
<!--  @focus="log('focus', $event)"-->
<!--  @blur="log('blur', $event)"-->
</template>

<script setup>
import {
  updateDebugTalk,
  getDebugTalk
} from '@/api/interfaceTemplate'
import { Codemirror } from 'vue-codemirror'
import { python } from '@codemirror/lang-python'
import { oneDark } from '@codemirror/theme-one-dark'
import { ref, computed, watch } from 'vue'
import {ElMessage} from "element-plus";

const props = defineProps({
  debugTalkType: ref(),
})

const code = ref(``)
const extensions = [python(), oneDark]
let content

const init = async () => {

  let data = {file_type:props.debugTalkType}
  const res = await getDebugTalk(data)
  if (res.code === 0) {
    code.value = res.data.reapicase.content
    content = res.data.reapicase.content
  }
}

const handleDebugTalk = async () => {
  let updateData = {file_type: props.debugTalkType, content: content}
  const res = await updateDebugTalk(updateData)
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '更新成功'
    })

  }
}




init()

const debugTalkChange = (event) => {
  content = event
}




</script>

<style scoped>

</style>
