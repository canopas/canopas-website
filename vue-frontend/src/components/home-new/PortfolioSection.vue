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
            ? 'tw-scale-105 md:tw-scale-[1.2] tw-z-[2] '
            : '',
        ]"
        class="tw-relative tw-transition-all tw-duration-500 tw-ease-in-out tw-transform tw-cursor-pointer"
        @mouseover="activeIndex = portfolio.id"
        @mouseleave="activeIndex = null"
        @touchstart.passive="activeIndex = portfolio.id"
        @click="openPortfolio(portfolio)"
      >
        <div
          class="tw-absolute tw-top-[18px] sm:tw-top-[13px] md:tw-top-5 xl:tw-top-[30px] tw-inset-x-0 tw-z-[2] tw-text-[1.375rem] tw-leading-8 lg:tw-text-[2rem] lg:tw-leading-8 xl:tw-text-[2.5rem] xl:tw-leading-8 tw-text-center tw-font-inter-bold"
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
            class="lozad tw-p-1 md:tw-p-8"
          >
            <source :data-src="portfolio.video[1]" type="video/webm" />
            <source :data-src="portfolio.video[0]" type="video/mp4" />
          </video>
          <img
            v-else
            :src="portfolio.images[0]"
            :srcset="`${portfolio.images[0]} 400w, ${portfolio.images[1]} 800w`"
            class="tw-h-full tw-w-full tw-object-cover tw-z-[1]"
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
          class="tw-absolute tw-inset-x-0 tw--mt-[50px] sm:-tw-mt-5 tw-z-[2] tw-text-[1rem] tw-leading-5 md:tw-text-[0.875rem] lg:tw-text-[1rem] md:tw-leading-4 lg:tw-leading-5 tw-text-center tw-font-inter-medium"
          :class="[
            activeIndex == portfolio.id
              ? 'tw-text-black-core/[0.87] tw-scale-[0.7] sm:tw-scale-[0.8]'
              : 'tw-text-transparent',
            portfolio.id == 7 ? 'lg:tw-px-2.5' : '',
            portfolio.id == 5 ? '!tw--mt-[50px]' : '',
            portfolio.video ? '-tw-mt-10 sm:!-tw-mt-5 lg:!tw--mt-[30px]' : '',
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

import tognessMp4 from "@/assets/images/portfolio/new-portfolio/togness.mp4";
import tognessWebm from "@/assets/images/portfolio/new-portfolio/togness.webm";

import luxeradio_1_400w from "@/assets/images/portfolio/new-portfolio/luxeradio-1-400w.webp";
import luxeradio_1_800w from "@/assets/images/portfolio/new-portfolio/luxeradio-1-800w.webp";

import luxeradio_2_400w from "@/assets/images/portfolio/new-portfolio/luxeradio-2-400w.webp";
import luxeradio_2_800w from "@/assets/images/portfolio/new-portfolio/luxeradio-2-800w.webp";

import smileMp4 from "@/assets/images/portfolio/new-portfolio/smile.mp4";
import smileWebm from "@/assets/images/portfolio/new-portfolio/smile.webm";

import justly_2_400w from "@/assets/images/portfolio/new-portfolio/justly-2-400w.webp";
import justly_2_800w from "@/assets/images/portfolio/new-portfolio/justly-2-800w.webp";

import togness_1_400w from "@/assets/images/portfolio/new-portfolio/togness-1-400w.webp";
import togness_1_800w from "@/assets/images/portfolio/new-portfolio/togness-1-800w.webp";

import justly_3_400w from "@/assets/images/portfolio/new-portfolio/justly-3-400w.webp";
import justly_3_800w from "@/assets/images/portfolio/new-portfolio/justly-3-800w.webp";

import config from "@/config.js";
import lozad from "lozad";

export default {
  data() {
    return {
      portfolios: [
        {
          id: 1,
          images: [justly_1_400w, justly_1_800w],
          classes: "tw-flex-[40%] tw-bg-[#F5E1B6]",
          title: "Justly",
          description:
            "Justly aims to tackle loneliness, depression, and mental health through innovative solutions.",
          url: "/portfolio/justly",
          event: "tap_justly_portfolio_card",
        },
        {
          id: 2,
          video: [tognessMp4, tognessWebm],
          classes: "tw-flex-[40%] tw-bg-[#B3D8FE]",
          title: "Togness",
          description:
            "Togness is a photo editor and slideshow maker app for your life’s most memorable events.",
          url: "/portfolio/togness",
          event: "tap_togness_portfolio_card",
        },
        {
          id: 3,
          images: [luxeradio_1_400w, luxeradio_1_800w],
          classes: "tw-flex-[20%] tw-hidden md:tw-block tw-bg-[#EBB7DB]",
          title: "Luxeradio",
          url: "/portfolio/luxeradio",
          event: "tap_luxeradio_portfolio_card",
        },
        {
          id: 4,
          images: [luxeradio_2_400w, luxeradio_2_800w],
          classes: "tw-flex-[50%] tw-bg-[#E9CFED]",
          title: "Luxeradio",
          description:
            "Luxe Radio displays the best of Moroccan and international creation, emphasizing taste, elegance, and refinement.",
          url: "/portfolio/luxeradio",
          event: "tap_luxeradio_portfolio_card",
        },
        {
          id: 5,
          video: [smileMp4, smileWebm],
          classes: "tw-flex-[50%] tw-bg-[#FED7CC]",
          title: "Smile+",
          description:
            "Smile+ app is designed for dentists to create a perfect smile for their patients automated with AI.",
          url: config.SMILEPLUS_URL,
          target: "_blank",
          event: "tap_smileplus_portfolio_card",
        },
        {
          id: 6,
          images: [justly_2_400w, justly_2_800w],
          classes: "tw-flex-[25%] tw-hidden md:tw-block tw-bg-[#CDE3F7]",
          title: "Justly",
          url: "/portfolio/justly",
          event: "tap_justly_portfolio_card",
        },
        {
          id: 7,
          images: [togness_1_400w, togness_1_800w],
          classes: "tw-flex-[50%] tw-hidden md:tw-block tw-bg-[#D6E1BF]",
          title: "Togness",
          description:
            "Togness is a photo editor and slideshow maker app for your life’s most memorable events.",
          url: "/portfolio/togness",
          event: "tap_togness_portfolio_card",
        },
        {
          id: 8,
          images: [justly_3_400w, justly_3_800w],
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
  mounted() {
    lozad().observe(); // lazy loads videos with default selector as '.lozad'
  },
  methods: {
    openPortfolio(portfolio) {
      window.open(portfolio.url, portfolio.target ? portfolio.target : "_self");
      this.mixpanel.track(portfolio.event);
    },
  },
};
</script>
