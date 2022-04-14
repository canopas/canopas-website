<template>
  <div>
    <ScreenHeader />
    <Portfolio />
    <ClientReview />
    <UserReview />
  </div>
</template>

<script>
import ScreenHeader from "@/components/partials/HeaderV2.vue";
import Portfolio from "@/components/home-new/PortfolioSection.vue";
import ClientReview from "@/components/home-new/ClientReview.vue";
import UserReview from "@/components/home-new/UserReview.vue";

export default {
  components: {
    ScreenHeader,
    Portfolio,
    ClientReview,
    UserReview,
  },
};
</script>

<style scoped></style>
