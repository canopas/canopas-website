<template>
  <div
    class="container flex md:items-center m-auto sm:w-[70%] md:w-full pb-16 md:pb-32"
  >
    <div class="flex flex-col md:flex-wrap md:flex-row">
      <div
        v-for="portfolio in portfolios"
        :key="portfolio"
        :class="[
          portfolio.classes,
          portfolio.id != 8 ? 'p-4 sm:p-8' : '',
          portfolio.id == 4 ? 'sm:!pb-12' : '',
          activeIndex == portfolio.id ? 'scale-105 md:scale-[1.2] z-[2] ' : '',
        ]"
        class="relative transition-all duration-500 ease-in-out transform cursor-pointer"
        @mouseover="activeIndex = portfolio.id"
        @mouseleave="activeIndex = null"
        @touchstart.passive="activeIndex = portfolio.id"
        @click="openPortfolio(portfolio)"
      >
        <div
          class="absolute top-[18px] sm:top-[13px] md:top-5 xl:top-[30px] inset-x-0 z-[2] text-[1.375rem] leading-8 lg:text-[2rem] lg:leading-8 xl:text-[2.5rem] xl:leading-8 text-center font-inter-bold"
          :class="[
            activeIndex == portfolio.id
              ? 'text-black-core/[0.87] scale-[0.7] sm:scale-[0.8]'
              : 'text-transparent',
          ]"
        >
          {{ portfolio.title }}
        </div>
        <div class="flex items-center h-full py-8 sm:p-0">
          <video
            v-if="portfolio.video"
            autoplay
            loop
            muted
            playsinline
            class="lozad p-1 md:p-8"
          >
            <source :data-src="portfolio.video[1]" type="video/webm" />
            <source :data-src="portfolio.video[0]" type="video/mp4" />
          </video>
          <img
            v-else
            :src="portfolio.images[0]"
            :srcset="`${portfolio.images[0]} 400w, ${portfolio.images[1]} 800w`"
            class="h-full w-full object-cover z-[1]"
            :class="[
              portfolio.id == 4
                ? 'mt-0 sm:mt-4 md:mt-0 sm:w-full sm:h-[90%]'
                : '',
            ]"
            :alt="portfolio.title + `-image`"
            loading="eager"
          />
        </div>
        <div
          v-if="portfolio.description"
          class="absolute inset-x-0 -mt-[50px] sm:-mt-5 z-[2] text-[1rem] leading-5 md:text-[0.875rem] lg:text-[1rem] md:leading-4 lg:leading-5 text-center font-inter-medium"
          :class="[
            activeIndex == portfolio.id
              ? 'text-black-core/[0.87] scale-[0.7] sm:scale-[0.8]'
              : 'text-transparent',
            portfolio.id == 7 ? 'lg:px-2.5' : '',
            portfolio.id == 5 ? '!-mt-[50px]' : '',
            portfolio.video ? '-mt-10 sm:!-mt-5 lg:!-mt-[30px]' : '',
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
          classes: "flex-[40%] bg-[#F5E1B6]",
          title: "Justly",
          description:
            "Justly aims to tackle loneliness, depression, and mental health through innovative solutions.",
          url: "/portfolio/justly",
          event: "tap_justly_portfolio_card",
        },
        {
          id: 2,
          video: [tognessMp4, tognessWebm],
          classes: "flex-[40%] bg-[#B3D8FE]",
          title: "Togness",
          description:
            "Togness is a photo editor and slideshow maker app for your life’s most memorable events.",
          url: "/portfolio/togness",
          event: "tap_togness_portfolio_card",
        },
        {
          id: 3,
          images: [luxeradio_1_400w, luxeradio_1_800w],
          classes: "flex-[20%] hidden md:block bg-[#EBB7DB]",
          title: "Luxeradio",
          url: "/portfolio/luxeradio",
          event: "tap_luxeradio_portfolio_card",
        },
        {
          id: 4,
          images: [luxeradio_2_400w, luxeradio_2_800w],
          classes: "flex-[50%] bg-[#E9CFED]",
          title: "Luxeradio",
          description:
            "Luxe Radio displays the best of Moroccan and international creation, emphasizing taste, elegance, and refinement.",
          url: "/portfolio/luxeradio",
          event: "tap_luxeradio_portfolio_card",
        },
        {
          id: 5,
          video: [smileMp4, smileWebm],
          classes: "flex-[50%] bg-[#FED7CC]",
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
          classes: "flex-[25%] hidden md:block bg-[#CDE3F7]",
          title: "Justly",
          url: "/portfolio/justly",
          event: "tap_justly_portfolio_card",
        },
        {
          id: 7,
          images: [togness_1_400w, togness_1_800w],
          classes: "flex-[50%] hidden md:block bg-[#D6E1BF]",
          title: "Togness",
          description:
            "Togness is a photo editor and slideshow maker app for your life’s most memorable events.",
          url: "/portfolio/togness",
          event: "tap_togness_portfolio_card",
        },
        {
          id: 8,
          images: [justly_3_400w, justly_3_800w],
          classes: "flex-[25%] hidden md:block bg-[#AED2F6]",
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
      this.$mixpanel.track(portfolio.event);
    },
  },
};
</script>
