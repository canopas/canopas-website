<template>
  <section class="lg:mt-60">
    <div class="container">
      <p class="text-center mobile-header-2 lg:desk-header-2 text-black-87">
        GitHub Contributions
      </p>
      <p
        class="mt-4 text-center sub-h1-regular lg:mobile-header-2-regular text-black-60"
      >
        Explore our repositories and join us on this thrilling journey of
        creating code that fuels global innovation.
      </p>
      <div
        class="m-auto mt-4 flex w-full justify-around gap-4 rounded-xl bg-white py-[0.313rem] drop-shadow-[0_2px_10px_rgba(61,61,61,0.1)] sm:w-[95%] lg:mt-[3.75rem] lg:w-3/4"
      >
        <button
          v-for="(navbar, index) in navbars"
          :key="index"
          class="portfolio-nav m-0 w-full rounded-lg py-2 text-center sub-h3-semibold lg:sub-h1-semibold active:scale-[0.98] md:py-[0.813rem] lg:hover:from-orange-300 lg:hover:to-pink-300 lg:hover:text-white lg:hover:bg-gradient-[270.11deg]"
          @click="(slideTo(index), changeContributions(index))"
          :class="
            activeIndex == index
              ? 'from-orange-300 to-pink-300 text-white bg-gradient-[270.11deg] '
              : 'border border-solid border-transparent text-black-60'
          "
        >
          {{ navbar }}
        </button>
      </div>
    </div>
    <!-- Mobile UI start -->
    <div class="swiper-content !mt-6 lg:hidden">
      <swiper
        :slidesPerView="1"
        :spaceBetween="10"
        :centeredSlides="true"
        @swiper="setSwiperRef"
        @slideChange="onSlideChange()"
        class="swiper-container"
        :breakpoints="{
          '768': {
            slidesPerView: 1.5,
          },
        }"
      >
        <swiper-slide
          v-for="(contribution, index) in allContributions"
          :key="index"
          @click="openBlog(contribution.link, contribution.event)"
          class="cursor-pointer"
        >
          <img
            v-if="contribution.image"
            :src="contribution.image[0]"
            :srcset="`${contribution.image[0]} 400w, ${contribution.image[1]} 800w`"
            class="relative h-full w-full object-cover"
            alt="contribution-image"
            loading="lazy"
          />
          <div
            class="z-[1] mx-auto mb-10 mt-[-3.125rem] sm:mt-[-3.75rem] flex max-w-[95%] flex-col rounded-[10px] bg-white py-4 drop-shadow-xl sm:max-w-[85%] md:max-w-[90%]"
          >
            <div class="flex flex-row items-center justify-start gap-8 px-4">
              <div
                class="flex h-[1.875rem] w-[4.75rem] items-center justify-center rounded-lg px-2.5 bg-primary-1 text-white sub-h3-semibold"
              >
                <Icon
                  class="fa h-[1.125rem] w-[1.125rem] pr-[0.313rem] text-white"
                  name="fa6-solid:star"
                />{{ contribution.stars }}
              </div>
              <div class="flex flex-row items-center gap-2">
                <span
                  class="h-4 w-4 rounded-full"
                  :class="contribution.color"
                ></span
                ><span class="sub-h3-medium text-black-87">{{
                  contribution.language
                }}</span>
              </div>
              <div class="flex flex-row items-center gap-2">
                <Icon
                  class="fa h-4 w-4 sub-h3-medium text-black-87"
                  name="fa6-solid:code-fork"
                />{{ contribution.forks }}
              </div>
            </div>
            <p class="mt-4 px-4 sub-h3-regular text-black-60">
              <span class="sub-h1-semibold text-black-87">{{
                contribution.title
              }}</span
              ><br />{{ contribution.description }}
            </p>
            <div class="mt-5 flex flex-row items-center px-4">
              <p class="mr-2 sub-h3-semibold">Contributors:</p>
              <div
                class="ml-2 flex flex-row items-center"
                v-for="(contributor, index) in contribution.contributors"
                :key="index"
              >
                <img
                  :src="contributor"
                  class="h-9 w-9 object-cover"
                  alt="contribution-image"
                  loading="lazy"
                />
              </div>
            </div>
          </div>
        </swiper-slide>
      </swiper>
    </div>
    <!-- Mobile UI end -->
    <!-- Desktop UI start -->
    <div
      class="container mx-auto mt-8 hidden 2xl:gap-12 flex-wrap justify-between lg:flex lg:my-10"
    >
      <div
        v-for="(contribution, index) in contributions"
        :key="index"
        class="flip-card-inner mx-auto relative h-[23rem] w-[48%] cursor-pointer md:-mt-12 lg:mt-4 lg:h-[27.5rem] xl:mt-8 xl:h-[31.25rem] 2xl:h-[34.375rem]"
        @click="openBlog(contribution.link, contribution.event)"
      >
        <div class="flip-card-front absolute w-full h-full overflow-hidden">
          <img
            :src="contribution.deskImage[1]"
            :srcset="`${contribution.deskImage[0]} 400w, ${contribution.deskImage[1]} 800w`"
            class="absolute w-full rounded-[20px] object-cover lg:h-full"
            alt="contribution-image"
            loading="lazy"
          />
        </div>
        <div
          class="flip-card-back absolute h-full w-full rounded-[20px] from-orange-300 to-pink-300 bg-gradient-[180deg]"
        >
          <div
            class="flex flex-row items-center justify-between pl-6 pr-4 pt-6 xl:px-12 xl:pt-10"
          >
            <div class="flex flex-row items-center gap-2 text-white">
              <span>
                <Icon class="fa h-6 w-6" name="fa6-solid:code-fork" /></span
              ><span class="mobile-header-2-semibold">{{
                contribution.forks
              }}</span>
            </div>
            <div class="flex flex-row items-center gap-2">
              <span class="mobile-header-2-semibold text-white">{{
                contribution.language
              }}</span>
            </div>
            <div
              class="flex flex-row items-center rounded-xl bg-white px-4 py-2 xl:px-6"
            >
              <Icon
                class="fab footer-icon h-[1.625rem] w-[1.625rem] pr-[0.313rem]"
                name="fa6-solid:star"
              />
              <span
                class="v2-canopas-gradient-text mt-0.5 mobile-header-2-semibold"
                >{{ contribution.stars }}</span
              >
            </div>
          </div>
          <p class="ml-6 mr-2 xl:ml-12 xl:mr-4 text-white mt-12 desk-header-3">
            {{ contribution.title }}
          </p>
          <p
            class="ml-6 mr-2 xl:ml-12 xl:mr-12 text-white-core-80 mobile-header-2-regular"
          >
            {{ contribution.description }}
          </p>
          <div class="mt-4 flex flex-row items-center px-6 xl:px-12 lg:mt-8">
            <p class="mr-2 mobile-header-2-semibold text-white">
              Contributors:
            </p>
            <div
              class="ml-2 flex flex-row items-center"
              v-for="(contributor, index) in contribution.contributors"
              :key="index"
            >
              <img
                :src="contributor"
                class="h-9 w-9 rounded-full object-cover"
                alt="contribution-image"
                loading="lazy"
              />
            </div>
          </div>
        </div>
      </div>
    </div>
    <!-- Desktop UI end -->
  </section>
