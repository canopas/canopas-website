<template>
  <section class="tw-my-20 tw-overflow-hidden md:tw-mt-16">
    <!-- Mobile UI start -->
    <p
      class="tw-container tw-text-center tw-font-inter-bold tw-text-[1.875rem] tw-leading-[2.438rem] tw-text-black-core/[0.87] md:tw-ml-[6%] md:tw-w-[39%] md:tw-text-left lg:tw-hidden lg:tw-text-[3.438rem] lg:tw-leading-[5.156rem]"
    >
      The Favorites You Can't Miss
    </p>
    <div class="tw-mt-8 tw-block md:tw-hidden">
      <swiper :slidesPerView="1.2" :spaceBetween="0" class="tw-p-[1rem]">
        <swiper-slide
          v-for="(favourite, index) in favourites"
          :key="index"
          class="tw-relative tw-flex tw-flex-col tw-pl-[1rem] sm:tw-pl-[2rem]"
        >
          <div
            class="tw-mb-2 tw-flex tw-flex-row tw-items-center tw-justify-end tw-gap-2"
          >
            <img
              :src="like_100w"
              alt="like"
              class="tw-h-7 tw-w-7 tw-object-cover"
              loading="lazy"
            /><span
              class="v2-canopas-gradient-text tw-font-inter-bold tw-text-[1rem] tw-leading-[1.25rem]"
              >{{ favourite.likes }}</span
            >
          </div>
          <div class="tw-overflow-hidden">
            <div class="tw-h-full tw-w-full">
              <img
                @click="openBlog(favourite.link)"
                :src="favourite.image[0]"
                :srcset="`${favourite.image[0]} 400w, ${favourite.image[1]} 800w`"
                alt="favourite"
                class="tw-h-full tw-w-full tw-rounded-[10px] tw-object-cover"
                loading="lazy"
              />
            </div>
            <p
              class="tw-mt-2 tw-pl-2 tw-font-inter-medium tw-text-[1rem] tw-leading-[1.5rem] tw-text-black-core/[0.87] sm:tw-text-[1.25rem]"
            >
              {{ favourite.description }}
            </p>
          </div>
        </swiper-slide>
      </swiper>
    </div>
    <!-- Mobile UI end -->
    <!-- Desktop UI start -->
    <p
      class="tw-container tw-hidden tw-font-inter-bold tw-text-[2rem] tw-leading-[3rem] tw-text-black-core/[0.87] lg:tw-block lg:tw-text-[3rem] lg:tw-leading-[4rem]"
    >
      The Favorites You <br />
      Can't Miss
    </p>
    <div class="tw-container tw-hidden tw-flex-row tw-gap-x-8 md:tw-flex">
      <div class="tw-flex tw-flex-col tw-justify-center tw-gap-8">
        <div
          class="tw-flex tw-flex-col"
          v-for="favourite in favourites.slice(0, 1)"
          :key="favourite.id"
        >
          <div
            class="tw-mb-2 tw-mr-6 tw-flex tw-flex-row tw-items-center tw-justify-end tw-gap-2 lg:tw-mr-[2rem] xl:tw-mr-[2.5rem] 2xl:tw-mr-[3rem]"
          >
            <img
              :src="like_100w"
              alt="like"
              class="tw-h-7 tw-w-7 tw-object-cover"
              loading="lazy"
            /><span
              class="v2-canopas-gradient-text tw-font-inter-bold tw-text-[1rem] tw-leading-[1.25rem] md:tw-text-[1.375rem] md:tw-leading-[1.375rem]"
              >{{ favourite.likes }}</span
            >
          </div>
          <div
            class="tw-group tw-relative tw-h-full tw-w-full tw-cursor-pointer tw-overflow-hidden"
            @click="openBlog(favourite.link)"
          >
            <img
              :src="favourite.deskImage[0]"
              :srcset="`${favourite.deskImage[0]} 400w, ${favourite.deskImage[1]} 800w`"
              alt="favourite"
              class="tw-h-full tw-w-full tw-rounded-[10px] tw-object-contain"
              loading="lazy"
            />
            <div
              class="tw-ease tw-absolute tw-left-[0] tw-top-[0] tw-h-[0] tw-w-[0] tw-rounded-[10px] tw-from-[#FF835B] tw-to-[#F2709C] tw-opacity-0 tw-transition-all tw-duration-1000 tw-bg-gradient-[180deg] group-hover:tw-h-full group-hover:tw-w-full group-hover:tw-opacity-[0.9]"
            ></div>
            <div
              class="tw-absolute tw-left-[4%] tw-top-[32%] tw-h-full tw-w-[82%] tw-translate-y-[50%] tw-text-center tw-font-inter-bold tw-text-[1rem] tw-leading-[1.8rem] tw-text-white tw-opacity-0 tw-transition-all tw-duration-1000 group-hover:tw-translate-y-[0%] group-hover:tw-opacity-[1] group-hover:tw-transition-all group-hover:tw-duration-1000 lg:tw-top-[35%] lg:tw-w-[80%] lg:tw-text-[1.5rem] lg:tw-leading-[2.25rem]"
            >
              {{ favourite.description }}
            </div>
          </div>
        </div>
      </div>
      <div class="tw-flex tw-flex-col tw-gap-8">
        <div
          class="tw-flex tw-flex-col"
          v-for="favourite in favourites.slice(1, 3)"
          :key="favourite.id"
        >
          <div
            class="tw-mb-2 tw-mr-6 tw-flex tw-flex-row tw-items-center tw-justify-end tw-gap-2 lg:tw-mr-[2rem] xl:tw-mr-[2.5rem] 2xl:tw-mr-[3rem]"
          >
            <img
              :src="like_100w"
              alt="like"
              class="tw-h-7 tw-w-7 tw-object-cover"
              loading="lazy"
            /><span
              class="v2-canopas-gradient-text tw-font-inter-bold tw-text-[1rem] tw-leading-[1.25rem] md:tw-text-[1.375rem] md:tw-leading-[1.375rem]"
              >{{ favourite.likes }}</span
            >
          </div>
          <div
            class="tw-group tw-relative tw-h-full tw-w-full tw-cursor-pointer tw-overflow-hidden"
            @click="openBlog(favourite.link)"
          >
            <img
              :src="favourite.deskImage[0]"
              :srcset="`${favourite.deskImage[0]} 400w, ${favourite.deskImage[1]} 800w`"
              alt="favourite"
              class="tw-h-full tw-w-full tw-rounded-[10px] tw-object-contain"
              loading="lazy"
            />
            <div
              class="tw-ease tw-absolute tw-left-[0] tw-top-[0] tw-h-[0] tw-w-[0] tw-rounded-[10px] tw-from-[#FF835B] tw-to-[#F2709C] tw-opacity-0 tw-transition-all tw-duration-1000 tw-bg-gradient-[180deg] group-hover:tw-h-full group-hover:tw-w-full group-hover:tw-opacity-[0.9]"
            ></div>
            <div
              class="tw-absolute tw-left-[4%] tw-top-[32%] tw-h-full tw-w-[82%] tw-translate-y-[50%] tw-text-center tw-font-inter-bold tw-text-[1rem] tw-leading-[1.8rem] tw-text-white tw-opacity-0 tw-transition-all tw-duration-1000 group-hover:tw-translate-y-[0%] group-hover:tw-opacity-[1] group-hover:tw-transition-all group-hover:tw-duration-1000 lg:tw-top-[35%] lg:tw-w-[80%] lg:tw-text-[1.5rem] lg:tw-leading-[2.25rem]"
            >
              {{ favourite.description }}
            </div>
          </div>
        </div>
      </div>
      <div class="tw-mt-[-5rem] tw-flex tw-flex-col lg:tw-mt-[-8rem]">
        <div
          class="tw-flex tw-flex-col"
          v-for="favourite in favourites.slice(3, 5)"
          :key="favourite.id"
          :class="favourite.id == 5 ? 'tw-mt-[2rem]' : ''"
        >
          <div
            class="tw-mb-2 tw-mr-6 tw-flex tw-flex-row tw-items-center tw-justify-end tw-gap-2 lg:tw-mr-[2rem] xl:tw-mr-[2.5rem] 2xl:tw-mr-[3rem]"
          >
            <img
              :src="like_100w"
              alt="like"
              class="tw-h-7 tw-w-7 tw-object-cover"
              loading="lazy"
            /><span
              class="v2-canopas-gradient-text tw-font-inter-bold tw-text-[1rem] tw-leading-[1.25rem] md:tw-text-[1.375rem] md:tw-leading-[1.375rem]"
              >{{ favourite.likes }}</span
            >
          </div>
          <div
            class="tw-group tw-relative tw-h-full tw-w-full tw-cursor-pointer tw-overflow-hidden"
            @click="openBlog(favourite.link)"
          >
            <img
              :src="favourite.deskImage[0]"
              :srcset="`${favourite.deskImage[0]} 400w, ${favourite.deskImage[1]} 800w`"
              alt="favourite"
              class="tw-h-full tw-w-full tw-rounded-[10px] tw-object-contain"
              loading="lazy"
            />
            <div
              class="tw-ease tw-absolute tw-left-[0] tw-top-[0] tw-h-[0] tw-w-[0] tw-rounded-[10px] tw-from-[#FF835B] tw-to-[#F2709C] tw-opacity-0 tw-transition-all tw-duration-1000 tw-bg-gradient-[180deg] group-hover:tw-h-full group-hover:tw-w-full group-hover:tw-opacity-[0.9]"
            ></div>
            <div
              class="tw-absolute tw-left-[4%] tw-top-[32%] tw-h-full tw-w-[82%] tw-translate-y-[50%] tw-text-center tw-font-inter-bold tw-text-[1rem] tw-leading-[1.8rem] tw-text-white tw-opacity-0 tw-transition-all tw-duration-1000 group-hover:tw-translate-y-[0%] group-hover:tw-opacity-[1] group-hover:tw-transition-all group-hover:tw-duration-1000 lg:tw-top-[35%] lg:tw-w-[80%] lg:tw-text-[1.5rem] lg:tw-leading-[2.25rem]"
            >
              {{ favourite.description }}
            </div>
          </div>
        </div>
      </div>
    </div>
    <!-- Desktop UI end -->
  </section>
