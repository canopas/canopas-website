<template>
  <div>
    <ScreenHeader />
    <div class="landing-background">
      <LandingSection />
      <CanopasDescription />
    </div>
    <Portfolio />
    <ProblemSolution />
    <WithCanopas />
    <ClientReview />
    <UserReview />
    <Blog />
    <Contribute />
    <CTA />
    <ScreenFooter2 />
  </div>
</template>

<script>
import { defineAsyncComponent } from "vue";

import ScreenHeader from "@/components/partials/ScreenHeaderV2.vue";
import LandingSection from "@/components/home-new/LandingSection.vue";
import CanopasDescription from "@/components/home-new/CanopasDescription.vue";
import Portfolio from "@/components/home-new/PortfolioSection.vue";
import ProblemSolution from "@/components/home-new/ProblemSolution.vue";
import ClientReview from "@/components/home-new/ClientReview.vue";
import WithCanopas from "@/components/home-new/WithCanopas.vue";
import CTA from "@/components/home-new/CTASection.vue";
import Blog from "@/components/home-new/BlogSection.vue";
import Contribute from "@/components/home-new/ContributeSection.vue";
import ScreenFooter2 from "@/components/partials/ScreenFooter2.vue";

import { library } from "@fortawesome/fontawesome-svg-core";
import { faPhone, faPaperPlane } from "@fortawesome/free-solid-svg-icons";

library.add(faPhone, faPaperPlane);

export default {
  components: {
    ScreenHeader,
    LandingSection,
    CanopasDescription,
    Portfolio,
    ProblemSolution,
    WithCanopas,
    ClientReview,
    UserReview: defineAsyncComponent(() =>
      import("@/components/home-new/UserReview.vue")
    ),
    CTA,
    Blog,
    Contribute,
    ScreenFooter2,
  },
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
