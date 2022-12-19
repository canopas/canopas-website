<template>
  <div>
    <metainfo>
      <template v-slot:title="{ content }">
        {{ content }}
      </template>
    </metainfo>
    <ScreenHeaderV2 />
    <Portfolio />
    <CTA />
  </div>
</template>

<script>
import ScreenHeaderV2 from "@/components/partials/ScreenHeaderV2.vue";
import Portfolio from "@/components/home/PortfolioSection.vue";
import CTA from "@/components/home/CTASection.vue";
import { useMeta } from "vue-meta";
import config from "@/config.js";

export default {
  setup() {
    var seoData = config.PORTFOLIO_SEO_META_DATA;
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
    ScreenHeaderV2,
    Portfolio,
    CTA,
  },
  mounted() {
    this.$gtag.event("view_page_portfolio_list");
  },
};
</script>
