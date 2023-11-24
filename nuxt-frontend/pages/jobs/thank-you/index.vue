<template>
  <div>
    <Header />
    <LandingSection :name="applicantName" />
    <InterviewProcess />
    <VirtueView />
    <LifeAtCanopasVue ref="lifeatcanopas" />
    <LatestBlog />
    <ContributionSection ref="contribution" />
    <NewFooter />
  </div>
</template>
<script setup>
import Header from "@/components/partials/NewHeader.vue";
import LandingSection from "@/components/contact/thank-you/LandingSection.vue";
import InterviewProcess from "@/components/jobs/thank-you/InterviewProcess.vue";
import { defineAsyncComponent } from "vue";
import { elementInViewPort } from "@/utils.js";
import config from "@/config.js";

const VirtueView = defineAsyncComponent(
  () => import("@/components/jobs/VirtuesView.vue"),
);
const LifeAtCanopasVue = defineAsyncComponent(
  () => import("@/components/jobs/PerksAndBenefits.vue"),
);
const LatestBlog = defineAsyncComponent(
  () => import("@/components/jobs/thank-you/LatestBlog.vue"),
);
const ContributionSection = defineAsyncComponent(
  () => import("@/components/jobs/thank-you/ContributionSection.vue"),
);
const NewFooter = defineAsyncComponent(
  () => import("@/components/partials/NewFooter.vue"),
);

const { $mixpanel } = useNuxtApp();

const lifeatcanopas = ref(null);
const contribution = ref(null);

const seoData = config.JOBS_THANKYOU_SEO_META_DATA;
useSeoMeta({
  title: seoData.title,
  description: seoData.description,
  ogTitle: seoData.title,
  ogType: seoData.type,
  ogUrl: seoData.url,
  ogImage: seoData.image,
});

let applicantName = "";
let event = "";
let events = {
  lifeatcanopas: "view_jobs_thankyou_lifeatcanopas_section",
  contribution: "view_jobs_thankyou_contribution_section",
};
let elements;

onMounted(() => {
  elements = ref({
    lifeatcanopas: lifeatcanopas,
    contribution: contribution,
  });
  window.addEventListener("scroll", sendEvent);
  $mixpanel.track("view_jobs_thankyou_page");
});
onUnmounted(() => {
  window.removeEventListener("scroll", sendEvent);
});
onBeforeMount(() => {
  if (localStorage.getItem("applicant-name")) {
    applicantName = JSON.parse(localStorage.getItem("applicant-name"));
  } else {
    navigateTo({
      name: "Error404Page",
      params: { pathMatch: ["jobs", "thank-you"] },
    });
  }
});
function sendEvent() {
  const newEvent = events[elementInViewPort(elements.value)];
  if (newEvent && event !== newEvent) {
    event = newEvent;
    $mixpanel.track(event);
  }
}
</script>
