<template>
  <div>
    <ScreenMeta v-bind:seoData="seoData" />
    <ScreenHeader />
    <LandingView v-on:scroll-to-career="scrollToCareer" />
    <VirtuesView />
    <LifeAtCanopas />
    <PerksAndBenifits v-on:scroll-to-career="scrollToCareer" />
    <WhyCanopas class="why-canopas-desktop" v-on:add-animation="handleScroll" />
    <WhyCanopasMobile
      class="why-canopas-mobile"
      v-on:add-animation="handleScroll"
    />
    <Career id="career" />
    <FaqSection v-on:add-animation="handleScroll" />
    <ScreenFooter2 />
  </div>
</template>

<script>
import ScreenMeta from "./partials/ScreenMeta.vue";
import ScreenHeader from "./partials/ScreenHeader.vue";
import LandingView from "./jobs/LandingView.vue";
import VirtuesView from "./jobs/VirtuesView.vue";
import LifeAtCanopas from "./jobs/LifeAtCanopas.vue";
import Career from "./jobs/CareerView.vue";
import PerksAndBenifits from "./jobs/PerksAndBenifits.vue";
import WhyCanopas from "./jobs/WhyCanopas.vue";
import WhyCanopasMobile from "./jobs/WhyCanopasMobile.vue";
import FaqSection from "./jobs/FaqSection.vue";
import ScreenFooter2 from "./partials/ScreenFooter2.vue";
import config from "@/config.js";
import { library } from "@fortawesome/fontawesome-svg-core";

import {
  faAlignLeft,
  faCheckCircle,
  faMinus,
} from "@fortawesome/free-solid-svg-icons";

library.add(faAlignLeft, faCheckCircle, faMinus);

export default {
  data() {
    return {
      seoData: config.JOBS_SEO_META_DATA,
    };
  },
  components: {
    ScreenHeader,
    LandingView,
    VirtuesView,
    LifeAtCanopas,
    Career,
    PerksAndBenifits,
    WhyCanopas,
    WhyCanopasMobile,
    FaqSection,
    ScreenFooter2,
    ScreenMeta,
  },
  methods: {
    scrollToCareer() {
      var careerDiv = document.getElementById("career");
      var top = careerDiv.offsetTop;
      window.scrollTo({ top: top, behavior: "smooth" });
    },
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
  mounted() {
    this.$gtag.event("view_page_jobs");
  },
};
</script>

<style lang="scss" scoped>
.why-canopas-mobile {
  display: block;
}
.why-canopas-desktop {
  display: none;
}

@include media-breakpoint-up(md) {
  .why-canopas-mobile {
    display: none;
  }
  .why-canopas-desktop {
    display: block;
  }
}
</style>
