<template>
  <section
    class="mb-48 2xl:mb-56 2xll:mb-64 xll:mb-[25rem] mt-[25px] bg-gradient-to-r from-[#FFFFFF] via-[#F2709C]/[0.05] to-[#FFFFFF] bg-gradient-[180.05deg]"
  >
    <div class="relative xll:container w-full h-full pb-16 lg:pt-12">
      <div
        class="flex flex-col items-center mx-auto w-[52%] lg:w-[55%] xl:w-[62%] 2xll:w-[45%] xll:w-[60%] 3xl:w-[60%] md:text-center"
      >
        <div
          class="text-[2rem] xl:text-[3.75rem] leading-[3rem] xl:leading-[5.625rem] font-inter-bold"
        >
          <h1>Crafting Code, Connecting Communities</h1>
        </div>
        <div
          class="mt-2.5 xl:mt-5 text-[1rem] xl:text-[1.1875rem] leading-6 xl:leading-[1.78125rem] font-inter-medium"
        >
          At Canopas, we believe in the power of open source software and the
          incredible community that makes it possible. We actively contribute to
          this vibrant ecosystem with the goal of advancing technology,
          fostering innovation, and empowering developers around the globe.
        </div>
      </div>
      <div
        class="absolute top-0 lg:top-[30px] left-[10%] flex flex-col justify-between gap-y-12 lg:gap-y-20 2xl:gap-y-24 w-[80%]"
      >
        <Icon
          class="-ml-3 w-5 lg:w-7 h-5 lg:h-7 animate-zoomIn"
          name="fa6-brands:github"
        />
        <div
          class="flex flex-row items-center gap-x-1.5 justify-center -ml-16 w-[65px] h-[25px] rounded-full drop-shadow-md text-center text-white from-[#ff835b] to-[#f2709c] bg-gradient-[270.11deg] animate-zoomOut"
        >
          <img
            :src="clapsmall"
            class="w-4 h-[18px]"
            alt="desktop-background-icon"
          />
          <span class="font-inter-semibold">315</span>
        </div>
        <img
          :src="heartlike"
          class="ml-[35px] lg:ml-[52px] xl:ml-24 -mt-[52px] lg:-mt-[90px] w-10 h-10 animate-zoomIn"
          alt="desktop-background-icon"
        />
      </div>

      <div
        class="absolute top-72 xl:top-[25rem] xll:top-[32rem] left-[10%] flex flex-row justify-between w-[80%]"
      >
        <div
          class="flex flex-row justify-between gap-x-[4.75rem] xl:gap-x-16 gap-y-20"
        >
          <div
            class="flex flex-col items-center justify-center -mt-6 w-[145px] h-[125px] rounded-md bg-white text-center drop-shadow-md"
          >
            <span
              class="-mt-[30px] font-inter-bold font-extrabold text-[1.5rem] xl:text-[1.75rem] leading-[2.1rem] v2-canopas-gradient-text"
            >
              100K+</span
            >
            <span
              class="font-inter-medium text-base text-black-core/[0.87] p-2.5"
              >Monthly minutes Read on blog
            </span>
          </div>
          <div
            class="flex justify-center items-center rounded-full bg-white drop-shadow-md lg:ml-5 2xl:ml-[8.25rem] lg:-mt-3 2xl:mt-5 w-[29px] h-[31px] animate-zoomIn"
          >
            <Icon
              class="w-3.5 h-3.5 p-px text-[#000]"
              name="fa6-solid:thumbs-up"
            />
          </div>
        </div>
      </div>

      <div
        class="absolute top-[19rem] xl:top-[25rem] xll:top-[32rem] left-[46%] flex flex-row justify-between w-[80%]"
      >
        <div
          class="flex flex-row justify-between gap-x-11 xl:gap-x-28 2xll:gap-x-36"
        >
          <div
            class="flex flex-col items-center justify-center -ml-4 w-[122px] lg:w-[150px] h-[100px] rounded-md drop-shadow-md text-center text-white from-[#ff835b] to-[#f2709c] bg-gradient-[270.11deg] lg:p-5"
          >
            <span
              class="font-inter-bold font-extrabold text-[1.5rem] leading-[1.1rem] text-white"
            >
              3K+</span
            >
            <span class="mt-2.5 font-inter-medium text-base text-white"
              >Daily Visitors on Our Blog
            </span>
          </div>

          <div
            class="flex justify-center items-center -mt-9 2xl:mt-[5.75rem] w-7 h-7 rounded-full bg-white drop-shadow-md animate-zoomOut"
          >
            <Icon
              class="fas w-3.5 h-3.5 p-px footer-icon"
              name="fa6-solid:star"
            />
          </div>
          <img
            :src="dribble"
            class="ml-24 mt-5 w-5 lg:w-7 h-5 lg:h-7 animate-zoomIn"
            alt="desktop-background-icon"
          />
        </div>
      </div>
      <div
        class="absolute left-[84%] top-[15%] flex flex-row justify-between w-[80%]"
      >
        <div
          class="flex flex-col justify-between gap-y-36 lg:gap-y-40 xl:gap-y-48 2xl:gap-y-56"
        >
          <Icon
            class="-ml-6 w-5 lg:w-7 h-4 animate-zoomIn"
            name="fa6-brands:medium"
          />
          <div
            class="flex flex-col items-center justify-center w-[95px] xl:w-[120px] h-[75px] xl:h-[95px] rounded-md bg-white text-center drop-shadow-md"
          >
            <img
              :src="githubicon"
              class="absolute flex flex-col items-center justify-center w-[70px] h-[70px]"
              alt="background-githubicon"
            />
            <span
              class="-mt-2.5 font-inter-bold font-extrabold text-[1.5rem] xl:text-[1.75rem] leading-[2.1rem] v2-canopas-gradient-text"
            >
              1K+</span
            >
            <span
              class="font-inter-medium text-base text-black-core/[0.87] p-2.5 xl:p-[18px]"
              >GitHub Stars
            </span>
          </div>
        </div>
      </div>
      <div
        class="absolute left-[90%] -top-[10%] lg:top-[3%] 2xll:-top-[7%] flex flex-row justify-between w-[80%]"
      >
        <div
          class="flex flex-col justify-between gap-y-20 lg:gap-20 xl:gap-28 2xl:gap-y-36"
        >
          <img
            :src="clap"
            class="w-14 h-14 animate-zoomOut"
            alt="desktop-background-icon"
          />
          <div
            class="flex flex-row gap-x-1.5 items-center justify-center -ml-[30px] lg:ml-5 w-[65px] h-[30px] rounded-full drop-shadow-md from-[#ff835b] to-[#f2709c] bg-gradient-[270.11deg] text-center text-white animate-zoomIn"
          >
            <img
              :src="clapsmall"
              class="w-4 h-[18px]"
              alt="desktop-background-icon"
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
