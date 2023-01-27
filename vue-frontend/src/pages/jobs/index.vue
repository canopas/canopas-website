<template>
  <div>
    <metainfo>
      <template v-slot:title="{ content }">
        {{ content }}
      </template>
    </metainfo>
    <ScreenHeaderV2 v-if="!isShowNewHomePage" />
    <NewHeader v-else />
    <LandingView v-on:scroll-to-career="scrollToCareer" ref="landingview" />
    <VirtuesView ref="virtue" />
    <LifeAtCanopas ref="life" />
    <PerksAndBenefits v-on:scroll-to-career="scrollToCareer" ref="perks" />
    <WhyCanopas
      class="tw-hidden md:tw-block"
      v-on:add-animation="handleAnimationOnScroll"
      ref="whycanopas"
    />
    <WhyCanopasMobile
      class="tw-block md:tw-hidden"
      v-on:add-animation="handleAnimationOnScroll"
      ref="whycanopas"
    />
    <Career id="career" ref="joblist" />
    <FaqSection v-on:add-animation="handleAnimationOnScroll" />
    <ScreenFooter2 />
  </div>
</template>

<script>
import ScreenHeaderV2 from "@/components/partials/ScreenHeaderV2.vue";
import NewHeader from "@/components/partials/NewHeader.vue";
import LandingView from "@/components/jobs/LandingView.vue";
import VirtuesView from "@/components/jobs/VirtuesView.vue";
import LifeAtCanopas from "@/components/jobs/LifeAtCanopas.vue";
import Career from "@/components/jobs/CareerView.vue";
import PerksAndBenefits from "@/components/jobs/PerksAndBenefits.vue";
import WhyCanopas from "@/components/jobs/WhyCanopas.vue";
import WhyCanopasMobile from "@/components/jobs/WhyCanopasMobile.vue";
import FaqSection from "@/components/jobs/FaqSection.vue";
import ScreenFooter2 from "@/components/partials/ScreenFooter2.vue";
import config from "@/config.js";
import { library } from "@fortawesome/fontawesome-svg-core";
import { useMeta } from "vue-meta";
import { analyticsEvent, handleAnimationOnScroll } from "@/utils.js";

import {
  faAlignLeft,
  faCheckCircle,
  faMinus,
  faPlus,
} from "@fortawesome/free-solid-svg-icons";

library.add(faAlignLeft, faCheckCircle, faMinus, faPlus);

export default {
  setup() {
    var seoData = config.JOBS_SEO_META_DATA;
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
    NewHeader,
    LandingView,
    VirtuesView,
    LifeAtCanopas,
    Career,
    PerksAndBenefits,
    WhyCanopas,
    WhyCanopasMobile,
    FaqSection,
    ScreenFooter2,
  },
  data() {
    return {
      handleAnimationOnScroll,
      event: "",
      isShowNewHomePage: config.IS_SHOW_NEW_HOME_PAGE,
    };
  },
  mounted() {
    this.$gtag.event("view_jobs_page");
    window.addEventListener("scroll", this.sendEvent);
  },
  unmounted() {
    window.removeEventListener("scroll", this.sendEvent);
  },
  methods: {
    sendEvent() {
      const event = analyticsEvent(this.$refs);
      if (event && this.event !== event) {
        this.event = event;
        this.$gtag.event(event);
      }
    },
    scrollToCareer() {
      var careerDiv = document.getElementById("career");
      var top = careerDiv.offsetTop;
      window.scrollTo({ top: top, behavior: "smooth" });
    },
  },
};
</script>
