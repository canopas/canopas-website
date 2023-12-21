<template>
  <div
    class="container justify-center mt-36 sm:mt-20 md:mt-24 mb-16 text-center"
  >
    <div class="m-auto w-[290px] sm:w-full pb-14 md:pb-8">
      <span class="mobile-header-2 lg:desk-header-2 text-black-87">
        We are here to Contribute
      </span>
      <div class="mt-4 sub-h1-regular w-full mx-auto text-black-60">
        <span>and that starts with the site itself, </span>
        <span>
          <a
            class="relative bg-gradient-[1deg] pb-[5px] v2-canopas-gradient-text sub-h3-semibold gradient-underline no-underline after:absolute after:bottom-0 after:left-0 after:w-full after:h-0.5 after:bg-gradient-to-l after:to-pink-300 after:from-orange-300"
            :href="websiteOpenSourceUrl"
            target="_blank"
            @click.native="$mixpanel.track('tap_canopas_website_github')"
            >canopas is open source!</a
          ></span
        >
      </div>
    </div>
    <div
      v-for="(contribution, index) in contributions"
      :key="contribution"
      class="relative flex mt-8 w-full text-left"
      :class="index % 2 != 0 ? 'flex-row-reverse' : 'flex-row'"
    >
      <div class="w-[45%] shadow-[2px_2px_10px_rgba(0,0,0,0.1)] rounded-xl">
        <video class="lozad rounded-xl" autoplay loop muted playsinline>
          <source :data-src="contribution.video[1]" type="video/webm" />
          <source :data-src="contribution.video[0]" type="video/mp4" />
        </video>
      </div>

      <div
        @click="
          openBlog(contribution.link, 'tap_home_github_contribution_section')
        "
        class="absolute top-[50%] flex flex-col w-[67%] md:w-[52%] translate-y-[-50%] bg-[#FFFFFF] p-2.5 sm:p-4 shadow-[0_0px_10px_rgba(0,0,0,0.1)] active:scale-[0.98] rounded-xl md:rounded"
        :class="[
          index % 2 == 0
            ? 'left-[32%] md:left-[48%]'
            : 'right-[32%] md:right-[48%]',
        ]"
      >
        <div>
          <div
            :class="index % 2 != 0 ? 'flex-row-reverse' : ''"
            class="flex flex-row justify-between items-center"
          >
            <div
              class="flex items-center justify-center w-[2.938rem] h-[1.375rem] md:w-[3.9375rem] md:h-[2.125rem] from-[#f2709c] to-[#ff835b] px-[3px] py-[2px] bg-gradient-[270.11deg] rounded sub-h4-semibold text-white"
            >
              <span>
                <Icon
                  class="fa w-3 h-3 pr-[5px] box-content mb-[0.225em]"
                  name="fa6-solid:star"
                />{{ contribution.stars }}</span
              >
            </div>
            <a
              class="flex items-center sub-h4-semibold"
              :href="contribution.link"
              target="_blank"
              @click.native="$mixpanel.track('tap_library')"
            >
              <Icon
                class="mr-1 md:mr-2.5 w-4 md:w-[1.75rem] h-4 md:h-7 footer-icon"
                name="fa6-brands:github"
              />
              <div class="v2-canopas-gradient-text">GitHub</div>
            </a>
          </div>
          <p class="mt-4 sub-h3-semibold text-black-87 w-[73%] md:w-full">
            {{ contribution.title }}
          </p>
          <div class="mt-3 hidden md:block">
            <span class="sub-h3-regular text-black-60">{{
              contribution.description
            }}</span>
          </div>
          <div
            class="flex flex-row justify-between mt-3 md:mt-6 sub-h4-regular text-black-60"
          >
            <span>by {{ contribution.author }}</span>
            <a
              :href="contribution.link"
              :aria-label="contribution.title"
              target="_blank"
              @click.native="$mixpanel.track('tap_library')"
            >
              <Icon
                class="arrow fa w-4 h-4 box-content text-black-87"
                name="fa6-solid:arrow-right-long"
                id="leftArrow"
              />
            </a>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script type="module">
import uIPilotMp4 from "@/assets/images/contribution/animations/UIPilot.mp4";
import uIPilotWebm from "@/assets/images/contribution/animations/UIPilot.webm";
import introShowCaseMp4 from "@/assets/images/contribution/animations/introShowCase.mp4";
import introShowCaseWebm from "@/assets/images/contribution/animations/introShowCase.webm";
import jcAnimationsMp4 from "@/assets/images/contribution/animations/JetpackComposeAnimations.mp4";
import jcAnimationsWebm from "@/assets/images/contribution/animations/JetpackComposeAnimations.webm";
import config from "@/config.js";
import lozad from "lozad";
import { setGithubStars, openBlog } from "@/utils.js";
import { mapState } from "pinia";
import { useContributionStore } from "@/stores/contribution";

export default {
  computed: {
    ...mapState(useContributionStore, {
      cData: "data",
    }),
  },
  data() {
    return {
      openBlog,
      websiteOpenSourceUrl: config.WEBSITE_OPEN_SOURCE_URL,
      contributions: [
        {
          title: "Intro Showcase view in jetpack compose",
          stars: config.INTRO_SHOWCASE_STARS,
          description:
            "Highlight different features of the app using Jetpack Compose",
          author: "Radhika S.",
          video: [introShowCaseMp4, introShowCaseWebm],
          link: "https://github.com/canopas/Intro-showcase-view",
        },
        {
          title: "UIPilot",
          stars: config.UIPILOT_STARS,
          description: "The missing typesafe SwiftUI navigation library",
          author: "Jimmy S.",
          video: [uIPilotMp4, uIPilotWebm],
          link: "https://github.com/canopas/UIPilot",
        },
        {
          title: "Jetpack Compose animations",
          stars: config.JETPACK_COMPOSE_ANIMATION_STARS,
          description: "Cool animations implemented with Jetpack compose",
          author: "Radhika S.",
          video: [jcAnimationsMp4, jcAnimationsWebm],
          link: "https://github.com/canopas/compose-animations-examples",
        },
      ],
    };
  },
  mounted() {
    if (!config.IS_PROD) {
      this.$watch(
        () => this.cData,
        (data) => {
          setGithubStars(this.contributions, data);
        },
      );
      if (this.cData != null) {
        setGithubStars(this.contributions, this.cData);
      }
    }
    lozad().observe(); // lazy loads videos with default selector as '.lozad'
  },
  inject: ["mixpanel"],
};
</script>
