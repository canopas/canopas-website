<template>
  <div
    class="xll:tw-container tw-my-[1rem] lg:tw-mt-[5rem] tw-mx-auto tw-text-center"
  >
    <span
      class="tw-container tw-text-center tw-font-inter-bold tw-text-black-core/[0.8] tw-text-[1.875rem] lg:tw-text-[3.4375rem] tw-leading-[2.4375rem] lg:tw-leading-[5.15625rem]"
      >WHY and How to Write Unit Tests?</span
    >
    <div class="md:tw-hidden swiper-content tw-mt-16 tw-h-full">
      <swiper
        :slidesPerView="1.1"
        :centeredSlides="true"
        :autoplay="{
          delay: 3000,
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
              :src="test.image[0]"
              :srcset="`${test.image[0]} 400w,${test.image[1]} 800w`"
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
      <div class="swiper-content tw-mt-[3rem] tw-pl-[25%]">
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
          :navigation="true"
          @swiper="onSwiper"
          class="swiper-container"
        >
          <swiper-slide v-for="(test, index) in tests" :key="index">
            <div>
              <img
                :src="test.image[0]"
                :srcset="`${test.image[0]} 400w,${test.image[1]} 800w`"
                class="tw-w-[800px] tw-h-[450px] tw-object-contain tw-rounded-t-lg tw-border-black-core/[0.87] tw-drop-shadow-md"
                alt="UnitTestImage"
              />
            </div>
          </swiper-slide>
        </swiper>
      </div>
      <div class="skewed tw-bg-white tw-h-[100px]"></div>
    </div>
  </div>
</template>

<script>
import SwiperCore, { Pagination, Autoplay } from "swiper";
import { Swiper, SwiperSlide } from "swiper/vue";
import unittest1_400w from "@/assets/images/contributions/unittest/unit_test_1-400w.webp";
import unittest1_800w from "@/assets/images/contributions/unittest/unit_test_1-800w.webp";
import unittest2_400w from "@/assets/images/contributions/unittest/unit_test_2-400w.webp";
import unittest2_800w from "@/assets/images/contributions/unittest/unit_test_2-800w.webp";
import unittest3_400w from "@/assets/images/contributions/unittest/unit_test_3-400w.webp";
import unittest3_800w from "@/assets/images/contributions/unittest/unit_test_3-800w.webp";
import unittest4_400w from "@/assets/images/contributions/unittest/unit_test_4-400w.webp";
import unittest4_800w from "@/assets/images/contributions/unittest/unit_test_4-800w.webp";

SwiperCore.use([Pagination, Autoplay]);

export default {
  data() {
    return {
      swiper: null,
      tests: [
        {
          image: [unittest1_400w, unittest1_800w],
          title:
            "The Ultimate Guide to iOS unit testing with best practices — Part 1",
        },
        {
          image: [unittest2_400w, unittest2_800w],
          title:
            "The Ultimate Guide to iOS unit testing with best practices — Part 1 ",
        },

        {
          image: [unittest3_400w, unittest3_800w],
          title:
            "The Ultimate Guide to iOS unit testing with best practices — Part 1 ",
        },
        {
          image: [unittest4_400w, unittest4_800w],
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
  @apply tw-opacity-40 tw--mt-[400px];
}
@screen md {
  .swiper-slide:not(.swiper-slide-active) {
    @apply !tw-opacity-100 tw-mt-0;
  }
}
.skewed {
  transform-origin: bottom right;
  transform: skew(-87deg);
}
</style>
