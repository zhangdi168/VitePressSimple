<template>
  <div v-if="!isEmptyArray(storeIndex.articleTreeData)" class="mt-1">
    {{ storeIndex.articleTreeData }}
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
            <q-item
              v-if="!title.endsWith('.md')"
              @click="showPopCreate(treeKey)"
              clickable
              v-close-popup
            >
              <q-item-section>新建文章</q-item-section>
            </q-item>
            <q-separator />
            <q-item
              clickable
              @click="showPopRename(title, treeKey)"
              v-close-popup
            >
              <q-item-section>重命名</q-item-section>
            </q-item>
            <q-separator />
            <!--            <q-item clickable v-close-popup>-->
            <!--              <q-item-section>移到根目录</q-item-section>-->
            <!--            </q-item>-->
            <q-separator />
            <q-item clickable @click="deletePath(treeKey)" v-close-popup>
              <q-item-section>删除</q-item-section>
            </q-item>
          </q-list>
        </q-menu>
        <div
          class="tree-node-wrapper flex items-start space-x-2"
          @click.prevent="handleClick(treeKey)"
        >
          <!-- 文字靠-->
          <div
            class="truncate"
            :class="
              title.endsWith('.md') ? 'text-base text-black' : 'text-gray-900'
            "
          >
            {{ title.replaceAll(".md", "") }}
          </div>
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
import { watch, ref, onMounted, createVNode, nextTick } from "vue";
import InputModel from "../../components/inputModel.vue";
import {
  FileTextOutlined,
  DownOutlined,
  ExclamationCircleOutlined,
} from "@ant-design/icons-vue";

import matter from "gray-matter";
import { IconPark } from "@icon-park/vue-next/es/all";
import {
  AntTreeNodeDragEnterEvent,
  AntTreeNodeDropEvent,
} from "ant-design-vue/es/tree";
import { useIndexStore } from "@/store";
import {
  DeletePath,
  ReadFileContent,
  Rename,
} from "../../../wailsjs/go/services/ArticleTreeData";
import { getParentDirectory, removeMdExtension } from "@/utils/file";
import { ToastError, ToastInfo } from "@/utils/Toast";
import { Modal } from "ant-design-vue";
import { isEmptyArray } from "@/utils/array";

const storeIndex = useIndexStore();
const moreIconShownKeys = ref<string[]>([]);
onMounted(async () => {
  nextTick(async () => {
    await storeIndex.loadTreeData();
  });
});

//拖拽移动的逻辑
const onDragEnter = (info: AntTreeNodeDragEnterEvent) => {
  console.log(info, "onDragEnter info -- console.log");

  // expandedKeys 需要展开时
  // expandedKeys.value = info.expandedKeys;
};

const onDrop = (info: AntTreeNodeDropEvent) => {
  console.log(info, "-----info value");
  storeIndex.moveTo(String(info.dragNode.key), String(info.node.key));
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

const deletePath = (key: string) => {
  Modal.confirm({
    content: "确定删除么？",
    okType: "danger",
    okText: "确定删除",
    icon: createVNode(ExclamationCircleOutlined),
    onOk() {
      DeletePath(key).then((res: string) => {
        if (res != "") {
          ToastError(res);
        } else {
          ToastInfo("删除成功");
          storeIndex.loadTreeData();
        }
      });
    },
    cancelText: "取消",
    onCancel() {
      Modal.destroyAll();
    },
  });
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
const currentCreateParentPath = ref<string>();
const showPopCreate = (path_: string) => {
  refModalCreate.value.showModal("");
  currentCreateParentPath.value = path_;
};
const onSubmitInputModalCreate = (value: string) => {};

//弹出层回调 重命名
const refModalRename = ref();
const currentRenamePath = ref<string>();
const showPopRename = (value: string, path_: string) => {
  refModalRename.value.showModal(removeMdExtension(value));
  currentRenamePath.value = path_;
};
const onSubmitInputModalRename = (value: string) => {
  let oldPath = currentRenamePath.value;
  let newPath = getParentDirectory(currentRenamePath.value ?? "") + "/" + value;
  if (oldPath?.endsWith(".md")) newPath = newPath + ".md";
  newPath = newPath.replace("//", "/");
  // console.log(currentRenamePath.value, "-----currentRenamePath value");
  // console.log(newPath, "-----newPath value");
  Rename(currentRenamePath.value ?? "", newPath).then(() => {
    storeIndex.loadTreeData();
  });
};
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
