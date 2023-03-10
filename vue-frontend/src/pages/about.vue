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
      <CTASection />
    </div>
    <NewFooter />
  </div>
</template>

<script>
import Header from "@/components/partials/NewHeader.vue";
import NewFooter from "@/components/partials/NewFooter.vue";
import LandingSection from "@/components/about/LandingSection.vue";
import CTASection from "@/components/about/CTASection.vue";
import config from "@/config.js";
import { useMeta } from "vue-meta";

export default {
  setup() {
    var seoData = config.ABOUT_SEO_META_DATA;
    useMeta({
      title: seoData.title,
      description: seoData.description,
      og: {
        type: seoData.type,
        title: seoData.title,
        url: seoData.url,
      },
    });
  },
  components: {
    Header,
    LandingSection,
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
