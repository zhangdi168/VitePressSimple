<template>
  <div class="flex justify-center">
    <a-button
      :class="props.btnClass"
      class="bg-blue-600 hover:bg-blue-500 text-white flex justify-center items-center"
    >
      <icon-park
        class="mr-1"
        strokeLinejoin="bevel"
        theme="outline"
        :type="props.icon"
      />
      {{ props.text }}
    </a-button>
  </div>
</template>
<script setup lang="ts">
import { IconPark } from "@icon-park/vue-next/es/all";
import { defineProps } from "vue";

interface Props {
  text: string;
  icon: string;
  btnClass?: string;
}

const props = defineProps<Props>();
</script>
<style scoped></style>
