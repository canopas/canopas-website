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
      <WeeklyUpdateSection />
      <WhatsTrending />
      <Favourite />
      <AnimatedCreation />
      <UnitTest />
      <DevOps />
      <ExploreDesign />
    </div>
    <NewFooter />
  </div>
</template>

<script>
import { defineAsyncComponent } from "vue";
import Header from "@/components/partials/NewHeader.vue";
import { useMeta } from "vue-meta";
import config from "@/config.js";
import MobileLandingSection from "@/components/contributions/MobileLanding.vue";
import DesktopLandingSection from "@/components/contributions/DesktopLanding.vue";
import GithubContribution from "@/components/contributions/GithubContribution.vue";

const WeeklyUpdateSection = defineAsyncComponent(() =>
  import("@/components/contributions/WeeklyUpdate.vue")
);
const WhatsTrending = defineAsyncComponent(() =>
  import("@/components/contributions/WhatsTrending.vue")
);
const Favourite = defineAsyncComponent(() =>
  import("@/components/contributions/Favourite.vue")
);
const AnimatedCreation = defineAsyncComponent(() =>
  import("@/components/contributions/AnimatedCreation.vue")
);
const UnitTest = defineAsyncComponent(() =>
  import("@/components/contributions/UnitTest.vue")
);

const DevOps = defineAsyncComponent(() =>
  import("@/components/contributions/DevOps.vue")
);
const ExploreDesign = defineAsyncComponent(() =>
  import("@/components/contributions/DesignExplore.vue")
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
    WeeklyUpdateSection,
    WhatsTrending,
    Favourite,
    AnimatedCreation,
    UnitTest,
    DevOps,
    ExploreDesign,
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
