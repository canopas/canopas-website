<template>
  <section class="mt-16 lg:mt-60">
    <div class="container flex flex-col items-center gap-y-2.5">
      <div class="text-center">
        <span class="mobile-header-2 lg:desk-header-2 text-black-87">
          Our weekly tech updates</span
        >
      </div>
      <div
        class="mb-8 mt-2.5 text-center lg:mt-6 sub-h3-regular lg:mobile-header-2-regular text-black-60 xl:w-4/5 2xl:w-[70%]"
      >
        Each week, we curate a hand-picked selection of the latest tech updates,
        delivering them straight to you. Immerse yourself in our weekly insights
        and equip yourself with the knowledge to navigate tomorrow.
      </div>
    </div>
    <!-- Mobile UI start -->
    <div class="swiper-content block lg:hidden">
      <swiper
        :slidesPerView="1.1"
        :effect="'cards'"
        :grabCursor="true"
        :modules="modules"
        :breakpoints="{
          '768': {
            slidesPerView: 1,
          },
        }"
        class="swiper-container"
      >
        <swiper-slide
          v-for="(weekly, index) in weeklies"
          :key="index"
          class="rounded-t-lg rounded-b-lg"
        >
          <div
            class="h-full w-full object-cover"
            @click="openBlog(weekly.url, weekly.event)"
          >
            <img
              :src="[weekly.image]"
              alt="Weekly-stack-image"
              class="h-fit w-fit object-contain"
              loading="lazy"
            />
            <div class="bg-white-smoke-1 p-4 pb-14">
              <div class="mt-4 flex flex-col">
                <div class="sub-h1-semibold text-black-87">
                  {{ weekly.title }}
                </div>
                <div class="mt-2.5 sub-h3-medium text-black-60">
                  {{ weekly.content }}
                </div>
              </div>
              <div class="mt-6 flex flex-row justify-between">
                <span class="sub-h4-medium text-black-87">{{
                  weekly.author
                }}</span>
                <span class="sub-h4-regular text-black-60">{{
                  weekly.readtime
                }}</span>
              </div>
            </div>
          </div>
        </swiper-slide>
      </swiper>
    </div>
    <!-- Mobile UI end -->
    <!-- Desktop UI start -->
    <ul id="scrollContainer" class="w-full mt-20 lg:mt-16 hidden lg:block">
      <li
        v-for="(weekly, index) in weeklies"
        :key="index"
        :ref="'weekly-' + index"
        :id="'weekly-' + index"
        @click="openBlog(weekly.url, weekly.event)"
        class="origin-[center top] sticky top-[8.625rem] xll:top-[15.625rem] 3xl:top-[20.625rem] mx-[1%] h-[24.375rem] cursor-pointer overflow-hidden drop-shadow-xl md:container xl:mx-auto xl:h-[30rem] 2xl:h-[36.875rem] xll:h-[34.125rem]"
        :style="{
          transform: `translateY(${weekly.translate}px) scale(${weekly.scale})`,
        }"
      >
        <div
          :class="weekly.color"
          class="mx-auto flex h-[21.875rem] flex-row xl:h-[28.125rem] 2xl:h-[32.188rem]"
        >
          <div class="basis-1/2 p-5">
            <img
              :src="[weekly.image]"
              alt="Weekly-stack-image"
              class="h-[72%] lg:h-full w-fit object-cover"
              loading="lazy"
            />
          </div>
          <div class="basis-1/2 p-5 sm:p-9">
            <div class="mb-28 xl:mb-40 flex flex-col">
              <div class="v2-canopas-gradient-text desk-header-3">
                {{ weekly.title }}
              </div>
              <div class="mt-4 mobile-header-2-regular text-black-60">
                {{ weekly.content }}
              </div>
            </div>
            <div class="flex flex-row justify-between">
              <span class="mobile-header-3-semibold text-black-60">{{
                weekly.author
              }}</span>
              <span class="sub-h1-regular text-black-60">{{
                weekly.readtime
              }}</span>
            </div>
          </div>
        </div>
      </li>
    </ul>
  </section>
  <!-- Desktop UI end -->
