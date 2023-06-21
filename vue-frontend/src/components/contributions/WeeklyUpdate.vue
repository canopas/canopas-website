<template>
  <div>
    <div
      class="tw-container tw-flex tw-flex-col tw-items-center tw-gap-y-2.5 tw-mt-24"
    >
      <div class="tw-text-center">
        <span
          class="tw-text-[1.875rem] lg:tw-text-[3.4375rem] tw-leading-[2.4375rem] lg:tw-leading-[5.15625rem] tw-text-black-core/[.87] tw-font-inter-bold"
        >
          Our Weekly Tech Updates</span
        >
      </div>
      <div
        class="tw-mb-8 tw-mt-2.5 lg:tw-mt-6 tw-text-center tw-text-[1rem] lg:tw-text-[1.5rem] tw-leading-[1.5rem] lg:tw-leading-[2.25rem] tw-text-black-core/[0.87] lg:tw-text-black-core/[0.6] tw-font-inter-regular lg:tw-font-inter-medium"
      >
        Each week, we curate a hand-picked selection of the latest tech updates,
        delivering them straight to you. Immerse yourself in our weekly insights
        and equip yourself with the knowledge to navigate tomorrow.
      </div>
    </div>
    <!-- Mobile UI start -->
    <div class="swiper-content tw-block md:tw-hidden tw-mb-10">
      <div
        class="tw-flex tw-justify-around tw-mx-auto tw-mb-[3.75rem] tw-w-[70%] xs:tw-w-[90%] xl:tw-w-[65%] 2xl:tw-w-[55%] xll:tw-w-[30%] tw-rounded-[30px] tw-py-[5px]"
      >
        <a
          v-for="(weekly, index) in weeklies"
          :key="index"
          target="_blank"
          class="md:tw-hidden tw-flex-[16.66%] tw-gap-x-4 tw-m-0 tw-w-full tw-px-[5px] tw-py-[0.7rem] md:tw-py-[15px] tw-text-center tw-tracking-[1px] tw-text-[0.75rem] sm:tw-text-[1rem] md:tw-text-[1.25rem] tw-leading-[0.938rem] md:tw-leading-[1.1875rem] tw-font-inter-regular md:hover:tw-bg-gradient-[270.11deg] md:hover:tw-from-[#ff835b] md:hover:tw-to-[#f2709c] md:hover:tw-text-white"
          :class="
            activeIndex == index
              ? 'v2-canopas-gradient-text tw-px-1.5 !tw-font-inter-bold'
              : 'tw-border-[1px] tw-border-solid tw-border-transparent tw-text-black-core/[0.6]'
          "
          @click="slideTo(index)"
        >
          {{ weekly.category }}
        </a>
      </div>
      <swiper
        :slidesPerView="1.1"
        :spaceBetween="0"
        :effect="'cards'"
        :grabCursor="true"
        @swiper="setSwiperRef"
        @slideChange="activeIndex = swiperRef.activeIndex"
        class="swiper-container"
      >
        <swiper-slide v-for="(weekly, index) in weeklies" :key="index">
          <div
            class="tw-w-full tw-h-full tw-object-cover"
            @click="openUrl(weekly.url)"
          >
            <img
              :src="[weekly.image]"
              alt="Weekly-stack-image"
              class="tw-h-fit tw-w-fit tw-object-contain tw-rounded-t-lg"
            />
            <div class="tw-bg-[#F2F2F2] !tw-rounded-b-lg tw-p-[2rem]">
              <div class="tw-flex tw-flex-row tw-justify-between">
                <span
                  class="v2-canopas-gradient-text tw-font-inter-regular tw-text-[0.875rem] tw-leading-[1.3125rem]"
                  >{{ weekly.author }}</span
                >
                <span
                  class="tw-text-[0.875rem] tw-leading-[1.3125rem] tw-text-black-core/[0.87]"
                  >{{ weekly.readtime }}</span
                >
              </div>
              <div class="tw-flex tw-flex-col tw-mt-4">
                <div
                  class="tw-text-[1.25rem] tw-leading-[1.5rem] tw-font-inter-medium sm:tw-text-black-core/[0.87]"
                >
                  {{ weekly.title }}
                </div>
                <div class="tw-mt-2.5 tw-text-black-core/[0.87]">
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
    <div
      class="tw-hidden md:tw-flex tw-justify-around tw-m-auto tw-mb-[7rem] tw-mt-8 md:tw-my-[3.75rem] tw-w-[100%] md:tw-w-[95%] xl:tw-w-[65%] 2xl:tw-w-[55%] xll:tw-w-[30%] tw-rounded-[30px] tw-py-[5px]"
    >
      <a
        v-for="(weekly, index) in weeklies"
        :key="index"
        target="_blank"
        class="tw-hidden md:tw-block tw-flex-[16.66%] tw-gap-x-4 tw-m-0 tw-w-full tw-rounded-[25px] tw-px-[5px] tw-py-[0.7rem] md:tw-py-[15px] tw-text-center tw-tracking-[1px] tw-text-[0.75rem] sm:tw-text-[1rem] md:tw-text-[1.25rem] tw-leading-[0.938rem] md:tw-leading-[1.1875rem] tw-font-inter-regular md:hover:tw-bg-gradient-[270.11deg] md:hover:tw-from-[#ff835b] md:hover:tw-to-[#f2709c] md:hover:tw-text-white"
        :class="
          activeIndex == index
            ? 'gradient-border-btn tw-px-1.5 !tw-font-inter-bold !tw-text-[#f2709c] hover:!tw-text-white'
            : 'tw-border-[1px] tw-border-solid tw-border-transparent tw-text-black-core/[0.6]'
        "
        @click="scrollToCard(index)"
      >
        {{ weekly.category }}
      </a>
    </div>
    <ul
      @scroll="handleScroll"
      id="scrollContainer"
      class="tw-relative tw-hidden md:tw-block tw-h-[50vh] tw-overflow-scroll hidden-scrollbar"
    >
      <li
        v-for="(weekly, index) in weeklies"
        :key="index"
        :ref="'weekly-' + index"
        :id="'weekly-' + index"
        @click="openUrl(weekly.url)"
        class="tw-sticky tw-top-[10px] tw-mx-auto tw-container tw-h-[50vh] tw-overflow-hidden tw-origin-[center top] tw-cursor-pointer"
        :style="{
          transform: `translateY(${weekly.translate}px) scale(${weekly.scale})`,
        }"
      >
        <div :class="weekly.color" class="tw-flex tw-flex-row tw-mx-auto">
          <div class="tw-basis-1/2 tw-p-[20px]">
            <img
              :src="[weekly.image]"
              alt="Weekly-stack-image"
              class="tw-h-full tw-w-fit tw-object-cover"
            />
          </div>
          <div class="tw-basis-1/2 tw-p-5 sm:tw-p-9">
            <div class="tw-flex tw-flex-row tw-justify-between">
              <span
                class="v2-canopas-gradient-text tw-font-inter-semibold md:tw-text-[1rem] md:tw-leading-[1.125rem] lg:tw-text-[1.125rem] lg:tw-leading-[1.5rem] xl:tw-text-[1.25rem] xl:tw-leading-[1.875rem]"
                >{{ weekly.author }}</span
              >
              <span
                class="tw-font-inter-semibold md:tw-text-[1rem] md:tw-leading-[1.125rem] lg:tw-text-[1.125rem] lg:tw-leading-[1.5rem] xl:tw-text-[1.25rem] xl:tw-leading-[1.875rem] tw-text-black-core/[0.87]"
                >{{ weekly.readtime }}</span
              >
            </div>
            <div class="tw-flex tw-flex-col tw-mt-8 tw-mb-5">
              <div
                class="md:tw-text-[1.125rem] md:tw-leading-[1.5rem] lg:tw-text-[1.5rem] lg:tw-leading-[1.875rem] xl:tw-text-[1.875rem] xl:tw-leading-[2.25rem] tw-font-inter-semibold tw-text-black-bore/[0.87]"
              >
                {{ weekly.title }}
              </div>
              <div
                class="tw-mt-2.5 tw-text-black-core/[0.87] tw-font-inter-regular md:tw-text-[1rem] md:tw-leading-[1.5rem] lg:tw-text-[1.25rem] lg:tw-leading-[1.875rem] xl:tw-text-[1.75rem] xl:tw-leading-[2.625rem]"
              >
                {{ weekly.content }}
              </div>
            </div>
          </div>
        </div>
      </li>
    </ul>
    <!-- Desktop UI end -->
  </div>
