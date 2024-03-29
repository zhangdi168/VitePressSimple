<template>
  <div style="front-size: 1rem" class="flex justify-between items-center">
    <icon-park
      class="mr-1"
      fill="#333"
      strokeLinejoin="bevel"
      theme="outline"
      :type="props.icon"
    />
    {{ props.title }}
  </div>
</template>
<script setup lang="ts">
import { IconPark } from "@icon-park/vue-next/es/all";

interface Props {
  title: string;
  icon: string;
}

const props = defineProps<Props>();
</script>

<style scoped></style>
