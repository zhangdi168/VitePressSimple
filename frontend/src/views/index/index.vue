<template>
  <div class="bg-white" style="height: 100vh">
    <!--    顶部菜单栏-->
    <!--    <index-bar></index-bar>-->

    <!-- 内容主体 -->
    <div class="border-t">
      <q-splitter v-model="splitterModel" unit="px">
        <template v-slot:before>
          <!--          文章列表区域-->
          <index-article></index-article>
        </template>

        <template v-slot:after>
          <!--          编辑器区域-->
          <index-editor></index-editor>
        </template>
      </q-splitter>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { ref } from "vue";
import IndexEditor from "./indexEditor.vue";
import IndexArticle from "./indexArticle.vue";

const splitterModel = ref(248);
</script>
