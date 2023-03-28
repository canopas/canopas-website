<template>
  <section class="tw-bg-white tw-relative" :ref="response.ref">
    <div class="tw-relative tw-container">
      <img
        v-if="response.backgroundImage[3]"
        :src="response.backgroundImage[3]"
        :srcset="`${response.backgroundImage[0]} 400w, ${response.backgroundImage[1]} 800w, ${response.backgroundImage[2]} 1200w, ${response.backgroundImage[3]} 1600w`"
        class="tw-hidden tw-absolute tw-object-cover tw-w-full tw-py-40 lg:tw-block"
        alt="luxeradio-music-background"
      />
      <div
        :class="[
          id == 'smileplus'
            ? 'tw-flex tw-flex-col tw-justify-between tw-pt-40 lg:tw-flex-row lg:tw-pt-80'
            : 'tw-flex tw-flex-col tw-justify-between tw-py-40 lg:tw-flex-row lg:tw-py-80',
        ]"
      >
        <div class="v2-normal-text tw-font-bold">{{ response.title }}</div>
        <div class="tw-pt-5 lg:tw-pl-16 lg:tw-w-4/5 lg:tw-pt-0">
          <div class="v2-normal-text tw-font-light">
            {{ response.description }}
          </div>

          <div
            class="tw-flex tw-w-fit tw-flex-row tw-pt-12 md:tw-pt-16 tw-pb-12 tw-absolute"
            v-if="response.buttons"
          >
            <div
              v-for="button in response.buttons"
              :key="button"
              class="tw-flex tw-items-center tw-pr-4 tw-py-2 tw-mb-6 sm:tw-mb-0"
            >
              <a
                target="_blank"
                :href="button.link"
                @click.native="mixpanel.track(button.event)"
                ><img :src="button.image" class="tw-w-48" :alt="button.name"
              /></a>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
  <section
    v-if="response.video"
    class="tw-container video tw-relative tw-z-[-1]"
  >
    <aspect-ratio height="56.25%" class="tw-overflow-hidden">
      <video
        class="tw-w-full tw-h-full tw-object-cover"
        preload="auto"
        loop
        muted
        autoplay
        playsinline
        :poster="response.videoThumb"
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
      <video
        autoplay
        loop
        muted
        playsinline
        v-if="response.animation"
        class="tw-absolute tw-right-0 tw-bottom-0 tw-left-0 tw-top-[10px] sm:tw-top-[18px] md:tw-top-[20px] xl:tw-top-[35px] tw-m-auto tw-h-3/3 tw-w-[19%] tw-rounded-md md:tw-rounded-xl lg:tw-rounded-3xl"
      >
        <source :src="response.animation" type="video/mp4" />
      </video>
    </aspect-ratio>
  </section>

  <!-- mobile screen slider -->
  <section
    v-if="response.slider && isShowSliderInMobile"
    class="tw-relative tw-mt-40 md:tw-mt-20"
  >
    <div
      class="mobile-background-video tw-opacity-50 tw-overflow-hidden tw-absolute tw-inset-0 tw-rounded-full tw-z-[-1]"
      :style="{ background: backgroundColor }"
    ></div>
    <div
      class="swiper-content tw-absolute tw-flex tw-items-center tw-w-full tw-h-full tw-z-0"
    >
      <swiper
        ref="swiper"
        :slidesPerView="1"
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
            slidesPerView: 1,
          },
        }"
        @slideChange="onSlideChange"
        class="swiper-container"
      >
        <swiper-slide v-for="(sider, index) in response.slider" :key="index">
          <aspect-ratio
            height="43%"
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
      height="92%"
      class="mobile-images tw-mx-auto tw-h-full tw-w-5/12 tw-z-[-1]"
    >
      <img
        :src="response.videoBackgroundImage"
        class="tw-w-full tw-h-full tw-object-cover"
        :alt="response.alt"
      />
    </aspect-ratio>
  </section>

  <!-- large screen slider -->
  <section v-else-if="response.slider" class="tw-relative tw-mt-20 lg:tw-mt-0">
    <div
      class="background-video tw-opacity-50 tw-overflow-hidden tw-absolute tw-inset-0 tw-rounded-full tw-z-[-1]"
      :style="{ background: backgroundColor }"
    ></div>
    <div
      class="swiper-content tw-absolute tw-flex tw-items-center tw-w-full tw-h-full tw-z-0"
    >
      <swiper
        ref="swiper"
        :centeredSlides="true"
        :autoplay="{
          delay: 1500,
          disableOnInteraction: true,
        }"
        :loop="true"
        :loopedSlides="50"
        :spaceBetween="50"
        :breakpoints="{
          '993': {
            slidesPerView: 3,
          },
          '1800': {
            slidesPerView: 3,
          },
        }"
        @slideChange="onSlideChange"
        class="swiper-container video-swiper"
      >
        <swiper-slide v-for="(sider, index) in response.slider" :key="index">
          <aspect-ratio
            height="62%"
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
      height="40%"
      class="mobile-image tw-mx-auto tw-h-full tw-z-[-1]"
    >
      <img
        :src="response.videoBackgroundImage"
        class="tw-w-full tw-h-full tw-object-cover"
        :alt="response.alt"
      />
    </aspect-ratio>
  </section>

  <section v-if="response.solution" class="tw-bg-white tw-relative">
    <div class="tw-container" :ref="response.solution.ref">
      <div
        class="tw-flex tw-flex-col tw-justify-between tw-pt-40 lg:tw-flex-row lg:tw-pt-80"
      >
        <div class="v2-normal-text tw-font-bold">
          {{ response.solution.title }}
        </div>
        <div class="tw-pt-5 lg:tw-pl-16 lg:tw-w-4/5 lg:tw-pt-0">
          <div class="v2-normal-text tw-font-light">
            {{ response.solution.description }}
          </div>

          <div
            class="tw-flex tw-flex-row tw-w-fit tw-pt-12 md:tw-pt-16"
            v-if="response.solution.buttons"
          >
            <div
              v-for="button in response.solution.buttons"
              :key="button"
              class="tw-flex tw-items-center tw-pr-4 tw-py-2"
            >
              <a
                target="_blank"
                :href="button.link"
                @click.native="mixpanel.track(button.event)"
                ><img :src="button.image" class="tw-w-48" :alt="button.name"
              /></a>
            </div>
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
import { elementInViewPort } from "@/utils.js";

