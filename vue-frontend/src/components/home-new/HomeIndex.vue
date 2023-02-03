<template>
  <div>
    <ScreenHeader />
    <LandingSection ref="landing" />
    <ServiceSection class="tw-hidden md:tw-block" ref="service" />
    <ServiceSectionMobile class="tw-block md:tw-hidden" ref="service" />
    <CaseStudy />
    <ClientReview ref="clientReview" />
    <BlogSection ref="blogs" />
    <ContributionSection ref="contributions" />
    <CTASection ref="cta" />
    <FooterV3 ref="footer" />
  </div>
</template>

<script>
import ScreenHeader from "@/components/partials/NewHeader.vue";
import LandingSection from "@/components/home-new/LandingSection.vue";
import CaseStudy from "@/components/home-new/CaseStudy.vue";
import ServiceSection from "@/components/home-new/ServiceSection.vue";
import ServiceSectionMobile from "@/components/home-new/ServiceSectionMobile.vue";
import ContributionSection from "@/components/home-new/ContributionSection.vue";
import ClientReview from "@/components/home-new/ClientReviewSection.vue";
import CTASection from "@/components/home-new/CTASection.vue";
import BlogSection from "@/components/home-new/BlogSection.vue";
import FooterV3 from "@/components/partials/FooterV3.vue";
import { library } from "@fortawesome/fontawesome-svg-core";
import { faPlus } from "@fortawesome/free-solid-svg-icons";
import { elementInViewPort } from "@/utils.js";
library.add(faPlus);

export default {
  components: {
    ScreenHeader,
    LandingSection,
    CaseStudy,
    ServiceSection,
    ServiceSectionMobile,
    ClientReview,
    ContributionSection,
    CTASection,
    BlogSection,
    FooterV3,
  },
  data() {
    return {
      event: "",
      events: {
        landing: "view_landing_section",
        service: "view_service_section",
        clientReview: "view_client_review_section",
        blogs: "view_blog_post_section",
        contributions: "view_open_source_contribution",
        cta: "view_home_cta",
        footer: "view_home_footer",
      },
    };
  },
  inject: ["mixpanel"],
  mounted() {
    this.mixpanel.track("view_landing_section");
    window.addEventListener("scroll", this.sendEvent);
  },
  unmounted() {
    window.removeEventListener("scroll", this.sendEvent);
  },
  methods: {
    sendEvent() {
      const event = this.events[elementInViewPort(this.$refs)];
      if (event && this.event !== event) {
        this.event = event;
        this.mixpanel.track(event);
      }
    },
  },
};
</script>
