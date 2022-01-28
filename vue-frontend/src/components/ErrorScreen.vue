<template>
  <div ref="contactScreen">
    <ScreenHeader />
    <NotFound />
    <ScreenFooter />
  </div>
</template>

<script>
import ScreenHeader from "./partials/ScreenHeader.vue";
import ScreenFooter from "./partials/ScreenFooter.vue";
import NotFound from "./error/NotFoundPage.vue";
export default {
  components: {
    ScreenHeader,
    NotFound,
    ScreenFooter,
  },
  mounted() {},
};
</script>

<style scoped></style>
