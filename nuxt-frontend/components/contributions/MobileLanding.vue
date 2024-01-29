<template>
  <section
    class="mb-16 h-[485px] bg-gradient-to-t from-orange-1 to-pink-1 bg-gradient-[270.11deg]"
  >
    <div class="relative">
      <div
        class="absolute left-[16%] top-4 flex w-4/5 flex-row justify-between"
      >
        <img
          :src="heartlike"
          class="mx-auto -mt-3 h-7 w-7 animate-zoomIn"
          alt="mobile-background-icon"
          loading="lazy"
        />

        <div
          class="flex h-8 w-[1.813rem] animate-zoomIn items-center justify-center gap-y-12 rounded-full bg-white drop-shadow-md lg:-mt-3 lg:ml-5 2xl:ml-[8.25rem] 2xl:mt-5"
        >
          <Icon
            class="h-3.5 w-3.5 p-px text-black-core"
            name="fa6-solid:thumbs-up"
          />
        </div>
      </div>
      <div
        class="absolute left-[6%] top-[11rem] xs:top-[11.5rem] mx-auto flex w-[90%] flex-row justify-between xs:top-44"
      >
        <div class="flex flex-col items-start justify-around mt-10">
          <div
            class="flex h-7 w-7 animate-zoomIn items-center justify-center rounded-full bg-white drop-shadow-md"
          >
            <Icon
              class="fas footer-icon h-3.5 w-3.5 p-px"
              name="fa6-solid:star"
            />
          </div>
          <div
            class="mt-[4.25rem] flex h-[3.4rem] w-[7.188rem] flex-col items-center justify-center rounded-md bg-white text-center drop-shadow-md"
          >
            <span class="text-primary-1 sub-h3-bold -mt-5"> 100K+</span>
            <span class="sub-h4-medium text-black-87"
              >Monthly minutes <br />
              Read on blog
            </span>
          </div>
          <Icon class="h-4 w-4 animate-zoomIn ml-5" name="fa6-brands:github" />
        </div>
        <div
          class="flex flex-col items-center justify-start gap-[3.75rem] mt-44"
        >
          <Icon
            class="mt-[-2.188rem] ml-[-3.438rem] h-2.5 w-[1.063rem] animate-zoomOut"
            name="fa6-brands:medium"
          />
          <div
            class="-mt-2 flex h-20 w-[6.25rem] flex-col items-center justify-center rounded-md text-center bg-pink-80"
          >
            <span class="text-primary-1 sub-h3-bold"> 3K+</span>
            <span class="mb-2.5 sub-h4-medium text-black-87"
              >Daily Visitors <br />
              on Our Blog
            </span>
          </div>
        </div>

        <div class="mt-10 flex flex-col items-center">
          <div
            class="flex h-[1.375rem] w-[3.25rem] animate-zoomIn flex-row items-center justify-center gap-x-0.5 rounded-full from-orange-300 to-pink-300 text-center text-white drop-shadow-md bg-gradient-[270.11deg] xs:ml-10"
          >
            <img
              :src="clapsmall"
              class="h-3.5 w-3"
              alt="desktop-background-icon"
              loading="lazy"
            />
            <span class="sub-h4-semibold">237</span>
          </div>

          <div
            class="mt-20 flex h-[4.5rem] w-20 flex-col items-center justify-center rounded-md bg-white text-center drop-shadow-md"
          >
            <img
              :src="githubicon"
              class="absolute mt-1 flex h-[3.125rem] w-[3.125rem] flex-col items-center justify-center"
              alt="mobile-background-icon"
              loading="lazy"
            />
            <span class="-mt-[0.938rem] text-primary-1 sub-h3-bold"> 1K+</span>
            <span class="p-2.5 sub-h4-medium text-black-87">GitHub Stars </span>
          </div>
          <img
            :src="dribble"
            class="mt-6 h-4 w-4 animate-zoomOut"
            alt="mobile-background-icon"
            loading="lazy"
          />
        </div>
      </div>
      <div class="container flex flex-col pt-11 text-center">
        <div class="mobile-header-1 text-black-87">
          <span
            >Crafting Code, <br />
            Connecting Communities</span
          >
        </div>
        <div class="mt-4 sub-h1-regular text-black-60">
          Our team at Canopas constantly seeks to <br />
          learn, grow, and give back to the developer <br />
          community.
        </div>
        <nuxt-link
          class="gradient-btn primary-btn"
          to="/contact"
          @click.native="$mixpanel.track('tap_landing_cta')"
        >
          <span class="sub-h3-semibold">Get Free Consultation </span>
        </nuxt-link>
      </div>
    </div>
  </section>
</template>

<script type="module">
import heartlike from "@/assets/images/contributions/landing/mobile_heart_like.webp";
import githubicon from "@/assets/images/contributions/landing/githubicon.webp";
import clapsmall from "@/assets/images/contributions/landing/clap-small.webp";

import dribble from "@/assets/images/contributions/landing/dribble.webp";

export default {
  data() {
    return {
      heartlike,
      dribble,
      clapsmall,
      githubicon,
    };
  },
  components: {},
};
</script>
