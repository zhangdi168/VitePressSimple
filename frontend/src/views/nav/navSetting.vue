<template>
  <empty-project></empty-project>
  <div v-if="!storeIndex.IsEmptyProject" class="flex justify-center mx-20 mt-4">
    <!--    选择操作的语言-->
    <div class="mx-2">
      <select-setting-lang></select-setting-lang>
    </div>
    <a-button
      v-if="storeConfig.IsUseI18n"
      class="bg-blue-100 mx-2 flex justify-center items-center hover:bg-blue-100"
      @click="copyNav()"
    >
      <icon-park
        class="mr-1"
        strokeLinejoin="bevel"
        theme="outline"
        type="copy-one"
      />
      {{ lang("pageNav.copy") }}
      <span class="text-red">{{ storeConfig.currSettingLang }}</span>
      {{ lang("pageNav.clipboard") }}
    </a-button>
    <a-button
      v-if="storeConfig.IsUseI18n"
      class="bg-green-200 mx-2 flex justify-center items-center hover:bg-blue-100"
      :disabled="!copyNavData"
      @click="cuttingNav()"
    >
      <q-tooltip> {{ lang("pageNav.coverCurrentNav") }}</q-tooltip>
      <icon-park
        class="mr-1"
        strokeLinejoin="bevel"
        theme="outline"
        type="cutting-one"
      />
      {{ lang("pageNav.paste") }}<span class="text-red">{{ copyNavLang }}</span
      >{{ lang("pageNav.toCurrentNav") }}
    </a-button>
    <!--    保存导航-->
    <div class="flex justify-end">
      <a-button
        class="bg-blue-200 mx-2 flex justify-center items-center hover:bg-blue-100"
        @click="addTopNav"
      >
        <icon-park
          class="mr-1"
          strokeLinejoin="bevel"
          theme="outline"
          type="add-one"
        />
        {{ lang("pageNav.addTopNav") }}
      </a-button>

      <a-button
        @click="saveNav()"
        class="bg-blue-600 mx-2 hover:bg-blue-500 text-white flex justify-center items-center"
      >
        <icon-park
          class="mr-1"
          strokeLinejoin="bevel"
          theme="outline"
          type="save"
        />
        {{ lang("pageNav.saveCurrentNav") }}
      </a-button>
    </div>
  </div>

  <hr class="my-3" />
  <div>
    <dy-add-nav
      ref="refNav"
      :show-add-top-nav="false"
      :level="1"
      v-model:nav-array="storeConfig.currLangConfig['themeConfig']['nav']"
    ></dy-add-nav>
  </div>
</template>
<script setup lang="ts">
import { onMounted, ref, watch } from "vue";
import { useVpconfigStore } from "@/store/vpconfig";
import DyAddNav from "@/components/dyAddNav.vue";
import { VpNav } from "@/utils/tree";
import { IconPark } from "@icon-park/vue-next/es/all";
import SelectSettingLang from "@/components/selectSettingLang.vue";
import { DeepClone } from "@/utils/deepClone";
import EmptyProject from "@/components/emptyProject.vue";
import { useIndexStore } from "@/store";
import { lang } from "../../utils/language";

const storeConfig = useVpconfigStore();
const storeIndex = useIndexStore();
onMounted(() => {
  // navList.value = storeConfig.configData["themeConfig"]["nav"];
  // console.log(navList.value, "navList -- console.log");
});

//新增顶级导航
const refNav = ref<InstanceType<typeof DyAddNav>>();

const addTopNav = () => {
  refNav.value?.addTopNav();
};
const copyNavData = ref();
const copyNavLang = ref("");
//复制导航
const copyNav = () => {
  copyNavLang.value = storeConfig.currSettingLang;
  copyNavData.value = DeepClone(
    storeConfig.currLangConfig["themeConfig"]["nav"],
  );
};
//粘贴导航
const cuttingNav = () => {
  storeConfig.currLangConfig["themeConfig"]["nav"] = copyNavData.value;
  copyNavData.value = null;
  copyNavLang.value = "";
};

const formatNavData = (data: VpNav[]) => {
  return data.map((item) => {
    const formattedItem: any = { text: item.text };
    if (item.items && item.items.length > 0) {
      // If 'items' exists and has a length greater than 0
      formattedItem.items = item.items.map((subItem) => ({
        text: subItem.text,
        link: subItem.items?.length ? undefined : subItem.link,
        // Recursive call to handle nested items
        items: subItem.items?.length ? formatNavData(subItem.items) : undefined,
      }));
    } else {
      // If 'items' doesn't exist or its length is 0
      formattedItem.link = item.link;
      formattedItem.text = item.text;
    }
    return formattedItem;
  });
};
const saveNav = () => {
  const formatData: VpNav[] = formatNavData(
    storeConfig.currLangConfig["themeConfig"]["nav"],
  );
  console.log(formatData, "formatData -- console.log");
  storeConfig.currLangConfig["themeConfig"]["nav"] = formatData;
  storeConfig.saveConfig();
};
</script>
<style scoped></style>
