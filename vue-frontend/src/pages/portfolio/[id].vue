<template>
  <div>
    <Header />
    <div class="parallax">
      <LandingSection v-bind:response="details.landing" />
      <VideoSection v-bind:response="details.video" />
      <BrandingSection v-bind:response="details.branding" />
      <DesignSection v-bind:response="details.design" />
      <ElementSection v-bind:response="details.element" />
    </div>
  </div>
</template>

<script>
import Header from "@/components/partials/ScreenHeaderV2.vue";
import LandingSection from "@/components/portfolio/LandingSection.vue";
import VideoSection from "@/components/portfolio/VideoSection.vue";
import BrandingSection from "@/components/portfolio/BrandingSection.vue";
import DesignSection from "@/components/portfolio/DesignSection.vue";
import ElementSection from "@/components/portfolio/ElementSection.vue";
import luxeradioResponse from "@/luxeradio-data.js";

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
      if (this.id == luxeradioResponse.name) {
        this.details = luxeradioResponse.detail;
      }
    },
  },

  components: {
    Header,
    LandingSection,
    VideoSection,
    BrandingSection,
    DesignSection,
    ElementSection,
  },
};
</script>

<style lang="postcss" scoped>
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
</style>
