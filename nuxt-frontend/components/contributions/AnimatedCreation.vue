<template>
  <section class="mb-[6.063rem] lg:my-[14.063rem] overflow-hidden">
    <div
      class="container text-center font-inter-bold text-[1.875rem] leading-[2.438rem] text-black-core/[0.87] lg:text-[3.438rem] lg:leading-[5.156rem]"
    >
      Explore Our Animated Creations
    </div>
    <!-- Mobile UI start -->
    <div class="mt-10 block md:hidden">
      <swiper :slidesPerView="1.1" :spaceBetween="0" class="p-[1rem]">
        <swiper-slide
          v-for="(animation, index) in animations"
          :key="index"
          class="pl-4 pr-4 sm:pl-8"
        >
          <div
            class="flex items-center justify-center rounded-[20px] border-1 bg-white drop-shadow-md"
          >
            <div class="flex-[39%] sm:flex-[45%]">
              <img
                @click="
                  openBlog(animation.link, 'tap_animation_creation_section')
                "
                :src="animation.image[0]"
                :srcset="`${animation.image[0]} 400w, ${animation.image[1]} 800w`"
                alt="animation"
                class="h-full w-full object-cover"
                loading="lazy"
              />
            </div>
            <div
              class="flex-[50%] px-2 font-inter-medium text-base text-black-core/[0.87] sm:text-[1.25rem] sm:leading-[1.688rem]"
            >
              {{ animation.title }}
            </div>
          </div>
        </swiper-slide>
      </swiper>
    </div>
    <!-- Mobile UI end -->
    <!-- Desktop UI start -->
    <div class="container mt-[50px] hidden flex-row gap-8 md:flex">
      <div class="flex w-[50%] flex-col gap-8">
        <div v-for="animation in animations.slice(0, 3)" :key="animation.id">
          <div
            class="group relative cursor-pointer overflow-hidden"
            @click="openBlog(animation.link, 'tap_animation_creation_section')"
          >
            <img
              :src="animation.deskImages[0]"
              :srcset="`${animation.deskImages[0]} 400w, ${animation.deskImages[1]} 800w`"
              alt="animation"
              class="h-full w-full rounded-bl-[50px] rounded-tr-[50px] object-cover"
              loading="lazy"
            />
            <div
              class="ease absolute left-[50%] top-[50%] h-[0%] w-[0%] -translate-x-1/2 -translate-y-1/2 rounded-bl-[50px] rounded-tr-[50px] bg-black-900 opacity-0 transition-all duration-1000 group-hover:h-full group-hover:w-full group-hover:opacity-[1]"
            ></div>
            <div
              class="absolute inset-x-0 top-[40%] m-auto h-full translate-y-[-50%] px-4 text-center font-inter-bold text-base text-white opacity-0 transition-all duration-1000 group-hover:translate-y-[0%] group-hover:opacity-[1] group-hover:transition-all group-hover:duration-1000 lg:text-[1.75rem] lg:leading-[2.625rem]"
            >
              {{ animation.title }}
            </div>
          </div>
        </div>
      </div>
      <div class="mt-[50px] flex w-[50%] flex-col gap-8">
        <div v-for="animation in animations.slice(3, 6)" :key="animation.id">
          <div
            class="group relative cursor-pointer overflow-hidden"
            @click="openBlog(animation.link, 'tap_animation_creation_section')"
          >
            <img
              :src="animation.deskImages[0]"
              :srcset="`${animation.deskImages[0]} 400w, ${animation.deskImages[1]} 800w`"
              alt="animation"
              class="h-full w-full rounded-bl-[50px] rounded-tr-[50px] object-cover"
              loading="lazy"
            />
            <div
              class="ease absolute left-[50%] top-[50%] h-0 w-0 -translate-x-1/2 -translate-y-1/2 rounded-bl-[50px] rounded-tr-[50px] bg-black-900 opacity-0 transition-all duration-1000 group-hover:h-full group-hover:w-full group-hover:opacity-[1]"
            ></div>
            <div
              class="absolute inset-x-0 top-[40%] m-auto h-full translate-y-[-50%] px-4 text-center font-inter-bold text-base text-white opacity-0 transition-all duration-1000 group-hover:translate-y-[0%] group-hover:opacity-[1] group-hover:transition-all group-hover:duration-1000 lg:text-[1.75rem] lg:leading-[2.625rem]"
            >
              {{ animation.title }}
            </div>
          </div>
        </div>
      </div>
    </div>
    <!-- Desktop UI end -->
  </section>
</template>

<script>
import { openBlog } from "@/utils.js";

