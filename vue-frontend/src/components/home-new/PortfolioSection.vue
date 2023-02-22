<template>
  <div
    class="tw-container tw-flex md:tw-items-center tw-m-auto sm:tw-w-[70%] md:tw-w-full tw-pb-16 md:tw-pb-32"
  >
    <div class="tw-flex tw-flex-col md:tw-flex-wrap md:tw-flex-row">
      <div
        v-for="portfolio in portfolios"
        :key="portfolio"
        :class="[
          portfolio.classes,
          portfolio.id != 8 ? 'tw-p-4 sm:tw-p-8' : '',
          portfolio.id == 4 ? 'sm:!tw-pb-12' : '',
          activeIndex == portfolio.id
            ? 'tw-scale-[1.05] md:tw-scale-[1.2] tw-z-[2] '
            : '',
        ]"
        class="tw-relative tw-transition-all tw-duration-500 tw-ease-in-out tw-transform"
        @mouseover="activeIndex = portfolio.id"
        @mouseleave="activeIndex = null"
        @touchstart.passive="activeIndex = portfolio.id"
        @click="openPortfolio(portfolio)"
      >
        <div
          class="tw-absolute tw-top-[18px] sm:tw-top-[13px] md:tw-top-[20px] xl:tw-top-[30px] tw-inset-x-0 tw-z-[2] tw-text-[1.375rem] tw-leading-[2rem] lg:tw-text-[2rem] lg:tw-leading-[2rem] xl:tw-text-[2.5rem] xl:tw-leading-[2rem] tw-text-center tw-font-inter-bold"
          :class="[
            activeIndex == portfolio.id
              ? 'tw-text-black-core/[0.87] tw-scale-[0.7] sm:tw-scale-[0.8]'
              : 'tw-text-transparent',
          ]"
        >
          {{ portfolio.title }}
        </div>
        <div class="tw-flex tw-items-center tw-h-full tw-py-8 sm:tw-p-0">
          <video
            v-if="portfolio.video"
            autoplay
            loop
            muted
            playsinline
            class="tw-p-1 md:tw-p-8"
          >
            <source :src="portfolio.video" type="video/mp4" />
          </video>
          <img
            v-else
            :src="portfolio.images[3]"
            :srcset="`${portfolio.images[0]} 400w, ${portfolio.images[1]} 800w, ${portfolio.images[2]} 1600w, ${portfolio.images[3]} 2400w`"
            class="tw-h-full tw-object-cover tw-z-[1]"
            :class="[
              portfolio.id == 4
                ? 'tw-mt-0 sm:tw-mt-4 md:tw-mt-0 sm:tw-w-full sm:tw-h-[90%]'
                : '',
            ]"
            :alt="portfolio.title + `-image`"
            loading="lazy"
          />
        </div>
        <div
          v-if="portfolio.description"
          class="tw-absolute tw-inset-x-0 tw--mt-[50px] sm:tw--mt-[20px] tw-z-[2] tw-text-[1rem] tw-leading-[1.25rem] md:tw-text-[0.875rem] lg:tw-text-[1rem] md:tw-leading-[1rem] lg:tw-leading-[1.25rem] tw-text-center tw-font-inter-medium"
          :class="[
            activeIndex == portfolio.id
              ? 'tw-text-black-core/[0.87] tw-scale-[0.7] sm:tw-scale-[0.8]'
              : 'tw-text-transparent',
            portfolio.id == 7 ? 'lg:tw-px-2.5' : '',
            portfolio.id == 5 ? '!tw--mt-[50px]' : '',
            portfolio.video
              ? 'tw--mt-[40px] sm:!tw--mt-[20px] lg:!tw--mt-[30px]'
              : '',
          ]"
        >
          {{ portfolio.description }}
        </div>
      </div>
    </div>
  </div>
</template>

<script type="module">
import justly_1_400w from "@/assets/images/portfolio/new-portfolio/justly-1-400w.webp";
import justly_1_800w from "@/assets/images/portfolio/new-portfolio/justly-1-800w.webp";
import justly_1_1600w from "@/assets/images/portfolio/new-portfolio/justly-1-1600w.webp";
import justly_1_2400w from "@/assets/images/portfolio/new-portfolio/justly-1-2400w.webp";

import togness from "@/assets/images/portfolio/new-portfolio/togness.mp4";

import luxeradio_1_400w from "@/assets/images/portfolio/new-portfolio/luxeradio-1-400w.webp";
import luxeradio_1_800w from "@/assets/images/portfolio/new-portfolio/luxeradio-1-800w.webp";
import luxeradio_1_1600w from "@/assets/images/portfolio/new-portfolio/luxeradio-1-1600w.webp";
import luxeradio_1_2400w from "@/assets/images/portfolio/new-portfolio/luxeradio-1-2400w.webp";

import luxeradio_2_400w from "@/assets/images/portfolio/new-portfolio/luxeradio-2-400w.webp";
import luxeradio_2_800w from "@/assets/images/portfolio/new-portfolio/luxeradio-2-800w.webp";
import luxeradio_2_1600w from "@/assets/images/portfolio/new-portfolio/luxeradio-2-1600w.webp";
import luxeradio_2_2400w from "@/assets/images/portfolio/new-portfolio/luxeradio-2-2400w.webp";

