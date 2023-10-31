<template>
  <div
    class="relative h-[28.125rem] xll:container 3xl:outer-container md:!h-[30.125rem] xl:md:!h-[33.125rem] xs:h-[23.125rem]"
  >
    <img
      :src="background400"
      :srcset="`${background400} 400w,${background800} 800w`"
      class="h-full w-full object-cover xs:hidden"
      alt="android-landing"
    />
    <img
      :src="background1200"
      :srcset="`${background1200} 1200w,${background2400} 2400w`"
      class="hidden h-full w-full object-cover xs:block"
      alt="android-landing"
    />
    <div
      class="container absolute p-4 w-full text-center text-[#FFF] transform left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2 xl:w-[84%] md:top-[49%] lg:top-[48%] 2xl:w-[69%] xll:w-[65%]"
    >
      <h1 class="header-1 text-white">
        Android app <br />
        development company
      </h1>
      <p class="mt-4 md:mt-6 sub-h1-regular text-white/[0.80]">
        We create bespoke Android applications known for their supreme
        performance, unwavering security, and utmost reliability. With an
        unshakable commitment, our team curates unique, captivating apps to
        promote your business amidst fierce competition.
      </p>

      <div
        class="mt-8 md:mt-10 flex flex-wrap justify-center md:justify-center"
      >
        <nuxt-link
          class="white-btn gradient-border-btn primary-btn border-none"
          to="/contact"
          @click.native="$mixpanel.track('tap_android_development_cta')"
        >
          <span class="sub-h3-semibold xl:sub-h4-semibold"
            >Consult Our Android Team</span
          >
        </nuxt-link>
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
