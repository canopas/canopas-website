import { createWebHistory, createRouter } from "vue-router";
import Config from "@/config.js";

const HomeScreen = () => import("@/components/HomeScreen.vue");
const Error404Page = () => import("@/components/Error404Page.vue");
const ContactScreen = () => import("@/components/ContactScreen.vue");
const JobsScreen = () => import("@/components/JobsScreen.vue");

const routes = [
  {
    path: "/",
    name: "HomeScreen",
    component: HomeScreen,
  },
  {
    path: "/contact",
    name: "ContactScreen",
    component: ContactScreen,
  },
  {
    path: "/jobs",
    name: "JobsScreen",
    component: Config.IS_SHOW_JOBS? JobsScreen : Error404Page
  },
  {
    path: "/:catchAll(.*)",
    name: "Error404Page",
    component: Error404Page,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition;
    } else {
      return { top: 0, behavior: "instant" };
    }
  },
});

export default router;
