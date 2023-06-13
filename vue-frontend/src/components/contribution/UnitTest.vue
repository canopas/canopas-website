<template>
  <div
    class="tw-container tw-my-[1rem] lg:tw-mt-[5rem] tw-mx-auto tw-text-center"
  >
    <span
      class="tw-text-center tw-font-inter-bold tw-text-black-core/[0.8] tw-text-[1.875rem] lg:tw-text-[3.4375rem] tw-leading-[2.4375rem] lg:tw-leading-[5.15625rem]"
      >WHY and How to Write Unit Tests?</span
    >
    <div class="swiper-content tw-mt-16">
      <swiper
        :slidesPerView="1"
        :centeredSlides="true"
        :autoplay="{
          delay: 3000,
          disableOnInteraction: false,
        }"
        :loop="true"
        :loopedSlides="50"
        :spaceBetween="10"
        :pagination="pagination"
        class="swiper-container"
      >
        <swiper-slide
          v-for="test in tests"
          :key="test"
          class="tw-cursor-pointer"
        >
          <div
            class="tw-w-[350px] tw-h-[286px] lg:tw-hidden tw-rounded-b-lg tw-drop-shadow-md tw-bg-white"
          >
            <img
              :src="test.image"
              class="tw-w-fit tw-h-fit tw-object-contain tw-border-b-[0.5px] tw-border-black-core/[0.87] tw-rounded-t-lg"
              alt="UnitTestImage"
            />
            <span class="tw-font-inter-medium tw-text-black-core/[0.87]">{{
              test.title
            }}</span>
          </div>
        </swiper-slide>
      </swiper>
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
    };
  },
  components: {
    Swiper,
    SwiperSlide,
  },
  inject: ["mixpanel"],
};
</script>

<style lang="postcss">
@import "swiper/css";

.swiper-wrapper {
  @apply !tw-items-center;
}
</style>
