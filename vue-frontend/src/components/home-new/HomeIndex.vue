<template>
  <div>
    <ScreenHeader />
    <Portfolio />
    <ClientReview />
  </div>
</template>

<script>
import ScreenHeader from "@/components/partials/HeaderV2.vue";
import Portfolio from "@/components/home-new/PortfolioSection.vue";
import ClientReview from "@/components/home-new/ClientReview.vue";
export default {
  components: {
    ScreenHeader,
    Portfolio,
    ClientReview,
  },
};
</script>

<style scoped></style>
