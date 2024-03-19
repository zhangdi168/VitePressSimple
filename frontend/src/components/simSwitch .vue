<template>
  <div class="mx-3 my-3" :class="props.isFullWidth ? 'w-full' : 'w-1/3'">
    <div class="flex items-center justify-start">
      <div>
        <q-toggle
          :label="label"
          left-label
          class="w-full"
          v-model="model"
        ></q-toggle>
      </div>
      <div>
        <a-tooltip :title="props.tooltip">
          <info-circle-outlined style="color: rgba(0, 0, 0, 0.45)" />
        </a-tooltip>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
const model = defineModel();
import { InfoCircleOutlined } from "@ant-design/icons-vue";

interface Props {
  tooltip: string;
  label: string;
  isFullWidth?: boolean;
}

const props = defineProps<Props>();
</script>

<style scoped></style>
