<template>
  <div>
    <metainfo>
      <template v-slot:title="{ content }">
        {{ content }}
      </template>
    </metainfo>
    <Header />

    <div>
      <Services />
    </div>

    <NewFooter />
  </div>
</template>

<script>
import Header from "@/components/partials/NewHeader.vue";
import Services from "@/components/mobile-app-development/Services.vue";
import config from "@/config.js";
import { useMeta } from "vue-meta";
import { defineAsyncComponent } from "vue";

const NewFooter = defineAsyncComponent(() =>
  import("@/components/partials/NewFooter.vue")
);

export default {
  setup() {
    var seoData = config.MOBILE_APP_DEVELOPMENT_SEO_META_DATA;
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
    Services,
    NewFooter,
  },
  beforeRouteEnter(to, from, next) {
    if (!config.SHOW_MOBILE_APP_DEVELOPMENT_PAGE) {
      next({
        name: "Error404Page",
        params: { pathMatch: "/mobile-app-development" },
      });
    } else {
      next();
    }
  },
};
</script>
