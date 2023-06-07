<template>
  <div>
    <metainfo>
      <template v-slot:title="{ content }">
        {{ content }}
      </template>
    </metainfo>
    <Header />
    <div>
      <GithubContribution />
    </div>
    <NewFooter />
  </div>
</template>

<script>
import Header from "@/components/partials/NewHeader.vue";
import GithubContribution from "@/components/contribution/GithubContribution.vue";
import config from "@/config.js";
import { useMeta } from "vue-meta";
import NewFooter from "@/components/partials/NewFooter.vue";

export default {
  setup() {
    var seoData = config.CONTRIBUTION_SEO_META_DATA;
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
    GithubContribution,
    NewFooter,
  },
  beforeRouteEnter(to, from, next) {
    if (!config.SHOW_CONTRIBUTION_PAGE) {
      next({ name: "Error404Page", params: { pathMatch: "/contributions" } });
    } else {
      next();
    }
  },
};
</script>
