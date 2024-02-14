<template>
  <div class="relative mt-16 lg:mt-60 overflow-hidden px-0 3xl:outer-container">
    <img
      :src="bg2400"
      :srcset="`${bg400} 400w, ${bg800} 800w, ${bg1200} 1200w, ${bg2400} 2400w`"
      class="absolute top-0 left-0 w-full -z-[1]"
      alt="canopas-cta"
    />
    <div class="container m-auto pt-14 lg:pt-28 text-center">
      <span class="w-full mobile-header-2 lg:desk-header-2 text-black-87">
        How do you know we’re a good fit for you?
      </span>
      <h2
        class="mt-4 lg:mt-6 mobile-header-3-regular lg:mobile-header-2-regular text-black-60"
      >
        Let’s have a short chat.
      </h2>
      <div class="mt-2 lg:mt-4 flex flex-wrap justify-center">
        <nuxt-link
          class="gradient-btn primary-btn"
          to="/contact"
          @click.native="$mixpanel.track('tap_landing_cta')"
        >
          <span class="sub-h3-semibold lg:sub-h1-semibold"
            >Get In Touch
          </span></nuxt-link
        >
      </div>
    </div>
    <div
      class="container mt-10 lg:mt-[4.5rem] text-left mb-16 lg:mb-60 lg:translate-x-[15%] 2xl:translate-x-[22%]"
    >
      <div class="sub-h1-semibold lg:sub-h2-medium text-black-60">
        Four reasons why do we do this call?
      </div>
      <ul
        class="lg:w-3/4 2xl:w-[64%] mt-4 list-none p-0 sub-h3-regular lg:mobile-header-3-regular text-black-87"
      >
        <li
          v-for="reason in reasons"
          :key="reason"
          class="flex pb-4 md:pb-6 px-0"
        >
          <span
            class="inline-block mt-1.5 mr-[0.938rem] md:mr-[1.125rem] w-3 h-3 md:w-[0.938rem] md:h-[0.938rem] rounded-sm md:rounded-[5px] bg-primary-1 content-none"
          ></span>
          <div class="w-[90%]">
            {{ reason }}
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>

<script type="module">
import bg400 from "@/assets/images/CTA/bg/400.webp";
import bg800 from "@/assets/images/CTA/bg/800.webp";
import bg1200 from "@/assets/images/CTA/bg/1200.webp";
import bg2400 from "@/assets/images/CTA/bg/2400.webp";

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
