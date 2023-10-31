<template>
  <div
    class="my-20 xll:container md:mt-[8.063rem] lg:mt-[14.063rem] md:mb-12 xl:mb-20"
  >
    <div
      class="container mb-20 text-center font-inter-bold text-[1.875rem] leading-[2.4375rem] text-black-core/[0.87] lg:text-[3.4375rem] lg:leading-[5.15625rem]"
    >
      Explore Our UI/UX Design
    </div>
    <!-- Mobile UI start -->
    <div class="my-10 block md:hidden">
      <div
        class="animate-gridAnimationReverse"
        :class="[pausedId == 0 ? gridAnimation : 'animation-paused']"
      >
        <div
          class="animate-gridAnimation"
          :class="[pausedId == 0 ? 'animation-running' : 'animation-paused']"
        >
          <div class="ml-40 flex" v-if="mobileGrid.length > 0">
            <div
              v-for="(item, index) in mobileGrid"
              :key="index"
              class="ml-8 flex flex-[0_0_320px] flex-col justify-center rounded-2xl text-left text-[1rem] leading-[1.125rem] sm:flex-[0_0_480px] md:text-[1.0625rem] md:leading-[1.5rem] lg:flex-[0_0_528px] lg:text-[1.1875rem] lg:leading-[1.875rem]"
              :class="pausedId == index ? 'scale-[0.97] cursor-pointer' : ''"
              @mouseover="pausedId = index"
              @mouseleave="pausedId = 0"
              @touchstart.passive="pausedId = index"
              @touchend="pausedId = 0"
              :ref="'card-1-' + index"
              @click="
                openBlog(item.link, 'tap_contribution_ui/ux_design_explore')
              "
            >
              <div class="image-container sm:px-5">
                <img
                  @click.native="$mixpanel.track('tap_explore_design')"
                  :src="item.images[0]"
                  :srcset="`${item.images[0]} 400w, ${item.images[1]} 800w`"
                  alt="explore_uiux_design"
                  class="h-full w-full rounded-lg object-cover"
                  :class="item.classname"
                  loading="lazy"
                />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <!-- Mobile UI end -->
    <!-- Desktop UI start -->

    <div class="mt-12 hidden md:block xl:mt-16">
      <div class="swiper-content -mt-6 w-full xll:container">
        <swiper
          :slidesPerView="3"
          :centeredSlides="true"
          :loop="true"
          :autoplay="{
            delay: 2000,
            disableOnInteraction: false,
          }"
          :speed="3000"
          :zoom="true"
          :spaceBetween="0"
          :modules="modules"
          :navigation="true"
          class="swiper-container !pt-[30px]"
        >
          <swiper-slide v-for="(item, index) in data" :key="index">
            <div
              class="image-content swiper-zoom-container cursor-pointer"
              @click="
                openBlog(item.link, 'tap_contribution_ui/ux_design_explore')
              "
            >
              <img
                :src="item.images[0]"
                :srcset="`${item.images[0]} 400w,${item.images[1]} 800w`"
                class="h-full w-full rounded-lg object-cover drop-shadow-md lg:h-fit lg:w-fit"
                alt="DesignImage"
                loading="lazy"
              />
            </div>
          </swiper-slide>
        </swiper>
      </div>
      <div class="swiper-content mt-4 xll:container xl:mt-8">
        <swiper
          dir="rtl"
          :slidesPerView="4"
          :centeredSlides="true"
          :loop="true"
          :autoplay="{
            delay: 2000,
            disableOnInteraction: false,
          }"
          :speed="3000"
          :spaceBetween="20"
          :modules="modules"
          :navigation="true"
          class="swiper-container"
        >
          <swiper-slide v-for="(item, index) in data2" :key="index">
            <div
              class="cursor-pointer"
              @click="
                openBlog(item.link, 'tap_contribution_ui/ux_design_explore')
              "
            >
              <img
                :src="item.images[0]"
                :srcset="`${item.images[0]} 400w,${item.images[1]} 800w`"
                class="h-full w-full rounded-lg object-cover drop-shadow-md lg:h-fit lg:w-fit"
                alt="DesignImage"
                loading="lazy"
              />
            </div>
          </swiper-slide>
        </swiper>
      </div>
    </div>
  </div>
  <!-- Desktop UI end -->
</template>

<script>
import { openBlog } from "@/utils.js";

import chatapp from "@/assets/images/contributions/explore/1-400w.webp";
import chatapp800w from "@/assets/images/contributions/explore/1-800w.webp";

import elearningapp from "@/assets/images/contributions/explore/2-400w.webp";
import elearningapp800w from "@/assets/images/contributions/explore/2-800w.webp";

import foodorderingapp from "@/assets/images/contributions/explore/3-400w.webp";
import foodorderingapp800w from "@/assets/images/contributions/explore/3-800w.webp";

import irctcapp from "@/assets/images/contributions/explore/4-400w.webp";
import irctcapp800w from "@/assets/images/contributions/explore/4-800w.webp";

import meditationapp from "@/assets/images/contributions/explore/5-400w.webp";
import meditationapp800w from "@/assets/images/contributions/explore/5-800w.webp";

import moviestreamingapp from "@/assets/images/contributions/explore/6-400w.webp";
import moviestreamingapp800w from "@/assets/images/contributions/explore/6-800w.webp";

import shoppingapp from "@/assets/images/contributions/explore/7-400w.webp";
import shoppingapp800w from "@/assets/images/contributions/explore/7-800w.webp";

