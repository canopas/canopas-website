<template>
  <div>
    <Header />
    <div>
      <LandingSection />
      <WhatWeOfferMobile class="block lg:hidden" />
      <WhatWeOfferDesktop class="hidden lg:block" />
      <CTA4 />
      <TechnologyStack />
      <BlogSection ref="blogs" />
      <ContributionSection />
      <SuccessStories />
      <ClientReviewSection />
      <CTASection ref="cta2" />
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
import { defineAsyncComponent } from "vue";
import { elementInViewPort } from "@/utils.js";

const SuccessStories = defineAsyncComponent(
  () => import("@/components/services/SuccessStories.vue"),
);
const BlogSection = defineAsyncComponent(
  () => import("@/components/home-new/BlogSection.vue"),
);
const CTA4 = defineAsyncComponent(() => import("@/components/CTA/CTA4.vue"));
const ContributionSection = defineAsyncComponent(
  () => import("@/components/home-new/ContributionSection.vue"),
);
const ClientReviewSection = defineAsyncComponent(
  () => import("@/components/home-new/ClientReviewSection.vue"),
);
const CTASection = defineAsyncComponent(
  () => import("@/components/services/CTASection.vue"),
);
const TechnologyStack = defineAsyncComponent(
  () => import("@/components/services/TechnologyStack.vue"),
);
const NewFooter = defineAsyncComponent(
  () => import("@/components/partials/NewFooter.vue"),
);

export default {
  setup() {
    const seoData = config.SERVICES_SEO_META_DATA;
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
    LandingSection,
    WhatWeOfferMobile,
    WhatWeOfferDesktop,
    CTA4,
    TechnologyStack,
    SuccessStories,
    BlogSection,
    ContributionSection,
    ClientReviewSection,
    CTASection,
    NewFooter,
  },
  data() {
    return {
      event: "",
      events: {
        blogs: "view_services_blogs_section",
        cta2: "view_mobileapp_development_cta2",
      },
    };
  },
  inject: ["mixpanel"],
  mounted() {
    window.addEventListener("scroll", this.sendEvent);
    this.$mixpanel.track("view_services_page");
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
