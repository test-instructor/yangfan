<template>
  <div>
    <codeEditor
        :debugTalkType="debugTalk"
    />
  </div>
</template>

<script>
export default {
  name: "debugtalkGen"
}
</script>

<script setup>
import codeEditor from '@/view/interface/debugtalk/debugtalk/codeEditor.vue'
import {ref} from "vue";
const debugTalk = ref(2)
</script>

<style scoped>

</style>
