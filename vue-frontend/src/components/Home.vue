<template>
  <div class="main-page-wrapper">
    <!-- Main Website -->
    <div class="main-wrapper">
      <Header />
      <LandingView />
      <ClientReview />
      <UserReviews />
      <DoYouKnow />
    </div>
  </div>
</template>

<script>
import Header from "./partials/Header.vue";
import LandingView from "./home/LandingView.vue";
import UserReviews from "./home/UserReviews.vue";
import ClientReview from "./home/ClientReview.vue";
import DoYouKnow from "./home/DoYouKnow.vue";

export default {
  name: "App",
  components: {
    Header,
    LandingView,
    UserReviews,
    ClientReview,
    DoYouKnow,
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
