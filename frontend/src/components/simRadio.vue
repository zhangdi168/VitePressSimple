<template>
  <div class="pl-3 my-3" :class="props.isFullWidth ? 'w-full' : 'w-1/3'">
    <div class="flex items-center justify-start">
      <a-form-item :label="label">
        <a-radio-group v-model:value="model">
          <a-radio
            v-for="item in props.items"
            :key="item.value"
            :value="item.value"
            >{{ item.label }}
          </a-radio>
        </a-radio-group>
        <a-tooltip :title="props.tooltip">
          <info-circle-outlined style="color: rgba(0, 0, 0, 0.45)" />
        </a-tooltip>
      </a-form-item>
    </div>
  </div>
</template>
<script setup lang="ts">
import { onMounted } from "vue";

import { InfoCircleOutlined } from "@ant-design/icons-vue";

const model = defineModel();

interface radioItems {
  label: string;
  value: string;
}

interface Props {
  tooltip: string;
  label: string;
  items: radioItems[];
  isFullWidth?: boolean;
}

const props = defineProps<Props>();
</script>

<style scoped></style>
