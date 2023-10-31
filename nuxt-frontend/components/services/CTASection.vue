<template>
  <section class="relative mt-32 xl:mt-40 mb-20 xl:mb-0">
    <div class="xl:hidden">
      <img
        :src="ctabgImages[0]"
        :srcset="`${ctabgImages[0]} 400w`"
        loading="lazy"
        alt="cta-bg-image"
        class="absolute top-[-6.5rem] sm:left-[4.5rem] md:left-[6.5rem] w-[100%] sm:w-[80%] tw object-content"
      />
    </div>
    <div class="container xl:flex xl:flex-row">
      <div
        class="flex flex-col items-center xl:items-start xl:mb-20 xl:w-[60%] 2xl:w-[50%]"
      >
        <p
          class="font-inter-bold text-center xl:text-left text-[1.875rem] md:text-[2.813rem] lg:text-[3.438rem] leading-9 md:leading-[3.75rem] lg:leading-[5.156rem] font-inter-bold text-black-core/[0.87]"
        >
          Want to build a new version of your existing app or add new features?
        </p>
        <p
          class="mt-4 text-[1rem] md:text-[1.25rem] leading-6 md:leading-[1.875rem] font-inter-medium text-center xl:text-left text-black-core/[0.87] lg:text-black-core/[0.60]"
        >
          Not sure where to start? We also offer code and architecture reviews,
          strategic planning, and more.
        </p>
        <nuxt-link
          class="flex items-center m-0 mt-6 w-max rounded-full p-3 text-center gradient-btn consultation-btn"
          to="/contact"
          @click.native="$mixpanel.track('tap_services_cta')"
        >
          <span
            class="mr-2.5 text-[1.188rem] leading-[1.425rem] md:text-[1.09375rem] md:leading-[1.3125rem] lg:text-[1.1875rem] lg:leading-[1.4375rem] font-inter-semibold !tracking-[0]"
            >Talk to our experts
          </span>
        </nuxt-link>
      </div>
      <div class="hidden xl:block relative w-[53%] xl:w-[40%] 2xl:w-[50%]">
        <img
          :src="ctaImages[0]"
          :srcset="`${ctaImages[0]} 400w, ${ctaImages[1]} 800w`"
          loading="lazy"
          class="absolute right-0 h-full object-contain z-[2] mr-[-50px]"
          alt="cta-image"
        />
        <div
          class="swiper-content flex xl:items-center w-full h-full pl-4 md:pl-0"
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
            :spaceBetween="20"
            :navigation="true"
            :modules="modules"
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
              class="cursor-pointer"
            >
              <div
                :class="slide.hasStar ? 'justify-between' : 'justify-center'"
                class="flex flex-col items-center rounded-2xl bg-gradient-to-r from-[#FF835B]/[0.3] to-[#F2709C]/[0.3] w-[50%] p-5 h-[140px]"
              >
                <span class="font-inter-bold text-[1.5rem] leading-9">{{
                  slide.title
                }}</span>
                <span
                  class="font-inter-regular text-[1.1875rem] leading-[1.78125rem] text-black-core/[0.6]"
                  >{{ slide.content }}</span
                >
                <div v-if="slide.hasStar" class="flex justify-between mt-[3px]">
                  <Icon
                    v-for="i in 5"
                    :key="i"
                    class="fa w-[15px] h-[15px] mr-[5px] text-[#F2709C]"
                    name="fa6-solid:star"
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
import { Navigation, Pagination, Autoplay } from "swiper/modules";
import { Swiper, SwiperSlide } from "swiper/vue";
import ctabg400w from "@/assets/images/services/cta/cta-bg-400w.webp";
import ctabg800w from "@/assets/images/services/cta/cta-bg-800w.webp";
import cta400w from "@/assets/images/services/cta/cta-400w.webp";
import cta800w from "@/assets/images/services/cta/cta-800w.webp";

export default {
  data() {
    return {
      swiper: null,
      modules: [Navigation, Pagination, Autoplay],
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
  @apply !items-center;
}
</style>
