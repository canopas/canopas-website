<template>
  <div>
    <Header />
    <LandingSection />
    <DevelopmentSection />
    <CtaSection />
    <SuccessStorySection ref="iosSuccessstory" />
    <BlogSection />
    <CtaSection2 />
    <FaqSection />
    <NewFooter ref="iosFooter" />
  </div>
</template>

<script>
import Header from "@/components/partials/NewHeader.vue";
import config from "@/config.js";
import { defineAsyncComponent } from "vue";
import { elementInViewPort } from "@/utils.js";
import LandingSection from "@/components/ios-app-development/LandingSection.vue";
import DevelopmentSection from "@/components/ios-app-development/DevelopmentSection.vue";
const SuccessStorySection = defineAsyncComponent(
  () => import("@/components/ios-app-development/SuccessStory.vue"),
);
const BlogSection = defineAsyncComponent(
  () => import("@/components/ios-app-development/BlogSection.vue"),
);
const FaqSection = defineAsyncComponent(
  () => import("@/components/ios-app-development/FaqSection.vue"),
);
const CtaSection = defineAsyncComponent(
  () => import("@/components/ios-app-development/CtaSection.vue"),
);
const CtaSection2 = defineAsyncComponent(
  () => import("@/components/ios-app-development/CtaSection2.vue"),
);
const NewFooter = defineAsyncComponent(
  () => import("@/components/partials/NewFooter.vue"),
);

export default {
  data() {
    return {
      event: "",
      events: {
        iosSuccessstory: "view_ios_app_development_success_story",
        iosFooter: "view_ios_app_development_footer",
      },
    };
  },
  setup() {
    const seoData = config.IOS_APP_DEVELOPMENT_SEO_META_DATA;
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
    DevelopmentSection,
    SuccessStorySection,
    BlogSection,
    FaqSection,
    CtaSection,
    CtaSection2,
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
    if (!config.SHOW_IOS_APP_DEVELOPMENT_PAGE) {
      next({
        name: "Error404Page",
        params: { pathMatch: "/ios-app-development" },
      });
    } else {
      next();
    }
  },
  inject: ["mixpanel"],
  mounted() {
    window.addEventListener("scroll", this.sendEvent);
    this.$mixpanel.track("view_ios_development_page");
  },
  unmounted() {
    window.removeEventListener("scroll", this.sendEvent);
  },
};
</script>
