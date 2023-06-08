<template>
  <section class="tw-my-[100px]">
    <p
      class="tw-container md:tw-hidden tw-text-center tw-text-[1.875rem] lg:tw-text-[3.438rem] tw-leading-[2.438rem] lg:tw-leading-[5.156rem] tw-font-inter-bold tw-text-black-core/[0.87]"
    >
      The Favorites You <br />
      Can't Miss
    </p>
    <div class="tw-block md:tw-hidden tw-mt-8">
      <swiper :slidesPerView="1.2" :spaceBetween="0" class="tw-p-[1rem]">
        <swiper-slide
          v-for="(favourite, index) in favourites"
          :key="index"
          class="tw-relative tw-flex tw-flex-col tw-pl-[1rem] sm:tw-pl-[2rem]"
        >
          <div
            class="tw-flex tw-flex-row tw-gap-2 tw-items-center tw-justify-end tw-mb-2"
          >
            <img
              :src="like_100w"
              alt="like"
              class="tw-w-7 tw-h-7 tw-object-cover"
              loading="lazy"
            /><span
              class="tw-font-inter-bold tw-text-[1rem] tw-leading-[1.25rem] v2-canopas-gradient-text"
              >{{ favourite.likes }}</span
            >
          </div>
          <div class="tw-overflow-hidden">
            <div class="tw-w-full tw-h-full">
              <img
                @click="openBlog(favourite)"
                :src="favourite.image[0]"
                :srcset="`${favourite.image[0]} 400w, ${favourite.image[1]} 800w`"
                alt="favourite"
                class="tw-w-full tw-h-full tw-object-cover tw-rounded-[10px]"
                loading="lazy"
              />
            </div>
            <p
              class="tw-pl-2 tw-text-[1rem] sm:tw-text-[1.25rem] tw-mt-2 tw-leading-[1.5rem] tw-font-inter-medium tw-text-black-core/[0.87]"
            >
              {{ favourite.description }}
            </p>
          </div>
        </swiper-slide>
      </swiper>
    </div>

    <p
      class="tw-container tw-hidden md:tw-block tw-font-inter-bold tw-text-[2rem] lg:tw-text-[3rem] tw-leading-[3rem] lg:tw-leading-[4rem] tw-text-black-core/[0.87]"
    >
      The Favorites You <br />
      Can't Miss
    </p>
    <div class="tw-container tw-hidden md:tw-flex tw-flex-row">
      <div
        class="tw-flex tw-flex-col tw-mt-[7.438rem]"
        v-for="favourite in favourites.slice(0, 1)"
        :key="favourite.id"
      >
        <div
          class="tw-flex tw-flex-row tw-gap-2 tw-items-center tw-justify-end tw-mb-2 tw-mr-6 lg:tw-mr-[2rem] xl:tw-mr-[2.5rem] 2xl:tw-mr-[3rem]"
        >
          <img
            :src="like_100w"
            alt="like"
            class="tw-w-7 tw-h-7 tw-object-cover"
            loading="lazy"
          /><span
            class="tw-font-inter-bold tw-text-[1rem] md:tw-text-[1.375rem] tw-leading-[1.25rem] md:tw-leading-[1.375rem] v2-canopas-gradient-text"
            >{{ favourite.likes }}</span
          >
        </div>
        <div class="tw-w-full tw-h-full">
          <img
            :src="favourite.deskImage[0]"
            :srcset="`${favourite.deskImage[0]} 400w, ${favourite.deskImage[1]} 800w`"
            @click="openBlog(favourite)"
            alt="favourite"
            class="tw-w-[90%] tw-h-[57%] tw-object-cover tw-rounded-[10px] tw-cursor-pointer"
            loading="lazy"
          />
        </div>
      </div>
      <div class="tw-flex tw-flex-col tw-gap-8">
        <div
          class="tw-flex tw-flex-col"
          v-for="favourite in favourites.slice(1, 3)"
          :key="favourite.id"
        >
          <div
            class="tw-flex tw-flex-row tw-gap-2 tw-items-center tw-justify-end tw-mb-2 tw-mr-6 lg:tw-mr-[2rem] xl:tw-mr-[2.5rem] 2xl:tw-mr-[3rem]"
          >
            <img
              :src="like_100w"
              alt="like"
              class="tw-w-7 tw-h-7 tw-object-cover"
              loading="lazy"
            /><span
              class="tw-font-inter-bold tw-text-[1rem] md:tw-text-[1.375rem] tw-leading-[1.25rem] md:tw-leading-[1.375rem] v2-canopas-gradient-text"
              >{{ favourite.likes }}</span
            >
          </div>
          <div class="tw-w-full tw-h-full">
            <img
              :src="favourite.deskImage[0]"
              :srcset="`${favourite.deskImage[0]} 400w, ${favourite.deskImage[1]} 800w`"
              @click="openBlog(favourite)"
              alt="favourite"
              class="tw-w-[90%] tw-h-[100%] tw-rounded-[10px] tw-object-cover tw-cursor-pointer"
              loading="lazy"
            />
          </div>
        </div>
      </div>
      <div class="tw-flex tw-flex-col tw-mt-[-5rem] lg:tw-mt-[-8rem]">
        <div
          class="tw-flex tw-flex-col"
          v-for="favourite in favourites.slice(3, 5)"
          :key="favourite.id"
          :class="favourite.id == 5 ? 'tw-mt-[2rem]' : ''"
        >
          <div
            class="tw-flex tw-flex-row tw-gap-2 tw-items-center tw-justify-end tw-mb-2 tw-mr-6 lg:tw-mr-[2rem] xl:tw-mr-[2.5rem] 2xl:tw-mr-[3rem]"
          >
            <img
              :src="like_100w"
              alt="like"
              class="tw-w-7 tw-h-7 tw-object-cover"
              loading="lazy"
            /><span
              class="tw-font-inter-bold tw-text-[1rem] md:tw-text-[1.375rem] tw-leading-[1.25rem] md:tw-leading-[1.375rem] v2-canopas-gradient-text"
              >{{ favourite.likes }}</span
            >
          </div>
          <div class="tw-w-full tw-h-full">
            <img
              :src="favourite.deskImage[0]"
              :srcset="`${favourite.deskImage[0]} 400w, ${favourite.deskImage[1]} 800w`"
              @click="openBlog(favourite)"
              alt="favourite"
              class="tw-w-[90%] tw-h-[100%] tw-object-cover tw-rounded-[10px] tw-cursor-pointer"
              loading="lazy"
            />
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import { Swiper, SwiperSlide } from "swiper/vue";
import golang_400w from "@/assets/images/favourites/mobile/golang-400w.webp";
import golang_800w from "@/assets/images/favourites/mobile/golang-800w.webp";
import android_400w from "@/assets/images/favourites/mobile/android-400w.webp";
import android_800w from "@/assets/images/favourites/mobile/android-800w.webp";
import jetpack_400w from "@/assets/images/favourites/mobile/jetpack-400w.webp";
import jetpack_800w from "@/assets/images/favourites/mobile/jetpack-800w.webp";
import kotlin_400w from "@/assets/images/favourites/mobile/kotlin-400w.webp";
import kotlin_800w from "@/assets/images/favourites/mobile/kotlin-800w.webp";
import uipilot_400w from "@/assets/images/favourites/mobile/uipilot-400w.webp";
import uipilot_800w from "@/assets/images/favourites/mobile/uipilot-800w.webp";

