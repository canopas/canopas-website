<template>
  <section
    class="bg-deep-charcoal text-white pt-6 pb-14 lg:py-36 3xl:outer-container"
  >
    <div
      class="container flex flex-col lg:flex-row text-center lg:text-left gap-8 lg:gap-4"
    >
      <div class="flex flex-col gap-4 lg:gap-6 lg:flex-1">
        <h1 class="mobile-header-1 lg:desk-header-1">
          Frontend <br class="sm:hidden" />development company
        </h1>
        <p class="sub-h1-regular lg:mobile-header-2-regular text-white-core-80">
          At Canopas, we merge art with technology, ensuring your website or app
          that feels and looks impeccable. Our team will help you design
          visually stunning, high-performing, and user-centric interfaces.
        </p>
        <nuxt-link
          class="gradient-btn primary-btn ml-0 mt-6 hidden lg:flex"
          to="/contact"
          @click.native="$mixpanel.track('tap_frontend_landing_cta')"
        >
          <span class="sub-h1-semibold"
            >Talk to Frontend Experts</span
          ></nuxt-link
        >
      </div>
      <div
        class="w-[17.25013rem] h-[17.25rem] mx-auto lg:w-[31.25rem] lg:h-[31.75rem] lg:flex-1"
      >
        <img
          :src="landing400"
          :srcset="`${landing400} 400w, ${landing800} 800w, ${landing1200} 1200w, ${landing2400} 2400w`"
          class="mx-auto w-full h-full object-contain"
          alt="frontend-landing"
        />
      </div>
      <nuxt-link
        class="gradient-btn primary-btn lg:hidden"
        to="/contact"
        @click.native="$mixpanel.track('tap_frontend_landing_cta')"
      >
        <span class="sub-h3-semibold">Talk to Frontend Experts</span></nuxt-link
      >
    </div>
  </section>
</template>
<script setup>
import landing400 from "@/assets/images/frontend-development/landing/landing400w.webp";
import landing800 from "@/assets/images/frontend-development/landing/landing800w.webp";
import landing1200 from "@/assets/images/frontend-development/landing/landing1200w.webp";
import landing2400 from "@/assets/images/frontend-development/landing/landing2400w.webp";
const { $mixpanel } = useNuxtApp();
</script>
