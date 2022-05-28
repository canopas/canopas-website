<template>
  <div>
    <PortfolioHeader />
    <LandingSection />
    <VideoSection />
    <BrandingSection />
    <DesignSection />
    <ElementSection />
    <CTA />
  </div>
</template>

<script>
import PortfolioHeader from "@/components/partials/PortfolioHeader.vue";
import LandingSection from "@/components/portfolio/LandingSection.vue";
import VideoSection from "@/components/portfolio/VideoSection.vue";
import BrandingSection from "@/components/portfolio/BrandingSection.vue";
import DesignSection from "@/components/portfolio/DesignSection.vue";
import ElementSection from "@/components/portfolio/ElementSection.vue";
import CTA from "@/components/home-new/CTASection.vue";

export default {
  components: {
    PortfolioHeader,
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
