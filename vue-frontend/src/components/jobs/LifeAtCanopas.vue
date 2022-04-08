<template>
  <div class="pb-32">
    <div class="container normal-text text-center">
      <div class="header-text canopas-gradient-text">
        <span class="underline-text">Life a</span>t Canopas
      </div>
      <div class="description">
        We wanted to create an environment that helps you get out of bed in the
        morning and an office that creates excellent work and encourages fun. As
        we play the infinite game, it is not about winning or losing, it's about
        striving to be better at whatever we do.
      </div>
    </div>
    <div class="swiper-content">
      <swiper
        :slidesPerView="1"
        :centeredSlides="true"
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
          <aspect-ratio height="66%" class="image-container">
            <img
              :src="sider.image"
              alt="Life at canopas"
              class="swiper-slide"
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
import "swiper.css";
import "swiper.css.pagination";

import life1 from "@/assets/images/life/jobs_canopas_life_1.jpg";
import life2 from "@/assets/images/life/jobs_canopas_life_2.jpg";
import life3 from "@/assets/images/life/jobs_canopas_life_3.jpg";
import life4 from "@/assets/images/life/jobs_canopas_life_4.jpg";
import life5 from "@/assets/images/life/jobs_canopas_life_5.jpg";
import life6 from "@/assets/images/life/jobs_canopas_life_6.jpg";
import life7 from "@/assets/images/life/jobs_canopas_life_7.jpg";
import life8 from "@/assets/images/life/jobs_canopas_life_8.jpg";
import life9 from "@/assets/images/life/jobs_canopas_life_9.jpg";
import life10 from "@/assets/images/life/jobs_canopas_life_10.jpg";
import life11 from "@/assets/images/life/jobs_canopas_life_11.jpg";

SwiperCore.use([Pagination, Autoplay]);

export default {
  data() {
    return {
      slides: [
        {
          id: 1,
          image: life2,
        },
        {
          id: 2,
          image: life1,
        },
        {
          id: 3,
          image: life4,
        },
        {
          id: 4,
          image: life10,
        },
        {
          id: 5,
          image: life3,
        },
        {
          id: 6,
          image: life6,
        },
        {
          id: 7,
          image: life9,
        },
        {
          id: 8,
          image: life7,
        },
        {
          id: 9,
          image: life8,
        },
        {
          id: 10,
          image: life5,
        },
        {
          id: 11,
          image: life11,
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

<style lang="scss" scoped>
.container {
  padding: 50px 20px 30px;
}

.description {
  margin: 30px auto;
  color: rgba(61, 61, 61, 0.8);
}

.swiper-content {
  display: flex;
}

.swiper {
  z-index: 0;
}

.image-container {
  margin-bottom: 80px;
  border: 1px solid transparent;

  img {
    object-fit: cover;
  }
}

@include media-breakpoint-up(md) {
  .container {
    padding: 100px 50px 50px;
  }
  .description {
    margin: 50px auto;
  }
}

@include media-breakpoint-up(lg) {
  .container {
    padding-top: 150px;
  }
}
</style>
