<template>
  <div class="mt-1">
    <a-tree
      v-model:expandedKeys="expandedKeys"
      v-model:selected-keys="selectKeys"
      :tree-data="storeIndex.ArticleTreeData"
      block-node
      class="mt-1.5 text-lg .ant-tree-treenode-selected customized-tree"
      draggable
      @dragenter="onDragEnter"
      @drop="onDrop"
    >
      <template #title="{ node, key: treeKey, title }">
        <!--     右键菜单-->
        <q-menu touch-position context-menu>
          <q-list dense style="min-width: 100px">
            <q-item v-if="!title.endsWith('.md')" clickable v-close-popup>
              <q-item-section>新建文章</q-item-section>
            </q-item>
            <q-separator />
            <q-item clickable v-close-popup>
              <q-item-section>重命名</q-item-section>
            </q-item>
            <q-separator />
            <q-item clickable v-close-popup>
              <q-item-section>剪切</q-item-section>
            </q-item>
            <q-separator />
            <q-item clickable v-close-popup>
              <q-item-section>删除</q-item-section>
            </q-item>
          </q-list>
        </q-menu>
        <div
          class="tree-node-wrapper flex items-start space-x-2"
          @click.prevent="handleClick(treeKey)"
        >
          <!-- 文字靠-->
          <div>{{ title.replaceAll(".md", "") }}</div>
        </div>
      </template>
      <template #switcherIcon="{ switcherCls }">
        <down-outlined :class="switcherCls" />
      </template>

      <template #icon="{ key }">
        <template v-if="key.endsWith('.md')">
          <FileTextOutlined />
        </template>
      </template>
    </a-tree>

    <input-model
      ref="refModalCreate"
      @submit-input-modal="onSubmitInputModalCreate"
    ></input-model>
    <input-model
      ref="refModalRename"
      @submit-input-modal="onSubmitInputModalRename"
    ></input-model>
  </div>
</template>
<script lang="ts" setup>
import { watch, ref, onMounted } from "vue";
import InputModel from "../../components/inputModel.vue";
import {
  FileTextOutlined,
  FileTextTwoTone,
  DownOutlined,
} from "@ant-design/icons-vue";

import matter from "gray-matter";
import { IconPark } from "@icon-park/vue-next/es/all";
import {
  AntTreeNodeDragEnterEvent,
  AntTreeNodeDropEvent,
} from "ant-design-vue/es/tree";
import { useIndexStore } from "@/store";
import { ReadFileContent } from "../../../wailsjs/go/services/ArticleTreeData";

const storeIndex = useIndexStore();
const moreIconShownKeys = ref<string[]>([]);
onMounted(async () => {
  await storeIndex.loadTreeData();
});

//拖拽移动的逻辑
const onDragEnter = (info: AntTreeNodeDragEnterEvent) => {
  console.log(info, "onDragEnter info -- console.log");

  // expandedKeys 需要展开时
  // expandedKeys.value = info.expandedKeys;
};

const onDrop = (info: AntTreeNodeDropEvent) => {
  console.log(info, "onDrop info -- console.log");
};

const showMoreIcon = (treeKey: string) => {
  moreIconShownKeys.value.push(treeKey);
};
const hideMoreIcon = (treeKey: string) => {
  const index = moreIconShownKeys.value.indexOf(treeKey);
  if (index > -1) {
    moreIconShownKeys.value.splice(index, 1);
  }
};

//双击标题展开菜单
const handleClick = (key: string) => {
  if (isDir(key)) {
    if (expandedKeys.value.includes(key)) {
      //删除
      expandedKeys.value = expandedKeys.value.filter((s) => s !== key);
    } else {
      //添加
      expandedKeys.value.push(key); // 或者添加逻辑来控制是否展开/收起
    }
  } else {
    //文件
    ReadFileContent(key).then((content: string) => {
      // console.log(content, "读取文件内容成功")
      let matterData = matter(content);

      // console.log(matterData, "matterData -- console.log")
      storeIndex.setCurrArticlePath(key);
      // //matter字符串
      storeIndex.setCurrArticleFrontMatter(matterData.data);

      let val = matterData.content
        ? matterData.content
        : "# hello vitePress client";
      console.log("val:" + val);
      storeIndex.Vditor?.setValue(val);
    });
  }
};

const isDir = (key: string) => {
  return !key.includes(".md");
};

const expandedKeys = ref<string[]>([]);
const selectKeys = ref<string[]>([]);

watch(expandedKeys, () => {
  console.log("expandedKeys", expandedKeys);
});

//弹出层回调 新建
const refModalCreate = ref();
const onSubmitInputModalCreate = (value: string) => {};

//弹出层回调 重命名
const refModalRename = ref();
const onSubmitInputModalRename = (value: string) => {};
</script>

<style scoped>
.tree-node-wrapper {
  display: flex;
  align-items: center;
  width: 100%;
  height: 100%;
}

.tree-node-wrapper {
  display: flex;
  align-items: center; /* 保持垂直居中 */
}

/* 确保图标始终在右侧 */
.tree-node-wrapper > div:last-child {
  justify-self: flex-end;
}
</style>
