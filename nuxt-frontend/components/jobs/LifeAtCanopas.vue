<template>
  <div>
    <div
      class="container normal-text px-5 pt-[3.125rem] pb-[1.875rem] md:px-[3.125rem] md:pt-[6.25rem] md:pb-[3.125rem] lg:pt-[14.063rem] text-center"
    >
      <div class="header-text canopas-gradient-text">
        <span class="underline-text">Life a</span>t Canopas
      </div>
      <div class="description mx-auto my-[30px] md:my-[50px]">
        We wanted to create an environment that helps you get out of bed in the
        morning and an office that creates excellent work and encourages fun. As
        we play the infinite game, it is not about winning or losing, it's about
        striving to be better at whatever we do.
      </div>
    </div>
    <div class="swiper-content flex">
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
        <swiper-slide v-for="(slider, index) in slides" :key="index">
          <aspect-ratio
            height="66%"
            class="mb-[80px] border-solid border-transparent border"
          >
            <img
              @click.native="$mixpanel.track('click_life_photo')"
              :src="slider.image[0]"
              :srcset="`${slider.image[0]} 400w, ${slider.image[1]} 800w, ${slider.image[2]} 1600w`"
              alt="Life at canopas"
              class="h-full w-full object-cover"
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

import life1_400w from "@/assets/images/life/jobs_canopas_life_1-400w.webp";
import life1_800w from "@/assets/images/life/jobs_canopas_life_1-800w.webp";
import life1_1600w from "@/assets/images/life/jobs_canopas_life_1-1600w.webp";

import life2_400w from "@/assets/images/life/jobs_canopas_life_2-400w.webp";
import life2_800w from "@/assets/images/life/jobs_canopas_life_2-800w.webp";
import life2_1600w from "@/assets/images/life/jobs_canopas_life_2-1600w.webp";

import life3_400w from "@/assets/images/life/jobs_canopas_life_3-400w.webp";
import life3_800w from "@/assets/images/life/jobs_canopas_life_3-800w.webp";
import life3_1600w from "@/assets/images/life/jobs_canopas_life_3-1600w.webp";

import life4_400w from "@/assets/images/life/jobs_canopas_life_4-400w.webp";
import life4_800w from "@/assets/images/life/jobs_canopas_life_4-800w.webp";
import life4_1600w from "@/assets/images/life/jobs_canopas_life_4-1600w.webp";

import life5_400w from "@/assets/images/life/jobs_canopas_life_5-400w.webp";
import life5_800w from "@/assets/images/life/jobs_canopas_life_5-800w.webp";
import life5_1600w from "@/assets/images/life/jobs_canopas_life_5-1600w.webp";

import life6_400w from "@/assets/images/life/jobs_canopas_life_6-400w.webp";
import life6_800w from "@/assets/images/life/jobs_canopas_life_6-800w.webp";
import life6_1600w from "@/assets/images/life/jobs_canopas_life_6-1600w.webp";

import life7_400w from "@/assets/images/life/jobs_canopas_life_7-400w.webp";
import life7_800w from "@/assets/images/life/jobs_canopas_life_7-800w.webp";
import life7_1600w from "@/assets/images/life/jobs_canopas_life_7-1600w.webp";

import life8_400w from "@/assets/images/life/jobs_canopas_life_8-400w.webp";
import life8_800w from "@/assets/images/life/jobs_canopas_life_8-800w.webp";
import life8_1600w from "@/assets/images/life/jobs_canopas_life_8-1600w.webp";

import life9_400w from "@/assets/images/life/jobs_canopas_life_9-400w.webp";
import life9_800w from "@/assets/images/life/jobs_canopas_life_9-800w.webp";
import life9_1600w from "@/assets/images/life/jobs_canopas_life_9-1600w.webp";

import life10_400w from "@/assets/images/life/jobs_canopas_life_10-400w.webp";
import life10_800w from "@/assets/images/life/jobs_canopas_life_10-800w.webp";
import life10_1600w from "@/assets/images/life/jobs_canopas_life_10-1600w.webp";

import life11_400w from "@/assets/images/life/jobs_canopas_life_11-400w.webp";
import life11_800w from "@/assets/images/life/jobs_canopas_life_11-800w.webp";
import life11_1600w from "@/assets/images/life/jobs_canopas_life_11-1600w.webp";

import life12_400w from "@/assets/images/life/jobs_canopas_life_12-400w.webp";
import life12_800w from "@/assets/images/life/jobs_canopas_life_12-800w.webp";
import life12_1600w from "@/assets/images/life/jobs_canopas_life_12-1600w.webp";

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
};
</script>

<style lang="postcss">
@import "swiper/css";
@import "swiper/css/pagination";

.swiper {
  @apply z-0;
}

.swiper-pagination-bullet {
  @apply !w-[15px] !h-[15px] !rounded-[3px] !bg-transparent !border !border-solid !border-[#3d3d3d];
}

.swiper-pagination-bullet-active {
  @apply !border-none !from-[#F69259] !to-[#F16975] !bg-gradient-[45deg];
}
</style>