</template>

<script>
import android from "@/assets/images/contributions/weekly/android.webp";
import flutter from "@/assets/images/contributions/weekly/flutter.webp";
import ios from "@/assets/images/contributions/weekly/ios.webp";
import web from "@/assets/images/contributions/weekly/web.webp";
import { elementInViewPort } from "@/utils";
import { Swiper, SwiperSlide } from "swiper/vue";
import { EffectCards } from "swiper/modules";
import { openBlog } from "@/utils.js";
import "swiper/css";

export default {
  data() {
    return {
      openBlog,
      swiper: null,
      activeIndex: 0,
      modules: [EffectCards],
      weeklies: [
        {
          category: "Android Weekly",
          image: android,
          author: "Radhika S",
          readtime: "2 min read",
          title: "Android Stack Weekly",
          content:
            "Welcome to Android Weekly — a newsletter on new development and updates of Android universe curated by Canopas team, delivered every Monday.",
          color: "bg-lavender-blush",
          url: "https://blog.canopas.com/tagged/canopas-android-weekly",
          translate: 0,
          scale: 1,
          event: "tap_contribution_android_weekly",
        },
        {
          category: "iOS Weekly",
          image: ios,
          author: "Amisha I",
          readtime: "2 min read",
          title: "iOS Stack Weekly",
          content:
            "Welcome to iOS Weekly — a newsletter on new development and updates of the iOS universe curated by Canopas team, delivered every Monday.",
          color: "bg-white-smoke",
          url: "https://blog.canopas.com/tagged/canopas-ios-weekly",
          translate: 40,
          scale: 1,
          event: "tap_contribution_ios_weekly",
        },
        {
          category: "Web Weekly",
          image: web,
          author: "Sumita K",
          readtime: "2 min read",
          title: "Web Stack Weekly",
          content:
            "Welcome to Web weekly — a weekly newsletter on new development and updates of Web universe curated by Canopas team, delivered every Monday.",
          color: "bg-lavender-blush",
          url: "https://blog.canopas.com/tagged/canopas-web-weekly",
          translate: 80,
          scale: 1,
          event: "tap_contribution_web_weekly",
        },
        {
          category: "Flutter Weekly",
          image: flutter,
          author: "Jimmy S",
          readtime: "2 min read",
          title: "Flutter Stack Weekly",
          content:
            "Welcome to Flutter Weekly — a newsletter on new development and updates of Flutter universe curated by Canopas team, delivered every Monday.",
          color: "bg-white-smoke",
          url: "https://blog.canopas.com/tagged/canopas-flutter-weekly",
          translate: 120,
          scale: 1,
          event: "tap_contribution_flutter_weekly",
        },
      ],
    };
  },
  components: {
    Swiper,
    SwiperSlide,
  },
  mounted() {
    if (window.innerWidth > 767) {
      window.addEventListener("scroll", this.handleScroll);
    }
  },
  unmounted() {
    if (window.innerWidth > 767) {
      window.removeEventListener("scroll", this.handleScroll);
    }
  },
  methods: {
    scrollToCard(index) {
      this.activeIndex = index;
      document.getElementById("scrollContainer").scrollTo({
        top: index * this.$refs["weekly-" + index][0].offsetHeight,
        behavior: "smooth",
      });
    },
    handleScroll() {
      const elementIdx = elementInViewPort(this.$refs);
      const index = elementIdx
        ? parseInt(elementIdx.charAt(elementIdx.length - 1))
        : 0;

      this.weeklies.forEach((weekly, j) => {
        if (j === index) {
          weekly.scale = 1;
          weekly.translate = 0;
        } else {
          weekly.scale = 1 - (index - j) * 0.03;
          weekly.translate = (index - j) * -40;
        }
      });
    },
  },
};
</script>
<style lang="postcss" scoped>
@import "swiper/css/effect-cards";

.swiper {
  @apply h-[460px] w-[380px];
}
.swiper-slide {
  @apply flex items-center justify-center;
}
</style>
