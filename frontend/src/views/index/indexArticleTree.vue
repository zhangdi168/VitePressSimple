<template>
  <div v-if="!isEmptyArray(storeIndex.articleTreeData)" class="mt-1">
    <a-tree
      v-model:expandedKeys="storeIndex.expandKeys"
      v-model:selected-keys="storeIndex.selectKeys"
      :tree-data="storeIndex.ArticleTreeData"
      block-node
      class="text-blue text-center items-center"
      draggable
      @drop="onDrop"
    >
      <!--      标题-->
      <template #title="{ key: treeKey, title }">
        <!--     右键菜单-->
        <q-menu touch-position context-menu>
          <q-list dense style="min-width: 100px">
            <q-item
              v-if="!title.endsWith('.md')"
              @click="showPopCreate(treeKey, 'file')"
              clickable
              v-close-popup
            >
              <q-item-section>
                <menu-item title="新建文章" icon="newlybuild"></menu-item>
              </q-item-section>
            </q-item>
            <q-separator />
            <q-item
              v-if="!title.endsWith('.md')"
              @click="showPopCreate(treeKey, 'dir')"
              clickable
              v-close-popup
            >
              <menu-item title="新建目录" icon="folder-plus"></menu-item>
            </q-item>
            <q-separator />
            <q-item
              clickable
              @click="showPopRename(title, treeKey)"
              v-close-popup
            >
              <menu-item title="重命名" icon="file-editing"></menu-item>
            </q-item>
            <q-separator />
            <q-item clickable @click="copyPath(treeKey)" v-close-popup>
              <menu-item title="复制" icon="copy-one"></menu-item>
            </q-item>
            <q-item
              v-if="title.endsWith('.md')"
              clickable
              @click="copyRouter(treeKey)"
              v-close-popup
            >
              <menu-item title="复制路由" icon="copy-link"></menu-item>
            </q-item>
            <q-separator />
            <q-item
              v-if="!title.endsWith('.md') && storeIndex.currCopyPath != ''"
              clickable
              @click="pastePath(treeKey)"
              v-close-popup
            >
              <menu-item title="粘贴" icon="intersection"></menu-item>
            </q-item>
            <q-separator />
            <q-item clickable @click="deletePath(treeKey)" v-close-popup>
              <menu-item title="删除" icon="delete-five"></menu-item>
            </q-item>
          </q-list>
        </q-menu>
        <div
          style="font-size: 1rem"
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

      <template #icon="{ key }">
        <template v-if="key.endsWith('.md')">
          <FileTextOutlined />
        </template>
      </template>
    </a-tree>

    <input-model
      ref="refModalCreate"
      placeholder="请输入不带.md后缀的文章标题"
      title="请输入文章标题"
      @submit-input-modal="onSubmitInputModalCreateFile"
    ></input-model>
    <input-model
      ref="refModalCreateDir"
      placeholder="请输入文件夹名称"
      @submit-input-modal="onSubmitInputModalCreateDir"
    ></input-model>
    <input-model
      ref="refModalRename"
      placeholder="请输入新名称"
      @submit-input-modal="onSubmitInputModalRename"
    ></input-model>
    <input-model
      ref="refPasteName"
      title="请确认新路径"
      placeholder="请确认或修改新路径"
      @submit-input-modal="onSubmitInputModalPasteName"
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
  CreateDir,
  CreateFile,
  DeletePath,
  ReadFileContent,
  Rename,
} from "../../../wailsjs/go/services/ArticleTreeData";
import { getParentDirectory, removeMdExtension } from "@/utils/file";
import { ToastCheck, ToastError, ToastInfo, ToastSuccess } from "@/utils/Toast";
import { Modal } from "ant-design-vue";
import { isEmptyArray } from "@/utils/array";
import { useVpconfigStore } from "@/store/vpconfig";
import { parseTagContent, regexScript, regexStyle } from "@/utils/parse";
import {
  ConfigGetBool,
  CopyPath,
  GetPathDir,
  GetPathFileName,
  PathExists,
  PathJoin,
} from "../../../wailsjs/go/system/SystemService";
import { IsEmptyValue } from "@/utils/utils";
import MenuItem from "@/components/menuItem.vue";
import { ConfigKeyChangeAutoSave } from "@/constant/keys/config";
import { replaceLocalStaticToImageUrl } from "@/utils/repalceStatic";

const storeIndex = useIndexStore();
const moreIconShownKeys = ref<string[]>([]);
const storeVpConfig = useVpconfigStore();
onMounted(async () => {
  nextTick(async () => {
    await storeIndex.loadTreeData();
    //打开
    if (storeIndex.selectKeys && storeIndex.selectKeys.length > 0) {
      openArticle(storeIndex.selectKeys[0]);
    }
  });
});

const onDrop = (info: AntTreeNodeDropEvent) => {
  // console.log(info, "-----info value");
  storeIndex.moveTo(String(info.dragNode.key), String(info.node.key));
};