</template>

<script>
import { Swiper, SwiperSlide } from "swiper/vue";

import golang_400w from "@/assets/images/contributions/favourites/mobile/golang-400w.webp";
import golang_800w from "@/assets/images/contributions/favourites/mobile/golang-800w.webp";
import android_400w from "@/assets/images/contributions/favourites/mobile/android-400w.webp";
import android_800w from "@/assets/images/contributions/favourites/mobile/android-800w.webp";
import jetpack_400w from "@/assets/images/contributions/favourites/mobile/jetpack-400w.webp";
import jetpack_800w from "@/assets/images/contributions/favourites/mobile/jetpack-800w.webp";
import kotlin_400w from "@/assets/images/contributions/favourites/mobile/kotlin-400w.webp";
import kotlin_800w from "@/assets/images/contributions/favourites/mobile/kotlin-800w.webp";
import dragDrop_400w from "@/assets/images/contributions/favourites/mobile/dragDrop-400w.webp";
import dragDrop_800w from "@/assets/images/contributions/favourites/mobile/dragDrop-800w.webp";

import desktop_golang_400w from "@/assets/images/contributions/favourites/desktop/golang-400w.webp";
import desktop_golang_800w from "@/assets/images/contributions/favourites/desktop/golang-800w.webp";
import desktop_android_400w from "@/assets/images/contributions/favourites/desktop/android-400w.webp";
import desktop_android_800w from "@/assets/images/contributions/favourites/desktop/android-800w.webp";
import desktop_jetpack_400w from "@/assets/images/contributions/favourites/desktop/jetpack-400w.webp";
import desktop_jetpack_800w from "@/assets/images/contributions/favourites/desktop/jetpack-800w.webp";
import desktop_dragDrop_400w from "@/assets/images/contributions/favourites/desktop/dragDrop-400w.webp";
import desktop_dragDrop_800w from "@/assets/images/contributions/favourites/desktop/dragDrop-800w.webp";
import desktop_kotlin_400w from "@/assets/images/contributions/favourites/desktop/kotlin-400w.webp";
import desktop_kotlin_800w from "@/assets/images/contributions/favourites/desktop/kotlin-800w.webp";
import like_100w from "@/assets/images/contributions/favourites/like-100w.webp";

