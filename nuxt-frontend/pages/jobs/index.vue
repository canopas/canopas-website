<template>
  <div>
    <Header />
    <LandingView v-on:scroll-to-career="scrollToCareer" />
    <VirtuesView />
    <LifeAtCanopas />
    <PerksAndBenefits v-on:scroll-to-career="scrollToCareer" ref="perks" />
    <WhyCanopas
      class="hidden md:block"
      v-on:add-animation="handleAnimationOnScroll"
    />
    <WhyCanopasMobile
      class="block md:hidden"
      v-on:add-animation="handleAnimationOnScroll"
    />
    <Career id="career" />
    <FaqSection v-on:add-animation="handleAnimationOnScroll" ref="faq" />
    <NewFooter />
  </div>
</template>

<script setup>
import Header from "@/components/partials/NewHeader.vue";
import LandingView from "@/components/jobs/LandingView.vue";
import { defineAsyncComponent } from "vue";
import config from "@/config.js";
import { elementInViewPort, handleAnimationOnScroll } from "@/utils.js";

const VirtuesView = defineAsyncComponent(
  () => import("@/components/jobs/VirtuesView.vue"),
);
const LifeAtCanopas = defineAsyncComponent(
  () => import("@/components/jobs/LifeAtCanopas.vue"),
);
const Career = defineAsyncComponent(
  () => import("@/components/jobs/CareerView.vue"),
);
const PerksAndBenefits = defineAsyncComponent(
  () => import("@/components/jobs/PerksAndBenefits.vue"),
);
const WhyCanopas = defineAsyncComponent(
  () => import("@/components/jobs/WhyCanopas.vue"),
);
const WhyCanopasMobile = defineAsyncComponent(
  () => import("@/components/jobs/WhyCanopasMobile.vue"),
);
const FaqSection = defineAsyncComponent(
  () => import("@/components/jobs/FaqSection.vue"),
);
const NewFooter = defineAsyncComponent(
  () => import("@/components/partials/NewFooter.vue"),
);

const { $mixpanel } = useNuxtApp();

const perks = ref(null);
const faq = ref(null);

const seoData = config.JOBS_SEO_META_DATA;
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
  perks: "view_perks_benefits",
  faq: "view_jobs_faq",
};

let elements;

onMounted(() => {
  elements = ref({
    perks: perks,
    faq: faq,
  });
  window.addEventListener("scroll", sendEvent);
  $mixpanel.track("view_jobs_page");
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
function scrollToCareer() {
  let careerDiv = document.getElementById("career");
  window.scrollTo({ top: careerDiv.offsetTop, behavior: "smooth" });
}
</script>
