<template>
  <div v-if="response">
    <metainfo>
      <template v-slot:title="{ content }">
        {{ content }}
      </template>
    </metainfo>
    <Header />
    <div
      class="parallax tw-w-full tw-p-0 tw-m-auto 4xl:tw-w-[2300px]"
      ref="response"
    >
      <LandingSection v-bind:json="details.landing" />
      <VideoSection v-bind:json="details.video" />
      <BrandingSection v-bind:json="details.branding" />
      <DesignSection v-bind:json="details.design" />
      <ElementSection v-bind:json="details.element" />
      <FooterSection v-bind:json="details.footer" />
      <CTASection :ref="ctaRef" />
    </div>
  </div>
  <div v-else><ErrorPage /></div>
</template>

<script>
import Header from "@/components/partials/NewHeader.vue";
import LandingSection from "@/components/portfolio/LandingSection.vue";
import VideoSection from "@/components/portfolio/VideoSection.vue";
import BrandingSection from "@/components/portfolio/BrandingSection.vue";
import DesignSection from "@/components/portfolio/DesignSection.vue";
import FooterSection from "@/components/portfolio/FooterSection.vue";
import CTASection from "@/components/home/CTASection.vue";
import ElementSection from "@/components/portfolio/ElementSection.vue";
import luxeradioResponse from "@/portfolio-json/luxeradio-data.js";
import justlyResponse from "@/portfolio-json/justly-data.js";
import tognessResponse from "@/portfolio-json/togness-data.js";
import ErrorPage from "@/components/error404/index.vue";
import { useMeta } from "vue-meta";
import { elementInViewPort } from "@/utils.js";

export default {
  setup() {
    const { meta } = useMeta({
      og: {
        type: "Website",
      },
    });
    return {
      meta,
    };
  },
  data() {
    return {
      details: "",
      response: null,
      ctaRef: null,
      event: "",
      events: {
        luxepagecta: "view_luxe_page_cta",
        tognesspagecta: "view_togness_page_cta",
        justlypagecta: "view_justly_page_cta",
      },
    };
  },
  watch: {
    "$route.params.id": {
      handler: function () {
        this.getData(this.$route.params.id);
      },
      deep: true,
      immediate: true,
    },
  },
  created() {
    this.getData(this.$route.params.id);
  },
  inject: ["mixpanel"],
  mounted() {
    window.addEventListener("scroll", this.sendEvent);
    if (this.mixpanel.__loaded) {
      this.mixpanel.track("view_page_portfolio");
    }
  },
  unmounted() {
    window.removeEventListener("scroll", this.sendEvent);
  },
  updated() {
    this.scrollToTop();
  },
  methods: {
    getData(id) {
      switch (id) {
        case luxeradioResponse.name:
          this.response = luxeradioResponse;
          this.ctaRef = "luxepagecta";
          break;

        case justlyResponse.name:
          this.response = justlyResponse;
          this.ctaRef = "justlypagecta";
          break;

        case tognessResponse.name:
          this.response = tognessResponse;
          this.ctaRef = "tognesspagecta";
          break;
      }

      if (this.response) {
        this.details = this.response.detail;
        this.meta.title = this.response.seoData.title;
        this.meta.description = this.response.seoData.description;
        this.meta.og.title = this.response.seoData.title;
        this.meta.og.url = this.response.seoData.url;
      }
    },
    sendEvent() {
      const event = this.events[elementInViewPort(this.$refs)];
      if (this.mixpanel.__loaded && event && this.event !== event) {
        this.event = event;
        this.mixpanel.track(event);
      }
    },
    scrollToTop() {
      var portfolioDiv = this.$refs.response;
      portfolioDiv.scrollTo({
        top: portfolioDiv.offsetTop,
        behavior: "instant",
      });
    },
  },
  components: {
    Header,
    LandingSection,
    VideoSection,
    BrandingSection,
    DesignSection,
    FooterSection,
    CTASection,
    ElementSection,
    ErrorPage,
  },
};
</script>

<style lang="postcss" scoped>
@screen lg {
  .parallax {
    perspective: 2px;

    @apply tw-h-screen tw-overflow-hidden tw-overflow-y-auto;
  }
}
</style>
