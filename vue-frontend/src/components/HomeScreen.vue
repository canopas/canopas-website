<template>
  <div>
    <ScreenHeader />
    <LandingView />
    <ClientReview />
    <UserReviews />
    <ProblemSolution />
    <DoYouKnow />
    <HowDoWeDiffer />
    <CanopasDescription />
    <HomeConsultation />
    <HomeCTA />
    <ScreenFooter />
  </div>
</template>

<script>
import ScreenHeader from "./partials/ScreenHeader.vue";
import LandingView from "./home/LandingView.vue";
import UserReviews from "./home/UserReviews.vue";
import ClientReview from "./home/ClientReview.vue";
import ProblemSolution from "./home/ProblemSolution.vue";
import CanopasDescription from "./home/CanopasDescription.vue";
import HomeConsultation from "./home/HomeConsultation.vue";
import HomeCTA from "./home/HomeCTA.vue";
import ScreenFooter from "./partials/ScreenFooter.vue";
import DoYouKnow from "./home/DoYouKnow.vue";
import HowDoWeDiffer from "./home/HowDoWeDiffer.vue";

import { library } from "@fortawesome/fontawesome-svg-core";
import {
  faArrowRight,
  faArrowLeft,
  faStar,
  faPlus,
  faCommentAlt,
  faPhoneAlt,
  faPaperPlane,
} from "@fortawesome/free-solid-svg-icons";

library.add(
  faArrowRight,
  faStar,
  faArrowLeft,
  faPlus,
  faCommentAlt,
  faPhoneAlt,
  faPaperPlane
);

export default {
  components: {
    ScreenHeader,
    LandingView,
    UserReviews,
    ClientReview,
    ProblemSolution,
    CanopasDescription,
    HomeConsultation,
    HomeCTA,
    ScreenFooter,
    DoYouKnow,
    HowDoWeDiffer,
  },
  mounted() {
    window.gtag("event", "view_canopas_home");
  },
};
</script>

<style scoped></style>