import socialnetworkingapp from "@/assets/images/contributions/explore/8-400w.webp";
import socialnetworkingapp800w from "@/assets/images/contributions/explore/8-800w.webp";

import AspectRatio from "@/components/utils/AspectRatio.vue";
import { Zoom, Pagination, Autoplay } from "swiper/modules";
import { Swiper, SwiperSlide } from "swiper/vue";
import "swiper/css";
import "swiper/css/pagination";
import "swiper/css/zoom";
export default {
  data() {
    return {
      modules: [Zoom, Pagination, Autoplay],
      openBlog,
      activeIndex: 0,
      gridAnimation: "animation-paused",
      gridAnimationReverse: "animation-running",
      pausedId: 0,
      lastWidth: null,
      lastScrollY: 0,
      viewTop: 0,
      mobileGrid: [],
      grid1: [],
      grid2: [],
      data2: [],
      data: [
        {
          images: [chatapp, chatapp800w],
          link: "https://dribbble.com/shots/20554178-Chat-app",
          classname: "rotate-6",
        },
        {
          images: [elearningapp, elearningapp800w],
          link: "https://dribbble.com/shots/20494364-E-learning-app",
          classname: "-rotate-6",
        },
        {
          images: [foodorderingapp, foodorderingapp800w],
          link: "https://dribbble.com/shots/20513671-Food-Delivery-App",
          classname: "rotate-6",
        },
        {
          images: [irctcapp, irctcapp800w],
          link: "https://dribbble.com/shots/20455362-Journey-Booking-App",
          classname: "-rotate-6",
        },
        {
          images: [meditationapp, meditationapp800w],
          link: "https://dribbble.com/shots/20749246-Meditation-app",
          classname: "rotate-6",
        },
        {
          images: [moviestreamingapp, moviestreamingapp800w],
          link: "https://dribbble.com/shots/20597815-Movie-Streaming-app",
          classname: "-rotate-6",
        },
        {
          images: [shoppingapp, shoppingapp800w],
          link: "https://dribbble.com/shots/20619022-Shopping-App",
          classname: "rotate-6",
        },
        {
          images: [socialnetworkingapp, socialnetworkingapp800w],
          link: "https://dribbble.com/shots/20577294-Social-Networking-App",
          classname: "-rotate-6",
        },
      ],
    };
  },
  components: {
    Swiper,
    SwiperSlide,
    AspectRatio,
  },
  inject: ["mixpanel"],
  created() {
    this.setGrid();
    this.setData();
  },
  mounted() {
    this.appendKeyframes();
    window.addEventListener("scroll", this.handleScroll);
  },
  unmounted() {
    window.removeEventListener("scroll", this.handleScroll);
  },
  methods: {
    handleScroll() {
      let scrollTop = window.scrollY;
      if (scrollTop < this.lastScrollY) {
        this.gridAnimation = "animation-running";
        this.gridAnimationReverse = "animation-paused";
      } else {
        this.gridAnimation = "animation-paused";
        this.gridAnimationReverse = "animation-running";
      }
      this.lastScrollY = scrollTop;
    },
    setData() {
      this.data2 = [...this.data.slice(4, 8), ...this.data.slice(0, 4)];
      this.data2 = new Array(50).fill(this.data2).flat();
    },
    setGrid() {
      var length = this.data.length;
      this.grid1 = new Array(3).fill(this.data.slice(0, length / 2)).flat();
      this.grid2 = new Array(3)
        .fill(this.data.slice(length / 2, length))
        .flat();

      this.mobileGrid = new Array(3).fill(this.data).flat();
    },
    appendKeyframes() {
      let offsetWidth =
        window.innerWidth < 768 ? this.$refs["card-1-0"].offsetWidth : "";

      let length =
        window.innerWidth < 768
          ? this.data.length * 2 // for mobile
          : this.data.length;

      let width = ((offsetWidth + 32) * length) / 2;

      const head = document.getElementsByTagName("head")[0];
      const styles = document.getElementsByTagName("style");
      let found = false;

      // find style titled with explore, if found replace width according to data
      for (var i = 0; i < styles.length; i++) {
        if (styles[i].title == "explore") {
          found = true;
          styles[i].innerText = styles[i].innerText.replaceAll(
            this.lastWidth,
            width,
          );
        }
      }
      // else append keyframes
      if (!found) {
        let keyframes =
          `@keyframes scroll { 0% {transform: translateX(0px)} 100% {transform:translateX(-` +
          width +
          `px)}}\n` +
          `@-webkit-keyframes scroll { 0% {transform: translateX(0px)}  100% {transform:translateX(-` +
          width +
          `px)}}\n` +
          `@-webkit-keyframes scroll-reverse { 0% {transform:translateX(-` +
          width +
          `px)} 100% {transform: translateX(0px)}}\n` +
          `@keyframes scroll-reverse { 0% {transform:translateX(-` +
          width +
          `px)} 100% {transform: translateX(0px)}}`;

        let style = document.createElement("style");
        style.title = "explore";
        style.innerHTML = keyframes;

        head.appendChild(style);
      }

      this.lastWidth = width;
    },
  },
};
</script>
<style lang="postcss" scoped>
@import "swiper/css";
@import "swiper/css/pagination";
@import "swiper/css/zoom";

.swiper-slide:not(.swiper-slide-active) .image-content {
  @apply scale-[0.8];
}

.swiper-slide .swiper-zoom-container {
  @apply origin-bottom duration-[3000ms];
}
</style>
