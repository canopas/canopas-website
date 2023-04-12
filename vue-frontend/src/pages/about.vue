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
import NewFooter from "@/components/partials/NewFooter.vue";
import LandingSection from "@/components/about/LandingSection.vue";
import HowItAllStartedSectionMobile from "@/components/about/HowItAllStartedSectionMobile.vue";
import HowItAllStartedSection from "@/components/about/HowItAllStartedSection.vue";
import AboutusVirtue from "@/components/about/AboutusVirtue.vue";
import WithCanopasSection from "@/components/home/WithCanopas.vue";
import ClientReviewSection from "@/components/home-new/ClientReviewSection.vue";
import CTASection from "@/components/about/CTASection.vue";
import config from "@/config.js";
import { useMeta } from "vue-meta";
import { elementInViewPort } from "@/utils.js";

export default {
  setup() {
    var seoData = config.ABOUT_SEO_META_DATA;
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
  inject: ["mixpanel"],
  beforeRouteEnter(to, from, next) {
    if (config.SHOW_ABOUT_US_PAGE === false) {
      next({ name: "Error404Page", params: { pathMatch: "/about" } });
    } else {
      next();
    }
  },
  mounted() {
    window.addEventListener("scroll", this.sendEvent);
    this.mixpanel.track("view_about_page");
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
