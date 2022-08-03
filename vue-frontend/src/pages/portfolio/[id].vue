<template>
  <div>
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
import luxeradioResponse from "@/luxeradio-data.js";
import nolonelyResponse from "@/nolonely-data.js";
import tognessResponse from "@/togness-data.js";

export default {
  data() {
    return {
      id: this.$route.params.id,
      details: "",
    };
  },
  created() {
    this.getData();
  },
  methods: {
    getData() {
      switch (this.id) {
        case luxeradioResponse.name:
          this.details = luxeradioResponse.detail;
          break;

        case nolonelyResponse.name:
          this.details = nolonelyResponse.detail;
          break;

        case tognessResponse.name:
          this.details = tognessResponse.detail;
          break;
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

@media (min-width: 5000px) {
  .container-fluid {
    width: 2300px;
  }
}
</style>
