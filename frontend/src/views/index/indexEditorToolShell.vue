<template>
  <a-dropdown class="mx-1" arrow>
    <icon-park
      class="select-none cursor-pointer"
      :size="storeLayout.editorToolIconSize"
      fill="black"
      strokeLinejoin="bevel"
      theme="outline"
      type="terminal"
    />
    <template #overlay>
      <a-menu>
        <a-menu-item @click="isShowShellEdit = true">
          <a href="javascript:;">终端配置</a>
        </a-menu-item>
        <a-menu-item @click="isShowShellLog = true">
          <a href="javascript:;">运行日志</a>
        </a-menu-item>
        <hr />
        <div v-for="(item, index) in cmdList" :key="index">
          <a-menu-item @click="createAndRunCmd(index)">
            <a href="javascript:;">{{ item.menuLabel }}</a>
          </a-menu-item>
          <hr v-if="[2].includes(index)" />
        </div>
      </a-menu>
    </template>
  </a-dropdown>

  <!--  弹出层：编辑shell配置-->
  <a-modal
    @ok="saveConfig"
    ok-text="Save"
    width="40%"
    :cancel-button-props="{ style: 'display:none' }"
    :ok-button-props="{ class: 'bg-blue-500' }"
    v-model:open="isShowShellEdit"
    title="终端配置"
  >
    <!--    命令运行的的起始位置-->
    <sim-input
      :is-full-width="true"
      v-model="storeShell.vpsimpleConfig.shellBaseDir"
      tooltip="命令运行目录的路径，默认值是项目根目录"
      label="命令运行目录"
    ></sim-input>
    <!--    文档dev命令-->
    <sim-input
      :is-full-width="true"
      v-model="storeShell.vpsimpleConfig.cmdDocsDev"
      :tooltip="'默认值:' + defaultVpSimple.cmdDocsDev"
      label="文档dev命令"
    ></sim-input>
    <!--    文档build命令-->
    <sim-input
      :is-full-width="true"
      v-model="storeShell.vpsimpleConfig.cmdDocsBuild"
      :tooltip="'默认值:' + defaultVpSimple.cmdDocsBuild"
      label="文档build命令"
    ></sim-input>
    <!--    npm install命令-->
    <sim-input
      :is-full-width="true"
      v-model="storeShell.vpsimpleConfig.cmdNpmInstall"
      :tooltip="'默认值:' + defaultVpSimple.cmdNpmInstall"
      label="npm install命令"
    ></sim-input>
    <!--    git pull命令-->
    <sim-input
      :is-full-width="true"
      v-model="storeShell.vpsimpleConfig.cmdGitPull"
      :tooltip="'默认值:' + defaultVpSimple.cmdGitPull"
      label="git pull命令"
    ></sim-input>
    <!--    git add命令-->
    <sim-input
      :is-full-width="true"
      v-model="storeShell.vpsimpleConfig.cmdGitAdd"
      :tooltip="'默认值:' + defaultVpSimple.cmdGitAdd"
      label="git add命令"
    ></sim-input>
    <sim-input
      :is-full-width="true"
      v-model="storeShell.vpsimpleConfig.cmdGitCommit"
      :tooltip="'默认值:' + defaultVpSimple.cmdGitCommit"
      label="git commit命令"
    ></sim-input>
    <!--    git push命令-->
    <sim-input
      :is-full-width="true"
      v-model="storeShell.vpsimpleConfig.cmdGitPush"
      :tooltip="'默认值:' + defaultVpSimple.cmdGitPush"
      label="git push命令"
    ></sim-input>
    <!--    git init命令-->
    <sim-input
      :is-full-width="true"
      v-model="storeShell.vpsimpleConfig.cmdGitInit"
      :tooltip="'默认值:' + defaultVpSimple.cmdGitInit"
      label="git init命令"
    ></sim-input>
  </a-modal>

  <!--  弹出层命令运行日志-->
  <a-modal
    width="50%"
    :cancel-button-props="{ style: 'display:none' }"
    :ok-button-props="{ class: 'bg-blue-500', style: 'display:none' }"
    v-model:open="isShowShellLog"
    title="运行日志"
  >
    <!--    顶部终端列表tabs-->
    <a-radio-group v-model:value="currShellIndex" button-style="solid">
      <a-radio-button
        v-for="(value, key) in storeShell.messages"
        :key="key"
        :value="Number(key)"
        >shell{{ key }}
      </a-radio-button>
    </a-radio-group>

    <!--    黑色背景的终端日志-->
    <div
      v-show="Object.keys(storeShell.messages).length > 0"
      style="height: 45vh; overflow-y: auto"
      class="bg-black text-white mt-4 px-5 pt-5 pb-10"
    >
      <pre v-html="storeShell.messages[currShellIndex]"></pre>
    </div>
    <!--    底部按钮操作组-->
    <div
      v-show="
        Object.keys(storeShell.messages).length > 0 &&
        storeIndex.systemType != SystemWindows
      "
      class="flex mt-5 justify-end items-center"
    >
      <!--          <a-button class="mx-2 bg-orange-200">移除当前终端</a-button>-->
      <a-button
        @click="storeShell.stopCmd(currShellIndex)"
        class="mx-2 bg-red-200"
        >中断命令
      </a-button>
    </div>

    <!--    空终端-->
    <div v-show="Object.keys(storeShell.messages).length == 0">
      <a-empty description="暂无运行中的终端"></a-empty>
    </div>
  </a-modal>
