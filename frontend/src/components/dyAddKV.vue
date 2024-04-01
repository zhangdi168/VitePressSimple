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
      class="flex mt-2 justify-start items-center"
    >
      <div class="mt-1 px-1 w-2/5">
        <a-input
          class="w-full"
          v-model:value="inputs[index].key"
          :placeholder="props.keyPlaceholder"
        />
      </div>
      <div class="mt-1 px-1 w-2/5">
        <a-textarea
          class="w-full"
          :auto-size="{ minRows: 1 }"
          v-model:value="inputs[index].value"
          :placeholder="props.valuePlaceholder"
        ></a-textarea>
      </div>
      <div class="mt-1 px-1">
        <a-button
          v-if="inputs.length > 0 && !props.removeConfirm"
          type="dashed"
          @click="removeInput(index)"
          >移除
        </a-button>
        <a-popconfirm
          v-if="props.removeConfirm"
          class="ml-3 pt-1"
          :title="props.removeConfirmText"
          ok-text="确认移除"
          cancel-text="算了"
          :ok-button-props="{ class: 'bg-blue-6 00 text-white' }"
          @confirm="removeInput(index)"
        >
          <a-button
            v-if="inputs.length > 0"
            type="dashed"
            @click="removeInput(index)"
            >移除
          </a-button>
        </a-popconfirm>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref, watch } from "vue";
import { IconPark } from "@icon-park/vue-next/es/all";

//双向绑定的数据
const arr = defineModel<Array<any>>("arr");
const obj = defineModel<any>("obj");
const objs = defineModel<any>("objs");
const arrayObject = defineModel<any>("arrayObject");
const inputs = ref<InputItem[]>([]);
watch(inputs.value, (newVal, oVal) => {
  arr.value = getArray(); //[{key:inpu1,value:val1},{}]
  obj.value = getObject(); //{propName1:vale1,propName2:value2}
  objs.value = getObjects(); // [{propName1:val1,propName2:val1},{}]

  //数组第一个值是outKeyName,第二个值是Object。很奇葩的格式。。。
  arrayObject.value = getArrayObject(); //[outKeyName, {keyName1:input1,keyName2:input2}]
});

export interface dyAddKvProps {
  keyPlaceholder: string;
  valuePlaceholder: string;
  addBtnText: string;
  defaultValue?: any;
  addBtnClass?: string;
  outKeyName?: string; //数组外层键名
  keyName?: string; //键名 默认key 仅对getArray有效
  valueName?: string; //值名 默认为value
  removeConfirm?: boolean; //是否需要确认删除
  removeConfirmText?: string; //确认删除文本
}

const props = defineProps<dyAddKvProps>();
const emits = defineEmits(["removeItem"]);

onMounted(() => {
  //判断props.defaultValue是数组还是对象
  if (arr.value) {
    setArrayValue(arr.value);
  } else {
    //遍历对象 依次将key 赋值key和value
    for (const key in obj.value) {
      if (obj && obj.value && key in obj.value) {
        inputs.value.push({
          key: key,
          value: obj.value[key],
        });
      } else {
        inputs.value.push({
          key: key,
          value: "",
        });
      }
    }
  }
});

interface InputItem {
  key: string;
  value: string;
}

//传入一个数组 转换并设置inputs的值
const setArrayValue = (arrData: any[]) => {
  //遍历数组 依次将defaultValue的props.keyName和 props.valueName赋值给inuts
  for (let i = 0; i < arrData.length; i++) {
    inputs.value.push({
      key: arrData[i][props.keyName ?? "key"],
      value: arrData[i][props.valueName ?? "value"],
    });
  }
};

// Method to add a new set of inputs
const addInput = () => {
  inputs.value.push({ key: "", value: "" });
};

// Method to remove a set of inputs
const removeInput = (index: number) => {
  let k = inputs.value[index].key;
  let v = inputs.value[index].value;
  inputs.value.splice(index, 1);
  emits("removeItem", k, v, index);
};

// Expose a method 返回数组
//数组中对象key和value可以通过props.keyName和props.valueName来修改，默认为key和value
const getArray = () => {
  if (!inputs.value || inputs.value.length == 0) return [];
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
  if (!inputs.value || inputs.value.length == 0) return {};
  let result: any = {};
  for (let i = 0; i < inputs.value.length; i++) {
    result[inputs.value[i].key] = inputs.value[i].value;
  }
  return result;
};
const getObjects = () => {
  if (!inputs.value || inputs.value.length == 0) return [];
  let result: any[] = [];
  for (let i = 0; i < inputs.value.length; i++) {
    let item: any = {};
    item[props.keyName ?? "key"] = inputs.value[i].key;
    item[props.valueName ?? "value"] = inputs.value[i].value;
    result.push(item);
  }
  return result;
};

const getArrayObject = () => {
  let keyName = props.keyName || "key";
  let valueName = props.valueName || "value";
  let result = [];
  for (let i = 0; i < inputs.value.length; i++) {
    let item = [];
    item.push(props.outKeyName ?? "meta");
    item.push({
      [keyName]: inputs.value[i].key,
      [valueName]: inputs.value[i].value,
    });
    result.push(item);
  }
  return result;
};
defineExpose({ getArray, getObject, setArrayValue });
</script>
