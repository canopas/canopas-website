<template>
  <div>
    <Header />
    <div class="overflow-hidden">
      <div class="bg-[#262626] md:bg-[#1A1A1A]">
        <LandingSection />
        <Services />
      </div>
      <Portfolio />
      <CTASection />
      <DevelopmentProcess ref="developmentprocess" />
      <div class="bg-black-core/[0.85]">
        <ClientReview />
        <CTASection2 />
      </div>
    </div>
    <NewFooter ref="footer" class="mt-[-15px] md:mt-auto" />
  </div>
</template>

<script>
import Header from "@/components/partials/NewHeader.vue";
import LandingSection from "@/components/mobile-app-development/LandingSection.vue";
import Services from "@/components/mobile-app-development/Services.vue";
import config from "@/config.js";
import { defineAsyncComponent } from "vue";
import { elementInViewPort } from "@/utils.js";

const Portfolio = defineAsyncComponent(() =>
  import("@/components/mobile-app-development/Portfolio.vue"),
);
const CTASection = defineAsyncComponent(() =>
  import("@/components/mobile-app-development/CTASection.vue"),
);
const DevelopmentProcess = defineAsyncComponent(() =>
  import("@/components/mobile-app-development/DevelopmentProcess.vue"),
);
const ClientReview = defineAsyncComponent(() =>
  import("@/components/mobile-app-development/ClientReview.vue"),
);
const CTASection2 = defineAsyncComponent(() =>
  import("@/components/mobile-app-development/CTASection2.vue"),
);
const NewFooter = defineAsyncComponent(() =>
  import("@/components/partials/NewFooter.vue"),
);

export default {
  setup() {
    const seoData = config.MOBILE_APP_DEVELOPMENT_SEO_META_DATA;
    useSeoMeta({
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
    Services,
    Portfolio,
    CTASection,
    DevelopmentProcess,
    ClientReview,
    CTASection2,
    NewFooter,
  },
  data() {
    return {
      event: "",
      events: {
        developmentprocess:
          "view_mobileapp_development_developmentprocess_section",
        footer: "view_mobileapp_development_footer",
      },
    };
  },
  inject: ["mixpanel"],
  mounted() {
    window.addEventListener("scroll", this.sendEvent);
    this.$mixpanel.track("view_mobileapp_development_page");
  },
  unmounted() {
    window.removeEventListener("scroll", this.sendEvent);
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
};
</script>
