<template>
  <div id="components-modal-demo-position">
    <a-modal
      v-model:open="modalVisible"
      :closable="true"
      :title="props.placeholder"
      centered
      okType="default"
      :ok-text="lang('pageProject.openProjectFromDirectory')"
      :ok-button-props="{
        class: 'bg-blue-300 text-black mt-4',
      }"
      width="42%"
      :cancel-button-props="{ ghost: true }"
      @ok="OpenDir"
    >
      <div class="mr-4 my-2">
        <a-alert
          closable
          :description="lang('pageProject.errorProjectNotVPSimple')"
          type="success"
        />
      </div>
      <div class="flex justify-start items-center">
        <span class="select-none text-h5">{{
          lang("pageProject.recentProject")
        }}</span>
      </div>
      <div class="mt-2">
        <!--        最近项目列表-->
        <div>
          <a-list item-layout="horizontal" :data-source="data">
            <template #renderItem="{ item }">
              <a-list-item>
                <template #actions>
                  <a-button
                    @click="openProject(item.title)"
                    class="bg-blue-500 text-white"
                    >{{ lang("common.open") }}
                  </a-button>
                </template>
                <a-list-item-meta>
                  <template #title>
                    <span>{{ item.title }}</span>
                  </template>
                  <!--                  <template #avatar>-->
                  <!--                    <a-avatar src="https://joeschmoe.io/api/v1/random" />-->
                  <!--                  </template>-->
                </a-list-item-meta>
              </a-list-item>
            </template>
          </a-list>
        </div>
      </div>
    </a-modal>
  </div>
</template>
<script lang="ts" setup>
import {
  PathExists,
  PathJoin,
  SelectDir,
} from "../../../wailsjs/go/system/SystemService";
import { nextTick, onMounted, ref, watch } from "vue";
import { HistoryProject } from "@/utils/historyProject";
import { useIndexStore } from "@/store";
import { useHistoryStore } from "@/store/history";
import { ToastError, ToastInfo, ToastSuccess } from "@/utils/Toast";
import { lang } from "@/utils/language";

const OpenDir = () => {
  SelectDir(lang("pageProject.errorProjectRootEmpty")).then((res: string) => {
    if (res != "") {
      openProject(res);
    }
  });
};

interface DataItem {
  title: string;
}

const data = ref<DataItem[]>([]);
const storeHistory = useHistoryStore();
onMounted(async () => {
  // await loadHistoryList();
});

//切换项目
const openProject = async (dir: string) => {
  await useIndexStore().changeProject(dir);
  modalVisible.value = false; //弹窗关闭
};

interface inputModelProps {
  defaultValue?: string;
  placeholder?: string;
}

const modalVisible = ref(false);
const props = defineProps<inputModelProps>();
watch(modalVisible, async (newVal: boolean) => {
  console.log(newVal, "newVal -- console.log");
  if (newVal) {
    await loadHistoryList();
  }
});
const loadHistoryList = async () => {
  await useHistoryStore().initList();
  data.value = [];
  for (let i = 0; i < storeHistory.currentList.length; i++) {
    data.value.push({
      title: storeHistory.currentList[i],
    });
  }
};
const showModal = () => {
  modalVisible.value = true;
};

defineExpose({ showModal });
</script>
