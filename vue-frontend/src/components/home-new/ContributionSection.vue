<template>
  <div class="tw-relative">
    <div class="tw-container tw-mt-2 md:tw-mt-4">
      <div
        class="tw-sticky tw-w-full tw-h-[30vh] sm:tw-h-[25vh] lg:tw-h-[35vh] tw-z-[1] tw-top-[6rem] tw-bg-white tw-pt-12 tw-text-center"
      >
        <h1
          class="tw-font-inter-bold tw-text-black-core/[.87] tw-text-[1.875rem] tw-leading-[2.5rem] md:tw-text-[3.4375rem] md:tw-leading-[5.5625rem]"
        >
          We are here to Contribute
        </h1>
        <span
          class="tw-font-inter-medium tw-text-black-core/[0.6] tw-text-[1rem] tw-leading-[1.375rem] md:tw-text-[1.5rem] md:tw-leading-[2.5625rem]"
          >and that starts with the site itself,
          <a
            class="tw-relative tw-pb-[5px] v2-canopas-gradient-text tw-font-inter-medium gradient-underline tw-no-underline after:tw-absolute after:tw-w-full after:tw-h-[2px] after:tw-bg-gradient-to-l tw-to-pink-300 tw-from-orange-300 tw-bg-gradient-[1deg] after:tw-bottom-0 after:tw-left-0"
            :href="websiteOpenSourceUrl"
            target="_blank"
            @click.native="$gtag.event('tap_canopas_website_github')"
            >canopas is open source!</a
          ></span
        >
      </div>

      <div
        v-for="(contribution, index) in contributions"
        :key="contribution"
        class="tw-flex tw-flex-row tw-justify-center tw-mt-14 sm:tw-mt-4"
        :class="index % 2 != 0 ? 'tw-flex-row-reverse' : ''"
      >
        <div
          class="tw-w-[13.625rem] md:tw-w-fit tw-h-[26.875rem] md:tw-h-fit tw-shadow-[2px_2px_10px_rgba(0,0,0,0.1)]"
        >
          <img
            :src="contribution.video"
            :alt="contribution.title + `image`"
            class="tw-w-full tw-h-full"
          />
        </div>

        <div
          class="tw-flex tw-flex-col tw-basis-[40%] md:tw-basis-[55%] lg:tw-basis-[35%] tw-my-auto tw-bg-[#FFFFFF] tw-p-2.5 sm:tw-p-4 tw-shadow-[0_0px_10px_rgba(0,0,0,0.1)]"
          :class="[
            isMobile && index % 2 == 0
              ? 'tw-absolute  tw-ml-40 sm:tw-ml-52 tw-mt-[8.5rem] sm:tw-mt-32  tw-w-[12.563rem]'
              : isMobile
              ? 'tw-absolute tw-mr-40 sm:tw-mr-52  tw-mt-[8.5rem]  sm:tw-mt-32  tw-w-[12.563rem]'
              : ' ',
          ]"
        >
          <div>
            <div
              @touchstart.passive="activeIndex = index"
              :class="index % 2 != 0 ? 'tw-flex-row-reverse' : ''"
              class="tw-flex tw-flex-row tw-justify-between tw-items-center"
            >
              <div
                class="tw-flex tw-items-center tw-justify-center tw-w-[2.625rem] md:tw-w-[4.75rem] tw-h-[1.125rem] md:tw-h-[1.875rem] tw-from-[#ff835b] tw-to-[#f2709c] tw-bg-gradient-[270.11deg] tw-p-px md:tw-px-2.5 tw-rounded-full tw-font-inter-bold md:tw-font-inter-semibold tw-text-[0.625rem] md:tw-text-[1rem] tw-text-white"
              >
                <span>
                  <font-awesome-icon
                    class="fa tw-w-[10px] md:tw-w-[18px] tw-h-[10px] md:tw-h-[18px] tw-pr-[5px] tw-text-white"
                    icon="star"
                  />{{ contribution.starPoints }}</span
                >
              </div>
              <div>
                <div>
                  <a
                    class="tw-flex tw-items-center tw-font-inter-bold tw-text-[0.75rem] md:tw-text-[1.25rem] tw-leading-[2.5625rem] tw-duration-300 tw-ease-in-out hover:tw-scale-[1.1] tw-hover-hover"
                    :href="contribution.link"
                    target="_blank"
                    @click.native="$gtag.event('tap_library')"
                  >
                    <font-awesome-icon
                      class="tw-mr-1 md:tw-mr-2.5 tw-w-[1rem] md:tw-w-[1.75rem] tw-h-[1rem] md:tw-h-[1.75rem] tw-text-pink-300"
                      :icon="icon"
                    />
                    <div class="v2-canopas-gradient-text">GitHub</div>
                  </a>
                </div>
              </div>
            </div>
            <div>
              <h1
                class="md:tw-mt-[41px] tw-font-inter-semibold tw-text-[1rem] md:tw-text-[1.875rem] tw-leading-[1.3rem] md:tw-leading-[2rem] tw-text-black-core/[0.87]"
              >
                {{ contribution.title }}
              </h1>
            </div>
            <div class="tw-mt-5">
              <span
                class="tw-hidden md:tw-block tw-font-inter-light tw-text-[1.25rem] tw-leading-[1.578rem] tw-text-black-core/[0.6]"
                >{{ contribution.description }}</span
              >
            </div>
            <div
              @touchstart.passive="activeIndex = index"
              class="tw-flex tw-flex-row tw-justify-between tw-mt-[10px] md:tw-mt-[30px] tw-font-inter-regular tw-text-[0.625rem] md:tw-text-[1rem] tw-leading-[0.9375rem] md:tw-leading-[1.5rem] tw-text-black-core/[0.87]"
            >
              <h1>by {{ contribution.author }}</h1>
              <a
                class="tw-duration-300 tw-ease-in-out hover:tw-scale-[1.1] tw-hover-hover"
                :href="contribution.link"
                target="_blank"
                @click.native="$gtag.event('tap_library')"
              >
                <font-awesome-icon
                  class="arrow fa tw-w-4 tw-h-4"
                  :icon="rightArrow"
                  id="leftArrow"
                />
              </a>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script type="module">
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { library } from "@fortawesome/fontawesome-svg-core";
import { faGithub } from "@fortawesome/free-brands-svg-icons";
import { faArrowRightLong } from "@fortawesome/free-solid-svg-icons";

