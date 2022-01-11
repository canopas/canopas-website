<template>
  <div class="main-page-wrapper">
    <!-- Main Website -->
    <div class="main-wrapper">
      <Header />
      <HomeLandingPage />
      <HomeClientReviewPage />
      <HomeUserReviewPage />
    </div>
  </div>
</template>

<script>
import Header from "./partials/Header.vue";
import HomeLandingPage from "./home/HomeLandingPage.vue";
import HomeUserReviewPage from "./home/HomeUserReviewPage.vue";
import HomeClientReviewPage from "./home/HomeClientReviewPage.vue";

export default {
  name: "App",
  components: {
    Header,
    HomeLandingPage,
    HomeUserReviewPage,
    HomeClientReviewPage,
  },
};
</script>

<style>
.normal-text {
  font-style: normal;
  font-weight: 400;
  font-size: 1.1rem;
  line-height: 1.3rem;
  color: rgba(0, 0, 0, 0.5);
}

.header-text {
  font-weight: 700;
  line-height: 1.813rem;
  font-size: 1.5rem;
}
</style>
