<template>
  <div
    class="xll:tw-container tw-flex tw-flex-col lg:tw-flex-row tw-my-20 lg:tw-mb-40"
  >
    <div
      class="tw-flex tw-flex-col xl:tw-mt-10 lg:tw-w-[50%] 2xl:tw-w-[40%] lg:tw-pl-[5%] 2xll:tw-pl-[13%] xll:tw-pl-0"
    >
      <span
        class="lg:tw-w-[87%] xl:tw-w-[80%] 2xl:tw-w-[88%] tw-text-center lg:tw-text-left tw-font-inter-bold tw-text-black-core/[.87] tw-text-[1.875rem] md:tw-text-[3.4375rem] tw-leading-[2.5rem] md:tw-leading-[5.5625rem]"
      >
        We are here to Contribute
      </span>
      <div
        class="tw-mt-2 md:tw-mt-0 tw-text-center lg:tw-text-left tw-font-inter-medium tw-text-black-core/[0.6] tw-text-[1rem] tw-leading-[1.375rem] md:tw-text-[1.5rem] md:tw-leading-[2.5625rem]"
      >
        <span>and that starts with the site itself,</span>
        <br />
        <span>
          <a
            class="tw-relative tw-bg-gradient-[1deg] tw-pb-[5px] v2-canopas-gradient-text tw-font-inter-medium gradient-underline tw-no-underline after:tw-absolute after:tw-bottom-0 after:tw-left-0 after:tw-w-full after:tw-h-0.5 after:tw-bg-gradient-to-l after:tw-to-pink-300 after:tw-from-orange-300"
            :href="websiteOpenSourceUrl"
            target="_blank"
            @click.native="mixpanel.track('tap_canopas_website_github')"
            >canopas is open source!</a
          ></span
        >

        <div class="tw-hidden lg:tw-block">
          <a
            class="tw-flex tw-items-center tw-mt-6 tw-ml-0 tw-w-max tw-rounded-full tw-p-3 tw-text-center gradient-btn consultation-btn"
            :href="contributionURL"
            target="_blank"
          >
            <span
              class="tw-mr-2.5 tw-font-inter-semibold tw-text-[1.1875rem] tw-leading-[1.425rem]"
              >GitHub
            </span>
          </a>
        </div>
      </div>
    </div>
    <div
      class="swiper-content tw-mt-12 lg:tw-mt-0 lg:tw-w-[50%] 2xl:tw-w-[60%] tw-pl-4 md:tw-pl-0"
    >
      <swiper
        :slidesPerView="1.1"
        :centeredSlides="true"
        :autoplay="{
          delay: 2000,
          disableOnInteraction: false,
        }"
        :loop="true"
        :spaceBetween="20"
        :pagination="pagination"
        :navigation="true"
        :modules="modules"
        :breakpoints="{
          '576': {
            slidesPerView: 1.2,
          },
          '768': {
            centeredSlides: true,
            slidesPerView: 1.525,
          },
          '992': {
            centeredSlides: false,
            slidesPerView: 2.2,
          },
        }"
        @swiper="onSwiper"
        class="swiper-container"
      >
        <swiper-slide
          v-for="(contribution, index) in contributions"
          :key="index"
          @mouseover="playSwiper(false)"
          @mouseleave="playSwiper(true)"
          @touchstart.passive="playSwiper(false)"
          @touchmove.passive="playSwiper(true)"
          @touchend="playSwiper(true)"
          class="tw-cursor-pointer"
          @click="
            openBlog(contribution.link, 'tap_jobs_thankyou_contributions')
          "
        >
          <img
            v-if="contribution.image"
            :src="contribution.image[1]"
            :srcset="`${contribution.image[0]} 400w, ${contribution.image[1]} 800w`"
            class="tw-w-full tw-h-full tw-object-cover tw-rounded-2xl tw-mb-5"
            alt="contribution-image"
          />

          <video
            class="lozad tw-rounded-2xl tw-mb-5"
            autoplay
            loop
            muted
            playsinline
            v-if="contribution.video"
          >
            <source :data-src="contribution.video[1]" type="video/webm" />
            <source :data-src="contribution.video[0]" type="video/mp4" />
          </video>

          <div>
            <a
              target="_blank"
              :href="contribution.link"
              class="tw-pr-5 tw-font-inter-semibold tw-text-black-core/[0.87] tw-text-[1.25rem] md:tw-text-[1.875rem] tw-leading-[1.875rem] md:tw-leading-[2.4rem]"
            >
              {{ contribution.title }}
            </a>
          </div>
          <div
            @touchstart.passive="activeIndex = index"
            class="tw-flex tw-flex-row tw-justify-between tw-font-inter-regular tw-text-black-core/[0.60] tw-text-[0.875rem] md:tw-text-[1rem] tw-leading-[1.3125rem] md:tw-leading-6"
          >
            <span>by {{ contribution.author }}</span>
            <a
              class="tw-duration-300 tw-ease-in-out hover:tw-scale-110"
              :href="contribution.link"
              :aria-label="contribution.title"
              target="_blank"
            >
            </a>
          </div>
        </swiper-slide>

        <div class="tw-mt-16 tw-mx-auto lg:tw-hidden">
          <div class="swiper-pagination"></div>
        </div>
      </swiper>
    </div>
    <div class="tw-block lg:tw-hidden tw-mx-auto">
      <a
        class="tw-flex tw-items-center tw-m-0 tw-mt-6 tw-w-max tw-rounded-full tw-p-3 tw-text-center gradient-btn consultation-btn"
        :href="contributionURL"
        target="_blank"
      >
        <span
          class="tw-mr-2.5 tw-font-inter-semibold tw-text-[1.1875rem] tw-leading-[1.425rem]"
          >GitHub
        </span>
      </a>
    </div>
  </div>