</template>

<script>
import { Pagination } from "swiper/modules";
import { Swiper, SwiperSlide } from "swiper/vue";
import { openBlog } from "@/utils.js";
import config from "@/config.js";

import android1_400w from "@/assets/images/contributions/github/mobile/android-1-400w.webp";
import android1_800w from "@/assets/images/contributions/github/mobile/android-1-800w.webp";
import android2_400w from "@/assets/images/contributions/github/mobile/android-2-400w.webp";
import android2_800w from "@/assets/images/contributions/github/mobile/android-2-800w.webp";
import android3_400w from "@/assets/images/contributions/github/mobile/android-3-400w.webp";
import android3_800w from "@/assets/images/contributions/github/mobile/android-3-800w.webp";

import ios1_400w from "@/assets/images/contributions/github/mobile/ios-1-400w.webp";
import ios1_800w from "@/assets/images/contributions/github/mobile/ios-1-800w.webp";
import ios2_400w from "@/assets/images/contributions/github/mobile/ios-2-400w.webp";
import ios2_800w from "@/assets/images/contributions/github/mobile/ios-2-800w.webp";
import ios3_400w from "@/assets/images/contributions/github/mobile/ios-3-400w.webp";
import ios3_800w from "@/assets/images/contributions/github/mobile/ios-3-800w.webp";

