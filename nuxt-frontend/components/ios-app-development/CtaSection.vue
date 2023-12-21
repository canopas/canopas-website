<template>
  <section>
    <div class="relative mb-16 h-[217px] md:h-[376px] lg:hidden">
      <img
        :src="BackgroundImage[0]"
        :srcset="`${BackgroundImage[0]} 400w`"
        alt="Background-Image"
        class="h-full lg:h-[22.5rem] xll:h-full w-full object-cover"
        loading="lazy"
      />
      <div
        class="container absolute top-0 md:top-[17%] w-full p-4 left-[2%] sm:left-[13%] xs:p-8 text-center sm:w-[79%]"
      >
        <div class="flex w-full">
          <div class="m-auto text-white">
            <p class="mobile-header-1 text-black-87">
              Let's Start a New Project Together
            </p>
          </div>
        </div>
        <div class="mx-auto mt-4">
          <router-link
            class="gradient-btn primary-btn m-0 mx-auto flex text-[#FFF] w-max items-center p-3 px-7 text-center"
            to="/contact"
            @click.native="$mixpanel.track('tap_ios_development_cta')"
          >
            <span class="mr-2.5 sub-h3-semibold lg:block">Contact Us </span>
          </router-link>
        </div>
      </div>
    </div>
    <!-- Desktop view start -->
    <div
      class="relative mb-16 lg:h-[476px] 3xl:outer-container hidden lg:block"
    >
      <img
        :src="BackgroundImage[1]"
        :srcset="`${BackgroundImage[1]} 800w,${BackgroundImage[2]} 1200w,${BackgroundImage[3]} 2400w`"
        alt="Background-Image"
        class="h-full w-full object-cover"
        loading="lazy"
      />
      <div
        class="container absolute left-1/2 flex w-full -translate-x-1/2 lg:-translate-y-[22rem] xl:-translate-y-[23rem] transform flex-row items-center"
      >
        <div class="flex w-[66%]">
          <div class="w-[95%] py-5 text-white xl:m-auto 2xl:w-4/5">
            <p class="desk-header-1">
              Are you ready to take your business app to the next level?
            </p>
          </div>
        </div>
        <div class="mx-auto mt-4">
          <router-link
            class="gradient-btn primary-btn m-0 mx-auto flex text-[#FFF] w-max items-center p-3 px-7"
            to="/contact"
            @click.native="$mixpanel.track('tap_ios_development_cta')"
          >
            <span class="sub-h4-semibold">Let's Chat </span>
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
