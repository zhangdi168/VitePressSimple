<template>
  <div class="flex items-center space-x-2">
    <input
      class="border rounded p-1"
      v-model="item.text"
      type="text"
      placeholder="Menu Text"
    />
    <input
      v-if="item.link"
      class="border rounded p-1 ml-2"
      v-model="item.link"
      type="text"
      placeholder="Link"
    />

    <!-- 修改为使用事件触发子菜单的添加和删除 -->
    <button
      v-if="!item.items || item.items.length === 0"
      class="bg-yellow-500 hover:bg-yellow-700 text-white font-bold py-1 px-2 rounded ml-2"
      @click="$emit('addChild', item)"
    >
      Add Submenu
    </button>

    <!-- 子菜单列表 -->
    <div v-if="item.items && item.items.length > 0">
      <recursive-menu-item
        v-for="(subItem, subIndex) in item.items"
        :key="subIndex"
        :item="subItem"
        @remove="
          (subItem) => $emit('removeChild', { parent: item, child: subItem })
        "
        @addChild="
          (subItem) =>
            $emit('addChildToChild', { parent: item, child: subItem })
        "
      />
    </div>
  </div>
</template>

<script lang="ts">
import { MenuItem } from "@/types/menu";
import { PropType } from "vue";

export default {
  components: { RecursiveMenuItem },
  props: {
    item: Object as PropType<MenuItem>,
  },
};
</script>
