<template>
  <div class="pt-14 px-0 font-inter-medium z-[1] bg-deep-charcoal">
    <div class="container text-center">
      <div class="mb-8 md:mb-[60px]">
        <div class="text-white mobile-header-3 lg:desk-header-3">
          Follow us on
        </div>
        <ul
          class="flex flex-wrap justify-center w-full m-auto mt-2 xl:mt-6 cursor-pointer list-none pl-0"
        >
          <li
            v-for="icon in socialMediaIcons"
            :key="icon"
            class="flex justify-center !items-center w-10 h-10 md:w-[62px] md:h-[62px] !rounded-[50%] text-center cursor-pointer pb-[1.2rem] md:pb-4"
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
        class="mb-8 md:mb-9 text-[0.75rem] leading-[0.9375rem] md:text-[0.84375rem] md:leading-[1.03125rem] lg:text-[0.9375rem] lg:leading-[1.125rem] text-white-core-80"
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

<script setup>
import config from "@/config.js";

const socialMediaIcons = [
  {
    url: config.FACEBOOK_URL,
    icon: "fa6-brands:facebook-f",
    event: "tap_footer_facebook",
  },
  {
    url: config.INSTAGRAM_URL,
    icon: "fa6-brands:instagram",
    event: "tap_footer_instagram",
  },
  {
    url: config.TWITTER_URL,
    icon: "fa6-brands:x-twitter",
    event: "tap_footer_twitter",
  },
  {
    url: config.LINKEDIN_URL,
    icon: "fa6-brands:linkedin-in",
    event: "tap_footer_linkedin",
  },
  {
    url: config.YOUTUBE_URL,
    icon: "fa6-brands:youtube",
    event: "tap_footer_youtube",
  },
  {
    url: config.GITHUB_URL,
    icon: "fa6-brands:github",
    event: "tap_footer_github",
  },
];
</script>
