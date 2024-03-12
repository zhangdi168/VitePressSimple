<script lang="ts" setup>
import DyAddHead from "../../components/dyAddKV.vue";
import { IconPark } from "@icon-park/vue-next/es/all";
import { nextTick, onMounted, reactive, ref, UnwrapRef } from "vue";
import IndexEditorViewVue from "./indexEditorViewVue.vue";
import { useIndexStore } from "../../store";

const storeIndex = useIndexStore();
const activeKey = ref("1");
const labelCol = { style: { width: "150px" } };
const wrapperCol = { span: 14 };
const pageSetting = reactive({
  navbar: true, //是否显示导航
  sideBar: true, //是否显示侧栏
  footer: true, //是否显示页脚
  outline: 2, //默认大纲显示标题数
  editLink: true, //是否显示编辑链接
  lastUpdated: true, //是否显示页脚更新时间
  aside: "left", //侧边栏位置
});
</script>

<template>
  <!--  <div v-if="storeIndex.currArticlePath == ''">-->
  <!--    <a-empty description="未选择文章" />-->
  <!--  </div>-->
  <div v-if="storeIndex.currArticlePath != ''">
    <a-tabs v-model:activeKey="activeKey" class="my-0.5 pl-1" type="card">
      <!-- tab   seo相关配置-->
      <a-tab-pane key="1" tab="SEO">
        <div>
          <div class="px-2 mb-2">
            <a-input placeholder="请输入文章标题" prefix="" suffix="标题" />
          </div>
          <div class="px-2 my-2">
            <a-textarea placeholder="请输入seo页面描述"></a-textarea>
          </div>

          <div class="mt-4">
            <dy-add-head
              ref="dyAddHead"
              add-btn-text="新增meta "
              class="mt-2"
              key-placeholder="meta-name"
              value-placeholder="meta-content"
            ></dy-add-head>
          </div>
        </div>
      </a-tab-pane>
      <!--    tab 页面属性-->
      <a-tab-pane key="2" tab="页面">
        <a-form
          :label-col="labelCol"
          :wrapper-col="wrapperCol"
          class="ml-5"
          label-align="left"
        >
          <a-form-item label="是否显示导航">
            <a-switch v-model:checked="pageSetting.navbar" />
          </a-form-item>
          <a-form-item label="是否显示侧边栏">
            <a-switch v-model:checked="pageSetting.sideBar" />
          </a-form-item>
          <a-form-item label="是否显示页脚">
            <a-switch v-model:checked="pageSetting.footer" />
          </a-form-item>
          <a-form-item label="默认大纲显示标题数">
            <a-input-number
              v-model:value="pageSetting.outline"
              :max="10"
              :min="1"
            />
          </a-form-item>
          <a-form-item label="是否显示编辑链接">
            <a-switch v-model:checked="pageSetting.editLink" />
          </a-form-item>
          <a-form-item label="是否显示页脚更新时间">
            <a-switch v-model:checked="pageSetting.lastUpdated" />
          </a-form-item>
          <a-form-item label="侧边栏位置">
            <a-radio-group v-model:value="pageSetting.aside">
              <a-radio value="left">左侧</a-radio>
              <a-radio value="right">右侧</a-radio>
            </a-radio-group>
          </a-form-item>
          <a-form-item label="是否显示页脚更新时间">
            <a-switch v-model:checked="pageSetting.lastUpdated" />
          </a-form-item>
        </a-form>
      </a-tab-pane>

      <!--tab    fontmatter-->
      <a-tab-pane key="3" tab="fontmatter">
        <dy-add-head
          ref="dyAddFontMatter"
          add-btn-text="新增自定义fontMatter"
          class="mb-2"
          key-placeholder="key"
          value-placeholder="value"
        ></dy-add-head>
      </a-tab-pane>

      <a-tab-pane key="4" tab="js">
        <div class="ml-1">
          <index-editor-view-vue></index-editor-view-vue>
        </div>
      </a-tab-pane>
      <!--    <a-tab-pane key="2" tab="fontmatter">font matter</a-tab-pane>-->
    </a-tabs>

    <div class="flex mt-6 border-t justify-end pr-3 pt-2">
      <a-button
        class="bg-green-200 flex justify-center items-center hover:bg-blue-100"
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
