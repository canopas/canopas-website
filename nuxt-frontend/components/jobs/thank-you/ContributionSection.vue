<template>
  <div class="xll:container flex flex-col lg:flex-row my-16 lg:mb-60">
    <div
      class="flex flex-col xl:mt-10 lg:w-1/2 2xl:w-2/5 lg:pl-[5%] 2xll:pl-[13%] xll:pl-0"
    >
      <span
        class="lg:w-[87%] xl:w-4/5 2xl:w-[88%] text-center lg:text-left mobile-header-2 lg:desk-header-2 text-black-87"
      >
        We are here to Contribute
      </span>
      <div
        class="mt-2 lg:mt-4 text-center lg:text-left sub-h1-regular lg:mobile-header-2-regular text-black-60"
      >
        <span>and that starts with the site itself,</span>
        <br />
        <span>
          <a
            class="relative bg-gradient-[1deg] pb-[5px] v2-canopas-gradient-text sub-h3-semibold lg:mobile-header-2-semibold gradient-underline no-underline after:absolute after:bottom-0 after:left-0 after:w-full after:h-0.5 after:bg-gradient-to-l after:to-pink-300 after:from-orange-300"
            :href="websiteOpenSourceUrl"
            target="_blank"
            @click.native="$mixpanel.track('tap_canopas_website_github')"
            >canopas is open source!</a
          ></span
        >

        <div class="hidden lg:block">
          <nuxt-link
            class="ml-0 mt-8 gradient-btn primary-btn"
            to="/contact"
            @click.native="$mixpanel.track('tap_landing_cta')"
          >
            <span class="sub-h1-semibold">GitHub </span>
          </nuxt-link>
        </div>
      </div>
    </div>
    <div class="swiper-content mt-8 lg:mt-0 lg:w-1/2 2xl:w-3/5 pl-4 md:pl-0">
      <swiper
        :slidesPerView="1.1"
        :centeredSlides="false"
        :loop="true"
        :spaceBetween="20"
        :pagination="pagination"
        :navigation="true"
        :modules="modules"
        :autoplay="{
          delay: 2000,
          disableOnInteraction: false,
        }"
        :breakpoints="{
          '576': {
            slidesPerView: 1.2,
          },
          '768': {
            slidesPerView: 2.2,
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
          class="cursor-pointer"
          @click="
            openBlog(contribution.link, 'tap_jobs_thankyou_contributions')
          "
        >
          <img
            v-if="contribution.image"
            :src="contribution.image[1]"
            :srcset="`${contribution.image[0]} 400w, ${contribution.image[1]} 800w`"
            class="w-full h-full object-cover rounded-xl mb-2"
            alt="contribution-image"
          />

          <video
            class="lozad rounded-2xl mb-2"
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
              class="pr-5 mobile-header-3-semibold lg:desk-header-3 text-black-87"
            >
              {{ contribution.title }}
            </a>
          </div>
          <div
            class="flex justify-between sub-h4-regular lg:sub-h4-medium text-black-60"
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

        <div class="pt-16 mx-auto lg:hidden">
          <div class="swiper-pagination"></div>
        </div>
      </swiper>
    </div>
    <div class="block lg:hidden mx-auto">
      <nuxt-link
        class="ml-0 mt-8 gradient-btn primary-btn"
        to="/contact"
        @click.native="$mixpanel.track('tap_landing_cta')"
      >
        <span class="sub-h3-semibold">GitHub </span>
      </nuxt-link>
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
  mounted() {
    lozad().observe(); // lazy loads videos with default selector as '.lozad'
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

.swiper-pagination-bullet {
  @apply !w-3 !h-3 !rounded-full !bg-transparent !border !border-solid !border-[#3d3d3d];
}

.swiper-pagination-bullet-active {
  @apply !border-none !bg-pink-300;
}
</style>
