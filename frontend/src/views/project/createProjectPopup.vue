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
      @ok="Create"
    >
      <div class="flex justify-start items-center">
        <span class="select-none text-h5">新建VitePress项目</span>
      </div>
      <div class="mt-2">
        <a-alert closable description="基于VitePress1.0.2创建" type="success" />
      </div>
      <div class="my-2">
        <a-input
          v-model:value="formData.title"
          placeholder="请输入项目名称"
          class="w-full"
        >
          <template #suffix>
            <a-tooltip title="项目的名称，请遵循文件的命名规则">
              <span class="text-gray-500">名称</span>
            </a-tooltip>
          </template>
        </a-input>
      </div>

      <!--      <div class="my-3">-->
      <!--        <a-input-->
      <!--          v-model:value="formData.identifier"-->
      <!--          placeholder="请输入项目标识（文件夹名）"-->
      <!--          class="w-full"-->
      <!--        >-->
      <!--          <template #suffix>-->
      <!--            <a-tooltip-->
      <!--              title="名称标识，作为项目根目录文件夹名称，建议使用英文并遵循变量命名法"-->
      <!--              ><span class="text-gray-500">名称</span>-->
      <!--            </a-tooltip>-->
      <!--          </template>-->
      <!--        </a-input>-->
      <!--      </div>-->

      <div class="my-2">
        <a-input
          v-model:value="formData.description"
          placeholder="请输入项目描述"
          class="w-full"
        >
          <template #suffix>
            <a-tooltip title="项目的描述，建议中文">
              <span class="text-gray-500">描述</span>
            </a-tooltip>
          </template>
        </a-input>
      </div>

      <div class="my-3 flex justify-between">
        <div class="flex-1">
          <a-input
            v-model:value="formData.dir"
            placeholder="项目根目录是尝试寻找 .vitepress 特殊目录的地方"
            class="w-full"
          >
          </a-input>
        </div>
        <div class="ml-2">
          <a-button class="bg-blue-200" @click="selectProjectDir"
            >选择根目录
          </a-button>
        </div>
      </div>
      <div class="my-3">
        <a-input
          disabled
          v-model:value="formData.docDir"
          placeholder=""
          class="w-full"
        >
          <template #suffix>
            <a-tooltip
              title="请源目录是 Markdown 源文件所在的位置，默认为./docs，创建完成后可修改"
            >
              <span class="text-gray-500">源目录(文档目录)</span>
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
import { PathJoin, SelectDir } from "../../../wailsjs/go/system/SystemService";
import { CreateProject } from "../../../wailsjs/go/vpsimpler/VpManager";
import { ToastError, ToastInfo } from "@/utils/Toast";
import { HistoryProject } from "@/utils/historyProject";
import { useHistoryStore } from "@/store/history";
import { CreateDir } from "../../../wailsjs/go/services/ArticleTreeData";
import { useIndexStore } from "@/store";

const formData = ref<ProjectCreate>({
  title: "VPSimpleProject",
  dir: "",
  docDir: "./docs",
} as ProjectCreate);
const test = ref("1");

interface inputModelProps {
  defaultValue?: string;
  placeholder?: string;
}

const Create = async () => {
  await CreateDir(formData.value.dir);
  CreateProject(
    formData.value.title,
    formData.value.description,
    formData.value.dir,
  ).then((res) => {
    if (res != "") {
      ToastError(res);
    } else {
      ToastInfo("创建完成");
      modalVisible.value = false;
      useHistoryStore().add(formData.value.dir); //添加到历史记录
      useIndexStore().changeProject(formData.value.dir);
    }
  });
};

const selectProjectDir = async () => {
  if (formData.value.title == "") {
    ToastError("请先输入项目名称");
    return;
  }
  let dir = await SelectDir("请选择项目存放目录");
  if (dir != "") {
    formData.value.dir = await PathJoin([dir, formData.value.title]);
  }
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

//隐藏弹出层并返回输入的值
const emits = defineEmits(["submitInputModal"]);
const hideModal = () => {
  modalVisible.value = false;
  emits("submitInputModal", val.value);
  return val.value;
};
const showModal = (val_: string = "") => {
  if (val_ != "") {
    val.value = val_;
  }
  modalVisible.value = true;
};

defineExpose({ showModal });
</script>
