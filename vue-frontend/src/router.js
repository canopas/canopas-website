import {
  createMemoryHistory,
  createRouter as _createRouter,
  createWebHistory,
} from "vue-router";
const Error404Page = () => import("@/components/error404/index.vue");
import routes from "~pages";
function createRouter() {
  return _createRouter({
    history: (import.meta.env.SSR ? createMemoryHistory : createWebHistory)(),
    routes: routes,
    scrollBehavior(to, from, savedPosition) {
      return savedPosition || { top: 0, behavior: "instant" };
    },
  });
}
routes.push({
  path: "/:pathMatch(.*)*",
  name: "Error404Page",
  component: Error404Page,
});
export { createRouter };
