<template>
  <section class="tw-bg-white tw-relative">
    <div class="tw-relative container">
      <img
        v-if="response.backgroundImage[3]"
        :src="response.backgroundImage[3]"
        :srcset="`${response.backgroundImage[0]} 400w, ${response.backgroundImage[1]} 800w, ${response.backgroundImage[2]} 1200w, ${response.backgroundImage[3]} 1600w`"
        class="tw-hidden tw-absolute tw-object-cover tw-w-full tw-py-40 lg:tw-block"
        alt="luxeradio-music-background"
      />

      <div
        class="tw-flex tw-flex-col tw-justify-between tw-py-20 sm:tw-py-40 lg:tw-flex-row lg:tw-py-80"
      >
        <div class="v2-normal-text tw-font-bold">{{ response.title }}</div>
        <div class="tw-pt-5 lg:tw-pl-16 lg:tw-w-4/5 lg:tw-pt-0">
          <div class="v2-normal-text tw-font-light">
            {{ response.description }}
          </div>
          <div class="tw-pt-16">
            <a
              v-for="button in response.buttons"
              :key="button"
              class="tw-pr-10 is-animation-tab tw-inline-block tw-relative v2-normal-2-text hover:tw-text-black-900 after:tw-absolute after:tw-w-1/4 after:tw-scale-x-0 after:tw-h-0.5 after:tw-bottom-0 after:tw-left-0 after:tw-bg-black-900 after:tw-origin-bottom-left after:tw-duration-300 hover:after:tw-scale-x-100 hover:after:tw-origin-bottom-left"
              :href="button.link"
              target="_blank"
            >
              {{ button.name }}
            </a>
          </div>
        </div>
      </div>
    </div>
  </section>

  <section v-if="response.video" class="container video tw-relative tw-z-[-1]">
    <aspect-ratio height="56.25%" class="tw-overflow-hidden">
      <video
        class="tw-w-full tw-h-full tw-object-cover"
        preload="auto"
        loop
        muted
        autoplay
        playsinline
      >
        <source :src="response.video" type="video/mp4" />
      </video>
    </aspect-ratio>
  </section>

  <section v-if="response.videoBackground" class="video tw-relative tw-z-[-1]">
    <aspect-ratio height="56.25%" class="tw-overflow-hidden">
      <img
        :src="response.videoBackground[3]"
        :srcset="`${response.videoBackground[0]} 400w, ${response.videoBackground[1]} 800w, ${response.videoBackground[2]} 1400w, ${response.videoBackground[3]} 2400w`"
        class="tw-w-full tw-h-full tw-object-cover"
        :alt="response.alt"
      />
      <img
        v-if="response.animation"
        :src="response.animation"
        class="tw-absolute tw-inset-0 tw-m-auto tw-h-5/5 tw-w-1/4 tw-rounded-lg"
        :alt="response.alt"
      />
    </aspect-ratio>
  </section>

  <section v-if="response.slider" class="tw-relative">
    <div
      class="background-video tw-opacity-50 tw-overflow-hidden tw-absolute tw-inset-0 tw-rounded-full tw-z-[-1]"
      :style="{ background: backgroundColor }"
    ></div>
    <div
      class="swiper-content tw-absolute tw-flex tw-items-center tw-w-full tw-h-full tw-z-0"
    >
      <swiper
        ref="swiper"
        :slidesPerView="2.5"
        :centeredSlides="true"
        :autoplay="{
          delay: 1500,
          disableOnInteraction: true,
        }"
        :loop="true"
        :loopedSlides="50"
        :spaceBetween="50"
        :breakpoints="{
          '768': {
            slidesPerView: 3,
          },
          '1024': {
            slidesPerView: 3,
          },
          '1800': {
            slidesPerView: 3,
          },
        }"
        @slideChange="onSlideChange"
        class="swiper-container"
      >
        <swiper-slide v-for="(sider, index) in response.slider" :key="index">
          <aspect-ratio
            height="85%"
            class="tw-border-solid tw-border-1 tw-border-transparent"
          >
            <img
              :src="sider.image"
              class="swiper-slide"
              loading="lazy"
              :alt="sider.alt ? sider.alt : ''"
            />
          </aspect-ratio>
        </swiper-slide>
      </swiper>
    </div>
    <aspect-ratio
      height="54%"
      class="mobile-image tw-mx-auto tw-h-full tw-w-1/4 tw-z-[-1]"
    >
      <img
        :src="response.videoBackgroundImage"
        class="tw-w-full tw-h-full tw-object-cover"
        :alt="response.alt"
      />
    </aspect-ratio>
  </section>

  <section v-if="response.solution" class="tw-bg-white tw-relative">
    <div class="container">
      <div
        class="tw-flex tw-flex-col tw-justify-between tw-pt-20 sm:tw-pt-40 lg:tw-flex-row lg:tw-pt-80"
      >
        <div class="v2-normal-text tw-font-bold">
          {{ response.title }}
        </div>
        <div class="tw-pt-5 lg:tw-pl-16 lg:tw-w-4/5 lg:tw-pt-0">
          <div class="v2-normal-text tw-font-light">
            {{ response.description }}
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import AspectRatio from "@/components/utils/AspectRatio.vue";
import SwiperCore, { Pagination, Autoplay } from "swiper";
import { Swiper, SwiperSlide } from "swiper/vue";

SwiperCore.use([Pagination, Autoplay]);

export default {
  props: ["response"],
  data() {
    return {
      backgroundColor: "",
    };
  },
  components: {
    Swiper,
    SwiperSlide,
    AspectRatio,
  },
  methods: {
    onSlideChange(swiper) {
      this.backgroundColor =
        this.response.slider[swiper.realIndex].backgroundColor;
    },
  },
};
</script>

<style lang="postcss">
@import "swiper/css";
@import "swiper/css/pagination";

.is-animation-tab:after {
  content: "";
}

section.video {
  transform: translateZ(-1px) scale(1.5);
}

.background-video {
  margin: 8% 30%;
}

.swiper-slide-active {
  transform: scale(1) !important;
  transition: 0.4s;
}

.swiper-slide-prev,
.swiper-slide-next {
  transform: scale(0.5) !important;
}
</style>
