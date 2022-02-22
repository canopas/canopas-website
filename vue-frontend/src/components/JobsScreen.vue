<template>
  <div>
    <ScreenHeader />
    <LandingView v-on:scroll-to-career="scrollToCareer" />
    <VirtuesView />
    <LifeAtCanopas />
    <PerksAndBenifits v-on:scroll-to-career="scrollToCareer" />
    <WhyCanopas class="why-canopas-desktop" />
    <WhyCanopasMobile class="why-canopas-mobile" />
    <Career id="career" />
    <ScreenFooter2 />
  </div>
</template>

<script>
import ScreenHeader from "./partials/ScreenHeader.vue";
import LandingView from "./jobs/LandingView.vue";
import VirtuesView from "./jobs/VirtuesView.vue";
import LifeAtCanopas from "./jobs/LifeAtCanopas.vue";
import Career from "./jobs/CareerView.vue";
import PerksAndBenifits from "./jobs/PerksAndBenifits.vue";
import WhyCanopas from "./jobs/WhyCanopas.vue";
import WhyCanopasMobile from "./jobs/WhyCanopasMobile.vue";
import { library } from "@fortawesome/fontawesome-svg-core";

import { faAlignLeft, faCheckCircle } from "@fortawesome/free-solid-svg-icons";

library.add(faAlignLeft, faCheckCircle);
import ScreenFooter2 from "./partials/ScreenFooter2.vue";

export default {
  components: {
    ScreenHeader,
    LandingView,
    VirtuesView,
    LifeAtCanopas,
    Career,
    PerksAndBenifits,
    WhyCanopas,
    WhyCanopasMobile,
    ScreenFooter2,
  },
  methods: {
    scrollToCareer() {
      var careerDiv = document.getElementById("career");
      var top = careerDiv.offsetTop;
      window.scrollTo({ top: top, behavior: "smooth" });
    },
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
