<template>
  <div class="main-page-wrapper">
    <!-- Main Website -->
    <div class="main-wrapper">
      <Header />
      <HomeLandingPage />
      <HomeClientReviewPage />
      <ProblemSolutionPage />
    </div>
  </div>
</template>

<script>
import Header from "./partials/Header.vue";
import HomeLandingPage from "./home/HomeLandingPage.vue";
import HomeClientReviewPage from "./home/HomeClientReviewPage.vue";
import ProblemSolutionPage from "./home/ProblemSolution.vue";

export default {
  name: "App",
  components: {
    Header,
    HomeLandingPage,
    HomeClientReviewPage,
    ProblemSolutionPage,
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
