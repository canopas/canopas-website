import { createWebHistory, createRouter } from "vue-router";
import HomeScreen from "@/components/HomeScreen.vue";
import ContactScreen from "@/components/ContactScreen.vue";

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
];

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior(_to, _from, savedPosition) {
    if (savedPosition) {
      return savedPosition;
    }
    return { left: 0, top: 0 };
  },
});

router.beforeEach(() => {
  return { top: 0 };
});

export default router;
