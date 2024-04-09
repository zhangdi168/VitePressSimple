<script setup lang="ts">
import { OpenURL } from "../../../wailsjs/go/system/SystemService";
</script>

<template>
  <div class="mx-20 mt-4 flex justify-start border p-3">
    <div class="w-full my-2">
      作者博客：<span
        class="cursor-pointer"
        @click="OpenURL('http://xiaod.co/')"
        >http://xiaod.co/</span
      >
    </div>
    <div class="w-full my-2">
      作者github：<span
        class="cursor-pointer"
        @click="OpenURL('https://github.com/zhangdi168')"
        >https://github.com/zhangdi168</span
      >
    </div>
    <div class="w-full my-2">
      作者gitee：<span
        class="cursor-pointer"
        @click="OpenURL('https://gitee.com/zhangdi168')"
        >https://gitee.com/zhangdi168</span
      >
    </div>
  </div>
</template>

<style scoped>
span {
  color: #007bff;
  text-decoration: none;
}

span:hover {
  text-decoration: underline;
}
</style>
