<template>
  <div>
    <ScreenHeader />
    <Portfolio />
  </div>
</template>

<script>
import ScreenHeader from "@/components/partials/ScreenHeader.vue";
import Portfolio from "@/components/home-new/PortfolioSection.vue";

export default {
  components: {
    ScreenHeader,
    Portfolio,
  },
};
</script>

<style scoped></style>
