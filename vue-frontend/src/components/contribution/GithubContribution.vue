<template>
  <section class="lg:tw-mt-[100px]">
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
          @click="
            [
              index == 0 ? slideTo(1) : '',
              index == 1 ? slideTo(4) : '',
              index == 2 ? slideTo(7) : '',
            ],
              (activeIndex = index),
              changeContributions(index)
          "
          :class="
            activeIndex == index
              ? 'tw-bg-gradient-[270.11deg] tw-from-[#ff835b] tw-to-[#f2709c] tw-text-white '
              : 'tw-border-[1px] tw-border-solid tw-border-transparent tw-text-black-core/[0.6]'
          "
        >
          {{ navbar.name }}
        </button>
      </div>
    </div>
    <div class="swiper-content tw-mt-4 lg:tw-hidden">
      <swiper
        :slidesPerView="1"
        :centeredSlides="true"
        :spaceBetween="20"
        @swiper="setSwiperRef"
        class="swiper-container"
      >
        <swiper-slide
          v-for="(contribution, index) in allContributions"
          :key="index"
          @touchstart.passive="false"
          @touchmove.passive="true, onSlideChange(index)"
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
            <div class="tw-flex tw-flex-row tw-items-center tw-gap-6 tw-px-4">
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
                :key="index"
                v-for="(contributor, index) in contribution.contributors"
              >
                <img
                  :src="contributor.image"
                  class="tw-w-9 tw-h-9 tw-object-cover"
                  alt="contribution-image"
                />
              </div>
            </div>
          </div>
        </swiper-slide>
      </swiper>
    </div>
    <div
      class="tw-container tw-hidden lg:tw-flex tw-flex-wrap tw-justify-between tw-mx-auto tw-my-10"
    >
      <div
        v-for="(contribution, index) in contributions"
        class="tw-relative tw-w-[48%] tw-h-[440px] xl:tw-h-[500px] 2xl:tw-h-[550px] contribution-card tw-mt-[2rem] xl:tw-mt-[2.5rem] 2xl:tw-mt-[3rem] tw-cursor-pointer"
        :class="[
          contributions.length == 3 && index == 2 ? '!tw-mx-auto ' : '',
          contributions.length == 4 ? '!tw-mx-0' : '',
        ]"
        :key="index"
        @click="openContribution(contribution)"
      >
        <img
          :src="contribution.deskImage[1]"
          :srcset="`${contribution.deskImage[0]} 400w, ${contribution.deskImage[1]} 800w`"
          class="tw-absolute tw-w-full tw-h-full tw-object-cover tw-rounded-[20px] card-image"
          alt="contribution-image"
        />
        <div
          class="tw-absolute tw-w-full tw-h-full tw-rounded-[20px] tw-from-[#FF835B] tw-to-[#F2709C] tw-bg-gradient-[180deg] card-image back-card"
        >
          <div
            class="tw-flex tw-flex-row tw-items-center tw-gap-[5rem] xl:tw-gap-[7rem] 2xl:tw-gap-[10rem] tw-pl-6 tw-pr-4 tw-pt-6"
            :class="
              contribution.language == 'JavaScript'
                ? 'lg:tw-gap-[3rem] xl:!tw-gap-[5rem] 2xl:tw-gap-[8rem]'
                : ''
            "
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
              :key="index"
              v-for="(contributor, index) in contribution.contributors"
            >
              <img
                :src="contributor.image"
                class="tw-w-9 tw-h-9 tw-object-cover"
                alt="contribution-image"
              />
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import SwiperCore, { Pagination, Autoplay } from "swiper";
import { Swiper, SwiperSlide } from "swiper/vue";
import android1_400w from "@/assets/images/contributions/github/android-1-400w.webp";
import android1_800w from "@/assets/images/contributions/github/android-1-800w.webp";
import android2_400w from "@/assets/images/contributions/github/android-2-400w.webp";
import android2_800w from "@/assets/images/contributions/github/android-2-800w.webp";
import android3_400w from "@/assets/images/contributions/github/android-3-400w.webp";
import android3_800w from "@/assets/images/contributions/github/android-3-800w.webp";
import ios1_400w from "@/assets/images/contributions/github/ios-1-400w.webp";
import ios1_800w from "@/assets/images/contributions/github/ios-1-800w.webp";
import ios2_400w from "@/assets/images/contributions/github/ios-2-400w.webp";
import ios2_800w from "@/assets/images/contributions/github/ios-2-800w.webp";
import ios3_400w from "@/assets/images/contributions/github/ios-3-400w.webp";
import ios3_800w from "@/assets/images/contributions/github/ios-3-800w.webp";
import web1_400w from "@/assets/images/contributions/github/web-1-400w.webp";
import web1_800w from "@/assets/images/contributions/github/web-1-800w.webp";
import web2_400w from "@/assets/images/contributions/github/web-2-400w.webp";
import web2_800w from "@/assets/images/contributions/github/web-2-800w.webp";
import web3_400w from "@/assets/images/contributions/github/web-3-400w.webp";
import web3_800w from "@/assets/images/contributions/github/web-3-800w.webp";
import web4_400w from "@/assets/images/contributions/github/web-4-400w.webp";
import web4_800w from "@/assets/images/contributions/github/web-4-800w.webp";
import deskandroid1_400w from "@/assets/images/contributions/github/desk-android-1-400w.webp";
import deskandroid1_800w from "@/assets/images/contributions/github/desk-android-1-800w.webp";
import deskandroid2_400w from "@/assets/images/contributions/github/desk-android-2-400w.webp";
import deskandroid2_800w from "@/assets/images/contributions/github/desk-android-2-800w.webp";
import deskandroid3_400w from "@/assets/images/contributions/github/desk-android-3-400w.webp";
import deskandroid3_800w from "@/assets/images/contributions/github/desk-android-3-800w.webp";
import deskios1_400w from "@/assets/images/contributions/github/desk-ios-1-400w.webp";
import deskios1_800w from "@/assets/images/contributions/github/desk-ios-1-800w.webp";
import deskios2_400w from "@/assets/images/contributions/github/desk-ios-2-400w.webp";
import deskios2_800w from "@/assets/images/contributions/github/desk-ios-2-800w.webp";
import deskios3_400w from "@/assets/images/contributions/github/desk-ios-3-400w.webp";
import deskios3_800w from "@/assets/images/contributions/github/desk-ios-3-800w.webp";

