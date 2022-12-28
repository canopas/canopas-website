<template>
  <div>
    <ScreenHeader />
    <LandingSection />
    <CTASection />
  </div>
</template>

<script>
import ScreenHeader from "@/components/partials/ScreenHeaderV2.vue";
import LandingSection from "@/components/home-new/LandingSection.vue";
import CTASection from "@/components/home-new/CTASection.vue";
import { library } from "@fortawesome/fontawesome-svg-core";
import { faPlus } from "@fortawesome/free-solid-svg-icons";

library.add(faPlus);

import { analyticsEvent } from "@/utils.js";
export default {
  components: {
    ScreenHeader,
    LandingSection,
    CTASection,
  },
  data() {
    return {
      event: "",
    };
  },
  mounted() {
    this.$gtag.event("view_landing_section");
    window.addEventListener("scroll", this.sendEvent);
  },
  unmounted() {
    window.removeEventListener("scroll", this.sendEvent);
  },
  methods: {
    sendEvent() {
      const event = analyticsEvent(this.$refs);
      if (event && this.event !== event) {
        this.event = event;
        this.$gtag.event(event);
      }
    },
  },
};
</script>
