<template>
  <div>
    <ScreenHeader v-if="!isShowNewHeader" />
    <ScreenHeaderV2 v-else />
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
import ScreenHeader from "@/components/partials/ScreenHeader.vue";
import ScreenHeaderV2 from "@/components/partials/ScreenHeaderV2.vue";
import LandingView from "@/components/home/LandingView.vue";
import ClientReview from "@/components/home/ClientReview.vue";
import UserReviews from "@/components/home/UserReviews.vue";
import ProblemSolution from "@/components/home/ProblemSolution.vue";
import DoYouKnow from "@/components/home/DoYouKnow.vue";
import HowDoWeDiffer from "@/components/home/HowDoWeDiffer.vue";
import CanopasDescription from "@/components/home/CanopasDescription.vue";
import HomeConsultation from "@/components/home/HomeConsultation.vue";
import HomeCTA from "@/components/home/HomeCTA.vue";
import ScreenFooter2 from "@/components/partials/ScreenFooter2.vue";
import Config from "@/config.js";
import { useMeta } from "vue-meta";
import config from "@/config.js";

import { library } from "@fortawesome/fontawesome-svg-core";
import {
  faArrowRight,
  faArrowLeft,
  faStar,
  faPlus,
  faCommentAlt,
  faPhone,
  faPaperPlane,
} from "@fortawesome/free-solid-svg-icons";

library.add(
  faArrowRight,
  faStar,
  faArrowLeft,
  faPlus,
  faCommentAlt,
  faPhone,
  faPaperPlane
);

export default {
  setup() {
    var seoData = Config.SEO_META_DATA;
    useMeta({
      title: seoData.title,
      description: seoData.description,
      og: {
        type: seoData.type,
        title: seoData.title,
        url: seoData.url,
      },
    });
  },
  data() {
    return {
      isShowNewHeader: config.IS_SHOW_NEW_HOME_PAGE,
    };
  },
  components: {
    ScreenHeader,
    ScreenHeaderV2,
    LandingView,
    ClientReview,
    UserReviews,
    ProblemSolution,
    DoYouKnow,
    CanopasDescription,
    HowDoWeDiffer,
    HomeConsultation,
    HomeCTA,
    ScreenFooter2,
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
