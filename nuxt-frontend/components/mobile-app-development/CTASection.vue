<template>
  <section
    class="relative my-[50px] bg-[#FDE0E2] bg-gradient-background md:mt-0 md:bg-none"
  >
    <!-- MOBILE UI START  -->
    <div class="md:hidden">
      <img
        :src="background400w"
        :srcset="`${background400w} 400w, ${background800w} 800w`"
        loading="lazy"
        class="absolute h-full w-full"
        alt="background-image"
      />
      <div
        class="container sticky flex flex-col items-center justify-center py-20"
      >
        <p
          class="v2-canopas-gradient-text flex text-center font-inter-medium text-[1rem] leading-[1.21rem]"
        >
          Get started today
        </p>
        <p
          class="mt-6 text-center font-inter-bold text-[2.25rem] leading-[2.723rem] text-black-core/[0.85]"
        >
          Let's build the next big thing!
        </p>
        <p
          class="mt-4 text-center font-inter-regular text-[0.875rem] leading-[1.313rem] text-black-core/[0.75]"
        >
          Let's improve your business's digital strategy and implement robust
          mobile apps to achieve your business objectives. Schedule Your Free
          Consultation Now.
        </p>
        <nuxt-link
          class="gradient-btn consultation-btn mt-10 flex w-max items-center rounded-full py-4 text-center"
          to="/contact"
          @click.native="$mixpanel.track('tap_mobileapp_development_cta1')"
        >
          <span
            class="mr-2.5 font-inter-semibold text-[1.1875rem] leading-[1.1875rem] md:text-[1.09375rem] md:leading-[1.3125rem] lg:text-[1.1875rem] lg:leading-[1.4375rem]"
            >Get Free Consultation
          </span>
        </nuxt-link>
      </div>
    </div>
    <!-- MOBILE UI END -->

    <!-- DESKTOP UI START -->
    <div class="hidden md:block">
      <img
        :src="desktop_background800w"
        :srcset="`${desktop_background800w} 800w, ${desktop_background1200w} 1200w, ${desktop_background1600w} 1600w, ${desktop_background2400w} 2400w `"
        loading="lazy"
        class="absolute h-full w-full"
        alt="background-image"
      />
      <div
        class="container sticky z-[1] flex flex-col items-center justify-center py-20 xll:py-32"
      >
        <p
          class="v2-canopas-gradient-text text-center font-inter-medium text-[1.125rem] leading-[1.375rem] xl:text-[1.25rem] xl:leading-[1.513rem]"
        >
          Get started today
        </p>
        <p
          class="mt-4 text-center font-inter-bold text-[2.813rem] leading-[3.438rem] text-black-core/[0.85] xl:text-[3.438rem] xl:leading-[4.16rem]"
        >
          Let's build the next big thing!
        </p>
        <p
          class="mt-4 w-[75%] text-center font-inter-regular text-[1.063rem] leading-[1.625rem] text-black-core/[0.75] xl:text-[1.25rem] xl:leading-[1.875rem]"
        >
          Let's improve your business's digital strategy and implement robust
          mobile apps to achieve your business objectives. Schedule Your Free
          Consultation Now.
        </p>
        <nuxt-link
          class="gradient-btn consultation-btn mt-10 flex w-max items-center rounded-full py-[1rem] text-center"
          to="/contact"
          @click.native="$mixpanel.track('tap_mobileapp_development_cta1')"
        >
          <span
            class="mr-2.5 font-inter-semibold text-[1.1875rem] leading-[1.1875rem] md:text-[1.09375rem] md:leading-[1.3125rem] lg:text-[1.1875rem] lg:leading-[1.4375rem]"
            >Get Free Consultation
          </span>
        </nuxt-link>
      </div>
    </div>
    <!-- DESKTOP UI END -->
  </section>
</template>
<script>
import background400w from "@/assets/images/mobile-app-development/cta-1/mobile/background-400w.webp";
import background800w from "@/assets/images/mobile-app-development/cta-1/mobile/background-800w.webp";
import desktop_background800w from "@/assets/images/mobile-app-development/cta-1/desktop/background-800w.webp";
import desktop_background1200w from "@/assets/images/mobile-app-development/cta-1/desktop/background-1200w.webp";
import desktop_background1600w from "@/assets/images/mobile-app-development/cta-1/desktop/background-1600w.webp";
import desktop_background2400w from "@/assets/images/mobile-app-development/cta-1/desktop/background-2400w.webp";

export default {
  data() {
    return {
      background400w,
      background800w,
      desktop_background800w,
      desktop_background1200w,
      desktop_background1600w,
      desktop_background2400w,
    };
  },
  inject: ["mixpanel"],
};
</script>
