<template>
  <div>
    <ScreenHeader />
    <LandingSection />
    <ServiceSection class="tw-hidden md:tw-block" />
    <ServiceSectionMobile class="tw-block md:tw-hidden" />
    <ClientReview />
    <CTASection />
    <FooterV3 />
  </div>
</template>

<script>
import ScreenHeader from "@/components/partials/NewHeader.vue";
import LandingSection from "@/components/home-new/LandingSection.vue";
import ServiceSection from "@/components/home-new/ServiceSection.vue";
import ServiceSectionMobile from "@/components/home-new/ServiceSectionMobile.vue";
import ClientReview from "@/components/home-new/ClientReviewSection.vue";
import CTASection from "@/components/home-new/CTASection.vue";
import FooterV3 from "@/components/partials/FooterV3.vue";
import { library } from "@fortawesome/fontawesome-svg-core";
import { faPlus } from "@fortawesome/free-solid-svg-icons";
library.add(faPlus);

export default {
  components: {
    ScreenHeader,
    LandingSection,
    ServiceSection,
    ServiceSectionMobile,
    ClientReview,
    CTASection,
    FooterV3,
  },
};
</script>
