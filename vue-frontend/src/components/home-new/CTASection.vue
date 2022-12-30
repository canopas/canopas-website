<template>
  <div class="tw-relative tw-mt-[150px] tw-px-0 tw-overflow-hidden">
    <img
      :src="bg2400"
      :srcset="`${bg400} 400w, ${bg800} 800w, ${bg1200} 1200w, ${bg2400} 2400w`"
      class="tw-absolute tw-top-0 tw-left-0 tw-w-full tw--z-[1]"
      alt="canopas-cta"
    />
    <div
      class="tw-container sm:tw-w-4/5 lg:tw-w-[70%] tw-m-auto tw-px-[25px] tw-pt-[20%] md:tw-pt-[10%] tw-pb-[5%] tw-text-center tw-opacity-1"
    >
      <div
        class="lg:tw-w-4/5 tw-m-auto tw-text-[2.1875rem] tw-leading-[2.75rem] tw-tracking-[1px] md:tw-text-[2.8125rem] md:tw-leading-[3.4375rem] lg:tw-text-[3.4375rem] lg:tw-leading-[4.0281rem] tw-font-roboto-bold canopas-gradient-text"
      >
        How do you know we’re a good fit for you?
      </div>
      <div
        class="tw-mt-6 tw-text-[1.25rem] tw-leading-[1.5625rem] md:tw-text-[1.375rem] md:tw-leading-[1.6875rem] lg:tw-text-[1.5rem] lg:tw-leading-[1.8156rem] tw-tracking-normal tw-text-black-core/[.87]"
      >
        Let’s have a short chat.
      </div>
      <div class="tw-flex tw-flex-wrap tw-justify-center">
        <router-link
          class="tw-flex tw-items-center tw-rounded-[2rem] gradient-btn"
          :to="contactURL"
        >
          <span
            class="!tw-font-semibold tw-text-[0.9375rem] tw-leading-[1.25rem] md:tw-text-[1.0625rem] md:tw-leading-[1.3125rem] lg:tw-text-[1.1875rem] lg:tw-leading-[1.4368rem] tw-tracking-wide"
            >Get in touch</span
          >
        </router-link>
      </div>
      <div
        class="lg:tw-w-[75%] tw-m-auto tw-mt-12 tw-text-[1.25rem] tw-leading-[1.5625rem] md:tw-text-[1.375rem] md:tw-leading-[1.6875rem] lg:tw-text-[1.5rem] lg:tw-leading-[1.8156rem] tw-text-left tw-font-medium"
      >
        <div class="v2-canopas-gradient-text">
          Four reasons why we do conversation?
        </div>
      </div>
      <ul
        class="lg:tw-w-[75%] tw-m-auto tw-mt-6 tw-list-none tw-p-0 tw-text-[1.0625rem] tw-leading-[1.5rem] md:tw-text-[1.1875rem] md:tw-leading-[1.625rem] lg:tw-text-[1.375rem] lg:tw-leading-[1.75rem] tw-text-left tw-text-black-core/[.87] tw-font-medium"
      >
        <li
          v-for="reason in reasons"
          :key="reason"
          class="tw-py-2.5 tw-px-0 tw-flex"
        >
          <span
            class="tw-inline-block tw-mt-[5px] tw-mr-[24px] tw-w-[16px] tw-h-[16px] tw-rounded-[30%] tw-bg-gradient-to-b tw-from-[#F2709C] tw-to-[#FF9472] tw-content-none"
          ></span>
          <div class="tw-w-[90%]">
            {{ reason }}
          </div>
        </li>
      </ul>
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
};
</script>
