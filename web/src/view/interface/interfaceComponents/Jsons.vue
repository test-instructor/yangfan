<template>
  <div>
    <div id="codeEditor" :style="contentStyleObj" @keyup="jsonDatas"></div>
  </div>
</template>

<script>
import "jsoneditor/dist/jsoneditor.min.css";
import jsoneditor from "jsoneditor";
import { ref } from "vue";

export default {
  props: {
    request: {
      require: false,
    },
    heights: ref(),
    jsons: ref(),
  },
  name: "Jsons",
  components: {
    jsoneditor,
  },

  methods: {
    editorInit() {
      require("brace/ext/language_tools");
      require("brace/mode/json");
      require("brace/theme/github");
      require("brace/snippets/json");
    },
    getHeight() {
      this.contentStyleObj.height = this.height + "px";
      this.contentStyleObj.width = "98%";
    },
    jsonDatas() {
      this.$emit("requestJsonData", this.codeEditor.get());
    },
  },

  data() {
    return {
      codeEditor: null,
      contentStyleObj: {
        height: "",
        width: "",
      },
      timeStamp: "",
    };
  },
  computed: {
    height() {
      return this.heights - 70;
    },
  },
  mounted: function () {
    let codeOptions = {
      mode: "code",
      modes: ["code", "tree"],
    };
    let codeEditorElement = document.getElementById("codeEditor");
    let json = {};
    if (this.jsons !== "") {
      json = this.jsons;
    }
    this.codeEditor = new jsoneditor(codeEditorElement, codeOptions, json);
    this.jsonDatas();
  },
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
