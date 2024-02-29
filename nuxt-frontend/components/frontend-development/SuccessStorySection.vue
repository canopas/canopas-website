<template>
  <section class="my-16 lg:my-60">
    <!-- Mobile UI start -->
    <div class="lg:hidden text-center">
      <h2 class="mobile-header-2 text-black-87">Success stories</h2>
      <div
        class="mt-6 h-[33rem] rounded-xl relative overflow-hidden bg-eerie-black mx-4"
      >
        <img
          :src="reviewBackgroundImage"
          :srcset="`${reviewBackgroundImage} 400w`"
          class="w-full h-full object-scale-down rounded-xl"
          alt="success-story-frontend-image"
          loading="lazy"
        />
        <div
          class="container absolute w-full p-4 text-white transform left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2"
        >
          <swiper
            :slidesPerView="1"
            :centeredSlides="true"
            :autoplay="{
              delay: 3000,
              disableOnInteraction: false,
            }"
            :speed="3000"
            :loop="true"
            :spaceBetween="20"
            :modules="modules"
            class="swiper-container"
          >
            <swiper-slide
              v-for="(client, index) in clients"
              :key="index"
              class="cursor-pointer"
            >
              <div class="mx-auto text-white flex flex-col text-center">
                <div class="transition-all ease duration-500 sub-h1-semibold">
                  {{ client.review }}
                </div>
                <div class="flex flex-col mt-8 mb-2.5">
                  <span class="flex justify-center">
                    <Icon
                      v-for="i in 5"
                      :key="i"
                      class="fa-sharp w-5 h-5 text-yellow"
                      name="fa6-solid:star"
                  /></span>
                </div>
                <div class="flex flex-col">
                  <span class="sub-h3-medium text-white-core-80">{{
                    client.name
                  }}</span>
                  <span class="sub-h4-regular text-white-core-60">{{
                    client.designation
                  }}</span>
                </div>
              </div>
            </swiper-slide>
          </swiper>
        </div>
      </div>
    </div>
    <!-- Mobile UI END -->
    <!-- Desktop UI -->
    <div class="hidden lg:flex container w-[69.25rem] h-[41.25rem]">
      <div
        class="!w-[23.5rem] bg-deep-charcoal text-white flex items-center rounded-tl-xl rounded-bl-xl"
      >
        <h2 class="desk-header-2 mx-auto">
          What our <br />
          clients says?
        </h2>
      </div>
      <div
        class="w-[38.75rem] xl:w-[45.75rem] flex flex-col gap-10 rounded-tr-xl rounded-br-xl shadow-3xl py-14 px-12"
      >
        <div>
          <img
            :src="imageleft"
            alt="quote-left-image"
            loading="lazy"
            class="w-8 h-7"
          />
        </div>
        <div class="my-auto">
          <swiper
            :slidesPerView="1"
            :speed="1000"
            :loop="true"
            :spaceBetween="20"
            :autoplay="{
              delay: 3000,
              disableOnInteraction: false,
            }"
            :pagination="pagination"
            :modules="modules"
            class="swiper-container cursor-pointer"
          >
            <swiper-slide v-for="(client, index) in clients" :key="index">
              <div class="my-auto text-black-87">
                <div class="mb-4">
                  <span class="flex">
                    <Icon
                      v-for="i in 5"
                      :key="i"
                      class="fa-star-sharp w-[1.375rem] h-[1.375rem] lg:w-6 lg:h-6 text-yellow"
                      name="fa6-solid:star"
                  /></span>
                </div>
                <div class="flex flex-col gap-6">
                  <div class="mobile-header-2-semibold">
                    {{ client.review }}
                  </div>
                  <div class="flex flex-col">
                    <span class="text-black-60 sub-h2-medium">{{
                      client.name
                    }}</span>
                    <span class="text-black-4 sub-h1-regular">{{
                      client.designation
                    }}</span>
                  </div>
                </div>
              </div>
            </swiper-slide>
          </swiper>
          <div class="flex justify-between mt-16 relative flex-row-reverse">
            <div class="swiper-pagination"></div>
            <img
              :src="imageright"
              alt="quote-left-image"
              loading="lazy"
              class="w-8 h-7"
            />
          </div>
        </div>
      </div>
      <div></div>
    </div>
    <!-- Desktop UI END-->
  </section>
</template>
<script setup>
import { Autoplay, Pagination } from "swiper/modules";
import { Swiper, SwiperSlide } from "swiper/vue";
import imageleft from "@/assets/images/frontend-development/quoteleft.webp";
import imageright from "@/assets/images/frontend-development/quoteright.webp";
import reviewBackgroundImage from "@/assets/images/frontend-development/review-background-400w.webp";

const modules = [Autoplay, Pagination];
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
<style lang="postcss" scoped>
@import "swiper/css";
@import "swiper/css/pagination";
.swiper-wrapper {
  @apply lg:items-center;
}
.swiper-pagination {
  @apply text-left;
}
.swiper-pagination-bullet {
  @apply !w-3 !h-3 !rounded-full !bg-white-smoke !border !border-solid !border-black-87;
}
.swiper-pagination-bullet-active {
  @apply !border-none !bg-black-87;
}
</style>
