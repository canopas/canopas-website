<template>
  <div
    class="tw-relative tw-h-[28.125rem] xll:tw-container 3xl:tw-outer-container md:!tw-h-[30.125rem] xl:md:!tw-h-[33.125rem] xs:tw-h-[23.125rem]"
  >
    <img
      :src="background400"
      :srcset="`${background400} 400w,${background800} 800w`"
      class="tw-h-full tw-w-full tw-object-cover xs:tw-hidden"
      alt="android-landing"
    />
    <img
      :src="background1200"
      :srcset="`${background1200} 1200w,${background2400} 2400w`"
      class="tw-hidden tw-h-full tw-w-full tw-object-cover xs:tw-block"
      alt="android-landing"
    />
    <div
      class="tw-container tw-absolute tw-p-4 tw-w-full tw-text-center tw-text-[#FFF] tw-transform tw-left-1/2 tw-top-1/2 tw--translate-x-1/2 tw--translate-y-1/2 xl:tw-w-[84%] md:tw-top-[49%] lg:tw-top-[48%] 2xl:tw-w-[69%] xll:tw-w-[65%]"
    >
      <h1 class="header-1 tw-text-white">
        Android app <br />
        development company
      </h1>
      <p class="tw-mt-4 md:tw-mt-6 sub-h1-regular tw-text-white/[0.80]">
        We create bespoke Android applications known for their supreme
        performance, unwavering security, and utmost reliability. With an
        unshakable commitment, our team curates unique, captivating apps to
        promote your business amidst fierce competition.
      </p>

      <div
        class="tw-mt-8 md:tw-mt-10 tw-flex tw-flex-wrap tw-justify-center md:tw-justify-center"
      >
        <router-link
          class="white-btn gradient-border-btn primary-btn tw-border-none"
          to="/contact"
          @click.native="mixpanel.track('tap_android_development_cta')"
        >
          <span class="sub-h3-semibold xl:sub-h4-semibold"
            >Consult Our Android Team</span
          >
        </router-link>
      </div>
    </div>
  </div>
</template>
<script>
import background400 from "@/assets/images/andriod-app-development/landing/background-400w.webp";
import background800 from "@/assets/images/andriod-app-development/landing/background-800w.webp";
import background1200 from "@/assets/images/andriod-app-development/landing/background-1200w.webp";
import background2400 from "@/assets/images/andriod-app-development/landing/background-2400w.webp";
export default {
  data() {
    return {
      background400,
      background800,
      background1200,
      background2400,
    };
  },
  inject: ["mixpanel"],
};
</script>
