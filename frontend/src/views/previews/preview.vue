<template>
  <div>
    <a-button @click="runCmd">执行命令</a-button>
    <a-button @click="stopCmd">停止命令</a-button>
  </div>
</template>
<script setup lang="ts">
import { ref } from "vue";
import { RunCmd, StopCmd } from "../../../wailsjs/go/services/ShellService";

const runCmd = () => {
  RunCmd("npm run docs:dev", "E:\\mydev\\vittepress-simple\\docs");
};
const stopCmd = () => {
  StopCmd();
};
</script>
<style scoped></style>
