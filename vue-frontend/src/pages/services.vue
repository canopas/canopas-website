<template>
  <div>
    <metainfo>
      <template v-slot:title="{ content }">
        {{ content }}
      </template>
    </metainfo>
    <Header />

    <div>
      <LandingSection />

      <WhatWeOfferSection class="tw-block lg:tw-hidden" />
      <WhatWeOfferDesktopSection class="tw-hidden lg:tw-block" />

      <ContactUsToday />
      <TechnologyStackMobile class="tw-block lg:tw-hidden" />
      <TechnologyStackDesktop class="tw-hidden lg:tw-block" />
      <BlogSection />
      <ContributionSection class="tw-hidden lg:tw-block" />
      <ContributionSectionMobile class="tw-block lg:tw-hidden" />
      <SuccessStories />
      <ClientReviewSection />
      <CTASection />
    </div>

    <NewFooter />
  </div>
</template>

<script>
import Header from "@/components/partials/NewHeader.vue";
import LandingSection from "@/components/services/LandingSection.vue";
import TechnologyStackMobile from "@/components/services/TechnologyStackMobile.vue";
import TechnologyStackDesktop from "@/components/services/TechnologyStackDesktop.vue";
import ContactUsToday from "@/components/services/ContactUsToday.vue";
import SuccessStories from "@/components/services/SuccessStories.vue";
import BlogSection from "@/components/home-new/BlogSection.vue";
import ContributionSection from "@/components/home-new/ContributionSection.vue";
import ContributionSectionMobile from "@/components/home-new/ContributionSectionMobile.vue";
import ClientReviewSection from "@/components/home-new/ClientReviewSection.vue";
import CTASection from "@/components/services/CTASection.vue";
import config from "@/config.js";
import { useMeta } from "vue-meta";
import { defineAsyncComponent } from "vue";

const WhatWeOfferSection = defineAsyncComponent(() =>
  import("@/components/services/WhatWeOfferSection.vue")
);
const WhatWeOfferDesktopSection = defineAsyncComponent(() =>
  import("@/components/services/WhatWeOfferDesktop.vue")
);
const SuccessStories = defineAsyncComponent(() =>
  import("@/components/services/SuccessStories.vue")
);
const BlogSection = defineAsyncComponent(() =>
  import("@/components/home-new/BlogSection.vue")
);
const ContactUsToday = defineAsyncComponent(() =>
  import("@/components/services/ContactUsToday.vue")
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
const CTASection = defineAsyncComponent(() =>
  import("@/components/services/CTASection.vue")
);
const NewFooter = defineAsyncComponent(() =>
  import("@/components/partials/NewFooter.vue")
);

export default {
  setup() {
    var seoData = config.SERVICES_SEO_META_DATA;
    useMeta({
      meta: [
        {
          name: "robots",
          content: "noindex, nofollow",
          vmid: "robots",
        },
      ],
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
    LandingSection,
    TechnologyStackMobile,
    TechnologyStackDesktop,
    WhatWeOfferSection,
    WhatWeOfferDesktopSection,
    ContactUsToday,
    SuccessStories,
    BlogSection,
    ContributionSection,
    ContributionSectionMobile,
    ClientReviewSection,
    CTASection,
    NewFooter,
  },
  inject: ["mixpanel"],
  beforeRouteEnter(to, from, next) {
    if (config.SHOW_SERVICES_PAGE === false) {
      next({ name: "Error404Page", params: { pathMatch: "/services" } });
    } else {
      next();
    }
  },
  mounted() {
    this.mixpanel.track("view_services_page");
  },
};
</script>
