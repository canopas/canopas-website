<template>
  <section class="tw-mb-48 2xl:tw-mb-[14rem] 2xll:tw-mb-[16rem] tw-mt-[25px]">
    <div
      class="tw-relative xll:tw-container tw-w-full tw-h-full tw-bg-gradient-to-r tw-from-[#FFFFFF] tw-via-[#F2709C]/[0.05] tw-to-[#FFFFFF] tw-bg-gradient-[180.05deg] tw-pb-16 lg:tw-pt-[3rem]"
    >
      <div
        class="tw-flex tw-flex-col tw-items-center tw-mx-auto tw-w-[52%] lg:tw-w-[55%] xl:tw-w-[62%] 2xl:tw-w-[55%] 2xll:tw-w-[45%] 3xl:tw-w-[55%] md:tw-text-center"
      >
        <div
          class="tw-text-[2rem] xl:tw-text-[3.75rem] tw-leading-[3rem] xl:tw-leading-[5.625rem] tw-font-inter-bold"
        >
          <h1>Crafting Code, Connecting Communities</h1>
        </div>
        <div
          class="tw-mt-[10px] xl:tw-mt-[20px] tw-text-[1rem] xl:tw-text-[1.1875rem] tw-leading-[1.5rem] xl:tw-leading-[1.78125rem] tw-font-inter-medium"
        >
          At Canopas, we believe in the power of open source software and the
          incredible community that makes it possible. We actively contribute to
          this vibrant ecosystem with the goal of advancing technology,
          fostering innovation, and empowering developers around the globe.
        </div>
      </div>
      <div
        class="tw-absolute tw-top-0 lg:tw-top-[30px] tw-left-[10%] tw-flex tw-flex-col tw-justify-between tw-gap-y-12 lg:tw-gap-y-20 2xl:tw-gap-y-24 tw-w-[80%]"
      >
        <font-awesome-icon
          class="tw--ml-3 tw-w-5 lg:tw-w-7 tw-h-[1.25rem] lg:tw-h-7 tw-animate-zoomIn"
          :icon="github"
        />
        <div
          class="tw-flex tw-flex-row tw-items-center tw-gap-x-1.5 tw-justify-center tw--ml-16 tw-w-[65px] tw-h-[25px] tw-rounded-full tw-drop-shadow-md tw-text-center tw-text-white tw-from-[#ff835b] tw-to-[#f2709c] tw-bg-gradient-[270.11deg] tw-animate-zoomOut"
        >
          <img
            :src="clapsmall"
            class="tw-w-[16px] tw-h-[18px]"
            alt="desktop-background-icon"
          />
          <span class="tw-font-inter-semibold">315</span>
        </div>
        <img
          :src="heartlike"
          class="tw-ml-[35px] lg:tw-ml-[52px] xl:tw-ml-[96px] tw--mt-[52px] lg:tw--mt-[90px] tw-w-10 tw-h-10 tw-animate-zoomIn"
          alt="desktop-background-icon"
        />
      </div>

      <div
        class="tw-absolute tw-top-[18rem] xl:tw-top-[25rem] xll:tw-top-[32rem] tw-left-[10%] tw-flex tw-flex-row tw-justify-between tw-w-[80%]"
      >
        <div
          class="tw-flex tw-flex-row tw-justify-between tw-gap-x-[4.75rem] xl:tw-gap-x-[4rem] tw-gap-y-20"
        >
          <div
            class="tw-flex tw-flex-col tw-items-center tw-justify-center tw--mt-6 tw-w-[145px] tw-h-[125px] tw-rounded-md tw-bg-white tw-text-center tw-drop-shadow-md"
          >
            <span
              class="tw--mt-[30px] tw-font-inter-bold tw-font-extrabold tw-text-[1.5rem] xl:tw-text-[1.75rem] tw-leading-[2.1rem] v2-canopas-gradient-text"
            >
              100K+</span
            >
            <span
              class="tw-font-inter-medium tw-text-[1rem] tw-leading-[1.5rem] tw-text-black-core/[0.87] tw-p-[10px]"
              >Monthly minutes Read on blog
            </span>
          </div>
          <div
            class="tw-flex tw-justify-center tw-items-center tw-rounded-full tw-bg-white tw-drop-shadow-md lg:tw-ml-5 2xl:tw-ml-[8.25rem] lg:tw--mt-3 2xl:tw-mt-[1.25rem] tw-w-[29px] tw-h-[31px] tw-animate-zoomIn"
          >
            <font-awesome-icon
              class="tw-w-[14px] tw-h-[14px] tw-p-[1px] tw-text-[#000]"
              icon="thumbs-up"
            />
          </div>
        </div>
      </div>

      <div
        class="tw-absolute tw-top-[19rem] xl:tw-top-[25rem] xll:tw-top-[32rem] tw-left-[46%] tw-flex tw-flex-row tw-justify-between tw-w-[80%]"
      >
        <div
          class="tw-flex tw-flex-row tw-justify-between tw-gap-x-11 xl:tw-gap-x-[7rem] 2xll:tw-gap-x-[9rem]"
        >
          <div
            class="tw-flex tw-flex-col tw-items-center tw-justify-center tw--ml-[1rem] tw-w-[122px] lg:tw-w-[150px] tw-h-[100px] tw-rounded-md tw-drop-shadow-md tw-text-center tw-text-white tw-from-[#ff835b] tw-to-[#f2709c] tw-bg-gradient-[270.11deg] lg:tw-p-[20px]"
          >
            <span
              class="tw-font-inter-bold tw-font-extrabold tw-text-[1.5rem] tw-leading-[1.1rem] tw-text-white"
            >
              3K+</span
            >
            <span
              class="tw-mt-[10px] tw-font-inter-medium tw-text-[1rem] tw-leading-[1.5rem] tw-text-white"
              >Daily Visitors on Our Blog
            </span>
          </div>

          <div
            class="tw-flex tw-justify-center tw-items-center tw--mt-9 2xl:tw-mt-[5.75rem] tw-w-7 tw-h-7 tw-rounded-full tw-bg-white tw-drop-shadow-md tw-animate-zoomOut"
          >
            <font-awesome-icon
              class="fas tw-w-[14px] tw-h-[14px] tw-p-[1px] tw-text-[#FF3D2E]"
              icon="star"
            />
          </div>
          <img
            :src="dribble"
            class="tw-ml-24 tw-mt-5 tw-w-5 lg:tw-w-7 tw-h-5 lg:tw-h-7 tw-animate-zoomIn"
            alt="desktop-background-icon"
          />
        </div>
      </div>
      <div
        class="tw-absolute tw-left-[84%] tw-top-[15%] tw-flex tw-flex-row tw-justify-between tw-w-[80%]"
      >
        <div
          class="tw-flex tw-flex-col tw-justify-between tw-gap-y-36 lg:tw-gap-y-40 xl:tw-gap-y-48 2xl:tw-gap-y-56"
        >
          <font-awesome-icon
            class="tw--ml-6 tw-w-5 lg:tw-w-7 tw-h-[1rem] tw-animate-zoomIn"
            :icon="medium"
          />
          <div
            class="tw-flex tw-flex-col tw-items-center tw-justify-center tw-w-[95px] xl:tw-w-[120px] tw-h-[75px] xl:tw-h-[95px] tw-rounded-md tw-bg-white tw-text-center tw-drop-shadow-md"
          >
            <img
              :src="githubicon"
              class="tw-absolute tw-flex tw-flex-col tw-items-center tw-justify-center tw-w-[70px] tw-h-[70px]"
              alt="background-githubicon"
            />
            <span
              class="tw--mt-[10px] tw-font-inter-bold tw-font-extrabold tw-text-[1.5rem] xl:tw-text-[1.75rem] tw-leading-[2.1rem] v2-canopas-gradient-text"
            >
              1K+</span
            >
            <span
              class="tw-font-inter-medium tw-text-[1rem] tw-leading-[1.5rem] tw-text-black-core/[0.87] tw-p-[10px] xl:tw-p-[18px]"
              >GitHub Stars
            </span>
          </div>
        </div>
      </div>
      <div
        class="tw-absolute tw-left-[90%] tw--top-[10%] lg:tw-top-[3%] 2xll:tw--top-[7%] tw-flex tw-flex-row tw-justify-between tw-w-[80%]"
      >
        <div
          class="tw-flex tw-flex-col tw-justify-between tw-gap-y-20 lg:tw-gap-20 xl:tw-gap-28 2xl:tw-gap-y-36"
        >
          <img
            :src="clap"
            class="tw-w-14 tw-h-14 tw-animate-zoomOut"
            alt="desktop-background-icon"
          />
          <div
            class="tw-flex tw-flex-row tw-gap-x-1.5 tw-items-center tw-justify-center tw--ml-[30px] lg:tw-ml-5 tw-w-[65px] tw-h-[30px] tw-rounded-full tw-drop-shadow-md tw-from-[#ff835b] tw-to-[#f2709c] tw-bg-gradient-[270.11deg] tw-text-center tw-text-white tw-animate-zoomIn"
          >
            <img
              :src="clapsmall"
              class="tw-w-[16px] tw-h-[18px]"
              alt="desktop-background-icon"
            />
            <span class="tw-font-inter-semibold">237</span>
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

import { faMediumM, faGithub } from "@fortawesome/free-brands-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

export default {
  data() {
    return {
      heartlike,
      clap,
      clapsmall,
      dribble,
      githubicon,
      medium: faMediumM,
      github: faGithub,
    };
  },
  components: {
    FontAwesomeIcon,
  },
};
</script>
