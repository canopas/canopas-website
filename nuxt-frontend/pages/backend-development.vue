<template>
  <div>
    <Header />
    <LandingSection />
    <CaseStudySection />
    <PartnerWithUsSection />
    <SuccessStorySection />
    <BlogSection />
    <BlackCtaSection />
    <FaqSection />
    <NewFooter />
  </div>
</template>
<script>
import Header from "@/components/partials/NewHeader.vue";
import config from "@/config.js";
import { defineAsyncComponent } from "vue";
import { elementInViewPort } from "@/utils.js";
import LandingSection from "@/components/backend-development/LandingSection.vue";
const CaseStudySection = defineAsyncComponent(
  () => import("@/components/backend-development/CaseStudySection.vue"),
);
const PartnerWithUsSection = defineAsyncComponent(
  () => import("@/components/backend-development/PartnerWithUsSection.vue"),
);
const SuccessStorySection = defineAsyncComponent(
  () => import("@/components/backend-development/SuccessStorySection.vue"),
);
const BlogSection = defineAsyncComponent(
  () => import("@/components/backend-development/BlogSection.vue"),
);
const BlackCtaSection = defineAsyncComponent(
  () => import("@/components/backend-development/BlackCtaSection.vue"),
);
const FaqSection = defineAsyncComponent(
  () => import("@/components/backend-development/FaqSection.vue"),
);
const NewFooter = defineAsyncComponent(
  () => import("@/components/partials/NewFooter.vue"),
);

export default {
  beforeRouteEnter(to, from, next) {
    if (!config.SHOW_BACKEND_DEVELOPMENT_PAGE) {
      next({
        name: "Error404Page",
        params: { pathMatch: "/backend-development" },
      });
    } else {
      next();
    }
  },
  components: {
    Header,
    LandingSection,
    CaseStudySection,
    PartnerWithUsSection,
    SuccessStorySection,
    BlogSection,
    BlackCtaSection,
    FaqSection,
    NewFooter,
  },
  setup() {
    const { $mixpanel } = useNuxtApp();
    const footer = ref(null);
    const seoData = config.BACKEND_DEVELOPMENT_SEO_META_DATA;
    useSeoMeta({
      robots: "noindex,nofollow",
      title: seoData.title,
      description: seoData.description,
      ogTitle: seoData.title,
      ogType: seoData.type,
      ogUrl: seoData.url,
      ogImage: seoData.image,
    });
    let event = "";
    let events = {
      footer: "view_backend_development_footer",
    };
    let elements;
    onMounted(() => {
      elements = ref({
        footer: footer,
      });
      window.addEventListener("scroll", sendEvent);
      $mixpanel.track("view_backend_development_page");
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
  },
};
</script>
