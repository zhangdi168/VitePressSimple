<template>
  <div>
    <div class="flex justify-start items-center">
      <div class="ml-1">{{ lang("pageIndex.pageType") }}</div>

      <q-radio
        v-model="storeIndex.currArticleFrontMatter['layout']"
        val="doc"
        :label="lang('pageIndex.pageTypeDoc')"
      />
      <q-radio
        v-model="storeIndex.currArticleFrontMatter['layout']"
        val="home"
        :label="lang('pageIndex.pageTypeHome')"
      />
      <q-radio
        v-model="storeIndex.currArticleFrontMatter['layout']"
        disable
        val="page"
        :label="lang('pageIndex.pageTypeCustom')"
      />
    </div>

    <div
      class="mt-2"
      v-if="storeIndex.currArticleFrontMatter['layout'] == 'home'"
    >
      <div class="px-1 mb-2">
        <a-input
          v-model:value="storeIndex.currArticleFrontMatter['hero']['name']"
          :placeholder="lang('pageIndex.hero.name')"
          prefix=""
          :suffix="lang('pageIndex.hero.name')"
        />
      </div>
      <div class="px-1 mb-2">
        <a-input
          v-model:value="storeIndex.currArticleFrontMatter['hero']['text']"
          :placeholder="lang('pageIndex.hero.text')"
          prefix=""
          :suffix="lang('pageIndex.hero.text')"
        />
      </div>
      <div class="px-1 mb-2">
        <a-input
          v-model:value="storeIndex.currArticleFrontMatter['hero']['tagline']"
          :placeholder="lang('pageIndex.hero.taglineExtra')"
          prefix=""
          :suffix="lang('pageIndex.hero.tagline')"
        />
      </div>
      <div class="my-3 flex justify-between">
        <div class="flex-1">
          <a-input
            disabled
            v-model:value="
              storeIndex.currArticleFrontMatter['hero']['image']['src']
            "
            :placeholder="lang('pageIndex.hero.image.src')"
            class="w-full"
          >
          </a-input>
        </div>
        <div class="mx-2">
          <a-button class="bg-blue-200" @click="selectLogo">
            <q-tooltip anchor="bottom left" self="bottom right"
              >{{ lang("pageIndex.hero.image.tips") }}
            </q-tooltip>
            {{ lang("pageIndex.hero.image.select") }}
          </a-button>
        </div>
      </div>
      <hr class="my-3" />
      <div>
        <dy-add-k-v
          v-model:objs="storeIndex.currArticleFrontMatter['hero']['actions']"
          :add-btn-text="lang('pageIndex.hero.actions.addBtnText')"
          :value-placeholder="lang('pageIndex.hero.actions.valuePlaceholder')"
          add-btn-class="bg-orange-200"
          :key-placeholder="lang('pageIndex.hero.actions.keyPlaceholder')"
          key-name="text"
          value-name="link"
        ></dy-add-k-v>
      </div>
      <hr class="my-3" />
      <div>
        <dy-add-k-v
          v-model:objs="storeIndex.currArticleFrontMatter['features']"
          :add-btn-text="lang('pageIndex.hero.features.addBtnText')"
          :value-placeholder="lang('pageIndex.hero.features.valuePlaceholder')"
          add-btn-class="bg-orange-200"
          :key-placeholder="lang('pageIndex.hero.features.keyPlaceholder')"
          key-name="title"
          value-name="details"
        ></dy-add-k-v>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import { useIndexStore } from "@/store";
import IndexPopAddFeatures from "@/views/index/indexPopAddFeatures.vue";
import { onBeforeMount, onMounted, ref } from "vue";
import DyAddKV from "@/components/dyAddKV.vue";
import {
  CopyPath,
  GetPathExt,
  GetPathFileName,
  PathJoin,
  SelectFile,
} from "../../../wailsjs/go/system/SystemService";
import { ToastCheck, ToastError } from "@/utils/Toast";
import { useVpconfigStore } from "@/store/vpconfig";
import { defaultFrontMatter } from "@/configs/defaultFrontMatter";
import { IsEmptyValue } from "@/utils/utils";
import { lang } from "../../utils/language";

const storeIndex = useIndexStore();
const storeVpConfig = useVpconfigStore();
onBeforeMount(() => {
  //初始化hero的默认值
  if (IsEmptyValue(storeIndex.currArticleFrontMatter["hero"])) {
    storeIndex.currArticleFrontMatter["hero"] = defaultFrontMatter.hero;
    return;
  }
  for (const key in defaultFrontMatter.hero) {
    if (IsEmptyValue(storeIndex.currArticleFrontMatter["hero"][key])) {
      storeIndex.currArticleFrontMatter["hero"][key] =
        defaultFrontMatter.hero[key];
    }
  }
  if (IsEmptyValue(storeIndex.currArticleFrontMatter["hero"]["image"]["src"]))
    storeIndex.currArticleFrontMatter["hero"]["image"]["src"] = "";
  if (IsEmptyValue(storeIndex.currArticleFrontMatter["hero"]["image"]["alt"]))
    storeIndex.currArticleFrontMatter["hero"]["image"]["alt"] = "";
});
const selectLogo = async () => {
  let oriImagePath = await SelectFile("选择主页图片", "");
  console.log(oriImagePath, "filePath -- console.log");
  let ext = await GetPathExt(oriImagePath);
  let allowExt = [".png", ".jpg", ".jpeg", ".bmp"];
  if (!allowExt.includes(ext)) {
    ToastError("请选则图片格式文件" + allowExt);
    return;
  }
  //当前文章的文件名
  let currFileName = await GetPathFileName(storeIndex.currArticlePath);
  currFileName = currFileName.replaceAll(".md", "");
  //组装新路径
  let publicDir = await PathJoin([storeVpConfig.fullSrcDir, "public"]);
  let newImagePath = await PathJoin([
    publicDir,
    "images",
    "home",
    currFileName + "_home" + ext,
  ]);
  let copyResult = await CopyPath(oriImagePath, newImagePath, false);
  ToastCheck(copyResult);
  storeIndex.currArticleFrontMatter["hero"]["image"]["src"] =
    newImagePath.replaceAll(publicDir, "");
};

const changeHome = () => {
  console.log(111, "111 -- console.log");
};
const refPopShowFeatures = ref();
const setFeatures = () => {
  refPopShowFeatures.value.showModal();
};
</script>
<style scoped></style>