import deskweb1_400w from "@/assets/images/contributions/github/desk-web-1-400w.webp";
import deskweb1_800w from "@/assets/images/contributions/github/desk-web-1-800w.webp";
import deskweb2_400w from "@/assets/images/contributions/github/desk-web-2-400w.webp";
import deskweb2_800w from "@/assets/images/contributions/github/desk-web-2-800w.webp";
import deskweb3_400w from "@/assets/images/contributions/github/desk-web-3-400w.webp";
import deskweb3_800w from "@/assets/images/contributions/github/desk-web-3-800w.webp";
import deskweb4_400w from "@/assets/images/contributions/github/desk-web-4-400w.webp";
import deskweb4_800w from "@/assets/images/contributions/github/desk-web-4-800w.webp";

import contributor1 from "@/assets/images/contributions/github/contributors/contributor-1-100w.webp";
import contributor2 from "@/assets/images/contributions/github/contributors/contributor-2-100w.webp";
import contributor3 from "@/assets/images/contributions/github/contributors/contributor-3-100w.webp";
import contributor4 from "@/assets/images/contributions/github/contributors/contributor-4-100w.webp";
import contributor5 from "@/assets/images/contributions/github/contributors/contributor-5-100w.webp";
import contributor6 from "@/assets/images/contributions/github/contributors/contributor-6-100w.webp";
import contributor7 from "@/assets/images/contributions/github/contributors/contributor-7-100w.webp";
import contributor8 from "@/assets/images/contributions/github/contributors/contributor-8-100w.webp";
import contributor9 from "@/assets/images/contributions/github/contributors/contributor-9-100w.webp";
import contributor10 from "@/assets/images/contributions/github/contributors/contributor-10-100w.webp";
import contributor11 from "@/assets/images/contributions/github/contributors/contributor-11-100w.webp";
import contributor12 from "@/assets/images/contributions/github/contributors/contributor-12-100w.webp";
import contributor13 from "@/assets/images/contributions/github/contributors/contributor-13-100w.webp";
import contributor14 from "@/assets/images/contributions/github/contributors/contributor-14-100w.webp";
import contributor15 from "@/assets/images/contributions/github/contributors/contributor-15-100w.webp";

