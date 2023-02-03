<template>
  <div>
    <ScreenHeader />
    <div
      class="tw-from-[#ffffff] tw-to-[#ff947226] tw-bg-gradient-[174.39deg] tw-z-[-5]"
    >
      <LandingSection ref="landing" />
      <CanopasDescription ref="budget" />
    </div>
    <Portfolio ref="portfolio" />
    <ProblemSolution ref="parallax" />
    <WithCanopas ref="service" />
    <ClientReview ref="clientreview" />
    <UserReview ref="userreview" />
    <Blog ref="blog" />
    <Contribute ref="contribution" />
    <CTA class="tw-mt-12 tw-mb-12" ref="CTA" />
    <ScreenFooter2 ref="footer" />
  </div>
</template>

<script>
import ScreenHeader from "@/components/partials/ScreenHeaderV2.vue";
import LandingSection from "@/components/home/LandingSection.vue";
import CanopasDescription from "@/components/home/CanopasDescription.vue";
import Portfolio from "@/components/home/PortfolioSection.vue";
import ProblemSolution from "@/components/home/ProblemSolution.vue";
import ClientReview from "@/components/home/ClientReview.vue";
import WithCanopas from "@/components/home/WithCanopas.vue";
import CTA from "@/components/home/CTASection.vue";
import Blog from "@/components/home/BlogSection.vue";
import Contribute from "@/components/home/ContributeSection.vue";
import ScreenFooter2 from "@/components/partials/ScreenFooter2.vue";
import UserReview from "@/components/home/UserReview.vue";
import { library } from "@fortawesome/fontawesome-svg-core";
import { faPhone, faPaperPlane } from "@fortawesome/free-solid-svg-icons";
import { elementInViewPort } from "@/utils.js";

library.add(faPhone, faPaperPlane);

export default {
  components: {
    ScreenHeader,
    LandingSection,
    CanopasDescription,
    Portfolio,
    ProblemSolution,
    WithCanopas,
    ClientReview,
    UserReview,
    CTA,
    Blog,
    Contribute,
    ScreenFooter2,
  },
  data() {
    return {
      event: "",
      events: {
        landing: "view_landing_section",
        budget: "view_ad_budget_section",
        portfolio: "view_home_portfolio_section",
        parallax: "view_parallax_section",
        clientreview: "view_client_review_section",
        userreview: "view_user_review_section",
        blog: "view_blog_post_section",
        contribution: "view_open_source_contribution",
        CTA: "view_home_page_end_cta",
        footer: "view_home_page_footer",
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
