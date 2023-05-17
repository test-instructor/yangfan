<template>
  <div>
    <div id="codeEditor" :style="contentStyleObj"></div>
  </div>
</template>
<script>
import "jsoneditor/dist/jsoneditor.min.css";
import jsoneditor from "jsoneditor";

export default {
  data() {
    return {
      codeEditor: null,
      treeEditor: null,
      syncData: true,
      contentStyleObj: {
        height: "",
      },
    };
  },
  mounted: function () {
    // let treeEditorElement = document.getElementById('treeEditor')
    let codeEditorElement = document.getElementById("codeEditor");
    let json = {
      Array: [1, 2, 3],
      Boolean: true,
      Null: null,
      Number: 123,
      Object: { a: "b", c: "d" },
      String: "Hello World",
    };
    let codeOptions = {
      mode: "code",
      modes: ["code", "tree"],
    };
    // this.treeEditor = new jsoneditor(treeEditorElement, treeOptions, json);
    this.codeEditor = new jsoneditor(codeEditorElement, codeOptions, json);
  },
  methods: {
    getHeight() {
      this.contentStyleObj.height = window.innerHeight - 180 + "px";
    },
  },
  computed: {},
  created() {
    window.addEventListener("resize", this.getHeight);
    this.getHeight();
  },
};
</script>

<style>
.ace_editor,
.ace_editor * {
  font-family: "Monaco", "Menlo", "Ubuntu Mono", "Droid Sans Mono", "Consolas",
    monospace !important;
  font-size: 14px !important;
  font-weight: 400 !important;
  letter-spacing: 0 !important;
}
</style>
