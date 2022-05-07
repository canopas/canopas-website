<template>
  <div>
    <ScreenHeader />
    <LandingSection />
    <Portfolio />
    <WithCanopas />
    <ClientReview />
    <UserReview />
    <Blog />
    <Contribute />
    <CTA />
    <ScreenFooter2 />
  </div>
</template>

<script>
import ScreenHeader from "@/components/partials/ScreenHeaderV2.vue";
import LandingSection from "@/components/home-new/LandingSection.vue";
import Portfolio from "@/components/home-new/PortfolioSection.vue";
import ClientReview from "@/components/home-new/ClientReview.vue";
import WithCanopas from "@/components/home-new/WithCanopas.vue";
import UserReview from "@/components/home-new/UserReview.vue";
import CTA from "@/components/home-new/CTASection.vue";
import Blog from "@/components/home-new/BlogSection.vue";
import Contribute from "@/components/home-new/ContributeSection.vue";
import ScreenFooter2 from "@/components/partials/ScreenFooter2.vue";

export default {
  components: {
    ScreenHeader,
    LandingSection,
    Portfolio,
    WithCanopas,
    ClientReview,
    UserReview,
    CTA,
    Blog,
    Contribute,
    ScreenFooter2,
  },
};
</script>

<style scoped></style>
