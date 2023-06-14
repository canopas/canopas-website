<template>
  <div class="tw-my-[1rem] lg:tw-mt-[5rem] tw-mx-auto tw-text-center">
    <span
      class="tw-container tw-text-center tw-font-inter-bold tw-text-black-core/[0.8] tw-text-[1.875rem] lg:tw-text-[3.4375rem] tw-leading-[2.4375rem] lg:tw-leading-[5.15625rem]"
      >WHY and How to Write Unit Tests?</span
    >
    <div class="md:tw-hidden swiper-content tw-mt-16 tw-h-full">
      <swiper
        :slidesPerView="1.1"
        :centeredSlides="true"
        :autoplay="{
          delay: 4000,
          disableOnInteraction: true,
        }"
        :loop="true"
        :loopedSlides="50"
        :spaceBetween="30"
        :pagination="pagination"
        class="swiper-container"
        @swiper="onSwiper"
      >
        <swiper-slide
          v-for="(test, index) in tests"
          :key="index"
          class="tw-cursor-pointer"
          @mouseover="playSwiper(false)"
          @mouseleave="playSwiper(true)"
          @touchstart.passive="playSwiper(false)"
          @touchmove.passive="playSwiper(true)"
          @touchend="playSwiper(true)"
        >
          <div>
            <img
              :src="test.image"
              class="tw-w-fit tw-h-fit tw-object-contain tw-border-b-[0.5px] tw-rounded-t-lg tw-border-black-core/[0.87] tw-drop-shadow-md tw-animate-slideLeft"
              alt="UnitTestImage"
            />
            <div
              class="tw-p-[20px] tw-bg-white tw-rounded-b-lg tw-drop-shadow-md tw-animate-slideDown"
            >
              <span class="tw-font-inter-medium tw-text-black-core/[0.87]">{{
                test.title
              }}</span>
            </div>
          </div>
        </swiper-slide>
      </swiper>
    </div>
    <div
      class="tw-hidden md:tw-block tw-bg-gradient-to-t tw-from-[#FF835B]/[0.1] tw-via-[#F2709C]/[0.1] tw-to-[#FFFFFF]"
    >
      <!-- <div class="swiper-content tw-mt-16 tw-ml-[17rem]">
        <swiper
          :slidesPerView="1.1"
          :centeredSlides="true"
          :autoplay="{
            delay: 3000,
            disableOnInteraction: true,
          }"
          :loop="true"
          :loopedSlides="50"
          :spaceBetween="10"
          :pagination="pagination"
          class="swiper-container"
        >
          <swiper-slide
            v-for="(test, index) in tests"
            :key="index"
            class="tw-cursor-pointer"
          >
            <div>
              <img
                :src="test.image"
                class="tw-w-[800px] tw-h-[450px] tw-object-contain tw-rounded-t-lg tw-border-black-core/[0.87] tw-drop-shadow-md"
                alt="UnitTestImage"
              />
            </div>
          </swiper-slide>
        </swiper>
      </div> -->
      <div
        class="swiper-content tw-mt-[3rem] lg:tw-w-[50%] 2xl:tw-w-[60%] tw-pl-4 md:tw-pl-0"
      >
        <swiper
          :slidesPerView="1.1"
          :centeredSlides="false"
          :autoplay="{
            delay: 2000,
            disableOnInteraction: false,
          }"
          :loop="true"
          :loopedSlides="50"
          :spaceBetween="20"
          :pagination="{ type: 'progressbar' }"
          :navigation="true"
          :breakpoints="{
            '576': {
              slidesPerView: 1.2,
            },
            '768': {
              centeredSlides: true,
              slidesPerView: 1.525,
            },
            '992': {
              centeredSlides: false,
              slidesPerView: 2.2,
            },
          }"
          @swiper="onSwiper"
          class="swiper-container"
        >
          <swiper-slide v-for="(test, index) in tests" :key="index">
            <div>
              <img
                :src="test.image"
                class="tw-w-[800px] tw-h-[450px] tw-object-contain tw-rounded-t-lg tw-border-black-core/[0.87] tw-drop-shadow-md"
                alt="UnitTestImage"
              />
            </div>
          </swiper-slide>
        </swiper>
      </div>
    </div>
  </div>
</template>

<script>
import SwiperCore, { Pagination, Autoplay } from "swiper";
import { Swiper, SwiperSlide } from "swiper/vue";
import unittest1 from "@/assets/images/ContributionPage/unittest/unit_test_1.webp";
import unittest2 from "@/assets/images/ContributionPage/unittest/unit_test_2.webp";
import unittest3 from "@/assets/images/ContributionPage/unittest/unit_test_3.webp";
import unittest4 from "@/assets/images/ContributionPage/unittest/unit_test_4.webp";
SwiperCore.use([Pagination, Autoplay]);

export default {
  data() {
    return {
      swiper: null,
      tests: [
        {
          image: unittest1,
          title:
            "The Ultimate Guide to iOS unit testing with best practices — Part 1",
        },
        {
          image: unittest2,
          title:
            "The Ultimate Guide to iOS unit testing with best practices — Part 1 ",
        },

        {
          image: unittest3,
          title:
            "The Ultimate Guide to iOS unit testing with best practices — Part 1 ",
        },
        {
          image: unittest4,
          title:
            "The Ultimate Guide to iOS unit testing with best practices — Part 1 ",
        },
      ],
      pagination: {
        el: ".swiper-pagination",
        clickable: true,
      },
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
  },

  components: {
    Swiper,
    SwiperSlide,
  },
  inject: ["mixpanel"],
};
</script>

<style lang="postcss" scoped>
@import "swiper/css";

.swiper-wrapper {
  @apply !tw-items-center;
}

.swiper-slide:not(.swiper-slide-active) {
  @apply md:tw--mt-[50px] tw-opacity-40 tw--mt-[400px];
}
</style>
