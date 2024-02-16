<template>
  <div>
    <ScreenHeader />
    <LandingSection ref="landing" />
    <ServiceSection class="hidden lg:block" ref="service" />
    <ServiceSectionMobile class="block lg:hidden" ref="service" />
    <CaseStudy ref="casestudies" />
    <ClientReview ref="clientReview" />
    <BlogSection ref="blogs" />
    <ContributionSection ref="contributions" />
    <CTASection ref="cta" class="pt-10 bg-[#121212]" />
    <NewFooter class="bg-gradient-to-b from-pink-0 to-pink-16" />
  </div>
</template>

<script>
import ScreenHeader from "@/components/partials/NewHeader.vue";
import LandingSection from "@/components/home-new/LandingSection.vue";
import { elementInViewPort } from "@/utils.js";
import { defineAsyncComponent } from "vue";

const CaseStudy = defineAsyncComponent(
  () => import("@/components/home-new/CaseStudy.vue"),
);
const ServiceSection = defineAsyncComponent(
  () => import("@/components/home-new/ServiceSection.vue"),
);
const ServiceSectionMobile = defineAsyncComponent(
  () => import("@/components/home-new/ServiceSectionMobile.vue"),
);
const ContributionSection = defineAsyncComponent(
  () => import("@/components/home-new/ContributionSection.vue"),
);
const ClientReview = defineAsyncComponent(
  () => import("@/components/home-new/ClientReviewSection.vue"),
);
const CTASection = defineAsyncComponent(
  () => import("@/components/mobile-app-development/CTASection2.vue"),
);
const BlogSection = defineAsyncComponent(
  () => import("@/components/home-new/BlogSection.vue"),
);
const NewFooter = defineAsyncComponent(
  () => import("@/components/partials/NewFooter.vue"),
);

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
    NewFooter,
  },
  data() {
    return {
      width: 0,
      event: "",
      events: {
        landing: "view_landing_section",
        service: "view_service_section",
        casestudies: "view_casestudies_section",
        clientReview: "view_client_review_section",
        blogs: "view_blog_post_section",
        contributions: "view_open_source_contribution",
        cta: "view_home_cta",
      },
    };
  },
  inject: ["mixpanel"],
  mounted() {
    this.width = window.innerWidth;
    window.addEventListener("scroll", this.sendEvent);
    this.$mixpanel.track("view_landing_section");
  },
  unmounted() {
    window.removeEventListener("scroll", this.sendEvent);
  },
  methods: {
    sendEvent() {
      const event = this.events[elementInViewPort(this.$refs)];
      if (event && this.event !== event) {
        this.event = event;
        this.$mixpanel.track(event);
      }
    },
  },
};
</script>