import smile from "@/assets/images/portfolio/new-portfolio/smile.mp4";

import justly_2_400w from "@/assets/images/portfolio/new-portfolio/justly-2-400w.webp";
import justly_2_800w from "@/assets/images/portfolio/new-portfolio/justly-2-800w.webp";
import justly_2_1600w from "@/assets/images/portfolio/new-portfolio/justly-2-1600w.webp";
import justly_2_2400w from "@/assets/images/portfolio/new-portfolio/justly-2-2400w.webp";

import togness_1_400w from "@/assets/images/portfolio/new-portfolio/togness-1-400w.webp";
import togness_1_800w from "@/assets/images/portfolio/new-portfolio/togness-1-800w.webp";
import togness_1_1600w from "@/assets/images/portfolio/new-portfolio/togness-1-1600w.webp";
import togness_1_2400w from "@/assets/images/portfolio/new-portfolio/togness-1-2400w.webp";

import justly_3_400w from "@/assets/images/portfolio/new-portfolio/justly-3-400w.webp";
import justly_3_800w from "@/assets/images/portfolio/new-portfolio/justly-3-800w.webp";
import justly_3_1600w from "@/assets/images/portfolio/new-portfolio/justly-3-1600w.webp";
import justly_3_2400w from "@/assets/images/portfolio/new-portfolio/justly-3-2400w.webp";

import config from "@/config.js";

export default {
  data() {
    return {
      portfolios: [
        {
          id: 1,
          images: [
            justly_1_400w,
            justly_1_800w,
            justly_1_1600w,
            justly_1_2400w,
          ],
          classes: "tw-flex-[40%] tw-bg-[#F5E1B6]",
          title: "Justly",
          description:
            "Justly aims to tackle loneliness, depression, and mental health through innovative solutions.",
          url: "/portfolio/justly",
          event: "tap_justly_portfolio_card",
        },
        {
          id: 2,
          video: togness,
          classes: "tw-flex-[40%] tw-bg-[#B3D8FE]",
          title: "Togness",
          description:
            "Togness is a photo editor and slideshow maker app for your life’s most memorable events.",
          url: "/portfolio/togness",
          event: "tap_togness_portfolio_card",
        },
        {
          id: 3,
          images: [
            luxeradio_1_400w,
            luxeradio_1_800w,
            luxeradio_1_1600w,
            luxeradio_1_2400w,
          ],
          classes: "tw-flex-[20%] tw-hidden md:tw-block tw-bg-[#EBB7DB]",
          title: "Luxeradio",
          url: "/portfolio/luxeradio",
          event: "tap_luxeradio_portfolio_card",
        },
        {
          id: 4,
          images: [
            luxeradio_2_400w,
            luxeradio_2_800w,
            luxeradio_2_1600w,
            luxeradio_2_2400w,
          ],
          classes: "tw-flex-[50%] tw-bg-[#E9CFED]",
          title: "Luxeradio",
          description:
            "Luxe Radio displays the best of Moroccan and international creation, emphasizing taste, elegance, and refinement.",
          url: "/portfolio/luxeradio",
          event: "tap_luxeradio_portfolio_card",
        },
        {
          id: 5,
          video: smile,
          classes: "tw-flex-[50%] tw-bg-[#FED7CC]",
          title: "Smile+",
          description:
            "Smile+ app is designed for dentists to create a perfect smile for their patients automated with AI.",
          url: config.SMILEPLUS_URL,
          event: "tap_smileplus_portfolio_card",
        },
        {
          id: 6,
          images: [
            justly_2_400w,
            justly_2_800w,
            justly_2_1600w,
            justly_2_2400w,
          ],
          classes: "tw-flex-[25%] tw-hidden md:tw-block tw-bg-[#CDE3F7]",
          title: "Justly",
          url: "/portfolio/justly",
          event: "tap_justly_portfolio_card",
        },
        {
          id: 7,
          images: [
            togness_1_400w,
            togness_1_800w,
            togness_1_1600w,
            togness_1_2400w,
          ],
          classes: "tw-flex-[50%] tw-hidden md:tw-block tw-bg-[#D6E1BF]",
          title: "Togness",
          description:
            "Togness is a photo editor and slideshow maker app for your life’s most memorable events.",
          url: "/portfolio/togness",
          event: "tap_togness_portfolio_card",
        },
        {
          id: 8,
          images: [
            justly_3_400w,
            justly_3_800w,
            justly_3_1600w,
            justly_3_2400w,
          ],
          classes: "tw-flex-[25%] tw-hidden md:tw-block tw-bg-[#AED2F6]",
          title: "Justly",
          url: "/portfolio/justly",
          event: "tap_justly_portfolio_card",
        },
      ],
      activeIndex: null,
    };
  },
  inject: ["mixpanel"],
  methods: {
    openPortfolio(portfolio) {
      window.open(portfolio.url, "_self");
      this.mixpanel.track(portfolio.event);
    },
  },
};
</script>
