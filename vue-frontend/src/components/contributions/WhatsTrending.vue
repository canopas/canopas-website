<template>
  <div class="tw-mt-5 xll:tw-container md:tw-mt-[14.063rem]">
    <div
      class="tw-container tw-text-center tw-font-inter-bold tw-text-[1.875rem] tw-leading-[2.25rem] lg:tw-text-[3.4375rem] lg:tw-leading-[5.1875rem]"
    >
      Discover What's Trending
    </div>
    <!-- Mobile UI start -->
    <div class="tw-my-10 tw-block tw-overflow-hidden md:tw-hidden">
      <div
        class="tw-animate-gridAnimationReverse"
        :class="[pausedId == 0 ? gridAnimation : 'animation-paused']"
      >
        <div
          class="tw-animate-gridAnimation"
          :class="[pausedId == 0 ? 'animation-running' : 'animation-paused']"
        >
          <div
            class="tw-ml-[160px] tw-flex tw-pb-2.5"
            v-if="mobileGrid.length > 0"
          >
            <div
              v-for="(post, index) in mobileGrid"
              :key="index"
              class="tw-ml-[32px] tw-flex tw-flex-[0_0_320px] tw-flex-col tw-justify-center tw-rounded-2xl tw-text-left tw-text-[1rem] tw-leading-[1.125rem] sm:tw-flex-[0_0_480px] md:tw-text-[1.0625rem] md:tw-leading-[1.5rem] lg:tw-flex-[0_0_528px] lg:tw-text-[1.1875rem] lg:tw-leading-[1.875rem]"
              :class="
                pausedId == post.id ? 'tw-scale-[0.97] tw-cursor-pointer' : ''
              "
              @mouseover="pausedId = post.id"
              @mouseleave="pausedId = 0"
              @touchstart.passive="pausedId = post.id"
              @touchend="pausedId = 0"
              :ref="'card-1-' + index"
              @click="openBlog(post.link)"
            >
              <div
                class="tw-overflow-hidden tw-rounded-3xl tw-border-1 tw-border-[#000000]/[0.6]"
              >
                <div class="">
                  <img
                    @click.native="mixpanel.track('tap_whats_trending_blog')"
                    :src="post.images[0]"
                    :srcset="`${post.images[0]} 400w, ${post.images[1]} 600w`"
                    alt="trending"
                    class="tw-h-[10%] tw-object-cover"
                    loading="lazy"
                  />
                </div>
                <div class="tw-flex tw-flex-col tw-gap-1 tw-p-[1rem]">
                  <div
                    class="tw-min-h-[5.625rem] tw-font-inter-semibold tw-text-[1rem] tw-leading-[1.875rem]"
                  >
                    {{ post.title }}
                  </div>
                  <div class="tw-w-fit tw-self-end">
                    <div
                      class="tw-rounded-full tw-border-[1px] tw-border-solid tw-border-transparent tw-from-[#ff835b] tw-to-[#f2709c] tw-px-[1.5rem] tw-py-[0.7rem] tw-text-center tw-font-inter-semibold tw-text-[1rem] tw-leading-[1.25rem] tw-tracking-normal tw-text-white tw-bg-gradient-[270.11deg]"
                    >
                      {{ post.views }} Views
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <!-- Mobile UI end -->
    <!-- Desktop UI start -->
    <div class="tw-mt-10 tw-hidden tw-overflow-hidden md:tw-block lg:tw-mt-20">
      <div
        class="tw-animate-gridAnimationReverse"
        :class="[pausedId == 0 ? gridAnimation : 'animation-paused']"
      >
        <div
          class="tw-animate-gridAnimation"
          :class="[pausedId == 0 ? 'animation-running' : 'animation-paused']"
        >
          <div
            class="tw-mb-[64px] tw-ml-[160px] tw-mt-[32px] tw-flex"
            v-if="grid1.length > 0"
          >
            <div
              v-for="(post, index) in grid1"
              :key="index"
              class="tw-relative tw-ml-[32px] tw-flex tw-flex-[0_0_320px] tw-flex-col tw-justify-center tw-rounded-2xl tw-text-left tw-text-[1rem] tw-leading-[1.125rem] sm:tw-flex-[0_0_480px] md:tw-text-[1.0625rem] md:tw-leading-[1.5rem] lg:tw-flex-[0_0_528px] lg:tw-text-[1.1875rem] lg:tw-leading-[1.875rem]"
              :class="
                pausedId == post.id ? 'tw-scale-[0.97] tw-cursor-pointer' : ''
              "
              @mouseover="pausedId = post.id"
              @mouseleave="pausedId = 0"
              @touchstart.passive="pausedId = post.id"
              @touchend="pausedId = 0"
              :ref="'card-2-' + index"
              @click="openBlog(post.link)"
            >
              <div
                :ref="'view-2-' + index"
                :style="{
                  top: `-${viewTop}px`,
                }"
                class="tw-absolute tw-inset-x-0 tw-z-10 tw-m-auto tw-w-fit tw-rounded-full tw-border-[1px] tw-border-solid tw-border-transparent tw-from-[#ff835b] tw-to-[#f2709c] tw-px-[1.5rem] tw-py-[0.3rem] tw-text-center tw-font-inter-semibold tw-text-[1rem] tw-leading-[1.25rem] tw-tracking-normal tw-text-white tw-bg-gradient-[270.11deg] xl:tw-text-[1.625rem] xl:tw-leading-[2.4375rem]"
              >
                {{ post.views }} Views
              </div>
              <aspect-ratio
                height="50%"
                class="tw-border-[1px] tw-border-solid tw-border-transparent"
              >
                <img
                  @click.native="mixpanel.track('tap_whats_trending_blog')"
                  :src="post.images[0]"
                  :srcset="`${post.images[0]} 400w, ${post.images[1]} 600w`"
                  alt="trending"
                  class="tw-rounded-md tw-object-contain"
                  loading="lazy"
                />
              </aspect-ratio>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="tw-mb-10 tw-hidden tw-overflow-hidden md:tw-block lg:tw-mb-20">
      <div
        class="tw-animate-gridAnimation"
        :class="[pauseId == 0 ? gridAnimation : 'animation-paused']"
      >
        <div
          class="tw-animate-gridAnimationReverse2"
          :class="[pauseId == 0 ? gridAnimationReverse : 'animation-paused']"
        >
          <div class="tw-mb-[64px] tw-flex" v-if="grid2.length > 0">
            <div
              v-for="(post, index) in grid2"
              :key="index"
              class="tw-relative tw-ml-[32px] tw-mt-[32px] tw-flex tw-flex-[0_0_320px] tw-flex-col tw-justify-center tw-rounded-2xl tw-text-left tw-text-[1rem] tw-leading-[1.125rem] sm:tw-flex-[0_0_480px] md:tw-text-[1.0625rem] md:tw-leading-[1.5rem] lg:tw-flex-[0_0_528px] lg:tw-text-[1.1875rem] lg:tw-leading-[1.875rem]"
              :class="
                pauseId == post.id ? 'tw-scale-[0.97] tw-cursor-pointer' : ''
              "
              @mouseover="pauseId = post.id"
              @mouseleave="pauseId = 0"
              @touchstart.passive="pauseId = post.id"
              @touchend="pauseId = 0"
              :ref="'card-3-' + index"
              @click="openBlog(post.link)"
            >
              <div
                :style="{
                  top: `-${viewTop}px`,
                }"
                class="tw-absolute tw-inset-x-0 tw-z-10 tw-m-auto tw-w-fit tw-rounded-full tw-border-[1px] tw-border-solid tw-border-transparent tw-from-[#ff835b] tw-to-[#f2709c] tw-px-[1.5rem] tw-py-[0.3rem] tw-text-center tw-font-inter-semibold tw-text-[1rem] tw-leading-[1.25rem] tw-tracking-normal tw-text-white tw-bg-gradient-[270.11deg] xl:tw-text-[1.625rem] xl:tw-leading-[2.4375rem]"
              >
                {{ post.views }} Views
              </div>
              <aspect-ratio
                height="50%"
                class="tw-border-[1px] tw-border-solid tw-border-transparent"
              >
                <img
                  @click.native="mixpanel.track('tap_whats_trending_blog')"
                  :src="post.images[0]"
                  :srcset="`${post.images[0]} 400w, ${post.images[1]} 600w`"
                  alt="trending"
                  class="tw-rounded-md tw-object-contain"
                  loading="lazy"
                />
              </aspect-ratio>
            </div>
          </div>
        </div>
      </div>
    </div>
    <!-- Desktop UI end -->
  </div>
