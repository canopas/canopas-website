<template>
  <div class="tw-mx-auto tw-text-center lg:tw-my-[14.063rem]">
    <span
      class="tw-container tw-text-center tw-font-inter-bold tw-text-[1.875rem] tw-leading-[2.4375rem] tw-text-black-core/[0.8] lg:tw-text-[3.4375rem] lg:tw-leading-[5.15625rem]"
      >WHY and How to Write Unit Tests?</span
    >
    <div
      class="swiper-content tw-mt-16 tw-block tw-h-full tw-pl-[5%] md:tw-hidden"
    >
      <swiper
        :slidesPerView="1.1"
        :centeredSlides="false"
        :autoplay="{
          delay: 4000,
          disableOnInteraction: true,
        }"
        :speed="3000"
        :loop="true"
        :loopedSlides="50"
        :spaceBetween="10"
        class="swiper-container"
      >
        <swiper-slide
          v-for="(test, index) in unitTests"
          :key="index"
          class="tw-cursor-pointer"
        >
          <div
            class="tw-cursor-pointer"
            @click="openBlog(test.url, 'tap_contribution_unit_test_section')"
          >
            <img
              :src="test.images[0]"
              :srcset="`${test.images[0]} 400w,${test.images[1]} 800w`"
              class="cover-image tw-h-fit tw-w-fit tw-rounded-t-lg tw-border-b-[0.5px] tw-border-black-core/[0.87] tw-object-contain tw-drop-shadow-md"
              alt="UnitTestImage"
            />
            <div
              class="content tw-mb-8 tw-rounded-b-lg tw-bg-white tw-p-5 tw-drop-shadow-md"
            >
              <span class="tw-font-inter-medium tw-text-black-core/[0.87]">{{
                test.title
              }}</span>
            </div>
          </div>
        </swiper-slide>
      </swiper>
    </div>
    <!-- ************************Desktop View****************************** -->
    <div
      class="tw-mt-12 tw-hidden tw-bg-gradient-to-t tw-from-[#FF835B]/[0.1] tw-via-[#F2709C]/[0.1] tw-to-[#FFFFFF] md:tw-block xl:tw-mt-16"
    >
      <div class="swiper-content tw-relative tw--mt-6 xll:tw-container">
        <div
          class="tw-absolute tw-left-0 tw-top-[18px] tw-z-0 tw-mx-auto tw-h-full tw-w-full tw-border-t-[0.5px] tw-border-black-core/[0.4]"
        ></div>

        <swiper
          :slidesPerView="3"
          :centeredSlides="true"
          :loop="true"
          :loopedSlides="50"
          :autoplay="{
            delay: 2000,
            disableOnInteraction: false,
          }"
          :speed="3000"
          :spaceBetween="20"
          :zoom="true"
          :navigation="true"
          class="swiper-container tw-absolute tw-left-0 tw-top-2.5 tw-z-30 tw-mx-auto tw-h-full tw-w-full"
        >
          <swiper-slide v-for="(test, index) in unitTests" :key="index">
            <div
              class="tw-cursor-pointer"
              @click="openBlog(test.url, 'tap_contribution_unit_test_section')"
            >
              <img
                :src="dot"
                alt="dot-image"
                class="tw-mx-auto tw-h-[26px] tw-w-[26px] tw-object-contain"
              />

              <img
                :src="test.images[0]"
                :srcset="`${test.images[0]} 400w,${test.images[1]} 800w`"
                class="image-content swiper-zoom-container tw-h-[80%] tw-w-fit tw-object-cover tw-drop-shadow-md lg:tw-h-fit"
                alt="UnitTestImage"
              />
            </div>
          </swiper-slide>
        </swiper>
      </div>
      <div
        class="tw-h-[100px] tw-origin-bottom-right tw-skew-x-[-87deg] tw-bg-white lg:tw-h-[135px] xll:tw-h-[220px]"
      ></div>
    </div>
  </div>
</template>