import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { faCodeFork } from "@fortawesome/free-solid-svg-icons";
import { faStar } from "@fortawesome/free-solid-svg-icons";

SwiperCore.use([Pagination, Autoplay]);

export default {
  data() {
    return {
      navbars: [
        {
          name: "Android",
        },
        {
          name: "iOS",
        },
        {
          name: "Web",
        },
      ],
      contributions: [],
      activeIndex: 0,
      fork: faCodeFork,
      star: faStar,
      swiper: null,
      contributor1,
      contributor2,
      contributor3,
      contributor4,
      contributor5,
      contributor6,
      contributor7,
      contributor8,
      contributor9,
      contributor10,
      contributor11,
      contributor12,
      contributor13,
      contributor14,
      contributor15,
      allContributions: [
        {
          id: 1,
          image: [android1_400w, android1_800w],
          deskImage: [deskandroid1_400w, deskandroid1_800w],
          link: "https://github.com/canopas/Intro-showcase-view",
          language: "Kotlin",
          color: "tw-bg-[#A97BFF]",
          forks: "20",
          title: "Intro Showcase View",
          description:
            "Intro showcase view in Jetpack compose. An implementation of Intro Showcase from ...",
          stars: "263",
          contributors: [
            {
              image: contributor1,
            },
            {
              image: contributor2,
            },
            {
              image: contributor3,
            },
          ],
        },
        {
          id: 2,
          image: [android2_400w, android2_800w],
          deskImage: [deskandroid2_400w, deskandroid2_800w],
          link: "https://github.com/canopas/Jetpack-compose-animations-examples",
          language: "Kotlin",
          color: "tw-bg-[#A97BFF]",
          forks: "19",
          title: "Jetpack Compose Animations",
          description: "Cool animations implemented with Jetpack compose....",
          stars: "263",
          contributors: [
            {
              image: contributor1,
            },
            {
              image: contributor2,
            },
            {
              image: contributor4,
            },
          ],
        },
        {
          id: 3,
          image: [android3_400w, android3_800w],
          deskImage: [deskandroid3_400w, deskandroid3_800w],
          link: "https://github.com/canopas/JetCountrypicker",
          language: "Kotlin",
          color: "tw-bg-[#A97BFF]",
          forks: "01",
          title: "JetCountryPicker",
          description:
            "Country code bottomsheet picker in Jetpack Compose with Search functionality...",
          stars: "35",
          contributors: [
            {
              image: contributor1,
            },
            {
              image: contributor2,
            },
            {
              image: contributor4,
            },
            {
              image: contributor5,
            },
          ],
        },
        {
          id: 4,
          image: [ios1_400w, ios1_800w],
          deskImage: [deskios1_400w, deskios1_800w],
          link: "https://github.com/canopas/UIPilot",
          language: "Swift",
          color: "tw-bg-[#F05138]",
          forks: "12",
          title: "UIPilot ",
          description: "The missing typesafe SwiftUI navigation library.",
          stars: "224",
          contributors: [
            {
              image: contributor1,
            },
            {
              image: contributor6,
            },
            {
              image: contributor2,
            },
            {
              image: contributor7,
            },
          ],
        },
        {
          id: 5,
          image: [ios2_400w, ios2_800w],
          deskImage: [deskios2_400w, deskios2_800w],
          link: "https://github.com/canopas/Swiftui-animations-examples",
          language: "Swift",
          color: "tw-bg-[#F05138]",
          forks: "02",
          title: "Swiftui Animations",
          description: "Cool animations implemented with SwiftUI",
          stars: "36",
          contributors: [
            {
              image: contributor1,
            },
            {
              image: contributor6,
            },
          ],
        },
        {
          id: 6,
          image: [ios3_400w, ios3_800w],
          deskImage: [deskios3_400w, deskios3_800w],
          link: "https://github.com/canopas/iOS-developer-roadmap",
          language: "Swift",
          color: "tw-bg-[#F05138]",
          forks: "10",
          title: "iOS Developer Roadmap 2022",
          description:
            "iOS Developer Roadmap 2022 is a learning path to understand iOS development...",
          stars: "29",
          contributors: [
            {
              image: contributor1,
            },
            {
              image: contributor6,
            },
          ],
        },
        {
          id: 7,
          image: [web1_400w, web1_800w],
          deskImage: [deskweb1_400w, deskweb1_800w],
          link: "https://github.com/canopas/tailwind-animations",
          language: "JavaScript",
          color: "tw-bg-[#F1E05A]",
          forks: "05",
          title: "Tailwind Animations",
          description:
            "This repository contains different animations implemented using tailwind css...",
          stars: "49",
          contributors: [
            {
              image: contributor8,
            },
            {
              image: contributor9,
            },
          ],
        },
        {
          id: 8,
          image: [web2_400w, web2_800w],
          deskImage: [deskweb2_400w, deskweb2_800w],

          link: "https://github.com/canopas/web-developer-roadmap-2023",
          language: "JAVA",
          color: "tw-bg-[#F1E05A]",
          forks: "10",
          title: "Web Development Roadmap",
          description:
            "Web Developer Roadmap is a path to understand web development including...",
          stars: "72",
          contributors: [
            {
              image: contributor8,
            },
            {
              image: contributor10,
            },
          ],
        },
        {
          id: 9,
          image: [web3_400w, web3_800w],
          deskImage: [deskweb3_400w, deskweb3_800w],

          link: "https://github.com/canopas/canopas-blog",
          language: "JavaScript",
          color: "tw-bg-[#F1E05A]",
          forks: "00",
          title: "Canopas Blog website",
          description:
            "We intend to keep this open source. Plan is to keep the repository up to date with...",
          stars: "02",
          contributors: [
            {
              image: contributor8,
            },
            {
              image: contributor11,
            },
            {
              image: contributor12,
            },
            {
              image: contributor13,
            },
          ],
        },
        {
          id: 10,
          image: [web4_400w, web4_800w],
          deskImage: [deskweb4_400w, deskweb4_800w],
          link: "https://github.com/canopas/canopas-website",
          language: "Vue",
          color: "tw-bg-[#41B883]",
          forks: "02",
          title: "Canopas Website",
          description:
            "We intend to keep this open source. Plan is to keep the repository up to date with latest...",
          stars: "07",
          contributors: [
            {
              image: contributor8,
            },
            {
              image: contributor1,
            },
            {
              image: contributor12,
            },
            {
              image: contributor14,
            },
            {
              image: contributor15,
            },
            {
              image: contributor10,
            },
            {
              image: contributor11,
            },
            {
              image: contributor13,
            },
          ],
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
    this.contributions = this.allContributions.slice(0, 3);
  },
  methods: {
    openContribution(contribution) {
      window.open(contribution.link, "_blank");
    },
    onSlideChange(index) {
      if (index <= 1) {
        this.activeIndex = 0;
      } else if (index > 1 && index <= 4) {
        this.activeIndex = 1;
      } else if (index > 4 && index <= 7) {
        this.activeIndex = 2;
      }
    },
    setSwiperRef(swiper) {
      this.swiperRef = swiper;
    },
    slideTo(index) {
      this.swiperRef.slideTo(index - 1, 0);
    },
    changeContributions(index) {
      if (index == 0) {
        this.contributions = this.allContributions.slice(0, 3);
      } else if (index == 1) {
        this.contributions = this.allContributions.slice(3, 6);
      } else {
        this.contributions = this.allContributions.slice(6, 10);
      }
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
.contribution-card {
  background-color: transparent;
  transition: transform 0.6s;
  transform-style: preserve-3d;
  perspective: 1000px;
}
.contribution-card:hover {
  transform: rotateY(180deg);
}
.card-image {
  backface-visibility: hidden;
}
.back-card {
  transform: rotateY(180deg);
}
</style>
