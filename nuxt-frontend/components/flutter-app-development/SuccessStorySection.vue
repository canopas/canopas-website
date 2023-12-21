<template>
  <div
    class="bg-black-core text-white text-center h-[600px] lg:h-[750px] 3xl:outer-container"
  >
    <div
      class="container flex gap-6 lg:gap-[4.5rem] flex-col pt-12 xs:pt-14 sm:pt-[4.5rem] lg:!pt-[7.5rem] xl:!pt-24 lg:px-48"
    >
      <div class="mobile-header-2 lg:desk-header-2">Our Success Stories</div>
      <div class="flex flex-col gap-6 lg:gap-10">
        <img
          :src="quoteimage"
          class="mx-auto w-10 lg:w-14 h-full object-contain"
          alt="quote-image"
        />
        <div class="swiper-content">
          <swiper
            :slidesPerView="1"
            :autoplay="{
              delay: 3000,
              disableOnInteraction: false,
            }"
            :speed="1500"
            :loop="true"
            :spaceBetween="20"
            :pagination="pagination"
            :modules="modules"
            class="swiper-container"
          >
            <swiper-slide v-for="(client, index) in clients" :key="index">
              <div class="flex flex-col gap-6 lg:gap-10">
                <div class="sub-h1-semibold lg:mobile-header-2-semibold">
                  "{{ client.review }}"
                </div>
                <div class="flex flex-col">
                  <span
                    class="sub-h3-semibold lg:mobile-header-3-semibold text-white-core-80"
                    >{{ client.name }}</span
                  >
                  <span
                    class="sub-h3-regular lg:mobile-header-3-regular text-white-core-60"
                    >{{ client.designation }}</span
                  >
                </div>
              </div>
            </swiper-slide>
            <div class="swiper-pagination"></div>
          </swiper>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup>
import { Autoplay, Pagination } from "swiper/modules";
import { Swiper, SwiperSlide } from "swiper/vue";
import quoteimage from "@/assets/images/flutter-app-development/quote.webp";
const modules = [Pagination, Autoplay];
const clients = [
  {
    id: 1,
    name: "Rebecca Kimura",
    designation: "Founder at Togness, Australia",
    review:
      "There was rarely ever a second explanation needed. Even if we struggled to explain technically what we wanted, they understood the first time.",
  },
  {
    id: 2,
    name: "Elyass Bouchater",
    designation: "Product Manager at Luxe, Morocco",
    review:
      "The Play Store is the hardest app store to get good reviews on, and we've just reached a five-star rating, which has been one of our biggest achievements, partly thanks to Canopas' work.",
  },
  {
    id: 3,
    name: "Rob Eberhard",
    designation: "Founder at ActivScout, Canada",
    review:
      "I was especially impressed with the skills of their backend developer and how well the project manager and she worked with one another to create a  high performing iOS app.",
  },
  {
    id: 4,
    name: "Lisa Weinstein",
    designation: "Founder at Brickandbatten, USA",
    review:
      "There is not enough space to say all the wonderful things I would want to share about Canopas. The team is incredibly helpful, stays calm even when we had to deal with tough issues on our app and always found a way to help us fix whatever was needed or roll out any new features for our app in both the iOS and Android stores.",
  },
  {
    id: 5,
    name: "Cyril Trosset",
    designation: "CTO at Udini, France",
    review:
      "Multiple versions of this Android app have been successfully delivered over time. They are always very responsive on bug resolution. They are very efficient at producing complex interfaces and  high quality apps.",
  },
];
const pagination = {
  el: ".swiper-pagination",
  clickable: true,
};
</script>
<style lang="postcss">
@import "swiper/css";
@import "swiper/css/pagination";
.swiper-pagination {
  @apply static pt-6 pl-0;
}
.swiper-pagination-bullet {
  @apply !w-3 !h-3 !rounded-full !bg-pink-90 !border !border-solid !border-white-core-80;
}
.swiper-pagination-bullet-active {
  @apply !border-none !opacity-100;
}
</style>
