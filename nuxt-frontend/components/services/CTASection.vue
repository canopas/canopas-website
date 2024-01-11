<template>
  <section
    class="mt-16 md:mt-32 lg:mt-60 mb-20 lg:mb-0 lg:pb-28 lg:bg-gradient-to-t lg:from-orange-1 lg:via-pink-1 lg:to-white 3xl:outer-container"
  >
    <div class="lg:hidden relative mx-auto w-[400px]">
      <img
        :src="ctabgImages[0]"
        :srcset="`${ctabgImages[0]} 400w`"
        loading="lazy"
        alt="cta-bg-image"
        class="absolute top-[-14] md:top-[-6.5rem] left-2 object-contain mx-auto"
      />
    </div>
    <div class="container flex flex-row">
      <div class="flex flex-col items-center lg:items-start lg:w-3/5">
        <p
          class="text-center lg:text-left mobile-header-2 lg:desk-header-2 text-black-87 w-[69%] lg:w-[87%] 2xl:w-[82%]"
        >
          Want to build a new version of your existing app or add new features?
        </p>
        <p
          class="mt-4 text-center lg:text-left sub-h1-regular lg:mobile-header-2-regular text-black-60 w-[78%] xl:w-[70%] 2xl:w-[65%]"
        >
          Not sure where to start? We also offer code and architecture reviews,
          strategic planning, and more.
        </p>
        <nuxt-link
          class="relative mt-8 lg:ml-0 gradient-btn primary-btn"
          to="/contact"
          @click.native="$mixpanel.track('tap_landing_cta')"
        >
          <span class="sub-h3-semibold lg:sub-h1-semibold"
            >Talk To Our Experts
          </span>
        </nuxt-link>
      </div>
      <div class="hidden lg:block relative w-2/5">
        <aspect-ratio class="w-[96%] 2xl:w-[80%] lg:block hidden" height="60%">
          <img
            :src="ctaImages[0]"
            :srcset="`${ctaImages[0]} 400w, ${ctaImages[1]} 800w`"
            loading="lazy"
            class="relative h-full w-full z-[2] object-cover ml-10"
            alt="cta-image"
          />
        </aspect-ratio>
        <div
          class="absolute top-[5.5rem] xl:-top-10 swiper-content flex xl:items-center w-full h-full pl-4 md:pl-0 lg:ml-[-3.75rem] xl:ml-[-1.625rem]"
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
              class="cursor-pointer"
            >
              <div
                :class="slide.hasStar ? 'justify-between' : 'justify-center'"
                class="flex flex-col items-center rounded-2xl bg-gradient-to-r from-orange-1 to-pink-1 w-64 p-5 h-32"
              >
                <span class="mobile-header-2-bold text-black-87">{{
                  slide.title
                }}</span>
                <span class="sub-h1-semibold text-black-60">{{
                  slide.content
                }}</span>
                <div
                  v-if="slide.hasStar"
                  class="flex justify-between mt-[0.188rem]"
                >
                  <Icon
                    v-for="i in 5"
                    :key="i"
                    class="fa w-[15px] h-[15px] mr-[0.313rem] secondary-color"
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
};
</script>
<style lang="postcss">
@import "swiper/css";

.swiper-wrapper {
  @apply !items-center;
}
</style>
