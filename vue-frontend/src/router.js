import {
  createMemoryHistory,
  createRouter as _createRouter,
  createWebHistory,
} from "vue-router";

const Error404Page = () => import("@/components/error404/index.vue");
import routes from "~pages";

routes.push({
  path: "/:pathMatch(.*)*",
  name: "Error404Page",
  component: Error404Page,
});

export function createRouter() {
  return _createRouter({
    // use appropriate history implementation for server/client
    // import.meta.env.SSR is injected by Vite.
    history: import.meta.env.SSR ? createMemoryHistory() : createWebHistory(),
    routes,
    scrollBehavior(to, from, savedPosition) {
      if (savedPosition) {
        return savedPosition;
      } else {
        return { top: 0, behavior: "instant" };
      }
    },
  });
}
