<template>
  <!--  使用界面模板1-->
  <nav-template1 v-if="useTemplateIndex === 1"></nav-template1>
  <!--  使用界面模板2-->
  <nav-template2 v-if="useTemplateIndex === 2"></nav-template2>
</template>
<script lang="ts" setup>
import NavTemplate1 from "@/layout/nav/navTemplate1.vue";
import NavTemplate2 from "@/layout/nav/navTemplate2.vue";
import { onMounted, ref } from "vue";
import { useVpconfigStore } from "@/store/vpconfig";
//在这里可以设置默认的模板
const useTemplateIndex = ref(2);
</script>
<style lang="scss"></style>
