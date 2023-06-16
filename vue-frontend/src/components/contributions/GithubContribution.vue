<template>
  <section class="lg:tw-my-[100px]">
    <div class="tw-container">
      <p
        class="tw-text-center tw-text-[1.875rem] lg:tw-text-[3.438rem] tw-leading-[2.438rem] lg:tw-leading-[5.156rem] tw-font-inter-bold tw-text-black-core/[0.87]"
      >
        GitHub Contributions
      </p>
      <p
        class="tw-mt-4 lg:tw-mt-8 tw-text-center tw-text-[1rem] lg:tw-text-[1.5rem] tw-leading-[1.5rem] lg:tw-leading-[2.25rem] tw-font-inter-regular lg:tw-font-inter-medium tw-text-black-core/[0.87] lg:tw-text-black-core/[0.6]"
      >
        Explore our repositories and join us on this thrilling journey of
        creating code that fuels global innovation.
      </p>
      <div
        class="tw-flex tw-justify-around tw-gap-4 tw-m-auto tw-mt-8 lg:tw-mt-[3.75rem] tw-w-[100%] sm:tw-w-[95%] lg:tw-w-[75%] tw-rounded-[30px] tw-bg-white tw-py-[5px] tw-drop-shadow-[0_2px_10px_rgba(61,61,61,0.1)]"
      >
        <button
          v-for="(navbar, index) in navbars"
          :key="index"
          class="tw-flex-[16.66%] tw-m-0 tw-w-full tw-rounded-[25px] tw-px-[5px] tw-py-2 md:tw-py-[13px] tw-text-center tw-tracking-[1px] tw-text-[0.75rem] tw-leading-[0.938rem] md:tw-text-[1rem] md:tw-leading-[1.1875rem] lg:tw-text-[1.188rem] lg:tw-leading-[1.438rem] tw-font-inter-semibold active:tw-scale-[0.98] portfolio-nav lg:hover:tw-bg-gradient-[270.11deg] lg:hover:tw-from-[#ff835b] lg:hover:tw-to-[#f2709c] lg:hover:tw-text-white"
          @click="slideTo(index), changeContributions(index)"
          :class="
            activeIndex == index
              ? 'tw-bg-gradient-[270.11deg] tw-from-[#ff835b] tw-to-[#f2709c] tw-text-white '
              : 'tw-border-[1px] tw-border-solid tw-border-transparent tw-text-black-core/[0.6]'
          "
        >
          {{ navbar }}
        </button>
      </div>
    </div>
    <!-- Mobile UI start -->
    <div class="sm:tw-container swiper-content !tw-mt-8 lg:tw-hidden">
      <swiper
        :slidesPerView="1"
        :centeredSlides="true"
        :spaceBetween="20"
        @swiper="setSwiperRef"
        @slideChange="onSlideChange()"
        class="swiper-container"
      >
        <swiper-slide
          v-for="(contribution, index) in allContributions"
          :key="index"
          @click="openContribution(contribution)"
          class="tw-cursor-pointer"
        >
          <img
            v-if="contribution.image"
            :src="contribution.image[0]"
            :srcset="`${contribution.image[0]} 400w, ${contribution.image[1]} 800w`"
            class="tw-w-full tw-h-full tw-object-cover tw-relative"
            alt="contribution-image"
          />
          <div
            class="tw-flex tw-flex-col tw-mt-[-23px] tw-mb-10 tw-mx-auto tw-max-w-[95%] sm:tw-max-w-[85%] md:tw-max-w-[70%] tw-z-[1] tw-rounded-[10px] tw-bg-white tw-drop-shadow-2xl tw-py-4"
          >
            <div
              class="tw-flex tw-flex-row tw-items-center tw-justify-start tw-gap-8 tw-px-4"
            >
              <div
                class="tw-flex tw-items-center tw-justify-center tw-w-[4.75rem] tw-h-[1.875rem] tw-rounded-full tw-bg-gradient-[270.11deg] tw-from-[#ff835b] tw-to-[#f2709c] tw-px-2.5 tw-font-inter-semibold tw-text-[1rem] tw-text-white"
              >
                <span>
                  <font-awesome-icon
                    class="fa tw-w-[18px] tw-h-[18px] tw-pr-[5px] tw-text-white"
                    :icon="star"
                  />{{ contribution.stars }}</span
                >
              </div>
              <div class="tw-flex tw-flex-row tw-items-center tw-gap-2">
                <span
                  class="tw-w-4 tw-h-4 tw-rounded-full"
                  :class="contribution.color"
                ></span
                ><span
                  class="tw-text-[1rem] tw-leading-[1rem] tw-font-inter-medium tw-text-black-core/[0.87]"
                  >{{ contribution.language }}</span
                >
              </div>
              <div class="tw-flex tw-flex-row tw-items-center tw-gap-2">
                <span>
                  <font-awesome-icon
                    class="fa tw-w-4 tw-h-4"
                    :icon="fork" /></span
                ><span
                  class="tw-text-[1rem] tw-leading-[1rem] tw-font-inter-medium tw-text-black-core/[0.87]"
                  >{{ contribution.forks }}</span
                >
              </div>
            </div>
            <p
              class="tw-mt-2 tw-px-4 tw-text-[1rem] tw-leading-[1.5rem] tw-text-black-core/[0.6] tw-font-inter-regular"
            >
              <span
                class="tw-text-[1.25rem] tw-leading-[1.875rem] tw-font-inter-medium tw-text-black-core/[0.87]"
                >{{ contribution.title }}</span
              ><br />{{ contribution.description }}
            </p>
            <div class="tw-flex tw-flex-row tw-items-center tw-mt-4 tw-px-4">
              <p
                class="tw-mr-2 tw-font-inter-semibold tw-text-[1rem] tw-leading-[1.5rem]"
              >
                Contributors:
              </p>
              <div
                class="tw-flex tw-flex-row tw-items-center tw-ml-2"
                v-for="(contributor, index) in contribution.contributors"
                :key="index"
              >
                <img
                  :src="contributor"
                  class="tw-w-9 tw-h-9 tw-object-cover"
                  alt="contribution-image"
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
      class="tw-container tw-hidden lg:tw-flex tw-flex-wrap tw-justify-between tw-mx-auto tw-my-10"
    >
      <div
        v-for="(contribution, index) in contributions"
        class="flip-card tw-relative tw-w-[48%] tw-h-[440px] xl:tw-h-[500px] 2xl:tw-h-[550px] tw-mt-[2rem] tw-mx-auto tw-cursor-pointer"
        :key="index"
        @click="openContribution(contribution)"
      >
        <div class="flip-card-inner tw-relative tw-w-full tw-h-full">
          <div class="flip-card-front">
            <img
              :src="contribution.deskImage[1]"
              :srcset="`${contribution.deskImage[0]} 400w, ${contribution.deskImage[1]} 800w`"
              class="tw-absolute tw-w-full tw-h-full tw-object-cover tw-rounded-[20px] card-image card-image-1"
              alt="contribution-image"
            />
          </div>
          <div
            class="flip-card-back tw-absolute tw-w-full tw-h-full tw-rounded-[20px] tw-from-[#FF835B] tw-to-[#F2709C] tw-bg-gradient-[180deg] card-image back-card"
          >
            <div
              class="tw-flex tw-flex-row tw-items-center tw-justify-between tw-pl-6 tw-pr-4 tw-pt-6"
            >
              <div
                class="tw-flex tw-flex-row tw-items-center tw-gap-2 tw-text-white"
              >
                <span>
                  <font-awesome-icon
                    class="fa tw-w-6 tw-h-6"
                    :icon="fork" /></span
                ><span
                  class="tw-text-[1.625rem] tw-leading-[1.625rem] tw-font-inter-semibold"
                  >{{ contribution.forks }}</span
                >
              </div>
              <div class="tw-flex tw-flex-row tw-items-center tw-gap-2">
                <span
                  class="tw-text-[1.625rem] tw-leading-[1.625rem] tw-font-inter-semibold tw-text-white"
                  >{{ contribution.language }}</span
                >
              </div>
              <div
                class="tw-flex tw-flex-row tw-items-center tw-rounded-[30px] tw-bg-white tw-px-[1rem] tw-py-[0.5rem] xl:tw-px-[1.5rem]"
              >
                <font-awesome-icon
                  class="fab footer-icon tw-w-[26px] tw-h-[26px] tw-pr-[5px]"
                  :icon="star"
                />
                <span
                  class="tw-font-inter-semibold tw-text-[1.625rem] tw-leading-[1.625rem] v2-canopas-gradient-text"
                  >{{ contribution.stars }}</span
                >
              </div>
            </div>
            <p
              class="tw-mt-[4rem] tw-ml-6 tw-mr-2 tw-font-inter-semibold tw-text-[2rem] tw-leading-[3rem] tw-text-white"
            >
              {{ contribution.title }}
            </p>
            <p
              class="tw-ml-6 tw-mr-2 tw-font-inter-regular tw-text-[1.75rem] tw-leading-[2.625rem] tw-text-white"
            >
              {{ contribution.description }}
            </p>
            <div class="tw-flex tw-flex-row tw-items-center tw-mt-8 tw-px-6">
              <p
                class="tw-mr-2 tw-font-inter-semibold tw-text-[1.625rem] tw-leading-[2.438rem] tw-text-white"
              >
                Contributors:
              </p>
              <div
                class="tw-flex tw-flex-row tw-items-center tw-ml-2"
                v-for="(contributor, index) in contribution.contributors"
                :key="index"
              >
                <img
                  :src="contributor"
                  class="tw-w-9 tw-h-9 tw-object-cover tw-rounded-full"
                  alt="contribution-image"
                />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <!-- Desktop UI end -->
  </section>
