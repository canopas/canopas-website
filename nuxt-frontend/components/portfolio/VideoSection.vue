<template>
  <section class="bg-white relative" :ref="response.ref">
    <div class="relative container">
      <img
        v-if="response.backgroundImage[3]"
        :src="response.backgroundImage[3]"
        :srcset="`${response.backgroundImage[0]} 400w, ${response.backgroundImage[1]} 800w, ${response.backgroundImage[2]} 1200w, ${response.backgroundImage[3]} 1600w`"
        class="hidden absolute object-cover w-full py-40 lg:block"
        alt="luxeradio-music-background"
      />
      <div
        class="flex flex-col justify-between py-20 lg:flex-row lg:py-80"
        :class="[id == 'justly' ? '!pt-16 xl:!pt-44 xl:!pb-56' : '']"
      >
        <div class="v2-normal-text font-bold">{{ response.title }}</div>
        <div class="pt-5 lg:pl-16 lg:w-4/5 lg:pt-0">
          <div class="v2-normal-text font-light">
            {{ response.description }}
          </div>

          <div
            class="flex w-fit flex-row pt-12 md:pt-16 pb-12 absolute"
            v-if="response.buttons"
          >
            <div
              v-for="button in response.buttons"
              :key="button"
              class="flex items-center pr-4 py-2 mb-6 sm:mb-0"
            >
              <a
                target="_blank"
                :href="button.link"
                @click.native="$mixpanel.track(button.event)"
                ><img
                  :src="button.image"
                  class="w-48 h-full"
                  :alt="button.name"
              /></a>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
  <section v-if="response.video" class="container video relative z-[-1]">
    <aspect-ratio height="56.25%" class="overflow-hidden">
      <video
        class="w-full h-full object-cover"
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

  <section v-if="response.videoBackground" class="video relative z-[-1]">
    <aspect-ratio height="56.25%" class="overflow-hidden">
      <img
        :src="response.videoBackground[3]"
        :srcset="`${response.videoBackground[0]} 400w, ${response.videoBackground[1]} 800w, ${response.videoBackground[2]} 1400w, ${response.videoBackground[3]} 2400w`"
        class="w-full h-full object-cover"
        :alt="response.alt"
      />
      <video
        autoplay
        loop
        muted
        playsinline
        v-if="response.animation"
        class="absolute right-0 bottom-0 left-0 top-2.5 sm:top-[18px] md:top-5 xl:top-[35px] m-auto h-3/3 w-[19%] rounded-md md:rounded-xl lg:rounded-3xl"
      >
        <source :src="response.animation" type="video/mp4" />
      </video>
    </aspect-ratio>
  </section>
  <section v-if="response.features" class="bg-white mt-20 lg:mt-0">
    <div class="container flex flex-col md:flex-row md:pr-6 lg:pr-10 xl:pr-14">
      <div v-if="response.features.gridData1" class="basis-1/2">
        <div v-for="data in response.features.gridData1" :key="data">
          <aspect-ratio
            :height="data.aspectRatio"
            class="mb-4 md:mb-6 lg:mb-8 xl:mb-10 2xl:mb-12"
            :class="data.id == 1 ? 'md:w-[92%]' : 'md:w-[110%]'"
          >
            <img
              v-if="data.image"
              class="w-full h-full object-cover"
              :src="data.image[2]"
              :srcset="`${data.image[0]} 400w, ${data.image[1]} 800w, ${data.image[2]} 1200w`"
              :alt="data.alt"
            />
          </aspect-ratio>
        </div>
      </div>
      <div v-if="response.features.gridData2" class="basis-1/2">
        <div v-for="data in response.features.gridData2" :key="data">
          <aspect-ratio
            :height="data.aspectRatio"
            class="mb-4 md:mb-6 lg:mb-8 xl:mb-10 2xl:mb-12"
            :class="
              data.id == 3
                ? 'md:w-[110%]'
                : 'md:w-[92%] md:ml-[3.8rem] lg:ml-20 xl:ml-24 2xl:ml-28'
            "
          >
            <img
              v-if="data.image"
              class="w-full h-full object-cover"
              :src="data.image[2]"
              :srcset="`${data.image[0]} 400w, ${data.image[1]} 800w, ${data.image[2]} 1200w`"
              :alt="data.alt"
            />
          </aspect-ratio>
        </div>
      </div>
    </div>
  </section>

  <section v-if="response.solution" class="bg-white relative">
    <div class="container" :ref="response.solution.ref">
      <div class="flex flex-col justify-between pt-20 lg:flex-row lg:pt-80">
        <div class="v2-normal-text font-bold">
          {{ response.solution.title }}
        </div>
        <div class="pt-5 lg:pl-16 lg:w-4/5 lg:pt-0">
          <div class="v2-normal-text font-light">
            {{ response.solution.description }}
          </div>

          <div
            class="flex flex-row w-fit pt-12 md:pt-16"
            v-if="response.solution.buttons"
          >
            <div
              v-for="button in response.solution.buttons"
              :key="button"
              class="flex items-center pr-4 py-2"
            >
              <a
                target="_blank"
                :href="button.link"
                @click.native="$mixpanel.track(button.event)"
                ><img :src="button.image" class="w-48" :alt="button.name"
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
  },
  inject: ["mixpanel"],
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
  @apply z-0;
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
