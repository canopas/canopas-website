<template>
    <div>
        <ScreenHeader />
    </div>
</template>

<script>
import ScreenHeader from "./partials/ScreenHeader.vue";

export default {
  components: {
      ScreenHeader,
  },
  mounted() {
    window.gtag("event", "view_canopas_jobs");
  },
};
</script>
