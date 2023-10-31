<template>
  <div>
    <Header />
    <div>
      <LandingSection />
      <HowItAllStartedSectionMobile class="block lg:hidden" />
      <HowItAllStartedSection class="hidden lg:block" />
      <AboutusVirtue />
      <WithCanopasSection ref="withcanopas" />
      <ClientReviewSection />
      <CTASection />
    </div>
    <NewFooter ref="footer" />
  </div>
</template>

<script>
import Header from "@/components/partials/NewHeader.vue";
import LandingSection from "@/components/about/LandingSection.vue";
import HowItAllStartedSection from "@/components/about/HowItAllStartedSection.vue";
import config from "@/config.js";
import { elementInViewPort } from "@/utils.js";
import { defineAsyncComponent } from "vue";

const HowItAllStartedSectionMobile = defineAsyncComponent(() =>
  import("@/components/about/HowItAllStartedSectionMobile.vue"),
);
const AboutusVirtue = defineAsyncComponent(() =>
  import("@/components/about/AboutusVirtue.vue"),
);
const WithCanopasSection = defineAsyncComponent(() =>
  import("@/components/about/WithCanopas.vue"),
);
const ClientReviewSection = defineAsyncComponent(() =>
  import("@/components/home-new/ClientReviewSection.vue"),
);
const CTASection = defineAsyncComponent(() =>
  import("@/components/about/CTASection.vue"),
);
const NewFooter = defineAsyncComponent(() =>
  import("@/components/partials/NewFooter.vue"),
);

export default {
  setup() {
    const seoData = config.ABOUT_SEO_META_DATA;
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
        withcanopas: "view_about_phases",
        footer: "view_about_footer",
      },
    };
  },
  inject: ["mixpanel"],
  mounted() {
    window.addEventListener("scroll", this.sendEvent);
    this.$mixpanel.track("view_about_page");
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
