<template>
  <div>
    <div class="flex justify-start pl-3">
      <a-button
        class="bg-blue-200 flex justify-center items-center hover:bg-blue-100"
        :class="props.addBtnClass"
        @click="addInput"
      >
        <icon-park
          class="mr-1"
          strokeLinejoin="bevel"
          theme="outline"
          type="add-one"
        />
        {{ props.addBtnText }}
      </a-button>
    </div>
    <div
      v-for="(item, index) in inputs"
      :key="index"
      class="flex mt-2 px-3 justify-between"
    >
      <div class="mt-1">
        <a-input
          v-model="item.key"
          :placeholder="props.keyPlaceholder"
          prefix=""
          suffix=""
        />
      </div>
      <div class="mt-1">
        <a-input
          v-model="item.value"
          :placeholder="props.valuePlaceholder"
          prefix=""
          suffix=""
        />
      </div>
      <div>
        <a-button
          v-if="inputs.length > 0"
          type="default"
          @click="removeInput(index)"
          >移除
        </a-button>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { IconPark } from "@icon-park/vue-next/es/all";

export interface dyAddKvProps {
  keyPlaceholder: string;
  valuePlaceholder: string;
  addBtnText: string;
  addBtnClass?: string;
}

const props = defineProps<dyAddKvProps>();

interface InputItem {
  key: string;
  value: string;
}

const inputs = ref<InputItem[]>([]);

// Method to add a new set of inputs
const addInput = () => {
  inputs.value.push({ key: "", value: "" });
};

// Method to remove a set of inputs
const removeInput = (index: number) => {
  inputs.value.splice(index, 1);
};

// Expose a method to return the input items array
const getInputs = () => inputs.value;
defineExpose({ getInputs });
// Form submit handler
const handleSubmit = () => {
  console.log("Form Submitted:", inputs.value);
  // You can call your function here to process the data and do whatever is needed.
};
</script>
