<template>
  <div class="ml-5 mt-5">
    <config-switch
      label="是否开机自启:"
      tooltip="重启程序后生效"
      :config-key="ConfigKeyIsStartup"
    ></config-switch>
  </div>
</template>
<script lang="ts" setup>
import ConfigSwitch from "@/components/ConfigSwitch.vue";
import { ConfigKeyIsStartup } from "@/constant/keys/config";
</script>

<style scoped></style>
