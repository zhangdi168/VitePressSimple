<template>
  <div class="flex justify-start">
    <sim-switch
      :label="props.boolLabel"
      :tooltip="props.boolTooltip"
      v-model="useChecked"
    ></sim-switch>

    <sim-input
      :tooltip="props.inputTooltip"
      v-model="useInput"
      :label="props.inputLabel"
      :disabled="!useChecked"
    ></sim-input>
  </div>
</template>
<script setup lang="ts">
import SimSwitch from "@/components/simSwitch .vue";
import SimInput from "@/components/simInput.vue";
import { onMounted, ref, watch } from "vue";

const model = defineModel<any>();
const useChecked = ref(false);
const useInput = ref("");
onMounted(() => {
  if (model.value === false) {
    useChecked.value = false;
    useInput.value = "";
  } else {
    useChecked.value = true;
    useInput.value = model.value;
  }
});
watch(useChecked, (newValue) => {
  if (newValue) {
    model.value = useInput.value;
  } else {
    model.value = false;
  }
});
watch(useInput, (newValue) => {
  if (newValue === "") {
    model.value = false;
  } else {
    model.value = newValue;
  }
});

interface Props {
  boolTooltip: string;
  boolLabel: string;
  boolIsFullWidth?: boolean;
  inputPlaceholder?: string;
  inputTooltip: string;
  inputLabel: string;
  inputIsFullWidth?: boolean;
}

const props = defineProps<Props>();
</script>

<style scoped></style>