<script>
import SwiperCore, { Zoom, Autoplay } from "swiper";
import { Swiper, SwiperSlide } from "swiper/vue";
import unittest1_400w from "@/assets/images/contributions/unittest/unit_test_1-400w.webp";
import unittest1_800w from "@/assets/images/contributions/unittest/unit_test_1-800w.webp";
import unittest2_400w from "@/assets/images/contributions/unittest/unit_test_2-400w.webp";
import unittest2_800w from "@/assets/images/contributions/unittest/unit_test_2-800w.webp";
import unittest3_400w from "@/assets/images/contributions/unittest/unit_test_3-400w.webp";
import unittest3_800w from "@/assets/images/contributions/unittest/unit_test_3-800w.webp";
import unittest4_400w from "@/assets/images/contributions/unittest/unit_test_4-400w.webp";
import unittest4_800w from "@/assets/images/contributions/unittest/unit_test_4-800w.webp";
import unittest5_400w from "@/assets/images/contributions/unittest/unit_test_5-400w.webp";
import unittest5_800w from "@/assets/images/contributions/unittest/unit_test_5-800w.webp";
import unittest6_400w from "@/assets/images/contributions/unittest/unit_test_6-400w.webp";
import unittest6_800w from "@/assets/images/contributions/unittest/unit_test_6-800w.webp";
import unittest7_400w from "@/assets/images/contributions/unittest/unit_test_7-400w.webp";
import unittest7_800w from "@/assets/images/contributions/unittest/unit_test_7-800w.webp";
import dot from "@/assets/images/why/why-canopas-dot-3.svg";
import { openBlog } from "@/utils.js";

SwiperCore.use([Zoom, Autoplay]);

export default {
  data() {
    return {
      openBlog,
      swiper: null,
      dot,
      unitTests: [
        {
          images: [unittest1_400w, unittest1_800w],
          title: "1 min guide to iOS development best practices in 2022",
          url: "https://blog.canopas.com/the-ultimate-guide-to-ios-unit-testing-with-best-practices-part-1-c3a484b8dc8d",
        },
        {
          images: [unittest2_400w, unittest2_800w],
          title:
            "The Ultimate Guide to iOS unit testing with best practices — Part 1",
          url: "https://blog.canopas.com/introduction-to-unit-testing-viewmodels-with-kotlin-coroutine-367630e95831",
        },

        {
          images: [unittest3_400w, unittest3_800w],
          title: "Vue 3 component testing with Jest?",
          url: "https://blog.canopas.com/vue-3-component-testing-with-jest-8b80a8a8946b",
        },
        {
          images: [unittest4_400w, unittest4_800w],
          title: "Swift — Essential tips for writing testable code",
          url: "https://blog.canopas.com/swift-essential-tips-to-know-for-writing-testable-code-65287584d83c",
        },
        {
          images: [unittest5_400w, unittest5_800w],
          title:
            "Developing a Todo app with android integration testing following the TDD approach — Part 3 of 3",
          url: "https://blog.canopas.com/how-to-use-tdd-effectively-to-write-android-integration-tests-5efc7ee48ba8",
        },
        {
          images: [unittest6_400w, unittest6_800w],
          title: "Unit testing in go with MySQL",
          url: "https://blog.canopas.com/unit-testing-in-go-with-mysql-8db449d86722",
        },
        {
          images: [unittest7_400w, unittest7_800w],
          title: "Golang: Unit tests with test gin context",
          url: "https://blog.canopas.com/golang-unit-tests-with-test-gin-context-80e1ac04adcd",
        },
      ],
    };
  },

  components: {
    Swiper,
    SwiperSlide,
  },
};
</script>

<style lang="postcss" scoped>
@import "swiper/css";
.swiper-slide:not(.swiper-slide-active) .image-content {
  @apply tw-scale-[0.8];
}
.swiper-slide .swiper-zoom-container {
  @apply tw-origin-bottom tw-duration-[3000ms];
}
.swiper-wrapper {
  @apply !tw-items-center;
}
.swiper-slide-active .content {
  @apply tw-animate-slideDown;
}
.swiper-slide-active .cover-image {
  @apply tw-animate-fadeSlide;
}
.swiper-slide-next {
  @apply tw--mt-[100%];
}
.swiper-slide-prev .content {
  @apply tw-animate-fadeUp tw-opacity-40;
}
.swiper-slide-prev .cover-image {
  @apply tw-animate-fadeSlideOut;
}
@screen md {
  .swiper-slide:not(.swiper-slide-active) {
    @apply tw-mt-0 !tw-opacity-100;
  }
}
</style>
