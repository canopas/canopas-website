<template>
  <div class="xll:tw-container tw-my-20 md:tw-mb-12 xl:tw-mb-20">
    <div
      class="tw-container tw-mb-20 tw-text-center tw-font-inter-bold tw-text-[1.875rem] lg:tw-text-[3.4375rem] tw-leading-[2.4375rem] lg:tw-leading-[5.15625rem] tw-text-black-core/[0.87]"
    >
      Explore Our UI/UX Design
    </div>
    <!-- Mobile UI start -->
    <div class="tw-block md:tw-hidden tw-my-10">
      <div
        class="tw-animate-gridAnimationReverse"
        :class="[pausedId == 0 ? gridAnimation : 'animation-paused']"
      >
        <div
          class="tw-animate-gridAnimation"
          :class="[pausedId == 0 ? 'animation-running' : 'animation-paused']"
        >
          <div class="tw-flex tw-ml-[160px]" v-if="mobileGrid.length > 0">
            <div
              v-for="(post, index) in mobileGrid"
              :key="index"
              class="tw-flex tw-flex-col tw-justify-center tw-flex-[0_0_320px] sm:tw-flex-[0_0_480px] lg:tw-flex-[0_0_528px] tw-ml-[32px] tw-rounded-2xl tw-text-[1rem] tw-leading-[1.125rem] md:tw-text-[1.0625rem] md:tw-leading-[1.5rem] lg:tw-text-[1.1875rem] lg:tw-leading-[1.875rem] tw-text-left"
              :class="
                pausedId == index ? 'tw-scale-[0.97] tw-cursor-pointer' : ''
              "
              @mouseover="pausedId = index"
              @mouseleave="pausedId = 0"
              @touchstart.passive="pausedId = index"
              @touchend="pausedId = 0"
              :ref="'card-1-' + index"
              @click="openBlog(post.link)"
            >
              <div class="image-container sm:tw-px-[20px]">
                <img
                  @click.native="mixpanel.track('tap_explore_design')"
                  :src="post.images[0]"
                  :srcset="`${post.images[0]} 400w, ${post.images[1]} 800w`"
                  alt="explore_uiux_design"
                  class="tw-object-cover tw-h-full tw-w-full tw-rounded-lg"
                  :class="post.classname"
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

    <div class="tw-hidden md:tw-block tw-mt-[3rem] xl:tw-mt-16">
      <div class="xll:tw-container swiper-content tw--mt-6">
        <swiper
          :slidesPerView="3"
          :centeredSlides="true"
          :autoplay="{
            delay: 2000,
            disableOnInteraction: false,
          }"
          :speed="3000"
          :spaceBetween="10"
          :freeMode="true"
          :navigation="true"
          :loop="true"
          :loopedSlides="50"
          class="swiper-container"
        >
          <swiper-slide v-for="(post, index) in data.slice(0, 5)" :key="index">
            <div
              class="tw-cursor-pointer image-content"
              @click="openBlog(post.link)"
            >
              <img
                :src="post.images[0]"
                :srcset="`${post.images[0]} 400w,${post.images[1]} 800w`"
                class="tw-w-full tw-h-full lg:tw-w-fit lg:tw-h-fit tw-object-cover tw-drop-shadow-md tw-rounded-lg"
                alt="DesignImage"
              />
            </div>
          </swiper-slide>
        </swiper>
      </div>
      <div class="xll:tw-container swiper-content">
        <swiper
          dir="rtl"
          :slidesPerView="4"
          :freeMode="true"
          :centeredSlides="true"
          :speed="3000"
          :autoplay="{
            delay: 2000,
            disableOnInteraction: false,
          }"
          :spaceBetween="20"
          :navigation="true"
          :loop="true"
          :loopedSlides="50"
          class="swiper-container"
        >
          <swiper-slide v-for="(post, index) in data.slice(5, 10)" :key="index">
            <div
              class="tw-cursor-pointer image2-content"
              @click="openBlog(post.link)"
            >
              <img
                :src="post.images[0]"
                :srcset="`${post.images[0]} 400w,${post.images[1]} 800w`"
                class="tw-w-full tw-h-full lg:tw-w-fit lg:tw-h-fit tw-object-cover tw-drop-shadow-md tw-rounded-lg"
                alt="DesignImage"
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
import chatapp from "@/assets/images/contributions/explore/1-400w.webp";
import chatapp800w from "@/assets/images/contributions/explore/1-800w.webp";

import elearningapp from "@/assets/images/contributions/explore/2-400w.webp";
import elearningapp800w from "@/assets/images/contributions/explore/2-800w.webp";

import fitnessapp from "@/assets/images/contributions/explore/3-400w.webp";
import fitnessapp800w from "@/assets/images/contributions/explore/3-800w.webp";

import foodorderingapp from "@/assets/images/contributions/explore/4-400w.webp";
import foodorderingapp800w from "@/assets/images/contributions/explore/4-800w.webp";

import irctcapp from "@/assets/images/contributions/explore/5-400w.webp";
import irctcapp800w from "@/assets/images/contributions/explore/5-800w.webp";

