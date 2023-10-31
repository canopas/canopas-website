<template>
  <div
    class="bg-gradient-to-t from-[#F77D8D]/[31.07%] md:from-[#F77D8D]/[27.59%] to-[#FFFFFF] pb-12"
  >
    <div class="container p-4 text-left md:flex xl:my-8">
      <div class="md:flex-1 lg:mt-16">
        <h1
          class="md:!-tracking-[2.4px] pr-2 mb-2 lg:w-2/4 sm:mb-0 font-inter-semibold md:font-inter-bold md:text-[2.75rem] xl:text-[3.75rem] text-[2rem] md:leading-[3.625rem] xl:leading-[5.625rem] leading-[2.6rem] text-black-core/[0.87] md:font-inter-bold"
        >
          iOS App Development Company
        </h1>
        <p
          class="font-inter-medium md:font-inter-regular pt-2 xl:text-[1.25rem] text text-[1rem] xl:leading-[1.875rem] 2xl:w-4/5 leading-normal text-black-core/[0.87]"
        >
          We craft tailor-made iOS apps that redefine user experience and put
          you ahead of the game. Our team work with passion and precision to
          transform your business vision into reality, giving your brand the
          edge it needs in today's competitive marketplace.
        </p>

        <nuxt-link
          class="relative z-9 ml-0 xl:w-[290px] hidden md:block my-8 lg:mt-8 lg:mb-4 xl:my-8 gradient-btn consultation-btn h-[3.3rem] w-max items-left rounded-full"
          to="/contact"
          @click.native="$mixpanel.track('tap_ios_development_cta')"
        >
          <span
            class="font-inter-semibold text-center text-[1.1875rem] leading-[0.425rem]"
            >Consult Our iOS Team</span
          >
        </nuxt-link>
      </div>
      <div
        class="flex flex-col md:flex-row items-center justify-center md:flex-1"
      >
        <img
          :src="landing400"
          :srcset="`${landing400} 400w,${landing800} 800w`"
          class="mt-7 overflow-visible md:mt-0 h-[21.49113rem] xs:h-[24.49113rem] md:hidden w-[23.77063rem] object-cover"
          alt="ios-landing"
        />
        <img
          v-if="width < 1400"
          :src="landingmd800"
          :srcset="`${landingmd800} 800w`"
          class="overflow-visible hidden md:block h-full xl:!h-full lg:!mt-[10%] lg:!h-[98%] -ml-[82%] lg:-ml-[52%] xl:-ml-[70%] w-[55%] object-cover"
          alt="ios-landing"
        />
        <img
          v-else
          :src="landing1200"
          :srcset="`${landing1200} 1200w`"
          class="overflow-visible hidden md:block h-[37.5rem] mt-8 xl:-ml-[70%] w-[55%] object-cover"
          alt="ios-landing"
        />

        <div class="-mt-[2.7rem] mr-28 xs:-mt-14 xs:mr-40 md:hidden">
          <nuxt-link
            class="gradient-btn consultation-btn w-max items-left rounded-full"
            to="/contact"
            @click.native="$mixpanel.track('tap_ios_development_cta')"
          >
            <span
              class="font-inter-medium px-[5px] text-center text-[1rem] leading-[1.21rem]"
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
