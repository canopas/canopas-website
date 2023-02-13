<template>
  <div>
    <metainfo>
      <template v-slot:title="{ content }">
        {{ content }}
      </template>
    </metainfo>
    <Header />
    <Portfolio />
    <NewFooter />
  </div>
</template>

<script>
import Header from "@/components/partials/NewHeader.vue";
import Portfolio from "@/components/home/PortfolioSection.vue";
import NewFooter from "@/components/partials/NewFooter.vue";
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
    Header,
    Portfolio,
    NewFooter,
  },
  inject: ["mixpanel"],
  mounted() {
    this.mixpanel.track("view_page_portfolio_list");
  },
};
</script>
