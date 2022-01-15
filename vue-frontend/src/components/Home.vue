<template>
  <div>
    <Header />
    <LandingView />
    <ClientReview />
    <UserReviews />
    <ProblemSolution />
    <CanopasDescription />
    <Consultation />
    <CTA />
    <Footer />
  </div>
</template>

<script>
import Header from "./partials/Header.vue";
import LandingView from "./home/LandingView.vue";
import UserReviews from "./home/UserReviews.vue";
import ClientReview from "./home/ClientReview.vue";
import ProblemSolution from "./home/ProblemSolution.vue";
import CanopasDescription from "./home/CanopasDescription.vue";
import Consultation from "./home/Consultation.vue";
import CTA from "./home/CTA.vue";
import Footer from "./partials/Footer.vue";

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
  name: "App",
  components: {
    Header,
    LandingView,
    UserReviews,
    ClientReview,
    ProblemSolution,
    CanopasDescription,
    Consultation,
    CTA,
    Footer,
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
