<template>
  <div class="relative h-[317px] lg:h-[476px] 3xl:outer-container">
    <img
      :src="background400"
      :srcset="`${background400} 400w, ${background800} 800w, ${background1200} 1200w, ${background2400} 2400w`"
      alt="Background-Image"
      class="h-full w-full object-cover"
      loading="lazy"
    />
    <div
      class="container absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 text-center sm:w-[90%]"
    >
      <div class="flex flex-col gap-3 lg:gap-7">
        <div class="flex flex-col gap-4 lg:gap-6">
          <h2 class="mobile-header-2 lg:desk-header-2 text-white">
            Not sure where to start?
          </h2>
          <p
            class="text-white-core-80 sub-h1-regular lg:mobile-header-2-regular"
          >
            We also offer code and architecture reviews, strategic planning,
            idea <br class="hidden lg:inline-block" />
            validation, and much more.
          </p>
        </div>
        <router-link
          class="gradient-btn primary-btn"
          to="/contact"
          @click.native="$mixpanel.track('tap_backend_development_cta')"
        >
          <span class="sub-h3-semibold xl:sub-h1-semibold"
            >Book a Free Consultation Call Today!</span
          >
        </router-link>
      </div>
    </div>
  </div>
</template>

<script setup type="module">
import background400 from "@/assets/images/backend-development/cta/ctabackground-400w.webp";
import background800 from "@/assets/images/backend-development/cta/ctabackground-800w.webp";
import background1200 from "@/assets/images/backend-development/cta/ctabackground-1200w.webp";
import background2400 from "@/assets/images/backend-development/cta/ctabackground-2400w.webp";
const { $mixpanel } = useNuxtApp();
</script>
