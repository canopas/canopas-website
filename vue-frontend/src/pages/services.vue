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
    </div>
    <NewFooter />
  </div>
</template>

<script>
import Header from "@/components/partials/NewHeader.vue";
import NewFooter from "@/components/partials/NewFooter.vue";
import LandingSection from "@/components/services/LandingSection.vue";
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
