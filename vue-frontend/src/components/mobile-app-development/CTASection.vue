<template>
  <section
    class="tw-relative tw-my-[50px] tw-bg-[#FDE0E2] tw-bg-gradient-background md:tw-mt-0 md:tw-bg-none"
  >
    <!-- MOBILE UI START  -->
    <div class="md:tw-hidden">
      <img
        :src="background400w"
        :srcset="`${background400w} 400w, ${background800w} 800w`"
        loading="lazy"
        class="tw-absolute tw-h-full tw-w-full"
        alt="background-image"
      />
      <div
        class="tw-container tw-sticky tw-flex tw-flex-col tw-items-center tw-justify-center tw-py-20"
      >
        <p
          class="v2-canopas-gradient-text tw-flex tw-text-center tw-font-inter-medium tw-text-[1rem] tw-leading-[1.21rem]"
        >
          Get started today
        </p>
        <p
          class="tw-mt-6 tw-text-center tw-font-inter-bold tw-text-[2.25rem] tw-leading-[2.723rem] tw-text-black-core/[0.85]"
        >
          Let's build the next big thing!
        </p>
        <p
          class="tw-mt-4 tw-text-center tw-font-inter-regular tw-text-[0.875rem] tw-leading-[1.313rem] tw-text-black-core/[0.75]"
        >
          Let's improve your business's digital strategy and implement robust
          mobile apps to achieve your business objectives. Schedule Your Free
          Consultation Now.
        </p>
        <router-link
          class="gradient-btn consultation-btn tw-mt-10 tw-flex tw-w-max tw-items-center tw-rounded-full tw-py-4 tw-text-center"
          to="/contact"
          @click.native="mixpanel.track('tap_mobileapp_development_cta1')"
        >
          <span
            class="tw-mr-2.5 tw-font-inter-semibold tw-text-[1.1875rem] tw-leading-[1.1875rem] md:tw-text-[1.09375rem] md:tw-leading-[1.3125rem] lg:tw-text-[1.1875rem] lg:tw-leading-[1.4375rem]"
            >Get Free Consultation
          </span>
        </router-link>
      </div>
    </div>
    <!-- MOBILE UI END -->

    <!-- DESKTOP UI START -->
    <div class="tw-hidden md:tw-block">
      <img
        :src="desktop_background800w"
        :srcset="`${desktop_background800w} 800w, ${desktop_background1200w} 1200w, ${desktop_background1600w} 1600w, ${desktop_background2400w} 2400w `"
        loading="lazy"
        class="tw-absolute tw-h-full tw-w-full"
        alt="background-image"
      />
      <div
        class="tw-container tw-sticky tw-z-[1] tw-flex tw-flex-col tw-items-center tw-justify-center tw-py-20 xll:tw-py-32"
      >
        <p
          class="v2-canopas-gradient-text tw-text-center tw-font-inter-medium tw-text-[1.125rem] tw-leading-[1.375rem] xl:tw-text-[1.25rem] xl:tw-leading-[1.513rem]"
        >
          Get started today
        </p>
        <p
          class="tw-mt-4 tw-text-center tw-font-inter-bold tw-text-[2.813rem] tw-leading-[3.438rem] tw-text-black-core/[0.85] xl:tw-text-[3.438rem] xl:tw-leading-[4.16rem]"
        >
          Let's build the next big thing!
        </p>
        <p
          class="tw-mt-4 tw-w-[75%] tw-text-center tw-font-inter-regular tw-text-[1.063rem] tw-leading-[1.625rem] tw-text-black-core/[0.75] xl:tw-text-[1.25rem] xl:tw-leading-[1.875rem]"
        >
          Let's improve your business's digital strategy and implement robust
          mobile apps to achieve your business objectives. Schedule Your Free
          Consultation Now.
        </p>
        <router-link
          class="gradient-btn consultation-btn tw-mt-10 tw-flex tw-w-max tw-items-center tw-rounded-full tw-py-[1rem] tw-text-center"
          to="/contact"
          @click.native="mixpanel.track('tap_mobileapp_development_cta1')"
        >
          <span
            class="tw-mr-2.5 tw-font-inter-semibold tw-text-[1.1875rem] tw-leading-[1.1875rem] md:tw-text-[1.09375rem] md:tw-leading-[1.3125rem] lg:tw-text-[1.1875rem] lg:tw-leading-[1.4375rem]"
            >Get Free Consultation
          </span>
        </router-link>
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
