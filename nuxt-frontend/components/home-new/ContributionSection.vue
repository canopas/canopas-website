<template>
  <div
    class="relative xll:container my-0 mx-auto w-[90%] xll:max-w-[1500px] p-0"
  >
    <div class="sticky top-0 w-full h-screen">
      <div class="flex flex-col justify-center w-[30%] h-full float-right">
        <span
          class="font-inter-bold text-black-core/[.87] text-[1.875rem] md:text-[3.4375rem] leading-10 md:leading-[5.5625rem]"
        >
          We are here to Contribute
        </span>
        <div
          class="mt-2 md:mt-0 font-inter-medium text-black-core/[0.6] text-[1.5rem] leading-[2.5625rem]"
        >
          <span>and that starts with the site itself, </span>
          <br />
          <a
            class="relative bg-gradient-[1deg] pb-[5px] v2-canopas-gradient-text font-inter-medium gradient-underline no-underline cursor-pointer after:absolute after:bottom-0 after:left-0 after:w-full after:h-0.5 after:bg-gradient-to-l after:to-pink-300 after:from-orange-300"
            :href="websiteOpenSourceUrl"
            target="_blank"
            @click.native="$mixpanel.track('tap_canopas_website_github')"
            >canopas is open source!</a
          >
        </div>
      </div>
    </div>
    <div
      class="flex flex-col -mt-[95vh] mb-0 ml-8 xl:ml-32 xll:ml-0 w-[55%] xll:w-[65%]"
    >
      <div
        v-for="(contribution, index) in contributions"
        :key="contribution"
        class="relative flex flex-row mt-8 mb-6 cursor-pointer"
        :class="index % 2 != 0 ? 'flex-row-reverse' : ''"
      >
        <div class="shadow-[2px_2px_10px_rgba(0,0,0,0.1)]">
          <video
            class="lozad"
            autoplay
            loop
            muted
            playsinline
            style="height: 690px"
          >
            <source :data-src="contribution.video[1]" type="video/webm" />
            <source :data-src="contribution.video[0]" type="video/mp4" />
          </video>
        </div>

        <div
          @click="
            openBlog(contribution.link, 'tap_home_github_contribution_section')
          "
          class="absolute xl:relative lg:top-[50%] flex flex-col basis-[40%] lg:basis-[43%] xl:basis-[50%] 2xl:basis-[40%] my-auto lg:w-[60%] xl:w-[100%] lg:translate-y-[-50%] xl:translate-y-0 bg-[#FFFFFF] p-4 xl:p-8 shadow-[0_0px_10px_rgba(0,0,0,0.1)] active:scale-[0.98]"
          :class="[
            index % 2 == 0
              ? 'lg:left-[40%] xl:left-[unset]'
              : 'lg:right-[40%] xl:right-[unset]',
          ]"
        >
          <div>
            <div
              :class="index % 2 != 0 ? 'flex-row-reverse' : ''"
              class="flex flex-row justify-between items-center"
            >
              <div
                class="flex items-center justify-center w-[4.75rem] h-[1.875rem] rounded-full bg-gradient-[270.11deg] from-[#ff835b] to-[#f2709c] px-2.5 font-inter-semibold text-[1rem] text-white"
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
                    class="flex items-center font-inter-bold text-[1.25rem] leading-[2.5625rem] duration-300 ease-in-out hover:scale-[1.1]"
                    :href="contribution.link"
                    target="_blank"
                    @click.native="$mixpanel.track('tap_library')"
                  >
                    <Icon
                      class="mr-2.5 w-[1.75rem] h-[1.75rem] text-pink-300"
                      name="fa6-brands:github"
                    />
                    <div class="v2-canopas-gradient-text">GitHub</div>
                  </a>
                </div>
              </div>
            </div>
            <div>
              <p
                class="mt-[41px] font-inter-semibold text-black-core/[0.87] text-[1.875rem] leading-8"
              >
                {{ contribution.title }}
              </p>
            </div>
            <div class="mt-5">
              <span
                class="font-inter-light text-black-core/[0.6] text-[1.25rem] leading-[1.578rem]"
                >{{ contribution.description }}</span
              >
            </div>
            <div
              class="flex flex-row justify-between mt-[30px] font-inter-regular text-black-core/[0.87] text-base"
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
                  class="arrow fa w-4 h-4 box-content"
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
          stars: "275",
          description:
            "Highlight different features of the app using Jetpack Compose",
          author: "Radhika S.",
          video: [introShowCaseMp4, introShowCaseWebm],
          link: "https://github.com/canopas/Intro-showcase-view",
        },
        {
          title: "UIPilot",
          stars: "228",
          description: "The missing typesafe SwiftUI navigation library",
          author: "Jimmy S.",
          video: [uIPilotMp4, uIPilotWebm],
          link: "https://github.com/canopas/UIPilot",
        },
        {
          title: "Jetpack Compose animations",
          stars: "203",
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
