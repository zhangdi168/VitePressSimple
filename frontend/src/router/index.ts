import { createRouter, createWebHashHistory } from "vue-router";
import HomeView from "../views/index/index.vue";

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: "/",
      name: "home",
      component: HomeView,
    },
    {
      path: "/about",
      name: "about",
      component: () => import("@/views/about/about.vue"),
    },
    {
      path: "/projectSetting",
      name: "projectSetting",
      component: () => import("@/views/project/project.vue"),
    },
    {
      path: "/preview",
      name: "preview",
      component: () => import("@/views/previews/preview.vue"),
    },
    {
      path: "/setting",
      name: "setting",
      component: () => import("@/views/setting/setting.vue"),
    },
  ],
});

export default router;
