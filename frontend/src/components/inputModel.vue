<template>
  <div id="components-modal-demo-position">
    <a-modal
      v-model:open="modalVisible"
      :closable="false"
      :title="props.placeholder"
      centered
      okType="default"
      :ok-button-props="{
        shape: 'round',
        class: 'bg-green-600',
      }"
      :cancel-button-props="{ ghost: true }"
      @ok="hideModal"
    >
      <div class="my-1">
        <a-textarea
          @keydown.prevent.enter="hideModal"
          v-model:value="val"
          :placeholder="props.placeholder"
          auto-size
        />
      </div>
    </a-modal>
  </div>
</template>
<script lang="ts" setup>
import { onMounted, ref } from "vue";

interface inputModelProps {
  defaultValue?: string;
  placeholder?: string;
}

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
