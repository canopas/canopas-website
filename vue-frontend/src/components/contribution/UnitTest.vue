<template>
  <div
    class="tw-container tw-my-[1rem] lg:tw-mt-[5rem] tw-mx-auto tw-text-center"
  >
    <span
      class="tw-text-center tw-font-inter-bold tw-text-black-core/[0.8] tw-text-[3.4375rem] tw-leading-[5.15625rem]"
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
        :spaceBetween="20"
        :pagination="pagination"
        @swiper="onSwiper"
        class="swiper-container"
      >
        <swiper-slide
          v-for="image in images"
          :key="image"
          class="tw-cursor-pointer"
        >
          <img
            :src="image"
            class="tw-w-fit tw-h-fit tw-object-contain"
            alt="UnitTestImage"
          />
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
import Config from "@/config.js";
export default {
  data() {
    return {
      swiper: null,
      images: [unittest1, unittest2, unittest3, unittest4],
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
