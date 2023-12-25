<template>
  <section class="my-16 lg:my-60 3xl:outer-container">
    <!-- Mobile UI start -->
    <div class="lg:hidden">
      <div class="mobile-header-2 text-center text-black-87 mb-6">
        Our success stories
      </div>
      <swiper
        :slidesPerView="1.1"
        :autoplay="{
          delay: 3000,
          disableOnInteraction: false,
        }"
        :spaceBetween="0"
        :centeredSlides="true"
        :speed="1500"
        :loop="true"
        :modules="modules"
        class="swiper-container"
        :breakpoints="{
          '420': {
            slidesPerView: 1.2,
          },
          '520': {
            slidesPerView: 1.5,
          },
          '630': {
            slidesPerView: 1.8,
          },
          '768': {
            slidesPerView: 2,
          },
        }"
      >
        <swiper-slide
          v-for="client in clients"
          :key="client"
          class="bg-deep-charcoal !scale-90 transition-all duration-200 rounded-xl text-left ease-out"
        >
          <div class="container py-6 xs:px-6">
            <img
              :src="imageleft"
              class="w-6 h-6 object-contain"
              loading="lazy"
              alt="left-quote"
            />
            <div class="mt-6">
              <div
                class="flex flex-col justify-between h-[25.25rem] xs:h-[22.25rem] sm:h-[19.25rem]"
              >
                <p
                  class="sub-h1-semibold text-white"
                  v-html="client.review"
                ></p>
                <div class="flex flex-col">
                  <p class="sub-h3-semibold text-white-core-80">
                    {{ client.name }}
                  </p>
                  <p class="sub-h3-regular text-white-core-60">
                    {{ client.designation }}
                  </p>
                </div>
              </div>
            </div>
          </div>
        </swiper-slide>
      </swiper>
    </div>
    <!-- Mobile UI END -->
    <!-- Desktop UI -->
    <div class="hidden lg:block xll:container">
      <div class="desk-header-2 text-center text-black-87 mb-[4.5rem]">
        Success Stories
      </div>
      <swiper
        :slidesPerView="1.5"
        :autoplay="{
          delay: 3000,
          disableOnInteraction: false,
        }"
        :centeredSlides="true"
        :speed="1500"
        :loop="true"
        :modules="modules"
        class="!p-8"
        :breakpoints="{
          '992': {
            slidesPerView: 1.5,
          },
          '1200': {
            slidesPerView: 1.7,
          },
          '1400': {
            slidesPerView: 2,
          },
          '2400': {
            slidesPerView: 2,
          },
        }"
      >
        <swiper-slide
          v-for="client in clients"
          :key="client"
          class="slide text-white bg-deep-charcoal !scale-75 transition-all duration-200 rounded-xl text-left ease-in"
          ><div class="container py-14 mx-14">
            <img
              :src="imageleft"
              class="w-6 h-6 object-contain"
              loading="lazy"
              alt="quote-left"
            />
            <div
              class="flex flex-col !justify-between mt-6 h-[25.26rem] 2xll:h-[20.25rem] xll:h-[25.26rem]"
            >
              <div class="mobile-header-2-semibold">{{ client.review }}</div>
              <div class="flex flex-col">
                <p class="sub-h2-medium text-white-core-80">
                  {{ client.name }}
                </p>
                <p class="mobile-header-3-regular text-white-core-60">
                  {{ client.designation }}
                </p>
              </div>
            </div>
          </div></swiper-slide
        >
      </swiper>
    </div>
    <!-- Desktop UI END-->
  </section>
</template>
<script setup>
import { Autoplay } from "swiper/modules";
import { Swiper, SwiperSlide } from "swiper/vue";
import imageleft from "@/assets/images/backend-development/quote-left.webp";
const modules = [Autoplay];
const clients = [
  {
    id: 1,
    name: "Rebecca Kimura",
    designation: "Founder at Togness, Australia",
    review: `"There was rarely ever a second explanation needed. Even if we struggled to explain technically what we wanted, they understood the first time."`,
  },
  {
    id: 2,
    name: "Elyass Bouchater",
    designation: "Product Manager at Luxe, Morocco",
    review: `"The Play Store is the hardest app store to get good reviews on, and we've just reached a five-star rating, which has been one of our biggest achievements, partly thanks to Canopas' work."`,
  },
  {
    id: 3,
    name: "Rob Eberhard",
    designation: "Founder at ActivScout, Canada",
    review: `"I was especially impressed with the skills of their backend developer and how well the project manager and she worked with one another to create a  high performing iOS app."`,
  },
  {
    id: 4,
    name: "Lisa Weinstein",
    designation: "Founder at Brickandbatten, USA",
    review: `"There is not enough space to say all the wonderful things I would want to share about Canopas. The team is incredibly helpful, stays calm even when we had to deal with tough issues on our app and always found a way to help us fix whatever was needed or roll out any new features for our app in both the iOS and Android stores."`,
  },
  {
    id: 5,
    name: "Cyril Trosset",
    designation: "CTO at Udini, France",
    review: `"Multiple versions of this Android app have been successfully delivered over time. They are always very responsive on bug resolution. They are very efficient at producing complex interfaces and  high quality apps."`,
  },
];
</script>
<style lang="postcss" scoped>
.swiper-slide-active {
  @apply !scale-100;
}
.slide {
  box-shadow: 1.5rem 1.5rem #ff9472 !important;
}
</style>
