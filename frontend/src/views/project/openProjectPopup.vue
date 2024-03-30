<template>
  <div id="components-modal-demo-position">
    <a-modal
      v-model:open="modalVisible"
      :closable="true"
      :title="props.placeholder"
      centered
      okType="default"
      ok-text="从目录中打开项目"
      :ok-button-props="{
        class: 'bg-blue-300 text-black mt-4',
      }"
      :cancel-button-props="{ ghost: true }"
      @ok="OpenDir"
    >
      <div class="mr-4 my-2">
        <a-alert
          closable
          description="若项目不是由VPSimple创建，配置文件可能读取失败"
          type="success"
        />
      </div>
      <div class="flex justify-start items-center">
        <span class="select-none text-h5">最近项目</span>
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
                    >打开</a-button
                  >
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
import { SelectDir } from "../../../wailsjs/go/system/SystemService";
import { onMounted, ref } from "vue";
import { HistoryProject } from "@/utils/historyProject";
import { useIndexStore } from "@/store";
import { useHistoryStore } from "@/store/history";

const OpenDir = () => {
  SelectDir("请选择项目存放目录").then((res: string) => {
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
onMounted(() => {
  for (let i = 0; i < storeHistory.currentList.length; i++) {
    data.value.push({
      title: storeHistory.currentList[i],
    });
  }
});

//切换项目
const openProject = (dir: string) => {
  useIndexStore().changeProject(dir);
  useHistoryStore().add(dir);
  modalVisible.value = false; //弹窗关闭
};
interface inputModelProps {
  defaultValue?: string;
  placeholder?: string;
}

const modalVisible = ref(false);
const props = defineProps<inputModelProps>();
const showModal = () => {
  modalVisible.value = true;
};

defineExpose({ showModal });
</script>
