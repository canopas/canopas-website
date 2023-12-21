<template>
  <div class="container mt-16 lg:mt-60 ml-auto mx-auto">
    <div class="swiper-content">
      <swiper
        :slidesPerView="1"
        :centeredSlides="true"
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
            class="mx-auto sm:px-5 md:px-[5.8rem] lg:px-8 xl:px-24 text-black-87"
            :class="[client.className]"
          >
            <a
              class="md:hidden"
              :href="clutchLink"
              target="_blank"
              @click.native="$mixpanel.track('tap_clutch_ratings')"
            >
              <div class="flex flex-row sm:justify-end w-full sm:w-fit">
                <div class="flex flex-col sm:justify-center mt-1 sm:mt-[7px]">
                  <span class="sm:justify-between sub-h4-regular">
                    REVIEWED ON
                  </span>

                  <img
                    :src="reviewImage"
                    class="mr-2 w-[5.0625rem] h-[1.4375rem]"
                    alt="reviewClutchImage"
                  />
                </div>
                <div class="flex flex-col mt-2.5 sm:mt-5 sm:justify-end">
                  <span class="flex flex-row">
                    <Icon
                      v-for="i in 5"
                      :key="i"
                      class="fa-star w-[12px] h-[12px] text-[#FFC700]"
                      name="fa6-solid:star"
                  /></span>

                  <span class="sub-h4-regular"> 5.0 RATINGS </span>
                </div>
              </div>
            </a>
            <div
              class="mt-8 md:mt-0 sub-h1-semibold lg:desk-header-3 transition-all ease duration-500"
              :class="reading ? 'scale-[0.98]' : ''"
              v-html="client.review"
            ></div>
            <div class="flex flex-row justify-between mt-8 md:mt-20">
              <div class="flex flex-col">
                <span
                  class="sub-h3-regular lg:mobile-header-3-regular xl:sub-h2-medium"
                  >{{ client.name }}</span
                >
                <span
                  class="sub-h4-regular lg:mobile-header-3-regular text-black-60"
                  >{{ client.designation }}</span
                >
              </div>
              <a
                class="hidden md:block"
                :href="clutchLink"
                target="_blank"
                @click.native="$mixpanel.track('tap_clutch_ratings')"
              >
                <div class="flex flex-row sm:justify-end w-[166px]">
                  <div class="flex flex-col sm:justify-center mt-1">
                    <span class="sm:justify-between mr-2 sub-h4-medium">
                      REVIEWED ON
                    </span>

                    <img
                      :src="reviewImage"
                      class="mr-2 w-[5.3125rem] h-[1.5rem]"
                      alt="reviewClutchImage"
                    />
                  </div>
                  <div class="flex flex-col mt-2.5 sm:mt-5 sm:justify-end">
                    <span class="flex flex-row">
                      <Icon
                        v-for="i in 5"
                        :key="i"
                        class="fa-star w-[14px] h-[14px] text-[#FFC700]"
                        name="fa6-solid:star"
                    /></span>

                    <span class="sub-h4-regular"> 5.0 RATINGS </span>
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
          review: `“There was rarely ever a second explanation needed. Even if we struggled to explain technically what we wanted, they <span class="lg:bg-single-color-underline
            ">understood the first time.</span>”`,
        },
        {
          id: 2,
          name: "Elyass Bouchater",
          className: "2xl:px-48",
          designation: "Product Manager at Luxe, Morocco",
          review: `“The Play Store is the hardest app store to get good reviews on, and we’ve just reached a five-star rating, which has been one of our biggest achievements, partly <span class="lg:bg-single-color-underline
            ">thanks to Canopas’ work.</span>”`,
        },
        {
          id: 3,
          name: "Rob Eberhard",
          className: "2xl:px-52",
          designation: "Founder at ActivScout, Canada",
          review: `“I was especially impressed with the skills of their backend developer and how well the project manager and she worked with one another to create a <span class="lg:bg-single-color-underline
            ">high performing iOS app.</span>”`,
        },
        {
          id: 4,
          name: "Lisa Weinstein",
          className: "2xl:px-52",
          designation: "Founder at Brickandbatten, USA",
          review: `“There is not enough space to say all the wonderful things I would want to share about Canopas. The team is incredibly helpful, stays calm even when we had to deal with tough issues on our app and always found a way to help us fix whatever was needed or roll out any new features for our app in both the
            <span class="lg:bg-single-color-underline
            "> iOS and Android stores.</span>”`,
        },
        {
          id: 5,
          name: "Cyril Trosset",
          className: "2xl:px-52",
          designation: "CTO at Udini, France",
          review: `“Multiple versions of this Android app have been successfully delivered over time. They are always very responsive on bug resolution. They are very efficient at producing complex interfaces and <span class="lg:bg-single-color-underline
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
