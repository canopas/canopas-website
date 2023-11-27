<template>
  <div>
    <Header />
    <LandingSection />
    <CtaSection />
    <SuccessStory />
    <BlogSection />
    <CtaSection2 />
    <NewFooter ref="flutterFooter" />
  </div>
</template>

<script setup>
import Header from "@/components/partials/NewHeader.vue";
import config from "@/config.js";
import { defineAsyncComponent } from "vue";
import { elementInViewPort } from "@/utils.js";
import LandingSection from "@/components/flutter-app-development/LandingSection.vue";
const CtaSection = defineAsyncComponent(
  () => import("@/components/flutter-app-development/CtaSection.vue"),
);
const SuccessStory = defineAsyncComponent(
  () => import("@/components/flutter-app-development/SuccessStorySection.vue"),
);
const BlogSection = defineAsyncComponent(
  () => import("@/components/flutter-app-development/BlogSection.vue"),
);
const CtaSection2 = defineAsyncComponent(
  () => import("@/components/flutter-app-development/CtaSection2.vue"),
);
const NewFooter = defineAsyncComponent(
  () => import("@/components/partials/NewFooter.vue"),
);
const { $mixpanel } = useNuxtApp();
const footer = ref(null);
const seoData = config.FLUTTER_APP_DEVELOPMENT_SEO_META_DATA;
useSeoMeta({
  robots: "noindex, nofollow",
  title: seoData.title,
  description: seoData.description,
  ogTitle: seoData.title,
  ogType: seoData.type,
  ogUrl: seoData.url,
  ogImage: seoData.image,
});

let event = "";
let events = {
  footer: "view_flutter_app_development_footer",
};
let elements;
onMounted(() => {
  elements = ref({
    footer: footer,
  });
  window.addEventListener("scroll", sendEvent);
  $mixpanel.track("view_flutter_app_development_page");
});
onUnmounted(() => {
  window.removeEventListener("scroll", sendEvent);
});
function sendEvent() {
  const newEvent = events[elementInViewPort(elements.value)];
  if (newEvent && event !== newEvent) {
    event = newEvent;
    $mixpanel.track(event);
  }
}
</script>
