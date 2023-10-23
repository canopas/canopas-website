<template>
  <section class="tw-container tw-py-16 md:tw-py-40 xl:tw-py-60">
    <!-- ---------------------MOBILE UI START------------------ -->
    <div class="lg:tw-hidden tw-relative tw-flex tw-flex-col">
      <img
        :src="CtaBoyImage"
        class="tw-z-[1] tw-mx-auto tw-h-full tw-w-[14.1875rem] md:tw-w-64 tw-object-cover xs:tw-w-[13.1875rem]"
        alt="android-cta-image"
        loading="lazy"
      />
      <div
        class="tw-w-[80%] tw-mx-auto sm:tw-w-full tw-text-center tw-mt-8 tw-text-black-core/[0.87] header-1"
      >
        Let’s talk about what we can build and scale together.
      </div>
      <div class="tw-mx-auto tw-mt-8 tw-z-[1]">
        <router-link
          class="gradient-btn primary-btn"
          to="/contact"
          @click.native="mixpanel.track('tap_android_development_cta')"
        >
          <span class="sub-h3-semibold">Contact Us </span>
        </router-link>
      </div>
      <div
        class="tw-top-12 md:tw-top-14 tw-absolute tw-inset-0 tw-z-[0] tw-mx-auto tw-h-[85%] sm:tw-h-[83%] tw-w-[100%] tw-rounded-[0.625rem] tw-border tw-border-black/[0.87]"
      ></div>
    </div>
    <!-- ---------------------MOBILE UI END------------------ -->

    <!-- ---------------------DESKTOP UI START------------------ -->
    <div class="tw-hidden tw-relative lg:tw-flex tw-flex-row">
      <div class="tw-relative tw-z-[1]">
        <img
          :src="CtaBoyImage"
          class="tw-h-full tw-w-[14.1875rem] tw-object-cover !tw-w-[32.625rem] xl:!tw-w-[34.625rem] xs:tw-w-[13.1875rem]"
          alt="android-cta-image"
          loading="lazy"
        />
      </div>
      <div
        class="tw-h-[82%] lg:tw-h-[80%] xl:tw-h-[75%] tw-border-[#00000]/[0.6] tw-absolute tw-left-[10%] tw-z-0 tw-mr-[0.2rem] tw-mt-8 tw-flex tw-flex-col tw-items-center tw-justify-center tw-rounded-lg tw-border-2 tw-py-2 lg:!tw-mt-[5.7rem] lg:!tw-py-12 xl:!tw-mt-[7.8rem] xl:!tw-py-4 2xl:!tw-pb-16 2xl:!tw-pt-12 2xl:tw-pr-8 xs:tw-mt-[2.8rem] xs:tw-py-4"
      >
        <div
          class="tw-ml-[27rem] xl:tw-ml-[27.5rem] tw-text-black-core/[0.87] header-1 lg:!tw-w-[47%] 2xl:!tw-w-[60%]"
        >
          <span>Let’s talk about what we can build and scale together.</span>
        </div>
        <div
          class="tw-ml-auto tw-mt-4 lg:tw-mt-10 xl:!tw-ml-[14%] 2xl:!tw-ml-[-7%]"
        >
          <router-link
            class="gradient-btn primary-btn"
            to="/contact"
            @click.native="mixpanel.track('tap_android_development_cta')"
          >
            <span class="sub-h4-semibold">Contact Us </span>
          </router-link>
        </div>
      </div>
    </div>
    <!-- ---------------------DESKTOP UI END------------------ -->
  </section>
</template>
<script>
import CtaBoyImage from "@/assets/images/andriod-app-development/cta/cta3.webp";

export default {
  data() {
    return {
      CtaBoyImage,
    };
  },

  inject: ["mixpanel"],
};
</script>
