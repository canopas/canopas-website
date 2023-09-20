<template>
  <div
    class="tw-relative tw-overflow-hidden tw-px-0 tw-font-inter-medium tw-z-[1]"
  >
    <img
      :src="bg"
      class="tw-absolute tw-top-[15px] sm:tw-top-0 tw-left-0 tw-w-full tw-h-full xl2:tw-h-[unset] tw--z-[1] tw-object-cover xl2:tw-object-fill"
      alt="canopas-contact-footer"
    />
    <div class="tw-container tw-text-center">
      <div class="tw-mt-11 tw-mb-8 md:tw-mb-[60px]">
        <div
          class="tw-mt-5 tw-text-[1.375rem] tw-leading-[1.6875rem] md:tw-text-[1.5rem] md:tw-leading-[1.8125rem] lg:tw-text-[1.625rem] lg:tw-leading-[1.9375rem] tw-text-white/[.87] tw-font-bold"
        >
          Follow us on
        </div>
        <ul
          class="tw-flex tw-flex-wrap tw-justify-center tw-w-full tw-m-auto tw-mt-2 xl:tw-mt-6 tw-cursor-pointer tw-list-none"
        >
          <li
            v-for="icon in socialMediaIcons"
            :key="icon"
            class="tw-flex tw-justify-center !tw-items-center tw-w-10 tw-h-10 md:tw-w-[62px] md:tw-h-[62px] !tw-rounded-[50%] tw-text-center tw-cursor-pointer"
            :class="
              icon.image ? 'tw-m-[5px]' : 'gradient-border-btn tw-border-0'
            "
          >
            <a :href="icon.url" target="_blank" aria-label="footerLink">
              <font-awesome-icon
                class="fab footer-icon tw-w-5 tw-h-5 md:tw-w-8 md:tw-h-8 tw-pt-[5px]"
                :icon="icon.icon"
              />
            </a>
          </li>
        </ul>
      </div>
      <div
        class="tw-mb-8 md:tw-mb-9 tw-text-[0.75rem] tw-leading-[0.9375rem] md:tw-text-[0.84375rem] md:tw-leading-[1.03125rem] lg:tw-text-[0.9375rem] lg:tw-leading-[1.125rem] tw-text-white/[.87]"
      >
        <font-awesome-icon :icon="copyrightIcon" />
        {{ new Date().getFullYear() }} Canopas Software LLP. All rights
        reserved.
      </div>
    </div>
    <svg width="0" height="0">
      <linearGradient id="lgrad" x1="100%" y1="100%" x2="0%" y2="0%">
        <stop offset="-24.42%" style="stop-color: #ff835b; stop-opacity: 1" />
        <stop offset="101.76%" style="stop-color: #f2709c; stop-opacity: 1" />
      </linearGradient>
    </svg>
  </div>
</template>

<script type="module">
import bg from "@/assets/images/footer/new-bg.svg";
import Config from "@/config.js";

import {
  faFacebookF,
  faInstagram,
  faTwitter,
  faMediumM,
  faLinkedinIn,
  faYoutube,
} from "@fortawesome/free-brands-svg-icons";
import { faCopyright } from "@fortawesome/free-regular-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

export default {
  data() {
    return {
      bg,
      hover: false,
      copyrightIcon: faCopyright,
      socialMediaIcons: [
        {
          url: Config.FACEBOOK_URL,
          icon: faFacebookF,
          event: "tap_footer_facebook",
        },
        {
          url: Config.INSTAGRAM_URL,
          icon: faInstagram,
          event: "tap_footer_instagram",
        },
        {
          url: Config.TWITTER_URL,
          icon: faTwitter,
          event: "tap_footer_twitter",
        },
        {
          url: Config.BLOG_URL,
          icon: faMediumM,
          event: "tap_footer_medium",
        },
        {
          url: Config.LINKEDIN_URL,
          icon: faLinkedinIn,
          event: "tap_footer_linkedin",
        },
        {
          url: Config.YOUTUBE_URL,
          icon: faYoutube,
          event: "tap_footer_youtube",
        },
      ],
    };
  },
  // inject: ["mixpanel"],
  components: {
    FontAwesomeIcon,
  },
};
</script>
