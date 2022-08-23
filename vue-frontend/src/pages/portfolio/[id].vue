<template>
  <div v-if="response">
    <metainfo>
      <template v-slot:title="{ content }">
        {{ content }}
      </template>
    </metainfo>
    <Header />
    <div class="parallax container-fluid tw-p-0">
      <LandingSection v-bind:response="details.landing" />
      <VideoSection v-bind:response="details.video" />
      <BrandingSection v-bind:response="details.branding" />
      <DesignSection v-bind:response="details.design" />
      <ElementSection v-bind:response="details.element" />
      <FooterSection v-bind:response="details.footer" />
      <CTASection />
    </div>
  </div>
</template>

<script>
import Header from "@/components/partials/ScreenHeaderV2.vue";
import LandingSection from "@/components/portfolio/LandingSection.vue";
import VideoSection from "@/components/portfolio/VideoSection.vue";
import BrandingSection from "@/components/portfolio/BrandingSection.vue";
import DesignSection from "@/components/portfolio/DesignSection.vue";
import FooterSection from "@/components/portfolio/FooterSection.vue";
import CTASection from "@/components/home-new/CTASection.vue";
import ElementSection from "@/components/portfolio/ElementSection.vue";
import luxeradioResponse from "@/portfolio-json/luxeradio-data.js";
import nolonelyResponse from "@/portfolio-json/nolonely-data.js";
import tognessResponse from "@/portfolio-json/togness-data.js";
import smileplusResponse from "@/portfolio-json/smileplus-data.js";
import { useMeta } from "vue-meta";

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
      id: this.$route.params.id,
      details: "",
      response: null,
    };
  },
  created() {
    this.getData();
  },
  mounted() {
    this.$gtag.event("view_page_portfolio");
  },
  methods: {
    getData() {
      switch (this.id) {
        case luxeradioResponse.name:
          this.response = luxeradioResponse;
          break;

        case nolonelyResponse.name:
          this.response = nolonelyResponse;
          break;

        case tognessResponse.name:
          this.response = tognessResponse;
          break;

        case smileplusResponse.name:
          this.response = smileplusResponse;
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
  },
};
</script>

<style lang="scss" scoped>
@include media-breakpoint-up(lg) {
  .parallax {
    height: 100vh;
    overflow: hidden;
    overflow-y: auto;
    perspective: 2px;
  }
  .parallax::-webkit-scrollbar {
    display: none; /* Chrome, Safari and Opera */
  }
  .parallax {
    -ms-overflow-style: none; /* IE and Edge */
    scrollbar-width: none; /* Firefox */
  }
}

@media (min-width: 5000px) {
  .container-fluid {
    width: 2300px;
  }
}
</style>
