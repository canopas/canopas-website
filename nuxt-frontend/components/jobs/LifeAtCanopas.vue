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
import axios from "axios";
import config from "@/config.js";
import { getDiffrentWidthImages } from "@/utils.js";

export default {
  data() {
    return {
      getDiffrentWidthImages,
      modules: [Pagination, Autoplay],
      slides: [],
      pagination: {
        el: ".swiper-pagination",
        clickable: true,
      },
    };
  },
  methods: {
    async fetchImages() {
      try {
        const response = await axios.get(config.API_BASE + "/api/lifeimages");
        this.slides = getDiffrentWidthImages(response);
      } catch (error) {
        console.error("Error fetching data from the API", error);
      }
    },
  },
  created() {
    this.fetchImages();
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
