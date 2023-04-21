<template>
  <div>
    <metainfo>
      <template v-slot:title="{ content }">
        {{ content }}
      </template>
    </metainfo>
    <Header />
    <CaseStudy ref="newPortfolio" />
    <UserReview
      ref="userReview"
      class="tw-pb-0 sm:tw-pb-10 md:tw-pb-8 lg:tw-pb-20 sm:tw-top-[40px]"
    />
    <CTASection ref="cta" />
    <NewFooter />
  </div>
</template>

<script>
import { useMeta } from "vue-meta";
import config from "@/config.js";
import Header from "@/components/partials/NewHeader.vue";
import CaseStudy from "@/components/home-new/CaseStudy.vue";
import NewFooter from "@/components/partials/NewFooter.vue";
import { defineAsyncComponent } from "vue";

const UserReview = defineAsyncComponent(() =>
  import("@/components/home/UserReview.vue")
);
const CTASection = defineAsyncComponent(() =>
  import("@/components/home-new/CTASection.vue")
);

export default {
  data() {
    return {
      event: "",
      events: {
        newPortfolio: "view_portfolio_section",
        cta: "view_cta_section",
        userReview: "view_user_review_section",
      },
    };
  },
  setup() {
    var seoData = config.PORTFOLIO_SEO_META_DATA;
    useMeta({
      title: seoData.title,
      description: seoData.description,
      og: {
        type: seoData.type,
        title: seoData.title,
        url: seoData.url,
        image: seoData.image,
      },
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
    this.mixpanel.track("view_portfolio_page");
  },
};
</script>
