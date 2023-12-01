<template>
  <section class="mt-6 lg:mt-28">
    <div class="container lg:text-left text-center lg:flex pb-16 lg:pb-60">
      <div class="flex flex-col gap-4 lg:gap-6 lg:flex-1">
        <h1 class="header-1 text-black-core/[0.87]">
          Flutter app development company
        </h1>
        <p class="sub-h1-regular text-black-core/[0.6]">
          We harness the power of Flutter to create visually stunning, fast, and
          highly responsive applications that effortlessly work on both iOS and
          Android platforms. Our team turns your unique business requirements
          into reality, giving your brand the competitive edge it needs in this
          dynamic mobile app market.
        </p>
        <div class="hidden lg:block">
          <nuxt-link
            class="gradient-btn primary-btn ml-0"
            to="/contact"
            @click.native="$mixpanel.track('tap_flutter_development_cta')"
          >
            <span class="sub-h3-semibold xl:sub-h4-semibold border-none"
              >Get Free Consultation</span
            >
          </nuxt-link>
        </div>
      </div>
      <div
        class="mx-auto w-[14.5625rem] xs:w-[20.5625rem] mt-8 lg:mt-0 lg:mx-0 md:gap-[1rem] lg:flex-1"
      >
        <img
          :src="landing400"
          :srcset="`${landing400} 400w, ${landing800} 800w, ${landing1200} 1200w, ${landing2400} 2400w`"
          class="h-full w-full object-contain"
          alt="flutter-landing"
        />
      </div>
      <div class="lg:hidden mt-3">
        <nuxt-link
          class="gradient-btn primary-btn"
          to="/contact"
          @click.native="$mixpanel.track('tap_flutter_development_cta')"
        >
          <span class="sub-h3-semibold border-none">Get Free Consultation</span>
        </nuxt-link>
      </div>
    </div>
  </section>
</template>

<script setup>
import landing400 from "@/assets/images/flutter-app-development/landing/landing400w.webp";
import landing800 from "@/assets/images/flutter-app-development/landing/landing800w.webp";
import landing1200 from "@/assets/images/flutter-app-development/landing/landing1200w.webp";
import landing2400 from "@/assets/images/flutter-app-development/landing/landing2400w.webp";
const { $mixpanel } = useNuxtApp();
</script>
