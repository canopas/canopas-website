<template>
  <div
    class="tw-relative tw-bg-gradient-to-t tw-from-pink-300/[0.05] tw-to-white landing-bg-mobile-clip-path md:landing-bg-clip-path tw-overflow-hidden tw-px-0 tw-pb-[50px] md:tw-pb-[100px] tw-font-inter-regular"
  >
    <img
      :src="background2100"
      :srcset="`${background400} 400w, ${background800} 800w, ${background1200} 1200w, ${background2100} 2100w`"
      class="tw-absolute tw-top-0 tw-left-0 tw-w-full tw-h-full tw--z-[1] tw-object-contain"
      alt="canopas-landing"
    />

    <div
      class="tw-container tw-flex tw-flex-col tw-items-center lg:tw-flex-row-reverse lg:tw-items-start tw-pb-[25px] md:tw-pb-[50px]"
    >
      <aspect-ratio
        class="tw-flex-[45%] lg:tw-block tw-hidden tw-mt-[40px]"
        height="60%"
      >
        <img
          class="tw-w-full tw-h-full tw-object-contain"
          :src="landing800"
          :srcset="`${landing400} 400w, ${landing800} 800w`"
          alt="landing-view-image"
        />
      </aspect-ratio>
      <div
        class="tw-flex-[55%] tw-flex tw-flex-col tw-mt-[20px] lg:tw-mt-[50px]"
      >
        <h1
          class="tw-w-[100%] sm:tw-w-[85%] tw-pb-5 lg:tw-pb-8 xl:tw-pb-10 v2-header-3-text tw-tracking-[0] tw-text-[2.1875rem] tw-leading-[2.8125rem] md:tw-text-[2.65625rem] md:tw-leading-[3.28125rem] lg:tw-text-[3.28125rem] lg:tw-leading-[3.70625rem] 2xl:tw-text-[4.0625rem] 2xl:tw-leading-[4.6875rem] tw-text-black-core/[.87] tw-font-roboto-bold"
        >
          For Over A Decade We Have Helped Businesses Grow Exponentially.
          <div>
            — <span class="canopas-gradient-text"> Is Yours Next?</span>
          </div>
        </h1>
        <div
          class="tw-w-[85%] md:tw-w-[65%] tw-pb-[29.5px] md:tw-pb-[47px] v2-normal-text tw-tracking-[0] tw-text-[1rem] tw-leading-[1.375rem] md:tw-text-[1.09375rem] md:tw-leading-[1.53125rem] lg:tw-text-[1.1875rem] lg:tw-leading-[1.75rem] tw-text-black-core/[0.6] tw-font-inter-regular"
        >
          We develop Mobile and Web apps that help you
          <span
            class="tw-bg-no-repeat tw-bg-bottom tw-bg-[length:100%_2.5px] tw-bg-gradient-underline tw-pb-[3px] tw-font-inter-bold"
            >generate more revenue,</span
          >&nbsp;
          <span
            class="tw-bg-no-repeat tw-bg-bottom tw-bg-[length:100%_2.5px] tw-bg-gradient-underline tw-pb-[3px] tw-font-inter-bold"
            >acquire new users,</span
          >
          and
          <span
            class="tw-bg-no-repeat tw-bg-bottom tw-bg-[length:100%_2.5px] tw-bg-gradient-underline tw-pb-[3px] tw-font-inter-bold"
            >deliver an exceptional user experience.</span
          >
        </div>

        <div class="tw-flex tw-items-center tw-flex-col-reverse lg:tw-flex-row">
          <div class="tw-text-center">
            <router-link
              class="tw-flex tw-items-center tw-w-max tw-m-0 tw-rounded-full tw-p-3 tw-text-center gradient-btn consultation-btn"
              :to="contactURL"
              @click.native="mixpanel.track('tap_landing_cta')"
            >
              <span
                class="tw-mr-2.5 tw-font-normal tw-text-[1rem] tw-leading-[1.1875rem] md:tw-text-[1.09375rem] md:tw-leading-[1.3125rem] lg:tw-text-[1.1875rem] lg:tw-leading-[1.4375rem] tw-font-inter-medium !tw-tracking-[0]"
                >Get Free Consultation
              </span> </router-link
            ><small
              class="lg:tw-hidden tw-font-inter-regular tw-text-black-core/[0.6] tw-text-[0.75rem] tw-leading-[1.125rem]"
              >100% MONEY BACK GUARANTEE <br />
              FOR 14 DAYS</small
            >
          </div>

          <div
            class="tw-flex tw-items-center tw-pb-3.5 md:tw-pl-7 lg:tw-pl-3 xl:tw-pl-7 lg:tw-pb-0"
          >
            <div v-for="client in clients" :key="client">
              <div
                v-if="client.image"
                class="tw-rounded-full"
                :class="client.className"
              >
                <img
                  :src="client.image"
                  class="tw-w-[40px] md:tw-w-[50px] tw-h-[40px] md:tw-h-[50px] tw-object-cover"
                  loading="lazy"
                  alt="client-image"
                />
              </div>
              <div v-if="client.plusIcon">
                <div
                  class="tw-w-[40px] md:tw-w-[50px] tw-h-[40px] md:tw-h-[50px] tw-border-[3px] tw-border-white tw-rounded-full tw-bg-[#ffc4c0]"
                  :class="client.className"
                >
                  <font-awesome-icon
                    class="fa tw-w-[35px] md:tw-w-[45px] tw-h-[23px] tw-pt-[5px] md:tw-pt-[10px] tw-text-white"
                    icon="plus"
                  />
                </div>
              </div>
            </div>
            <div class="tw-pl-[10px] md:tw-pl-[6px]">
              <div class="tw-flex tw-justify-between tw-mt-[3px]">
                <font-awesome-icon
                  v-for="i in 5"
                  :key="i"
                  class="fas tw-w-[14px] tw-h-[14px] tw-p-[1px] tw-text-[#FF3D2E]"
                  icon="star"
                />
              </div>
              <div class="">
                <span
                  class="tw-font-medium tw-text-[12px] tw-leading-[22px] tw-uppercase tw-text-black-core/[0.6] tw-tracking-[0]"
                >
                  40+ Reviews
                </span>
              </div>
            </div>
          </div>
        </div>
        <small
          class="tw-hidden lg:tw-block tw-w-[74%] xl:tw-w-[56%] tw-pt-3 lg:tw-pt-4 tw-text-[0.875rem] tw-leading-[1.125rem] tw-text-black-core/[0.6]"
        >
          We offer a
          <span class="tw-font-inter-bold"
            >100% MONEY BACK GUARANTEE FOR 14 DAYS
          </span>
          if you don't like our work. No questions asked.
        </small>
      </div>
    </div>
  </div>
</template>
<script>
import landing400 from "@/assets/images/landing/landing-400w.webp";
import landing800 from "@/assets/images/landing/landing-800w.webp";
import client1 from "@/assets/images/landing/client1-100w.webp";
import client2 from "@/assets/images/landing/client2-100w.webp";
import client3 from "@/assets/images/landing/client3-100w.webp";
import client4 from "@/assets/images/landing/client4-100w.webp";
import background400 from "@/assets/images/landing/bg/background-400w.webp";
import background800 from "@/assets/images/landing/bg/background-800w.webp";
import background1200 from "@/assets/images/landing/bg/background-1200w.webp";
import background2100 from "@/assets/images/landing/bg/background-2100w.webp";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import AspectRatio from "../utils/AspectRatio.vue";
export default {
  data() {
    return {
      width: 680,
      landing400,
      landing800,
      background400,
      background800,
      background1200,
      background2100,
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
    AspectRatio,
  },
  inject: ["mixpanel"],
  mounted() {
    this.width = window.innerWidth;
  },
};
</script>
