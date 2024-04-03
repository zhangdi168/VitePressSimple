<template>
  <div class="flex items-center">
    <a-form-item :label="label">
      <a-radio-group v-model:value="value">
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
</template>
<script setup lang="ts">
import { nextTick, onMounted, ref, watch } from "vue";

import { InfoCircleOutlined } from "@ant-design/icons-vue";
import { ConfigGet, ConfigSet } from "../../wailsjs/go/system/SystemService";

const value = ref();

interface radioItems {
  label: string;
  value: string;
}

interface Props {
  tooltip: string;
  label: string;
  configKey: string;
  items: radioItems[];
  isFullWidth?: boolean;
}

onMounted(() => {
  ConfigGet(props.configKey).then((res) => {
    value.value = res;
  });
});
const props = defineProps<Props>();
watch(value, (newVal) => {
  ConfigSet(props.configKey, newVal);
});
</script>

<style scoped></style>