import { Swiper, SwiperSlide } from "swiper/vue";
import animation1_400w from "@/assets/images/contributions/animations/mobile/animation-1-400w.webp";
import animation1_800w from "@/assets/images/contributions/animations/mobile/animation-1-800w.webp";
import animation2_400w from "@/assets/images/contributions/animations/mobile/animation-2-400w.webp";
import animation2_800w from "@/assets/images/contributions/animations/mobile/animation-2-800w.webp";
import animation3_400w from "@/assets/images/contributions/animations/mobile/animation-3-400w.webp";
import animation3_800w from "@/assets/images/contributions/animations/mobile/animation-3-800w.webp";
import animation4_400w from "@/assets/images/contributions/animations/mobile/animation-4-400w.webp";
import animation4_800w from "@/assets/images/contributions/animations/mobile/animation-4-800w.webp";
import animation5_400w from "@/assets/images/contributions/animations/mobile/animation-5-400w.webp";
import animation5_800w from "@/assets/images/contributions/animations/mobile/animation-5-800w.webp";
import animation6_400w from "@/assets/images/contributions/animations/mobile/animation-6-400w.webp";
import animation6_800w from "@/assets/images/contributions/animations/mobile/animation-6-800w.webp";

import desktop_animation1_400w from "@/assets/images/contributions/animations/desktop/animation-1-400w.webp";
import desktop_animation1_800w from "@/assets/images/contributions/animations/desktop/animation-1-800w.webp";
import desktop_animation2_400w from "@/assets/images/contributions/animations/desktop/animation-2-400w.webp";
import desktop_animation2_800w from "@/assets/images/contributions/animations/desktop/animation-2-800w.webp";
import desktop_animation3_400w from "@/assets/images/contributions/animations/desktop/animation-3-400w.webp";
import desktop_animation3_800w from "@/assets/images/contributions/animations/desktop/animation-3-800w.webp";
import desktop_animation4_400w from "@/assets/images/contributions/animations/desktop/animation-4-400w.webp";
import desktop_animation4_800w from "@/assets/images/contributions/animations/desktop/animation-4-800w.webp";
import desktop_animation5_400w from "@/assets/images/contributions/animations/desktop/animation-5-400w.webp";
import desktop_animation5_800w from "@/assets/images/contributions/animations/desktop/animation-5-800w.webp";
import desktop_animation6_400w from "@/assets/images/contributions/animations/desktop/animation-6-400w.webp";
import desktop_animation6_800w from "@/assets/images/contributions/animations/desktop/animation-6-800w.webp";

export default {
  data() {
    return {
      openBlog,
      animations: [
        {
          id: 1,
          image: [animation1_400w, animation1_800w],
          deskImages: [desktop_animation1_400w, desktop_animation1_800w],
          title: "Part 3 — Animations in SwiftUI with Examples",
          link: "https://blog.canopas.com/animations-in-swiftui-with-examples-part-3-56411ae2ffe0",
        },
        {
          id: 2,
          image: [animation3_400w, animation3_800w],
          deskImages: [desktop_animation3_400w, desktop_animation3_800w],
          title:
            "CSS Made Simple: A Guide To Animations, Transitions, And Transformations",
          link: "https://blog.canopas.com/css-made-simple-a-guide-to-animations-transitions-and-transformations-f9b2ecc8620f",
        },
        {
          id: 3,
          image: [animation5_400w, animation5_800w],
          deskImages: [desktop_animation5_400w, desktop_animation5_800w],
          title:
            "Flutter Animations Made Fun: A Smiley That Reacts to Your Every Move",
          link: "https://blog.canopas.com/flutter-animations-made-fun-a-smiley-that-reacts-to-your-every-move-9ebfb0a0657f",
        },
        {
          id: 4,
          image: [animation2_400w, animation2_800w],
          deskImages: [desktop_animation2_400w, desktop_animation2_800w],
          title: "iOS — Animations in SwiftUI with Examples — Part 2",
          link: "https://blog.canopas.com/ios-animations-in-swiftui-with-examples-part-2-980a2df69145",
        },

        {
          id: 5,
          image: [animation4_400w, animation4_800w],
          deskImages: [desktop_animation4_400w, desktop_animation4_800w],
          title: "Jetpack Compose Typewriter animation with highlighted texts",
          link: "https://blog.canopas.com/jetpack-compose-typewriter-animation-with-highlighted-texts-74397fee42f1",
        },
        {
          id: 6,
          image: [animation6_400w, animation6_800w],
          deskImages: [desktop_animation6_400w, desktop_animation6_800w],
          title: "Animations in Jetpack Compose with examples",
          link: "https://medium.com/canopas/animations-in-jetpack-compose-with-examples-48307ba9dff1",
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
