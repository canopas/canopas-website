<template>
  <div>
    <div
      class="tw-container tw-flex tw-flex-col tw-items-center tw-gap-y-2.5 tw-mt-24"
    >
      <div>
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

    <div class="swiper-content lg:tw-hidden tw-mb-10">
      <div
        class="lg:tw-hidden swiper-pagination tw-flex tw-justify-around tw-mx-auto tw-mb-[3.75rem] tw-w-[100%] md:tw-w-[90%] xl:tw-w-[65%] 2xl:tw-w-[55%] xll:tw-w-[30%] tw-rounded-[30px] tw-py-[5px]"
      >
        <a
          v-for="(weekly, index) in weeklies"
          :key="index"
          target="_blank"
          class="lg:tw-hidden tw-flex-[16.66%] tw-gap-x-4 tw-m-0 tw-w-full tw-px-[5px] tw-py-[0.7rem] md:tw-py-[15px] tw-text-center tw-tracking-[1px] tw-text-[0.75rem] sm:tw-text-[1rem] lg:tw-text-[1.25rem] tw-leading-[0.938rem] lg:tw-leading-[1.1875rem] tw-font-inter-regular portfolio-nav lg:hover:tw-bg-gradient-[270.11deg] lg:hover:tw-from-[#ff835b] lg:hover:tw-to-[#f2709c] lg:hover:tw-text-white"
          :class="
            activeSlide == index
              ? 'v2-canopas-gradient-text tw-px-1.5 !tw-font-inter-bold'
              : 'tw-border-[1px] tw-border-solid tw-border-transparent tw-text-black-core/[0.6]'
          "
          @click="slideTo(index), activeSlide == index"
          :href="weekly.url"
        >
          {{ weekly.category }}
        </a>
      </div>
      <swiper
        :slidesPerView="1.1"
        :spaceBetween="0"
        :effect="'cards'"
        :grabCursor="true"
        :pagination="pagination"
        @swiper="setSwiperRef"
        class="swiper-container pagination-slider"
      >
        <swiper-slide
          v-for="(weekly, index) in weeklies"
          :key="index"
          @touchstart.passive="false"
          @touchmove.passive="true, onSlideChange(index)"
        >
          <a class="tw-w-[380px] tw-h-[460px]" :href="weekly.url">
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
          </a>
        </swiper-slide>
      </swiper>
    </div>

    <div
      class="tw-hidden lg:tw-flex tw-justify-around tw-m-auto tw-mb-[7rem] tw-mt-8 lg:tw-my-[3.75rem] tw-w-[100%] md:tw-w-[90%] xl:tw-w-[65%] 2xl:tw-w-[55%] xll:tw-w-[30%] tw-rounded-[30px] tw-py-[5px]"
    >
      <a
        v-for="(weekly, index) in weeklies"
        :key="index"
        target="_blank"
        class="tw-hidden lg:tw-block tw-flex-[16.66%] tw-gap-x-4 tw-m-0 tw-w-full tw-rounded-[25px] tw-px-[5px] tw-py-[0.7rem] md:tw-py-[15px] tw-text-center tw-tracking-[1px] tw-text-[0.75rem] sm:tw-text-[1rem] lg:tw-text-[1.25rem] tw-leading-[0.938rem] lg:tw-leading-[1.1875rem] tw-font-inter-regular portfolio-nav lg:hover:tw-bg-gradient-[270.11deg] lg:hover:tw-from-[#ff835b] lg:hover:tw-to-[#f2709c] lg:hover:tw-text-white"
        :class="
          activeIndex == index || isScrolled
            ? 'gradient-border-btn tw-px-1.5 !tw-font-inter-bold !tw-text-[#f2709c] hover:!tw-text-white'
            : 'tw-border-[1px] tw-border-solid tw-border-transparent tw-text-black-core/[0.6]'
        "
        :href="weekly.url"
        @click="changeCard(index)"
      >
        {{ weekly.category }}
      </a>
    </div>

    <div
      ref="scrollContainer"
      @scroll="handleScroll"
      class="tw-relative tw-hidden lg:tw-block tw-h-[52vh] tw-overflow-scroll hidden-scrollbar"
    >
      <div
        v-for="(weekly, index) in weeklies"
        :key="index"
        :ref="'divItems' + index"
        :class="['weekly', { active: activeIndex === index }]"
        class="tw-w-[1060px] tw-h-[466px] tw-mx-auto tw-sticky tw-overflow-hidden tw-h-[50vh] tw-top-[2vh]"
      >
        <div :class="weekly.color" class="tw-flex tw-flex-row tw-mx-auto">
          <div class="tw-basis-1/2 tw-p-[20px]">
            <img
              :src="[weekly.image]"
              alt="Weekly-stack-image"
              class="tw-h-full tw-w-full tw-object-cover"
            />
          </div>
          <div class="tw-basis-1/2 tw-p-5 sm:tw-p-9">
            <div class="tw-flex tw-flex-row tw-justify-between">
              <span
                class="v2-canopas-gradient-text tw-font-inter-semibold tw-text-[1.25rem] tw-leading-[1.875rem]"
                >{{ weekly.author }}</span
              >
              <span
                class="tw-font-inter-semibold tw-text-[1.25rem] tw-leading-[1.875rem] tw-text-black-core/[0.87]"
                >{{ weekly.readtime }}</span
              >
            </div>
            <div class="tw-flex tw-flex-col tw-mt-8 tw-mb-5">
              <div
                class="tw-text-[1.875rem] tw-leading-[2.25rem] tw-font-inter-semibold tw-text-black-bore/[0.87]"
              >
                {{ weekly.title }}
              </div>
              <div
                class="tw-mt-2.5 tw-text-black-core/[0.87] tw-font-inter-regular tw-text-[1.75rem] tw-leading-[2.625rem]"
              >
                {{ weekly.content }}
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import android from "@/assets/images/ContributionPage/weekly/android.webp";
import flutter from "@/assets/images/ContributionPage/weekly/flutter.webp";
import ios from "@/assets/images/ContributionPage/weekly/ios.webp";
import web from "@/assets/images/ContributionPage/weekly/web.webp";
import { elementInViewPort } from "@/utils";
import { Swiper, SwiperSlide } from "swiper/vue";
import SwiperCore, { Pagination, Autoplay, EffectCards } from "swiper";
SwiperCore.use([Pagination, Autoplay, EffectCards]);
export default {
  data() {
    return {
      swiper: null,
      activeIndex: 0,
      activeSlide: 0,
      isScrolled: false,
      weeklies: [
        {
          id: 1,
          category: "Android Weekly",
          image: android,
          author: "Radhika S",
          readtime: "2 min read",
          title: "Android Stack Weekly - Issue #70",
          content:
            "Welcome to iOS Weekly — a newsletter on new development and updates of the iOS universe curated by the Canopas team, delivered every Monday.",
          color: "tw-bg-[#FFEDF0]",
          url: "https://blog.canopas.com/tagged/canopas-android-weekly",
        },
        {
          id: 2,
          category: "iOS Weekly",
          image: ios,
          author: "Amisha L",
          readtime: "1 min read",
          title: "iOS Stack Weekly - Issue #70",
          content:
            "Welcome to iOS Weekly — a newsletter on new development and updates of the iOS universe curated by the Canopas team, delivered every Monday.",
          color: "tw-bg-[#FFF8EF]",
          url: "https://blog.canopas.com/tagged/canopas-ios-weekly",
        },
        {
          id: 3,
          category: "Web Weekly",
          image: web,
          author: "Sumita K",
          readtime: "2 min read",
          title: "Web Stack Weekly - Issue #70",
          content:
            "Welcome to iOS Weekly — a newsletter on new development and updates of the iOS universe curated by the Canopas team, delivered every Monday.",
          color: "tw-bg-[#EDEFFF]",
          url: "https://blog.canopas.com/tagged/canopas-web-weekly",
        },
        {
          id: 4,
          category: "Flutter Weekly",
          image: flutter,
          author: "Jimmy S",
          readtime: "3 min read",
          title: "Flutter Stack Weekly - Issue #70",
          content:
            "Welcome to iOS Weekly — a newsletter on new development and updates of the iOS universe curated by the Canopas team, delivered every Monday.",
          color: "tw-bg-[#EBF9EF]",
          url: "https://blog.canopas.com/tagged/canopas-flutter-weekly",
        },
      ],
      pagination: {
        el: ".swiper-pagination",
        clickable: true,
        renderBullet: function (index, className) {
          return (
            '<span class=" ' +
            className +
            '">' +
            this.weeklies[index].category +
            "</span>"
          );
        },
      },
    };
  },
  components: {
    Swiper,
    SwiperSlide,
  },
  mounted() {
    this.height = window.innerHeight;
    window.addEventListener("scroll", this.handleScroll);
  },
  unmounted() {
    window.removeEventListener("scroll", this.handleScroll);
  },

  methods: {
    // Active slide for desktop
    changeCard(index) {
      this.activeIndex = index;
      this.$refs.divItems[index].scrollIntoView();
    },
    handleScroll() {
      const scrollContainer = this.$refs.scrollContainer;
      const elementIdx = elementInViewPort(this.$refs);
      this.isScrolled = scrollContainer.scrollTop > 0;
    },

    // Active index for mobile view
    onSwiper(swiper) {
      this.swiper = swiper;
    },
    slideTo(index) {
      this.activeSlide = this.swiper.realIndex;
      this.swiperRef.slideTo(index - 1, 0);
    },
    onSlideChange(index) {
      this.activeSlide = index;
    },
    setSwiperRef(swiper) {
      this.swiperRef = swiper;
    },
  },
};
</script>
<style lang="postcss" scoped>
@import "swiper/css";
@import "swiper/css/pagination";
@import "swiper/css/effect-cards";

.swiper {
  @apply tw-w-[380px] tw-h-[460px];
}
.swiper-slide {
  @apply tw-flex tw-items-center tw-justify-center;
}

.swiper-pagination {
  @apply tw-border-[1px] tw-border-solid tw-border-transparent tw-text-black-core/[0.6];
  position: unset !important;
}
</style>
