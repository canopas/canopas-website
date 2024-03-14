<template>
  <section
    class="relative mt-16 bg-pink-16 bg-gradient-background lg:mt-0 lg:bg-none 3xl:outer-container"
  >
    <!-- MOBILE UI START  -->
    <div class="lg:hidden">
      <img
        :src="background400w"
        :srcset="`${background400w} 400w, ${background800w} 800w`"
        loading="lazy"
        class="absolute h-full w-full"
        alt="background-image"
      />
      <div
        class="container sticky flex flex-col items-center justify-center py-20"
      >
        <p class="text-primary-1 sub-h3-semibold">Get started today</p>
        <p class="mt-2 text-center mobile-header-1 text-black-87">
          Let's build the next <br />
          big thing!
        </p>
        <p class="mt-4 text-center sub-h1-regular text-black-60">
          Let's improve your business's digital strategy and implement robust
          mobile apps to achieve your business objectives. Schedule Your Free
          Consultation Now.
        </p>
        <nuxt-link
          class="mt-6 gradient-btn primary-btn"
          to="/contact"
          @click.native="$mixpanel.track('tap_landing_cta')"
        >
          <span class="sub-h3-semibold lg:sub-h1-semibold"
            >Get Free Consultation
          </span>
        </nuxt-link>
      </div>
    </div>
    <!-- MOBILE UI END -->

    <!-- DESKTOP UI START -->
    <div class="hidden lg:block">
      <img
        :src="desktop_background800w"
        :srcset="`${desktop_background800w} 800w, ${desktop_background1200w} 1200w, ${desktop_background1600w} 1600w, ${desktop_background2400w} 2400w `"
        loading="lazy"
        class="absolute h-full w-full"
        alt="background-image"
      />
      <div
        class="container sticky z-[1] flex flex-col items-center justify-center py-20 xll:py-32"
      >
        <p class="text-primary-1 text-center mobile-header-3-semibold">
          Get started today
        </p>
        <p class="mt-2 text-center desk-header-2 text-black-87">
          Let's build the next big thing!
        </p>
        <p
          class="mt-6 w-[75%] text-center mobile-header-2-regular text-black-60"
        >
          Let's improve your business's digital strategy and implement robust
          mobile apps to achieve your business objectives. Schedule Your Free
          Consultation Now.
        </p>
        <nuxt-link
          class="mt-10 gradient-btn primary-btn"
          to="/contact"
          @click.native="$mixpanel.track('tap_landing_cta')"
        >
          <span class="sub-h3-semibold lg:sub-h1-semibold"
            >Get Free Consultation
          </span>
        </nuxt-link>
      </div>
    </div>
    <!-- DESKTOP UI END -->
  </section>
</template>
<script>
import background400w from "@/assets/images/mobile-app-development/cta-1/mobile/background-400w.webp";
import background800w from "@/assets/images/mobile-app-development/cta-1/mobile/background-800w.webp";
import desktop_background800w from "@/assets/images/mobile-app-development/cta-1/desktop/background-800w.webp";
import desktop_background1200w from "@/assets/images/mobile-app-development/cta-1/desktop/background-1200w.webp";
import desktop_background1600w from "@/assets/images/mobile-app-development/cta-1/desktop/background-1600w.webp";
import desktop_background2400w from "@/assets/images/mobile-app-development/cta-1/desktop/background-2400w.webp";

export default {
  data() {
    return {
      background400w,
      background800w,
      desktop_background800w,
      desktop_background1200w,
      desktop_background1600w,
      desktop_background2400w,
    };
  },
  inject: ["mixpanel"],
};
</script>
