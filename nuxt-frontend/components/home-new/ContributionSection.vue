<template>
  <div
    class="tw-relative xll:tw-container tw-my-0 tw-mx-auto tw-w-[90%] xll:tw-max-w-[1500px] tw-p-0"
  >
    <div class="tw-sticky tw-top-0 tw-w-full tw-h-screen">
      <div
        class="tw-flex tw-flex-col tw-justify-center tw-w-[30%] tw-h-full tw-float-right"
      >
        <span
          class="tw-font-inter-bold tw-text-black-core/[.87] tw-text-[1.875rem] md:tw-text-[3.4375rem] tw-leading-10 md:tw-leading-[5.5625rem]"
        >
          We are here to Contribute
        </span>
        <div
          class="tw-mt-2 md:tw-mt-0 tw-font-inter-medium tw-text-black-core/[0.6] tw-text-[1.5rem] tw-leading-[2.5625rem]"
        >
          <span>and that starts with the site itself, </span>
          <br />
          <a
            class="tw-relative tw-bg-gradient-[1deg] tw-pb-[5px] v2-canopas-gradient-text tw-font-inter-medium gradient-underline tw-no-underline tw-cursor-pointer after:tw-absolute after:tw-bottom-0 after:tw-left-0 after:tw-w-full after:tw-h-0.5 after:tw-bg-gradient-to-l after:tw-to-pink-300 after:tw-from-orange-300"
            :href="websiteOpenSourceUrl"
            target="_blank"
            @click.native="mixpanel.track('tap_canopas_website_github')"
            >canopas is open source!</a
          >
        </div>
      </div>
    </div>
    <div
      class="tw-flex tw-flex-col tw--mt-[95vh] tw-mb-0 tw-ml-8 xl:tw-ml-32 xll:tw-ml-0 tw-w-[55%] xll:tw-w-[65%]"
    >
      <div
        v-for="(contribution, index) in contributions"
        :key="contribution"
        class="tw-relative tw-flex tw-flex-row tw-mt-8 tw-mb-6 tw-cursor-pointer"
        :class="index % 2 != 0 ? 'tw-flex-row-reverse' : ''"
      >
        <div class="tw-shadow-[2px_2px_10px_rgba(0,0,0,0.1)]">
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
          class="tw-absolute xl:tw-relative lg:tw-top-[50%] tw-flex tw-flex-col tw-basis-[40%] lg:tw-basis-[43%] xl:tw-basis-[50%] 2xl:tw-basis-[40%] tw-my-auto lg:tw-w-[60%] xl:tw-w-[100%] lg:tw-translate-y-[-50%] xl:tw-translate-y-0 tw-bg-[#FFFFFF] tw-p-4 xl:tw-p-8 tw-shadow-[0_0px_10px_rgba(0,0,0,0.1)] active:tw-scale-[0.98]"
          :class="[
            index % 2 == 0
              ? 'lg:tw-left-[40%] xl:tw-left-[unset]'
              : 'lg:tw-right-[40%] xl:tw-right-[unset]',
          ]"
        >
          <div>
            <div
              @touchstart.passive="activeIndex = index"
              :class="index % 2 != 0 ? 'tw-flex-row-reverse' : ''"
              class="tw-flex tw-flex-row tw-justify-between tw-items-center"
            >
              <div
                class="tw-flex tw-items-center tw-justify-center tw-w-[4.75rem] tw-h-[1.875rem] tw-rounded-full tw-bg-gradient-[270.11deg] tw-from-[#ff835b] tw-to-[#f2709c] tw-px-2.5 tw-font-inter-semibold tw-text-[1rem] tw-text-white"
              >
                <!-- <span>
                  <font-awesome-icon
                    class="fa tw-w-[18px] tw-h-[18px] tw-pr-[5px] tw-text-white"
                    icon="star"
                  />{{ contribution.stars }}</span
                > -->
              </div>
              <div>
                <div>
                  <a
                    class="tw-flex tw-items-center tw-font-inter-bold tw-text-[1.25rem] tw-leading-[2.5625rem] tw-duration-300 tw-ease-in-out hover:tw-scale-[1.1]"
                    :href="contribution.link"
                    target="_blank"
                    @click.native="mixpanel.track('tap_library')"
                  >
                    <font-awesome-icon
                      class="tw-mr-2.5 tw-w-[1.75rem] tw-h-[1.75rem] tw-text-pink-300"
                      :icon="icon"
                    />
                    <div class="v2-canopas-gradient-text">GitHub</div>
                  </a>
                </div>
              </div>
            </div>
            <div>
              <p
                class="tw-mt-[41px] tw-font-inter-semibold tw-text-black-core/[0.87] tw-text-[1.875rem] tw-leading-8"
              >
                {{ contribution.title }}
              </p>
            </div>
            <div class="tw-mt-5">
              <span
                class="tw-font-inter-light tw-text-black-core/[0.6] tw-text-[1.25rem] tw-leading-[1.578rem]"
                >{{ contribution.description }}</span
              >
            </div>
            <div
              @touchstart.passive="activeIndex = index"
              class="tw-flex tw-flex-row tw-justify-between tw-mt-[30px] tw-font-inter-regular tw-text-black-core/[0.87] tw-text-base"
            >
              <span>by {{ contribution.author }}</span>
              <a
                class="tw-duration-300 tw-ease-in-out hover:tw-scale-[1.1]"
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
import uIPilotMp4 from "@/assets/images/contribution/animations/UIPilot.mp4";
import uIPilotWebm from "@/assets/images/contribution/animations/UIPilot.webm";
import introShowCaseMp4 from "@/assets/images/contribution/animations/introShowCase.mp4";
import introShowCaseWebm from "@/assets/images/contribution/animations/introShowCase.webm";
import jcAnimationsMp4 from "@/assets/images/contribution/animations/JetpackComposeAnimations.mp4";
import jcAnimationsWebm from "@/assets/images/contribution/animations/JetpackComposeAnimations.webm";
import config from "@/config.js";
import lozad from "lozad";
import { setGithubStars } from "@/utils.js";
import { mapState } from "pinia";
import { useContributionStore } from "@/stores/contribution";
import { openBlog } from "@/utils.js";

library.add(faGithub);

export default {
  computed: {
    ...mapState(useContributionStore, {
      cData: "data",
    }),
  },
  data() {
    return {
      openBlog,
      icon: faGithub,
      rightArrow: faArrowRightLong,
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
        }
      );
      if (this.cData != null) {
        setGithubStars(this.contributions, this.cData);
      }
    }

    lozad().observe(); // lazy loads videos with default selector as '.lozad'
  },
  components: {
    FontAwesomeIcon,
  },
  // inject: ["mixpanel"],
};
</script>
