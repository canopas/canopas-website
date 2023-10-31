<template>
  <div>
    <Header />
    <CaseStudy />
    <UserReview
      ref="userReview"
      class="pb-0 sm:pb-10 md:pb-8 lg:pb-20 sm:top-10"
    />
    <CTASection />
    <NewFooter ref="footer" />
  </div>
</template>

<script>
import config from "@/config.js";
import Header from "@/components/partials/NewHeader.vue";
import CaseStudy from "@/components/home-new/CaseStudy.vue";
import { defineAsyncComponent } from "vue";
import { elementInViewPort } from "@/utils.js";

const UserReview = defineAsyncComponent(() =>
  import("@/components/home-new/UserReview.vue"),
);
const CTASection = defineAsyncComponent(() =>
  import("@/components/home-new/CTASection.vue"),
);
const NewFooter = defineAsyncComponent(() =>
  import("@/components/partials/NewFooter.vue"),
);

export default {
  data() {
    return {
      event: "",
      events: {
        footer: "view_portfolio_footer_section",
        userReview: "view_portfolio_user_review_section",
      },
    };
  },
  setup() {
    const seoData = config.PORTFOLIO_SEO_META_DATA;
    useSeoMeta({
      title: seoData.title,
      description: seoData.description,
      ogTitle: seoData.title,
      ogType: seoData.type,
      ogUrl: seoData.url,
      ogImage: seoData.image,
    });
  },
  components: {
    Header,
    CaseStudy,
    UserReview,
    CTASection,
    NewFooter,
  },
  inject: ["mixpanel"],
  mounted() {
    window.addEventListener("scroll", this.sendEvent);
    this.$mixpanel.track("view_portfolio_page");
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
