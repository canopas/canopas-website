<template>
  <section class="bg-ios-landing-background bg-opacity-10 3xl:outer-container">
    <div
      class="container flex flex-col lg:flex-row pt-6 lg:py-[7.5rem] gap-4 text-center lg:text-left"
    >
      <div class="flex flex-col gap-4 lg:gap-6 lg:flex-1">
        <h1 class="mobile-header-1 lg:desk-header-1 text-black-87 lg:w-1/2">
          Backend development company
        </h1>
        <p
          class="sub-h1-regular lg:mobile-header-2-regular xs:px-4 lg:!px-0 text-black-60"
        >
          Our backend development team isn’t just about writing code; it's about
          crafting powerful, efficient, scalable, and secure foundations that
          drive innovative applications. We shape, optimize, and elevate digital
          platforms ensuring they're ready for tomorrow's challenges.
        </p>
        <nuxt-link
          class="gradient-btn primary-btn ml-0 hidden lg:flex"
          to="/contact"
          @click.native="$mixpanel.track('tap_backend_landing_cta')"
        >
          <span class="sub-h3-semibold lg:sub-h1-semibold"
            >Be Future-Ready With Us</span
          ></nuxt-link
        >
      </div>
      <div class="flex flex-col -gap-1 lg:flex-1 lg:my-auto">
        <div class="w-4/5 md:w-3/5 lg:w-full mx-auto">
          <img
            :src="landing"
            alt="backend"
            class="w-full h-full object-cover"
          />
        </div>
        <nuxt-link
          class="gradient-btn primary-btn mb-16 lg:hidden"
          to="/contact"
          @click.native="$mixpanel.track('tap_backend_landing_cta')"
        >
          <span class="sub-h3-semibold"
            >Be Future-Ready With Us</span
          ></nuxt-link
        >
      </div>
    </div>
  </section>
</template>
<script setup>
import landing from "@/assets/images/backend-development/landing/landing.webm";
const { $mixpanel } = useNuxtApp();
</script>
