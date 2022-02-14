<template>
  <div>
    <ScreenHeader />
    <LandingView />
    <VirtuesView />
    <LifeAtCanopas />
    <PerksAndBenifits />
    <WhyCanopas v-if="!isMobile" />
    <WhyCanopasMobile v-if="isMobile" />
  </div>
</template>

<script>
import ScreenHeader from "./partials/ScreenHeader.vue";
import LandingView from "./jobs/LandingView.vue";
import VirtuesView from "./jobs/VirtuesView.vue";
import LifeAtCanopas from "./jobs/LifeAtCanopas.vue";
import PerksAndBenifits from "./jobs/PerksAndBenifits.vue";
import WhyCanopas from "./jobs/WhyCanopas.vue";
import WhyCanopasMobile from "./jobs/WhyCanopasMobile.vue";

export default {
  data() {
    return {
      isMobile: false,
    };
  },
  components: {
    ScreenHeader,
    LandingView,
    VirtuesView,
    LifeAtCanopas,
    PerksAndBenifits,
    WhyCanopas,
    WhyCanopasMobile,
  },
  mounted: function () {
    this.showWhyCanopas();
    window.addEventListener("resize", this.showWhyCanopas);
  },
  unmounted: function () {
    window.removeEventListener("resize", this.showWhyCanopas);
  },
  methods: {
    showWhyCanopas() {
      this.isMobile = window.innerWidth <= 768;
    },
  },
};
</script>
