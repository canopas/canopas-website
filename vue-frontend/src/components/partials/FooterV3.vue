<template>
  <div class="tw-relative tw-overflow-hidden tw-px-0 tw-font-inter-medium">
    <img
      v-if="width < 600"
      :src="bg400"
      class="tw-absolute tw-w-full tw--z-[1]"
      alt="canopas-contact-footer"
    />
    <img
      :src="bg2400"
      :srcset="`${bg400} 400w, ${bg800} 800w, ${bg1200} 1200w, ${bg2400} 2400w`"
      class="tw-absolute tw-top-[15px] sm:tw-top-0 tw-left-0 tw-w-full tw-h-full xl2:tw-h-[unset] tw--z-[1] tw-object-cover xl2:tw-object-fill"
      alt="canopas-contact-footer"
    />
    <div class="tw-container tw-text-center">
      <div
        class="tw-mt-11 md:tw-mt-20 lg:tw-mt-32 md:tw-mb-20 canopas-gradient-text tw-text-[1.875rem] tw-leading-[2.25rem] md:tw-text-[3rem] md:tw-leading-[3.625rem] lg:tw-text-[4.0625rem] lg:tw-leading-[4.9375rem] tw-font-inter-bold"
      >
        Say Hello!
      </div>
      <ContactForm />
      <div class="tw-mb-8 md:tw-mb-[60px]">
        <div
          class="tw-mb-5 tw-text-[1.375rem] tw-leading-[1.6875rem] md:tw-text-[1.5rem] md:tw-leading-[1.8125rem] lg:tw-text-[1.625rem] lg:tw-leading-[1.9375rem] tw-text-white/[.87] tw-font-bold"
        >
          Follow us on
        </div>
        <ul
          class="tw-flex tw-flex-wrap tw-justify-center tw-w-full tw-m-auto tw-cursor-pointer tw-list-none"
        >
          <li
            v-for="icon in socialMediaIcons"
            :key="icon"
            class="tw-flex tw-justify-center tw-items-center tw-w-[40px] tw-h-[40px] md:tw-w-[62px] md:tw-h-[62px] tw-rounded-[50%] tw-text-center tw-cursor-pointer"
            :class="icon.image ? 'tw-m-[5px]' : 'gradient-border-btn'"
          >
            <a :href="icon.url" target="_blank" aria-label="footerLink">
              <img
                class="tw-h-full tw-w-full active:tw-scale-[0.98]"
                v-if="icon.image"
                :src="hover ? designRushHover : designRush"
                @mouseover="hover = true"
                @mouseleave="hover = false"
                alt="designrush-icon"
              />
              <font-awesome-icon
                v-else
                class="fab footer-icon tw-w-[20px] tw-h-[20px] md:tw-w-[32px] md:tw-h-[32px] tw-pt-[5px]"
                :icon="icon.icon"
              />
            </a>
          </li>
        </ul>
      </div>
      <div
        class="tw-mb-5 md:tw-mb-9 tw-text-[0.75rem] tw-leading-[0.9375rem] md:tw-text-[0.84375rem] md:tw-leading-[1.03125rem] lg:tw-text-[0.9375rem] lg:tw-leading-[1.125rem] tw-text-white/[.87]"
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
import designRushHover from "@/assets/images/footer/design-rush-gradient.svg";
import designRush from "@/assets/images/footer/design-rush-white.svg";
import bg400 from "@/assets/images/footer/bg/400.svg";
import bg800 from "@/assets/images/footer/bg/800.svg";
import bg1200 from "@/assets/images/footer/bg/1200.svg";
import bg2400 from "@/assets/images/footer/bg/2400.svg";
import Config from "@/config.js";
import ContactForm from "@/components/contact/ContactFormNew.vue";

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
      width: 680,
      bg400,
      bg800,
      bg1200,
      bg2400,
      designRush,
      designRushHover,
      hover: false,
      copyrightIcon: faCopyright,
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
    };
  },
  components: {
    FontAwesomeIcon,
    ContactForm,
  },
  mounted() {
    this.width = window.innerWidth;
  },
};
</script>
