<template>
  <div>
    <ScreenHeader />
    <VirtuesView />
    <LifeAtCanopas />
  </div>
</template>

<script>
import ScreenHeader from "./partials/ScreenHeader.vue";
import VirtuesView from "./jobs/VirtuesView.vue";
import LifeAtCanopas from "./jobs/LifeAtCanopas.vue";

export default {
  components: {
    ScreenHeader,
    VirtuesView,
    LifeAtCanopas,
  },
};
</script>