import ComplexRouting from "@/assets/images/contribution/animations/ComplexRouting.gif";
import IntroShowCase from "@/assets/images/contribution/animations/introShowCase.gif";
import JetpackComposeAnimations from "@/assets/images/contribution/animations/JetpackComposeAnimations.gif";
import SwiftUIAnimation from "@/assets/images/contribution/animations/SwiftUIAnimation.gif";
import Tailwind from "@/assets/images/contribution/animations/Tailwind.gif";

import AspectRatio from "@/components/utils/AspectRatio.vue";
import Config from "@/config.js";

library.add(faGithub);

export default {
  data() {
    return {
      icon: faGithub,
      rightArrow: faArrowRightLong,
      isMobile: false,
      websiteOpenSourceUrl: Config.WEBSITE_OPEN_SOURCE_URL,
      contributions: [
        {
          title: "Intro Showcase view in jetpack compose",
          description:
            "Highlight different features of the app using Jetpack Compose",
          author: "Radhika S.",
          video: IntroShowCase,
          starPoints: "237",
          link: "https://github.com/canopas/Intro-showcase-view",
        },
        {
          title: "UIPilot",
          description: "The missing typesafe SwiftUI navigation library",
          author: "Jimmy S.",
          video: ComplexRouting,
          starPoints: "201",
          link: "https://github.com/canopas/UIPilot",
        },
        {
          title: "Jetpack Compose animations",
          description: "Cool animations implemented with Jetpack compose",
          author: "Radhika S.",
          video: JetpackComposeAnimations,
          starPoints: "186",
          link: "https://github.com/canopas/Jetpack-compose-animations-examples",
        },
        {
          title: "Tailwind Animations",
          description: "Cool animations implemented with Tailwind CSS",
          author: "Sumita K.",
          video: Tailwind,
          starPoints: "13",
          link: "https://github.com/canopas/tailwind-animations",
        },
        {
          title: "Swiftui-animations",
          description: "Cool animations implemented with SwiftUI",
          author: "Amisha I.",
          video: SwiftUIAnimation,
          starPoints: "30",
          link: "https://github.com/canopas/Swiftui-animations-examples",
        },
      ],
    };
  },
  components: {
    AspectRatio,
    FontAwesomeIcon,
  },
  mounted() {
    if (window.innerWidth <= 768) {
      this.isMobile = true;
    }
  },
};
</script>
