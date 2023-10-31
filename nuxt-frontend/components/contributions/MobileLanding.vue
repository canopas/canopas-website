<template>
  <section
    class="mb-16 h-[470px] bg-gradient-to-t from-[#FF835B]/[0.1] to-[#F2709C]/[0.1] bg-gradient-[270.11deg]"
  >
    <div class="relative">
      <div
        class="absolute left-[16%] top-4 flex w-[80%] flex-row justify-between"
      >
        <img
          :src="heartlike"
          class="mx-auto -mt-3 h-7 w-7 animate-zoomIn"
          alt="mobile-background-icon"
        />

        <div
          class="flex h-[31px] w-[29px] animate-zoomIn items-center justify-center gap-y-12 rounded-full bg-white drop-shadow-md lg:-mt-3 lg:ml-5 2xl:ml-[8.25rem] 2xl:mt-5"
        >
          <Icon
            class="h-3.5 w-3.5 p-px text-[#000]"
            name="fa6-solid:thumbs-up"
          />
        </div>
      </div>
      <div
        class="absolute left-[6%] top-[13.5rem] mx-auto flex w-[90%] flex-row justify-between xs:top-44"
      >
        <div
          class="mt-[45px] flex flex-col items-center justify-around gap-y-16 xs:gap-y-20"
        >
          <div
            class="-mt-6 mt-5 flex h-20 w-[140px] flex-col items-center justify-center gap-y-1.5 rounded-md bg-white text-center drop-shadow-md"
          >
            <span
              class="v2-canopas-gradient-text -mt-[30px] font-inter-bold text-[1.5rem] font-extrabold leading-[2.1rem]"
            >
              100K+</span
            >
            <span class="font-inter-medium text-base text-black-core/[0.87]"
              >Monthly minutes Read on blog
            </span>
          </div>
          <Icon class="h-4 w-4 animate-zoomIn" name="fa6-brands:github" />
        </div>
        <div
          class="flex flex-col items-center justify-start gap-[3.75rem] xs:justify-end xs:gap-y-[4.5rem]"
        >
          <div
            class="mt-[25px] flex h-7 w-7 animate-zoomIn items-center justify-center rounded-full bg-white drop-shadow-md sm:-mr-56 xs:-mr-24"
          >
            <Icon
              class="fas footer-icon h-3.5 w-3.5 p-px"
              name="fa6-solid:star"
            />
          </div>

          <Icon
            class="h-2.5 w-[17px] animate-zoomOut"
            name="fa6-brands:medium"
          />
          <div
            class="-mt-6 flex h-[70px] w-[98px] flex-col items-center justify-center rounded-md from-[#ff835b] to-[#f2709c] text-center text-white drop-shadow-md bg-gradient-[270.11deg] lg:w-[150px] lg:p-5 xs:h-[100px] xs:w-[120px]"
          >
            <span
              class="font-inter-bold text-[1rem] font-extrabold leading-[2.1rem] text-white xs:text-[1.5rem]"
            >
              3K+</span
            >
            <span
              class="mb-2.5 font-inter-medium text-[0.8rem] leading-[1rem] text-white xs:text-[1rem] xs:leading-6"
              >Daily Visitors on Our Blog
            </span>
          </div>
        </div>

        <div
          class="mt-[50px] flex flex-col items-center justify-between gap-y-5"
        >
          <div
            class="ml-[30px] flex h-[22px] w-[52px] animate-zoomIn flex-row items-center justify-center gap-x-0.5 rounded-full from-[#ff835b] to-[#f2709c] text-center text-white drop-shadow-md bg-gradient-[270.11deg] xs:ml-10"
          >
            <img
              :src="clapsmall"
              class="h-3.5 w-3"
              alt="desktop-background-icon"
            />
            <span class="font-inter-semibold">237</span>
          </div>

          <div
            class="flex h-20 w-20 flex-col items-center justify-center rounded-md bg-white text-center drop-shadow-md"
          >
            <img
              :src="githubicon"
              class="absolute mt-5 flex h-[60px] w-[60px] flex-col items-center justify-center"
              alt="mobile-background-icon"
            />
            <span
              class="v2-canopas-gradient-text -mt-[15px] font-inter-bold text-[1.5rem] font-extrabold leading-[2.1rem] xl:text-[1.75rem]"
            >
              1K+</span
            >
            <span
              class="p-2.5 font-inter-medium text-base text-black-core/[0.87]"
              >GitHub Stars
            </span>
          </div>
          <img
            :src="dribble"
            class="h-4 w-4 animate-zoomOut"
            alt="mobile-background-icon"
          />
        </div>
      </div>
      <div class="container flex flex-col pt-11">
        <div class="font-inter-bold text-[2rem] leading-[3rem]">
          <h1>Crafting Code, Connecting Communities</h1>
        </div>
        <div class="font-inter-regular text-base">
          Our team at Canopas constantly seeks to learn, grow, and give back to
          the developer community.
        </div>
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
