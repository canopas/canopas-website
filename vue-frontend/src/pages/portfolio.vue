<template>
  <div>
    <ScreenHeader />
    <LandingSection />
    <VideoSection />
    <BrandingSection />
    <DesignSection />
    <ElementSection />
    <CTA />
  </div>
</template>

<script>
import ScreenHeader from "@/components/partials/ScreenHeaderV2.vue";
import LandingSection from "@/components/portfolio/LandingSection.vue";
import VideoSection from "@/components/portfolio/VideoSection.vue";
import BrandingSection from "@/components/portfolio/BrandingSection.vue";
import DesignSection from "@/components/portfolio/DesignSection.vue";
import ElementSection from "@/components/portfolio/ElementSection.vue";
import CTA from "@/components/home-new/CTASection.vue";

export default {
  components: {
    ScreenHeader,
    LandingSection,
    VideoSection,
    BrandingSection,
    DesignSection,
    ElementSection,
    CTA,
  },
};
</script>

<style scoped></style>