import meditationapp from "@/assets/images/contributions/explore/6-400w.webp";
import meditationapp800w from "@/assets/images/contributions/explore/6-800w.webp";

import moviestreamingapp from "@/assets/images/contributions/explore/7-400w.webp";
import moviestreamingapp800w from "@/assets/images/contributions/explore/7-800w.webp";

import musicapp from "@/assets/images/contributions/explore/8-400w.webp";
import musicapp800w from "@/assets/images/contributions/explore/8-800w.webp";

import shoppingapp from "@/assets/images/contributions/explore/9-400w.webp";
import shoppingapp800w from "@/assets/images/contributions/explore/9-800w.webp";

import socialnetworkingapp from "@/assets/images/contributions/explore/10-400w.webp";
import socialnetworkingapp800w from "@/assets/images/contributions/explore/10-800w.webp";

import AspectRatio from "@/components/utils/AspectRatio.vue";
import SwiperCore, { FreeMode, Pagination, Autoplay } from "swiper";
import { Swiper, SwiperSlide } from "swiper/vue";
SwiperCore.use([FreeMode, Pagination, Autoplay]);

export default {
  data() {
    return {
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
      data: [
        {
          images: [chatapp, chatapp800w],
          link: "https://dribbble.com/shots/20554178-Chat-app",
          classname: "tw-rotate-6",
        },
        {
          images: [elearningapp, elearningapp800w],
          link: "https://dribbble.com/shots/20494364-E-learning-app",
          classname: "tw--rotate-6",
        },
        {
          images: [fitnessapp, fitnessapp800w],
          link: "https://dribbble.com/shots/20682872-Fitness-App",
          classname: "tw-rotate-6",
        },
        {
          images: [foodorderingapp, foodorderingapp800w],
          link: "https://dribbble.com/shots/20513671-Food-Delivery-App",
          classname: "tw--rotate-6",
        },
        {
          images: [irctcapp, irctcapp800w],
          link: "https://dribbble.com/shots/20455362-Journey-Booking-App",
          classname: "tw-rotate-6",
        },
        {
          images: [meditationapp, meditationapp800w],
          link: "https://dribbble.com/shots/20749246-Meditation-app",
          classname: "tw--rotate-6",
        },
        {
          images: [moviestreamingapp, moviestreamingapp800w],
          link: "https://dribbble.com/shots/20597815-Movie-Streaming-app",
          classname: "tw-rotate-6",
        },
        {
          images: [musicapp, musicapp800w],
          link: "https://dribbble.com/shots/20781387-Music-Streaming-app",
          classname: "tw--rotate-6",
        },
        {
          images: [shoppingapp, shoppingapp800w],
          link: "https://dribbble.com/shots/20619022-Shopping-App",
          classname: "tw-rotate-6",
        },
        {
          images: [socialnetworkingapp, socialnetworkingapp800w],
          link: "https://dribbble.com/shots/20577294-Social-Networking-App",
          classname: "tw--rotate-6",
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
  },
  mounted() {
    this.appendKeyframes();
    window.addEventListener("scroll", this.handleScroll);
  },
  unmounted() {
    window.removeEventListener("scroll", this.handleScroll);
  },
  methods: {
    openBlog(link) {
      window.open(link, "_blank");
    },
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

    setGrid() {
      var length = this.data.length;
      this.grid1 = new Array(3).fill(this.data.slice(0, length / 2)).flat();
      this.grid2 = new Array(3)
        .fill(this.data.slice(length / 2, length))
        .flat();

      this.mobileGrid = new Array(3).fill(this.data).flat();
    },
    appendKeyframes() {
      var offsetWidth =
        window.innerWidth < 768 ? this.$refs["card-1-0"].offsetWidth : "";

      var length =
        window.innerWidth < 768
          ? this.data.length * 2 // for mobile
          : this.data.length;

      let width = ((offsetWidth + 32) * length) / 2;

      const head = document.getElementsByTagName("head")[0];
      const styles = document.getElementsByTagName("style");
      var found = false;

      // find style titled with explore, if found replace width according to data
      for (var i = 0; i < styles.length; i++) {
        if (styles[i].title == "explore") {
          found = true;
          styles[i].innerText = styles[i].innerText.replaceAll(
            this.lastWidth,
            width
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
@import "swiper/css/free-mode";

.swiper-slide-active .image-content {
  @apply tw-w-full tw-h-full tw-animate-verticalZoom tw-mb-[10%] 2xll:tw-mb-[11%] tw-scale-100;
}
.swiper-slide:not(.swiper-slide-active) .image-content {
  @apply tw-w-fit tw-h-fit tw-scale-[0.8];
}
.swiper-slide:not(.swiper-slide-active) {
  @apply tw-w-fit tw-h-fit;
}

.swiper-slide-prev .image-content {
  @apply tw-animate-verticalZoomOut tw-mb-[12%] 2xll:tw-mb-[11%];
}
.swiper-slide-next .image-content {
  @apply tw--mb-[2%] 2xll:tw--mb-[2%];
}
</style>
