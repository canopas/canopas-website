<template>
  <div class="container flex lg:items-center m-auto lg:w-full pb-16 lg:pb-60">
    <div class="flex flex-col lg:flex-wrap lg:flex-row">
      <div
        v-for="portfolio in portfolios"
        :key="portfolio"
        :class="[
          portfolio.classes,
          portfolio.id != 8 ? 'p-4 sm:p-8' : '',
          portfolio.id == 4 ? 'sm:!pb-12' : '',
          activeIndex == portfolio.id
            ? 'lg:scale-105 lg::scale-[1.2] lg:z-[2] '
            : '',
          portfolio.video ? 'sm:!p-0' : '',
        ]"
        class="relative lg:transition-all lg:duration-500 lg:ease-in-out lg:transform cursor-pointer rounded-xl mt-4 lg:mt-0 md:w-[80%] md:mx-auto"
        @mouseover="(activeIndex = portfolio.id)"
        @mouseleave="(activeIndex = null)"
        @touchstart.passive="(activeIndex = portfolio.id)"
      >
        <nuxt-link
          :to="portfolio.url"
          @click.native="$mixpanel.track(portfolio.event)"
        >
          <div
            class="text-black-87 absolute top-[18px] sm:top-[13px] lg:top-[8px] xl:top-[10px] 2xl:top-[19px] inset-x-0 z-[2] text-center mobile-header-3-semibold lg:desk-header-3"
            :class="[
              activeIndex == portfolio.id
                ? ' lg:scale-[0.8]'
                : 'lg:text-transparent',
              portfolio.video ? 'xl:!mt-[8px]' : '',
              portfolio.id == 7 ? 'lg:!top-[-6px] 2xl:!top-0' : '',
            ]"
          >
            {{ portfolio.title }}
          </div>
          <div
            class="flex items-center h-full"
            :class="portfolio.video ? 'p-0' : 'py-8 sm:p-0 lg:pt-0 lg:pb-2'"
          >
            <video
              v-if="portfolio.video"
              autoplay
              loop
              muted
              playsinline
              class="lozad mx-auto sm:h-[400px] 2xl:h-[509px]"
            >
              <source :data-src="portfolio.video[1]" type="video/webm" />
              <source :data-src="portfolio.video[0]" type="video/mp4" />
            </video>
            <img
              v-else
              :src="portfolio.images[0]"
              :srcset="`${portfolio.images[0]} 400w, ${portfolio.images[1]} 800w`"
              class="h-full w-full object-contain z-[1]"
              :class="[
                portfolio.id == 4
                  ? 'mt-0 sm:mt-4 lg:mt-0 sm:w-full sm:h-[90%]'
                  : '',
                portfolio.id == 8 ? 'object-cover' : '',
              ]"
              :alt="portfolio.title + `-image`"
              height="100"
              width="100"
              loading="eager"
            />
          </div>
          <div
            v-if="portfolio.description"
            class="text-black-60 absolute inset-x-0 -mt-[35px] sm:-mt-[10px] xl:-mt-[25px] z-[2] w-[92%] sm:w-[90%] md:w-[75%] lg:w-full mx-auto text-center sub-h4-semibold lg:sub-h1-regular"
            :class="[
              activeIndex == portfolio.id
                ? 'lg:scale-[0.8]'
                : 'lg:text-transparent',
              portfolio.id == 7
                ? 'lg:px-2.5 lg:-mt-[20px] 2xl:-mt-[23px]'
                : 'lg:-mt-[46px] 2xl:-mt-[41px]',
              portfolio.video
                ? 'sm:-mt-[50px] lg:-mt-[74px] xl:-mt-[70px] 2xl:-mt-[75px]'
                : '',
            ]"
          >
            {{ portfolio.description }}
          </div>
        </nuxt-link>
      </div>
    </div>
  </div>
</template>

<script type="module">
import config from "@/config.js";
import lozad from "lozad";

