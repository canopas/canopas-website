<template>
  <div>
    <Header />
    <LandingSection />
    <CtaSection />
    <SuccessStory />
    <NewFooter ref="flutterFooter" />
  </div>
</template>

<script>
import Header from "@/components/partials/NewHeader.vue";
import config from "@/config.js";
import { defineAsyncComponent } from "vue";
import { elementInViewPort } from "@/utils.js";
import LandingSection from "@/components/flutter-app-development/LandingSection.vue";
const CtaSection = defineAsyncComponent(() =>
  import("@/components/flutter-app-development/CtaSection.vue"),
);
const SuccessStory = defineAsyncComponent(() =>
  import("@/components/flutter-app-development/SuccessStorySection.vue"),
);
const NewFooter = defineAsyncComponent(() =>
  import("@/components/partials/NewFooter.vue"),
);
export default {
  data() {
    return {
      event: "",
      events: {
        flutterFooter: "view_flutter_app_development_footer",
      },
    };
  },
  setup() {
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
  },
  components: {
    Header,
    LandingSection,
    CtaSection,
    SuccessStory,
    NewFooter,
  },
  methods: {
    sendEvent() {
      const event = this.events[elementInViewPort(this.$refs)];
      if (event && this.event !== event) {
        this.event = event;
        this.$mixpanel.track(event);
      }
    },
  },
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
  inject: ["mixpanel"],
  mounted() {
    window.addEventListener("scroll", this.sendEvent);
    this.$mixpanel.track("view_flutter_development_page");
  },
  unmounted() {
    window.removeEventListener("scroll", this.sendEvent);
  },
};
</script>
