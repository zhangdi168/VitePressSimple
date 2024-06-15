<template>
  <div>
    <div
      v-if="showAddTopSidebar"
      class="flex justify-start ml-5"
      :style="StyleNoDrag"
    >
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
        新增顶级侧栏
      </a-button>
    </div>
    <div v-for="(item, index) in SidebarArray" :key="index">
      <div class="flex mt-1 justify-start items-center" :style="marginLeft">
        <div class="mt-1 px-1 w-1/5">
          <a-input
            :style="StyleNoDrag"
            :class="hasChildren(index) ? '' : 'text-blue-600'"
            placeholder="侧栏文本"
            class="w-full"
            v-model:value="SidebarArray[index].text"
          />
        </div>
        <div class="mt-1 px-1 w-1/5">
          <a-input
            :style="StyleNoDrag"
            :class="hasChildren(index) ? '' : 'text-blue-600'"
            placeholder="跳转链接"
            :disabled="hasChildren(index)"
            class="w-full"
            v-model:value="SidebarArray[index].link"
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
        <div class="flex items-center mt-1 mx-1">
          <a-dropdown>
            <a-button type="dashed">
              <icon-park class="mr-1" theme="outline" type="move-one" />
              移动
            </a-button>
            <template #overlay>
              <a-menu>
                <a-menu-item>
                  <a href="javascript:;" @click="moveNav(index, 'forward')"
                    >↑上移</a
                  >
                </a-menu-item>
                <a-menu-item>
                  <a href="javascript:;" @click="moveNav(index, 'backward')"
                    >↓下移</a
                  >
                </a-menu-item>
              </a-menu>
            </template>
          </a-dropdown>
        </div>

        <div class="mt-1 px-1">
          <a-tooltip title="目前最高支持到第三级侧栏">
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
          <a-popconfirm
            title="将删除当前侧栏以及子侧栏，是否继续？"
            ok-text="Yes"
            cancel-text="No"
            :ok-button-props="{ danger: true }"
            ok-type="default"
            @confirm="confirmRemove"
            @cancel="cancel"
          >
            <a-button type="primary" @click="setCurrIndex(index)" danger ghost
              >删除
            </a-button>
          </a-popconfirm>
        </div>
      </div>
      <div
        v-if="
          SidebarArray[index].items &&
          SidebarArray[index].items.length > 0 &&
          isExpand(index)
        "
      >
        <dy-add-sidebar
          :level="props.level + 1"
          v-model:sidebar-array="SidebarArray[index].items"
        ></dy-add-sidebar>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import { IconPark } from "@icon-park/vue-next/es/all";
import { computed, ref } from "vue";
import { VpNav } from "@/utils/tree";
import { StyleNoDrag } from "@/configs/cnts";

interface Prop {
  level: number;
  showAddTopSidebar?: boolean;
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
//是否存在子侧栏(不存在说明是链接)
const hasChildren = (index: number) => {
  if (!SidebarArray.value[index]) return false;
  if (!SidebarArray.value[index].items) return false;
  if (!SidebarArray.value[index].items?.length) return false;
  return true;
};
const props = defineProps<Prop>();
const SidebarArray = defineModel<any[]>("sidebarArray", {
  required: true,
});
const marginLeft = computed(() => {
  return `margin-left:${(props.level - 1) * 30}px`;
});
//新增顶级侧栏
const addTopNav = () => {
  SidebarArray.value.push({ text: "", link: "", items: [] });
};
const removeNav = (index: number) => {
  SidebarArray.value.splice(index, 1);
};

const addSubNav = (index: number) => {
  isShowChildren.value[index] = true; //将当前父级展开
  if (!hasChildren(index)) {
    SidebarArray.value[index].items = [{ text: "", link: "" }];
  } else {
    SidebarArray.value[index].items?.push({ text: "", link: "" });
  }
};

const currentIndex = ref<number>(0);
const setCurrIndex = (index: number) => {
  currentIndex.value = index;
};
const confirmRemove = (e: MouseEvent) => {
  removeNav(currentIndex.value);
};

const cancel = (e: MouseEvent) => {
  console.log(e);
};

defineExpose({ addTopNav, addSubNav, removeNav, moveNav });

function moveNav(index: number, direction: "forward" | "backward") {
  // 检查索引和方向的合法性
  if (
    index < 0 ||
    index >= SidebarArray.value.length ||
    (direction !== "forward" && direction !== "backward")
  ) {
    return; // 不进行操作
  }

  let newIndex =
    direction === "forward"
      ? index === 0
        ? SidebarArray.value.length - 1
        : index - 1
      : index === SidebarArray.value.length - 1
        ? 0
        : index + 1;

  // 使用解构赋值交换两个元素的位置
  [SidebarArray.value[index], SidebarArray.value[newIndex]] = [
    SidebarArray.value[newIndex],
    SidebarArray.value[index],
  ];
}
</script>

<style scoped></style>
