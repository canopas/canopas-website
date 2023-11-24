<template>
  <div>
    <Header />
    <CaseStudy />
    <UserReview
      ref="userReview"
      class="pb-0 sm:pb-10 md:pb-8 lg:pb-20 sm:top-10"
    />
    <CTASection ref="cta" />
    <NewFooter />
  </div>
</template>

<script setup>
import config from "@/config.js";
import Header from "@/components/partials/NewHeader.vue";
import CaseStudy from "@/components/home-new/CaseStudy.vue";
import { defineAsyncComponent } from "vue";
import { elementInViewPort } from "@/utils.js";

const UserReview = defineAsyncComponent(
  () => import("@/components/home-new/UserReview.vue"),
);
const CTASection = defineAsyncComponent(
  () => import("@/components/home-new/CTASection.vue"),
);
const NewFooter = defineAsyncComponent(
  () => import("@/components/partials/NewFooter.vue"),
);
const { $mixpanel } = useNuxtApp();

const userReview = ref(null);
const cta = ref(null);

const seoData = config.PORTFOLIO_SEO_META_DATA;
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
  userReview: "view_portfolio_user_review_section",
  cta: "view_portfolio_cta_section",
};

let elements;

onMounted(() => {
  elements = ref({
    userReview: userReview,
    cta: cta,
  });
  window.addEventListener("scroll", sendEvent);
  $mixpanel.track("view_portfolio_page");
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
