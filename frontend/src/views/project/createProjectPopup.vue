<template>
  <div id="components-modal-demo-position">
    <a-modal
      v-model:open="modalVisible"
      :closable="true"
      :title="props.placeholder"
      centered
      okType="default"
      ok-text="创建"
      :ok-button-props="{
        class: 'bg-blue-300 text-black mt-4',
      }"
      :cancel-button-props="{ ghost: true }"
      @ok="hideModal"
    >
      <div class="flex justify-start items-center">
        <span class="select-none text-h5">新建VitePress项目</span>
      </div>
      <div class="mt-2">
        <a-alert
          description="基于VitePress1.0.0-rc.45创建,创建前请确保已经安装nodejs和npm"
          type="success"
        />
      </div>
      <div class="my-2">
        <a-input
          v-model:value="formData.title"
          placeholder="请输入项目标题"
          class="w-full"
        >
          <template #suffix>
            <a-tooltip title="项目的标题，建议中文">
              <span class="text-gray-500">项目标题</span>
            </a-tooltip>
          </template>
        </a-input>
      </div>

      <div class="my-3">
        <a-input
          v-model:value="formData.identifier"
          placeholder="请输入项目标识（文件夹名）"
          class="w-full"
        >
          <template #suffix>
            <a-tooltip
              title="作为项目根目录文件夹名称，建议使用英文并遵循变量命名法"
              ><span class="text-gray-500">项目标识</span>
            </a-tooltip>
          </template>
        </a-input>
      </div>
      <div class="my-3 flex justify-between">
        <div class="flex-1">
          <a-input
            v-model:value="formData.dir"
            placeholder="请选择项目存放目录"
            class="w-full"
          >
          </a-input>
        </div>
        <div class="ml-2">
          <a-button class="bg-blue-200" @click="selectProjectDir"
            >选择目录</a-button
          >
        </div>
      </div>
      <div class="my-3">
        <a-input
          v-model:value="formData.docDir"
          placeholder="请输入markdown文档所在子目录"
          class="w-full"
        >
          <template #suffix>
            <a-tooltip
              title="请输入markdown文档存放的子目录（基于项目目录），如：/docs"
            >
              <span class="text-gray-500">文档目录</span>
            </a-tooltip>
          </template>
        </a-input>
      </div>
    </a-modal>
  </div>
</template>
<script lang="ts" setup>
import { h, onMounted, ref } from "vue";
import { ProjectCreate } from "@/types/project";
import { InfoCircleOutlined, FileOutlined } from "@ant-design/icons-vue";
import { SelectDir } from "../../../wailsjs/go/system/SystemService";

const formData = ref<ProjectCreate>({
  title: "我的vitepress站点",
  identifier: "mySite",
  dir: "",
  docDir: "/docs",
} as ProjectCreate);
const test = ref("1");
interface inputModelProps {
  defaultValue?: string;
  placeholder?: string;
}

const selectProjectDir = () => {
  SelectDir("请选择项目存放目录").then((res: string) => {
    console.log(res, "-----res value");
    formData.value.dir = res;
  });
};

const props = defineProps<inputModelProps>();
onMounted(() => {
  if (
    props.defaultValue != undefined ||
    props.defaultValue != null ||
    props.defaultValue != ""
  ) {
    val.value = props.defaultValue;
  }
});

const modalVisible = ref<boolean>(false);
const val = ref();

const showModal = (val_: string = "") => {
  if (val_ != "") {
    val.value = val_;
  }
  modalVisible.value = true;
};

//隐藏弹出层并返回输入的值
const emits = defineEmits(["submitInputModal"]);
const hideModal = () => {
  modalVisible.value = false;
  emits("submitInputModal", val.value);
  return val.value;
};

defineExpose({ showModal });
</script>
