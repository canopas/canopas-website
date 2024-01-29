<template>
  <div class="mt-16 mx-auto lg:my-60">
    <div
      class="container text-center mobile-header-2 lg:desk-header-2 text-black-87"
    >
      WHY and How to write unit tests?
    </div>
    <div class="swiper-content mt-8 block h-full pl-[5%] md:hidden">
      <swiper
        :slidesPerView="1.1"
        :centeredSlides="false"
        :autoplay="{
          delay: 4000,
        }"
        :speed="3000"
        :loop="true"
        :spaceBetween="10"
        :modules="modules"
        class="swiper-container"
      >
        <swiper-slide
          v-for="(test, index) in unitTests"
          :key="index"
          class="cursor-pointer"
        >
          <div
            class="cursor-pointer"
            @click="openBlog(test.url, 'tap_contribution_unit_test_section')"
          >
            <img
              :src="test.images[0]"
              :srcset="`${test.images[0]} 400w,${test.images[1]} 800w`"
              class="cover-image h-fit w-fit rounded-t-lg border-b-[0.5px] border-black-8 object-contain drop-shadow-md"
              alt="UnitTestImage"
              loading="lazy"
            />
            <div class="content mb-8 rounded-b-lg bg-white p-5 drop-shadow-md">
              <span class="sub-h1-semibold text-black-87">{{
                test.title
              }}</span>
            </div>
          </div>
        </swiper-slide>
      </swiper>
    </div>
    <!-- ************************Desktop View****************************** -->
    <div
      class="mt-12 hidden bg-gradient-to-t from-orange-1 via-pink-1 to-white md:block xl:mt-16 3xl:outer-container"
    >
      <div class="swiper-content relative -mt-6 xll:container">
        <div
          class="absolute left-0 top-[1.125rem] z-0 mx-auto h-full w-full border-t-[0.5px] border-black-4"
        ></div>

        <swiper
          :slidesPerView="3"
          :centeredSlides="true"
          :loop="true"
          :autoplay="{
            delay: 2000,
            disableOnInteraction: false,
          }"
          :speed="3000"
          :spaceBetween="20"
          :zoom="true"
          :modules="modules"
          class="swiper-container absolute left-0 top-2.5 z-30 mx-auto h-full w-full"
        >
          <swiper-slide v-for="(test, index) in unitTests" :key="index">
            <div
              class="cursor-pointer"
              @click="openBlog(test.url, 'tap_contribution_unit_test_section')"
            >
              <img
                :src="dot"
                alt="dot-image"
                class="mx-auto h-[1.625rem] w-[1.625rem] object-contain"
              />

              <img
                :src="test.images[0]"
                :srcset="`${test.images[0]} 400w,${test.images[1]} 800w`"
                class="image-content swiper-zoom-container h-4/5 w-fit object-cover drop-shadow-md lg:h-fit"
                alt="UnitTestImage"
                loading="lazy"
              />
            </div>
          </swiper-slide>
        </swiper>
      </div>
      <div
        class="h-[6.25rem] origin-bottom-right skew-x-[-87deg] bg-white lg:h-[8.438rem] xll:h-[13.75rem]"
      ></div>
    </div>
  </div>
</template>

<script>
import { Zoom, Autoplay } from "swiper/modules";
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

export default {
  data() {
    return {
      modules: [Zoom, Autoplay],
      openBlog,
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
  @apply scale-[0.8];
}
.swiper-slide .swiper-zoom-container {
  @apply origin-bottom duration-[3000ms];
}
.swiper-wrapper {
  @apply !items-center;
}
.swiper-slide-active .content {
  @apply animate-slideDown;
}
.swiper-slide-active .cover-image {
  @apply animate-fadeSlide;
}
.swiper-slide-next {
  @apply -mt-[100%];
}
.swiper-slide-prev .content {
  @apply animate-fadeUp opacity-40;
}
.swiper-slide-prev .cover-image {
  @apply animate-fadeSlideOut;
}
@screen md {
  .swiper-slide:not(.swiper-slide-active) {
    @apply mt-0 !opacity-100;
  }
}
</style>
