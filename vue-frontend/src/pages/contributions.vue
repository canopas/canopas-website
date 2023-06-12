<template>
  <div>
    <metainfo>
      <template v-slot:title="{ content }">
        {{ content }}
      </template>
    </metainfo>
    <Header />
    <div class="tw-overflow-hidden">
      <MobileLandingSection class="tw-block md:tw-hidden" />
      <DesktopLandingSection class="tw-hidden md:tw-block" />
      <GithubContribution />
      <Favourite />
      <AnimatedCreation />
    </div>
    <NewFooter />
  </div>
</template>

<script>
import Header from "@/components/partials/NewHeader.vue";
import { useMeta } from "vue-meta";
import MobileLandingSection from "@/components/contribution/MobileLanding.vue";
import DesktopLandingSection from "@/components/contribution/DesktopLanding.vue";
import config from "@/config.js";

import { defineAsyncComponent } from "vue";

const GithubContribution = defineAsyncComponent(() =>
  import("@/components/contribution/GithubContribution.vue")
);
const Favourite = defineAsyncComponent(() =>
  import("@/components/contribution/Favourite.vue")
);
const AnimatedCreation = defineAsyncComponent(() =>
  import("@/components/contribution/AnimatedCreation.vue")
);
const NewFooter = defineAsyncComponent(() =>
  import("@/components/partials/NewFooter.vue")
);
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
    MobileLandingSection,
    DesktopLandingSection,
    GithubContribution,
    Favourite,
    AnimatedCreation,
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
