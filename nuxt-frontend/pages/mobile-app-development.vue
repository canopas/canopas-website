<template>
  <div>
    <Header />
    <div class="overflow-hidden">
      <div class="bg-[#262626] md:bg-[#1A1A1A]">
        <LandingSection />
        <Services />
      </div>
      <Portfolio />
      <CTASection />
      <DevelopmentProcess ref="developmentprocess" />
      <div class="bg-black-core/[0.85]">
        <ClientReview />
        <CTASection2 />
      </div>
    </div>
    <NewFooter ref="footer" class="mt-[-15px] md:mt-auto" />
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
const CTASection = defineAsyncComponent(
  () => import("@/components/mobile-app-development/CTASection.vue"),
);
const DevelopmentProcess = defineAsyncComponent(
  () => import("@/components/mobile-app-development/DevelopmentProcess.vue"),
);
const ClientReview = defineAsyncComponent(
  () => import("@/components/mobile-app-development/ClientReview.vue"),
);
const CTASection2 = defineAsyncComponent(
  () => import("@/components/mobile-app-development/CTASection2.vue"),
);
const NewFooter = defineAsyncComponent(
  () => import("@/components/partials/NewFooter.vue"),
);

const { $mixpanel } = useNuxtApp();

const developmentprocess = ref(null);
const footer = ref(null);

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
  footer: "view_mobileapp_development_footer",
};

let elements;

onMounted(() => {
  elements = ref({
    developmentprocess: developmentprocess,
    footer: footer,
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
