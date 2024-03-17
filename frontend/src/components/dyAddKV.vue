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
      class="flex mt-2 justify-between items-center"
    >
      <div class="mt-1 w-2/5">
        <a-input
          class="w-full"
          v-model:value="inputs[index].key"
          :placeholder="props.keyPlaceholder"
        />
      </div>
      <div class="mt-1 w-2/5">
        <a-textarea
          class="w-full"
          :auto-size="{ minRows: 1 }"
          v-model:value="inputs[index].value"
          :placeholder="props.valuePlaceholder"
        ></a-textarea>
      </div>
      <div class="mt-1 flex">
        <a-button
          size="small"
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
import { onMounted, ref } from "vue";
import { IconPark } from "@icon-park/vue-next/es/all";

export interface dyAddKvProps {
  keyPlaceholder: string;
  valuePlaceholder: string;
  addBtnText: string;
  defaultValue?: any;
  addBtnClass?: string;
  keyName?: string; //键名 默认key
  valueName?: string; //值名 默认为value
}

const props = defineProps<dyAddKvProps>();
onMounted(() => {
  //判断props.defaultValue是数组还是对象
  if (Array.isArray(props.defaultValue)) {
    //遍历数组 依次将defaultValue的props.keyName和 props.valueName赋值给inuts
    for (let i = 0; i < props.defaultValue.length; i++) {
      inputs.value.push({
        key: props.defaultValue[i][props.keyName ?? "key"],
        value: props.defaultValue[i][props.valueName ?? "value"],
      });
    }
  } else {
    //遍历对象 依次将key 赋值key和value
    for (const key in props.defaultValue) {
      inputs.value.push({
        key: key,
        value: props.defaultValue[key],
      });
    }
    // console.log(props.defaultValue, "遍历对象 得到inputs -- console.log");
  }
});

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

// Expose a method 返回数组
//数组中对象key和value可以通过props.keyName和props.valueName来修改，默认为key和value
const getArray = () => {
  let keyName = props.keyName || "key";
  let valueName = props.valueName || "value";
  let result = [];
  for (let i = 0; i < inputs.value.length; i++) {
    result.push({
      [keyName]: inputs.value[i].key,
      [valueName]: inputs.value[i].value,
    });
  }
  result = result.filter((item) => {
    return item[keyName] !== "" && item[valueName] !== "";
  });
  return result;
};

// Get the inputs as an object 返回一个对象key是输入的key,value是输入的value
const getObject = () => {
  let result: any = {};
  for (let i = 0; i < inputs.value.length; i++) {
    result[inputs.value[i].key] = inputs.value[i].value;
  }
  return result;
};
defineExpose({ getArray, getObject });
</script>
