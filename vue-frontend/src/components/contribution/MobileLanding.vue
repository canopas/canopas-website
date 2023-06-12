<template>
  <section
    class="tw-mb-16 tw-h-[470px] tw-bg-gradient-to-t tw-from-[#FF835B]/[0.1] tw-to-[#F2709C]/[0.1] tw-bg-gradient-[270.11deg]"
  >
    <div class="tw-relative tw-block md:tw-hidden">
      <div
        class="tw-absolute tw-w-[80%] tw-top-4 tw-left-[16%] tw-flex tw-flex-row tw-justify-between"
      >
        <img
          :src="heartlike"
          class="tw-mx-auto tw--mt-3 tw-w-7 tw-h-7 tw-animate-zoomIn"
          alt="mobile-background-icon"
        />

        <div
          class="tw-flex tw-justify-center tw-items-center tw-gap-y-12 tw-rounded-full tw-bg-white tw-drop-shadow-md lg:tw-ml-5 2xl:tw-ml-[8.25rem] lg:tw--mt-3 2xl:tw-mt-[1.25rem] tw-w-[29px] tw-h-[31px] tw-animate-zoomIn"
        >
          <font-awesome-icon
            class="tw-w-[14px] tw-h-[14px] tw-p-[1px] tw-text-[#000]"
            icon="thumbs-up"
          />
        </div>
      </div>
      <div
        class="tw-absolute tw-w-[90%] tw-left-[6%] tw-top-[13.5rem] xs:tw-top-[11rem] tw-flex tw-flex-row tw-justify-between tw-mx-auto"
      >
        <div
          class="tw-flex tw-flex-col tw-justify-around tw-items-center tw-mt-[45px] tw-gap-y-16 xs:tw-gap-y-20"
        >
          <div
            class="tw-mt-[20px] tw-flex tw-flex-col tw-items-center tw-justify-center tw--mt-6 tw-w-[140px] tw-h-[80px] tw-rounded-md tw-bg-white tw-text-center tw-gap-y-1.5 tw-drop-shadow-md"
          >
            <span
              class="tw--mt-[30px] tw-font-inter-bold tw-font-extrabold tw-text-[1.5rem] tw-leading-[2.1rem] v2-canopas-gradient-text"
            >
              100K+</span
            >
            <span
              class="tw-font-inter-medium tw-text-[1rem] tw-leading-[1.5rem] tw-text-black-core/[0.87]"
              >Monthly minutes Read on blog
            </span>
          </div>
          <font-awesome-icon
            class="tw-w-4 tw-h-4 tw-animate-zoomIn"
            :icon="github"
          />
        </div>
        <div
          class="tw-flex tw-flex-col tw-justify-start xs:tw-justify-end tw-items-center tw-gap-[3.75rem] xs:tw-gap-y-[4.5rem]"
        >
          <div
            class="tw-flex tw-justify-center tw-items-center tw-rounded-full tw-bg-white tw-drop-shadow-md xs:tw--mr-24 tw-mt-[25px] sm:tw--mr-[14rem] tw-w-[28px] tw-h-[28px] tw-animate-zoomIn"
          >
            <font-awesome-icon
              class="fas tw-w-[14px] tw-h-[14px] tw-p-[1px] tw-text-[#FF3D2E]"
              icon="star"
            />
          </div>

          <font-awesome-icon
            class="tw-w-[17px] tw-h-[10px] tw-animate-zoomOut"
            :icon="medium"
          />
          <div
            class="tw-flex tw-flex-col tw-items-center tw-justify-center tw--mt-6 tw-w-[98px] xs:tw-w-[120px] lg:tw-w-[150px] tw-h-[70px] xs:tw-h-[100px] tw-rounded-md tw-drop-shadow-md tw-text-center tw-text-white tw-from-[#ff835b] tw-to-[#f2709c] tw-bg-gradient-[270.11deg] lg:tw-p-[20px]"
          >
            <span
              class="tw-font-inter-bold tw-font-extrabold tw-text-[1rem] xs:tw-text-[1.5rem] tw-leading-[2.1rem] tw-text-white"
            >
              3K+</span
            >
            <span
              class="tw-mb-[10px] tw-font-inter-medium tw-text-[0.8rem] xs:tw-text-[1rem] tw-leading-[1rem] xs:tw-leading-[1.5rem] tw-text-white"
              >Daily Visitors on Our Blog
            </span>
          </div>
        </div>

        <div
          class="tw-flex tw-flex-col tw-justify-between tw-items-center tw-gap-y-5 tw-mt-[50px]"
        >
          <div
            class="tw-flex tw-flex-row tw-gap-x-[2px] tw-items-center tw-justify-center tw-ml-[30px] xs:tw-ml-[40px] tw-w-[52px] tw-h-[22px] tw-rounded-full tw-drop-shadow-md tw-from-[#ff835b] tw-to-[#f2709c] tw-bg-gradient-[270.11deg] tw-text-center tw-text-white tw-animate-zoomIn"
          >
            <img
              :src="clapsmall"
              class="tw-w-[12px] tw-h-[14px]"
              alt="desktop-background-icon"
            />
            <span class="tw-font-inter-semibold">237</span>
          </div>

          <div
            class="tw-flex tw-flex-col tw-items-center tw-justify-center tw-w-[80px] tw-h-[80px] tw-rounded-md tw-bg-white tw-text-center tw-drop-shadow-md"
          >
            <img
              :src="githubicon"
              class="tw-absolute tw-flex tw-flex-col tw-items-center tw-justify-center tw-mt-[20px] tw-w-[60px] tw-h-[60px]"
              alt="mobile-background-icon"
            />
            <span
              class="tw--mt-[15px] tw-font-inter-bold tw-font-extrabold tw-text-[1.5rem] xl:tw-text-[1.75rem] tw-leading-[2.1rem] v2-canopas-gradient-text"
            >
              1K+</span
            >
            <span
              class="tw-font-inter-medium tw-text-[1rem] tw-leading-[1.5rem] tw-text-black-core/[0.87] tw-p-[10px]"
              >GitHub Stars
            </span>
          </div>
          <img
            :src="dribble"
            class="tw-w-4 tw-h-4 tw-animate-zoomOut"
            alt="mobile-background-icon"
          />
        </div>
      </div>
      <div class="tw-container tw-flex tw-flex-col tw-pt-11">
        <div class="tw-text-[2rem] tw-leading-[3rem] tw-font-inter-bold">
          <h1>Crafting Code, Connecting Communities</h1>
        </div>
        <div class="tw-text-[1rem] tw-leading-[1.5rem] tw-font-inter-regular">
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
