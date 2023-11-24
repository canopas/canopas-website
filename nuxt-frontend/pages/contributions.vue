<template>
  <div>
    <Header />
    <div class="overflow-hidden">
      <MobileLandingSection class="block md:hidden" />
      <DesktopLandingSection class="hidden md:block" />
      <GithubContribution />
    </div>
    <div>
      <WeeklyUpdateSection class="overflow-hidden md:overflow-visible" />
      <WhatsTrending class="overflow-hidden" />
      <Favourite ref="favourite" />
      <AnimatedCreation />
      <UnitTest class="md:overflow-hidden" />
      <DevOps />
      <ExploreDesign class="overflow-hidden" ref="design" />
    </div>
    <NewFooter />
  </div>
</template>

<script setup>
import { defineAsyncComponent } from "vue";
import Header from "@/components/partials/NewHeader.vue";
import config from "@/config.js";
import MobileLandingSection from "@/components/contributions/MobileLanding.vue";
import DesktopLandingSection from "@/components/contributions/DesktopLanding.vue";
import GithubContribution from "@/components/contributions/GithubContribution.vue";
import { elementInViewPort } from "@/utils.js";
const WeeklyUpdateSection = defineAsyncComponent(
  () => import("@/components/contributions/WeeklyUpdate.vue"),
);
const WhatsTrending = defineAsyncComponent(
  () => import("@/components/contributions/WhatsTrending.vue"),
);
const Favourite = defineAsyncComponent(
  () => import("@/components/contributions/Favourite.vue"),
);
const AnimatedCreation = defineAsyncComponent(
  () => import("@/components/contributions/AnimatedCreation.vue"),
);
const UnitTest = defineAsyncComponent(
  () => import("@/components/contributions/UnitTest.vue"),
);
const DevOps = defineAsyncComponent(
  () => import("@/components/contributions/DevOps.vue"),
);
const ExploreDesign = defineAsyncComponent(
  () => import("@/components/contributions/DesignExplore.vue"),
);
const NewFooter = defineAsyncComponent(
  () => import("@/components/partials/NewFooter.vue"),
);
const { $mixpanel } = useNuxtApp();

const favourite = ref(null);
const design = ref(null);

const seoData = config.CONTRIBUTION_SEO_META_DATA;
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
  favourite: "view_favourite_contribution",
  design: "view_design_contribution",
};

let elements;

onMounted(() => {
  elements = ref({
    favourite: favourite,
    design: design,
  });
  window.addEventListener("scroll", sendEvent);
  $mixpanel.track("view_contribution_page");
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
