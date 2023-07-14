<template>
  <div class="tw-container tw-relative tw-flex tw-flex-row tw-py-[2rem]">
    <div class="tw-relative tw-z-[1]">
      <img
        :src="CtaBoyImage"
        class="tw-h-full tw-w-[14.1875rem] tw-object-cover md:!tw-w-[22.64594rem] md:tw-w-[25.1875rem] lg:!tw-w-[32.625rem] xl:!tw-w-[34.625rem] xs:tw-w-[13.1875rem]"
        alt="android-cta-image"
        loading="lazy"
      />
    </div>
    <div
      class="tw-border-[#00000]/[0.6] tw-absolute tw-left-[10%] tw-z-0 tw-mr-[0.2rem] tw-mt-[2rem] tw-flex tw-flex-col tw-items-center tw-justify-center tw-rounded-lg tw-border-2 tw-py-[0.5rem] md:!tw-mt-[3.7rem] lg:!tw-mt-[5.7rem] lg:!tw-py-[3rem] xl:!tw-mt-[7.8rem] xl:!tw-py-[1rem] 2xl:!tw-pb-[4rem] 2xl:!tw-pt-[3rem] 2xl:tw-pr-[2rem] xs:tw-mt-[2.8rem] xs:tw-py-4"
    >
      <div
        class="tw-ml-auto tw-w-[38%] tw-font-inter-semibold tw-text-[1.125rem] tw-leading-[1.4625rem] tw-text-black-core/[0.87] md:tw-text-[2.4375rem] md:tw-leading-[3.15625rem] lg:!tw-w-[47%] lg:tw-text-[2.935rem] lg:tw-leading-[4.15625rem] xl:tw-text-[3.4375rem] xl:tw-leading-[5.15625rem] 2xl:!tw-w-[60%] xs:tw-w-[50%]"
      >
        <span>Let’s talk about what we can build and scale together.</span>
      </div>
      <div
        class="tw-ml-auto tw-mt-4 sm:!tw-ml-[25%] md:!tw-ml-[23%] xl:!tw-ml-[19%] 2xl:!tw-ml-[-7%] xs:tw-ml-[30%]"
      >
        <router-link
          class="gradient-btn consultation-btn tw-m-0 tw-mx-auto tw-flex tw-w-max tw-items-center tw-rounded-full tw-p-3 tw-text-center"
          :to="contactURL"
          @click.native="mixpanel.track('tap_android_development_cta')"
        >
          <span
            class="tw-mr-2.5 tw-font-inter-medium tw-text-[1rem] tw-font-normal tw-leading-[1.1875rem] !tw-tracking-[0] md:tw-block md:tw-text-[1.09375rem] md:tw-text-[1.1875rem] md:tw-leading-[1.3125rem] md:tw-leading-[1.436875rem] lg:tw-text-[1.1875rem] lg:tw-leading-[1.4375rem]"
            >Contact Us
          </span>
        </router-link>
      </div>
    </div>
  </div>
</template>
<script>
import CtaBoyImage from "@/assets/images/andriod-app-development/cta/cta3.webp";

export default {
  data() {
    return {
      CtaBoyImage,
      contactURL: "/contact",
    };
  },

  inject: ["mixpanel"],
};
</script>
