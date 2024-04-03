<template>
  <div class="flex items-center">
    <q-input filled v-model="color" :label="props.label" class="my-input">
      <template v-slot:append>
        <q-icon name="colorize" class="cursor-pointer">
          <q-popup-proxy cover transition-show="scale" transition-hide="scale">
            <q-color v-model="color" />
          </q-popup-proxy>
        </q-icon>
      </template>
    </q-input>
  </div>
</template>
<script setup lang="ts">
// const model = defineModel();

import { onMounted, ref, watch } from "vue";
import { ConfigGet, ConfigSet } from "../../wailsjs/go/system/SystemService";

const color = ref();

interface Props {
  label: string;
  configKey: string;
}

const props = defineProps<Props>();
onMounted(() => {
  ConfigGet(props.configKey).then((res) => {
    color.value = res;
  });
});

//监听变化
watch(color, (newVal) => {
  ConfigSet(props.configKey, newVal);
});
</script>

<style scoped></style>
