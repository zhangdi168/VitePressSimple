<template>
  <div id="components-modal-demo-position">
    <a-modal
      v-model:open="modalVisible"
      :closable="true"
      :title="props.title ? props.title : props.placeholder"
      centered
      okType="default"
      :ok-button-props="{
        shape: '',
        class: 'bg-blue-600 text-white',
      }"
      :cancel-button-props="{ ghost: true }"
      @ok="hideModal"
    >
      <div class="mt-4 mb-2">
        <a-input
          ref="refInput"
          v-model:value="val"
          @keydown.prevent.enter="hideModal"
          :placeholder="props.placeholder"
          auto-size
        />
      </div>
    </a-modal>
  </div>
</template>
<script lang="ts" setup>
import { nextTick, onMounted, ref } from "vue";

interface inputModelProps {
  defaultValue?: string;
  placeholder?: string;
  title?: string;
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
const refInput = ref();
const showModal = (val_: string = "") => {
  val.value = val_;
  modalVisible.value = true;
  nextTick(() => {
    refInput.value.focus();
  });
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
