<template>
  <div
    class="bg-ios-landing-background bg-opacity-10 py-10 3xl:outer-container lg:mt-28"
  >
    <div
      class="container px-4 pb-4 3xl:pb-28 text-left lg:flex mb-4 md:mb-7 lg:mb-14"
    >
      <div class="lg:flex-1 mb-6">
        <h1 class="header-1 mb-4 pr-2 lg:w-2/4 text-black-core/[0.87]">
          iOS app development company
        </h1>
        <p class="sub-h1-regular 2xl:w-4/5 text-black-core/[0.60]">
          We craft tailor-made iOS apps that redefine user experience and put
          you ahead of the game. Our team work with passion and precision to
          transform your business vision into reality, giving your brand the
          edge it needs in today's competitive marketplace.
        </p>
        <div
          class="relative z-9 lg:mt-[2.2rem] xl:mt-[2.5rem] 2xl:mt-[2.4rem] hidden lg:block"
        >
          <nuxt-link
            class="ml-0 gradient-btn primary-btn"
            to="/contact"
            @click.native="$mixpanel.track('tap_ios_development_cta')"
          >
            <span class="sub-h3-semibold xl:sub-h4-semibold"
              >Consult Our iOS Team</span
            >
          </nuxt-link>
        </div>
      </div>
      <div
        class="flex flex-col lg:flex-row items-center justify-center lg:flex-1 mt-8 xs:mt-0"
      >
        <img
          :src="landing400"
          :srcset="`${landing400} 400w,${landing800} 800w`"
          class="overflow-visible md:mt-0 h-[21.49113rem] xs:h-[24.49113rem] lg:hidden w-[23.77063rem] object-cover"
          alt="ios-landing"
        />
        <img
          v-if="width < 1400"
          :src="landingmd800"
          :srcset="`${landingmd800} 800w`"
          class="overflow-visible hidden lg:block h-full lg:!-mt-[0.5%] xl:!mt-[0.5%] lg:!h-[31rem] xl:!h-[38rem] -ml-[82%] lg:-ml-[61%] xl:-ml-[59%] w-[55%] object-cover"
          alt="ios-landing"
        />
        <img
          v-else
          :src="landing1200"
          :srcset="`${landing1200} 1200w`"
          class="overflow-visible hidden lg:block lg:h-[40rem] -mt-[4%] xl:-ml-[80%] object-cover"
          alt="ios-landing"
        />

        <div class="-mt-[5.3rem] mr-28 xs:mr-40 lg:hidden">
          <nuxt-link
            class="gradient-btn primary-btn"
            to="/contact"
            @click.native="$mixpanel.track('tap_ios_development_cta')"
          >
            <span
              class="px-[5px] text-center sub-h3-semibold xl:sub-h4-semibold"
              >Consult Our iOS Team</span
            >
          </nuxt-link>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import landing400 from "@/assets/images/ios-app-development/landing/landing-400w.webp";
import landing800 from "@/assets/images/ios-app-development/landing/landing-800w.webp";
import landingmd800 from "@/assets/images/ios-app-development/landing/landing-md-800w.webp";
import landing1200 from "@/assets/images/ios-app-development/landing/landing-1200w.webp";

export default {
  data() {
    return {
      width: 1400,
      landing400,
      landing800,
      landingmd800,
      landing1200,
    };
  },
  mounted() {
    this.width = window.innerWidth;
  },
  inject: ["mixpanel"],
};
</script>
