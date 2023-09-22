<template>
  <div>
    <div
      class="tw-container normal-text tw-px-5 tw-pt-[3.125rem] tw-pb-[1.875rem] md:tw-px-[3.125rem] md:tw-pt-[6.25rem] md:tw-pb-[3.125rem] lg:tw-pt-[14.063rem] tw-text-center"
    >
      <div class="header-text canopas-gradient-text">
        <span class="underline-text">Life a</span>t Canopas
      </div>
      <div class="description tw-mx-auto tw-my-[30px] md:tw-my-[50px]">
        We wanted to create an environment that helps you get out of bed in the
        morning and an office that creates excellent work and encourages fun. As
        we play the infinite game, it is not about winning or losing, it's about
        striving to be better at whatever we do.
      </div>
    </div>
    <div class="swiper-content tw-flex">
      <swiper
        :slidesPerView="1"
        :centeredSlides="true"
        :preloadImages="false"
        :lazy="true"
        :autoplay="{
          delay: 3000,
          disableOnInteraction: false,
        }"
        :loop="true"
        :spaceBetween="20"
        :pagination="pagination"
        :modules="modules"
        :breakpoints="{
          '768': {
            slidesPerView: 1.125,
          },
          '1024': {
            slidesPerView: 1.5,
          },
          '1800': {
            slidesPerView: 2.1,
          },
        }"
        class="swiper-container"
      >
        <swiper-slide v-for="slider in slides" :key="slider.id">
          <aspect-ratio
            height="66%"
            class="tw-mb-[80px] tw-border-solid tw-border-transparent tw-border"
          >
            <img
              @click.native="trackLifePhotoClick"
              :src="slider.image[0]"
              :srcset="`${slider.image[0]} 400w, ${slider.image[1]} 800w, ${slider.image[2]} 1600w`"
              alt="Life at Canopas"
              class="tw-h-full tw-w-full tw-object-cover"
              loading="lazy"
            />
          </aspect-ratio>
        </swiper-slide>
        <div class="swiper-pagination"></div>
      </swiper>
    </div>
  </div>
</template>

<script>
import AspectRatio from "@/components/utils/AspectRatio.vue";
import { Autoplay, Pagination } from "swiper/modules";
import { Swiper, SwiperSlide } from "swiper/vue";

import {
  life1_400w,
  life1_800w,
  life1_1600w,
  life2_400w,
  life2_800w,
  life2_1600w,
  life3_400w,
  life3_800w,
  life3_1600w,
  life4_400w,
  life4_800w,
  life4_1600w,
  life5_400w,
  life5_800w,
  life5_1600w,
  life6_400w,
  life6_800w,
  life6_1600w,
  life7_400w,
  life7_800w,
  life7_1600w,
  life8_400w,
  life8_800w,
  life8_1600w,
  life9_400w,
  life9_800w,
  life9_1600w,
  life10_400w,
  life10_800w,
  life10_1600w,
  life11_400w,
  life11_800w,
  life11_1600w,
  life12_400w,
  life12_800w,
  life12_1600w,
} from "@/assets/images/life";

export default {
  data() {
    return {
      modules: [Pagination, Autoplay],
      slides: [
        {
          id: 2,
          image: [life2_400w, life2_800w, life2_1600w],
        },
        {
          id: 12,
          image: [life12_400w, life12_800w, life12_1600w],
        },
        {
          id: 7,
          image: [life7_400w, life7_800w, life7_1600w],
        },
        {
          id: 1,
          image: [life1_400w, life1_800w, life1_1600w],
        },
        {
          id: 4,
          image: [life4_400w, life4_800w, life4_1600w],
        },
        {
          id: 10,
          image: [life10_400w, life10_800w, life10_1600w],
        },
        {
          id: 3,
          image: [life3_400w, life3_800w, life3_1600w],
        },
        {
          id: 9,
          image: [life9_400w, life9_800w, life9_1600w],
        },
        {
          id: 6,
          image: [life6_400w, life6_800w, life6_1600w],
        },
        {
          id: 8,
          image: [life8_400w, life8_800w, life8_1600w],
        },
        {
          id: 5,
          image: [life5_400w, life5_800w, life5_1600w],
        },
        {
          id: 11,
          image: [life11_400w, life11_800w, life11_1600w],
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
    AspectRatio,
  },
  inject: ["mixpanel"],
  methods: {
    trackLifePhotoClick() {
      this.mixpanel.track("click_life_photo");
    },
  },
};
</script>

<style lang="postcss">
@import "swiper/css";
@import "swiper/css/pagination";

.swiper {
  @apply tw-z-0;
}

.swiper-pagination-bullet {
  @apply !tw-w-[15px] !tw-h-[15px] !tw-rounded-[3px] !tw-bg-transparent !tw-border !tw-border-solid !tw-border-[#3d3d3d];
}

.swiper-pagination-bullet-active {
  @apply !tw-border-none !tw-from-[#F69259] !tw-to-[#F16975] !tw-bg-gradient-[45deg];
}
</style>
