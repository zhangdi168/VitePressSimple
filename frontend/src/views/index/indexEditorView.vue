<script lang="ts" setup>
import DyAddHead from "../../components/dyAddKV.vue";
import { IconPark } from "@icon-park/vue-next/es/all";
import { nextTick, onMounted, reactive, ref, UnwrapRef } from "vue";
import IndexEditorViewVue from "./indexEditorViewVue.vue";
import { useIndexStore } from "../../store";
import { useLayoutStore } from "@/store/layout";
import IndexEditorViewBase from "@/views/index/indexEditorViewBase.vue";

const storeIndex = useIndexStore();
const storeLayout = useLayoutStore();
const activeKey = ref("1");
</script>

<template>
  <!--  <div v-if="storeIndex.currArticlePath == ''">-->
  <!--    <a-empty description="未选择文章" />-->
  <!--  </div>-->
  <div v-if="storeIndex.currArticlePath != ''" style="height: 100vh">
    <a-tabs v-model:activeKey="activeKey" class="my-0.5 pl-1" type="card">
      <!-- tab   seo相关配置-->
      <a-tab-pane key="1" tab="基础">
        <index-editor-view-base></index-editor-view-base>
      </a-tab-pane>

      <a-tab-pane key="4" tab="js">
        <div class="ml-1">
          <index-editor-view-vue></index-editor-view-vue>
        </div>
      </a-tab-pane>
      <!--    <a-tab-pane key="2" tab="fontmatter">font matter</a-tab-pane>-->
    </a-tabs>
    <div class="mt-6">
      <q-separator inset />
    </div>

    <div class="flex justify-end pr-3 pt-2">
      <a-button
        class="bg-blue-600 hover:bg-blue-500 text-white flex justify-center items-center"
      >
        <icon-park
          class="mr-1"
          strokeLinejoin="bevel"
          theme="outline"
          type="save"
        />
        保存
      </a-button>
    </div>
  </div>
</template>

<style scoped></style>
