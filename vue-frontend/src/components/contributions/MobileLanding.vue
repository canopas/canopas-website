<template>
  <section
    class="tw-mb-16 tw-h-[470px] tw-bg-gradient-to-t tw-from-[#FF835B]/[0.1] tw-to-[#F2709C]/[0.1] tw-bg-gradient-[270.11deg]"
  >
    <div class="tw-relative">
      <div
        class="tw-absolute tw-left-[16%] tw-top-4 tw-flex tw-w-[80%] tw-flex-row tw-justify-between"
      >
        <img
          :src="heartlike"
          class="tw-mx-auto tw--mt-3 tw-h-7 tw-w-7 tw-animate-zoomIn"
          alt="mobile-background-icon"
        />

        <div
          class="tw-flex tw-h-[31px] tw-w-[29px] tw-animate-zoomIn tw-items-center tw-justify-center tw-gap-y-12 tw-rounded-full tw-bg-white tw-drop-shadow-md lg:tw--mt-3 lg:tw-ml-5 2xl:tw-ml-[8.25rem] 2xl:tw-mt-[1.25rem]"
        >
          <font-awesome-icon
            class="tw-h-[14px] tw-w-[14px] tw-p-[1px] tw-text-[#000]"
            icon="thumbs-up"
          />
        </div>
      </div>
      <div
        class="tw-absolute tw-left-[6%] tw-top-[13.5rem] tw-mx-auto tw-flex tw-w-[90%] tw-flex-row tw-justify-between xs:tw-top-[11rem]"
      >
        <div
          class="tw-mt-[45px] tw-flex tw-flex-col tw-items-center tw-justify-around tw-gap-y-16 xs:tw-gap-y-20"
        >
          <div
            class="tw--mt-6 tw-mt-[20px] tw-flex tw-h-[80px] tw-w-[140px] tw-flex-col tw-items-center tw-justify-center tw-gap-y-1.5 tw-rounded-md tw-bg-white tw-text-center tw-drop-shadow-md"
          >
            <span
              class="v2-canopas-gradient-text tw--mt-[30px] tw-font-inter-bold tw-text-[1.5rem] tw-font-extrabold tw-leading-[2.1rem]"
            >
              100K+</span
            >
            <span
              class="tw-font-inter-medium tw-text-[1rem] tw-leading-[1.5rem] tw-text-black-core/[0.87]"
              >Monthly minutes Read on blog
            </span>
          </div>
          <font-awesome-icon
            class="tw-h-4 tw-w-4 tw-animate-zoomIn"
            :icon="github"
          />
        </div>
        <div
          class="tw-flex tw-flex-col tw-items-center tw-justify-start tw-gap-[3.75rem] xs:tw-justify-end xs:tw-gap-y-[4.5rem]"
        >
          <div
            class="tw-mt-[25px] tw-flex tw-h-[28px] tw-w-[28px] tw-animate-zoomIn tw-items-center tw-justify-center tw-rounded-full tw-bg-white tw-drop-shadow-md sm:tw--mr-[14rem] xs:tw--mr-24"
          >
            <font-awesome-icon
              class="fas footer-icon tw-h-[14px] tw-w-[14px] tw-p-[1px]"
              icon="star"
            />
          </div>

          <font-awesome-icon
            class="tw-h-[10px] tw-w-[17px] tw-animate-zoomOut"
            :icon="medium"
          />
          <div
            class="tw--mt-6 tw-flex tw-h-[70px] tw-w-[98px] tw-flex-col tw-items-center tw-justify-center tw-rounded-md tw-from-[#ff835b] tw-to-[#f2709c] tw-text-center tw-text-white tw-drop-shadow-md tw-bg-gradient-[270.11deg] lg:tw-w-[150px] lg:tw-p-[20px] xs:tw-h-[100px] xs:tw-w-[120px]"
          >
            <span
              class="tw-font-inter-bold tw-text-[1rem] tw-font-extrabold tw-leading-[2.1rem] tw-text-white xs:tw-text-[1.5rem]"
            >
              3K+</span
            >
            <span
              class="tw-mb-[10px] tw-font-inter-medium tw-text-[0.8rem] tw-leading-[1rem] tw-text-white xs:tw-text-[1rem] xs:tw-leading-[1.5rem]"
              >Daily Visitors on Our Blog
            </span>
          </div>
        </div>

        <div
          class="tw-mt-[50px] tw-flex tw-flex-col tw-items-center tw-justify-between tw-gap-y-5"
        >
          <div
            class="tw-ml-[30px] tw-flex tw-h-[22px] tw-w-[52px] tw-animate-zoomIn tw-flex-row tw-items-center tw-justify-center tw-gap-x-[2px] tw-rounded-full tw-from-[#ff835b] tw-to-[#f2709c] tw-text-center tw-text-white tw-drop-shadow-md tw-bg-gradient-[270.11deg] xs:tw-ml-[40px]"
          >
            <img
              :src="clapsmall"
              class="tw-h-[14px] tw-w-[12px]"
              alt="desktop-background-icon"
            />
            <span class="tw-font-inter-semibold">237</span>
          </div>

          <div
            class="tw-flex tw-h-[80px] tw-w-[80px] tw-flex-col tw-items-center tw-justify-center tw-rounded-md tw-bg-white tw-text-center tw-drop-shadow-md"
          >
            <img
              :src="githubicon"
              class="tw-absolute tw-mt-[20px] tw-flex tw-h-[60px] tw-w-[60px] tw-flex-col tw-items-center tw-justify-center"
              alt="mobile-background-icon"
            />
            <span
              class="v2-canopas-gradient-text tw--mt-[15px] tw-font-inter-bold tw-text-[1.5rem] tw-font-extrabold tw-leading-[2.1rem] xl:tw-text-[1.75rem]"
            >
              1K+</span
            >
            <span
              class="tw-p-[10px] tw-font-inter-medium tw-text-[1rem] tw-leading-[1.5rem] tw-text-black-core/[0.87]"
              >GitHub Stars
            </span>
          </div>
          <img
            :src="dribble"
            class="tw-h-4 tw-w-4 tw-animate-zoomOut"
            alt="mobile-background-icon"
          />
        </div>
      </div>
      <div class="tw-container tw-flex tw-flex-col tw-pt-11">
        <div class="tw-font-inter-bold tw-text-[2rem] tw-leading-[3rem]">
          <h1>Crafting Code, Connecting Communities</h1>
        </div>
        <div class="tw-font-inter-regular tw-text-[1rem] tw-leading-[1.5rem]">
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
import { faMediumM, faGithub } from "@fortawesome/free-brands-svg-icons";

import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

export default {
  data() {
    return {
      heartlike,
      dribble,
      clapsmall,
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
