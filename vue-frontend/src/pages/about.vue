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
      <AboutusVirtue />
      <WithCanopasSection />
      <ClientReviewSection />
      <CTASection />
    </div>
    <NewFooter />
  </div>
</template>

<script>
import Header from "@/components/partials/NewHeader.vue";
import NewFooter from "@/components/partials/NewFooter.vue";
import LandingSection from "@/components/about/LandingSection.vue";
import AboutusVirtue from "@/components/about/AboutusVirtue.vue";
import WithCanopasSection from "@/components/home/WithCanopas.vue";
import ClientReviewSection from "@/components/home-new/ClientReviewSection.vue";
import CTASection from "@/components/about/CTASection.vue";
import config from "@/config.js";
import { useMeta } from "vue-meta";

export default {
  setup() {
    var seoData = config.ABOUT_SEO_META_DATA;
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
    AboutusVirtue,

    WithCanopasSection,
    ClientReviewSection,
    CTASection,
    NewFooter,
  },
  inject: ["mixpanel"],
  beforeRouteEnter(to, from, next) {
    if (config.SHOW_ABOUT_US_PAGE === false) {
      next({ name: "Error404Page", params: { pathMatch: "/about" } });
    } else {
      next();
    }
  },
  mounted() {
    this.mixpanel.track("view_about_page");
  },
};
</script>
