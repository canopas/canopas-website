<template>
  <div>
    <Header />
    <div>
      <LandingSection />
      <HowItAllStartedSectionMobile class="block lg:hidden" />
      <HowItAllStartedSection class="hidden lg:block" />
      <AboutusVirtue />
      <WithCanopasSection ref="withcanopas" />
      <ClientReviewSection />
      <CTASection ref="cta" />
    </div>
    <NewFooter />
  </div>
</template>

<script setup>
import Header from "@/components/partials/NewHeader.vue";
import LandingSection from "@/components/about/LandingSection.vue";
import HowItAllStartedSection from "@/components/about/HowItAllStartedSection.vue";
import config from "@/config.js";
import { elementInViewPort } from "@/utils.js";
import { defineAsyncComponent } from "vue";

const HowItAllStartedSectionMobile = defineAsyncComponent(
  () => import("@/components/about/HowItAllStartedSectionMobile.vue"),
);
const AboutusVirtue = defineAsyncComponent(
  () => import("@/components/about/AboutusVirtue.vue"),
);
const WithCanopasSection = defineAsyncComponent(
  () => import("@/components/about/WithCanopas.vue"),
);
const ClientReviewSection = defineAsyncComponent(
  () => import("@/components/home-new/ClientReviewSection.vue"),
);
const CTASection = defineAsyncComponent(
  () => import("@/components/about/CTASection.vue"),
);
const NewFooter = defineAsyncComponent(
  () => import("@/components/partials/NewFooter.vue"),
);
const { $mixpanel } = useNuxtApp();

const withcanopas = ref(null);
const cta = ref(null);

const seoData = config.ABOUT_SEO_META_DATA;
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
  withcanopas: "view_about_phases",
  cta: "view_about_cta",
};
let elements;
onMounted(() => {
  elements = ref({
    withcanopas: withcanopas,
    cta: cta,
  });
  window.addEventListener("scroll", sendEvent);
  $mixpanel.track("view_about_page");
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
