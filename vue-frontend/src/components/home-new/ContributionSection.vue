<template>
  <div
    class="tw-container lg:tw-flex lg:tw-flex-row-reverse tw-justify-center tw-max-w-[1878px] lg:tw-pt-20 tw-text-center lg:tw-text-left"
  >
    <div
      class="lg:tw-basis-[25%] tw-my-auto tw-mx-auto lg:tw-ml-11 2xl:tw-mx-0 tw-w-[290px] sm:tw-w-full tw-pb-14 md:tw-pb-20 lg:tw-pb-0"
    >
      <h1
        class="tw-font-inter-bold tw-text-black-core/[.87] tw-text-[1.875rem] tw-leading-[2.5rem] md:tw-text-[3.4375rem] md:tw-leading-[5.5625rem]"
      >
        We are here to Contribute
      </h1>
      <div class="tw-mt-2 md:tw-mt-0">
        <span
          class="tw-font-inter-medium tw-text-black-core/[0.6] tw-text-[1rem] tw-leading-[1.375rem] md:tw-text-[1.5rem] md:tw-leading-[2.5625rem]"
          >and that starts with the site itself,
          <a
            class="tw-relative tw-bg-gradient-[1deg] tw-pb-[5px] v2-canopas-gradient-text tw-font-inter-medium gradient-underline tw-no-underline after:tw-absolute after:tw-w-full after:tw-h-[2px] after:tw-bg-gradient-to-l after:tw-to-pink-300 after:tw-from-orange-300 after:tw-bottom-0 after:tw-left-0"
            :href="websiteOpenSourceUrl"
            target="_blank"
            @click.native="mixpanel.track('tap_canopas_website_github')"
            >canopas is open source!</a
          ></span
        >
      </div>
    </div>
    <div
      class="tw-relative lg:tw-basis-[70%] 2xl:tw-basis-[60%] tw-h-[35rem] md:tw-h-[50rem] tw-text-left tw-overflow-y-scroll hidden-scrollbar"
    >
      <div
        v-for="(contribution, index) in contributions"
        :key="contribution"
        class="tw-flex tw-flex-row tw-justify-center tw-mt-14 sm:tw-mt-4 md:tw-mt-8 md:tw-mb-6"
        :class="index % 2 != 0 ? 'tw-flex-row-reverse' : ''"
      >
        <div
          class="tw-w-[13.625rem] md:tw-w-fit tw-h-[26.875rem] md:tw-h-fit tw-shadow-[2px_2px_10px_rgba(0,0,0,0.1)]"
        >
          <img
            :src="contribution.video"
            :alt="contribution.title + `image`"
            class="tw-w-full tw-h-full"
            loading="lazy"
          />
        </div>

        <div
          @click="openLink(contribution)"
          class="tw-flex tw-flex-col tw-basis-[40%] lg:tw-basis-[43%] xl:tw-basis-[50%] 2xl:tw-basis-[40%] tw-my-auto md:tw-w-[21.563rem] tw-bg-[#FFFFFF] tw-p-2.5 sm:tw-p-4 xl:tw-p-8 tw-shadow-[0_0px_10px_rgba(0,0,0,0.1)] active:tw-scale-[0.98]"
          :class="[
            isMobile && index % 2 == 0
              ? 'tw-absolute  tw-ml-48 sm:tw-ml-56 md:tw-ml-[22rem] tw-mt-[8.5rem] sm:tw-mt-32 tw-w-[10.563rem] sm:tw-w-[12.563rem]'
              : isMobile
              ? 'tw-absolute tw-mr-48 sm:tw-mr-56 md:tw-mr-[22rem] tw-mt-[8.5rem] sm:tw-mt-32 tw-w-[10.563rem] sm:tw-w-[12.563rem]'
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
                    @click.native="mixpanel.track('tap_library')"
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
                :aria-label="contribution.title"
                target="_blank"
                @click.native="mixpanel.track('tap_library')"
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

import UIPilot from "@/assets/images/contribution/animations/UIPilot.gif";
import IntroShowCase from "@/assets/images/contribution/animations/introShowCase.gif";
import JetpackComposeAnimations from "@/assets/images/contribution/animations/JetpackComposeAnimations.gif";
import SwiftUIAnimations from "@/assets/images/contribution/animations/SwiftUIAnimations.gif";
import TailwindAnimations from "@/assets/images/contribution/animations/TailwindAnimations.gif";

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
          video: UIPilot,
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
          video: TailwindAnimations,
          starPoints: "13",
          link: "https://github.com/canopas/tailwind-animations",
        },
        {
          title: "Swiftui-animations",
          description: "Cool animations implemented with SwiftUI",
          author: "Amisha I.",
          video: SwiftUIAnimations,
          starPoints: "30",
          link: "https://github.com/canopas/Swiftui-animations-examples",
        },
      ],
    };
  },
  methods: {
    openLink(contributions) {
      window.open(contributions.link, "_blank");
    },
  },
  components: {
    AspectRatio,
    FontAwesomeIcon,
  },
  inject: ["mixpanel"],
  mounted() {
    if (window.innerWidth <= 768) {
      this.isMobile = true;
    }
  },
};
</script>
