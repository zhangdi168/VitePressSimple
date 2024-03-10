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
      component: () => import("../views/AboutView.vue"),
    },
    {
      path: "/setting",
      name: "setting",
      component: () => import("../views/SettingView.vue"),
    },
  ],
});

export default router;
