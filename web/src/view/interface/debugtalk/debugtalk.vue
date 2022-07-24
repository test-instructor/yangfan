<template>
  <div>
    <codeEditor
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
import codeEditor from '@/view/interface/debugtalk/debugtalk/codeEditor.vue'
import {ref} from "vue";
const debugTalk = ref(1)

</script>

<style scoped>

</style>
