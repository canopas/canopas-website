<template>
  <div>
    <metainfo>
      <template v-slot:title="{ content }">
        {{ content }}
      </template>
    </metainfo>
    <Header />
    <LandingSection ref="iosLanding" />
    <DevelopmentSection id="developmentSection" ref="iosDevelopment" />
    <CaseStudySection
      ref="iosCaseStudy"
      v-on:scroll-to-next="scrollToCta"
      v-on:scroll-to-previous="scrollToDevelopment"
    />
    <CtaSection id="ctasection" ref="iosCta1" />
    <SuccessStorySection ref="iosSuccessstory" />
    <BlogSection ref="iosblog" />
    <CtaSection2 ref="iosCta2" />
    <FaqSection ref="iosFaq" />
    <NewFooter ref="iosFooter" />
  </div>
</template>

<script>
import Header from "@/components/partials/NewHeader.vue";
import config from "@/config.js";
import { useMeta } from "vue-meta";
import { defineAsyncComponent } from "vue";
import { elementInViewPort } from "@/utils.js";
import LandingSection from "@/components/ios-app-development/LandingSection.vue";
import DevelopmentSection from "@/components/ios-app-development/DevelopmentSection.vue";
const SuccessStorySection = defineAsyncComponent(() =>
  import("@/components/ios-app-development/SuccessStory.vue"),
);
const BlogSection = defineAsyncComponent(() =>
  import("@/components/ios-app-development/BlogSection.vue"),
);
const FaqSection = defineAsyncComponent(() =>
  import("@/components/ios-app-development/FaqSection.vue"),
);
const CaseStudySection = defineAsyncComponent(() =>
  import("@/components/ios-app-development/CaseStudySection.vue"),
);
const CtaSection = defineAsyncComponent(() =>
  import("@/components/ios-app-development/CtaSection.vue"),
);
const CtaSection2 = defineAsyncComponent(() =>
  import("@/components/ios-app-development/CtaSection2.vue"),
);
const NewFooter = defineAsyncComponent(() =>
  import("@/components/partials/NewFooter.vue"),
);

export default {
  data() {
    return {
      event: "",
      events: {
        iosLanding: "view_ios_development_landing_section",
        iosDevelopment: "view_ios_app_development_section",
        iosSuccessstory: "view_ios_app_development_section",
        iosFaq: "view_ios_development_faq_section",
        iosCaseStudy: "view_ios_CaseStudy_section",
        iosCta1: "view_cta_1_ios_app_development_section",
        iosblog: "view_ios_app_blog_section",
        iosCta2: "view_cta_2_ios_app_development_section",
      },
    };
  },
  setup() {
    var seoData = config.IOS_APP_DEVELOPMENT_SEO_META_DATA;
    useMeta({
      meta: [
        {
          name: "robots",
          content: "noindex, nofollow",
          vmid: "robots",
        },
      ],
      title: seoData.title,
      description: seoData.description,
      og: {
        type: seoData.type,
        title: seoData.title,
        url: seoData.url,
        image: seoData.image,
      },
    });
  },
  components: {
    Header,
    LandingSection,
    DevelopmentSection,
    SuccessStorySection,
    BlogSection,
    FaqSection,
    CaseStudySection,
    CtaSection,
    CtaSection2,
    NewFooter,
  },
  methods: {
    sendEvent() {
      const event = this.events[elementInViewPort(this.$refs)];
      if (event && this.event !== event) {
        this.event = event;
        this.mixpanel.track(event);
      }
    },
    scrollToCta() {
      var ctaDiv = document.getElementById("ctasection");
      var top = ctaDiv.offsetTop;
      window.scrollTo({ top: top, behavior: "smooth" });
    },
    scrollToDevelopment() {
      var developmentDiv = document.getElementById("developmentSection");
      var top = developmentDiv.offsetTop;
      window.scrollTo({ top: top, behavior: "smooth" });
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
    this.mixpanel.track("view_ios_development_page");
  },
  unmounted() {
    window.removeEventListener("scroll", this.sendEvent);
  },
};
</script>