</template>

<script>
import android from "@/assets/images/ContributionPage/weekly/android.webp";
import flutter from "@/assets/images/ContributionPage/weekly/flutter.webp";
import ios from "@/assets/images/ContributionPage/weekly/ios.webp";
import web from "@/assets/images/ContributionPage/weekly/web.webp";
import { elementInViewPort } from "@/utils";
import { Swiper, SwiperSlide } from "swiper/vue";
import SwiperCore, { EffectCards } from "swiper";

SwiperCore.use([EffectCards]);

export default {
  data() {
    return {
      swiper: null,
      activeIndex: 0,
      weeklies: [
        {
          category: "Android Weekly",
          image: android,
          author: "Radhika S",
          readtime: "2 min read",
          title: "Android Stack Weekly",
          content:
            "Welcome to Android Weekly — a newsletter on new development and updates of Android universe curated by Canopas team, delivered every Monday.",
          color: "tw-bg-[#FFEDF0]",
          url: "https://blog.canopas.com/tagged/canopas-android-weekly",
          translate: 0,
          scale: 1,
        },
        {
          category: "iOS Weekly",
          image: ios,
          author: "Amisha L",
          readtime: "2 min read",
          title: "iOS Stack Weekly",
          content:
            "Welcome to iOS Weekly — a newsletter on new development and updates of the iOS universe curated by Canopas team, delivered every Monday.",
          color: "tw-bg-[#FFF8EF]",
          url: "https://blog.canopas.com/tagged/canopas-ios-weekly",
          translate: 30,
          scale: 1,
        },
        {
          category: "Web Weekly",
          image: web,
          author: "Sumita K",
          readtime: "2 min read",
          title: "Web Stack Weekly",
          content:
            "Welcome to Web weekly — a weekly newsletter on new development and updates of Web universe curated by Canopas team, delivered every Monday.",
          color: "tw-bg-[#EDEFFF]",
          url: "https://blog.canopas.com/tagged/canopas-web-weekly",
          translate: 60,
          scale: 1,
        },
        {
          category: "Flutter Weekly",
          image: flutter,
          author: "Jimmy S",
          readtime: "2 min read",
          title: "Flutter Stack Weekly",
          content:
            "Welcome to Flutter Weekly — a newsletter on new development and updates of Flutter universe curated by Canopas team, delivered every Monday.",
          color: "tw-bg-[#EBF9EF]",
          url: "https://blog.canopas.com/tagged/canopas-flutter-weekly",
          translate: 90,
          scale: 1,
        },
      ],
    };
  },
  components: {
    Swiper,
    SwiperSlide,
  },
  mounted() {
    if (window.innerWidth > 992) {
      window.addEventListener("scroll", this.handleScroll);
    }
  },
  unmounted() {
    window.removeEventListener("scroll", this.handleScroll);
  },
  methods: {
    openUrl(link) {
      window.open(link, "_blank");
    },
    scrollToCard(index) {
      this.activeIndex = index;
      document.getElementById("scrollContainer").scrollTo({
        top: index * this.$refs["weekly-" + index][0].offsetHeight,
        behavior: "smooth",
      });
    },
    handleScroll() {
      const elementIdx = elementInViewPort(this.$refs);
      if (!elementIdx) {
        return;
      }

      const index = elementIdx
        ? parseInt(elementIdx.charAt(elementIdx.length - 1))
        : 0;

      for (var j = index; j >= 0; j--) {
        this.weeklies[j].scale =
          1 -
          (document.getElementById("scrollContainer").scrollTop * (index - j)) /
            (index * 10000);
      }

      // get current active element from offset
      const counts = {};
      this.weeklies.forEach((_, index) => {
        var offsetTop = this.$refs["weekly-" + index][0].offsetTop;
        counts[offsetTop] = (counts[offsetTop] || 0) + 1;
      });

      this.activeIndex = this.weeklies.length - Object.keys(counts).length;
    },
    onSlideChange() {
      this.activeIndex = this.swiperRef.activeIndex;
    },
    setSwiperRef(swiper) {
      this.swiperRef = swiper;
    },
    slideTo(index) {
      this.activeIndex = index;
      this.swiperRef.slideTo(index, 500);
    },
  },
};
</script>
<style lang="postcss" scoped>
@import "swiper/css";
@import "swiper/css/effect-cards";

.swiper {
  @apply tw-w-[380px] tw-h-[460px];
}

.swiper-slide {
  @apply tw-flex tw-items-center tw-justify-center;
}
</style>
