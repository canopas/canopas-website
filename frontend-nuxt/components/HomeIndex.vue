<template>
  <div>
    <ScreenHeaderV2 />
    <div class="landing-background">
      <HomeLandingSection />
      <HomeCanopasDescription />
    </div>
    <HomePortfolioSection />
    <HomeProblemSolution />
    <HomeWithCanopas />
    <HomeClientReview />
    <HomeUserReview />
    <HomeBlogSection />
    <HomeContributeSection />
    <HomeCTASection />
    <ScreenFooter2 />
  </div>
</template>

<script>
import { library } from "@fortawesome/fontawesome-svg-core";
import { faPhone, faPaperPlane } from "@fortawesome/free-solid-svg-icons";

library.add(faPhone, faPaperPlane);

export default {
};
</script>

<style scoped>
.landing-background {
  background: linear-gradient(
    174.39deg,
    #ffffff 9.06%,
    rgba(255, 148, 114, 0.15) 90%
  );
  z-index: -5;
}
</style>
