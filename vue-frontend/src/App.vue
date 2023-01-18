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
} from "@fortawesome/free-solid-svg-icons";
import Config from "@/config.js";
import { useMeta } from "vue-meta";

library.add(faArrowRight, faStar, faArrowLeft);

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
  @apply tw-font-product-sans tw-select-none;
}

body::-webkit-scrollbar,
.parallax::-webkit-scrollbar {
  @apply tw-hidden;
}

a {
  @apply tw-no-underline;
}

.canopas-gradient-text {
  @apply tw-bg-clip-text tw-text-transparent tw-from-[#f2709c] tw-to-[#ff9472] tw-bg-gradient-[89.88deg];
}

.v2-canopas-gradient-text {
  @apply tw-bg-clip-text tw-text-transparent tw-from-[#FF835B] tw-to-[#F2709C] tw-bg-gradient-[270.11deg];
}

.abs-fit-parent {
  @apply tw-absolute tw-w-full tw-h-full tw-top-0 tw-left-0;
}

.normal-text {
  @apply tw-not-italic tw-font-normal tw-text-[1.15rem] tw-leading-[1.5rem] tw-text-[#3d3d3dcc] tw-tracking-[0.3px] md:tw-text-[1.2rem] md:tw-leading-[1.375rem] lg:tw-text-[1.5rem] lg:tw-leading-[1.8rem];
}

.normal-2-text {
  @apply tw-not-italic tw-font-normal tw-text-[1.1rem] tw-leading-[1.3rem] tw-text-[#3d3d3dcc] lg:tw-text-[1.1875rem] lg:tw-leading-[1.4rem];
}

.header-text {
  @apply tw-font-bold tw-text-[1.5rem] tw-leading-[1.813rem] md:tw-text-[1.875rem] md:tw-leading-[2rem] lg:tw-text-[3rem] lg:tw-leading-[3.6rem];
}

.header-2-text {
  @apply tw-text-[#3d3d3d] tw-font-bold tw-text-[1.25rem] tw-leading-[1.6rem] tw-tracking-[0.5px] md:tw-text-[1.75rem] md:tw-leading-[2rem] lg:tw-text-[2.25rem] lg:tw-leading-[2.729rem];
}

.grecaptcha-badge {
  @apply tw-invisible;
}

/* 120px/119px - 90px/107px - 65px/94px */
.v2-header-text {
  @apply tw-font-bold tw-text-[4.0625rem] tw-leading-[5.875rem] tw-tracking-[1px] tw-text-[#3d3d3d] md:tw-text-[5.625rem] md:tw-leading-[6.6875rem] lg:tw-text-[7.5rem] lg:tw-leading-[7.4375rem];
}

/* 90px/107px - 65px/94px - 50px/80px */
.v2-header-2-text {
  @apply tw-font-semibold tw-text-[3.125rem] tw-leading-[5rem] tw-tracking-[1px] tw-text-[#3d3d3d] md:tw-text-[4.0625rem] md:tw-leading-[5.875rem] lg:tw-text-[5.625rem] lg:tw-leading-[6.6875rem];
}

/* 65px/94px - 50px/70px - 35px/50px */
.v2-header-3-text {
  @apply tw-font-semibold tw-text-[2.1875rem] tw-leading-[3.125rem] tw-tracking-[1px] tw-text-[#3d3d3d] md:tw-text-[3.125rem] md:tw-leading-[4.375rem] lg:tw-text-[4.0625rem] lg:tw-leading-[5.875rem];
}

/* 80px/95px - 60px/75px - 40px/55px */
.v2-title-text {
  @apply tw-font-medium tw-text-[2.5rem] tw-leading-[3.4375rem] tw-tracking-[1px] tw-text-[#3d3d3d] md:tw-text-[3.75rem] md:tw-leading-[4.6875rem] lg:tw-text-[5rem] lg:tw-leading-[5.9375rem];
}

/* 35px/68px - 30px/60px - 25px/45px */
.v2-title-2-text {
  @apply tw-font-medium tw-text-[1.5625rem] tw-leading-[2.8125rem] tw-tracking-[1px] tw-text-[#3d3d3d] md:tw-text-[1.875rem] md:tw-leading-[3.75rem] lg:tw-text-[2.1875rem] lg:tw-leading-[4.25rem];
}

/* 24px/29px - 20px/24px - 16px/20px */
.v2-title-3-text {
  @apply tw-font-medium tw-text-[1rem] tw-leading-[1.25rem] tw-tracking-[1px] tw-text-[#3d3d3d] md:tw-text-[1.25rem] md:tw-leading-[1.5rem] lg:tw-text-[1.5rem] lg:tw-leading-[1.8125rem];
}

/* 27px/43px - 22px/32px - 19px/30px */
.v2-normal-text {
  @apply tw-font-normal tw-text-[1.1875rem] tw-leading-[1.875rem] tw-tracking-[1px] tw-text-[#3d3d3d] md:tw-text-[1.375rem] md:tw-leading-[2rem] lg:tw-text-[1.6875rem] lg:tw-leading-[2.6875rem];
}

/* 22px/32px - 19px/30px - 17px/19px */
.v2-normal-2-text {
  @apply tw-font-normal tw-text-[1.0625rem] tw-leading-[1.1875rem] tw-tracking-[1px] tw-text-[#3d3d3d] md:tw-text-[1.1875rem] md:tw-leading-[1.875rem] lg:tw-text-[1.375rem] lg:tw-leading-[2rem];
}

/* 18px/30px - 17px/24px - 16px/18px */
.v2-normal-3-text {
  @apply tw-font-normal tw-text-[1rem] tw-leading-[1.125rem] tw-tracking-[1px] tw-text-[#3d3d3d] md:tw-text-[1.0625rem] md:tw-leading-[1.5rem] lg:tw-text-[1.125rem] lg:tw-leading-[1.875rem];
}

.border-text {
  -webkit-text-stroke: 1px #3d3d3d;
  color: #fff !important;
}

.underline-text {
  @apply tw-underline tw-decoration-solid tw-decoration-[#3d3d3dcc] tw-decoration-2 tw-underline-offset-[1rem] md:tw-decoration-[3px];
}

.swiper-pagination-bullet {
  @apply !tw-w-[15px] !tw-h-[15px] !tw-rounded-[3px] !tw-bg-transparent !tw-border-[1px] !tw-border-solid !tw-border-[#3d3d3d];
}

.swiper-pagination-bullet-active {
  @apply !tw-border-none !tw-from-[#F69259] !tw-to-[#F16975] !tw-bg-gradient-[45deg];
}

.v2-button {
  @apply tw-rounded-[3rem] tw-py-[0.7rem] tw-px-[1.5rem] tw-text-center tw-border-[1px] tw-border-solid tw-border-[#3d3d3d] tw-bg-white tw-text-[#3d3d3d] hover:tw-bg-[#3d3d3d] hover:tw-text-[#fff] active:tw-scale-[0.98];
}

.v2-button > span {
  @apply tw-mr-[10px];
}

.v2-button:hover > .fa,
.cta-button:hover > span,
.portfolio-nav:hover > span {
  @apply tw-text-[#fff];
}

.gradient-btn {
  @apply tw-rounded-[0.6rem] tw-p-[1rem] tw-mt-[1.25rem] tw-mr-auto tw-mb-0 tw-ml-auto tw-text-center tw-text-white tw-from-[#ff835b] tw-to-[#f2709c] tw-bg-gradient-[270.11deg] tw-border-[1px] tw-border-solid tw-border-transparent lg:tw-py-[1rem] lg:tw-px-[1rem] lg:tw-mt-[1.25rem] lg:tw-mr-[0.5rem] lg:tw-mb-0 lg:tw-ml-[0.5rem] active:tw-scale-[0.98] hover:tw-from-[#ff9472] hover:tw-via-[#ff909c] hover:tw-to-[#f2709c] hover:tw-bg-gradient-to-r hover:tw-shadow-[inset_2px_1000px_1px_#fff];
}

.gradient-btn > span {
  @apply tw-my-0 tw-mx-[6px] tw-font-bold;
}

.gradient-btn:hover > span,
.gradient-btn:hover > a > span {
  @apply tw-bg-clip-text tw-text-transparent tw-from-[#f2709c] tw-to-[#ff9472] tw-bg-gradient-[89.88deg];
}

.gradient-btn:hover > .fa,
.gradient-btn:hover > .fas,
.gradient-btn:hover > a > .fab {
  @apply tw-text-[#f2709c];
}

.gradient-border-btn {
  @apply tw-rounded-[0.6rem] tw-p-[1rem] tw-m-[5px] tw-text-center tw-border-[1px] tw-border-solid tw-border-transparent tw-from-[#ff9472] tw-via-[#ff909c] tw-to-[#f2709c] tw-bg-gradient-to-r tw-shadow-[inset_2px_1000px_1px_#fff] active:tw-scale-[0.98] hover:tw-from-[#f2709c] hover:tw-to-[#ff9472] hover:tw-bg-gradient-[91.53deg] hover:tw-border-[1px] hover:tw-border-solid hover:tw-border-transparent hover:tw-shadow-none;
}

.gradient-border-btn > a > .fab {
  @apply tw-text-[#f2709c];
}

.gradient-border-btn:hover > span,
.gradient-border-btn:hover > span > span,
.gradient-border-btn:hover > .fa,
.gradient-border-btn:hover > .arrow,
.gradient-border-btn:hover > a > .fab {
  @apply tw-text-white;
}
</style>
