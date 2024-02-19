<template>
  <div>
    <div class="container mt-16 lg:mt-60 text-center">
      <div class="mobile-header-2 lg:desk-header-2 text-primary-1">
        Life at Canopas
      </div>
      <div
        class="xl:w-[90%] 2xl:w-[77%] xl:mx-auto sub-h1-regular lg:mobile-header-2-regular text-black-60 mt-4 lg:mt-6 mb-8 lg:mb-[4.5rem]"
      >
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
            class="mb-20 border-solid border-transparent border"
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

.swiper-pagination-bullet {
  @apply !w-[0.938rem] !h-[0.938rem] !rounded-[3px] !bg-transparent !border !border-solid !border-black-80;
}

.swiper-pagination-bullet-active {
  @apply !border-none !from-orange-300 !to-pink-300 !bg-gradient-[45deg];
}
</style>
