<template>
  <!--  <div class="mb-1">-->
  <!--    <a-alert message="Success Text" type="success"> asas</a-alert>-->
  <!--  </div>-->
  <div v-show="storeIndex.currVueCode == ''" class="flex justify-start mb-1.5">
    <a-button
      class="bg-blue-300 flex justify-center items-center hover:bg-blue-100"
      @click="useDemoCode"
    >
      <icon-park
        class="mr-1"
        strokeLinejoin="bevel"
        theme="outline"
        type="page-template"
      />
      使用示例代码
    </a-button>
  </div>
  <Codemirror
    ref="cmRef"
    v-model:value="storeIndex.currVueCode"
    :options="cmOptions"
    border
    height="70vh"
    width="100%"
    @change="onChange"
    @input="onInput"
    @ready="onReady"
  >
  </Codemirror>
</template>
<script lang="ts" setup>
import { ref, onMounted, onUnmounted, nextTick } from "vue";
import "codemirror/mode/javascript/javascript.js";
import Codemirror from "codemirror-editor-vue3";
import type { CmComponentRef } from "codemirror-editor-vue3";
import type { Editor, EditorConfiguration } from "codemirror";
import { useIndexStore } from "../../store";
import { IconPark } from "@icon-park/vue-next/es/all";

const storeIndex = useIndexStore();
const code = ref<string>(
  `<script setup>\nimport { ref } from 'vue'\nconst count = ref(0)\n</\script>\n\n<style module>\n.button {color: red; font-weight: bold;}\n</style>\n`,
);

const useDemoCode = () => {
  storeIndex.setCurrVueCode(code.value);
};
onMounted(() => {
  // nextTick(() => {
  //   storeIndex.setCurrVueCode(code.value)
  // })
  // setTimeout(() => {
  //   cmRef.value?.refresh()
  // }, 1000)
  //
  // setTimeout(() => {
  //   cmRef.value?.resize(300, 200)
  // }, 2000)
  //
  // setTimeout(() => {
  //   cmRef.value?.cminstance.isClean()
  // }, 3000)
});

const cmRef = ref<CmComponentRef>();
const cmOptions: EditorConfiguration = {
  mode: "text/javascript",
};

const onChange = (val: string, cm: Editor) => {
  console.log(val);
  console.log(cm.getValue());
};

const onInput = (val: string) => {
  console.log(val);
};

const onReady = (cm: Editor) => {
  console.log(cm.focus());
};

onUnmounted(() => {
  cmRef.value?.destroy();
});
</script>
