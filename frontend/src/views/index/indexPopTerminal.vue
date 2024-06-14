<script setup lang="ts">
import ConfigInput from "@/components/configInput.vue";
import { ConfigKeySysTerminalDir } from "@/constant/keys/config";
</script>

<template>
  <div>
    <div class="flex">
      <config-input
        label="命令运行目录"
        :config-key="ConfigKeySysTerminalDir"
      ></config-input>
    </div>
    <div></div>
  </div>
</template>

<style scoped></style>
