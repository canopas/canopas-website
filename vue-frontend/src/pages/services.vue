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
      <WhatWeOfferMobile class="tw-block lg:tw-hidden" />
      <WhatWeOfferDesktop class="tw-hidden lg:tw-block" />
      <ContactUsToday />
      <TechnologyStack />
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
import WhatWeOfferMobile from "@/components/services/WhatWeOfferMobile.vue";
import WhatWeOfferDesktop from "@/components/services/WhatWeOfferDesktop.vue";
import config from "@/config.js";
import { useMeta } from "vue-meta";
import { defineAsyncComponent } from "vue";

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
  import("@/components/home-new/PortfolioPageCTASection.vue")
);
const NewFooter = defineAsyncComponent(() =>
  import("@/components/partials/NewFooter.vue")
);
const TechnologyStack = defineAsyncComponent(() =>
  import("@/components/services/TechnologyStack.vue")
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
    WhatWeOfferMobile,
    WhatWeOfferDesktop,
    ContactUsToday,
    TechnologyStack,
    SuccessStories,
    BlogSection,
    ContributionSection,
    ContributionSectionMobile,
    ClientReviewSection,
    CTASection,
    NewFooter,
  },
  inject: ["mixpanel"],
  mounted() {
    this.mixpanel.track("view_services_page");
  },
};
</script>