export default {
  data() {
    return {
      portfolios: [
        {
          id: 1,
          images: [
            "/images/portfolio/new-portfolio/justly-1-400w.webp",
            "/images/portfolio/new-portfolio/justly-1-800w.webp",
          ],
          classes: "flex-[40%] bg-[#F6EAD0] lg:rounded-none lg:rounded-tl-xl",
          title: "Justly",
          description:
            "Justly aims to tackle loneliness, depression, and mental health through innovative solutions.",
          url: "/portfolio/justly",
          event: "tap_justly_portfolio_card",
        },
        {
          id: 2,
          video: [
            "/images/portfolio/new-portfolio/togness.mp4",
            "/images/portfolio/new-portfolio/togness.webm",
          ],
          classes: "flex-[40%] bg-[#E1F1FF] lg:rounded-none",
          title: "Togness",
          description:
            "Togness is a photo editor and slideshow maker app for your life’s most memorable events.",
          url: "/portfolio/togness",
          event: "tap_togness_portfolio_card",
        },
        {
          id: 3,
          images: [
            "/images/portfolio/new-portfolio/luxeradio-1-400w.webp",
            "/images/portfolio/new-portfolio/luxeradio-1-800w.webp",
          ],
          classes:
            "flex-[20%] hidden lg:block bg-[#EBB7DB] lg:rounded-none lg:rounded-tr-xl",
          title: "Luxeradio",
          url: "/portfolio/luxeradio",
          event: "tap_luxeradio_portfolio_card",
        },
        {
          id: 4,
          images: [
            "/images/portfolio/new-portfolio/luxeradio-2-400w.webp",
            "/images/portfolio/new-portfolio/luxeradio-2-800w.webp",
          ],
          classes: "flex-[50%] bg-[#F1E8F2] lg:rounded-none",
          title: "Luxeradio",
          description:
            "Luxe Radio displays the best of Moroccan and international creation, emphasizing taste, elegance, and refinement.",
          url: "/portfolio/luxeradio",
          event: "tap_luxeradio_portfolio_card",
        },
        {
          id: 5,
          video: [
            "/images/portfolio/new-portfolio/smile.mp4",
            "/images/portfolio/new-portfolio/smile.webm",
          ],
          classes: "flex-[50%] bg-[#FFE9E5] lg:rounded-none",
          title: "Smile+",
          description:
            "Smile+ app is designed for dentists to create a perfect smile for their patients automated with AI.",
          url: config.SMILEPLUS_URL,
          target: "_blank",
          event: "tap_smileplus_portfolio_card",
        },
        {
          id: 6,
          images: [
            "/images/portfolio/new-portfolio/justly-2-400w.webp",
            "/images/portfolio/new-portfolio/justly-2-800w.webp",
          ],
          classes:
            "flex-[25%] hidden lg:block bg-[#CDE3F7] lg:rounded-none lg:rounded-bl-xl",
          title: "Justly",
          url: "/portfolio/justly",
          event: "tap_justly_portfolio_card",
        },
        {
          id: 7,
          images: [
            "/images/portfolio/new-portfolio/togness-1-400w.webp",
            "/images/portfolio/new-portfolio/togness-1-800w.webp",
          ],
          classes: "flex-[50%] hidden lg:block bg-[#E6F0D0] lg:rounded-none",
          title: "Togness",
          description:
            "Togness is a photo editor and slideshow maker app for your life’s most memorable events.",
          url: "/portfolio/togness",
          event: "tap_togness_portfolio_card",
        },
        {
          id: 8,
          images: [
            "/images/portfolio/new-portfolio/justly-3-400w.webp",
            "/images/portfolio/new-portfolio/justly-3-800w.webp",
          ],
          classes:
            "flex-[25%] hidden lg:block bg-[#AED2F6] lg:rounded-none lg:rounded-br-xl",
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
};
</script>
