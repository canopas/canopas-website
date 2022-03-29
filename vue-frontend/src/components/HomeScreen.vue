<template>
  <div>
    <ScreenMeta v-bind:seoData="seoData" />
    <ScreenHeader />
    <LandingView />
    <ClientReview />
    <UserReviews />
    <ProblemSolution />
    <DoYouKnow v-on:add-animation="handleScroll" />
    <HowDoWeDiffer v-on:add-animation="handleScroll" />
    <CanopasDescription v-on:add-animation="handleScroll" />
    <HomeConsultation />
    <HomeCTA />
    <ScreenFooter2 />
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
import ScreenFooter2 from "./partials/ScreenFooter2.vue";
import ScreenMeta from "./partials/ScreenMeta.vue";
import DoYouKnow from "./home/DoYouKnow.vue";
import HowDoWeDiffer from "./home/HowDoWeDiffer.vue";
import Config from "@/config.js";

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
  data() {
    return {
      seoData: Config.SEO_META_DATA,
    };
  },
  components: {
    ScreenMeta,
    ScreenHeader,
    LandingView,
    ClientReview,
    UserReviews,
    ProblemSolution,
    DoYouKnow,
    HowDoWeDiffer,
    CanopasDescription,
    HomeConsultation,
    HomeCTA,
    ScreenFooter2,
  },
  mounted() {
    this.$gtag.event("view_page_home");
  },
  methods: {
    handleScroll(data) {
      if (data) {
        let { top, bottom } = data.name.getBoundingClientRect();
        let height = document.documentElement.clientHeight;

        if (top < height && bottom > 0) {
          if (data.childRef.length > 0) {
            for (let i = 0; i < data.childRef.length; i++) {
              data.childRef[i].name.classList.add(data.childRef[i].animation);
            }
          } else {
            data.name.classList.add(data.animation);
          }
        }
      }
    },
  },
};
</script>

<style scoped></style>