import web1_400w from "@/assets/images/contributions/github/mobile/web-1-400w.webp";
import web1_800w from "@/assets/images/contributions/github/mobile/web-1-800w.webp";
import web2_400w from "@/assets/images/contributions/github/mobile/web-2-400w.webp";
import web2_800w from "@/assets/images/contributions/github/mobile/web-2-800w.webp";
import web3_400w from "@/assets/images/contributions/github/mobile/web-3-400w.webp";
import web3_800w from "@/assets/images/contributions/github/mobile/web-3-800w.webp";
import web4_400w from "@/assets/images/contributions/github/mobile/web-4-400w.webp";
import web4_800w from "@/assets/images/contributions/github/mobile/web-4-800w.webp";

import desktop_android1_400w from "@/assets/images/contributions/github/desktop/android-1-400w.webp";
import desktop_android1_800w from "@/assets/images/contributions/github/desktop/android-1-800w.webp";
import desktop_android2_400w from "@/assets/images/contributions/github/desktop/android-2-400w.webp";
import desktop_android2_800w from "@/assets/images/contributions/github/desktop/android-2-800w.webp";
import desktop_android3_400w from "@/assets/images/contributions/github/desktop/android-3-400w.webp";
import desktop_android3_800w from "@/assets/images/contributions/github/desktop/android-3-800w.webp";

import desktop_ios1_400w from "@/assets/images/contributions/github/desktop/ios-1-400w.webp";
import desktop_ios1_800w from "@/assets/images/contributions/github/desktop/ios-1-800w.webp";
import desktop_ios2_400w from "@/assets/images/contributions/github/desktop/ios-2-400w.webp";
import desktop_ios2_800w from "@/assets/images/contributions/github/desktop/ios-2-800w.webp";
import desktop_ios3_400w from "@/assets/images/contributions/github/desktop/ios-3-400w.webp";
import desktop_ios3_800w from "@/assets/images/contributions/github/desktop/ios-3-800w.webp";

import desktop_web1_400w from "@/assets/images/contributions/github/desktop/web-1-400w.webp";
import desktop_web1_800w from "@/assets/images/contributions/github/desktop/web-1-800w.webp";
import desktop_web2_400w from "@/assets/images/contributions/github/desktop/web-2-400w.webp";
import desktop_web2_800w from "@/assets/images/contributions/github/desktop/web-2-800w.webp";
import desktop_web3_400w from "@/assets/images/contributions/github/desktop/web-3-400w.webp";
import desktop_web3_800w from "@/assets/images/contributions/github/desktop/web-3-800w.webp";
import desktop_web4_400w from "@/assets/images/contributions/github/desktop/web-4-400w.webp";
import desktop_web4_800w from "@/assets/images/contributions/github/desktop/web-4-800w.webp";

import contributor1 from "@/assets/images/contributions/github/contributors/contributor-1-100w.webp";
import contributor2 from "@/assets/images/contributions/github/contributors/contributor-2-100w.webp";
import contributor3 from "@/assets/images/contributions/github/contributors/contributor-3-100w.webp";
import contributor4 from "@/assets/images/contributions/github/contributors/contributor-4-100w.webp";
import contributor6 from "@/assets/images/contributions/github/contributors/contributor-6-100w.webp";
import contributor8 from "@/assets/images/contributions/github/contributors/contributor-8-100w.webp";
import contributor9 from "@/assets/images/contributions/github/contributors/contributor-9-100w.webp";
import contributor10 from "@/assets/images/contributions/github/contributors/contributor-10-100w.webp";
import contributor11 from "@/assets/images/contributions/github/contributors/contributor-11-100w.webp";
import contributor12 from "@/assets/images/contributions/github/contributors/contributor-12-100w.webp";
import contributor14 from "@/assets/images/contributions/github/contributors/contributor-14-100w.webp";

