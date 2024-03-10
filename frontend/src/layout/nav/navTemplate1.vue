<script setup lang="ts">
import { useNavStore } from "@/store/nav";
import { NavList } from "@/configs/navs";
import { useRouter } from "vue-router";
const navRouter = useNavStore();
const router = useRouter();
const routerClick = (path: string, name: string) => {
  navRouter.setActiveNav(path, name);
  router.push(path);
};
</script>

<template>
  <!-- Header -->
  <div class="header">
    <div class="flex bg-blue-400 items-center justify-between">
      <!-- 菜单 -->
      <div class="flex select-none pl-2 justify-start">
        <div
          v-for="item in NavList()"
          :key="item.name"
          :class="navRouter.ActiveName == item.name ? 'bg-blue-600' : ''"
          class="text-white text-base cursor-pointer mx-1 my-1 font-bold hover:bg-orange-500 px-3 py-1 rounded-lg"
          @click="routerClick(item.path, item.name)"
        >
          {{ item.title }}
        </div>
      </div>

      <!--      logo-->
      <div class="menu">
        <div class="mr-2">logo</div>
      </div>
    </div>
  </div>

  <!-- 路由 Page -->
  <div class="view">
    <router-view />
  </div>
</template>

<style scoped lang="scss"></style>
