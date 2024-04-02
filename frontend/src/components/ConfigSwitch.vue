<template>
  <div class="flex items-center">
    <q-toggle left-label :label="props.label" v-model="CurrSwitchValue">
    </q-toggle>
    <a-tooltip :title="props.tooltip">
      <info-circle-outlined style="color: rgba(0, 0, 0, 0.45)" />
    </a-tooltip>
  </div>
</template>
<script setup lang="ts">
import { defineProps, onMounted, ref, watch } from "vue";
import { ConfigGet, ConfigSet } from "../../wailsjs/go/system/SystemService";
import { InfoCircleOutlined } from "@ant-design/icons-vue";

const props = defineProps({
  configKey: {
    type: String,
    default: "",
  },
  label: {
    type: String,
    default: "请输入",
  },
  disEnable: {
    type: Boolean,
    default: false,
  },
  tooltip: {
    type: String,
    default: "",
  },
});
const emits = defineEmits(["onChange"]);
let CurrSwitchValue = ref(false);
onMounted(async () => {
  let val = await ConfigGet(props.configKey);
  CurrSwitchValue.value = val == "yes";
});
watch(CurrSwitchValue, (nval: boolean, oval: boolean) => {
  ConfigSet(props.configKey, nval ? "yes" : "no");
  emits("onChange", nval);
});
</script>

<style scoped></style>