const ANDROID = 0;
const IOS = 1;
const WEB = 2;

export default {
  data() {
    return {
      modules: [Pagination],
      openBlog,
      navbars: ["Android", "iOS", "Web"],
      contributions: [],
      activeIndex: 0,
      swiperRef: null,
      allContributions: [
        {
          id: 1,
          key: ANDROID,
          image: [android1_400w, android1_800w],
          deskImage: [desktop_android1_400w, desktop_android1_800w],
          link: "https://github.com/canopas/compose-intro-showcase",
          language: "Kotlin",
          color: "bg-[#A97BFF]",
          forks: "23",
          title: "Compose Intro Showcase",
          description:
            "Intro showcase view in Jetpack compose. An implementation of Intro Showcase from ...",
          stars: config.INTRO_SHOWCASE_STARS,
          contributors: [contributor1, contributor2, contributor3],
          event: "tap_contribution_intro_showcase",
        },
        {
          id: 2,
          key: ANDROID,
          image: [android2_400w, android2_800w],
          deskImage: [desktop_android2_400w, desktop_android2_800w],
          link: "https://github.com/canopas/compose-animations-examples",
          language: "Kotlin",
          color: "bg-[#A97BFF]",
          forks: "26",
          title: "Compose Animation Examples",
          description: "Cool animations implemented with Jetpack compose....",
          stars: config.JETPACK_COMPOSE_ANIMATION_STARS,
          contributors: [contributor1, contributor2, contributor4],
          event: "tap_contribution_jetpack_compose_animation",
        },
        {
          id: 3,
          key: ANDROID,
          image: [android3_400w, android3_800w],
          deskImage: [desktop_android3_400w, desktop_android3_800w],
          link: "https://github.com/canopas/compose-country-picker",
          language: "Kotlin",
          color: "bg-[#A97BFF]",
          forks: "02",
          title: "Compose CountryPicker",
          description:
            "Country code bottomsheet picker in Jetpack Compose with Search functionality...",
          stars: config.COMPOSE_COUNTRY_PICKER_STARS,
          contributors: [contributor1, contributor2, contributor4],
          event: "tap_contribution_jetcountry_picker",
        },
        {
          id: 4,
          key: IOS,
          image: [ios1_400w, ios1_800w],
          deskImage: [desktop_ios1_400w, desktop_ios1_800w],
          link: "https://github.com/canopas/UIPilot",
          language: "Swift",
          color: "bg-[#F05138]",
          forks: "21",
          title: "UIPilot ",
          description: "The missing typesafe SwiftUI navigation library.",
          stars: config.UIPILOT_STARS,
          contributors: [contributor1, contributor6, contributor2],
          event: "tap_contribution_jetpack_compose_animation",
        },
        {
          id: 5,
          key: IOS,
          image: [ios2_400w, ios2_800w],
          deskImage: [desktop_ios2_400w, desktop_ios2_800w],
          link: "https://github.com/canopas/Swiftui-animations-examples",
          language: "Swift",
          color: "bg-[#F05138]",
          forks: "02",
          title: "Swiftui Animations",
          description: "Cool animations implemented with SwiftUI",
          stars: config.SWIFT_UI_ANIMATION_STARS,
          contributors: [contributor1, contributor6],
          event: "tab_contribution_swiftui_animation",
        },
        {
          id: 6,
          key: IOS,
          image: [ios3_400w, ios3_800w],
          deskImage: [desktop_ios3_400w, desktop_ios3_800w],
          link: "https://github.com/canopas/iOS-developer-roadmap",
          language: "Swift",
          color: "bg-[#F05138]",
          forks: "13",
          title: "iOS Developer Roadmap 2022",
          description:
            "iOS Developer Roadmap 2022 is a learning path to understand iOS development...",
          stars: config.IOS_DEVELOPER_ROADMAP_STARS,
          contributors: [contributor1, contributor6],
          event: "tab_contribution_ios_roadmap",
        },
        {
          id: 7,
          key: WEB,
          image: [web1_400w, web1_800w],
          deskImage: [desktop_web1_400w, desktop_web1_800w],
          link: "https://github.com/canopas/tailwind-animations-examples",
          language: "JavaScript",
          color: "bg-[#F1E05A]",
          forks: "05",
          title: "Tailwind Animations Examples",
          description:
            "This repository contains different animations implemented using tailwind css...",
          stars: config.TAILWIND_ANIMATIONS_STARS,
          contributors: [contributor8, contributor9],
          event: "tab_contribution_tailwind_animation",
        },
        {
          id: 8,
          key: WEB,
          image: [web2_400w, web2_800w],
          deskImage: [desktop_web2_400w, desktop_web2_800w],
          link: "https://github.com/canopas/web-developer-roadmap",
          language: "",
          color: "",
          forks: "14",
          title: "Web Development Roadmap",
          description:
            "Web Developer Roadmap is a path to understand web development including...",
          stars: config.WEB_DEVELOPMENT_ROADMAP_STARS,
          contributors: [contributor8, contributor10],
          event: "tab_contribution_webdevelopment_roadmap",
        },
        {
          id: 9,
          key: WEB,
          image: [web3_400w, web3_800w],
          deskImage: [desktop_web4_400w, desktop_web4_800w],
          link: "https://github.com/canopas/canopas-blog",
          language: "JavaScript",
          color: "bg-[#F1E05A]",
          forks: "03",
          title: "Canopas Blog website",
          description:
            "Source code of canopas resources by following best practices...",
          stars: config.CANOPAS_BLOG_WEBSITE_STARS,
          contributors: [contributor8, contributor11, contributor12],
          event: "tab_contribution_canopas_blog",
        },
        {
          id: 10,
          key: WEB,
          image: [web4_400w, web4_800w],
          deskImage: [desktop_web3_400w, desktop_web3_800w],
          link: "https://github.com/canopas/canopas-website",
          language: "Vue",
          color: "bg-[#41B883]",
          forks: "03",
          title: "Canopas Website",
          description:
            "Source code of canopas website by following best practices...",
          stars: config.CANOPAS_WEBSITE_STARS,
          contributors: [contributor8, contributor14, contributor12],
          event: "tap_contribution_canopas_website",
        },
      ],
    };
  },
  components: {
    Swiper,
    SwiperSlide,
  },
  mounted() {
    this.contributions = this.allContributions.filter((obj) => {
      return obj.key === ANDROID;
    });
  },
  inject: ["mixpanel"],
  methods: {
    openContribution(contribution) {
      window.open(contribution.link, "_blank");
      this.$mixpanel.track(contribution.event);
    },
    onSlideChange() {
      this.activeIndex = this.allContributions[this.swiperRef.activeIndex].key;
    },
    setSwiperRef(swiper) {
      this.swiperRef = swiper;
    },
    slideTo(index) {
      this.activeIndex = index;
      this.swiperRef.slideTo(index * 3, 1000);
    },
    changeContributions(index) {
      this.activeIndex = index;
      this.contributions = this.allContributions.filter((obj) => {
        return obj.key === index;
      });
    },
  },
};
</script>
<style lang="postcss" scoped>
@import "swiper/css";
@import "swiper/css/pagination";

.swiper-wrapper {
  @apply !items-center;
}

.swiper {
  @apply z-0;
}

.flip-card-inner {
  transform-style: preserve-3d;
  perspective: 1000px;
}

.flip-card-front,
.flip-card-back {
  backface-visibility: hidden;
  transform-style: preserve-3d;
  transition: transform 0.7s ease 0.1s;
}
.flip-card-back {
  transform: rotateY(180deg);
  box-shadow: 0 27px 80px 0 #f2709c;
}
.flip-card-front {
  transform: rotateY(0deg);
}
.flip-card-inner:hover .flip-card-back {
  transform: rotateY(0deg);
}
.flip-card-inner:hover .flip-card-front {
  transform: rotateY(-180deg);
}

.flip-card-back p,
div,
img,
span {
  transform: translateZ(80px);
}
@media (min-width: 768px) and (max-width: 1199px) {
  .flip-card-back p,
  div,
  img,
  span {
    transform: translateZ(40px);
  }
}
</style>
