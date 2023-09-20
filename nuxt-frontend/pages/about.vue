<template>
  <div>
    <metainfo>
      <template v-slot:title="{ content }">
        {{ content }}
      </template>
    </metainfo>
    <Header />
    <div>
      <LandingSection ref="landing" />
      <HowItAllStartedSectionMobile
        ref="howstarted"
        class="tw-block lg:tw-hidden"
      />
      <HowItAllStartedSection ref="howstarted" class="tw-hidden lg:tw-block" />
      <AboutusVirtue ref="aboutvirtue" />
      <WithCanopasSection ref="withcanopas" />
      <ClientReviewSection ref="aboutclientreview" />
      <CTASection ref="aboutcta" />
    </div>
    <NewFooter />
  </div>
</template>

<script>
import Header from "@/components/partials/NewHeader.vue";
import LandingSection from "@/components/about/LandingSection.vue";
import HowItAllStartedSection from "@/components/about/HowItAllStartedSection.vue";
import config from "@/config.js";
// import { useMeta } from "vue-meta";
import { elementInViewPort } from "@/utils.js";
import { defineAsyncComponent } from "vue";

const HowItAllStartedSectionMobile = defineAsyncComponent(() =>
  import("@/components/about/HowItAllStartedSectionMobile.vue")
);
const AboutusVirtue = defineAsyncComponent(() =>
  import("@/components/about/AboutusVirtue.vue")
);
const WithCanopasSection = defineAsyncComponent(() =>
  import("@/components/about/WithCanopas.vue")
);
const ClientReviewSection = defineAsyncComponent(() =>
  import("@/components/home-new/ClientReviewSection.vue")
);
const CTASection = defineAsyncComponent(() =>
  import("@/components/about/CTASection.vue")
);

const NewFooter = defineAsyncComponent(() =>
  import("@/components/partials/NewFooter.vue")
);

export default {
  setup() {
    // var seoData = config.ABOUT_SEO_META_DATA;
    // useMeta({
    //   title: seoData.title,
    //   description: seoData.description,
    //   og: {
    //     type: seoData.type,
    //     title: seoData.title,
    //     url: seoData.url,
    //     image: seoData.image,
    //   },
    // });
  },
  components: {
    Header,
    LandingSection,
    HowItAllStartedSectionMobile,
    HowItAllStartedSection,
    AboutusVirtue,
    WithCanopasSection,
    ClientReviewSection,
    CTASection,
    NewFooter,
  },
  data() {
    return {
      event: "",
      events: {
        landing: "view_about_landing",
        howstarted: "view_about_how_all_started",
        aboutvirtue: "view_about_virtues",
        withcanopas: "view_about_phases",
        aboutclientreview: "view_about_client_review",
        aboutcta: "view_about_cta",
      },
    };
  },
  // inject: ["mixpanel"],
  mounted() {
    window.addEventListener("scroll", this.sendEvent);
    // // this.mixPanel.track("view_about_page");
  },
  unmounted() {
    window.removeEventListener("scroll", this.sendEvent);
  },
  methods: {
    sendEvent() {
      const event = this.events[elementInViewPort(this.$refs)];
      if (event && this.event !== event) {
        this.event = event;
        // // this.mixPanel.track(event);
      }
    },
  },
};
</script>
