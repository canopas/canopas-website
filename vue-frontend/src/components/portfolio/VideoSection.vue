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
        class="tw-flex tw-flex-col tw-justify-between tw-py-20 lg:tw-flex-row lg:tw-py-80"
        :class="[id == 'justly' ? '!tw-pt-16 xl:!tw-pt-44 xl:!tw-pb-56' : '']"
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
                ><img
                  :src="button.image"
                  class="tw-w-48 tw-h-full"
                  :alt="button.name"
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
        class="tw-absolute tw-right-0 tw-bottom-0 tw-left-0 tw-top-2.5 sm:tw-top-[18px] md:tw-top-5 xl:tw-top-[35px] tw-m-auto tw-h-3/3 tw-w-[19%] tw-rounded-md md:tw-rounded-xl lg:tw-rounded-3xl"
      >
        <source :src="response.animation" type="video/mp4" />
      </video>
    </aspect-ratio>
  </section>
  <section v-if="response.features" class="tw-bg-white tw-mt-20 lg:tw-mt-0">
    <div
      class="tw-container tw-flex tw-flex-col md:tw-flex-row md:tw-pr-6 lg:tw-pr-10 xl:tw-pr-14"
    >
      <div v-if="response.features.gridData1" class="tw-basis-1/2">
        <div v-for="data in response.features.gridData1" :key="data">
          <aspect-ratio
            :height="data.aspectRatio"
            class="tw-mb-4 md:tw-mb-6 lg:tw-mb-8 xl:tw-mb-10 2xl:tw-mb-12"
            :class="data.id == 1 ? 'md:tw-w-[92%]' : 'md:tw-w-[110%]'"
          >
            <img
              v-if="data.image"
              class="tw-w-full tw-h-full tw-object-cover"
              :src="data.image[2]"
              :srcset="`${data.image[0]} 400w, ${data.image[1]} 800w, ${data.image[2]} 1200w`"
              :alt="data.alt"
            />
          </aspect-ratio>
        </div>
      </div>
      <div v-if="response.features.gridData2" class="tw-basis-1/2">
        <div v-for="data in response.features.gridData2" :key="data">
          <aspect-ratio
            :height="data.aspectRatio"
            class="tw-mb-4 md:tw-mb-6 lg:tw-mb-8 xl:tw-mb-10 2xl:tw-mb-12"
            :class="
              data.id == 3
                ? 'md:tw-w-[110%]'
                : 'md:tw-w-[92%] md:tw-ml-[3.8rem] lg:tw-ml-20 xl:tw-ml-24 2xl:tw-ml-28'
            "
          >
            <img
              v-if="data.image"
              class="tw-w-full tw-h-full tw-object-cover"
              :src="data.image[2]"
              :srcset="`${data.image[0]} 400w, ${data.image[1]} 800w, ${data.image[2]} 1200w`"
              :alt="data.alt"
            />
          </aspect-ratio>
        </div>
      </div>
    </div>
  </section>

  <section v-if="response.solution" class="tw-bg-white tw-relative">
    <div class="tw-container" :ref="response.solution.ref">
      <div
        class="tw-flex tw-flex-col tw-justify-between tw-pt-20 lg:tw-flex-row lg:tw-pt-80"
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
import { Autoplay, Pagination } from "swiper/modules";
import { Swiper, SwiperSlide } from "swiper/vue";
import { elementInViewPort } from "@/utils.js";

export default {
  props: ["json"],

  data() {
    return {
      modules: [Pagination, Autoplay],
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
        justlyfeature: "view_justly_feature",
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
