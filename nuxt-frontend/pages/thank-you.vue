<template>
  <div v-if="!isNotFound">
    <Header />
    <div>
      <LandingSection :name="clientName" />
      <BenefitSection
        class="overflow-y-hidden 2xl:overflow-y-visible 3xl:overflow-y-hidden"
        ref="benefits"
      />
      <HappyClientSection />
      <InformationPoints />
    </div>
    <NewFooter />
  </div>
  <NotFound v-else />
  <!-- Show Calendly Iframe -->
  <div v-if="openCalendlyIframeModal">
    <transition name="modal">
      <div
        class="modal-mask fixed top-0 left-0 w-full h-full table mask bg-[#00000080] z-[5]"
      >
        <div
          class="flex mx-auto left-auto sm:mx-auto h-full login-modal modal-xl"
          role="document"
        >
          <div
            class="relative flex flex-col w-full border-1 border-gray border-solid rounded-md bg-white bg-clip-padding outline-0"
          >
            <div class="relative flex-auto">
              <CalendlyIframe
                class="w-full h-screen overflow-hidden"
                @close="closeCalendlyIframeModal"
              />

              <button
                type="button"
                class="absolute right-2 md:right-5 top-[5.5rem] md:top-28 close modal-close-btn border-none text-pink-300 !bg-pink-100 rounded-full md:text-2xl h-7 w-7 md:h-10 md:w-10 !bg-opacity-80 font-bold focus:outline-none"
                data-dismiss="modal"
                aria-label="Close"
              >
                <span
                  aria-hidden="true"
                  @click="closeCalendlyIframeModal"
                  class="h-6 w-full flex items-center justify-center"
                  >&times;</span
                >
              </button>
            </div>
          </div>
        </div>
      </div>
    </transition>
  </div>
</template>

<script setup>
import { defineAsyncComponent } from "vue";
import config from "@/config.js";
import Header from "@/components/partials/NewHeader.vue";
import LandingSection from "@/components/contact/thank-you/LandingSection.vue";
import { elementInViewPort } from "@/utils.js";
import CalendlyIframe from "@/components/contact/CalendlyIframe.vue";
import NotFound from "@/components/error404/index.vue";

const BenefitSection = defineAsyncComponent(
  () => import("@/components/contact/thank-you/BenefitSection.vue"),
);
const HappyClientSection = defineAsyncComponent(
  () => import("@/components/contact/thank-you/HappyClient.vue"),
);
const InformationPoints = defineAsyncComponent(
  () => import("@/components/contact/thank-you/InformationPoints.vue"),
);
const NewFooter = defineAsyncComponent(
  () => import("@/components/partials/NewFooter.vue"),
);

const { $mixpanel } = useNuxtApp();

const benefits = ref(null);

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
  benefits: "view_thankyou_benefits_section",
};

let elements;

let openCalendlyIframeModal = ref(false);
let isNotFound = ref(false);

onMounted(() => {
  elements = ref({
    benefits: benefits,
  });
  window.addEventListener("scroll", sendEvent);
  $mixpanel.track("view_thankyou_page");
  if (clientName != "") {
    openCalendlyIframe();
  }
});

onUnmounted(() => {
  openCalendlyIframeModal.value = false;
  localStorage.removeItem("client-name");
  document.body.classList.remove("overflow-hidden");
  window.removeEventListener("scroll", sendEvent);
});

onBeforeMount(() => {
  let cName = localStorage.getItem("client-name");
  if (cName) {
    clientName = JSON.parse(cName);
  } else {
    isNotFound.value = true;
  }
});

function sendEvent() {
  const newEvent = events[elementInViewPort(elements.value)];
  if (newEvent && event !== newEvent) {
    event = newEvent;
    $mixpanel.track(event);
  }
}

function openCalendlyIframe() {
  openCalendlyIframeModal.value = true;
  document.body.classList.add("overflow-hidden");
  $mixpanel.track("tap_schedule_meeting");
}

function closeCalendlyIframeModal() {
  openCalendlyIframeModal.value = false;
  document.body.classList.remove("overflow-hidden");
}
</script>
