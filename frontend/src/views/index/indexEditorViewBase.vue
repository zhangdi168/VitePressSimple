<template>
  <div>
    <div class="px-1 mb-2">
      <a-input
        v-model:value="storeIndex.currArticleFrontMatter['title']"
        :placeholder="lang('pageIndex.inputArticleTitle')"
        prefix=""
        :suffix="lang('common.title')"
      />
    </div>

    <div class="px-1 my-2">
      <a-textarea
        class="text-gray-500"
        v-model:value="storeIndex.currArticleFrontMatter['description']"
        :auto-size="{ minRows: 1 }"
        :placeholder="lang('pageIndex.inputPageSeoDescription')"
      ></a-textarea>
    </div>
    <div class="px-2 mb-2">
      <a-input
        v-model:value="storeIndex.currArticleFrontMatter['outline'][1]"
        class="text-gray-500"
        prefix=""
        :placeholder="lang('pageIndex.inputOutlineLevel')"
        :suffix="lang('pageIndex.outlineLevel')"
      />
    </div>
  </div>

  <div>
    <a-form
      :label-col="labelCol"
      :wrapper-col="wrapperCol"
      class="ml-5"
      label-align="left"
    >
      <div>
        <q-toggle
          :label="lang('pageIndex.showNav')"
          left-label
          v-model="storeIndex.currArticleFrontMatter['navbar']"
          color="blue"
        />
      </div>
      <div>
        <q-toggle
          :label="lang('pageIndex.showSidebar')"
          left-label
          v-model="storeIndex.currArticleFrontMatter['sideBar']"
          color="blue"
        />
      </div>
      <div>
        <q-toggle
          :label="lang('pageIndex.showFooter')"
          left-label
          v-model="storeIndex.currArticleFrontMatter['footer']"
          color="blue"
        />
      </div>

      <div>
        <q-toggle
          :label="lang('pageIndex.showEditLink')"
          left-label
          v-model="storeIndex.currArticleFrontMatter['editLink']"
          color="blue"
        />
      </div>
      <div>
        <q-toggle
          :label="lang('pageIndex.showUpdateTime')"
          left-label
          v-model="storeIndex.currArticleFrontMatter['lastUpdated']"
          color="blue"
        />
      </div>

      <a-form-item :label="lang('pageIndex.outlinePosition')">
        <a-radio-group
          v-model:value="storeIndex.currArticleFrontMatter['aside']"
        >
          <a-radio value="left">{{ lang("pageIndex.outlineLeft") }}</a-radio>
          <a-radio value="right">{{ lang("pageIndex.outlineRight") }}</a-radio>
        </a-radio-group>
      </a-form-item>
    </a-form>
  </div>
  <q-separator inset />

  <div class="mt-3">
    <dy-add-head
      v-model:meta="storeIndex.currArticleFrontMatter['head']"
      ref="refDyAddHead"
      :add-btn-text="lang('pageIndex.addMeta')"
      add-btn-class="bg-blue-500 text-white hover:bg-blue-600"
      class="mt-2"
      key-placeholder="name"
      value-placeholder="content"
      key-name="name"
      value-name="content"
    ></dy-add-head>
  </div>
  <div class="mt-3">
    <q-separator inset />
  </div>

  <!--  自定义formatter-->
  <div class="mt-3">
    <a-tooltip :title="lang('pageIndex.customFormatterTip')">
      <dy-add-head
        v-model:obj="storeIndex.currArticleFrontMatter['custom']"
        ref="refDyAddFontMatter"
        add-btn-class="bg-blue-500 text-white hover:bg-blue-600"
        :add-btn-text="lang('pageIndex.addCustomFormatter')"
        key-placeholder="key"
        value-placeholder="value"
      >
      </dy-add-head>
    </a-tooltip>
  </div>
</template>
<script setup lang="ts">
import DyAddHead from "@/components/dyAddKV.vue";
import { onMounted, ref } from "vue";
import { useIndexStore } from "@/store";
import { lang } from "@/utils/language";

const labelCol = { style: { width: "150px" } };
const wrapperCol = { span: 14 };
const storeIndex = useIndexStore();

const refDyAddHead = ref();
const refDyAddFontMatter = ref();
onMounted(() => {});
</script>
<style scoped></style>
