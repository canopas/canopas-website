<template>
  <div>
    <Header />
    <div>
      <LandingSection :name="clientName" />
      <BenefitSection
        class="overflow-y-hidden 2xl:overflow-y-visible 3xl:overflow-y-hidden"
      />
      <HappyClientSection ref="clientreview" />
      <ScheduleMeeting />
    </div>
    <NewFooter ref="footer" />
  </div>
</template>

<script setup>
import { defineAsyncComponent } from "vue";
import config from "@/config.js";
import Header from "@/components/partials/NewHeader.vue";
import LandingSection from "@/components/contact/thank-you/LandingSection.vue";
import BenefitSection from "@/components/contact/thank-you/BenefitSection.vue";
import { elementInViewPort } from "@/utils.js";
const HappyClientSection = defineAsyncComponent(
  () => import("@/components/contact/thank-you/HappyClient.vue"),
);
const ScheduleMeeting = defineAsyncComponent(
  () => import("@/components/contact/thank-you/ScheduleMeeting.vue"),
);
const NewFooter = defineAsyncComponent(
  () => import("@/components/partials/NewFooter.vue"),
);
const { $mixpanel } = useNuxtApp();

const clientreview = ref(null);
const footer = ref(null);

const seoData = config.CLIENT_THANKYOU_SEO_META_DATA;
useSeoMeta({
  title: seoData.title,
  description: seoData.description,
  ogTitle: seoData.title,
  ogType: seoData.type,
  ogUrl: seoData.url,
  ogImage: seoData.image,
});

let clientName = "";
let event = "";
let events = {
  clientreview: "view_thankyou_clientreview_section",
  footer: "view_thankyou_footer",
};
let elements;

onMounted(() => {
  elements = ref({
    clientreview: clientreview,
    footer: footer,
  });
  window.addEventListener("scroll", sendEvent);
  $mixpanel.track("view_thankyou_page");
});
onUnmounted(() => {
  window.removeEventListener("scroll", sendEvent);
});
onBeforeMount(() => {
  let cName = localStorage.getItem("client-name");
  if (cName) {
    clientName = JSON.parse(cName);
  } else {
    navigateTo({
      name: "Error404Page",
      params: { pathMatch: ["thank-you"] },
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
