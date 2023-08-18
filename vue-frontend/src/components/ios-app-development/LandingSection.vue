<template>
  <div
    class="tw-bg-gradient-to-t tw-from-[#F77D8D]/[31.07%] md:tw-from-[#F77D8D]/[27.59%] tw-to-[#FFFFFF] tw-pb-12"
  >
    <div class="tw-container tw-p-4 tw-text-left md:tw-flex xl:tw-my-8">
      <div class="md:tw-flex-1 lg:tw-mt-16">
        <h1
          class="md:!tw--tracking-[2.4px] tw-pr-2 tw-mb-2 lg:tw-w-2/4 sm:tw-mb-0 tw-font-inter-semibold md:tw-font-inter-bold md:tw-text-[2.75rem] xl:tw-text-[3.75rem] tw-text-[2rem] md:tw-leading-[3.625rem] xl:tw-leading-[5.625rem] tw-leading-[2.6rem] tw-text-black-core/[0.87] md:tw-font-inter-bold"
        >
          iOS App Development Company
        </h1>
        <p
          class="tw-font-inter-medium md:tw-font-inter-regular tw-pt-2 xl:tw-text-[1.25rem] tw-text tw-text-[1rem] xl:tw-leading-[1.875rem] 2xl:tw-w-4/5 tw-leading-normal tw-text-black-core/[0.87]"
        >
          We craft tailor-made iOS apps that redefine user experience and put
          you ahead of the game. Our team work with passion and precision to
          transform your business vision into reality, giving your brand the
          edge it needs in today's competitive marketplace.
        </p>

        <router-link
          class="tw-relative tw-z-10 tw-ml-0 xl:tw-w-[290px] tw-hidden md:tw-block tw-my-8 lg:tw-mt-8 lg:tw-mb-4 xl:tw-my-8 gradient-btn consultation-btn tw-h-[3.3rem] tw-w-max tw-items-left tw-rounded-full"
          to="/contact"
          @click.native="mixpanel.track('tap_ios_development_cta')"
        >
          <span
            class="tw-font-inter-semibold tw-text-center tw-text-[1.1875rem] tw-leading-[0.425rem]"
            >Consult Our iOS Team</span
          >
        </router-link>
      </div>
      <div
        class="tw-flex tw-flex-col md:tw-flex-row tw-items-center tw-justify-center md:tw-flex-1"
      >
        <img
          :src="landing400"
          :srcset="`${landing400} 400w,${landing800} 800w`"
          class="tw-mt-7 tw-overflow-visible md:tw-mt-0 tw-h-[21.49113rem] xs:tw-h-[24.49113rem] md:tw-hidden tw-w-[23.77063rem] tw-object-cover"
          alt="ios-landing"
        />
        <img
          v-if="width < 1400"
          :src="landingmd800"
          :srcset="`${landingmd800} 800w`"
          class="tw-overflow-visible tw-hidden md:tw-block tw-h-full xl:!tw-h-full lg:!tw-mt-[10%] lg:!tw-h-[98%] tw--ml-[82%] lg:tw--ml-[52%] xl:tw--ml-[70%] tw-w-[55%] tw-object-cover"
          alt="ios-landing"
        />
        <img
          v-else
          :src="landing1200"
          :srcset="`${landing1200} 1200w`"
          class="tw-overflow-visible tw-hidden md:tw-block tw-h-[37.5rem] tw-mt-8 xl:tw--ml-[70%] tw-w-[55%] tw-object-cover"
          alt="ios-landing"
        />

        <div
          class="tw--mt-[2.7rem] tw-mr-28 xs:tw--mt-14 xs:tw-mr-40 md:tw-hidden"
        >
          <router-link
            class="gradient-btn consultation-btn tw-w-max tw-items-left tw-rounded-full"
            to="/contact"
            @click.native="mixpanel.track('tap_ios_development_cta')"
          >
            <span
              class="tw-font-inter-medium tw-px-[5px] tw-text-center tw-text-[1rem] tw-leading-[1.21rem]"
              >Consult Our iOS Team</span
            >
          </router-link>
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
