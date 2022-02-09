<template>
  <div>
    <ScreenHeader />
    <LandingView />
    <VirtuesView />
    <LifeAtCanopas />
    <PreksAndBenifits />
  </div>
</template>

<script>
import ScreenHeader from "./partials/ScreenHeader.vue";
import LandingView from "./jobs/LandingView.vue";
import VirtuesView from "./jobs/VirtuesView.vue";
import LifeAtCanopas from "./jobs/LifeAtCanopas.vue";
import PreksAndBenifits from "./jobs/PerksAndBenifits.vue";

export default {
  components: {
    ScreenHeader,
    LandingView,
    VirtuesView,
    LifeAtCanopas,
    PreksAndBenifits,
  },
};
</script>