</template>

<script>
import SwiperCore, { Pagination, Autoplay } from "swiper";
import { Swiper, SwiperSlide } from "swiper/vue";

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

import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { faCodeFork } from "@fortawesome/free-solid-svg-icons";
import { faStar } from "@fortawesome/free-solid-svg-icons";

SwiperCore.use([Pagination, Autoplay]);

const ANDROID = 0;
const IOS = 1;
const WEB = 2;

export default {
  data() {
    return {
      navbars: ["Android", "iOS", "Web"],
      contributions: [],
      activeIndex: 0,
      fork: faCodeFork,
      star: faStar,
      swiperRef: null,
      allContributions: [
        {
          id: 1,
          key: ANDROID,
          image: [android1_400w, android1_800w],
          deskImage: [desktop_android1_400w, desktop_android1_800w],
          link: "https://github.com/canopas/Intro-showcase-view",
          language: "Kotlin",
          color: "tw-bg-[#A97BFF]",
          forks: "22",
          title: "Intro Showcase View",
          description:
            "Intro showcase view in Jetpack compose. An implementation of Intro Showcase from ...",
          stars: "274",
          contributors: [contributor1, contributor2, contributor3],
        },
        {
          id: 2,
          key: ANDROID,
          image: [android2_400w, android2_800w],
          deskImage: [desktop_android2_400w, desktop_android2_800w],
          link: "https://github.com/canopas/Jetpack-compose-animations-examples",
          language: "Kotlin",
          color: "tw-bg-[#A97BFF]",
          forks: "19",
          title: "Jetpack Compose Animations",
          description: "Cool animations implemented with Jetpack compose....",
          stars: "203",
          contributors: [contributor1, contributor2, contributor4],
        },
        {
          id: 3,
          key: ANDROID,
          image: [android3_400w, android3_800w],
          deskImage: [desktop_android3_400w, desktop_android3_800w],
          link: "https://github.com/canopas/JetCountrypicker",
          language: "Kotlin",
          color: "tw-bg-[#A97BFF]",
          forks: "01",
          title: "JetCountryPicker",
          description:
            "Country code bottomsheet picker in Jetpack Compose with Search functionality...",
          stars: "34",
          contributors: [contributor1, contributor2, contributor4],
        },
        {
          id: 4,
          key: IOS,
          image: [ios1_400w, ios1_800w],
          deskImage: [desktop_ios1_400w, desktop_ios1_800w],
          link: "https://github.com/canopas/UIPilot",
          language: "Swift",
          color: "tw-bg-[#F05138]",
          forks: "15",
          title: "UIPilot ",
          description: "The missing typesafe SwiftUI navigation library.",
          stars: "227",
          contributors: [contributor1, contributor6, contributor2],
        },
        {
          id: 5,
          key: IOS,
          image: [ios2_400w, ios2_800w],
          deskImage: [desktop_ios2_400w, desktop_ios2_800w],
          link: "https://github.com/canopas/Swiftui-animations-examples",
          language: "Swift",
          color: "tw-bg-[#F05138]",
          forks: "02",
          title: "Swiftui Animations",
          description: "Cool animations implemented with SwiftUI",
          stars: "37",
          contributors: [contributor1, contributor6],
        },
        {
          id: 6,
          key: IOS,
          image: [ios3_400w, ios3_800w],
          deskImage: [desktop_ios3_400w, desktop_ios3_800w],
          link: "https://github.com/canopas/iOS-developer-roadmap",
          language: "Swift",
          color: "tw-bg-[#F05138]",
          forks: "11",
          title: "iOS Developer Roadmap 2022",
          description:
            "iOS Developer Roadmap 2022 is a learning path to understand iOS development...",
          stars: "43",
          contributors: [contributor1, contributor6],
        },
        {
          id: 7,
          key: WEB,
          image: [web1_400w, web1_800w],
          deskImage: [desktop_web1_400w, desktop_web1_800w],
          link: "https://github.com/canopas/tailwind-animations",
          language: "JavaScript",
          color: "tw-bg-[#F1E05A]",
          forks: "05",
          title: "Tailwind Animations",
          description:
            "This repository contains different animations implemented using tailwind css...",
          stars: "59",
          contributors: [contributor8, contributor9],
        },
        {
          id: 8,
          key: WEB,
          image: [web2_400w, web2_800w],
          deskImage: [desktop_web2_400w, desktop_web2_800w],
          link: "https://github.com/canopas/web-developer-roadmap-2023",
          language: "",
          color: "",
          forks: "10",
          title: "Web Development Roadmap",
          description:
            "Web Developer Roadmap is a path to understand web development including...",
          stars: "00",
          contributors: [contributor8, contributor10],
        },
        {
          id: 9,
          key: WEB,
          image: [web3_400w, web3_800w],
          deskImage: [desktop_web4_400w, desktop_web4_800w],
          link: "https://github.com/canopas/canopas-blog",
          language: "JavaScript",
          color: "tw-bg-[#F1E05A]",
          forks: "00",
          title: "Canopas Blog website",
          description:
            "We intend to keep this open source. Plan is to keep the repository up to date with...",
          stars: "02",
          contributors: [contributor8, contributor11, contributor12],
        },
        {
          id: 10,
          key: WEB,
          image: [web4_400w, web4_800w],
          deskImage: [desktop_web3_400w, desktop_web3_800w],
          link: "https://github.com/canopas/canopas-website",
          language: "Vue",
          color: "tw-bg-[#41B883]",
          forks: "02",
          title: "Canopas Website",
          description:
            "We intend to keep this open source. Plan is to keep the repository up to date with latest...",
          stars: "07",
          contributors: [contributor8, contributor14, contributor12],
        },
      ],
    };
  },
  components: {
    Swiper,
    SwiperSlide,
    FontAwesomeIcon,
  },
  mounted() {
    this.contributions = this.allContributions.filter((obj) => {
      return obj.key === ANDROID;
    });
  },
  methods: {
    openContribution(contribution) {
      window.open(contribution.link, "_blank");
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
<style lang="postcss">
@import "swiper/css";
@import "swiper/css/pagination";

.swiper-wrapper {
  @apply !tw-items-center;
}

.swiper {
  @apply tw-z-0;
}

.flip-card {
  background-color: transparent;
  perspective: 1000px;
}

.flip-card-inner {
  transition: transform 0.6s;
  transform-style: preserve-3d;
}

.flip-card:hover .flip-card-inner {
  transform: rotateY(180deg);
}

.flip-card-front,
.flip-card-back {
  -webkit-backface-visibility: hidden;
  backface-visibility: hidden;
}

.flip-card-back {
  transform: rotateY(180deg);
}
</style>
