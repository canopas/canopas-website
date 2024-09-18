<template>
  <section
    class="pb-28 mt-[1.563rem] from-white via-pink-5 to-white bg-gradient-[180.05deg]"
  >
    <div class="relative xll:container w-full h-full pb-16 lg:pt-12">
      <div class="container flex flex-col items-center mx-auto md:text-center">
        <div class="desk-header-2 xl:desk-header-1 text-black-87">
          <h1>
            Crafting code,<br />
            connecting communities
          </h1>
        </div>
        <div
          class="mt-2.5 xl:mt-5 mobile-header-3-regular text-black-60 w-[73%] 2xl:w-[54%]"
        >
          At Canopas, we believe in the power of open source software and the
          incredible community that makes it possible. We actively contribute to
          this vibrant ecosystem with the goal of advancing technology,
          fostering innovation, and empowering developers around the globe.
        </div>
        <nuxt-link
          class="mt-10 gradient-btn primary-btn z-10"
          to="/contact"
          @click.native="$mixpanel.track('tap_landing_cta')"
        >
          <span class="sub-h3-semibold lg:sub-h1-semibold"
            >Get Free Consultation
          </span>
        </nuxt-link>
      </div>
      <div
        class="absolute top-0 lg:top-[1.875rem] left-[10%] flex flex-col justify-between gap-y-12 lg:gap-y-20 2xl:gap-y-24 w-4/5"
      >
        <Icon
          class="-ml-3 w-5 lg:w-7 h-5 lg:h-7 animate-zoomIn"
          name="fa6-brands:github"
        />
        <div
          class="flex flex-row items-center gap-x-1.5 justify-center -ml-16 w-[4.063rem] h-[1.563rem] rounded-full drop-shadow-md text-center text-white from-orange-300 to-pink-300 bg-gradient-[270.11deg] animate-zoomOut"
        >
          <img
            :src="clapsmall"
            class="w-4 h-[1.125rem]"
            alt="desktop-background-icon"
            loading="lazy"
          />
          <span class="sub-h3-semibold">315</span>
        </div>
        <img
          :src="heartlike"
          class="ml-0.5 xl:ml-24 xl:ml-12 2xl:ml-40 xll:ml-24 -mt-[3.563rem] xl:-mt-[1.813rem] 2xl:-mt-[3.063rem] w-10 h-10 animate-zoomIn"
          alt="desktop-background-icon"
          loading="lazy"
        />
      </div>

      <div
        class="absolute top-[23rem] xl:top-[25rem] left-[10%] 2xl:left-[15%] flex flex-row justify-between w-4/5"
      >
        <div
          class="flex flex-row justify-between gap-x-[4.75rem] xl:gap-x-16 2xl:gap-x-16 xll:gap-x-4 gap-y-20"
        >
          <div
            class="flex flex-col items-center justify-center -mt-6 w-[8.75rem] xl:w-[11.25rem] h-[7.313rem] xl:h-[9.313rem] rounded-md bg-white text-center drop-shadow-md"
          >
            <span
              class="-mt-[1.188rem] xl:-mt-[1.438rem] mobile-header-3-bold text-primary-1"
            >
              100K+</span
            >
            <span
              class="sub-h1-semibold xl:mobile-header-2-semibold text-black-60 p-2.5"
              >Monthly minutes Read on blog
            </span>
          </div>
          <div
            class="flex justify-center items-center rounded-full bg-white drop-shadow-md lg:ml-5 2xl:ml-[4.25rem] mt-[6.25rem] w-[1.813rem] h-[1.938rem] animate-zoomIn"
          >
            <Icon
              class="w-3.5 h-3.5 p-px text-black-core"
              name="fa6-solid:thumbs-up"
            />
          </div>
        </div>
      </div>

      <div
        class="absolute top-[19rem] xl:top-[25rem] left-[46%] flex flex-row justify-between w-fit"
      >
        <div
          class="flex flex-row justify-between gap-x-11 xl:gap-x-28 2xll:gap-x-36"
        >
          <div
            class="mt-[8.5rem] xl:mt-[4.5rem] flex flex-col items-center justify-center -ml-4 xl:-ml-12 2xl:-ml-4 xll:-ml-8 w-[9.688rem] xl:w-[12.063rem] h-[6.25rem] xl:h-[8.125rem] rounded-md text-center bg-pink-80 lg:p-5"
          >
            <span class="mobile-header-3-bold text-primary-1"> 3K+</span>
            <span
              class="sub-h1-semibold xl:mobile-header-2-semibold text-black-60"
              >Daily Visitors on Our Blog
            </span>
          </div>

          <div
            class="flex justify-center items-center mt-12 xl:mt-8 ml-[5.25rem] xl:ml-[2.25rem] xll:ml-[-2.75rem] 2xl:mt-[1.75rem] w-7 h-7 rounded-full bg-white drop-shadow-md animate-zoomOut"
          >
            <Icon
              class="fas w-3.5 h-3.5 p-px footer-icon"
              name="fa6-solid:star"
            />
          </div>
          <img
            :src="dribble"
            class="ml- mt-48 xl:mt-36 w-5 lg:w-7 h-5 lg:h-7 animate-zoomIn"
            alt="desktop-background-icon"
            loading="lazy"
          />
        </div>
      </div>
      <div
        class="absolute left-[86%] 2xl:left-[79%] xll:left-[82%] top-[35%] xll:top-[42%] flex flex-row justify-between w-fit"
      >
        <div
          class="flex flex-col justify-between gap-y-36 lg:gap-y-40 xl:gap-y-48 2xl:gap-y-36"
        >
          <Icon
            class="-ml-6 w-5 lg:w-7 h-4 animate-zoomIn"
            name="fa6-brands:medium"
          />
          <div
            class="flex flex-col items-center justify-center w-[6.25rem] xl:w-[6.625rem] h-[5.625rem] xl:h-[6.75rem] rounded-md bg-white text-center drop-shadow-md"
          >
            <img
              :src="githubicon"
              class="absolute flex flex-col items-center justify-center w-[4.375rem] h-[4.375rem]"
              alt="background-githubicon"
              loading="lazy"
            />
            <span
              class="-mt-[1.2rem] xl:-mt-[0.2rem] mobile-header-3-bold text-primary-1"
            >
              1K+</span
            >
            <span
              class="sub-h1-semibold xl:mobile-header-2-semibold p-2.5 xl:p-[1.125rem] text-black-60"
              >GitHub Stars
            </span>
          </div>
        </div>
      </div>
      <div
        class="absolute left-[90%] top-[13%] flex flex-row justify-between w-fit"
      >
        <div
          class="flex flex-col justify-between gap-y-20 lg:gap-20 xl:gap-28 2xl:gap-y-36"
        >
          <img
            :src="clap"
            class="w-14 h-14 animate-zoomOut"
            alt="desktop-background-icon"
            loading="lazy"
          />
          <div
            class="flex flex-row gap-x-1.5 items-center justify-center -ml-[1.875rem] lg:ml-5 w-[4.063rem] h-[1.875rem] rounded-full drop-shadow-md from-orange-300 to-pink-300 bg-gradient-[270.11deg] text-center text-white animate-zoomIn"
          >
            <img
              :src="clapsmall"
              class="w-4 h-[1.125rem]"
              alt="desktop-background-icon"
              loading="lazy"
            />
            <span class="font-inter-semibold">237</span>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script type="module">
import heartlike from "@/assets/images/contributions/landing/heart_like.webp";
import clap from "@/assets/images/contributions/landing/clap.webp";
import clapsmall from "@/assets/images/contributions/landing/clap-small.webp";

import githubicon from "@/assets/images/contributions/landing/githubicon.webp";
import dribble from "@/assets/images/contributions/landing/dribble.webp";

export default {
  data() {
    return {
      heartlike,
      clap,
      clapsmall,
      dribble,
      githubicon,
    };
  },
  components: {},
};
</script>