</template>
<script setup lang="ts">
import { IconPark } from "@icon-park/vue-next/es/all";
import { useLayoutStore } from "@/store/layout";
import { onMounted, ref } from "vue";

import SimInput from "@/components/simInput.vue";
import { useVpconfigStore } from "@/store/vpconfig";
import { IsEmptyValue } from "@/utils/utils";
import { useIndexStore } from "@/store";
import { defaultVpSimple } from "@/configs/defaultVpSimple";

import { useShellStore } from "@/store/shell";
import { ToastError } from "@/utils/Toast";
import { SystemWindows } from "@/constant/enums/system";

const storeIndex = useIndexStore();
const storeLayout = useLayoutStore();
const storeConfig = ref(useVpconfigStore());
const storeShell = useShellStore();
const currShellIndex = ref(0);
const isShowShellEdit = ref(false);
const isShowShellLog = ref(false);

interface CmdItem {
  menuLabel: string;

  cmd: string;
  isAlone: boolean;
  isSystemShell: boolean;
}

const cmdList = ref<CmdItem[]>([]);

// const vpsimpleConfig = ref<VPSimpleConfig>({} as VPSimpleConfig);
onMounted(async () => {
  await storeIndex.loadTreeData();
  storeShell.loadVpSimpleConfig();
  setCmdList();
});

const setCmdList = () => {
  cmdList.value = [
    {
      menuLabel: "npm install",
      cmd: storeShell.vpsimpleConfig.cmdNpmInstall,
      isAlone: false,
      isSystemShell: false,
    },
    {
      menuLabel: "npm run docs:dev",
      cmd: storeShell.vpsimpleConfig.cmdDocsDev,
      isAlone: true,
      isSystemShell: false,
    },
    {
      menuLabel: "npm run docs:build",
      cmd: storeShell.vpsimpleConfig.cmdDocsBuild,
      isAlone: false,
      isSystemShell: false,
    },

    {
      menuLabel: "git init",
      cmd: storeShell.vpsimpleConfig.cmdGitInit,
      isAlone: false,
      isSystemShell: false,
    },
    {
      menuLabel: "git add",
      cmd: storeShell.vpsimpleConfig.cmdGitAdd,
      isAlone: false,
      isSystemShell: false,
    },
    {
      menuLabel: "git commit",
      cmd: storeShell.vpsimpleConfig.cmdGitCommit,
      isAlone: false,
      isSystemShell: false,
    },
    {
      menuLabel: "git push",
      cmd: storeShell.vpsimpleConfig.cmdGitPush,
      isAlone: false,
      isSystemShell: false,
    },

    {
      menuLabel: "git pull",
      cmd: storeShell.vpsimpleConfig.cmdGitPull,
      isAlone: false,
      isSystemShell: false,
    },
  ];
};
// 创建并运行命令
//cmd 命令
//isAlone 是否单独创建一个终端(如果是则不显示日志)
//isSystemShell 是否使用系统自带的终端
const createAndRunCmd = async (index: number) => {
  const cmd = cmdList.value[index].cmd;
  const isAlone = cmdList.value[index].isAlone;
  const isSystemShell = cmdList.value[index].isSystemShell;
  if (IsEmptyValue(storeShell.vpsimpleConfig.shellBaseDir)) {
    ToastError("起始目录不能为空");
    return;
  }
  if (IsEmptyValue(cmd)) {
    ToastError("cmd命令不能为空");
    return;
  }

  if (isSystemShell) {
    currShellIndex.value = await storeShell.createSystemShellAndRun(
      storeShell.vpsimpleConfig.shellBaseDir,
      cmd,
    );
  } else {
    currShellIndex.value = await storeShell.createShellAndRun(
      storeShell.vpsimpleConfig.shellBaseDir,
      cmd,
      isAlone,
    );
  }
  if (isAlone && storeIndex.systemType == SystemWindows) {
    isShowShellLog.value = false;
  } else {
    isShowShellLog.value = true;
  }
};

const saveConfig = () => {
  storeConfig.value.configData["vpsimple"] = storeShell.vpsimpleConfig;
  storeConfig.value.saveConfig();
  setCmdList();
  isShowShellEdit.value = false;
};
</script>
<style scoped>
pre {
  white-space: pre-wrap; /* 保留空格和换行，并允许文本换行 */
  font-family: monospace; /* 使用等宽字体 */
}
</style>
