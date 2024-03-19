<template>
  <!--  <div class="text-h4 mb-2 ml-2 text-center">导航配置</div>-->
  <q-splitter :model-value="modelValue">
    <template v-slot:before>
      <div class="border px-6" style="min-height: 80vh">
        <div class="mt-3">
          <a-tree
            class="draggable-tree"
            draggable
            block-node
            v-model:selectedKeys="selectedKeys"
            :tree-data="gData"
            @dragenter="onDragEnter"
            @drop="onDrop"
          />
        </div>
      </div>
    </template>
    <template v-slot:after>
      <div class="border px-6" style="min-height: 80vh">
        <a-tabs centered class="mt-3" v-model:activeKey="activeKey" type="card">
          <a-tab-pane disabled key="1" tab="新增顶级导航">
            <sim-input
              v-model="inputTopTitle"
              tooltip=""
              label="一级导航文本"
              is-full-width
            ></sim-input>
            <sim-input
              v-model="inputTopLink"
              placeholder="如果该导航不是链接，则该项为空不填"
              tooltip="如果该导航不是链接，则该项为空不填"
              label="顶级导航链接"
              is-full-width
            ></sim-input>
            <div class="flex justify-end items-center">
              <a-button @click="addTopNode()" class="mx-1" type="primary" ghost
                >新增顶级导航</a-button
              >
            </div></a-tab-pane
          >
          <a-tab-pane disabled key="2" tab="操作选中导航" force-render>
            <div v-show="selectedKeys.length > 0">
              <sim-input
                v-model="inputSelectTitle"
                tooltip=""
                label="当前选择导航文本"
                is-full-width
              ></sim-input>
              <sim-input
                v-model="inputSelectLink"
                tooltip="如果该导航不是链接，则该项为空不填"
                label="当前选择导航连接"
                is-full-width
              ></sim-input>
              <div class="flex justify-end items-center">
                <a-button
                  class="mx-1"
                  @click="deleteSelectNav()"
                  type="primary"
                  danger
                  >删除当前选中导航</a-button
                >

                <a-button
                  @click="
                    TreeUtil.updateNodeByKey(gData, selectedKeys[0], {
                      title: inputSelectTitle,
                      link: inputSelectLink,
                      children: currSelectNode.children,
                    } as DataNode)
                  "
                  class="mx-1"
                  type="primary"
                  ghost
                  >修改选中</a-button
                >
              </div>
            </div></a-tab-pane
          >
        </a-tabs>
      </div>
    </template>
  </q-splitter>
</template>
<script lang="ts" setup>
import { computed, ref, watch } from "vue";
import type {
  AntTreeNodeDragEnterEvent,
  AntTreeNodeDropEvent,
  DataNode,
  TreeProps,
} from "ant-design-vue/es/tree";
import SimInput from "@/components/simInput.vue";
import { IsEmptyValue } from "@/utils/utils";
import { ToastError } from "@/utils/Toast";
import { TreeUtil } from "@/utils/tree";

const x = 3;
const y = 2;
const z = 1;
const genData: DataNode[] | undefined = [];

type TreeDataItem = TreeProps["treeData"][number];
const gData = ref<DataNode[]>([]);
const inputTopTitle = ref("");
const inputTopLink = ref("");
const inputSelectTitle = ref();
const inputSelectLink = ref();
const modelValue = ref(50);
const selectedKeys = ref<string[]>([]);
const activeKey = ref("1");
const currSelectNode = ref<DataNode>({} as DataNode);
watch(selectedKeys, (val) => {
  if (val.length > 0) {
    currSelectNode.value = TreeUtil.findNodeByKey(
      gData.value,
      selectedKeys.value[0],
    );
    activeKey.value = "2";
    console.log(currSelectNode.value, "-----currSelectNode value");
    inputSelectTitle.value = currSelectNode.value.title;
    inputSelectLink.value = currSelectNode.value.key;
  } else {
    activeKey.value = "1";
  }
});

const addTopNode = () => {
  if (IsEmptyValue(inputTopTitle.value)) {
    ToastError("请输入导航文本");
    return;
  }
  gData.value = gData.value ? gData.value : [];
  gData.value?.push({
    title: inputTopTitle.value,
    key: inputTopLink.value,
  });
  console.log(gData.value, "-----gData value");
};

const deleteSelectNav = () => {
  if (TreeUtil.deleteNodeByKey(gData.value, selectedKeys.value[0])) {
    selectedKeys.value = [];
  }
};
const onDragEnter = (info: AntTreeNodeDragEnterEvent) => {
  console.log(info);
};

const onDrop = (info: AntTreeNodeDropEvent) => {
  console.log(info);
  const dropKey = info.node.key;
  const dragKey = info.dragNode.key;
  const dropPos = info.node.pos?.split("-");
  const dropPosition =
    info.dropPosition - Number(dropPos ? dropPos[dropPos.length - 1] : 0);
  const loop = (
    data: TreeProps["treeData"],
    key: string | number,
    callback: any,
  ) => {
    data?.forEach((item, index) => {
      if (item.key === key) {
        return callback(item, index, data);
      }
      if (item.children) {
        return loop(item.children, key, callback);
      }
    });
  };
  const data = [...gData.value];

  // Find dragObject
  let dragObj: TreeDataItem;
  loop(
    data,
    dragKey,
    (item: TreeDataItem, index: number, arr: TreeProps["treeData"]) => {
      arr?.splice(index, 1);
      dragObj = item;
    },
  );
  if (!info.dropToGap) {
    // Drop on the content
    loop(data, dropKey, (item: TreeDataItem) => {
      item.children = item.children || [];
      /// where to insert 示例添加到头部，可以是随意位置
      item.children.unshift(dragObj);
    });
  } else if (
    (info.node.children || []).length > 0 && // Has children
    info.node.expanded && // Is expanded
    dropPosition === 1 // On the bottom gap
  ) {
    loop(data, dropKey, (item: TreeDataItem) => {
      item.children = item.children || [];
      // where to insert 示例添加到头部，可以是随意位置
      item.children.unshift(dragObj);
    });
  } else {
    let ar: TreeProps["treeData"] = [];
    let i = 0;
    loop(
      data,
      dropKey,
      (_item: TreeDataItem, index: number, arr: TreeProps["treeData"]) => {
        ar = arr;
        i = index;
      },
    );
    if (dropPosition === -1) {
      ar.splice(i, 0, dragObj);
    } else {
      ar.splice(i + 1, 0, dragObj);
    }
  }
  gData.value = data;
};
</script>
