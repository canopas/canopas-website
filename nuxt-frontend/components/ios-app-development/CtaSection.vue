<template>
  <section>
    <div class="relative mb-16 h-48 md:h-72 xll:h-[30rem] md:hidden">
      <img
        :src="BackgroundImage[0]"
        :srcset="`${BackgroundImage[0]} 400w`"
        alt="Background-Image"
        class="h-full lg:h-[22.5rem] xll:h-full w-full object-cover"
        loading="lazy"
      />
      <div
        class="container absolute top-4 w-full p-4 left-[2%] sm:left-[13%] xs:p-8 text-center sm:w-[79%]"
      >
        <div class="flex w-full">
          <div class="m-auto text-white">
            <p
              class="font-inter-semibold text-[1.375rem] leading-[2.0625rem] tracking-[-0.04rem] text-black-core/[0.87]"
            >
              Let's Start a New Project Together
            </p>
          </div>
        </div>
        <div class="mx-auto mt-4">
          <router-link
            class="gradient-btn consultation-btn m-0 mx-auto flex text-[#FFF] w-max items-center rounded-full p-3 px-7 text-center"
            to="/contact"
            @click.native="$mixpanel.track('tap_ios_development_cta')"
          >
            <span
              class="mr-2.5 font-inter-semibold text-base leading-[1.21rem] md:block"
              >Contact Us
            </span>
          </router-link>
        </div>
      </div>
    </div>
    <!-- Desktop view start -->
    <div class="relative mb-16 h-80 lg:h-[476px] hidden md:block">
      <img
        :src="BackgroundImage[1]"
        :srcset="`${BackgroundImage[1]} 800w,${BackgroundImage[2]} 1200w,${BackgroundImage[3]} 2400w`"
        alt="Background-Image"
        class="h-full w-full object-cover"
        loading="lazy"
      />
      <div
        class="container absolute left-1/2 flex w-full -translate-x-1/2 -translate-y-[15.55rem] lg:-translate-y-[24rem] transform flex-row items-center"
      >
        <div class="flex w-[66%]">
          <div class="w-[95%] py-5 text-white xl:m-auto 2xl:w-4/5">
            <p
              class="font-inter-bold text-[2.188rem] leading-[2.625rem] lg:text-[3.4375rem] lg:leading-[5.15625rem]"
            >
              Are you ready to take your business app to the next level?
            </p>
          </div>
        </div>
        <div class="mx-auto mt-4">
          <router-link
            class="gradient-btn consultation-btn m-0 mx-auto flex text-[#FFF] w-max items-center rounded-full p-3 px-7 text-center"
            to="/contact"
            @click.native="$mixpanel.track('tap_ios_development_cta')"
          >
            <span
              class="mr-2.5 font-inter-medium text-base leading-[1.21rem] lg:text-[1.1875rem] lg:leading-[1.436875rem]"
              >Let's Chat
            </span>
          </router-link>
        </div>
      </div>
    </div>
  </section>
</template>

<script type="module">
import background400 from "@/assets/images/ios-app-development/cta/mobile.webp";
import background800 from "@/assets/images/ios-app-development/cta/desktopCta-1-800w.webp";
import background1200 from "@/assets/images/ios-app-development/cta/desktopCta-1-1200w.webp";
import background2400 from "@/assets/images/ios-app-development/cta/desktopCta-1-2400w.webp";

export default {
  data() {
    return {
      BackgroundImage: [
        background400,
        background800,
        background1200,
        background2400,
      ],
    };
  },
  inject: ["mixpanel"],
};
</script>
