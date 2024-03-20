<template>
  <div>
    <div v-if="showAddTopNav" class="flex justify-center ml-5">
      <a-button
        class="bg-blue-200 flex justify-center items-center hover:bg-blue-100"
        @click="addTopNav"
      >
        <icon-park
          class="mr-1"
          strokeLinejoin="bevel"
          theme="outline"
          type="add-one"
        />
        新增顶级导航
      </a-button>
    </div>
    <div v-for="(item, index) in NavArray" :key="index">
      <div class="flex mt-1 justify-center items-center" :style="marginLeft">
        <div class="mt-1 px-1 w-1/5">
          <a-input
            :class="hasChildren(index) ? '' : 'text-blue-600'"
            placeholder="导航文本"
            class="w-full"
            v-model:value="NavArray[index].text"
          />
        </div>
        <div class="mt-1 px-1 w-1/5">
          <a-input
            :class="hasChildren(index) ? '' : 'text-blue-600'"
            placeholder="跳转链接"
            :disabled="hasChildren(index)"
            class="w-full"
            v-model:value="NavArray[index].link"
          ></a-input>
        </div>
        <div class="flex items-center mt-1">
          <a-button
            :disabled="!hasChildren(index)"
            type="dashed"
            @click="toggleIndex(index)"
          >
            <icon-park
              v-if="isExpand(index)"
              class="mr-1"
              theme="outline"
              type="expand-up"
            />
            <icon-park
              v-if="!isExpand(index)"
              class="mr-1"
              theme="outline"
              type="expand-down"
            />
            {{ isExpand(index) ? "收起" : "展开" }}
          </a-button>
        </div>
        <div class="flex items-center mt-1">
          <a-button
            :disabled="!hasChildren(index)"
            type="dashed"
            @click="toggleIndex(index)"
          >
            <icon-park
              v-if="isExpand(index)"
              class="mr-1"
              theme="outline"
              type="expand-up"
            />
            <icon-park
              v-if="!isExpand(index)"
              class="mr-1"
              theme="outline"
              type="expand-down"
            />
            移动
          </a-button>
        </div>

        <div class="mt-1 px-1">
          <a-tooltip title="目前最高支持到第三级导航">
            <a-button
              :disabled="props.level == 3"
              type="primary"
              ghost
              @click="addSubNav(index)"
              >添加子级
            </a-button>
          </a-tooltip>
        </div>
        <div class="mt-1 px-1">
          <a-button type="primary" danger ghost @click="removeNav(index)"
            >删除
          </a-button>
        </div>
      </div>
      <div
        v-if="
          NavArray[index].items &&
          NavArray[index].items.length > 0 &&
          isExpand(index)
        "
      >
        <dy-add-nav
          :level="props.level + 1"
          v-model:nav-array="NavArray[index].items"
        ></dy-add-nav>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import { IconPark } from "@icon-park/vue-next/es/all";
import { computed, ref } from "vue";
import { VpNav } from "@/utils/tree";

interface Prop {
  level: number;
  showAddTopNav?: boolean;
}

const isShowChildren = ref<Record<number, boolean>>({});
const isExpand = (index: number) => {
  if (isShowChildren.value[index] === undefined) return false; //默认都为折叠
  return isShowChildren.value[index];
};
const toggleIndex = (index: number) => {
  if (isExpand(index)) {
    isShowChildren.value[index] = false;
  } else {
    isShowChildren.value[index] = true;
  }
};
//是否存在子导航(不存在说明是链接)
const hasChildren = (index: number) => {
  if (!NavArray.value[index]) return false;
  if (!NavArray.value[index].items) return false;
  if (!NavArray.value[index].items?.length) return false;
  return true;
};
const props = defineProps<Prop>();
const NavArray = defineModel<VpNav[]>("navArray", {
  required: true,
});
const marginLeft = computed(() => {
  return `margin-left:${(props.level - 1) * 30}px`;
});
//新增顶级导航
const addTopNav = () => {
  NavArray.value.push({ text: "", link: "", items: [] });
};
const removeNav = (index: number) => {
  NavArray.value.splice(index, 1);
};

const addSubNav = (index: number) => {
  if (!NavArray.value[index].items) {
    NavArray.value[index].items = [{ text: "", link: "" }];
  } else {
    NavArray.value[index].items?.push({ text: "", link: "" });
  }
};
</script>

<style scoped></style>
