<template>
  <div>
    <metainfo>
      <template v-slot:title="{ content }">
        {{ content }}
      </template>
    </metainfo>
    <Header />

    <div class="tw-overflow-hidden">
      <div class="tw-bg-[#262626] md:tw-bg-[#1A1A1A]">
        <LandingSection ref="landing" />
        <Services ref="services" />
      </div>
      <Portfolio ref="portfolio" />
      <CTASection ref="cta1" />
      <DevelopmentProcess ref="developmentprocess" />
      <div class="tw-bg-black-core/[0.85]">
        <ClientReview ref="clientreview" />
        <CTASection2 ref="cta2" />
      </div>
    </div>

    <NewFooter class="tw-mt-[-15px] md:tw-mt-auto" />
  </div>
</template>

<script>
import Header from "@/components/partials/NewHeader.vue";
import LandingSection from "@/components/mobile-app-development/LandingSection.vue";
import Services from "@/components/mobile-app-development/Services.vue";
import config from "@/config.js";
import { useMeta } from "vue-meta";
import { defineAsyncComponent } from "vue";
import { elementInViewPort } from "@/utils.js";

const Portfolio = defineAsyncComponent(() =>
  import("@/components/mobile-app-development/Portfolio.vue")
);
const CTASection = defineAsyncComponent(() =>
  import("@/components/mobile-app-development/CTASection.vue")
);
const DevelopmentProcess = defineAsyncComponent(() =>
  import("@/components/mobile-app-development/DevelopmentProcess.vue")
);

const ClientReview = defineAsyncComponent(() =>
  import("@/components/mobile-app-development/ClientReview.vue")
);
const CTASection2 = defineAsyncComponent(() =>
  import("@/components/mobile-app-development/CTASection2.vue")
);
const NewFooter = defineAsyncComponent(() =>
  import("@/components/partials/NewFooter.vue")
);

export default {
  setup() {
    var seoData = config.MOBILE_APP_DEVELOPMENT_SEO_META_DATA;
    useMeta({
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
        landing: "view_mobileapp_development_landing_section",
        services: "view_mobileapp_development_services_section",
        portfolio: "view_mobileapp_development_portfolio_section",
        cta1: "view_mobileapp_development_cta_section",
        developmentprocess:
          "view_mobileapp_development_developmentprocess_section",
        clientreview: "view_mobileapp_development_clientreview_section",
        cta2: "view_mobileapp_development_cta2_section",
        footer: "view_mobileapp_development_footer",
      },
    };
  },
  inject: ["mixpanel"],
  mounted() {
    window.addEventListener("scroll", this.sendEvent);
    this.mixpanel.track("view_mobileapp_landing_section");
  },
  unmounted() {
    window.removeEventListener("scroll", this.sendEvent);
  },
  methods: {
    sendEvent() {
      const event = this.events[elementInViewPort(this.$refs)];
      if (event && this.event !== event) {
        this.event = event;
        this.mixpanel.track(event);
      }
    },
  },
};
</script>
