<script setup lang="ts">
import { OpenURL } from "../../../wailsjs/go/system/SystemService";
import { lang } from "../../utils/language";
import { StyleNoDrag } from "@/configs/cnts";
</script>

<template>
  <div class="mx-20 mt-4 flex justify-start border p-3">
    <div class="w-full my-2">
      {{ lang("pageAbout.authorBlog")
      }}<span class="cursor-pointer" @click="OpenURL('http://xiaod.co/')"
        >http://xiaod.co/</span
      >
    </div>
    <div class="w-full my-2">
      {{ lang("pageAbout.authorGithub")
      }}<span
        class="cursor-pointer"
        @click="OpenURL('https://github.com/zhangdi168')"
        >https://github.com/zhangdi168</span
      >
    </div>
    <div class="w-full my-2">
      {{ lang("pageAbout.authorGitee")
      }}<span
        class="cursor-pointer"
        @click="OpenURL('https://gitee.com/zhangdi168')"
        >https://gitee.com/zhangdi168</span
      >
    </div>
    <div class="w-full my-2">
      {{ lang("pageAbout.authorWx")
      }}<span
        class="cursor-pointer"
        @click="OpenURL('https://gitee.com/zhangdi168')"
        >bcxiaod</span
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
