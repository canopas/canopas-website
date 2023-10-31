<template>
  <div
    class="relative mt-[4.375rem] md:mt-[14.063rem] overflow-hidden px-0 font-inter-regular md:font-inter-medium"
  >
    <img
      :src="bg2400"
      :srcset="`${bg400} 400w, ${bg800} 800w, ${bg1200} 1200w, ${bg2400} 2400w`"
      class="absolute top-0 left-0 w-full -z-[1]"
      alt="canopas-cta"
    />
    <div
      class="container m-auto md:w-[86%] lg:w-4/5 xl:w-[86%] px-[25px] pt-[61px] md:pt-[162px] pb-[50px] md:pb-[70px] opacity-1 text-center"
    >
      <div class="md:w-[86%] lg:w-4/5 xl:w-[86%] m-auto">
        <span
          style="-webkit-box-decoration-break: clone"
          class="w-full text-[2.1875rem] leading-[2.75rem] tracking-[1px] md:text-[2.8125rem] md:leading-[3.4375rem] lg:text-[3.4375rem] lg:leading-[4.0281rem] font-roboto-bold canopas-gradient-text"
        >
          How do you know we’re a good fit for you?
        </span>
        <h2
          class="mt-6 text-[1.25rem] leading-[1.5625rem] md:text-[1.375rem] md:leading-[1.6875rem] lg:text-[1.5rem] lg:leading-[1.8156rem] tracking-normal text-black-core/[.87]"
        >
          Let’s have a short chat.
        </h2>
        <div class="flex flex-wrap justify-center">
          <nuxt-link
            class="flex w-max items-center rounded-full py-3 px-3 gradient-btn consultation-btn"
            to="/contact"
            @click.native="$mixpanel.track('tap_home_touch_cta')"
          >
            <span
              class="mr-2.5 font-normal text-[1rem] leading-[1.1875rem] md:text-[1.09375rem] md:leading-[1.3125rem] lg:text-[1.1875rem] lg:leading-[1.4375rem] font-inter-medium !tracking-[0]"
              >Get in touch</span
            >
          </nuxt-link>
        </div>
        <div
          class="mt-12 w-full text-[1.25rem] leading-[1.5625rem] md:text-[1.375rem] md:leading-[1.6875rem] lg:text-[1.5rem] lg:leading-[1.8156rem] text-left"
        >
          <div class="v2-canopas-gradient-text">
            Four reasons why do we do this call?
          </div>
        </div>
        <ul
          class="w-full mt-4 list-none p-0 text-[1rem] leading-[1.1875rem] md:text-[1.1875rem] md:leading-[1.625rem] lg:text-[1.375rem] lg:leading-[1.75rem] text-left text-black-core/[.87]"
        >
          <li
            v-for="reason in reasons"
            :key="reason"
            class="flex py-[7.5px] md:py-[17.5px] px-0"
          >
            <span
              class="inline-block mt-[5px] mr-[15px] md:mr-[18px] w-3 h-3 md:w-[15px] md:h-[15px] rounded-[2px] md:rounded-[5px] bg-gradient-to-b from-[#F2709C] to-[#FF9472] content-none"
            ></span>
            <div class="w-[90%]">
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
      bg400,
      bg800,
      bg1200,
      bg2400,
    };
  },
  inject: ["mixpanel"],
};
</script>
