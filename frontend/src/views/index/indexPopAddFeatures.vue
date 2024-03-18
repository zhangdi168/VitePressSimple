<template>
  <div>
    <a-modal
      v-model:open="modalVisible"
      :closable="true"
      :title="props.placeholder"
      centered
      okType="default"
      :ok-button-props="{
        shape: 'round',
        class: 'bg-green-600',
      }"
      :cancel-button-props="{ ghost: true }"
      @ok="okClick"
    >
      <div>
        <div></div>
      </div>
    </a-modal>
  </div>
</template>
<script setup lang="ts">
import { ref } from "vue";

export interface dyAddFeaturesProps {
  title: string;
  placeholder: string;
}

const props = defineProps<dyAddFeaturesProps>();
const modalVisible = ref();

const okClick = () => {};

const showModal = () => {
  modalVisible.value = true;
};

defineExpose({ showModal });
</script>
<style scoped></style>
