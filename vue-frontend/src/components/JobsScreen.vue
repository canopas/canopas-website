<template>
    <div>
        <ScreenHeader />
        <VirtuesView />
    </div>
</template>

<script>
import ScreenHeader from "./partials/ScreenHeader.vue";
import VirtuesView from "./jobs/VirtuesView.vue";

export default {
  components: {
      ScreenHeader,
      VirtuesView,
  },
};
</script>
