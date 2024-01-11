<template>
  <div
    class="h-[352px] lg:h-[420px] xl:h-[476px] from-[#FF835B] to-[#F2709C] bg-gradient-[174deg] lg:bg-gradient-[180deg] 3xl:outer-container"
  >
    <div class="container flex flex-col lg:flex-row py-4 lg:py-0">
      <div class="hidden lg:block lg:w-full">
        <img
          :src="CTAImage[3]"
          :srcset="`${CTAImage[0]} 400w, ${CTAImage[1]} 800w, ${CTAImage[2]} 1600w, ${CTAImage[3]} 2400w`"
          alt="portfolio-cta-image"
          class="w-full object-fill 2xl:object-contain h-[420px] xl:h-[476px]"
          loading="lazy"
        />
      </div>

      <div
        class="flex flex-col md:flex-row lg:flex-col m-auto lg:ml-[25px] xl:m-auto w-[88%] sm:w-[85%] md:w-[94%] lg:w-[90%] 2xl:w-[92%] py-5 md:py-12"
      >
        <div class="text-white">
          <p class="mobile-header-1 xl:desk-header-2">Whether you need...</p>
          <ul
            class="pt-6 sub-h3-semibold lg:sub-h1-semibold xl:mobile-header-2-semibold pl-0"
          >
            <li v-for="reason in reasons" :key="reason" class="flex">
              <span class="pr-1">*</span>
              <div class="sm:w-[90%]">{{ reason }}</div>
            </li>
          </ul>
          <div
            class="md:w-[76%] lg:w-full xl:w-[77%] pt-6 pb-8 sub-h4-regular text-white-core-80 lg:sub-h3-regular xl:mobile-header-3-regular"
          >
            Bring us your toughest challenge and we'll show you the path to a
            sleek solution.
          </div>
        </div>
        <nuxt-link
          class="m-auto lg:ml-0 white-border-btn primary-btn border-none"
          to="/contact"
          @click.native="$mixpanel.track('tap_home_portfolio_cta')"
        >
          <span class="sub-h3-semibold lg:sub-h1-semibold"
            >Talk to our experts</span
          >
        </nuxt-link>
      </div>
    </div>
  </div>
</template>

<script type="module">
import cta400 from "@/assets/images/portfolio/new-portfolio/cta-400w.webp";
import cta800 from "@/assets/images/portfolio/new-portfolio/cta-800w.webp";
import cta1600 from "@/assets/images/portfolio/new-portfolio/cta-1600w.webp";
import cta2400 from "@/assets/images/portfolio/new-portfolio/cta-2400w.webp";

export default {
  data() {
    return {
      CTAImage: [cta400, cta800, cta1600, cta2400],
      reasons: [
        "High-performing mobile apps",
        "Bulletproof cloud solutions",
        "Custom solutions for your business.",
      ],
    };
  },
  inject: ["mixpanel"],
};
</script>
