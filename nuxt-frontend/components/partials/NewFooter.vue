<template>
  <div class="relative overflow-hidden px-0 font-inter-medium z-[1]">
    <img
      :src="bg"
      class="absolute top-[15px] sm:top-0 left-0 w-full h-full xl2:h-[unset] -z-[1] object-cover xl2:object-fill"
      alt="canopas-contact-footer"
    />
    <div class="container text-center">
      <div class="mt-11 mb-8 md:mb-[60px]">
        <div
          class="mt-5 text-[1.375rem] leading-[1.6875rem] md:text-[1.5rem] md:leading-[1.8125rem] lg:text-[1.625rem] lg:leading-[1.9375rem] text-white/[.87] font-bold"
        >
          Follow us on
        </div>
        <ul
          class="flex flex-wrap justify-center w-full m-auto mt-2 xl:mt-6 cursor-pointer list-none"
        >
          <li
            v-for="icon in socialMediaIcons"
            :key="icon"
            class="flex justify-center !items-center w-10 h-10 md:w-[62px] md:h-[62px] !rounded-[50%] text-center cursor-pointer"
            :class="icon.image ? 'm-[5px]' : 'gradient-border-btn border-0'"
          >
            <a :href="icon.url" target="_blank" aria-label="footerLink">
              <Icon
                class="fab footer-icon w-5 h-5 md:w-8 md:h-8 box-content"
                :name="icon.icon"
              />
            </a>
          </li>
        </ul>
      </div>
      <div
        class="mb-8 md:mb-9 text-[0.75rem] leading-[0.9375rem] md:text-[0.84375rem] md:leading-[1.03125rem] lg:text-[0.9375rem] lg:leading-[1.125rem] text-white/[.87]"
      >
        <Icon name="fa6-regular:copyright" class="box-content" />
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

export default {
  data() {
    return {
      bg,
      hover: false,
      socialMediaIcons: [
        {
          url: Config.FACEBOOK_URL,
          icon: "fa6-brands:facebook-f",
          event: "tap_footer_facebook",
        },
        {
          url: Config.INSTAGRAM_URL,
          icon: "fa6-brands:instagram",
          event: "tap_footer_instagram",
        },
        {
          url: Config.TWITTER_URL,
          icon: "fa6-brands:x-twitter",
          event: "tap_footer_twitter",
        },
        {
          url: Config.BLOG_URL,
          icon: "fa6-brands:medium",
          event: "tap_footer_medium",
        },
        {
          url: Config.LINKEDIN_URL,
          icon: "fa6-brands:linkedin-in",
          event: "tap_footer_linkedin",
        },
        {
          url: Config.YOUTUBE_URL,
          icon: "fa6-brands:youtube",
          event: "tap_footer_youtube",
        },
      ],
    };
  },
};
</script>
