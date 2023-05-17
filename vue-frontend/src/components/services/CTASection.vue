<template>
  <section class="tw-relative tw-mt-32 xl:tw-mt-20 tw-mb-20 xl:tw-mb-0">
    <div class="xl:tw-hidden">
      <img
        :src="ctabgImages[0]"
        :srcset="`${ctabgImages[0]} 400w`"
        loading="lazy"
        alt="cta-bg-image"
        class="tw-absolute tw-top-[-6.5rem] sm:tw-left-[4.5rem] md:tw-left-[6.5rem] tw-w-[100%] sm:tw-w-[80%] tw tw-object-content"
      />
    </div>
    <div class="tw-container xl:tw-flex xl:tw-flex-row">
      <div
        class="tw-flex tw-flex-col tw-items-center xl:tw-items-start xl:tw-mb-10 xl:tw-w-[60%] 2xl:tw-w-[50%]"
      >
        <p
          class="tw-font-inter-bold tw-text-[1.875rem] md:tw-text-[2.813rem] tw-leading-[2.813rem] md:tw-leading-[4.063rem] tw-text-center xl:tw-text-left tw-text-black-core/[0.87]"
        >
          Want to build a new version of your existing app or add new features?
        </p>
        <p
          class="tw-mt-4 tw-text-[1rem] md:tw-text-[1.25rem] tw-leading-[1.5rem] md:tw-leading-[1.875rem] tw-font-inter-medium tw-text-center xl:tw-text-left tw-text-black-core/[0.87]"
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
        class="tw-hidden xl:tw-block tw-relative tw-w-[53%] xl:tw-w-[40%] 2xl:tw-w-[50%]"
      >
        <img
          :src="ctaImages[0]"
          :srcset="`${ctaImages[0]} 400w, ${ctaImages[1]} 800w`"
          loading="lazy"
          class="tw-absolute tw-right-0 tw-w-[100%] xl:tw-w-[95%] 2xl:tw-w-[66%] tw-h-full tw-object-cover tw-z-[2] tw-mr-[-15%] 2xl:tw-mr-0"
          alt="cta-image"
        />
        <div
          class="swiper-content tw-flex xl:tw-items-center tw-w-full tw-h-full tw-pl-4 md:tw-pl-0"
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
              <div
                :class="
                  slide.hasStar ? 'tw-justify-between' : 'tw-justify-center'
                "
                class="tw-flex tw-flex-col tw-items-center tw-rounded-[16px] tw-bg-gradient-to-r tw-from-[#FF835B]/[0.3] tw-to-[#F2709C]/[0.3] tw-w-[50%] tw-p-[20px] tw-h-[140px]"
              >
                <span
                  class="tw-font-inter-bold tw-text-[1.5rem] tw-leading-[2.25rem]"
                  >{{ slide.title }}</span
                >
                <span
                  class="tw-font-inter-regular tw-text-[1.1875rem] tw-leading-[1.78125rem] tw-text-black-core/[0.6]"
                  >{{ slide.content }}</span
                >
                <div
                  v-if="slide.hasStar"
                  class="tw-flex tw-justify-between tw-mt-[3px]"
                >
                  <font-awesome-icon
                    v-for="i in 5"
                    :key="i"
                    class="fa tw-w-[15px] tw-h-[15px] tw-mr-[5px] tw-text-[#F2709C]"
                    icon="star"
                  />
                </div>
              </div>
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
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

SwiperCore.use([Pagination, Navigation, Autoplay]);

export default {
  data() {
    return {
      swiper: null,
      ctabgImages: [ctabg400w, ctabg800w],
      ctaImages: [cta400w, cta800w],
      slides: [
        {
          title: "50+",
          content: "Projects",
          hasStar: false,
        },
        {
          title: "5.0",
          content: "Clutch Reviews",
          hasStar: true,
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
    FontAwesomeIcon,
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