import desktop_golang_400w from "@/assets/images/favourites/desktop/golang-400w.webp";
import desktop_golang_800w from "@/assets/images/favourites/desktop/golang-800w.webp";
import desktop_android_400w from "@/assets/images/favourites/desktop/android-400w.webp";
import desktop_android_800w from "@/assets/images/favourites/desktop/android-800w.webp";
import desktop_jetpack_400w from "@/assets/images/favourites/desktop/jetpack-400w.webp";
import desktop_jetpack_800w from "@/assets/images/favourites/desktop/jetpack-800w.webp";
import desktop_uipilot_400w from "@/assets/images/favourites/desktop/uipilot-400w.webp";
import desktop_uipilot_800w from "@/assets/images/favourites/desktop/uipilot-800w.webp";
import desktop_kotlin_400w from "@/assets/images/favourites/desktop/kotlin-400w.webp";
import desktop_kotlin_800w from "@/assets/images/favourites/desktop/kotlin-800w.webp";
import like_100w from "@/assets/images/favourites/like-100w.webp";

export default {
  data() {
    return {
      like_100w,
      favourites: [
        {
          id: 1,
          image: [golang_400w, golang_800w],
          deskImage: [desktop_golang_400w, desktop_golang_800w],
          description: "Golang — 1 Minute guide to Useful Tips and...",
          likes: "496",
          link: "https://blog.canopas.com/1-min-guide-to-golang-development-best-practices-in-2022-b50d846fd6c",
        },
        {
          id: 2,
          image: [android_400w, android_800w],
          deskImage: [desktop_android_400w, desktop_android_800w],
          description: "Android — 1 Minute guide to Useful Tips and ...",
          likes: "248",
          link: "https://blog.canopas.com/android-development-best-practices-2022-203682a440f5",
        },
        {
          id: 3,
          image: [jetpack_400w, jetpack_800w],
          deskImage: [desktop_jetpack_400w, desktop_jetpack_800w],
          description: "Jetpack compose: Cool Button Click Effects",
          likes: "121",
          link: "https://blog.canopas.com/jetpack-compose-cool-button-click-effects-c6bbecec7bcb",
        },
        {
          id: 4,
          image: [kotlin_400w, kotlin_800w],
          deskImage: [desktop_kotlin_400w, desktop_kotlin_800w],
          description: "Android — 9 Useful Kotlin Flow Operators You...",
          likes: "236",
          link: "https://blog.canopas.com/android-9-useful-kotlin-flow-operators-you-need-to-know-b9daef4b630f",
        },
        {
          id: 5,
          image: [uipilot_400w, uipilot_800w],
          deskImage: [desktop_uipilot_400w, desktop_uipilot_800w],
          description: "SwiftUI — Complex navigation made easier... ",
          likes: "98",
          link: "https://blog.canopas.com/swiftui-complex-navigation-made-easier-with-uipilot-5b33279f3476",
        },
      ],
    };
  },
  methods: {
    openBlog(favourite) {
      window.open(favourite.link, "_blank");
    },
  },
  components: {
    Swiper,
    SwiperSlide,
  },
};
</script>
<style lang="postcss">
@import "swiper/css";
</style>
