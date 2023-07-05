<template>
  <section class="tw-my-[60px] tw-overflow-hidden">
    <div
      class="tw-container tw-text-center tw-font-inter-bold tw-text-[1.875rem] tw-leading-[2.438rem] tw-text-black-core/[0.87] lg:tw-text-[3.438rem] lg:tw-leading-[5.156rem]"
    >
      Explore Our Animated Creations
    </div>
    <!-- Mobile UI start -->
    <div class="tw-mt-10 tw-block md:tw-hidden">
      <swiper :slidesPerView="1.1" :spaceBetween="0" class="tw-p-[1rem]">
        <swiper-slide
          v-for="(animation, index) in animations"
          :key="index"
          class="tw-pl-[1rem] tw-pr-[1rem] sm:tw-pl-[2rem]"
        >
          <div
            class="tw-flex tw-items-center tw-justify-center tw-rounded-[20px] tw-border-1 tw-bg-white tw-drop-shadow-md"
          >
            <div class="tw-flex-[39%] sm:tw-flex-[45%]">
              <img
                @click="openBlog(animation.link)"
                :src="animation.image[0]"
                :srcset="`${animation.image[0]} 400w, ${animation.image[1]} 800w`"
                alt="animation"
                class="tw-h-full tw-w-full tw-object-cover"
                loading="lazy"
              />
            </div>
            <div
              class="tw-flex-[50%] tw-px-[0.5rem] tw-font-inter-medium tw-text-[1rem] tw-leading-[1.5rem] tw-text-black-core/[0.87] sm:tw-text-[1.25rem] sm:tw-leading-[1.688rem]"
            >
              {{ animation.title }}
            </div>
          </div>
        </swiper-slide>
      </swiper>
    </div>
    <!-- Mobile UI end -->
    <!-- Desktop UI start -->
    <div
      class="tw-container tw-mt-[50px] tw-hidden tw-flex-row tw-gap-8 md:tw-flex"
    >
      <div class="tw-flex tw-w-[50%] tw-flex-col tw-gap-8">
        <div v-for="animation in animations.slice(0, 3)" :key="animation.id">
          <div
            class="tw-group tw-relative tw-cursor-pointer tw-overflow-hidden"
            @click="openBlog(animation.link)"
          >
            <img
              :src="animation.deskImages[0]"
              :srcset="`${animation.deskImages[0]} 400w, ${animation.deskImages[1]} 800w`"
              alt="animation"
              class="tw-h-full tw-w-full tw-rounded-bl-[50px] tw-rounded-tr-[50px] tw-object-cover"
              loading="lazy"
            />
            <div
              class="tw-ease tw-absolute tw-left-[50%] tw-top-[50%] tw-h-[0%] tw-w-[0%] tw--translate-x-1/2 tw--translate-y-1/2 tw-rounded-bl-[50px] tw-rounded-tr-[50px] tw-bg-black-900 tw-opacity-0 tw-transition-all tw-duration-1000 group-hover:tw-h-[100%] group-hover:tw-w-[100%] group-hover:tw-opacity-[1]"
            ></div>
            <div
              class="tw-absolute tw-inset-x-0 tw-top-[40%] tw-m-auto tw-h-full tw-translate-y-[-50%] tw-px-4 tw-text-center tw-font-inter-bold tw-text-[1rem] tw-leading-[1.5rem] tw-text-white tw-opacity-0 tw-transition-all tw-duration-1000 group-hover:tw-translate-y-[0%] group-hover:tw-opacity-[1] group-hover:tw-transition-all group-hover:tw-duration-1000 lg:tw-text-[1.75rem] lg:tw-leading-[2.625rem]"
            >
              {{ animation.title }}
            </div>
          </div>
        </div>
      </div>
      <div class="tw-mt-[50px] tw-flex tw-w-[50%] tw-flex-col tw-gap-8">
        <div v-for="animation in animations.slice(3, 6)" :key="animation.id">
          <div
            class="tw-group tw-relative tw-cursor-pointer tw-overflow-hidden"
            @click="openBlog(animation.link)"
          >
            <img
              :src="animation.deskImages[0]"
              :srcset="`${animation.deskImages[0]} 400w, ${animation.deskImages[1]} 800w`"
              alt="animation"
              class="tw-h-full tw-w-full tw-rounded-bl-[50px] tw-rounded-tr-[50px] tw-object-cover"
              loading="lazy"
            />
            <div
              class="tw-ease tw-absolute tw-left-[50%] tw-top-[50%] tw-h-0 tw-w-0 tw--translate-x-1/2 tw--translate-y-1/2 tw-rounded-bl-[50px] tw-rounded-tr-[50px] tw-bg-black-900 tw-opacity-0 tw-transition-all tw-duration-1000 group-hover:tw-h-[100%] group-hover:tw-w-[100%] group-hover:tw-opacity-[1]"
            ></div>
            <div
              class="tw-absolute tw-inset-x-0 tw-top-[40%] tw-m-auto tw-h-full tw-translate-y-[-50%] tw-px-4 tw-text-center tw-font-inter-bold tw-text-[1rem] tw-leading-[1.5rem] tw-text-white tw-opacity-0 tw-transition-all tw-duration-1000 group-hover:tw-translate-y-[0%] group-hover:tw-opacity-[1] group-hover:tw-transition-all group-hover:tw-duration-1000 lg:tw-text-[1.75rem] lg:tw-leading-[2.625rem]"
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
  methods: {
    openBlog(link) {
      window.open(link, "_blank");
      this.mixpanel.track("tap_animation_creation_section");
    },
  },
};
</script>
