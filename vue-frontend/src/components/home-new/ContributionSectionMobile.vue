<template>
  <div class="tw-container tw-justify-center tw-mt-20 tw-text-center">
    <div class="tw-m-auto tw-w-[290px] sm:tw-w-full tw-pb-14 md:tw-pb-20">
      <span
        class="tw-font-inter-bold tw-text-black-core/[.87] tw-text-[1.875rem] md:tw-text-[3.4375rem] tw-leading-[2.5rem] md:tw-leading-[5.5625rem]"
      >
        We are here to Contribute
      </span>
      <div
        class="tw-mt-2 md:tw-mt-0 tw-font-inter-medium tw-text-black-core/[0.6] tw-text-[1rem] tw-leading-[1.375rem] md:tw-text-[1.5rem] md:tw-leading-[2.5625rem]"
      >
        <span>and that starts with the site itself, </span>
        <br />
        <span>
          <a
            class="tw-relative tw-bg-gradient-[1deg] tw-pb-[5px] v2-canopas-gradient-text tw-font-inter-medium gradient-underline tw-no-underline after:tw-absolute after:tw-bottom-0 after:tw-left-0 after:tw-w-full after:tw-h-[2px] after:tw-bg-gradient-to-l after:tw-to-pink-300 after:tw-from-orange-300"
            :href="websiteOpenSourceUrl"
            target="_blank"
            @click.native="mixpanel.track('tap_canopas_website_github')"
            >canopas is open source!</a
          ></span
        >
      </div>
    </div>
    <div class="tw-text-left">
      <div
        v-for="(contribution, index) in contributions"
        :key="contribution"
        class="tw-relative tw-flex tw-justify-center tw-mt-8 tw-w-full"
        :class="index % 2 != 0 ? 'tw-flex-row-reverse' : 'tw-flex-row'"
      >
        <div class="tw-w-[45%] tw-shadow-[2px_2px_10px_rgba(0,0,0,0.1)]">
          <video class="lozad" autoplay loop muted playsinline>
            <source :data-src="contribution.video[1]" type="video/webm" />
            <source :data-src="contribution.video[0]" type="video/mp4" />
          </video>
        </div>

        <div
          @click="openLink(contribution)"
          class="tw-absolute tw-top-[50%] tw-flex tw-flex-col tw-w-[50%] tw-translate-y-[-50%] tw-bg-[#FFFFFF] tw-p-2.5 sm:tw-p-4 tw-shadow-[0_0px_10px_rgba(0,0,0,0.1)] active:tw-scale-[0.98]"
          :class="[index % 2 == 0 ? 'tw-left-[50%]' : 'tw-right-[50%]']"
        >
          <div>
            <div
              @touchstart.passive="activeIndex = index"
              :class="index % 2 != 0 ? 'tw-flex-row-reverse' : ''"
              class="tw-flex tw-flex-row tw-justify-between tw-items-center"
            >
              <div
                class="tw-flex tw-items-center tw-justify-center tw-w-[2.625rem] md:tw-w-[4.75rem] tw-h-[1.125rem] md:tw-h-[1.875rem] tw-from-[#ff835b] tw-to-[#f2709c] tw-bg-gradient-[270.11deg] tw-rounded-full md:tw-px-2.5 tw-font-inter-bold md:tw-font-inter-semibold tw-text-[0.625rem] md:tw-text-[1rem] tw-text-white"
              >
                <span>
                  <font-awesome-icon
                    class="fa tw-w-[10px] md:tw-w-[18px] tw-h-[10px] md:tw-h-[18px] tw-pr-[5px] tw-text-white"
                    icon="star"
                  />{{ contribution.stars }}</span
                >
              </div>
              <div>
                <div>
                  <a
                    class="tw-flex tw-items-center tw-font-inter-bold tw-text-[0.75rem] md:tw-text-[1.25rem] tw-leading-[2.5625rem] tw-duration-300 tw-ease-in-out hover:tw-scale-[1.1]"
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
              <p
                class="md:tw-mt-[41px] tw-font-inter-semibold tw-text-black-core/[0.87] tw-text-[1rem] md:tw-text-[1.875rem] tw-leading-[1.3rem] md:tw-leading-[2rem]"
              >
                {{ contribution.title }}
              </p>
            </div>
            <div class="tw-mt-5">
              <span
                class="tw-hidden md:tw-block tw-font-inter-light tw-text-[1.25rem] tw-leading-[1.578rem] tw-text-black-core/[0.6]"
                >{{ contribution.description }}</span
              >
            </div>
            <div
              @touchstart.passive="activeIndex = index"
              class="tw-flex tw-flex-row tw-justify-between tw-mt-[10px] md:tw-mt-[30px] tw-font-inter-regular tw-text-black-core/[0.87] tw-text-[0.625rem] md:tw-text-[1rem] tw-leading-[0.9375rem] md:tw-leading-[1.5rem]"
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

library.add(faGithub);

export default {
  computed: {
    ...mapState(useContributionStore, {
      cData: "data",
    }),
  },
  data() {
    return {
      icon: faGithub,
      rightArrow: faArrowRightLong,
      websiteOpenSourceUrl: config.WEBSITE_OPEN_SOURCE_URL,
      contributions: [
        {
          title: "Intro Showcase view in jetpack compose",
          stars: "271",
          author: "Radhika S.",
          video: [introShowCaseMp4, introShowCaseWebm],
          link: "https://github.com/canopas/Intro-showcase-view",
        },
        {
          title: "UIPilot",
          stars: "226",
          author: "Jimmy S.",
          video: [uIPilotMp4, uIPilotWebm],
          link: "https://github.com/canopas/UIPilot",
        },
        {
          title: "Jetpack Compose animations",
          stars: "202",
          author: "Radhika S.",
          video: [jcAnimationsMp4, jcAnimationsWebm],
          link: "https://github.com/canopas/Jetpack-compose-animations-examples",
        },
      ],
    };
  },
  computed: {
    ...mapState(useContributionStore, {
      cData: "data",
    }),
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
  methods: {
    openLink(contributions) {
      window.open(contributions.link, "_blank");
    },
  },
  components: {
    FontAwesomeIcon,
  },
  inject: ["mixpanel"],
};
</script>
