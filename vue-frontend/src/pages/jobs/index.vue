<template>
  <div>
    <metainfo>
      <template v-slot:title="{ content }">
        {{ content }}
      </template>
    </metainfo>
    <Header />
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
    <NewFooter />
  </div>
</template>

<script>
import Header from "@/components/partials/NewHeader.vue";
import LandingView from "@/components/jobs/LandingView.vue";
import { defineAsyncComponent } from "vue";
import config from "@/config.js";
import { library } from "@fortawesome/fontawesome-svg-core";
import { useMeta } from "vue-meta";
import { elementInViewPort, handleAnimationOnScroll } from "@/utils.js";
import {
  faAlignLeft,
  faCheckCircle,
  faMinus,
  faPlus,
} from "@fortawesome/free-solid-svg-icons";

const VirtuesView = defineAsyncComponent(() =>
  import("@/components/jobs/VirtuesView.vue")
);
const LifeAtCanopas = defineAsyncComponent(() =>
  import("@/components/jobs/LifeAtCanopas.vue")
);
const Career = defineAsyncComponent(() =>
  import("@/components/jobs/CareerView.vue")
);
const PerksAndBenefits = defineAsyncComponent(() =>
  import("@/components/jobs/PerksAndBenefits.vue")
);
const WhyCanopas = defineAsyncComponent(() =>
  import("@/components/jobs/WhyCanopas.vue")
);
const WhyCanopasMobile = defineAsyncComponent(() =>
  import("@/components/jobs/WhyCanopasMobile.vue")
);
const FaqSection = defineAsyncComponent(() =>
  import("@/components/jobs/FaqSection.vue")
);
const NewFooter = defineAsyncComponent(() =>
  import("@/components/partials/NewFooter.vue")
);

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
        image: seoData.image,
      },
    });
  },
  components: {
    Header,
    LandingView,
    VirtuesView,
    LifeAtCanopas,
    Career,
    PerksAndBenefits,
    WhyCanopas,
    WhyCanopasMobile,
    FaqSection,
    NewFooter,
  },
  data() {
    return {
      handleAnimationOnScroll,
      event: "",
      events: {
        landingview: "view_jobs_landing_section",
        virtue: "view_virtues_section",
        life: "view_life_at_canopas_section",
        perks: "view_perks_benefits",
        whycanopas: "view_why_canopas",
        joblist: "view_jobs_list",
      },
    };
  },
  inject: ["mixpanel"],
  mounted() {
    window.addEventListener("scroll", this.sendEvent);
    this.mixpanel.track("view_jobs_page");
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
    scrollToCareer() {
      var careerDiv = document.getElementById("career");
      var top = careerDiv.offsetTop;
      window.scrollTo({ top: top, behavior: "smooth" });
    },
  },
};
</script>
