<template>
  <!--  <div v-if="storeIndex.currArticlePath == ''">-->
  <!--    <a-empty description="未选择文章" />-->
  <!--  </div>-->
  <div v-if="storeIndex.currArticlePath != ''" style="height: 100vh">
    <a-tabs v-model:activeKey="activeKey" class="my-0.5 pl-1" type="card">
      <!-- tab   seo相关配置-->
      <a-tab-pane key="1" :tab="lang('common.base')">
        <index-editor-view-base></index-editor-view-base>
      </a-tab-pane>
      <a-tab-pane key="2" :tab="lang('common.home')">
        <index-editor-view-hero></index-editor-view-hero>
      </a-tab-pane>
      <a-tab-pane key="4" :tab="lang('common.code')">
        <div class="ml-1">
          <index-editor-view-vue></index-editor-view-vue>
        </div>
      </a-tab-pane>
      <!--    <a-tab-pane key="2" tab="fontmatter">font matter</a-tab-pane>-->
    </a-tabs>
    <div class="mt-6">
      <q-separator inset />
    </div>
  </div>
</template>
<script lang="ts" setup>
import DyAddHead from "../../components/dyAddKV.vue";
import { IconPark } from "@icon-park/vue-next/es/all";
import { nextTick, onMounted, reactive, ref, UnwrapRef } from "vue";
import IndexEditorViewVue from "./indexEditorViewVue.vue";
import { useIndexStore } from "../../store";
import { useLayoutStore } from "@/store/layout";
import IndexEditorViewBase from "@/views/index/indexEditorViewBase.vue";
import IndexEditorViewHero from "@/views/index/indexEditorViewHero.vue";
import { lang } from "@/utils/language";

const storeIndex = useIndexStore();
const storeLayout = useLayoutStore();
const activeKey = ref("1");
const savePageConfig = () => {};
</script>

<style scoped></style>
