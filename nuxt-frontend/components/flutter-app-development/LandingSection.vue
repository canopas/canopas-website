<template>
  <section class="mt-6 lg:mt-[6.75rem]">
    <div class="container lg:text-left text-center lg:flex xl:my-8 pb-16">
      <div class="flex flex-col gap-4 lg:gap-6 lg:flex-1">
        <h1 class="header-1">Flutter App Development Company</h1>
        <p class="sub-h1-regular text-black-core/[0.89]">
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
            <span class="sub-h3-semibold border-none"
              >Get Free Consultation</span
            >
          </nuxt-link>
        </div>
      </div>
      <div
        class="mx-20 mt-8 lg:mt-0 lg:mx-0 flex flex-col md:gap-[1rem] lg:flex-1"
      >
        <img
          :src="landing400"
          :srcset="`${landing400} 400w, ${landing800} 800w, ${landing1200} 1200w, ${landing2400} 2400w`"
          class="h-full w-full object-contain"
          alt="flutter-landing"
        />
        <div class="lg:hidden">
          <nuxt-link
            class="gradient-btn primary-btn"
            to="/contact"
            @click.native="$mixpanel.track('tap_flutter_development_cta')"
          >
            <span class="sub-h3-semibold border-none"
              >Get Free Consultation</span
            >
          </nuxt-link>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import landing400 from "@/assets/images/flutter-app-development/landing/landing400w.webp";
import landing800 from "@/assets/images/flutter-app-development/landing/landing800w.webp";
import landing1200 from "@/assets/images/flutter-app-development/landing/landing1200w.webp";
import landing2400 from "@/assets/images/flutter-app-development/landing/landing2400w.webp";
export default {
  data() {
    return {
      landing400,
      landing800,
      landing1200,
      landing2400,
    };
  },
  inject: ["mixpanel"],
};
</script>
