import { createWebHistory, createRouter } from "vue-router";
import Config from "@/config.js";

// Don't lazy load home screen
import HomeScreen from "@/components/HomeScreen.vue";
import NewHomeScreen from "@/components/NewHomeScreen.vue";

const Error404Page = () => import("@/components/Error404Page.vue");
const ContactScreen = () => import("@/components/ContactScreen.vue");
const JobsScreen = () => import("@/components/JobsScreen.vue");
const ApplyForCareerPage = () => import("@/components/ApplyForCareer.vue");
const CareerDetailsPage = () => import("@/components/CareerDetails.vue");

const routes = [
  {
    path: "/contact",
    name: "ContactScreen",
    component: ContactScreen,
  },
  {
    path: "/:catchAll(.*)",
    name: "Error404Page",
    component: Error404Page,
  },
  {
    path: "/jobs",
    name: "JobsScreen",
    component: JobsScreen,
  },
  {
    path: "/jobs/:id",
    name: "CareerDetailsPage",
    component: CareerDetailsPage,
  },
  {
    path: "/jobs/apply/:id",
    name: "ApplyForCareerPage",
    component: ApplyForCareerPage,
  },
];

if (Config.IS_SHOW_NEW_HOME_PAGE) {
  routes.push({
    path: "/",
    name: "NewHomeScreen",
    component: NewHomeScreen,
  });
} else {
  routes.push({
    path: "/",
    name: "HomeScreen",
    component: HomeScreen,
  });
}

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
