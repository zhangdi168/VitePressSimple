<template>
  <div class="mt-2 items-end text-right">
    <q-input
      v-if="props.formType == 'text'"
      outlined
      type="text"
      :disable="props.disEnable"
      v-model="CurrInputValue"
      :label="props.label"
    />

    <q-input
      v-if="props.formType == 'password'"
      outlined
      type="password"
      :disable="props.disEnable"
      v-model="CurrInputValue"
      :label="props.label"
    />
    <q-input
      v-if="props.formType == 'textarea'"
      outlined
      type="textarea"
      :disable="props.disEnable"
      v-model="CurrInputValue"
      :label="props.label"
    />
  </div>
</template>

<script setup lang="ts">
import { defineProps, onMounted, ref, watch } from "vue";
import { ConfigGet, ConfigSet } from "../../wailsjs/go/system/SystemService";

const props = defineProps({
  configKey: {
    type: String,
    default: "",
  },
  label: {
    type: String,
    default: "请输入",
  },
  formType: {
    type: String,
    default: "text",
  },
  disEnable: {
    type: Boolean,
    default: false,
  },
});

onMounted(async () => {
  CurrInputValue.value = await ConfigGet(props.configKey);
});

let CurrInputValue = ref("");
watch(CurrInputValue, (nval: string, oval: string) => {
  //Set(props.ConfigKey,nval,props.ConfigType)
  ConfigSet(props.configKey, nval);
});

const SetModelValue = (val: string) => {
  LoadConfigValue();
};
const LoadConfigValue = () => {
  ConfigGet(props.configKey).then((res) => {
    CurrInputValue.value = res;
  });
};

defineExpose({ SetModelValue, LoadConfigValue });
</script>

<style scoped></style>
