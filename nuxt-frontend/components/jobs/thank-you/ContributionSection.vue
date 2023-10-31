<template>
  <div class="xll:container flex flex-col lg:flex-row my-20 lg:mb-40">
    <div
      class="flex flex-col xl:mt-10 lg:w-[50%] 2xl:w-[40%] lg:pl-[5%] 2xll:pl-[13%] xll:pl-0"
    >
      <span
        class="lg:w-[87%] xl:w-[80%] 2xl:w-[88%] text-center lg:text-left font-inter-bold text-black-core/[.87] text-[1.875rem] md:text-[3.4375rem] leading-[2.5rem] md:leading-[5.5625rem]"
      >
        We are here to Contribute
      </span>
      <div
        class="mt-2 md:mt-0 text-center lg:text-left font-inter-medium text-black-core/[0.6] text-[1rem] leading-[1.375rem] md:text-[1.5rem] md:leading-[2.5625rem]"
      >
        <span>and that starts with the site itself,</span>
        <br />
        <span>
          <a
            class="relative bg-gradient-[1deg] pb-[5px] v2-canopas-gradient-text font-inter-medium gradient-underline no-underline after:absolute after:bottom-0 after:left-0 after:w-full after:h-0.5 after:bg-gradient-to-l after:to-pink-300 after:from-orange-300"
            :href="websiteOpenSourceUrl"
            target="_blank"
            @click.native="$mixpanel.track('tap_canopas_website_github')"
            >canopas is open source!</a
          ></span
        >

        <div class="hidden lg:block">
          <a
            class="flex items-center mt-6 ml-0 w-max rounded-full p-3 text-center gradient-btn consultation-btn"
            :href="contributionURL"
            target="_blank"
          >
            <span
              class="mr-2.5 font-inter-semibold text-[1.1875rem] leading-[1.425rem]"
              >GitHub
            </span>
          </a>
        </div>
      </div>
    </div>
    <div
      class="swiper-content mt-12 lg:mt-0 lg:w-[50%] 2xl:w-[60%] pl-4 md:pl-0"
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
          class="cursor-pointer"
          @click="
            openBlog(contribution.link, 'tap_jobs_thankyou_contributions')
          "
        >
          <img
            v-if="contribution.image"
            :src="contribution.image[1]"
            :srcset="`${contribution.image[0]} 400w, ${contribution.image[1]} 800w`"
            class="w-full h-full object-cover rounded-2xl mb-5"
            alt="contribution-image"
          />

          <video
            class="lozad rounded-2xl mb-5"
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
              class="pr-5 font-inter-semibold text-black-core/[0.87] text-[1.25rem] md:text-[1.875rem] leading-[1.875rem] md:leading-[2.4rem]"
            >
              {{ contribution.title }}
            </a>
          </div>
          <div
            class="flex flex-row justify-between font-inter-regular text-black-core/[0.60] text-[0.875rem] md:text-[1rem] leading-[1.3125rem] md:leading-6"
          >
            <span>by {{ contribution.author }}</span>
            <a
              class="duration-300 ease-in-out hover:scale-110"
              :href="contribution.link"
              :aria-label="contribution.title"
              target="_blank"
            >
            </a>
          </div>
        </swiper-slide>

        <div class="mt-16 mx-auto lg:hidden">
          <div class="swiper-pagination"></div>
        </div>
      </swiper>
    </div>
    <div class="block lg:hidden mx-auto">
      <a
        class="flex items-center m-0 mt-6 w-max rounded-full p-3 text-center gradient-btn consultation-btn"
        :href="contributionURL"
        target="_blank"
      >
        <span
          class="mr-2.5 font-inter-semibold text-[1.1875rem] leading-[1.425rem]"
          >GitHub
        </span>
      </a>
    </div>
  </div>
</template>
<script type="module">
import { Autoplay, Pagination, Navigation } from "swiper/modules";
import { Swiper, SwiperSlide } from "swiper/vue";

import uIPilot400 from "@/assets/images/jobsthankyou/contribution/UIPilot-400w.webp";
import uIPilot800 from "@/assets/images/jobsthankyou/contribution/UIPilot-800w.webp";
import introShowCase400 from "@/assets/images/jobsthankyou/contribution/introShowCase-400w.webp";
import introShowCase800 from "@/assets/images/jobsthankyou/contribution/introShowCase-800w.webp";
import jcAnimationsMp4 from "@/assets/images/jobsthankyou/contribution/JetpackComposeAnimations.mp4";
import jcAnimationsWebm from "@/assets/images/jobsthankyou/contribution/JetpackComposeAnimations.webm";
import Config from "@/config.js";
import lozad from "lozad";
import { openBlog } from "@/utils.js";

export default {
  data() {
    return {
      openBlog,
      swiper: null,
      modules: [Autoplay, Pagination, Navigation],
      contributionURL: Config.GITHUB_URL,
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
  inject: ["mixpanel"],
};
</script>
<style lang="postcss">
@import "swiper/css";
@import "swiper/css/pagination";

.swiper {
  @apply z-0;
}

.swiper-pagination-bullet {
  @apply !w-3 !h-3 !rounded-full !bg-transparent !border !border-solid !border-[#3d3d3d];
}

.swiper-pagination-bullet-active {
  @apply !border-none !from-[#F69259] !to-[#F16975] !bg-gradient-[45deg];
}
</style>
