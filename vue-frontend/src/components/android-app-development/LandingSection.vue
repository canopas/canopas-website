<template>
  <div class="xll:tw-container lg:tw-relative">
    <img
      :src="background400"
      :srcset="`${background400} 400w,${background800} 800w`"
      class="tw-relative tw-h-full tw-w-full tw-object-contain sm:tw-hidden"
      alt="android-landing"
    />

    <aspect-ratio height="70%" class="tw-hidden sm:tw-block">
      <img
        :src="background1200"
        :srcset="`${background1200} 1200w,${background2400} 2400w`"
        class="tw-h-full tw-w-full tw-object-cover"
        alt="android-landing"
      />
    </aspect-ratio>

    <div
      class="tw-container tw-absolute tw-left-0 tw-top-[8rem] tw-w-[75%] tw-flex-col tw-items-center tw-justify-center tw-p-4 tw-text-left tw-text-[#FFF] md:tw-left-1/2 md:tw-top-[38%] md:tw--translate-x-1/2 md:tw--translate-y-1/2 md:tw-transform md:tw-text-center lg:tw-top-1/2 xl:tw-w-[60%] 2xl:tw-top-[40%] 2xl:tw-w-[49%] xl2:tw-w-[55%] 2xll:tw-w-[45%] xll:tw-top-[45%] xll:tw-w-[61%]"
    >
      <h1
        class="tracking-[0.08rem] tw-mb-2 tw-font-inter-semibold tw-text-[2rem] tw-leading-[2.6rem] tw-text-[#FFF] sm:tw-mb-[1.5rem] lg:tw-font-inter-bold lg:tw-text-[3.75rem] lg:tw-leading-[5.625rem]"
      >
        Android App<br />
        Development Company
      </h1>
      <p
        class="tw-font-inter-medium tw-text-[1rem] tw-text-[1rem] tw-leading-[1.5rem] tw-text-[#FFF] lg:tw-text-[1.25rem] lg:tw-leading-[1.875rem]"
      >
        We create bespoke Android applications known for their supreme
        performance, unwavering security, and utmost reliability. With an
        unshakable commitment, our team curates unique, captivating apps to
        promote your business amidst fierce competition.
      </p>

      <div
        class="tw-justify-left tw-mt-[1rem] tw-flex tw-flex-wrap md:tw-justify-center xs:tw-my-[2rem]"
      >
        <router-link
          class="white-btn gradient-border-btn tw-flex tw-h-[3.3rem] tw-w-max tw-items-center tw-rounded-full tw-border-none"
          to="/contact"
          @click.native="mixpanel.track('tap_android_development_cta')"
        >
          <span
            class="tw-font-inter-semibold tw-text-[1rem] tw-font-normal tw-leading-[1.21rem] !tw-tracking-[0] md:tw-text-[1.09375rem] md:tw-text-[1.1875rem] md:tw-leading-[1.3125rem] md:tw-leading-[1.4375rem] lg:tw-text-[1.1875rem] lg:tw-leading-[1.4375rem]"
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
import AspectRatio from "@/components/utils/AspectRatio.vue";
export default {
  data() {
    return {
      background400,
      background800,
      background1200,
      background2400,
    };
  },
  components: {
    AspectRatio,
  },
  inject: ["mixpanel"],
};
</script>
