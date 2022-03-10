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
        :loop="true"
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
import SwiperCore, { Pagination } from "swiper";
import { Swiper, SwiperSlide } from "swiper/vue";
import AspectRatio from "./../utils/AspectRatio.vue";
import "swiper/css";
import "swiper/css/pagination";

SwiperCore.use([Pagination]);

export default {
  data() {
    return {
      slides: [
        {
          id: 1,
          image: require("./../../assets/images/life/jobs_canopas_life_2.jpg"),
        },
        {
          id: 2,
          image: require("./../../assets/images/life/jobs_canopas_life_1.jpg"),
        },
        {
          id: 3,
          image: require("./../../assets/images/life/jobs_canopas_life_4.jpg"),
        },
        {
          id: 4,
          image: require("./../../assets/images/life/jobs_canopas_life_3.jpg"),
        },
        {
          id: 5,
          image: require("./../../assets/images/life/jobs_canopas_life_6.jpg"),
        },
        {
          id: 6,
          image: require("./../../assets/images/life/jobs_canopas_life_5.jpg"),
        },
        {
          id: 7,
          image: require("./../../assets/images/life/jobs_canopas_life_7.jpg"),
        },
        {
          id: 8,
          image: require("./../../assets/images/life/jobs_canopas_life_8.jpg"),
        },
      ],
      pagination: {
        el: ".swiper-pagination",
        clickable: true,
      },
      modules: [Pagination],
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
