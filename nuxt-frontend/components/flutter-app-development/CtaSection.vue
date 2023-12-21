<template>
  <section>
    <div
      class="relative h-[352px] md:h-[430px] lg:h-[552px] 3xl:outer-container"
    >
      <img
        :src="background400"
        :srcset="`${background400} 400w, ${background800} 800w`"
        class="h-full w-full object-cover md:object-top lg:hidden"
        alt="flutter-cta"
        loading="lazy"
      />
      <img
        :src="background1200"
        :srcset="`${background1200} 1200w, ${background2400} 2400w`"
        class="h-full w-full object-cover hidden lg:block 2xll:object-right-top"
        alt="flutter-cta"
        loading="lazy"
      />
      <div class="absolute top-16 lg:top-12 xl:top-8 2xl:top-12 w-full">
        <div
          class="flex flex-col lg:flex-row container gap-1 lg:gap-[2.25rem] text-center lg:[&_img]:h-[calc(552px-theme(spacing.16)-1rem)]"
        >
          <img
            :src="hex"
            class="w-full h-full object-contain flex-1 hidden lg:block"
            alt="flutter-hex-image"
            loading="lazy"
          />
          <div
            class="flex flex-col gap-4 lg:gap-6 lg:text-left lg:flex-1 my-auto"
          >
            <div class="text-black-87">
              <p class="mobile-header-1 lg:desk-header-1">
                Let's code your success story together!
              </p>
            </div>
            <div class="text-black-60">
              <span class="sub-h1-regular lg:mobile-header-2-regular">
                Join hands with us for your next app endeavour.
              </span>
            </div>
            <div>
              <nuxt-link
                class="gradient-btn primary-btn lg:mx-0"
                to="/contact"
                @click.native="$mixpanel.track('tap_flutter_development_cta')"
              >
                <span class="sub-h3-semibold lg:sub-h1-semibold border-none"
                  >Connect With Us
                </span>
                <Icon name="fa6-solid:arrow-right" class="fa ml-3" />
              </nuxt-link>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup type="module">
import background400 from "@/assets/images/flutter-app-development/cta/background-400w.webp";
import background800 from "@/assets/images/flutter-app-development/cta/background-800w.webp";
import background1200 from "@/assets/images/flutter-app-development/cta/background-1200w.webp";
import background2400 from "@/assets/images/flutter-app-development/cta/background-2400w.webp";
import hex from "@/assets/images/flutter-app-development/cta/hex-img.webp";

const { $mixpanel } = useNuxtApp();
</script>
