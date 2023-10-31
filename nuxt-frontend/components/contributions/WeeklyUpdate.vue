<template>
  <section class="mb-[200px]">
    <div class="container mt-12 flex flex-col items-center gap-y-2.5 lg:mt-24">
      <div class="text-center">
        <span
          class="font-inter-bold text-[1.875rem] leading-[2.4375rem] text-black-core/[.87] lg:text-[3.4375rem] lg:leading-[5.15625rem]"
        >
          Our Weekly Tech Updates</span
        >
      </div>
      <div
        class="mb-8 mt-2.5 text-center font-inter-regular text-base text-black-core/[0.87] lg:mt-6 lg:font-inter-medium lg:text-[1.5rem] lg:leading-9 lg:text-black-core/[0.6]"
      >
        Each week, we curate a hand-picked selection of the latest tech updates,
        delivering them straight to you. Immerse yourself in our weekly insights
        and equip yourself with the knowledge to navigate tomorrow.
      </div>
    </div>
    <!-- Mobile UI start -->
    <div class="swiper-content mb-10 block md:hidden">
      <swiper
        :slidesPerView="1.1"
        :spaceBetween="0"
        :effect="'cards'"
        :grabCursor="true"
        :modules="modules"
        class="swiper-container"
      >
        <swiper-slide v-for="(weekly, index) in weeklies" :key="index">
          <div
            class="h-full w-full object-cover"
            @click="openBlog(weekly.url, weekly.event)"
          >
            <img
              :src="[weekly.image]"
              alt="Weekly-stack-image"
              class="h-fit w-fit rounded-t-lg object-contain"
            />
            <div class="!rounded-b-lg bg-[#F2F2F2] p-8">
              <div class="flex flex-row justify-between">
                <span
                  class="v2-canopas-gradient-text font-inter-regular text-[0.875rem] leading-[1.3125rem]"
                  >{{ weekly.author }}</span
                >
                <span
                  class="text-[0.875rem] leading-[1.3125rem] text-black-core/[0.87]"
                  >{{ weekly.readtime }}</span
                >
              </div>
              <div class="mt-4 flex flex-col">
                <div
                  class="font-inter-medium text-[1.25rem] leading-6 sm:text-black-core/[0.87]"
                >
                  {{ weekly.title }}
                </div>
                <div class="mt-2.5 text-black-core/[0.87]">
                  {{ weekly.content }}
                </div>
              </div>
            </div>
          </div>
        </swiper-slide>
      </swiper>
    </div>
    <!-- Mobile UI end -->
    <!-- Desktop UI start -->
    <ul id="scrollContainer" class="w-full mt-20 lg:mt-28 hidden md:block">
      <li
        v-for="(weekly, index) in weeklies"
        :key="index"
        :ref="'weekly-' + index"
        :id="'weekly-' + index"
        @click="openBlog(weekly.url, weekly.event)"
        class="origin-[center top] sticky top-[8.625rem] xll:top-[15.625rem] 3xl:top-[20.625rem] mx-[1%] h-[270px] lg:h-[390px] cursor-pointer overflow-hidden drop-shadow-xl md:container lg:h-[365px] xl:mx-auto xl:h-[480px] 2xl:h-[590px] xll:h-[546px]"
        :style="{
          transform: `translateY(${weekly.translate}px) scale(${weekly.scale})`,
        }"
      >
        <div
          :class="weekly.color"
          class="mx-auto flex h-[350px] flex-row xl:h-[450px] 2xl:h-[515px]"
        >
          <div class="basis-1/2 p-5">
            <img
              :src="[weekly.image]"
              alt="Weekly-stack-image"
              class="h-[72%] lg:h-full w-fit object-cover"
            />
          </div>
          <div class="basis-1/2 p-5 sm:p-9">
            <div class="flex flex-row justify-between">
              <span
                class="v2-canopas-gradient-text font-inter-semibold md:text-[1rem] md:leading-[1.125rem] lg:text-[1.125rem] lg:leading-6 xl:text-[1.25rem] xl:leading-[1.875rem]"
                >{{ weekly.author }}</span
              >
              <span
                class="font-inter-semibold text-black-core/[0.87] md:text-[1rem] md:leading-[1.125rem] lg:text-[1.125rem] lg:leading-6 xl:text-[1.25rem] xl:leading-[1.875rem]"
                >{{ weekly.readtime }}</span
              >
            </div>
            <div class="mb-5 mt-8 flex flex-col">
              <div
                class="text-black-bore/[0.87] font-inter-semibold md:text-[1.125rem] md:leading-6 lg:text-[1.5rem] lg:leading-[1.875rem] xl:text-[1.875rem] xl:leading-9"
              >
                {{ weekly.title }}
              </div>
              <div
                class="mt-2.5 font-inter-regular text-black-core/[0.87] md:text-[1rem] md:leading-6 lg:text-[1.25rem] lg:leading-[1.875rem] xl:text-[1.75rem] xl:leading-[2.625rem]"
              >
                {{ weekly.content }}
              </div>
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
          color: "bg-[#FFEDF0]",
          url: "https://blog.canopas.com/tagged/canopas-android-weekly",
          translate: 0,
          scale: 1,
          event: "tap_contribution_android_weekly",
        },
        {
          category: "iOS Weekly",
          image: ios,
          author: "Amisha L",
          readtime: "2 min read",
          title: "iOS Stack Weekly",
          content:
            "Welcome to iOS Weekly — a newsletter on new development and updates of the iOS universe curated by Canopas team, delivered every Monday.",
          color: "bg-[#FFF8EF]",
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
          color: "bg-[#EDEFFF]",
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
          color: "bg-[#EBF9EF]",
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
