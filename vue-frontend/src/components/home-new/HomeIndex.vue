<template>
  <div>
    <ScreenHeader />
    <LandingSection ref="landing" />
    <ServiceSection class="tw-hidden md:tw-block" ref="service" />
    <ServiceSectionMobile class="tw-block md:tw-hidden" ref="service" />
    <CaseStudy ref="casestudies" />
    <ClientReview ref="clientReview" />
    <BlogSection ref="blogs" />
    <ContributionSection class="tw-hidden lg:tw-block" ref="contributions" />
    <ContributionSectionMobile
      class="tw-block lg:tw-hidden"
      ref="contributions"
    />
    <CTASection
      ref="cta"
      class="!tw-mt-[100px] md:!tw-mt-[220px] tw-pt-10 tw-bg-black-core/[0.85] tw-mb-[-15px]"
    />
    <NewFooter ref="footer" />
  </div>
</template>

<script>
import ScreenHeader from "@/components/partials/NewHeader.vue";
import LandingSection from "@/components/home-new/LandingSection.vue";
import { library } from "@fortawesome/fontawesome-svg-core";
import { faPlus } from "@fortawesome/free-solid-svg-icons";
import { elementInViewPort } from "@/utils.js";
import { defineAsyncComponent } from "vue";

const CaseStudy = defineAsyncComponent(() =>
  import("@/components/home-new/CaseStudy.vue"),
);
const ServiceSection = defineAsyncComponent(() =>
  import("@/components/home-new/ServiceSection.vue"),
);
const ServiceSectionMobile = defineAsyncComponent(() =>
  import("@/components/home-new/ServiceSectionMobile.vue"),
);
const ContributionSection = defineAsyncComponent(() =>
  import("@/components/home-new/ContributionSection.vue"),
);
const ContributionSectionMobile = defineAsyncComponent(() =>
  import("@/components/home-new/ContributionSectionMobile.vue"),
);
const ClientReview = defineAsyncComponent(() =>
  import("@/components/home-new/ClientReviewSection.vue"),
);
const CTASection = defineAsyncComponent(() =>
  import("@/components/mobile-app-development/CTASection2.vue"),
);
const BlogSection = defineAsyncComponent(() =>
  import("@/components/home-new/BlogSection.vue"),
);
const NewFooter = defineAsyncComponent(() =>
  import("@/components/partials/NewFooter.vue"),
);

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
    ContributionSectionMobile,
    CTASection,
    BlogSection,
    NewFooter,
  },
  data() {
    return {
      event: "",
      events: {
        landing: "view_landing_section",
        service: "view_service_section",
        casestudies: "view_casestudies_section",
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
    window.addEventListener("scroll", this.sendEvent);
    this.mixpanel.track("view_landing_section");
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
