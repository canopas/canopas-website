<template>
  <div>
    <ScreenHeader />
    <Portfolio />
  </div>
</template>

<script>
import ScreenHeader from "./partials/ScreenHeader.vue";
import Portfolio from "./home-new/PortfolioSection.vue";

export default {
  data() {
    return {};
  },
  components: {
    ScreenHeader,
    Portfolio,
  },
};
</script>

<style scoped></style>