</template>

<script>
import retrofit from "@/assets/images/contributions/trending/1-400w.webp";
import retrofit600w from "@/assets/images/contributions/trending/1-600w.webp";

import goGuide from "@/assets/images/contributions/trending/2-400w.webp";
import goGuide600w from "@/assets/images/contributions/trending/2-600w.webp";

import keyboardHandling from "@/assets/images/contributions/trending/3-400w.webp";
import keyboardHandling600w from "@/assets/images/contributions/trending/3-600w.webp";

import kotlinFlow from "@/assets/images/contributions/trending/4-400w.webp";
import kotlinFlow600w from "@/assets/images/contributions/trending/4-600w.webp";

import gorm from "@/assets/images/contributions/trending/5-400w.webp";
import gorm600w from "@/assets/images/contributions/trending/5-600w.webp";

import androidGuide from "@/assets/images/contributions/trending/6-400w.webp";
import androidGuide600w from "@/assets/images/contributions/trending/6-600w.webp";

import globalVars from "@/assets/images/contributions/trending/7-400w.webp";
import globalVars600w from "@/assets/images/contributions/trending/7-600w.webp";

import dragAndDrop from "@/assets/images/contributions/trending/8-400w.webp";
import dragAndDrop600w from "@/assets/images/contributions/trending/8-600w.webp";