const deletePath = (key: string) => {
  Modal.confirm({
    content: "确定删除么？ " + key,
    okType: "danger",
    okText: "确定删除",
    icon: createVNode(ExclamationCircleOutlined),
    onOk() {
      DeletePath(key).then((res: string) => {
        if (res != "") {
          ToastError(res);
        } else {
          ToastInfo(key + "：删除成功");
          //更新当前文件

          if (storeIndex.currArticlePath == key) {
            storeIndex.currArticlePath = "";
          }
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
//复制一个路径
const copyPath = (path: string) => {
  storeIndex.currCopyPath = path;
  ToastInfo("路径：‘" + path + "’已经复制");
};
//粘贴路径
const refPasteName = ref();
const pastePath = async (baseDir: string) => {
  if (storeIndex.currCopyPath == "") {
    ToastError("没有复制路径");
    return;
  }
  let oldName = await GetPathFileName(storeIndex.currCopyPath);
  let fullPath = await PathJoin([baseDir, oldName]);
  refPasteName.value.showModal(fullPath);
};
//提交粘贴新路径
const onSubmitInputModalPasteName = async (fullPath: string) => {
  let isExists = await PathExists(fullPath);
  let isDir = !fullPath.endsWith(".md");
  if (isExists) {
    ToastError(`${isDir ? "目录" : "文章"}路径已存在：${fullPath}`);
    return;
  }
  let res = await CopyPath(storeIndex.currCopyPath, fullPath, isDir);
  if (ToastCheck(res, "已成功复制")) {
    storeIndex.currCopyPath = "";
    await storeIndex.loadTreeData();
  }
};

//双击标题展开菜单
const handleClick = (key: string) => {
  if (isDir(key)) {
    if (storeIndex.expandKeys.includes(key)) {
      //删除
      storeIndex.expandKeys = storeIndex.expandKeys.filter((s) => s !== key);
    } else {
      //删除
      storeIndex.expandKeys.push(key); // 或者添加逻辑来控制是否展开/收起
    }
  } else {
    //文件
    openArticle(key);
  }
};

const openArticle = async (path: string) => {
  let isAutoSave = await ConfigGetBool(ConfigKeyChangeAutoSave);
  if (isAutoSave === true) {
    await storeIndex.saveCurrArticle(false); //先保存
  }
  const content = await ReadFileContent(path);
  // console.log(content, "content -- console.log");

  let matterData = matter(content);
  // console.log(matterData.data, "matterData.data -- console.log");
  let matchScriptArray = parseTagContent(matterData.content ?? "", regexScript);
  let matchStyleArray = parseTagContent(matterData.content ?? "", regexStyle);
  let scriptContent = matchScriptArray[0] ?? "";
  let styleContent = matchStyleArray[0] ?? "";
  await storeIndex.setCurrScriptContent(scriptContent);
  await storeIndex.setCurrStyleContent(styleContent);
  storeIndex.setCurrArticlePath(path);
  // //matter字符串

  storeIndex.setCurrArticleFrontMatter(matterData.data);

  let vditorContent = (matterData.content ?? "")
    .replace(scriptContent, "")
    .replace(styleContent, "");
  let val = vditorContent ? vditorContent : "# hello vitePress client";
  // console.log("val:" + val);
  //相对路径转换成 域名替换
  val = replaceLocalStaticToImageUrl(vditorContent);
  storeIndex.Vditor?.setValue(val); //设置编辑器的值
};
const isDir = (key: string) => {
  return !key.includes(".md");
};

// const expandedKeys = ref<string[]>([]);
// const selectKeys = ref<string[]>([]);

//弹出层回调 新建
const refModalCreate = ref();
const currentCreateParentPath = ref<string>();
const showPopCreate = (path_: string, typ: string) => {
  currentCreateParentPath.value = path_;
  if (typ == "dir") {
    refModalCreateDir.value.showModal("");
  } else if (typ == "file") {
    refModalCreate.value.showModal("");
  }
};
const onSubmitInputModalCreateFile = async (value: string) => {
  let fullFile = currentCreateParentPath.value + "/" + value + ".md";
  let res = await CreateFile(fullFile);
  if (ToastCheck(res, "文件创建完成")) await storeIndex.loadTreeData();
};

//创建子文件夹
const refModalCreateDir = ref();
const onSubmitInputModalCreateDir = async (value: string) => {
  let res = await CreateDir(currentCreateParentPath.value + "/" + value);
  if (ToastCheck(res, "文件夹创建完成")) await storeIndex.loadTreeData();
};

//弹出层回调 重命名
const refModalRename = ref();
const currentRenamePath = ref<string>();
const showPopRename = (value: string, path_: string) => {
  refModalRename.value.showModal(removeMdExtension(value));
  currentRenamePath.value = path_;
};

const copyRouter = async (path: string) => {
  try {
    let router = path
      .replaceAll(storeVpConfig.fullSrcDir, "")
      .replaceAll("\\", "/")
      .replaceAll(".md", "");
    await navigator.clipboard.writeText(router);
    ToastSuccess("已复制：'" + router + "'");
  } catch (err: any) {
    ToastError("复制失败" + err.message);
  }
};
const onSubmitInputModalRename = async (value: string) => {
  let oldPath = currentRenamePath.value;
  let dir = await GetPathDir(
    currentRenamePath.value ?? storeVpConfig.fullSrcDir,
  );
  let newPath = await PathJoin([dir, value]);
  if (oldPath?.endsWith(".md")) newPath = newPath + ".md";
  // newPath = newPath.replace("//", "/");
  storeIndex.currArticlePath = newPath;
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
