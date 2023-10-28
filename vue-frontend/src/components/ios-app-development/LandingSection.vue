<template>
  <div
    class="tw-bg-ios-landing-background tw-bg-opacity-10 tw-py-10 3xl:tw-outer-container lg:tw-mt-28"
  >
    <div
      class="tw-container tw-px-4 tw-pb-4 3xl:tw-pb-28 tw-text-left lg:tw-flex tw-mb-4 md:tw-mb-7 lg:tw-mb-14"
    >
      <div class="lg:tw-flex-1 tw-mb-6">
        <h1
          class="header-1 tw-mb-4 tw-pr-2 lg:tw-w-2/4 tw-text-black-core/[0.87]"
        >
          iOS app development company
        </h1>
        <p class="sub-h1-regular 2xl:tw-w-4/5 tw-text-black-core/[0.60]">
          We craft tailor-made iOS apps that redefine user experience and put
          you ahead of the game. Our team work with passion and precision to
          transform your business vision into reality, giving your brand the
          edge it needs in today's competitive marketplace.
        </p>
        <router-link
          class="tw-relative tw-z-9 tw-ml-0 lg:tw-mt-8 tw-hidden tw-h-[3.6rem] tw-w-[51%] lg:tw-block gradient-btn primary-btn"
          to="/contact"
          @click.native="mixpanel.track('tap_ios_development_cta')"
        >
          <span
            class="tw-px-[5px] tw-text-center sub-h3-semibold xl:sub-h4-semibold"
            >Consult Our iOS Team</span
          >
        </router-link>
      </div>
      <div
        class="tw-flex tw-flex-col lg:tw-flex-row tw-items-center tw-justify-center lg:tw-flex-1 tw-mt-8 xs:tw-mt-0"
      >
        <img
          :src="landing400"
          :srcset="`${landing400} 400w,${landing800} 800w`"
          class="tw-overflow-visible md:tw-mt-0 tw-h-[21.49113rem] xs:tw-h-[24.49113rem] lg:tw-hidden tw-w-[23.77063rem] tw-object-cover"
          alt="ios-landing"
        />
        <img
          v-if="width < 1400"
          :src="landingmd800"
          :srcset="`${landingmd800} 800w`"
          class="tw-overflow-visible tw-hidden lg:tw-block tw-h-full lg:!tw--mt-[0.5%] xl:!tw-mt-[0.5%] xl:tw- lg:!tw-h-[31rem] xl:!tw-h-[38rem] tw--ml-[82%] lg:tw--ml-[61%] xl:tw--ml-[59%] tw-w-[55%] tw-object-cover"
          alt="ios-landing"
        />
        <img
          v-else
          :src="landing1200"
          :srcset="`${landing1200} 1200w`"
          class="tw-overflow-visible tw-hidden lg:tw-block lg:tw-h-[40rem] tw--mt-[4%] xl:tw--ml-[80%] tw-object-cover"
          alt="ios-landing"
        />

        <div class="tw--mt-[5.3rem] tw-mr-28 xs:tw-mr-40 lg:tw-hidden">
          <router-link
            class="gradient-btn primary-btn"
            to="/contact"
            @click.native="mixpanel.track('tap_ios_development_cta')"
          >
            <span
              class="tw-px-[5px] tw-text-center sub-h3-semibold xl:sub-h4-semibold"
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
