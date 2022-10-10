<template>
  <div>
    <metainfo>
      <template v-slot:title="{ content }">
        {{ content }}
      </template>
    </metainfo>
    <ScreenHeaderV2 />
    <LandingView v-on:scroll-to-career="scrollToCareer" ref="landingview" />
    <VirtuesView ref="virtue" />
    <LifeAtCanopas />
    <PerksAndBenifits v-on:scroll-to-career="scrollToCareer" ref="perks" />
    <WhyCanopas
      class="tw-hidden md:tw-block"
      v-on:add-animation="handleScroll"
      ref="whycanopas"
    />
    <WhyCanopasMobile
      class="tw-block md:tw-hidden"
      v-on:add-animation="handleScroll"
      ref="whycanopas"
    />
    <Career id="career" ref="joblist" />
    <FaqSection v-on:add-animation="handleScroll" />
    <ScreenFooter2 />
  </div>
</template>

<script>
import ScreenHeaderV2 from "@/components/partials/ScreenHeaderV2.vue";
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
import { analyticsEvent } from "@/utils.js";

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
  data() {
    return {
      event: "",
    };
  },
  mounted() {
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
};
</script>
