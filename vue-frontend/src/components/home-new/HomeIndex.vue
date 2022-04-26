<template>
  <div>
    <ScreenHeader />
    <Portfolio />
    <ClientReview />
    <UserReview />
    <Blog />
    <CTA />
  </div>
</template>

<script>
import ScreenHeader from "@/components/partials/ScreenHeader.vue";
import Portfolio from "@/components/home-new/PortfolioSection.vue";
import ClientReview from "@/components/home-new/ClientReview.vue";
import UserReview from "@/components/home-new/UserReview.vue";
import CTA from "@/components/home-new/CTASection.vue";
import Blog from "@/components/home-new/BlogSection.vue";

export default {
  components: {
    ScreenHeader,
    Portfolio,
    ClientReview,
    UserReview,
    CTA,
    Blog,
  },
};
</script>

<style scoped></style>