SwiperCore.use([Pagination, Autoplay]);

export default {
  props: ["json"],

  data() {
    return {
      id: this.$route.params.id,
      portfolioRef: null,
      backgroundColor: "",
      isShowSliderInMobile: false,
      response: this.json,
      event: "",
      events: {
        luxevideo: "view_luxe_radio_video",
        tognessvideo: "view_togness_problem",
        justlyvideo: "view_justly_horizontal_slider",
      },
    };
  },
  watch: {
    json: function (newVal, oldVal) {
      this.response = newVal;
    },
  },
  components: {
    Swiper,
    SwiperSlide,
    AspectRatio,
  },
  mounted() {
    this.isShowSliderInMobile = window.innerWidth <= 992;
    this.portfolioRef = document.getElementById("response");
    this.portfolioRef.addEventListener("scroll", this.sendEvent);
    window.addEventListener("scroll", this.sendEvent);
  },
  unmounted() {
    this.portfolioRef.removeEventListener("scroll", this.sendEvent);
    window.removeEventListener("scroll", this.sendEvent);
  },
  inject: ["mixpanel"],
  methods: {
    onSlideChange(swiper) {
      this.backgroundColor =
        this.response.slider[swiper.realIndex].backgroundColor;
    },
    sendEvent() {
      const event = this.events[elementInViewPort(this.$refs)];
      if (event && this.event !== event) {
        this.event = event;
        this.mixpanel.track(event);
      }
    },
  },
};
</script>

<style lang="postcss">
@import "swiper/css";

.is-animation-tab:after {
  content: "";
}

.background-video {
  margin: 6% 35%;
}

.mobile-background-video {
  margin: 9% 11%;
}

.video-animation {
  width: 18%;
}

.swiper-slide-active {
  transform: scale(1) !important;
  transition: 0.4s;
}

.video-swiper > .swiper-wrapper > .swiper-slide-prev,
.video-swiper > .swiper-wrapper > .swiper-slide-next {
  transform: scale(0.5);
}

.swiper-wrapper {
  display: flex;
  align-items: center;
}
.swiper {
  @apply tw-z-0;
}

.mobile-image {
  width: 18.5%;
}

@screen lg {
  section.video {
    transform: translateZ(-1px) scale(1.5);
  }
}

@screen 3xl {
  section.video {
    transform: translateZ(-1px) scale(1.3);
  }
}
</style>
