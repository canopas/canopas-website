<template>
  <div>
    <ScreenHeader />
    <LandingSection />
    <CTASection />
    <FooterV3 />
  </div>
</template>

<script>
import ScreenHeader from "@/components/partials/ScreenHeaderV2.vue";
import LandingSection from "@/components/home-new/LandingSection.vue";
import CTASection from "@/components/home-new/CTASection.vue";
import FooterV3 from "@/components/partials/FooterV3.vue";
import { library } from "@fortawesome/fontawesome-svg-core";
import { faPlus } from "@fortawesome/free-solid-svg-icons";

library.add(faPlus);

import { analyticsEvent } from "@/utils.js";
export default {
  components: {
    ScreenHeader,
    LandingSection,
    CTASection,
    FooterV3,
  },
  data() {
    return {
      event: "",
    };
  },
  mounted() {
    window.addEventListener("scroll", this.sendEvent);
    let recaptchaScript = document.createElement("script");
    recaptchaScript.setAttribute(
      "src",
      "https://www.google.com/recaptcha/enterprise.js?render=" +
        import.meta.env.VITE_RECAPTCHA_SITE_KEY
    );
    recaptchaScript.setAttribute("async", "true");
    recaptchaScript.setAttribute("defer", "true");
    document.head.appendChild(recaptchaScript);
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