export default {
  data() {
    return {
      like_100w,
      favourites: [
        {
          id: 1,
          image: [android_400w, android_800w],
          deskImage: [desktop_android_400w, desktop_android_800w],
          description:
            "Android — 1 Minute guide to Useful Tips and Libraries in 2022",
          likes: "500",
          link: "https://blog.canopas.com/android-development-best-practices-2022-203682a440f5",
        },
        {
          id: 2,
          image: [kotlin_400w, kotlin_800w],
          deskImage: [desktop_kotlin_400w, desktop_kotlin_800w],
          description:
            "Android — 9 Useful Kotlin Flow Operators You Need to know",
          likes: "249",
          link: "https://blog.canopas.com/android-9-useful-kotlin-flow-operators-you-need-to-know-b9daef4b630f",
        },
        {
          id: 3,
          image: [golang_400w, golang_800w],
          deskImage: [desktop_golang_400w, desktop_golang_800w],
          description:
            "Golang — 1 Minute guide to Useful Tips and Libraries in 2022",
          likes: "237",
          link: "https://blog.canopas.com/1-min-guide-to-golang-development-best-practices-in-2022-b50d846fd6c",
        },
        {
          id: 4,
          image: [jetpack_400w, jetpack_800w],
          deskImage: [desktop_jetpack_400w, desktop_jetpack_800w],
          description: "Jetpack compose: Cool Button Click Effects",
          likes: "129",
          link: "https://blog.canopas.com/jetpack-compose-cool-button-click-effects-c6bbecec7bcb",
        },
        {
          id: 5,
          image: [dragDrop_400w, dragDrop_800w],
          deskImage: [desktop_dragDrop_400w, desktop_dragDrop_800w],
          description:
            "Android — How to Drag And Drop Views in Jetpack Compose",
          likes: "101",
          link: "https://blog.canopas.com/android-drag-and-drop-ui-element-in-jetpack-compose-14922073b3f1",
        },
      ],
    };
  },
  methods: {
    openBlog(link) {
      window.open(link, "_blank");
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
