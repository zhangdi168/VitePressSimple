<script lang="ts" setup>
import { IconPark } from "@icon-park/vue-next/es/all";
import { ref } from "vue";
import { NavList } from "@/configs/navs";
import { useNavStore } from "@/store/nav";
import { useRouter } from "vue-router";
const navStore = useNavStore();
const iconSize = ref<number>(28);

const router = useRouter();
const routerClick = (path: string, name: string) => {
  navStore.setActiveNav(path, name);
  router.push(path);
};
</script>

<template>
  <div class="flex bg-white justify-start items-start w-screen h-screen">
    <!-- 左侧导航 Start -->
    <div
      style="background-color: #ebebeb"
      class="flex flex-col w-14 justify-center items-center h-full overflow-hidden text-gray-700 rounded"
    >
      <!--    <a class="flex items-center justify-start mt-3" href="#">-->
      <!--      <icon-park :size="navSize" fill="#333" strokeLinecap="square" strokeLinejoin="bevel" theme="outline"  type="robot-one"/>-->
      <!--    </a>-->
      <div
        :key="item.path"
        v-for="item in NavList()"
        :class="item.borderTop ? 'mt-2 border-t' : 'mt-1'"
        class="flex flex-col items-center justify-center border-gray-300"
      >
        <router-link
          :class="navStore.ActiveName === item.name && 'bg-gray-400'"
          :to="item.path"
          class="flex items-center w-12 h-12 justify-center rounded hover:bg-gray-300"
          href="#"
          @click="routerClick(item.path, item.name)"
        >
          <q-tooltip v-if="item.title">{{ item.title }} </q-tooltip>
          <icon-park
            :size="iconSize"
            :type="item.icon"
            fill="#190e19"
            strokeLinejoin="bevel"
            theme="outline"
          />
        </router-link>
      </div>
      <a
        class="flex items-center w-12 h-12 justify-center mt-auto bg-gray-200 hover:bg-gray-300"
        href="#"
      >
        <icon-park
          :size="iconSize"
          :type="'user'"
          fill="#333"
          strokeLinejoin="bevel"
          theme="outline"
        />
      </a>
    </div>
    <!-- 左侧导航 End  -->

    <!-- 右侧主体 Start -->
    <div class="flex-1">
      <!-- 路由匹配到的组件将渲染在这里 -->
      <router-view></router-view>
    </div>
    <!-- 右侧主体 Start -->
  </div>
</template>

<style scoped></style>
