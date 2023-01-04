<template>
  <div
    class="tw-relative tw-mt-[-50px] tw-bg-gradient-to-t tw-from-pink-300/[0.05] tw-to-white landing-bg-mobile-clip-path md:landing-bg-clip-path tw-overflow-hidden tw-px-0 tw-pb-[50px] md:tw-pb-[100px] tw-font-inter-regular"
  >
    <img
      v-if="width < 680"
      :src="bgMobile"
      class="tw-absolute tw-top-0 tw-left-0 tw-w-full tw-h-full tw--z-[1] tw-object-cover"
      alt="canopas-landing"
    />
    <img
      v-else
      :src="bg"
      class="tw-absolute tw-top-0 tw-left-0 tw-w-full tw-h-full tw--z-[1] tw-object-cover 2xl:tw-object-contain"
      alt="canopas-landing"
    />
    <div
      class="tw-container tw-flex tw-flex-col tw-items-center lg:tw-flex-row-reverse lg:tw-items-start tw-py-[25px] md:tw-py-[50px]"
    >
      <div class="tw-flex-[45%] lg:tw-block tw-hidden tw-mt-[40px]">
        <img :src="landing" alt="landing-view-image" loading="lazy" />
      </div>
      <div class="tw-flex-[55%] tw-flex tw-flex-col tw-mt-[50px]">
        <div
          class="tw-w-[100%] sm:tw-w-[85%] tw-pb-5 lg:tw-pb-8 xl:tw-pb-10 v2-header-3-text tw-tracking-[0] tw-text-[1.875rem] tw-leading-[2.5rem] md:tw-text-[2.65625rem] md:tw-leading-[3.28125rem] lg:tw-text-[3.28125rem] lg:tw-leading-[3.70625rem] 2xl:tw-text-[4.0625rem] 2xl:tw-leading-[4.6875rem] tw-text-black-core/[.87] tw-font-roboto-bold"
        >
          For Over A Decade We Have Helped Businesses Grow Exponentially.
          <div>
            — <span class="canopas-gradient-text"> Is Yours Next?</span>
          </div>
        </div>
        <div
          class="tw-w-[85%] md:tw-w-[65%] tw-pb-[29.5px] md:tw-pb-[47px] v2-normal-text tw-tracking-[0] tw-text-[1rem] tw-leading-[1.375rem] md:tw-text-[1.09375rem] md:tw-leading-[1.53125rem] lg:tw-text-[1.1875rem] lg:tw-leading-[1.75rem] tw-text-black-core/[0.6] !tw-font-sfpro-regular"
        >
          We develop Mobile and Web apps that help you
          <span
            class="tw-bg-no-repeat tw-bg-top-bottom-160 lg:tw-bg-top-bottom-150 tw-bg-[length:100%_10px] tw-bg-gradient-underline tw-pb-[3px] tw-font-bold"
            >generate more revenue,</span
          >&nbsp;
          <span
            class="tw-bg-no-repeat tw-bg-top-bottom-160 lg:tw-bg-top-bottom-150 tw-bg-[length:100%_10px] tw-bg-gradient-underline tw-pb-[3px] tw-font-bold"
            >acquire new users,
          </span>
          and
          <span
            class="tw-bg-no-repeat tw-bg-top-bottom-160 lg:tw-bg-top-bottom-150 tw-bg-[length:100%_10px] tw-bg-gradient-underline tw-pb-[3px] tw-font-bold"
            >deliver an exceptional user experience</span
          >.
        </div>
        <div class="tw-flex tw-items-center tw-flex-col-reverse lg:tw-flex-row">
          <div class="tw-text-center">
            <router-link
              class="tw-flex tw-items-center tw-w-max tw-m-0 tw-rounded-[3rem] tw-py-[0.7rem] tw-px-5 tw-text-center gradient-btn"
              :to="contactURL"
            >
              <span
                class="tw-mr-2.5 tw-font-normal tw-text-[1rem] tw-leading-[1.1875rem] md:tw-text-[1.09375rem] md:tw-leading-[1.3125rem] lg:tw-text-[1.1875rem] lg:tw-leading-[1.4375rem] tw-font-inter-medium !tw-tracking-[0.5px]"
                >Get Free Consultation
              </span> </router-link
            ><small
              class="lg:tw-hidden tw-font-inter-regular tw-text-black-core/[0.6] tw-text-[0.75rem] tw-leading-[1.125rem]"
              >100% MONEY BACK GUARANTEE</small
            >
          </div>

          <div class="tw-flex tw-items-center tw-pb-3.5 lg:tw-pb-0 md:tw-pl-7">
            <div v-for="client in clients" :key="client">
              <div
                v-if="client.image"
                class="tw-w-[40px] md:tw-w-[50px] tw-h-[40px] md:tw-h-[50px] tw-rounded-full"
                :class="client.className"
              >
                <img
                  :src="client.image"
                  class="tw-w-full tw-h-full tw-object-cover"
                  loading="lazy"
                  alt="phase.alt"
                />
              </div>
              <div v-if="client.plusIcon">
                <div
                  class="tw-w-[40px] md:tw-w-[50px] tw-h-[40px] md:tw-h-[50px] tw-border-[3px] tw-border-white tw-rounded-full tw-bg-[#ffc4c0]"
                  :class="client.className"
                >
                  <font-awesome-icon
                    class="fa tw-w-[35px] md:tw-w-[45px] tw-h-[23px] md:tw-pt-[10px] tw-pt-[5px] tw-text-white"
                    icon="plus"
                  />
                </div>
              </div>
            </div>
            <div class="tw-pl-[10px] md:tw-pl-[6px]">
              <div class="tw-flex tw-justify-between">
                <font-awesome-icon
                  v-for="i in 5"
                  :key="i"
                  class="fas tw-w-[14px] tw-h-[14px] tw-p-[1px] tw-text-[#FF3D2E]"
                  icon="star"
                />
              </div>

              <p
                class="tw-font-medium tw-text-[12px] tw-leading-[22px] tw-uppercase tw-text-black-core/[0.6] tw-tracking-[0]"
              >
                40+ Reviews
              </p>
            </div>
          </div>
        </div>
        <small
          class="tw-hidden lg:tw-block tw-w-[74%] xl:tw-w-[56%] tw-pt-3 lg:tw-pt-4 tw-text-[0.875rem] tw-leading-[1.125rem] tw-text-black-core/[0.6]"
        >
          We offer a
          <span class="tw-font-inter-bold"
            >100% MONEY BACK GUARANTEE FOR 14 DAYS</span
          >
          if you don't like our work. No questions asked.
        </small>
      </div>
    </div>
  </div>
</template>

<script>
import landing from "@/assets/images/landing/landing-new.webp";
import client1 from "@/assets/images/landing/client1.webp";
import client2 from "@/assets/images/landing/client2.webp";
import client3 from "@/assets/images/landing/client3.webp";
import client4 from "@/assets/images/landing/client4.webp";
import bgMobile from "@/assets/images/landing/bg/bg800.svg";
import bg from "@/assets/images/landing/bg/bg2400.svg";

import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
export default {
  data() {
    return {
      width: 680,
      landing,
      bgMobile,
      bg,
      contactURL: "/contact",
      clients: [
        {
          className: "tw-z-0",
          image: client4,
        },
        {
          className: "tw-z-10 -tw-ml-[21px]",
          image: client3,
        },
        {
          className: "tw-z-20 -tw-ml-[21px]",
          image: client2,
        },
        {
          className: "tw-z-30 -tw-ml-[21px]",
          image: client1,
        },
        {
          className: "tw-z-40 -tw-ml-[21px]",
          plusIcon: true,
        },
      ],
    };
  },
  components: {
    FontAwesomeIcon,
  },
  mounted() {
    this.width = window.innerWidth;
  },
};
</script>
