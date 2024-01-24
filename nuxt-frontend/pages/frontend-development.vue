<template>
  <div>
    <Header />
    <LandingSection />
    <NewFooter class="bg-gradient-to-b from-pink-0 to-pink-16" />
  </div>
</template>
<script>
import Header from "@/components/partials/NewHeader.vue";
import NewFooter from "@/components/partials/NewFooter.vue";
import config from "@/config.js";
import { elementInViewPort } from "@/utils.js";
import LandingSection from "@/components/frontend-development/LandingSection.vue";
export default {
  beforeRouteEnter(to, from, next) {
    if (!config.SHOW_FRONTEND_DEVELOPMENT_PAGE) {
      next({
        name: "Error404Page",
        params: { pathMatch: "/frontend-development" },
      });
    } else {
      next();
    }
  },
  components: {
    Header,
    LandingSection,
    NewFooter,
  },
  setup() {
    const { $mixpanel } = useNuxtApp();
    const footer = ref(null);
    const seoData = config.FRONTEND_DEVELOPMENT_SEO_META_DATA;
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
      footer: "view_frontend_development_footer",
    };
    let elements;
    onMounted(() => {
      elements = ref({
        footer: footer,
      });
      window.addEventListener("scroll", sendEvent);
      $mixpanel.track("view_frontend_development_page");
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
