<template>
    <div>
        <ScreenHeader />
        <VirtuesView />
        <LandingView />
    </div>
</template>

<script>
import ScreenHeader from "./partials/ScreenHeader.vue";
import VirtuesView from "./jobs/VirtuesView.vue";
import VirtuesView from "./jobs/LandingView.vue";

export default {
  components: {
      ScreenHeader,
      VirtuesView,
  },
};
</script>
