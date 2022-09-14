<template>
  <div
    class="tw-m-0 tw-p-0 tw-relative tw-flex tw-flex-col tw-items-center tw-justify-between"
  >
    <img
      src="@/assets/images/footer/Canopas-footer-img.svg"
      class="tw-w-full tw-h-full sm:tw-object-cover tw-oject-fill tw-absolute tw-top-0 tw-left-0 -tw-z-[5]"
      alt="footer-background-image"
    />
    <div
      class="tw-container tw-text-center tw-flex tw-flex-col tw-items-center"
    >
      <div class="tw-mb-16 xl:tw-mb-24" v-if="isJobsUrl">
        <a :href="glassdoorLink">
          <img
            src="@/assets/images/footer/glassdoor.webp"
            class="tw-w-4/5 sm:tw-w-4/5 tw-h-full"
            alt="glassdoor-image"
          />
        </a>
      </div>
      <div class="tw-mb-16 xl:tw-mb-24">
        <div
          class="normal-text tw-font-bold tw-text-[1.375rem] tw-leading-[1.75rem] lg:tw-text-[1.75rem] lg:tw-leading-[2.125rem] tw-mb-6"
        >
          Follow us on
        </div>

        <ul
          class="tw-my-0 tw-mx-1 tw-w-full tw-h-full tw-cursor-pointer tw-list-none tw-flex tw-flex-wrap sm:tw-flex-nowrap"
        >
          <li
            class="sm:tw-w-[50px] sm:tw-h-[50px] sm:tw-mb-0"
            v-for="socialMedia in socialMediaIcons"
            :key="socialMedia"
            :class="
              socialMedia.image
                ? ' tw-mt-[20px] sm:tw-my-0 tw-mx-1 tw-w-[50px] tw-h-[50px] tw-cursor-pointer'
                : ' gradient-btn sm:tw-my-0  tw-mx-1 tw-w-[50px] tw-h-[50px] tw-cursor-pointer tw-rounded-[50%] tw-p-[0.7rem] '
            "
          >
            <a :href="socialMedia.url" target="_blank">
              <img
                class="active:tw-scale-[0.98] tw-h-full tw-w-full"
                v-if="socialMedia.image"
                :src="hover ? designRushHover : designRush"
                @mouseover="hover = true"
                @mouseleave="hover = false"
              />
              <font-awesome-icon
                v-else
                class="fab tw-h-full tw-w-full"
                :icon="socialMedia.icon"
              />
            </a>
          </li>
        </ul>
      </div>
      <div class="tw-mb-16 xl:tw-mb-24">
        <p class="normal-2-text">
          <font-awesome-icon :icon="copyrightIcon" /> &nbsp;
          <span
            >{{ new Date().getFullYear() }} Canopas Software LLP. All rights
            reserved.</span
          >
        </p>
      </div>
    </div>
  </div>
</template>

<script type="module">
import { faCopyright } from "@fortawesome/free-regular-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import designRush from "@/assets/images/footer/design-rush.svg";
import designRushHover from "@/assets/images/footer/design-rush-hover.svg";

import {
  faFacebookF,
  faInstagram,
  faTwitter,
  faMediumM,
  faLinkedinIn,
  faYoutube,
} from "@fortawesome/free-brands-svg-icons";

import Config from "@/config.js";

export default {
  data() {
    return {
      glassdoorLink:
        "https://www.glassdoor.co.in/Overview/Working-at-Canopas-EI_IE3194462.11,18.htm",
      copyrightIcon: faCopyright,
      designRush: designRush,
      designRushHover: designRushHover,
      hover: false,
      socialMediaIcons: [
        {
          url: Config.FACEBOOK_URL,
          icon: faFacebookF,
        },
        {
          url: Config.INSTAGRAM_URL,
          icon: faInstagram,
        },
        {
          url: Config.TWITTER_URL,
          icon: faTwitter,
        },
        {
          url: Config.BLOG_URL,
          icon: faMediumM,
        },
        {
          url: Config.LINKEDIN_URL,
          icon: faLinkedinIn,
        },
        {
          url: Config.YOUTUBE_URL,
          icon: faYoutube,
        },
        {
          url: Config.DESIGNRUSH_URL,
          image: true,
        },
      ],
      isJobsUrl: this.$router.currentRoute.value.fullPath.indexOf("/jobs") > -1,
    };
  },
  components: {
    FontAwesomeIcon,
  },
};
</script>