</template>
<script type="module">
import { Autoplay, Pagination, Navigation } from "swiper/modules";
import { Swiper, SwiperSlide } from "swiper/vue";
import { library } from "@fortawesome/fontawesome-svg-core";
import { faGithub } from "@fortawesome/free-brands-svg-icons";
import { faArrowRightLong } from "@fortawesome/free-solid-svg-icons";

import uIPilot400 from "@/assets/images/jobsthankyou/contribution/UIPilot-400w.webp";
import uIPilot800 from "@/assets/images/jobsthankyou/contribution/UIPilot-800w.webp";
import introShowCase400 from "@/assets/images/jobsthankyou/contribution/introShowCase-400w.webp";
import introShowCase800 from "@/assets/images/jobsthankyou/contribution/introShowCase-800w.webp";
import jcAnimationsMp4 from "@/assets/images/jobsthankyou/contribution/JetpackComposeAnimations.mp4";
import jcAnimationsWebm from "@/assets/images/jobsthankyou/contribution/JetpackComposeAnimations.webm";
import Config from "@/config.js";
import lozad from "lozad";
import { openBlog } from "@/utils.js";

library.add(faGithub);

export default {
  data() {
    return {
      openBlog,
      swiper: null,
      modules: [Autoplay, Pagination, Navigation],
      contributionURL: Config.GITHUB_URL,
      icon: faGithub,
      rightArrow: faArrowRightLong,
      websiteOpenSourceUrl: Config.WEBSITE_OPEN_SOURCE_URL,
      contributions: [
        {
          title: "Intro Showcase view",
          author: "Radhika S.",
          image: [introShowCase400, introShowCase800],
          video: "",
          link: "https://github.com/canopas/Intro-showcase-view",
        },
        {
          title: "UIPilot",
          author: "Jimmy S.",
          image: [uIPilot400, uIPilot800],
          video: "",
          link: "https://github.com/canopas/UIPilot",
        },
        {
          title: "Jetpack Compose animations",
          author: "Radhika S.",
          video: [jcAnimationsMp4, jcAnimationsWebm],
          image: "",
          link: "https://github.com/canopas/Jetpack-compose-animations-examples",
        },
      ],
      pagination: {
        el: ".swiper-pagination",
        clickable: true,
      },
    };
  },
  computed: {
    contributions() {
      return Array(50).fill(this.contributions).flat();
    },
  },
  mounted() {
    lozad().observe(); // lazy loads videos with default selector as '.lozad'
  },
  methods: {
    onSwiper(swiper) {
      this.swiper = swiper;
    },
    playSwiper(play) {
      play ? this.swiper.autoplay.start() : this.swiper.autoplay.stop();
      this.reading = !play;
    },
  },
  components: {
    Swiper,
    SwiperSlide,
  },
  // inject: ["mixpanel"],
};
</script>
<style lang="postcss">
@import "swiper/css";
@import "swiper/css/pagination";

.swiper {
  @apply tw-z-0;
}

.swiper-pagination-bullet {
  @apply !tw-w-3 !tw-h-3 !tw-rounded-full !tw-bg-transparent !tw-border !tw-border-solid !tw-border-[#3d3d3d];
}

.swiper-pagination-bullet-active {
  @apply !tw-border-none !tw-from-[#F69259] !tw-to-[#F16975] !tw-bg-gradient-[45deg];
}
</style>
