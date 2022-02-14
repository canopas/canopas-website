<template>
  <div>
    <ScreenHeader />
    <LandingView />
    <VirtuesView />
    <LifeAtCanopas />
    <PerksAndBenifits />
    <WhyCanopas class="why-canopas-desktop" />
    <WhyCanopasMobile class="why-canopas-mobile" />
    <Career />
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
