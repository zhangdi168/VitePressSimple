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
        class: 'bg-cyan-400 text-black mt-4',
      }"
      :cancel-button-props="{ ghost: true }"
      @ok="hideModal"
    >
      <div class="flex justify-start items-center">
        <span class="select-none text-h5">新建VitePress项目</span>
      </div>
      <div>
        <a-alert message="Success Text" type="success" />
      </div>
      <div class="my-5">
        <a-input
          v-model="formData.title"
          placeholder="请输入项目标题"
          class="w-full"
        >
          <template #suffix>
            <a-tooltip title="项目的标题，建议中文">
              <info-circle-outlined style="color: rgba(0, 0, 0, 0.45)" />
            </a-tooltip>
          </template>
        </a-input>
      </div>

      <div class="my-3">
        <a-input
          v-model="formData.identifier"
          placeholder="请输入项目标识（文件夹名）"
          class="w-full"
        >
          <template #suffix>
            <a-tooltip
              title="作为项目根目录文件夹名称，建议使用英文并遵循变量命名法"
            >
              <info-circle-outlined style="color: rgba(0, 0, 0, 0.45)" />
            </a-tooltip>
          </template>
        </a-input>
      </div>
      <div class="my-3">
        <a-input
          disabled
          v-model="formData.docDir"
          placeholder="请选择项目存放目录"
          class="w-full"
        >
          <template #suffix>
            <a-tooltip title="请选择项目存放目录" @click="selectProjectDir">
              <FileOutlined style="color: rgba(0, 0, 0, 0.45)" />
            </a-tooltip>
          </template>
        </a-input>
      </div>
      <div class="my-3">
        <a-input
          v-model="formData.docDir"
          placeholder="请输入markdown文档所在子目录"
          class="w-full"
        >
          <template #suffix>
            <a-tooltip title="请输入markdown文档存放的子目录，如：docs">
              <info-circle-outlined style="color: rgba(0, 0, 0, 0.45)" />
            </a-tooltip>
          </template>
        </a-input>
      </div>
    </a-modal>
  </div>
</template>
<script lang="ts" setup>
import { onMounted, ref } from "vue";
import { ProjectCreate } from "@/types/project";
import { InfoCircleOutlined, FileOutlined } from "@ant-design/icons-vue";
import { SelectDir } from "../../../wailsjs/go/system/SystemService";

const formData = ref<ProjectCreate>({} as ProjectCreate);

interface inputModelProps {
  defaultValue?: string;
  placeholder?: string;
}

const selectProjectDir = () => {
  SelectDir("请选择项目存放目录").then((res) => {
    formData.value.docDir;
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
