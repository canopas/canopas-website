<template>
  <section class="tw-relative tw-mt-32 lg:tw-mt-20 tw-mb-20 lg:tw-mb-0">
    <aspect-ratio class="lg:tw-hidden" height="30%">
      <img
        :src="ctabgImages[0]"
        :srcset="`${ctabgImages[0]} 400w`"
        loading="lazy"
        alt="cta-bg-image"
        class="tw-absolute tw-top-[-6.5rem] sm:tw-left-[4.5rem] md:tw-left-[6.5rem] tw-w-[100%] sm:tw-w-[80%] tw tw-object-content"
      />
    </aspect-ratio>
    <div class="tw-container lg:tw-flex lg:tw-flex-row">
      <div
        class="tw-flex tw-flex-col tw-items-center lg:tw-items-start lg:tw-mb-10 lg:tw-w-[47%] xl:tw-w-[60%] 2xl:tw-w-[50%]"
      >
        <p
          class="tw-font-inter-bold tw-text-[1.875rem] md:tw-text-[2.813rem] tw-leading-[2.813rem] md:tw-leading-[4.063rem] tw-text-center lg:tw-text-left tw-text-black-core/[0.87]"
        >
          Want to build a new version of your existing app or add new features?
        </p>
        <p
          class="tw-mt-4 tw-text-[1rem] md:tw-text-[1.25rem] tw-leading-[1.5rem] md:tw-leading-[1.875rem] tw-font-inter-medium tw-text-center lg:tw-text-left tw-text-black-core/[0.87]"
        >
          Not sure where to start? <br />
          We also offer code and architecture reviews, strategic planning, and
          more.
        </p>
        <router-link
          class="tw-flex tw-items-center tw-m-0 tw-mt-6 tw-w-max tw-rounded-full tw-p-3 tw-text-center gradient-btn consultation-btn"
          to="/contact"
          @click.native="mixpanel.track('tap_services_cta')"
        >
          <span
            class="tw-mr-2.5 tw-text-[1.188rem] tw-leading-[1.425rem] md:tw-text-[1.09375rem] md:tw-leading-[1.3125rem] lg:tw-text-[1.1875rem] lg:tw-leading-[1.4375rem] tw-font-inter-semibold !tw-tracking-[0]"
            >Talk to our experts
          </span>
        </router-link>
      </div>
      <div
        class="tw-hidden lg:tw-block tw-relative tw-w-[53%] xl:tw-w-[40%] 2xl:tw-w-[50%]"
      >
        <aspect-ratio height="50%">
          <img
            :src="ctaImages[0]"
            :srcset="`${ctaImages[0]} 400w, ${ctaImages[1]} 800w`"
            loading="lazy"
            class="tw-absolute tw-right-0 tw-w-[100%] xl:tw-w-[95%] 2xl:tw-w-[66%] tw-h-full tw-object-cover tw-z-[2]"
            alt="cta-image"
          />
        </aspect-ratio>
        <div
          class="swiper-content tw-mt-[3rem] xl:tw-mt-[6rem] lg:tw-w-[43%] xl:tw-w-[45%] 2xl:tw-w-[60%] tw-pl-4 md:tw-pl-0"
        >
          <swiper
            :slidesPerView="1"
            :centeredSlides="false"
            :speed="5000"
            :autoplay="{
              delay: 3000,
              disableOnInteraction: false,
            }"
            :loop="true"
            :loopedSlides="50"
            :spaceBetween="20"
            :navigation="true"
            @swiper="onSwiper"
            class="swiper-container"
          >
            <swiper-slide
              v-for="(slide, index) in slides"
              :key="index"
              @mouseover="playSwiper(false)"
              @mouseleave="playSwiper(true)"
              @touchstart.passive="playSwiper(false)"
              @touchmove.passive="playSwiper(true)"
              @touchend="playSwiper(true)"
              class="tw-cursor-pointer"
            >
              <aspect-ratio
                height="50%"
                class="tw-mb-[80px] tw-border-solid tw-border-transparent tw-border-[1px]"
              >
                <img
                  :src="slide.image[0]"
                  class="tw-mb-10 tw-mr-0 tw-w-full tw-h-full tw-object-cover"
                  alt="cta-slider-image"
                />
              </aspect-ratio>
            </swiper-slide>
          </swiper>
        </div>
      </div>
    </div>
  </section>
</template>
<script>
import SwiperCore, { Pagination, Navigation, Autoplay } from "swiper";
import { Swiper, SwiperSlide } from "swiper/vue";
import ctabg400w from "@/assets/images/services/cta/cta-bg-400w.webp";
import ctabg800w from "@/assets/images/services/cta/cta-bg-800w.webp";
import cta400w from "@/assets/images/services/cta/cta-400w.webp";
import cta800w from "@/assets/images/services/cta/cta-800w.webp";
import slide1400w from "@/assets/images/services/cta/slide-1-400w.webp";
import slide2400w from "@/assets/images/services/cta/slide-2-400w.webp";
import AspectRatio from "../utils/AspectRatio.vue";

SwiperCore.use([Pagination, Navigation, Autoplay]);
export default {
  components: { AspectRatio },
  data() {
    return {
      swiper: null,
      ctabgImages: [ctabg400w, ctabg800w],
      ctaImages: [cta400w, cta800w],
      slides: [
        {
          image: [slide1400w],
        },
        {
          image: [slide2400w],
        },
      ],
      pagination: {
        el: ".swiper-pagination",
        clickable: true,
      },
    };
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
    },
  },
};
</script>
<style lang="postcss">
@import "swiper/css";

.swiper-wrapper {
  @apply !tw-items-center;
}
</style>
