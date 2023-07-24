<template>
  <div
    class="tw-relative tw-mt-[4.375rem] md:tw-mt-[14.063rem] tw-overflow-hidden tw-px-0 tw-font-inter-regular md:tw-font-inter-medium"
  >
    <img
      :src="bg2400"
      :srcset="`${bg400} 400w, ${bg800} 800w, ${bg1200} 1200w, ${bg2400} 2400w`"
      class="tw-absolute tw-top-0 tw-left-0 tw-w-full tw--z-[1]"
      alt="canopas-cta"
    />
    <div
      class="tw-container tw-m-auto md:tw-w-[86%] lg:tw-w-4/5 xl:tw-w-[86%] tw-px-[25px] tw-pt-[61px] md:tw-pt-[162px] tw-pb-[50px] md:tw-pb-[70px] tw-opacity-1 tw-text-center"
    >
      <div class="md:tw-w-[86%] lg:tw-w-4/5 xl:tw-w-[86%] tw-m-auto">
        <span
          style="-webkit-box-decoration-break: clone"
          class="tw-w-full tw-text-[2.1875rem] tw-leading-[2.75rem] tw-tracking-[1px] md:tw-text-[2.8125rem] md:tw-leading-[3.4375rem] lg:tw-text-[3.4375rem] lg:tw-leading-[4.0281rem] tw-font-roboto-bold canopas-gradient-text"
        >
          How do you know we’re a good fit for you?
        </span>
        <h2
          class="tw-mt-6 tw-text-[1.25rem] tw-leading-[1.5625rem] md:tw-text-[1.375rem] md:tw-leading-[1.6875rem] lg:tw-text-[1.5rem] lg:tw-leading-[1.8156rem] tw-tracking-normal tw-text-black-core/[.87]"
        >
          Let’s have a short chat.
        </h2>
        <div class="tw-flex tw-flex-wrap tw-justify-center">
          <router-link
            class="tw-flex tw-w-max tw-items-center tw-rounded-full tw-py-3 tw-px-3 gradient-btn consultation-btn"
            :to="contactURL"
            @click.native="mixpanel.track('tap_home_touch_cta')"
          >
            <span
              class="tw-mr-2.5 tw-font-normal tw-text-[1rem] tw-leading-[1.1875rem] md:tw-text-[1.09375rem] md:tw-leading-[1.3125rem] lg:tw-text-[1.1875rem] lg:tw-leading-[1.4375rem] tw-font-inter-medium !tw-tracking-[0]"
              >Get in touch</span
            >
          </router-link>
        </div>
        <div
          class="tw-mt-12 tw-w-full tw-text-[1.25rem] tw-leading-[1.5625rem] md:tw-text-[1.375rem] md:tw-leading-[1.6875rem] lg:tw-text-[1.5rem] lg:tw-leading-[1.8156rem] tw-text-left"
        >
          <div class="v2-canopas-gradient-text">
            Four reasons why do we do this call?
          </div>
        </div>
        <ul
          class="tw-w-full tw-mt-4 tw-list-none tw-p-0 tw-text-[1rem] tw-leading-[1.1875rem] md:tw-text-[1.1875rem] md:tw-leading-[1.625rem] lg:tw-text-[1.375rem] lg:tw-leading-[1.75rem] tw-text-left tw-text-black-core/[.87]"
        >
          <li
            v-for="reason in reasons"
            :key="reason"
            class="tw-flex tw-py-[7.5px] md:tw-py-[17.5px] tw-px-0"
          >
            <span
              class="tw-inline-block tw-mt-[5px] tw-mr-[15px] md:tw-mr-[18px] tw-w-[12px] tw-h-[12px] md:tw-w-[15px] md:tw-h-[15px] tw-rounded-[2px] md:tw-rounded-[5px] tw-bg-gradient-to-b tw-from-[#F2709C] tw-to-[#FF9472] tw-content-none"
            ></span>
            <div class="tw-w-[90%]">
              {{ reason }}
            </div>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script type="module">
import bg400 from "@/assets/images/CTA/bg/400.svg";
import bg800 from "@/assets/images/CTA/bg/800.svg";
import bg1200 from "@/assets/images/CTA/bg/1200.svg";
import bg2400 from "@/assets/images/CTA/bg/2400.svg";

export default {
  data() {
    return {
      reasons: [
        "It gives an opportunity to learn more about each other.",
        "We want to give you a chance to put your idea into action. You might have faced this situation before where you wish you had someone to help you.",
        "Sometimes, you want to check the feasibility of your idea. We will give you honest feedback based on our industry experience.",
        "We offer 100% Money-back guarantee. so, there’s no risk to you.",
      ],
      contactURL: "/contact",
      bg400,
      bg800,
      bg1200,
      bg2400,
    };
  },
  inject: ["mixpanel"],
};
</script>
