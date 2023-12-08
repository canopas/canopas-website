<template>
  <div>
    <Header />
    <LandingSection />
    <DevelopmentSection />
    <CaseStudySection />
    <CtaSection />
    <SuccessStory />
    <BlogSection />
    <CtaSection2 />
    <FaqSection />
    <NewFooter class="-mt-4 md:mt-auto" ref="flutterFooter" />
  </div>
</template>
<script>
import Header from "@/components/partials/NewHeader.vue";
import config from "@/config.js";
import { defineAsyncComponent } from "vue";
import { elementInViewPort } from "@/utils.js";
import LandingSection from "@/components/flutter-app-development/LandingSection.vue";
import DevelopmentSection from "@/components/flutter-app-development/DevelopmentSection.vue";
const CaseStudySection = defineAsyncComponent(
  () => import("@/components/flutter-app-development/CaseStudySection.vue"),
);
const CtaSection = defineAsyncComponent(
  () => import("@/components/flutter-app-development/CtaSection.vue"),
);
const SuccessStory = defineAsyncComponent(
  () => import("@/components/flutter-app-development/SuccessStorySection.vue"),
);
const BlogSection = defineAsyncComponent(
  () => import("@/components/flutter-app-development/BlogSection.vue"),
);
const CtaSection2 = defineAsyncComponent(
  () => import("@/components/flutter-app-development/CtaSection2.vue"),
);
const FaqSection = defineAsyncComponent(
  () => import("@/components/flutter-app-development/FaqSection.vue"),
);
const NewFooter = defineAsyncComponent(
  () => import("@/components/partials/NewFooter.vue"),
);

export default {
  beforeRouteEnter(to, from, next) {
    if (!config.SHOW_FLUTTER_APP_DEVELOPMENT_PAGE) {
      next({
        name: "Error404Page",
        params: { pathMatch: "/flutter-app-development" },
      });
    } else {
      next();
    }
  },
  components: {
    Header,
    LandingSection,
    DevelopmentSection,
    CaseStudySection,
    CtaSection,
    SuccessStory,
    BlogSection,
    CtaSection2,
    FaqSection,
    NewFooter,
  },
  setup() {
    const { $mixpanel } = useNuxtApp();
    const footer = ref(null);
    const seoData = config.FLUTTER_APP_DEVELOPMENT_SEO_META_DATA;
    useSeoMeta({
      robots: "noindex, nofollow",
      title: seoData.title,
      description: seoData.description,
      ogTitle: seoData.title,
      ogType: seoData.type,
      ogUrl: seoData.url,
      ogImage: seoData.image,
    });

    let event = "";
    let events = {
      footer: "view_flutter_app_development_footer",
    };
    let elements;
    onMounted(() => {
      elements = ref({
        footer: footer,
      });
      window.addEventListener("scroll", sendEvent);
      $mixpanel.track("view_flutter_app_development_page");
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
