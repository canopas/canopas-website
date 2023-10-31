<template>
  <section class="container py-16 md:py-40 xl:py-60">
    <!-- ---------------------MOBILE UI START------------------ -->
    <div class="lg:hidden relative flex flex-col">
      <img
        :src="CtaBoyImage"
        class="z-[1] mx-auto h-full w-[14.1875rem] md:w-64 object-cover xs:w-[13.1875rem]"
        alt="android-cta-image"
        loading="lazy"
      />
      <div
        class="w-[80%] mx-auto sm:w-full text-center mt-8 text-black-core/[0.87] header-1"
      >
        Let’s talk about what we can build and scale together.
      </div>
      <div class="mx-auto mt-8 z-[1]">
        <router-link
          class="gradient-btn primary-btn"
          to="/contact"
          @click.native="$mixpanel.track('tap_android_development_cta')"
        >
          <span class="sub-h3-semibold">Contact Us </span>
        </router-link>
      </div>
      <div
        class="top-12 md:top-14 absolute inset-0 z-[0] mx-auto h-[85%] sm:h-[83%] w-[100%] rounded-[0.625rem] border border-black/[0.87]"
      ></div>
    </div>
    <!-- ---------------------MOBILE UI END------------------ -->

    <!-- ---------------------DESKTOP UI START------------------ -->
    <div class="hidden relative lg:flex flex-row">
      <div class="relative z-[1]">
        <img
          :src="CtaBoyImage"
          class="h-full w-[14.1875rem] object-cover !w-[32.625rem] xl:!w-[34.625rem] xs:w-[13.1875rem]"
          alt="android-cta-image"
          loading="lazy"
        />
      </div>
      <div
        class="h-[82%] lg:h-[80%] xl:h-[75%] border-[#00000]/[0.6] absolute left-[10%] z-0 mr-[0.2rem] mt-8 flex flex-col items-center justify-center rounded-lg border-2 py-2 lg:!mt-[5.7rem] lg:!py-12 xl:!mt-[7.8rem] xl:!py-4 2xl:!pb-16 2xl:!pt-12 2xl:pr-8 xs:mt-[2.8rem] xs:py-4"
      >
        <div
          class="ml-[27rem] xl:ml-[27.5rem] text-black-core/[0.87] header-1 lg:!w-[47%] 2xl:!w-[60%]"
        >
          <span>Let’s talk about what we can build and scale together.</span>
        </div>
        <div class="ml-auto mt-4 lg:mt-10 xl:!ml-[14%] 2xl:!ml-[-7%]">
          <router-link
            class="gradient-btn primary-btn"
            to="/contact"
            @click.native="$mixpanel.track('tap_android_development_cta')"
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
