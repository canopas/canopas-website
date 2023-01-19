<template>
  <div>
    <ScreenHeader />
    <LandingSection ref="newLanding" />
    <ServiceSection class="tw-hidden md:tw-block" ref="newService" />
    <ServiceSectionMobile class="tw-block md:tw-hidden" ref="newService" />
    <CaseStudy />
    <ClientReview ref="newClientReview" />
    <CTASection ref="newCTA" />
    <FooterV3 />
  </div>
</template>

<script>
import ScreenHeader from "@/components/partials/NewHeader.vue";
import LandingSection from "@/components/home-new/LandingSection.vue";
import CaseStudy from "@/components/home-new/CaseStudy.vue";
import ServiceSection from "@/components/home-new/ServiceSection.vue";
import ServiceSectionMobile from "@/components/home-new/ServiceSectionMobile.vue";
import ClientReview from "@/components/home-new/ClientReviewSection.vue";
import CTASection from "@/components/home-new/CTASection.vue";
import FooterV3 from "@/components/partials/FooterV3.vue";
import { library } from "@fortawesome/fontawesome-svg-core";
import { faPlus } from "@fortawesome/free-solid-svg-icons";
import { analyticsEvent } from "@/utils.js";
library.add(faPlus);

export default {
  components: {
    ScreenHeader,
    LandingSection,
    CaseStudy,
    ServiceSection,
    ServiceSectionMobile,
    ClientReview,
    CTASection,
    FooterV3,
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
