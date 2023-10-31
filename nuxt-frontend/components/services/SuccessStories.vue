<template>
  <section
    class="xll:container flex flex-col lg:flex-row mt-20 mb-20 lg:my-[14.063rem]"
  >
    <div
      class="flex flex-col mt-10 lg:mt-[4.563rem] lg:w-[50%] 2xl:w-[40%] lg:pl-[8%] 2xl:pl-[10%] xll:pl-0"
    >
      <p
        class="lg:w-[87%] xl:w-[80%] 2xl:w-[88%] text-center lg:text-left text-[1.875rem] md:text-[2.813rem] lg:text-[3.438rem] leading-[2.5rem] md:leading-[3.75rem] lg:leading-[5.156rem] font-inter-bold text-black-core/[0.87]"
      >
        Our Success Stories
      </p>
      <p
        class="mt-4 text-center lg:text-left font-inter-regular text-[1.2rem] md:text-[1.25rem] lg:text-[1.5rem] leading-[1.575rem] md:leading-[1.875rem] lg:leading-9 text-black-core/[0.6]"
      >
        Check out the best from our vault of<br />
        50+ projects.
      </p>
      <nuxt-link
        class="hidden lg:block flex items-center m-0 mt-6 w-max rounded-full p-3 text-center gradient-btn consultation-btn"
        :to="portfolioURL"
        @click.native="$mixpanel.track('tap_services_portfolio')"
      >
        <span
          class="mr-2.5 font-normal text-[1rem] leading-[1.1875rem] md:text-[1.09375rem] md:leading-[1.3125rem] lg:text-[1.1875rem] lg:leading-[1.4375rem] font-inter-medium !tracking-[0]"
          >View Work
        </span>
      </nuxt-link>
    </div>
    <div class="swiper-content mt-12 lg:w-[50%] 2xl:w-[60%] pl-4 md:pl-0">
      <swiper
        :slidesPerView="1.1"
        :centeredSlides="false"
        :autoplay="{
          delay: 2000,
          disableOnInteraction: false,
        }"
        :loop="true"
        :spaceBetween="20"
        :pagination="{ type: 'progressbar' }"
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
          v-for="(story, index) in stories"
          :key="index"
          @mouseover="playSwiper(false)"
          @mouseleave="playSwiper(true)"
          @touchstart.passive="playSwiper(false)"
          @touchmove.passive="playSwiper(true)"
          @touchend="playSwiper(true)"
          class="cursor-pointer"
          @click="openPortfolio(story)"
        >
          <img
            :src="story.image[0]"
            :srcset="`${story.image[0]} 400w, ${story.image[1]} 800w`"
            class="mb-2 w-full h-[600px] lg:h-full object-cover"
            alt="story-image"
            loading="lazy"
          />
        </swiper-slide>
        <div class="swiper-pagination"></div>
      </swiper>
    </div>
  </section>
</template>

<script>
import { Pagination, Navigation, Autoplay } from "swiper/modules";
import { Swiper, SwiperSlide } from "swiper/vue";
import justly400w from "@/assets/images/services/success-stories/justly-400w.webp";
import justly800w from "@/assets/images/services/success-stories/justly-800w.webp";
import luxeradio400w from "@/assets/images/services/success-stories/luxeradio-400w.webp";
import luxeradio800w from "@/assets/images/services/success-stories/luxeradio-800w.webp";
import togness400w from "@/assets/images/services/success-stories/togness-400w.webp";
import togness800w from "@/assets/images/services/success-stories/togness-800w.webp";
import smile400w from "@/assets/images/services/success-stories/smile+400w.webp";
import smile800w from "@/assets/images/services/success-stories/smile+800w.webp";
import config from "@/config.js";

export default {
  data() {
    return {
      modules: [Pagination, Navigation, Autoplay],
      swiper: null,
      reading: false,
      portfolioURL: "/portfolio",
      stories: [
        {
          image: [justly400w, justly800w],
          url: "/portfolio/justly",
          event: "tap_services_justly_portfolio",
        },
        {
          image: [luxeradio400w, luxeradio800w],
          url: "/portfolio/luxeradio",
          event: "tap_services_luxe_radio_portfolio",
        },
        {
          image: [togness400w, togness800w],
          url: "/portfolio/togness",
          event: "tap_services_togness_portfolio",
        },
        {
          image: [smile400w, smile800w],
          url: config.SMILEPLUS_URL,
          target: "_blank",
          event: "tap_services_smile_portfolio",
        },
      ],
      pagination: {
        el: ".swiper-pagination",
        clickable: true,
      },
    };
  },
  computed: {
    stories() {
      return Array(50).fill(this.stories).flat();
    },
  },
  components: {
    Swiper,
    SwiperSlide,
  },
  inject: ["mixpanel"],
  methods: {
    onSwiper(swiper) {
      this.swiper = swiper;
    },
    playSwiper(play) {
      play ? this.swiper.autoplay.start() : this.swiper.autoplay.stop();
      this.reading = !play;
    },
    openPortfolio(story) {
      window.open(story.url, story.target ? story.target : "_self");
      this.$mixpanel.track(story.event);
    },
  },
};
</script>

<style lang="postcss">
@import "swiper/css";
@import "swiper/css/pagination";
.swiper-wrapper {
  @apply !items-center;
}
.swiper-horizontal > .swiper-pagination-progressbar,
.swiper-pagination-horizontal,
.swiper-pagination-progressbar-fill {
  @apply lg:hidden !bottom-0 !top-auto !left-auto md:!left-[0.2rem] !right-4 w-[99%] !h-1.5 !rounded-[18px] !bg-neutral-100;
}
.swiper-pagination-progressbar-fill {
  @apply !right-0  md:!left-[0.2rem]  !h-1.5 !rounded-[18px]  !bg-gradient-to-r !from-[#FF835B] !to-[#F2709C];
}
</style>
