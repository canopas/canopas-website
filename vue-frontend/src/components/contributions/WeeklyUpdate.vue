<template>
  <div>
    <div
      class="tw-container tw-mt-[3rem] tw-flex tw-flex-col tw-items-center tw-gap-y-2.5 lg:tw-mt-24"
    >
      <div class="tw-text-center">
        <span
          class="tw-font-inter-bold tw-text-[1.875rem] tw-leading-[2.4375rem] tw-text-black-core/[.87] lg:tw-text-[3.4375rem] lg:tw-leading-[5.15625rem]"
        >
          Our Weekly Tech Updates</span
        >
      </div>
      <div
        class="tw-mb-8 tw-mt-2.5 tw-text-center tw-font-inter-regular tw-text-[1rem] tw-leading-[1.5rem] tw-text-black-core/[0.87] lg:tw-mt-6 lg:tw-font-inter-medium lg:tw-text-[1.5rem] lg:tw-leading-[2.25rem] lg:tw-text-black-core/[0.6]"
      >
        Each week, we curate a hand-picked selection of the latest tech updates,
        delivering them straight to you. Immerse yourself in our weekly insights
        and equip yourself with the knowledge to navigate tomorrow.
      </div>
    </div>
    <!-- Mobile UI start -->
    <div class="swiper-content tw-mb-10 tw-block md:tw-hidden">
      <swiper
        :slidesPerView="1.1"
        :spaceBetween="0"
        :effect="'cards'"
        :grabCursor="true"
        class="swiper-container"
      >
        <swiper-slide v-for="(weekly, index) in weeklies" :key="index">
          <div
            class="tw-h-full tw-w-full tw-object-cover"
            @click="openUrl(weekly.url)"
          >
            <img
              :src="[weekly.image]"
              alt="Weekly-stack-image"
              class="tw-h-fit tw-w-fit tw-rounded-t-lg tw-object-contain"
            />
            <div class="!tw-rounded-b-lg tw-bg-[#F2F2F2] tw-p-[2rem]">
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
              <div class="tw-mt-4 tw-flex tw-flex-col">
                <div
                  class="tw-font-inter-medium tw-text-[1.25rem] tw-leading-[1.5rem] sm:tw-text-black-core/[0.87]"
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
      class="tw-m-auto tw-mb-[7rem] tw-mt-8 tw-hidden tw-justify-around tw-rounded-[30px] tw-py-[5px] md:tw-my-[3.75rem] md:tw-flex md:tw-w-[95%] lg:tw-w-[77%] xl:tw-w-[65%] 2xl:tw-w-[55%] xll:tw-w-[30%]"
    >
      <div
        v-for="(weekly, index) in weeklies"
        :key="index"
        target="_blank"
        class="tw-m-0 tw-hidden tw-w-full tw-flex-[16.66%] tw-gap-x-4 tw-rounded-[25px] tw-px-[5px] tw-py-[0.7rem] tw-text-center tw-font-inter-regular tw-text-[0.75rem] tw-leading-[0.938rem] tw-tracking-[1px] sm:tw-text-[1rem] md:tw-block md:tw-py-[15px] md:tw-text-[1.25rem] md:tw-leading-[1.1875rem] md:hover:tw-from-[#ff835b] md:hover:tw-to-[#f2709c] md:hover:tw-text-white md:hover:tw-bg-gradient-[270.11deg]"
        :class="
          activeIndex == index
            ? 'gradient-border-btn tw-px-1.5 !tw-font-inter-bold !tw-text-[#f2709c] hover:!tw-text-white'
            : 'tw-border-[1px] tw-border-solid tw-border-transparent tw-text-black-core/[0.6]'
        "
        @click="scrollToCard(index)"
      >
        {{ weekly.category }}
      </div>
    </div>
    <ul
      @scroll="handleScroll"
      id="scrollContainer"
      class="hidden-scrollbar tw-relative tw-hidden tw-h-[54.85vh] tw-overflow-y-scroll md:tw-block xl:tw-h-[65.5vh] 2xl:tw-h-[73vh] 2xll:tw-h-[76vh] xll:tw-h-[37vh]"
    >
      <li
        v-for="(weekly, index) in weeklies"
        :key="index"
        :ref="'weekly-' + index"
        :id="'weekly-' + index"
        @click="openUrl(weekly.url)"
        class="tw-origin-[center top] tw-sticky tw-top-[10px] tw-mx-[1%] tw-h-[45vh] tw-cursor-pointer tw-overflow-hidden tw-drop-shadow-xl lg:tw-container lg:tw-h-[45vh] xl:tw-mx-auto xl:tw-h-[55vh] 2xl:tw-h-[65vh] xll:tw-h-[33vh]"
        :style="{
          transform: `translateY(${weekly.translate}px) scale(${weekly.scale})`,
        }"
      >
        <div
          :class="weekly.color"
          class="tw-mx-auto tw-flex tw-h-[350px] tw-flex-row xl:tw-h-[450px] 2xl:tw-h-[515px]"
        >
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
                class="tw-font-inter-semibold tw-text-black-core/[0.87] md:tw-text-[1rem] md:tw-leading-[1.125rem] lg:tw-text-[1.125rem] lg:tw-leading-[1.5rem] xl:tw-text-[1.25rem] xl:tw-leading-[1.875rem]"
                >{{ weekly.readtime }}</span
              >
            </div>
            <div class="tw-mb-5 tw-mt-8 tw-flex tw-flex-col">
              <div
                class="tw-text-black-bore/[0.87] tw-font-inter-semibold md:tw-text-[1.125rem] md:tw-leading-[1.5rem] lg:tw-text-[1.5rem] lg:tw-leading-[1.875rem] xl:tw-text-[1.875rem] xl:tw-leading-[2.25rem]"
              >
                {{ weekly.title }}
              </div>
              <div
                class="tw-mt-2.5 tw-font-inter-regular tw-text-black-core/[0.87] md:tw-text-[1rem] md:tw-leading-[1.5rem] lg:tw-text-[1.25rem] lg:tw-leading-[1.875rem] xl:tw-text-[1.75rem] xl:tw-leading-[2.625rem]"
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
import android from "@/assets/images/contributions/weekly/android.webp";
import flutter from "@/assets/images/contributions/weekly/flutter.webp";
import ios from "@/assets/images/contributions/weekly/ios.webp";
import web from "@/assets/images/contributions/weekly/web.webp";
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
          translate: 40,
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
          translate: 80,
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
          translate: 120,
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
  },
};
</script>
<style lang="postcss" scoped>
@import "swiper/css";
@import "swiper/css/effect-cards";

.swiper {
  @apply tw-h-[460px] tw-w-[380px];
}
.swiper-slide {
  @apply tw-flex tw-items-center tw-justify-center;
}
</style>
