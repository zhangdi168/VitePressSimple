<template>
  <q-bar class="bg-white border-b">
    <div class="cursor-pointer non-selectable">
      文件
      <q-menu>
        <q-list dense style="min-width: 100px">
          <q-item clickable v-close-popup @click="createProject()">
            <q-item-section>新建项目</q-item-section>
          </q-item>

          <q-separator />
          <q-item clickable v-close-popup @click="openProject()">
            <q-item-section>打开项目</q-item-section>
          </q-item>

          <q-separator />
          <q-item clickable>
            <q-item-section>Preferences</q-item-section>
            <q-item-section side>
              <q-icon name="keyboard_arrow_right" />
            </q-item-section>

            <q-menu anchor="top end" self="top start">
              <q-list>
                <q-item v-for="n in 3" :key="n" dense clickable>
                  <q-item-section>Submenu Label</q-item-section>
                  <q-item-section side>
                    <q-icon name="keyboard_arrow_right" />
                  </q-item-section>
                  <q-menu auto-close anchor="top end" self="top start">
                    <q-list>
                      <q-item v-for="n in 3" :key="n" dense clickable>
                        <q-item-section>3rd level Label</q-item-section>
                      </q-item>
                    </q-list>
                  </q-menu>
                </q-item>
              </q-list>
            </q-menu>
          </q-item>

          <q-separator />

          <q-item clickable v-close-popup>
            <q-item-section>Quit</q-item-section>
          </q-item>
        </q-list>
      </q-menu>
    </div>

    <div class="q-ml-md cursor-pointer non-selectable">
      项目
      <q-menu auto-close>
        <q-list dense style="min-width: 100px">
          <q-item clickable>
            <q-item-section>Cut</q-item-section>
          </q-item>
          <q-item clickable>
            <q-item-section>Copy</q-item-section>
          </q-item>
          <q-item clickable>
            <q-item-section>Paste</q-item-section>
          </q-item>
          <q-separator />
          <q-item clickable>
            <q-item-section>Select All</q-item-section>
          </q-item>
        </q-list>
      </q-menu>
    </div>
  </q-bar>
</template>
<script setup lang="ts">
import CreateProjectPopup from "@/views/project/createProjectPopup.vue";
import { ref } from "vue";
import OpenProjectPopup from "@/views/project/openProjectPopup.vue";
</script>
