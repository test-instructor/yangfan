<template>
  <div>
    <MonacoEditor
        :debugTalkType="debugTalk"
    />
  </div>
</template>

<script>
export default {
  name: "debugtalk"
}
</script>

<script setup>
import MonacoEditor from '@/view/interface/debugtalk/debugtalk/MonacoEditor.vue'
import {ref} from "vue";
const debugTalk = ref(1)

</script>

<style scoped>

</style>
