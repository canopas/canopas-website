<template>
  <div v-if="response">
    <Header />
    <div class="parallax w-full p-0 m-auto 4xl:w-[2300px]" id="response">
      <LandingSection v-bind:json="details.landing" />
      <VideoSection v-bind:json="details.video" />
      <BrandingSection v-bind:json="details.branding" />
      <DesignSection v-bind:json="details.design" />
      <ElementSection v-bind:json="details.element" />
      <FooterSection v-bind:json="details.footer" />
      <NewCTASection />
      <NewFooter />
    </div>
  </div>
  <div v-else><ErrorPage /></div>
</template>

<script setup>
import Header from "@/components/partials/NewHeader.vue";
import NewFooter from "@/components/partials/NewFooter.vue";
import LandingSection from "@/components/portfolio/LandingSection.vue";
import VideoSection from "@/components/portfolio/VideoSection.vue";
import BrandingSection from "@/components/portfolio/BrandingSection.vue";
import DesignSection from "@/components/portfolio/DesignSection.vue";
import FooterSection from "@/components/portfolio/FooterSection.vue";
import NewCTASection from "@/components/portfolio/NewCTASection.vue";
import ElementSection from "@/components/portfolio/ElementSection.vue";
import luxeradioResponse from "@/json-data/luxeradio-portfolio.js";
import justlyResponse from "@/json-data/justly-portfolio.js";
import tognessResponse from "@/json-data/togness-portfolio.js";
import ErrorPage from "@/error.vue";
import { elementInViewPort } from "@/utils.js";
import config from "@/config.js";

import { useRoute } from "vue-router";

const { $mixpanel } = useNuxtApp();

const route = useRoute();
const id = route.params.id;

const details = ref("");
const response = ref(null);
const ctaRef = ref(null);
const portfolioRef = ref(null);
const event = ref("");

const events = {
  luxepagecta: "view_luxe_page_cta",
  tognesspagecta: "view_togness_page_cta",
  justlypagecta: "view_justly_page_cta",
};

getData(id);

useSeoMeta({
  title: response.value.seoData.title,
  description: response.value.seoData.description,
  ogTitle: response.value.seoData.title,
  ogUrl: response.value.seoData.url,
  ogType: "Website",
  ogImage: config.OG_IMAGE_URL,
});

function getData(routeId) {
  switch (routeId) {
    case luxeradioResponse.name:
      response.value = luxeradioResponse;
      ctaRef.value = "luxepagecta";
      break;

    case justlyResponse.name:
      response.value = justlyResponse;
      ctaRef.value = "justlypagecta";
      break;

    case tognessResponse.name:
      response.value = tognessResponse;
      ctaRef.value = "tognesspagecta";
      break;
  }

  if (response.value) {
    details.value = response.value.detail;
  }
}

function sendEvent() {
  const newEvent = events[elementInViewPort(this.$refs)];
  if (newEvent && event.value !== newEvent) {
    event.value = newEvent;
    $mixpanel.track(event);
  }
}

function scrollToTop() {
  portfolioRef.value.scrollTo({
    top: portfolioRef.value.offsetTop,
    behavior: "instant",
  });
}

onMounted(() => {
  portfolioRef.value = document.getElementById("response");
  portfolioRef.value.addEventListener("scroll", sendEvent);
  window.addEventListener("scroll", sendEvent);
  $mixpanel.track("view_" + id + "_portfolio_page");
});

onUnmounted(() => {
  portfolioRef.value.removeEventListener("scroll", sendEvent);
  window.removeEventListener("scroll", sendEvent);
});

onUpdated(() => {
  scrollToTop();
});

watch(
  id,
  (currentValue, oldValue) => {
    getData(currentValue);
  },
  { immediate: true, deep: true },
);
</script>

<style lang="postcss" scoped>
@screen lg {
  .parallax {
    perspective: 2px;
    @apply h-screen overflow-hidden overflow-y-auto;
  }
}
</style>
