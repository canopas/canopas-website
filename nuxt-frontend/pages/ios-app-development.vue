<template>
  <div>
    <Header v-if="isHeaderVisible" />
    <LandingSection />
    <DevelopmentSection />
    <CaseStudySection @sectionFullscreen="handleSectionFullscreen" />
    <CtaSection />
    <SuccessStorySection ref="iosSuccessstory" />
    <BlogSection />
    <CtaSection2 />
    <FaqSection ref="iosFaq" />
    <NewFooter />
  </div>
</template>

<script setup>
import Header from "@/components/partials/NewHeader.vue";
import config from "@/config.js";
import { defineAsyncComponent } from "vue";
import { elementInViewPort } from "@/utils.js";
import LandingSection from "@/components/ios-app-development/LandingSection.vue";
import DevelopmentSection from "@/components/ios-app-development/DevelopmentSection.vue";

const CaseStudySection = defineAsyncComponent(
  () => import("@/components/ios-app-development/CaseStudySection.vue"),
);
const SuccessStorySection = defineAsyncComponent(
  () => import("@/components/ios-app-development/SuccessStory.vue"),
);
const BlogSection = defineAsyncComponent(
  () => import("@/components/ios-app-development/BlogSection.vue"),
);
const FaqSection = defineAsyncComponent(
  () => import("@/components/ios-app-development/FaqSection.vue"),
);
const CtaSection = defineAsyncComponent(
  () => import("@/components/ios-app-development/CtaSection.vue"),
);
const CtaSection2 = defineAsyncComponent(
  () => import("@/components/ios-app-development/CtaSection2.vue"),
);
const NewFooter = defineAsyncComponent(
  () => import("@/components/partials/NewFooter.vue"),
);
const { $mixpanel } = useNuxtApp();

const iosSuccessstory = ref(null);
const iosFaq = ref(null);

const seoData = config.IOS_APP_DEVELOPMENT_SEO_META_DATA;
useSeoMeta({
  title: seoData.title,
  description: seoData.description,
  ogTitle: seoData.title,
  ogType: seoData.type,
  ogUrl: seoData.url,
  ogImage: seoData.image,
});

let isHeaderVisible = ref(true);
let event = "";
let events = {
  iosSuccessstory: "view_ios_app_development_success_story",
  iosFaq: "view_ios_app_development_faq_section",
};
let elements;

onMounted(() => {
  elements = ref({
    iosSuccessstory: iosSuccessstory,
    iosFaq: iosFaq,
  });
  window.addEventListener("scroll", sendEvent);
  $mixpanel.track("view_ios_development_page");
});

onUnmounted(() => {
  window.removeEventListener("scroll", sendEvent);
});
function handleSectionFullscreen(isFullscreen) {
  isHeaderVisible.value = isFullscreen;
}
function sendEvent() {
  const newEvent = events[elementInViewPort(elements.value)];
  if (newEvent && event !== newEvent) {
    event = newEvent;
    $mixpanel.track(event);
  }
}
</script>
