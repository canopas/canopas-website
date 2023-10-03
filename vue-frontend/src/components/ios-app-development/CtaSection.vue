<template>
  <section>
    <div
      class="tw-relative tw-mb-16 tw-h-48 md:tw-h-72 xll:tw-h-[30rem] md:tw-hidden"
    >
      <img
        :src="BackgroundImage[0]"
        :srcset="`${BackgroundImage[0]} 400w`"
        alt="Background-Image"
        class="tw-h-full lg:tw-h-[22.5rem] xll:tw-h-full tw-w-full tw-object-cover"
        loading="lazy"
      />
      <div
        class="tw-container tw-absolute tw-top-4 tw-w-full tw-p-4 tw-left-[2%] sm:tw-left-[13%] xs:tw-p-8 tw-text-center sm:tw-w-[79%]"
      >
        <div class="tw-flex tw-w-full">
          <div class="tw-m-auto tw-text-white">
            <p
              class="tw-font-inter-semibold tw-text-[1.375rem] tw-leading-[2.0625rem] tw-tracking-[-4%] tw-text-black-core/[0.87]"
            >
              Let's Start a New Project Together
            </p>
          </div>
        </div>
        <div class="tw-mx-auto tw-mt-4">
          <router-link
            class="gradient-btn consultation-btn tw-m-0 tw-mx-auto tw-flex tw-text-[#FFF] tw-w-max tw-items-center tw-rounded-full tw-p-3 tw-px-7 tw-text-center"
            to="/contact"
            @click.native="mixpanel.track('tap_ios_development_cta')"
          >
            <span
              class="tw-mr-2.5 tw-font-inter-semibold tw-text-base tw-leading-[1.21rem] md:tw-block"
              >Contact Us
            </span>
          </router-link>
        </div>
      </div>
    </div>
    <!-- Desktop view start -->
    <div
      class="tw-relative tw-mb-16 tw-h-80 lg:tw-h-[476px] tw-hidden md:tw-block"
    >
      <img
        :src="BackgroundImage[1]"
        :srcset="`${BackgroundImage[1]} 800w,${BackgroundImage[2]} 1200w,${BackgroundImage[3]} 2400w`"
        alt="Background-Image"
        class="tw-h-full tw-w-full tw-object-cover"
        loading="lazy"
      />
      <div
        class="tw-container tw-absolute tw-left-1/2 tw-flex tw-w-full tw--translate-x-1/2 tw--translate-y-[15.55rem] lg:tw--translate-y-[24rem] tw-transform tw-flex-row tw-items-center"
      >
        <div class="tw-flex tw-w-[66%]">
          <div
            class="tw-w-[95%] tw-py-5 tw-text-white xl:tw-m-auto 2xl:tw-w-4/5"
          >
            <p
              class="tw-font-inter-bold tw-text-[2.188rem] tw-leading-[2.625rem] lg:tw-text-[3.4375rem] lg:tw-leading-[5.15625rem]"
            >
              Are you ready to take your business app to the next level?
            </p>
          </div>
        </div>
        <div class="tw-mx-auto tw-mt-4">
          <router-link
            class="gradient-btn consultation-btn tw-m-0 tw-mx-auto tw-flex tw-text-[#FFF] tw-w-max tw-items-center tw-rounded-full tw-p-3 tw-px-7 tw-text-center"
            to="/contact"
            @click.native="mixpanel.track('tap_ios_development_cta')"
          >
            <span
              class="tw-mr-2.5 tw-font-inter-medium tw-text-base tw-leading-[1.21rem] lg:tw-text-[1.1875rem] lg:tw-leading-[1.436875rem]"
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
