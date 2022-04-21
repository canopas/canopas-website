<template>
  <div>
    <HeaderV2 v-if="isNewHeader" />
    <ScreenHeader v-else />
    <LandingView v-on:scroll-to-career="scrollToCareer" />
    <VirtuesView />
    <LifeAtCanopas />
    <PerksAndBenifits v-on:scroll-to-career="scrollToCareer" />
    <WhyCanopas class="why-canopas-desktop" v-on:add-animation="handleScroll" />
    <WhyCanopasMobile
      class="why-canopas-mobile"
      v-on:add-animation="handleScroll"
    />
    <Career id="career" />
    <FaqSection v-on:add-animation="handleScroll" />
    <ScreenFooter2 />
  </div>
</template>

<script>
import ScreenHeader from "@/components/partials/ScreenHeader.vue";
import HeaderV2 from "@/components/partials/HeaderV2.vue";
import LandingView from "@/components/jobs/LandingView.vue";
import VirtuesView from "@/components/jobs/VirtuesView.vue";
import LifeAtCanopas from "@/components/jobs/LifeAtCanopas.vue";
import Career from "@/components/jobs/CareerView.vue";
import PerksAndBenifits from "@/components/jobs/PerksAndBenifits.vue";
import WhyCanopas from "@/components/jobs/WhyCanopas.vue";
import WhyCanopasMobile from "@/components/jobs/WhyCanopasMobile.vue";
import FaqSection from "@/components/jobs/FaqSection.vue";
import ScreenFooter2 from "@/components/partials/ScreenFooter2.vue";
import config from "@/config.js";
import { library } from "@fortawesome/fontawesome-svg-core";
import { useMeta } from "vue-meta";

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
    ScreenHeader,
    HeaderV2,
    LandingView,
    VirtuesView,
    LifeAtCanopas,
    Career,
    PerksAndBenifits,
    WhyCanopas,
    WhyCanopasMobile,
    FaqSection,
    ScreenFooter2,
  },
  methods: {
    scrollToCareer() {
      var careerDiv = document.getElementById("career");
      var top = careerDiv.offsetTop;
      window.scrollTo({ top: top, behavior: "smooth" });
    },
    handleScroll(data) {
      if (data) {
        let { top, bottom } = data.name.getBoundingClientRect();
        let height = document.documentElement.clientHeight;

        if (top < height && bottom > 0) {
          if (data.childRef.length > 0) {
            for (let i = 0; i < data.childRef.length; i++) {
              data.childRef[i].name.classList.add(data.childRef[i].animation);
            }
          } else {
            data.name.classList.add(data.animation);
          }
        }
      }
    },
  },
  data() {
    return {
      isNewHeader: config.IS_NEW_HEADER,
    };
  },
  mounted() {
    this.$gtag.event("view_page_jobs");
  },
};
</script>

<style lang="scss" scoped>
.why-canopas-mobile {
  display: block;
}
.why-canopas-desktop {
  display: none;
}

@include media-breakpoint-up(md) {
  .why-canopas-mobile {
    display: none;
  }
  .why-canopas-desktop {
    display: block;
  }
}
</style>
