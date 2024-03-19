<template>
  <div class="flex flex-col space-y-2">
    <!-- 主菜单操作按钮 -->
    <button
      class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
      @click="addTopMenuItem"
    >
      Add Main Menu Item
    </button>

    <!-- 递归渲染主菜单及子菜单 -->
    <recursive-menu-item
      v-for="(item, index) in menuItems"
      :key="index"
      :item="item"
      @remove="removeMenuItem(index)"
      @addChild="addChildMenuItem(item)"
    />

    <!-- 新增子菜单项模态框（仅在有菜单项被选中时显示） -->
    <transition name="fade">
      <div v-if="selectedMenuItem" class="fixed inset-0 z-10 overflow-hidden">
        <div class="absolute inset-0 bg-gray-500 opacity-75"></div>
        <div
          class="relative mx-auto max-w-sm p-6 bg-white rounded-lg shadow-xl"
        >
          <h3 class="text-lg font-medium text-gray-900">Add Submenu Item</h3>
          <input
            class="border rounded p-1 mt-2"
            v-model="newChildMenuItem.text"
            type="text"
            placeholder="Submenu Text"
          />
          <input
            class="border rounded p-1 ml-2 mt-2"
            v-model="newChildMenuItem.link"
            type="text"
            placeholder="Link"
          />
          <button
            class="bg-green-500 hover:bg-green-700 text-white font-bold py-1 px-2 rounded mt-2"
            @click="addChildToSelectedMenuItem"
          >
            Add Submenu
          </button>
          <button
            class="bg-red-500 hover:bg-red-700 text-white font-bold py-1 px-2 rounded ml-2 mt-2"
            @click="closeAddChildModal"
          >
            Cancel
          </button>
        </div>
      </div>
    </transition>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from "vue";
import RecursiveMenuItem from "@/components/recursiveMenuItem.vue";
import { MenuItem } from "@/types/menu";

const menuItems = ref<MenuItem[]>([{ text: "Home", link: "/home" }]);

let selectedMenuItem: MenuItem | null = null;
const newChildMenuItem = reactive<MenuItem>({ text: "", link: "" });

// 新增顶级菜单项
function addTopMenuItem() {
  menuItems.value.push({ text: "New Menu Item", items: [] });
}

// 删除菜单项
function removeMenuItem(index: number) {
  menuItems.value.splice(index, 1);
}

// 新增子菜单项
function addChildMenuItem(parentItem: MenuItem) {
  parentItem.items ||= [];
  parentItem.items.push({ text: "New Submenu Item", items: [] });
  selectedMenuItem = parentItem;
}

// 将新建的子菜单项添加到选中的菜单项下
function addChildToSelectedMenuItem() {
  if (selectedMenuItem && newChildMenuItem.text) {
    selectedMenuItem.items?.push(newChildMenuItem);
    newChildMenuItem.text = "";
    newChildMenuItem.link = "";
    selectedMenuItem = null;
  }
}

// 关闭新增子菜单项模态框
function closeAddChildModal() {
  selectedMenuItem = null;
}
</script>
