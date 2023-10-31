<template>
  <div class="container my-4 mt-[6.25rem] md:mt-[14.063rem] ml-auto mx-auto">
    <div class="swiper-content">
      <swiper
        :slidesPerView="1"
        :centeredSlides="true"
        :autoplay="{
          delay: 3000,
          disableOnInteraction: false,
        }"
        :loop="true"
        :spaceBetween="20"
        :pagination="pagination"
        :modules="modules"
        @swiper="onSwiper"
        class="swiper-container"
      >
        <swiper-slide
          v-for="(client, index) in clients"
          :key="index"
          @mouseover="playSwiper(false)"
          @mouseleave="playSwiper(true)"
          @touchstart.passive="playSwiper(false)"
          @touchmove.passive="playSwiper(true)"
          @touchend="playSwiper(true)"
          class="cursor-pointer"
        >
          <div
            class="mx-auto sm:px-10 md:px-6 lg:px-4 xl:px-24 sm:text-black-core/[0.87] text-black-900"
            :class="[client.className]"
          >
            <div
              class="font-roboto-medium text-[1.375rem] md:text-[2.25rem] lg:text-[3.125rem] leading-[1.875rem] md:leading-[2.75rem] lg:leading-[3.75rem] transition-all ease duration-500"
              :class="reading ? 'scale-[0.98]' : ''"
              v-html="client.review"
            ></div>
            <div class="flex flex-row justify-between mt-11 md:mt-20">
              <div class="flex flex-col">
                <span
                  class="v2-canopas-gradient-text font-inter-semibold text-[1rem] md:text-[1.625rem]"
                  >{{ client.name }}</span
                >
                <span
                  class="text-[0.75rem] md:text-[1.25rem] leading-[1.375rem] text-black-core/[0.87]"
                  >{{ client.designation }}</span
                >
              </div>
              <a
                :href="clutchLink"
                target="_blank"
                @click.native="$mixpanel.track('tap_clutch_ratings')"
              >
                <div class="flex flex-row sm:justify-end w-[132px] sm:w-fit">
                  <div class="flex flex-col sm:justify-center mt-1 sm:mt-[7px]">
                    <span
                      class="sm:justify-between mr-1 md:mr-2 text-[0.5rem] md:text-[0.625rem] leading-[1.375rem] text-black-core/[0.87]"
                    >
                      REVIEWED ON
                    </span>

                    <img
                      :src="reviewImage"
                      class="mr-2 w-[3.313rem] md:w-[4.25rem] h-[0.938rem] md:h-[1.188rem]"
                      alt="reviewClutchImage"
                    />
                  </div>
                  <div class="flex flex-col mt-2.5 sm:mt-5 sm:justify-end">
                    <span class="flex flex-row">
                      <Icon
                        v-for="i in 5"
                        :key="i"
                        class="fa-star w-[11px] md:w-4 h-[11px] md:h-4 text-[#FF3D2E]"
                        name="fa6-solid:star"
                    /></span>

                    <span
                      class="md:font-inter-regular text-[0.5rem] md:text-[0.75rem] leading-[1.375rem] text-black-core/[0.87]"
                    >
                      5.0 RATINGS
                    </span>
                  </div>
                </div>
              </a>
            </div>
          </div>
        </swiper-slide>
      </swiper>
    </div>
  </div>
</template>

<script>
import { Autoplay, Pagination } from "swiper/modules";
import { Swiper, SwiperSlide } from "swiper/vue";
import reviewImage from "@/assets/images/clients/v3-client/ReviewImage.webp";
import Config from "@/config.js";
export default {
  data() {
    return {
      swiper: null,
      reading: false,
      modules: [Pagination, Autoplay],
      reviewImage: reviewImage,
      clutchLink: Config.CLUTCH_URL,
      clients: [
        {
          id: 1,
          name: "Rebecca Kimura",
          className: "2xl:px-48",
          designation: "Founder at Togness, Australia",
          review: `“There was rarely ever a second explanation needed. Even if we struggled to explain technically what we wanted, they <span class="bg-single-color-underline
            ">understood the first time.</span>”`,
        },
        {
          id: 2,
          name: "Elyass Bouchater",
          className: "2xl:px-48",
          designation: "Product Manager at Luxe, Morocco",
          review: `“The Play Store is the hardest app store to get good reviews on, and we’ve just reached a five-star rating, which has been one of our biggest achievements, partly <span class="bg-single-color-underline
            ">thanks to Canopas’ work.</span>”`,
        },
        {
          id: 3,
          name: "Rob Eberhard",
          className: "2xl:px-52",
          designation: "Founder at ActivScout, Canada",
          review: `“I was especially impressed with the skills of their backend developer and how well the project manager and she worked with one another to create a <span class="bg-single-color-underline
            ">high performing iOS app.</span>”`,
        },
        {
          id: 4,
          name: "Lisa Weinstein",
          className: "2xl:px-52",
          designation: "Founder at Brickandbatten, USA",
          review: `“There is not enough space to say all the wonderful things I would want to share about Canopas. The team is incredibly helpful, stays calm even when we had to deal with tough issues on our app and always found a way to help us fix whatever was needed or roll out any new features for our app in both the
            <span class="bg-single-color-underline
            "> iOS and Android stores.</span>”`,
        },
        {
          id: 5,
          name: "Cyril Trosset",
          className: "2xl:px-52",
          designation: "CTO at Udini, France",
          review: `“Multiple versions of this Android app have been successfully delivered over time. They are always very responsive on bug resolution. They are very efficient at producing complex interfaces and <span class="bg-single-color-underline
            ">high quality apps.</span>”`,
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
  },
  inject: ["mixpanel"],
  methods: {
    onSwiper(swiper) {
      this.swiper = swiper;
    },
    playSwiper(play) {
      play ? this.swiper.autoplay.start() : this.swiper.autoplay.stop();
      this.reading = !play;
    },
  },
};
</script>

<style lang="postcss">
@import "swiper/css";
</style>
