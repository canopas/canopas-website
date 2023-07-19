<template>
  <div>
    <metainfo>
      <template v-slot:title="{ content }">
        {{ content }}
      </template>
    </metainfo>
    <Header />
    <div class="tw-overflow-hidden">
      <MobileLandingSection ref="landing" class="tw-block md:tw-hidden" />
      <DesktopLandingSection ref="landing" class="tw-hidden md:tw-block" />
      <GithubContribution ref="contribution" />
      <WeeklyUpdateSection ref="weeklyupdate" />
      <WhatsTrending ref="whatstrending" />
      <Favourite ref="favourite" />
      <AnimatedCreation ref="animationcreation" />
      <UnitTest ref="unittest" />
      <DevOps ref="devops" />
      <ExploreDesign ref="exploredesign" />
    </div>
    <NewFooter ref="footer" />
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
import { elementInViewPort } from "@/utils.js";
const WeeklyUpdateSection = defineAsyncComponent(() =>
  import("@/components/contributions/WeeklyUpdate.vue"),
);
const WhatsTrending = defineAsyncComponent(() =>
  import("@/components/contributions/WhatsTrending.vue"),
);
const Favourite = defineAsyncComponent(() =>
  import("@/components/contributions/Favourite.vue"),
);
const AnimatedCreation = defineAsyncComponent(() =>
  import("@/components/contributions/AnimatedCreation.vue"),
);
const UnitTest = defineAsyncComponent(() =>
  import("@/components/contributions/UnitTest.vue"),
);

const DevOps = defineAsyncComponent(() =>
  import("@/components/contributions/DevOps.vue"),
);
const ExploreDesign = defineAsyncComponent(() =>
  import("@/components/contributions/DesignExplore.vue"),
);

const NewFooter = defineAsyncComponent(() =>
  import("@/components/partials/NewFooter.vue"),
);

export default {
  setup() {
    var seoData = config.CONTRIBUTION_SEO_META_DATA;
    useMeta({
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
  data() {
    return {
      event: "",
      events: {
        landing: "view_contribution_landing_section",
        contribution: "view_contribution_section",
        weeklyupdate: "view_ weekly_update_section",
        whatstrending: "view_whats_trending_section",
        favourite: "view_favourite_contribution",
        animationcreation: "view_animation_creation",
        unittest: "view_ unit_test",
        devops: "view_devops_section",
        exploredesign: "view_explore_design_section",
        footer: "view_contribution_footer",
      },
    };
  },
  inject: ["mixpanel"],
  mounted() {
    window.addEventListener("scroll", this.sendEvent);
    this.mixpanel.track("view_contribution_landing_section");
  },
  unmounted() {
    window.removeEventListener("scroll", this.sendEvent);
  },
  methods: {
    sendEvent() {
      const event = this.events[elementInViewPort(this.$refs)];
      if (event && this.event !== event) {
        this.event = event;
        this.mixpanel.track(event);
      }
    },
  },
};
</script>
