<template>
  <div>
    <metainfo>
      <template v-slot:title="{ content }">
        {{ content }}
      </template>
    </metainfo>
    <Header />

    <div>
      <LandingSection ref="landing" />
      <WhatWeOfferMobile class="tw-block lg:tw-hidden" ref="whatweoffer" />
      <WhatWeOfferDesktop class="tw-hidden lg:tw-block" ref="whatweoffer" />
      <CTASection class="!tw-mt-24 xl:!tw-mt-40" ref="cta" />
      <TechnologyStack ref="technologystack" />
      <BlogSection ref="blogs" />
      <ContributionSection class="tw-hidden lg:tw-block" ref="contributions" />
      <ContributionSectionMobile
        class="tw-block lg:tw-hidden"
        ref="contributions"
      />
      <SuccessStories ref="successstories" />
      <ClientReviewSection ref="clientreview" />
      <CTASection2 ref="cta2" />
    </div>

    <NewFooter />
  </div>
</template>

<script>
import Header from "@/components/partials/NewHeader.vue";
import LandingSection from "@/components/services/LandingSection.vue";
import WhatWeOfferMobile from "@/components/services/WhatWeOfferMobile.vue";
import WhatWeOfferDesktop from "@/components/services/WhatWeOfferDesktop.vue";
import config from "@/config.js";
// import { useMeta } from "vue-meta";
import { defineAsyncComponent } from "vue";
import { elementInViewPort } from "@/utils.js";

const SuccessStories = defineAsyncComponent(() =>
  import("@/components/services/SuccessStories.vue")
);
const BlogSection = defineAsyncComponent(() =>
  import("@/components/home-new/BlogSection.vue")
);
const CTASection = defineAsyncComponent(() =>
  import("@/components/home-new/PortfolioPageCTASection.vue")
);
const ContributionSection = defineAsyncComponent(() =>
  import("@/components/home-new/ContributionSection.vue")
);
const ContributionSectionMobile = defineAsyncComponent(() =>
  import("@/components/home-new/ContributionSectionMobile.vue")
);
const ClientReviewSection = defineAsyncComponent(() =>
  import("@/components/home-new/ClientReviewSection.vue")
);
const CTASection2 = defineAsyncComponent(() =>
  import("@/components/services/CTASection.vue")
);
const TechnologyStack = defineAsyncComponent(() =>
  import("@/components/services/TechnologyStack.vue")
);
const NewFooter = defineAsyncComponent(() =>
  import("@/components/partials/NewFooter.vue")
);

export default {
  setup() {
    var seoData = config.SERVICES_SEO_META_DATA;
    // useMeta({
    //   title: seoData.title,
    //   description: seoData.description,
    //   og: {
    //     type: seoData.type,
    //     title: seoData.title,
    //     url: seoData.url,
    //     image: seoData.image,
    //   },
    // });
  },
  components: {
    Header,
    LandingSection,
    WhatWeOfferMobile,
    WhatWeOfferDesktop,
    CTASection,
    TechnologyStack,
    SuccessStories,
    BlogSection,
    ContributionSection,
    ContributionSectionMobile,
    ClientReviewSection,
    CTASection2,
    NewFooter,
  },
  data() {
    return {
      event: "",
      events: {
        landing: "view_services_landing_section",
        whatweoffer: "view_services_whatweoffer_section",
        cta: "view_services_cta_section",
        technologystack: "view_services_technologystack_section",
        blogs: "view_services_blogs_section",
        contributions: "view_services_contributions_section",
        successstories: "view_services_successstories_section",
        clientreview: "view_services_clientreview_section",
        footer: "view_mobileapp_development_footer",
      },
    };
  },
  // inject: ["mixpanel"],
  mounted() {
    window.addEventListener("scroll", this.sendEvent);
    // // this.mixPanel.track("view_services_page");
  },
  unmounted() {
    window.removeEventListener("scroll", this.sendEvent);
  },
  methods: {
    sendEvent() {
      const event = this.events[elementInViewPort(this.$refs)];
      if (event && this.event !== event) {
        this.event = event;
        // // this.mixPanel.track(event);
      }
    },
  },
};
</script>
