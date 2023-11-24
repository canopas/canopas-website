<template>
  <div>
    <Header />
    <div class="xll:bg-black-core/[90%] 3xl:bg-white">
      <LandingSection />
    </div>
    <DevelopmentSection />
    <CaseStudy ref="casestudy" />
    <PinkCtaSection class="hidden lg:block" />
    <SuccessStorySection />
    <BlackCtaSection />
    <BlogSection />
    <FaqSection />
    <CtaSection ref="cta" />
    <NewFooter />
  </div>
</template>

<script setup>
import Header from "@/components/partials/NewHeader.vue";
import config from "@/config.js";
import LandingSection from "@/components/android-app-development/LandingSection.vue";
import { elementInViewPort } from "@/utils.js";
import { defineAsyncComponent } from "vue";

const DevelopmentSection = defineAsyncComponent(
  () => import("@/components/android-app-development/DevelopmentSection.vue"),
);
const CaseStudy = defineAsyncComponent(
  () => import("@/components/android-app-development/CaseStudySection.vue"),
);
const SuccessStorySection = defineAsyncComponent(
  () => import("@/components/android-app-development/SuccessStorySection.vue"),
);
const FaqSection = defineAsyncComponent(
  () => import("@/components/android-app-development/FaqSection.vue"),
);
const PinkCtaSection = defineAsyncComponent(
  () => import("@/components/android-app-development/PinkCtaSection.vue"),
);
const CtaSection = defineAsyncComponent(
  () => import("@/components/android-app-development/CtaSection.vue"),
);
const BlackCtaSection = defineAsyncComponent(
  () => import("@/components/android-app-development/BlackCtaSection.vue"),
);
const BlogSection = defineAsyncComponent(
  () => import("@/components/android-app-development/BlogSection.vue"),
);
const NewFooter = defineAsyncComponent(
  () => import("@/components/partials/NewFooter.vue"),
);
const { $mixpanel } = useNuxtApp();

const casestudy = ref(null);
const cta = ref(null);

const seoData = config.ANDRIOD_APP_DEVELOPMENT_SEO_META_DATA;
useSeoMeta({
  title: seoData.title,
  description: seoData.description,
  ogTitle: seoData.title,
  ogType: seoData.type,
  ogUrl: seoData.url,
  ogImage: seoData.image,
});

let event = "";
let events = {
  casestudy: "view_android_development_casestudy_sections",
  cta: "view_android_development_cta",
};

let elements;

onMounted(() => {
  elements = ref({
    casestudy: casestudy,
    cta: cta,
  });
  window.addEventListener("scroll", sendEvent);
  $mixpanel.track("view_android_development_page");
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
