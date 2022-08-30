<template>
  <div>
    <div
      class="normal-text tw-px-[20px] tw-pt-[50px] tw-pb-[30px] md:tw-px-[50px] md:tw-pt-[100px] md:tw-pb-[50px] lg:tw-pt-[150px] tw-text-center"
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
          disableOnInteraction: true,
        }"
        :loop="true"
        :loopedSlides="50"
        :spaceBetween="20"
        :pagination="pagination"
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
        <swiper-slide v-for="(sider, index) in slides" :key="index">
          <aspect-ratio
            height="66%"
            class="tw-mb-[80px] tw-border-solid tw-border-transparent tw-border-[1px]"
          >
            <img
              :src="sider.image[3]"
              :srcset="`${sider.image[0]} 400w, ${sider.image[1]} 800w, ${sider.image[2]} 1600w, ${sider.image[3]} 2400w`"
              alt="Life at canopas"
              class="swiper-slide tw-object-cover"
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
import SwiperCore, { Pagination, Autoplay } from "swiper";
import { Swiper, SwiperSlide } from "swiper/vue";

import life1_400w from "@/assets/images/life/jobs_canopas_life_1-400w.webp";
import life1_800w from "@/assets/images/life/jobs_canopas_life_1-800w.webp";
import life1_1600w from "@/assets/images/life/jobs_canopas_life_1-1600w.webp";
import life1_2400w from "@/assets/images/life/jobs_canopas_life_1-2400w.webp";

import life2_400w from "@/assets/images/life/jobs_canopas_life_2-400w.webp";
import life2_800w from "@/assets/images/life/jobs_canopas_life_2-800w.webp";
import life2_1600w from "@/assets/images/life/jobs_canopas_life_2-1600w.webp";
import life2_2400w from "@/assets/images/life/jobs_canopas_life_2-2400w.webp";

import life3_400w from "@/assets/images/life/jobs_canopas_life_3-400w.webp";
import life3_800w from "@/assets/images/life/jobs_canopas_life_3-800w.webp";
import life3_1600w from "@/assets/images/life/jobs_canopas_life_3-1600w.webp";
import life3_2400w from "@/assets/images/life/jobs_canopas_life_3-2400w.webp";

import life4_400w from "@/assets/images/life/jobs_canopas_life_4-400w.webp";
import life4_800w from "@/assets/images/life/jobs_canopas_life_4-800w.webp";
import life4_1600w from "@/assets/images/life/jobs_canopas_life_4-1600w.webp";
import life4_2400w from "@/assets/images/life/jobs_canopas_life_4-2400w.webp";

import life5_400w from "@/assets/images/life/jobs_canopas_life_5-400w.webp";
import life5_800w from "@/assets/images/life/jobs_canopas_life_5-800w.webp";
import life5_1600w from "@/assets/images/life/jobs_canopas_life_5-1600w.webp";
import life5_2400w from "@/assets/images/life/jobs_canopas_life_5-2400w.webp";

import life6_400w from "@/assets/images/life/jobs_canopas_life_6-400w.webp";
import life6_800w from "@/assets/images/life/jobs_canopas_life_6-800w.webp";
import life6_1600w from "@/assets/images/life/jobs_canopas_life_6-1600w.webp";
import life6_2400w from "@/assets/images/life/jobs_canopas_life_6-2400w.webp";

import life7_400w from "@/assets/images/life/jobs_canopas_life_7-400w.webp";
import life7_800w from "@/assets/images/life/jobs_canopas_life_7-800w.webp";
import life7_1600w from "@/assets/images/life/jobs_canopas_life_7-1600w.webp";
import life7_2400w from "@/assets/images/life/jobs_canopas_life_7-2400w.webp";

import life8_400w from "@/assets/images/life/jobs_canopas_life_8-400w.webp";
import life8_800w from "@/assets/images/life/jobs_canopas_life_8-800w.webp";
import life8_1600w from "@/assets/images/life/jobs_canopas_life_8-1600w.webp";
import life8_2400w from "@/assets/images/life/jobs_canopas_life_8-2400w.webp";

import life9_400w from "@/assets/images/life/jobs_canopas_life_9-400w.webp";
import life9_800w from "@/assets/images/life/jobs_canopas_life_9-800w.webp";
import life9_1600w from "@/assets/images/life/jobs_canopas_life_9-1600w.webp";
import life9_2400w from "@/assets/images/life/jobs_canopas_life_9-2400w.webp";

import life10_400w from "@/assets/images/life/jobs_canopas_life_10-400w.webp";
import life10_800w from "@/assets/images/life/jobs_canopas_life_10-800w.webp";
import life10_1600w from "@/assets/images/life/jobs_canopas_life_10-1600w.webp";
import life10_2400w from "@/assets/images/life/jobs_canopas_life_10-2400w.webp";

import life11_400w from "@/assets/images/life/jobs_canopas_life_11-400w.webp";
import life11_800w from "@/assets/images/life/jobs_canopas_life_11-800w.webp";
import life11_1600w from "@/assets/images/life/jobs_canopas_life_11-1600w.webp";
import life11_2400w from "@/assets/images/life/jobs_canopas_life_11-2400w.webp";

import life12_400w from "@/assets/images/life/jobs_canopas_life_12-400w.webp";
import life12_800w from "@/assets/images/life/jobs_canopas_life_12-800w.webp";
import life12_1600w from "@/assets/images/life/jobs_canopas_life_12-1600w.webp";
import life12_2400w from "@/assets/images/life/jobs_canopas_life_12-2400w.webp";

SwiperCore.use([Pagination, Autoplay]);

export default {
  data() {
    return {
      slides: [
        {
          id: 2,
          image: [life2_400w, life2_800w, life2_1600w, life2_2400w],
        },
        {
          id: 12,
          image: [life12_400w, life12_800w, life12_1600w, life12_2400w],
        },
        {
          id: 7,
          image: [life7_400w, life7_800w, life7_1600w, life7_2400w],
        },
        {
          id: 1,
          image: [life1_400w, life1_800w, life1_1600w, life1_2400w],
        },
        {
          id: 4,
          image: [life4_400w, life4_800w, life4_1600w, life4_2400w],
        },
        {
          id: 10,
          image: [life10_400w, life10_800w, life10_1600w, life10_2400w],
        },
        {
          id: 3,
          image: [life3_400w, life3_800w, life3_1600w, life3_2400w],
        },
        {
          id: 9,
          image: [life9_400w, life9_800w, life9_1600w, life9_2400w],
        },
        {
          id: 6,
          image: [life6_400w, life6_800w, life6_1600w, life6_2400w],
        },
        {
          id: 8,
          image: [life8_400w, life8_800w, life8_1600w, life8_2400w],
        },
        {
          id: 5,
          image: [life5_400w, life5_800w, life5_1600w, life5_2400w],
        },
        {
          id: 11,
          image: [life11_400w, life11_800w, life11_1600w, life11_2400w],
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
};
</script>

<style lang="postcss">
@import "swiper/css";
@import "swiper/css/pagination";

.swiper {
  @apply tw-z-0;
}
</style>
