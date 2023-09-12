<template>
  <div>
    <router-view></router-view>
  </div>
</template>

<script>
import { library } from "@fortawesome/fontawesome-svg-core";
import {
  faArrowRight,
  faArrowLeft,
  faStar,
  faThumbsUp,
  faTimes,
  faAngleUp,
  faAngleDown,
  faAngleRight,
  faAngleLeft,
} from "@fortawesome/free-solid-svg-icons";
import { faMediumM, faGithub } from "@fortawesome/free-brands-svg-icons";
import Config from "@/config.js";
import { useMeta } from "vue-meta";
library.add(
  faTimes,
  faArrowRight,
  faStar,
  faArrowLeft,
  faThumbsUp,
  faAngleUp,
  faAngleDown,
  faAngleRight,
  faAngleLeft,
  faMediumM,
  faGithub,
);
export default {
  name: "App",
  components: {},
  setup() {
    useMeta({
      meta: [
        {
          name: "robots",
          content: Config.IS_PROD ? "" : "noindex",
          vmid: "robots",
        },
      ],
    });
  },
};
</script>

<style lang="postcss">
body {
  @apply tw-select-none tw-font-product-sans;
}
.hidden-scrollbar::-webkit-scrollbar,
.parallax::-webkit-scrollbar,
body::-webkit-scrollbar {
  @apply tw-hidden;
}
a {
  @apply tw-no-underline;
}
.canopas-gradient-text {
  @apply tw-from-[#f2709c] tw-to-[#ff9472] tw-bg-clip-text tw-text-transparent tw-bg-gradient-[89.88deg];
}
.v2-canopas-gradient-text {
  @apply tw-from-[#FF835B] tw-to-[#F2709C] tw-bg-clip-text tw-text-transparent tw-bg-gradient-[270.11deg];
}
.normal-text {
  @apply tw-text-[1.15rem] tw-font-normal tw-not-italic tw-leading-6 tw-tracking-[0.3px] tw-text-[#3d3d3dcc] md:tw-text-[1.2rem] md:tw-leading-[1.375rem] lg:tw-text-[1.5rem] lg:tw-leading-[1.8rem];
}
.normal-2-text {
  @apply tw-text-[1.1rem] tw-font-normal tw-not-italic tw-leading-[1.3rem] tw-text-[#3d3d3dcc] lg:tw-text-[1.1875rem] lg:tw-leading-[1.4rem];
}
.header-text {
  @apply tw-text-[1.5rem] tw-font-bold tw-leading-[1.813rem] md:tw-text-[1.875rem] md:tw-leading-8 lg:tw-text-[3rem] lg:tw-leading-[3.6rem];
}
.header-2-text {
  @apply tw-text-[1.25rem] tw-font-bold tw-leading-[1.6rem] tw-tracking-[0.5px] tw-text-[#3d3d3d] md:tw-text-[1.75rem] md:tw-leading-8 lg:tw-text-[2.25rem] lg:tw-leading-[2.729rem];
}
.grecaptcha-badge {
  @apply tw-invisible;
}
.v2-header-text {
  @apply tw-text-[4.0625rem] tw-font-bold tw-leading-[5.875rem] tw-tracking-[1px] tw-text-[#3d3d3d] md:tw-text-[5.625rem] md:tw-leading-[6.6875rem] lg:tw-text-[7.5rem] lg:tw-leading-[7.4375rem];
}
.v2-header-2-text {
  @apply tw-text-[3.125rem] tw-font-semibold tw-leading-[5rem] tw-tracking-[1px] tw-text-[#3d3d3d] md:tw-text-[4.0625rem] md:tw-leading-[5.875rem] lg:tw-text-[5.625rem] lg:tw-leading-[6.6875rem];
}
.v2-header-3-text {
  @apply tw-text-[2.1875rem] tw-font-semibold tw-leading-[3.125rem] tw-tracking-[1px] tw-text-[#3d3d3d] md:tw-text-[3.125rem] md:tw-leading-[4.375rem] lg:tw-text-[4.0625rem] lg:tw-leading-[5.875rem];
}
.v2-title-text {
  @apply tw-text-[2.5rem] tw-font-medium tw-leading-[3.4375rem] tw-tracking-[1px] tw-text-[#3d3d3d] md:tw-text-[3.75rem] md:tw-leading-[4.6875rem] lg:tw-text-[5rem] lg:tw-leading-[5.9375rem];
}
.v2-title-2-text {
  @apply tw-text-[1.5625rem] tw-font-medium tw-leading-[2.8125rem] tw-tracking-[1px] tw-text-[#3d3d3d] md:tw-text-[1.875rem] md:tw-leading-[3.75rem] lg:tw-text-[2.1875rem] lg:tw-leading-[4.25rem];
}
.v2-title-3-text {
  @apply tw-text-[1rem] tw-font-medium tw-leading-[1.25rem] tw-tracking-[1px] tw-text-[#3d3d3d] md:tw-text-[1.25rem] md:tw-leading-6 lg:tw-text-[1.5rem] lg:tw-leading-[1.8125rem];
}
.v2-normal-text {
  @apply tw-text-[1.1875rem] tw-font-normal tw-leading-[1.875rem] tw-tracking-[1px] tw-text-[#3d3d3d] md:tw-text-[1.375rem] md:tw-leading-8 lg:tw-text-[1.6875rem] lg:tw-leading-[2.6875rem];
}
.v2-normal-2-text {
  @apply tw-text-[1.0625rem] tw-font-normal tw-leading-[1.1875rem] tw-tracking-[1px] tw-text-[#3d3d3d] md:tw-text-[1.1875rem] md:tw-leading-[1.875rem] lg:tw-text-[1.375rem] lg:tw-leading-8;
}
.v2-normal-3-text {
  @apply tw-text-[1rem] tw-font-normal tw-leading-[1.125rem] tw-tracking-[1px] tw-text-[#3d3d3d] md:tw-text-[1.0625rem] md:tw-leading-6 lg:tw-text-[1.125rem] lg:tw-leading-[1.875rem];
}
.border-text {
  -webkit-text-stroke: 1px #3d3d3d;
  color: #fff !important;
}
.underline-text {
  @apply tw-underline tw-decoration-[#3d3d3dcc] tw-decoration-solid tw-decoration-2 tw-underline-offset-[1rem] md:tw-decoration-[3px];
}
.v2-button {
  @apply tw-rounded-[3rem] tw-border-[1px] tw-border-solid tw-border-[#3d3d3d] tw-bg-white tw-px-[1.5rem] tw-py-[0.7rem] tw-text-center tw-text-[#3d3d3d] hover:tw-bg-[#3d3d3d] hover:tw-text-[#fff] active:tw-scale-[0.98];
}
.v2-button > span {
  @apply tw-mr-2.5;
}
.cta-button:hover > span,
.portfolio-nav:hover > span,
.v2-button:hover > .fa {
  @apply tw-text-[#fff];
}
.gradient-btn {
  @apply tw-mb-0 tw-ml-auto tw-mr-auto tw-mt-5 tw-rounded-[0.6rem] tw-border tw-border-solid tw-border-transparent tw-from-[#ff835b] tw-to-[#f2709c] tw-p-4 tw-text-center tw-text-white tw-bg-gradient-[270.11deg] hover:tw-bg-gradient-to-r hover:tw-from-[#ff9472] hover:tw-via-[#ff909c] hover:tw-to-[#f2709c] hover:tw-shadow-[inset_2px_1000px_1px_#fff] active:tw-scale-[0.98] lg:tw-mb-0 lg:tw-ml-2 lg:tw-mr-2 lg:tw-mt-5 lg:tw-px-4 lg:tw-py-4;
}
.gradient-btn > span {
  @apply tw-mx-1.5 tw-my-0 tw-font-bold;
}
.gradient-btn:hover > a > span,
.gradient-btn:hover > span {
  @apply tw-from-[#f2709c] tw-to-[#ff9472] tw-bg-clip-text tw-text-transparent tw-bg-gradient-[89.88deg];
}
.white-btn > span {
  @apply tw-from-[#FF835B] tw-to-[#F2709C] tw-bg-clip-text tw-text-transparent tw-bg-gradient-[270.11deg];
}
.gradient-btn:hover > .fa,
.gradient-btn:hover > .fas,
.gradient-btn:hover > a > .fab {
  @apply tw-text-[#f2709c];
}
.gradient-border-btn {
  @apply tw-m-[5px] tw-rounded-[0.6rem] tw-border tw-border-solid tw-border-transparent tw-bg-gradient-to-r tw-from-[#ff9472] tw-via-[#ff909c] tw-to-[#f2709c] tw-p-[1rem] tw-text-center tw-shadow-[inset_2px_1000px_1px_#fff] hover:tw-border hover:tw-border-solid hover:tw-border-transparent hover:tw-from-[#f2709c] hover:tw-to-[#ff9472] hover:tw-shadow-none hover:tw-bg-gradient-[91.53deg] active:tw-scale-[0.98];
}
.gradient-border-btn > a > .fab {
  @apply tw-text-[#f2709c];
}
.gradient-border-btn:hover > .arrow,
.gradient-border-btn:hover > .fa,
.gradient-border-btn:hover > a > .fab,
.gradient-border-btn:hover > span,
.gradient-border-btn:hover > span > span {
  @apply tw-text-white;
}
</style>
