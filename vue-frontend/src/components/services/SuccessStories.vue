<template>
  <section
    class="xll:tw-container tw-flex tw-flex-col lg:tw-flex-row tw-mt-20 tw-mb-20 lg:tw-my-[14.063rem]"
  >
    <div
      class="tw-flex tw-flex-col tw-mt-10 lg:tw-mt-[4.563rem] lg:tw-w-[50%] 2xl:tw-w-[40%] lg:tw-pl-[8%] 2xl:tw-pl-[10%] xll:tw-pl-0"
    >
      <p
        class="lg:tw-w-[87%] xl:tw-w-[80%] 2xl:tw-w-[88%] tw-text-center lg:tw-text-left tw-text-[1.875rem] md:tw-text-[2.813rem] lg:tw-text-[3.438rem] tw-leading-[2.5rem] md:tw-leading-[3.75rem] lg:tw-leading-[5.156rem] tw-font-inter-bold tw-text-black-core/[0.87]"
      >
        Our Success Stories
      </p>
      <p
        class="tw-mt-4 tw-text-center lg:tw-text-left tw-font-inter-regular tw-text-[1.2rem] md:tw-text-[1.25rem] lg:tw-text-[1.5rem] tw-leading-[1.575rem] md:tw-leading-[1.875rem] lg:tw-leading-9 tw-text-black-core/[0.6]"
      >
        Check out the best from our vault of<br />
        50+ projects.
      </p>
      <router-link
        class="tw-hidden lg:tw-block tw-flex tw-items-center tw-m-0 tw-mt-6 tw-w-max tw-rounded-full tw-p-3 tw-text-center gradient-btn consultation-btn"
        :to="portfolioURL"
        @click.native="mixpanel.track('tap_services_portfolio')"
      >
        <span
          class="tw-mr-2.5 tw-font-normal tw-text-[1rem] tw-leading-[1.1875rem] md:tw-text-[1.09375rem] md:tw-leading-[1.3125rem] lg:tw-text-[1.1875rem] lg:tw-leading-[1.4375rem] tw-font-inter-medium !tw-tracking-[0]"
          >View Work
        </span>
      </router-link>
    </div>
    <div
      class="swiper-content tw-mt-12 lg:tw-w-[50%] 2xl:tw-w-[60%] tw-pl-4 md:tw-pl-0"
    >
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
          class="tw-cursor-pointer"
          @click="openPortfolio(story)"
        >
          <img
            :src="story.image[0]"
            :srcset="`${story.image[0]} 400w, ${story.image[1]} 800w`"
            class="tw-mb-2 tw-w-full tw-h-[600px] lg:tw-h-full tw-object-cover"
            alt="story-image"
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
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
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
    FontAwesomeIcon,
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
      this.mixpanel.track(story.event);
    },
  },
};
</script>

<style lang="postcss">
@import "swiper/css";
@import "swiper/css/pagination";
.swiper-wrapper {
  @apply !tw-items-center;
}
.swiper-horizontal > .swiper-pagination-progressbar,
.swiper-pagination-horizontal,
.swiper-pagination-progressbar-fill {
  @apply lg:tw-hidden !tw-bottom-0 !tw-top-auto !tw-left-auto md:!tw-left-[0.2rem] !tw-right-4 tw-w-[99%] !tw-h-1.5 !tw-rounded-[18px] !tw-bg-neutral-100;
}
.swiper-pagination-progressbar-fill {
  @apply !tw-right-0  md:!tw-left-[0.2rem]  !tw-h-1.5 !tw-rounded-[18px]  !tw-bg-gradient-to-r !tw-from-[#FF835B] !tw-to-[#F2709C];
}
</style>
