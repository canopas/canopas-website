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
      <ContactUsToday />
      <BlogSection />
      <ContributionSection class="tw-hidden lg:tw-block" />
      <ContributionSectionMobile class="tw-block lg:tw-hidden" />
      <SuccessStories />
      <ClientReviewSection />
    </div>
    <NewFooter />
  </div>
</template>

<script>
import Header from "@/components/partials/NewHeader.vue";
import NewFooter from "@/components/partials/NewFooter.vue";
import LandingSection from "@/components/services/LandingSection.vue";
import ContactUsToday from "@/components/services/ContactUsToday.vue";
import SuccessStories from "@/components/services/SuccessStories.vue";
import BlogSection from "@/components/home-new/BlogSection.vue";
import ContributionSection from "@/components/home-new/ContributionSection.vue";
import ContributionSectionMobile from "@/components/home-new/ContributionSectionMobile.vue";
import ClientReviewSection from "@/components/home-new/ClientReviewSection.vue";
import config from "@/config.js";
import { useMeta } from "vue-meta";
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
    ContactUsToday,
    SuccessStories,
    BlogSection,
    ContributionSection,
    ContributionSectionMobile,
    ClientReviewSection,
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
