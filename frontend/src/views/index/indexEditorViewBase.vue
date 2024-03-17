<template>
  <div>
    <div class="px-1 mb-2">
      <a-input
        :value="storeIndex.GetArticleFrontMatter['title']"
        placeholder="请输入文章标题"
        prefix=""
        suffix="标题"
      />
    </div>

    <div class="px-1 my-2">
      <a-textarea
        class="text-gray-500"
        :value="storeIndex.GetArticleFrontMatter['description']"
        :auto-size="{ minRows: 1 }"
        placeholder="请输入页面seo描述"
      ></a-textarea>
    </div>
    <div class="px-2 mb-2">
      <a-input
        :value="storeIndex.GetArticleFrontMatter['outline']"
        placeholder="选填：1-6"
        class="text-gray-500"
        prefix=""
        suffix="大纲显示级别(1-6)"
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
          label="是否显示导航:"
          left-label
          v-model="storeIndex.GetArticleFrontMatter['navbar']"
          color="blue"
        />
      </div>
      <div>
        <q-toggle
          label="是否显示侧边栏:"
          left-label
          v-model="storeIndex.GetArticleFrontMatter['sideBar']"
          color="blue"
        />
      </div>
      <div>
        <q-toggle
          label="是否显示页脚:"
          left-label
          v-model="storeIndex.GetArticleFrontMatter['footer']"
          color="blue"
        />
      </div>

      <div>
        <q-toggle
          label="是否显示编辑链接:"
          left-label
          v-model="storeIndex.GetArticleFrontMatter['editLink']"
          color="blue"
        />
      </div>
      <div>
        <q-toggle
          label="是否显示更新时间:"
          left-label
          v-model="storeIndex.GetArticleFrontMatter['lastUpdated']"
          color="blue"
        />
      </div>

      <a-form-item label="侧边栏位置">
        <a-radio-group
          v-model:value="storeIndex.GetArticleFrontMatter['aside']"
        >
          <a-radio value="left">左侧</a-radio>
          <a-radio value="right">右侧</a-radio>
        </a-radio-group>
      </a-form-item>
    </a-form>
  </div>
  <q-separator inset />

  <div class="mt-3">
    <dy-add-head
      :default-value="storeIndex.GetArticleFrontMatter['head']['meta']"
      ref="refDyAddHead"
      add-btn-text="新增meta"
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
    <a-tooltip
      title="变量数据存放于custom属性中，通过{frontmatter}.custom.{xx}进行访问"
    >
      <dy-add-head
        :default-value="storeIndex.GetArticleFrontMatter['custom']"
        ref="refDyAddFontMatter"
        add-btn-class="bg-blue-500 text-white hover:bg-blue-600"
        add-btn-text="新增自定义fontMatter"
        key-placeholder="key"
        value-placeholder="value"
      >
      </dy-add-head>
    </a-tooltip>
  </div>
</template>
<script setup lang="ts">
import DyAddHead from "@/components/dyAddKV.vue";
import { nextTick, onMounted, ref } from "vue";
import { useIndexStore } from "@/store";
import { useLayoutStore } from "@/store/layout";

const labelCol = { style: { width: "150px" } };
const wrapperCol = { span: 14 };
const storeIndex = useIndexStore();
const layoutStore = useLayoutStore();
const refDyAddHead = ref();
const refDyAddFontMatter = ref();
onMounted(() => {
  nextTick(() => {
    layoutStore.setComponentDyAddHeader(refDyAddHead.value);
    layoutStore.setComponentDyAddCustomFrontMatter(refDyAddFontMatter.value);
  });
});
</script>
<style scoped></style>
