<template>
  <div>
    <metainfo>
      <template v-slot:title="{ content }">
        {{ content }}
      </template>
    </metainfo>
    <ScreenHeaderV2 />
    <Portfolio />
    <CTA class="mb-5" />
    <ScreenFooterV2 />
  </div>
</template>

<script>
import ScreenHeaderV2 from "@/components/partials/ScreenHeaderV2.vue";
import Portfolio from "@/components/home-new/PortfolioSection.vue";
import CTA from "@/components/home-new/CTASection.vue";
import ScreenFooterV2 from "@/components/partials/ScreenFooter2.vue";
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
    ScreenFooterV2,
    CTA,
  },
  mounted() {
    this.$gtag.event("view_page_portfolio_list");
  },
};
</script>
