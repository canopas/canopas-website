<template>
    <div>
        <ScreenHeader />
        <VirtuesView />
        <LandingView />
    </div>
</template>

<script>
import ScreenHeader from "./partials/ScreenHeader.vue";
import LandingView from "./jobs/LandingView.vue";
import VirtuesView from "./jobs/VirtuesView.vue";


export default {
  components: {
      ScreenHeader,
      LandingView,
      VirtuesView,
      
  },
};
</script>
