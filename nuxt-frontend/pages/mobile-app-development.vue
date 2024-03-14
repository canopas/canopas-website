<template>
  <div>
    <Header />
    <div>
      <div class="bg-deep-charcoal 3xl:outer-container">
        <LandingSection />
        <Services />
      </div>
      <Portfolio />
      <CTA1 />
      <DevelopmentProcess ref="developmentprocess" />
      <div class="bg-deep-charcoal 3xl:outer-container">
        <ClientReview />
        <CTA5 />
      </div>
    </div>
    <NewFooter class="bg-gradient-to-b from-pink-0 to-pink-16" />
  </div>
</template>

<script setup>
import Header from "@/components/partials/NewHeader.vue";
import LandingSection from "@/components/mobile-app-development/LandingSection.vue";
import Services from "@/components/mobile-app-development/Services.vue";
import config from "@/config.js";
import { defineAsyncComponent } from "vue";
import { elementInViewPort } from "@/utils.js";

const Portfolio = defineAsyncComponent(
  () => import("@/components/mobile-app-development/Portfolio.vue"),
);
const CTA1 = defineAsyncComponent(() => import("@/components/CTA/CTA1.vue"));
const DevelopmentProcess = defineAsyncComponent(
  () => import("@/components/mobile-app-development/DevelopmentProcess.vue"),
);
const ClientReview = defineAsyncComponent(
  () => import("@/components/mobile-app-development/ClientReview.vue"),
);
const CTA5 = defineAsyncComponent(() => import("@/components/CTA/CTA5.vue"));
const NewFooter = defineAsyncComponent(
  () => import("@/components/partials/NewFooter.vue"),
);

const { $mixpanel } = useNuxtApp();

const developmentprocess = ref(null);
const cta = ref(null);

const seoData = config.MOBILE_APP_DEVELOPMENT_SEO_META_DATA;
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
  developmentprocess: "view_mobileapp_development_developmentprocess_section",
  cta: "view_mobileapp_development_cta",
};

let elements;

onMounted(() => {
  elements = ref({
    developmentprocess: developmentprocess,
    cta: cta,
  });
  window.addEventListener("scroll", sendEvent);
  $mixpanel.track("view_mobileapp_development_page");
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
