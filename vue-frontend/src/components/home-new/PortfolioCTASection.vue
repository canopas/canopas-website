<template>
  <div
    class="tw-from-[#F98584] tw-to-[#F47695] tw-bg-gradient-[180deg] tw-font-inter-regular tw-mt-[7.813rem]"
  >
    <div class="tw-container tw-flex tw-flex-col lg:tw-flex-row">
      <div class="tw-hidden lg:tw-block lg:tw-w-[45%] xl:tw-w-[50%]">
        <img
          :src="CTAImage[3]"
          :srcset="`${CTAImage[0]} 400w, ${CTAImage[1]} 800w, ${CTAImage[2]} 1600w, ${CTAImage[3]} 2400w`"
          alt="portfolio-cta-image"
          class="tw-w-full tw-h-full tw-object-fill xl:tw-object-contain"
          loading="lazy"
        />
      </div>

      <div class="tw-flex lg:tw-w-[55%] xl:tw-w-[50%]">
        <div
          class="tw-m-auto lg:tw-ml-auto xl:tw-m-auto tw-w-[80%] sm:tw-w-[85%] md:tw-w-[69%] lg:tw-w-[100%] xl:tw-w-[90%] 2xl:tw-w-[92%] tw-py-[20px] tw-text-white"
        >
          <p
            class="tw-text-[1.5rem] tw-leading-[1.813rem] md:tw-text-[2.188rem] md:tw-leading-[2.625rem] lg:tw-text-[3.125rem] lg:tw-leading-[3.813rem] tw-font-inter-bold"
          >
            Whether you need...
          </p>
          <ul
            class="tw-pt-6 tw-text-[1rem] tw-leading-[1.25rem] md:tw-text-[1.188rem] md:tw-leading-[1.5rem] lg:tw-text-[1.375rem] lg:tw-leading-[1.8125rem] tw-font-inter-semibold"
          >
            <li v-for="reason in reasons" :key="reason" class="tw-flex">
              <span class="tw-pr-1">*</span>
              <div class="sm:tw-w-[90%]">{{ reason }}</div>
            </li>
          </ul>
          <div
            class="xl:tw-w-[100%] tw-pt-6 tw-pb-8 tw-text-[0.875rem] tw-leading-[1.125rem] md:tw-text-[1.125rem] md:tw-leading-[1.375rem] lg:tw-text-[1.375rem] lg:tw-leading-[1.813rem]"
          >
            Bring us your toughest challenge and we'll show you the path to a
            sleek solution.
          </div>
          <router-link
            class="tw-flex tw-items-center tw-mx-auto lg:tw-mx-0 tw-w-max tw-rounded-full tw-p-3 tw-text-center tw-font-normal tw-text-[1rem] tw-leading-[1.1875rem] md:tw-text-[1.09375rem] md:tw-leading-[1.3125rem] lg:tw-text-[1.1875rem] lg:tw-leading-[1.4375rem] tw-font-inter-medium !tw-tracking-[0] tw-border-[1px] tw-border-solid tw-border-white tw-bg-white hover:tw-bg-transparent hover:tw-text-white active:tw-scale-[0.98] cta-button"
            to="/contact"
            @click.native="mixpanel.track('tap_home_portfolio_cta')"
          >
            <span
              class="tw-bg-gradient-[270.11deg] tw-bg-clip-text tw-text-transparent tw-from-[#ff9472] tw-to-[#f2709c] tw-font-inter-semibold"
              >Talk to our experts</span
            >
          </router-link>
        </div>
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
