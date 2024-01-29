<template>
  <section class="my-16 overflow-hidden lg:mt-60">
    <!-- Mobile UI start -->
    <p class="container text-center mobile-header-2 lg:hidden">
      The favorites you can't miss
    </p>
    <div class="mt-8 block lg:hidden">
      <swiper :slidesPerView="1.2" :spaceBetween="0" class="p-4">
        <swiper-slide
          v-for="(favourite, index) in favourites"
          :key="index"
          class="relative flex flex-col pl-4 sm:pl-8"
        >
          <div class="mb-2 flex flex-row items-center justify-end gap-2">
            <img
              :src="like_100w"
              alt="like"
              class="h-7 w-7 object-cover"
              loading="lazy"
            /><span class="text-primary-1 sub-h3-semibold">{{
              favourite.likes
            }}</span>
          </div>
          <div class="overflow-hidden">
            <div class="h-full w-full">
              <img
                @click="
                  openBlog(favourite.link, 'tap_contribution_favourite_section')
                "
                :src="favourite.image[0]"
                :srcset="`${favourite.image[0]} 400w, ${favourite.image[1]} 800w`"
                alt="favourite"
                class="h-full w-full rounded-xl object-cover"
                loading="lazy"
              />
            </div>
            <p class="mt-3 sub-h1-semibold text-black-87">
              {{ favourite.description }}
            </p>
          </div>
        </swiper-slide>
      </swiper>
    </div>
    <!-- Mobile UI end -->
    <!-- Desktop UI start -->
    <p class="container hidden lg:block desk-header-2 text-black-87">
      The Favorites You <br />
      Can't Miss
    </p>
    <div class="container hidden flex-row gap-x-8 lg:flex">
      <div class="flex flex-col justify-center gap-8">
        <div
          class="flex flex-col"
          v-for="favourite in favourites.slice(0, 1)"
          :key="favourite.id"
        >
          <div class="mb-2 flex flex-row items-center justify-end gap-2">
            <img
              :src="like_100w"
              alt="like"
              class="h-7 w-7 object-cover"
              loading="lazy"
            /><span class="text-primary-1 mobile-header-3-semibold">{{
              favourite.likes
            }}</span>
          </div>
          <div
            class="group relative h-full w-full cursor-pointer overflow-hidden"
            @click="
              openBlog(favourite.link, 'tap_contribution_favourite_section')
            "
          >
            <img
              :src="favourite.deskImage[0]"
              :srcset="`${favourite.deskImage[0]} 400w, ${favourite.deskImage[1]} 800w`"
              alt="favourite"
              class="h-full w-full rounded-xl object-contain"
              loading="lazy"
            />
            <div
              class="ease absolute left-0 top-0 h-0 w-0 rounded-[10px] from-pink-300 to-orange-300 opacity-0 transition-all duration-1000 bg-gradient-[180deg] group-hover:h-full group-hover:w-full group-hover:opacity-90"
            ></div>
            <div
              class="mobile-header-2-semibold absolute left-[10%] xl:left-[8%] top-[25%] xl:top-[32%] 2xl:top-[35%] 2xl:left-[9%] h-full w-[82%] translate-y-1/2 text-center text-white opacity-0 transition-all duration-1000 group-hover:translate-y-0 group-hover:opacity-100 group-hover:transition-all group-hover:duration-1000 lg:w-4/5"
            >
              {{ favourite.description }}
            </div>
          </div>
        </div>
      </div>
      <div class="flex flex-col gap-8">
        <div
          class="flex flex-col"
          v-for="favourite in favourites.slice(1, 3)"
          :key="favourite.id"
        >
          <div class="mb-2 flex flex-row items-center justify-end gap-2">
            <img
              :src="like_100w"
              alt="like"
              class="h-7 w-7 object-cover"
              loading="lazy"
            /><span class="text-primary-1 mobile-header-3-semibold">{{
              favourite.likes
            }}</span>
          </div>
          <div
            class="group relative h-full w-full cursor-pointer overflow-hidden"
            @click="
              openBlog(favourite.link, 'tap_contribution_favourite_section')
            "
          >
            <img
              :src="favourite.deskImage[0]"
              :srcset="`${favourite.deskImage[0]} 400w, ${favourite.deskImage[1]} 800w`"
              alt="favourite"
              class="h-full w-full rounded-xl object-contain"
              loading="lazy"
            />
            <div
              class="ease absolute left-0 top-0 h-0 w-0 rounded-[10px] from-pink-300 to-orange-300 opacity-0 transition-all duration-1000 bg-gradient-[180deg] group-hover:h-full group-hover:w-full group-hover:opacity-90"
            ></div>
            <div
              class="mobile-header-2-semibold absolute left-[10%] xl:left-[8%] top-[25%] xl:top-[32%] 2xl:top-[37%] 2xl:left-[9%] h-full w-[82%] translate-y-1/2 text-center text-white opacity-0 transition-all duration-1000 group-hover:translate-y-0 group-hover:opacity-100 group-hover:transition-all group-hover:duration-1000 lg:w-4/5"
            >
              {{ favourite.description }}
            </div>
          </div>
        </div>
      </div>
      <div class="-mt-20 flex flex-col lg:-mt-32">
        <div
          class="flex flex-col"
          v-for="favourite in favourites.slice(3, 5)"
          :key="favourite.id"
          :class="favourite.id == 5 ? 'mt-8' : ''"
        >
          <div class="mb-2 flex flex-row items-center justify-end gap-2">
            <img
              :src="like_100w"
              alt="like"
              class="h-7 w-7 object-cover"
              loading="lazy"
            /><span class="text-primary-1 mobile-header-3-semibold">{{
              favourite.likes
            }}</span>
          </div>
          <div
            class="group relative h-full w-full cursor-pointer overflow-hidden"
            @click="
              openBlog(favourite.link, 'tap_contribution_favourite_section')
            "
          >
            <img
              :src="favourite.deskImage[0]"
              :srcset="`${favourite.deskImage[0]} 400w, ${favourite.deskImage[1]} 800w`"
              alt="favourite"
              class="h-full w-full rounded-xl object-contain"
              loading="lazy"
            />
            <div
              class="ease absolute left-0 top-0 h-0 w-0 rounded-[10px] from-pink-300 to-orange-300 opacity-0 transition-all duration-1000 bg-gradient-[180deg] group-hover:h-full group-hover:w-full group-hover:opacity-90"
            ></div>
            <div
              class="mobile-header-2-semibold absolute left-[10%] xl:left-[8%] top-[25%] xl:top-[38%] 2xl:left-[9%] h-full w-[82%] translate-y-1/2 text-center text-white opacity-0 transition-all duration-1000 group-hover:translate-y-0 group-hover:opacity-100 group-hover:transition-all group-hover:duration-1000 lg:w-4/5"
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
import { openBlog } from "@/utils.js";

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
      openBlog,
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

  components: {
    Swiper,
    SwiperSlide,
  },
};
</script>
<style lang="postcss">
@import "swiper/css";
</style>
