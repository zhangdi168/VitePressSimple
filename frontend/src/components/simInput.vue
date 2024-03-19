<template>
  <div class="px-1.5 my-3" :class="props.isFullWidth ? 'w-full' : 'w-1/3'">
    <a-input
      :auto-size="{ minRows: 1, maxRows: 3 }"
      v-model:value="model"
      :placeholder="props.placeholder ?? ''"
      class="w-full"
    >
      <template #prefix>
        <span class="text-gray-700">{{ props.label }}：</span>
      </template>
      <template #suffix>
        <a-tooltip :title="props.tooltip">
          <info-circle-outlined style="color: rgba(0, 0, 0, 0.45)" />
        </a-tooltip>
      </template>
    </a-input>
  </div>
</template>
<script setup lang="ts">
const model = defineModel();
import { InfoCircleOutlined } from "@ant-design/icons-vue";

interface Props {
  placeholder?: string;
  tooltip: string;
  label: string;
  isFullWidth?: boolean;
}

const props = defineProps<Props>();
</script>

<style scoped></style>
