<template>
  <div class="relative container my-0 lg:mb-60 mx-auto">
    <div class="sticky top-0 w-full h-screen">
      <div
        class="flex flex-col justify-center w-[36%] xl:w-[39%] h-full float-right xl:pl-12"
      >
        <span
          class="text-black-87 mobile-header-2 lg:mobile-header-1 xl:desk-header-2"
        >
          We are here to Contribute
        </span>
        <div class="mt-2 md:mt-0 text-black-60 mobile-header-2-regular">
          <span>and that starts with the site itself, </span>
          <a
            class="mobile-header-2-semibold relative bg-gradient-[1deg] pb-[5px] v2-canopas-gradient-text gradient-underline no-underline cursor-pointer after:absolute after:bottom-0 after:left-0 after:w-full after:h-0.5 after:bg-gradient-to-l after:to-pink-300 after:from-orange-300"
            :href="websiteOpenSourceUrl"
            target="_blank"
            @click.native="$mixpanel.track('tap_canopas_website_github')"
            >canopas is open source!</a
          >
        </div>
      </div>
    </div>
    <div class="flex flex-col -mt-[90vh] mb-0 w-[62%] xl:w-[60%]">
      <div
        v-for="(contribution, index) in contributions"
        :key="contribution"
        class="relative flex flex-row mt-8 mb-6 cursor-pointer"
        :class="index % 2 != 0 ? 'flex-row-reverse' : ''"
      >
        <video
          class="lozad rounded-xl z-[1] shadow-[2px_2px_10px_rgba(0,0,0,0.1)]"
          autoplay
          loop
          muted
          playsinline
          style="height: 630px"
        >
          <source :data-src="contribution.video[1]" type="video/webm" />
          <source :data-src="contribution.video[0]" type="video/mp4" />
        </video>

        <div
          @click="
            openBlog(contribution.link, 'tap_home_github_contribution_section')
          "
          class="relative flex flex-col basis-[50%] xl:basis-[60%] my-auto w-[100%] translate-y-0 bg-[#FFFFFF] p-4 xl:p-8 shadow-[0_0px_10px_rgba(0,0,0,0.1)] active:scale-[0.98]"
          :class="[index % 2 == 0 ? 'rounded-r-xl' : 'rounded-l-xl']"
        >
          <div>
            <div
              :class="index % 2 != 0 ? 'flex-row-reverse' : ''"
              class="flex flex-row justify-between items-center"
            >
              <div
                class="flex items-center xl:sub-h1-semibold sub-h4-medium justify-center w-[3.9375rem] h-[2.125rem] xl:w-[5.0625rem] xl:h-[2.25rem] rounded bg-gradient-[270.11deg] from-[#ff835b] to-[#f2709c] xl:px-2.5 text-white"
              >
                <span>
                  <Icon
                    class="fa w-[18px] h-[18px] pr-[5px] text-white box-content mb-[0.225em]"
                    name="fa6-solid:star"
                  />{{ contribution.stars }}</span
                >
              </div>
              <div>
                <div>
                  <a
                    class="flex items-center lg:sub-h4-medium xl:mobile-header-3-semibold duration-300 ease-in-out hover:scale-[1.1]"
                    :href="contribution.link"
                    target="_blank"
                    @click.native="$mixpanel.track('tap_library')"
                  >
                    <Icon
                      class="mr-2.5 w-4 h-4 xl:w-7 xl:h-7 footer-icon"
                      name="fa6-brands:github"
                    />
                    <div class="v2-canopas-gradient-text">GitHub</div>
                  </a>
                </div>
              </div>
            </div>
            <div>
              <p class="mt-[41px] text-black-87 mobile-header-2-semibold">
                {{ contribution.title }}
              </p>
            </div>
            <div class="mt-5">
              <span
                class="text-black-60 sub-h3-regular xl:mobile-header-3-regular"
                >{{ contribution.description }}</span
              >
            </div>
            <div
              class="flex flex-row justify-between mt-[30px] text-black-87 sub-h4-regular xl:sub-h1-regular"
            >
              <span>by {{ contribution.author }}</span>
              <a
                class="duration-300 ease-in-out hover:scale-[1.1]"
                :href="contribution.link"
                :aria-label="contribution.title"
                target="_blank"
                @click.native="$mixpanel.track('tap_library')"
              >
                <Icon
                  class="arrow fa w-5 h-5 box-content"
                  name="fa6-solid:arrow-right-long"
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
          link: "https://github.com/canopas/Jetpack-compose-animations-examples",
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