import jacoco from "@/assets/images/contributions/trending/9-400w.webp";
import jacoco600w from "@/assets/images/contributions/trending/9-600w.webp";

import lifecycle from "@/assets/images/contributions/trending/10-400w.webp";
import lifecycle600w from "@/assets/images/contributions/trending/10-600w.webp";

import AspectRatio from "@/components/utils/AspectRatio.vue";

export default {
  data() {
    return {
      gridAnimation: "animation-paused",
      gridAnimationReverse: "animation-running",
      pausedId: 0,
      pauseId: 0,
      lastWidth: null,
      lastScrollY: 0,
      viewTop: 0,
      mobileGrid: [],
      grid1: [],
      grid2: [],
      data: [
        {
          id: 1,
          title:
            "Retrofit — Effective error handling with Kotlin Coroutine and Result API",
          images: [retrofit, retrofit600w],
          views: "26k",
          link: "https://blog.canopas.com/retrofit-effective-error-handling-with-kotlin-coroutine-and-result-api-405217e9a73d",
        },
        {
          id: 2,
          title: "Golang — 1 Minute guide to Useful Tips and Libraries in 2022",
          images: [goGuide, goGuide600w],
          views: "18k",
          link: "https://blog.canopas.com/1-min-guide-to-golang-development-best-practices-in-2022-b50d846fd6c",
        },
        {
          id: 3,
          title: "Keyboard Handling In Jetpack Compose — All You Need To Know",
          images: [keyboardHandling, keyboardHandling600w],
          views: "17k",
          link: "https://blog.canopas.com/keyboard-handling-in-jetpack-compose-all-you-need-to-know-3e6fddd30d9a",
        },
        {
          id: 4,
          title: "Android — 9 Useful Kotlin Flow Operators You Need to Know",
          images: [kotlinFlow, kotlinFlow600w],
          views: "14.5k",
          link: "https://blog.canopas.com/android-9-useful-kotlin-flow-operators-you-need-to-know-b9daef4b630f",
        },
        {
          id: 5,
          title: "Golang: gorm with MySQL and gin",
          images: [gorm, gorm600w],
          views: "14.2k",
          link: "https://blog.canopas.com/golang-gorm-with-mysql-and-gin-ab876f406244",
        },
        {
          id: 6,
          title:
            "Android — 1 Minute guide to Useful Tips and Libraries in 2022",
          images: [androidGuide, androidGuide600w],
          views: "13.5k",
          link: "https://blog.canopas.com/android-development-best-practices-2022-203682a440f5",
        },
        {
          id: 7,
          title: "Approach To Avoid Accessing Variables Globally in Golang",
          images: [globalVars, globalVars600w],
          views: "12.4k",
          link: "https://blog.canopas.com/approach-to-avoid-accessing-variables-globally-in-golang-2019b234762",
        },
        {
          id: 8,
          title: "Android — How to Drag And Drop Views in Jetpack Compose",
          images: [dragAndDrop, dragAndDrop600w],
          views: "12.2k",
          link: "https://blog.canopas.com/android-drag-and-drop-ui-element-in-jetpack-compose-14922073b3f1",
        },
        {
          id: 9,
          title: "Android code coverage using JaCoCo",
          images: [jacoco, jacoco600w],
          views: "11.8k",
          link: "https://blog.canopas.com/android-code-coverage-using-jacoco-6639a1fc4293",
        },
        {
          id: 10,
          title: "Vue 3 lifecycle hooks with real-time example",
          images: [lifecycle, lifecycle600w],
          views: "11k",
          link: "https://blog.canopas.com/vue-3-lifecycle-hooks-with-real-time-example-1b772b89e085",
        },
      ],
    };
  },
  components: {
    AspectRatio,
  },
  inject: ["mixpanel"],
  created() {
    this.setGrid();
  },
  mounted() {
    this.setViews();
    this.appendKeyframes();
    window.addEventListener("scroll", this.handleScroll);
  },
  unmounted() {
    window.removeEventListener("scroll", this.handleScroll);
  },
  methods: {
    openBlog(link) {
      window.open(link, "_blank");
      this.mixpanel.track("tap_contribution_trending_section");
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
    setViews() {
      this.viewTop = this.$refs["view-2-0"][0].offsetHeight / 2;
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
        window.innerWidth < 768
          ? this.$refs["card-1-0"][0].offsetWidth
          : this.$refs["card-2-0"][0].offsetWidth;

      var length =
        window.innerWidth < 768
          ? this.data.length * 2 // for mobile
          : this.data.length;

      let width = ((offsetWidth + 32) * length) / 2;

      const head = document.getElementsByTagName("head")[0];
      const styles = document.getElementsByTagName("style");
      var found = false;

      // find style titled with trending, if found replace width according to data
      for (var i = 0; i < styles.length; i++) {
        if (styles[i].title == "trending") {
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
          `@-webkit-keyframes scroll { 0% {transform: translateX(0px)} 100% {transform:translateX(-` +
          width +
          `px)}}\n` +
          `@-webkit-keyframes scroll-reverse { 0% {transform:translateX(-` +
          width +
          `px)} 100% {transform: translateX(0px)}}\n` +
          `@keyframes scroll-reverse { 0% {transform:translateX(-` +
          width +
          `px)} 100% {transform: translateX(0px)}}`;

        let style = document.createElement("style");
        style.title = "trending";
        style.innerHTML = keyframes;

        head.appendChild(style);
      }

      this.lastWidth = width;
    },
  },
};
</script>
