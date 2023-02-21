<template>
  <div
    class="tw-mt-[20px] lg:tw-mt-[100px] xl:tw-mt-[200px] 2xl:tw-mt-[250px] tw-bg-[#F8F8F8] tw-font-inter-regular"
  >
    <div class="tw-container tw-flex tw-flex-col lg:tw-flex-row">
      <div class="lg:tw-relative tw-order-2 lg:tw-order-first lg:tw-w-[50%]">
        <img
          :src="CTAImage[0]"
          :srcset="`${CTAImage[0]} 400w, ${CTAImage[1]} 800w, ${CTAImage[2]} 1600w, ${CTAImage[3]} 2400w`"
          alt="portfolio-cta-image"
          class="-tw-mb-[0.75rem] sm:-tw-mb-[1.5rem] md:-tw-mb-[1.75rem] lg:tw-absolute lg:tw-right-[2.5rem] lg:-tw-top-[5rem] xl:tw-right-[1.5rem] xl:-tw-top-[7.5rem] 2xl:tw-right-[3.5rem] 2xl:-tw-top-[11.5rem]"
          loading="lazy"
        />
      </div>

      <div
        class="tw-flex lg:tw-w-[50%] tw-py-[10px] lg:tw-py-[40px] 2xl:tw-py-[50px] 2xl:tw-pl-[8.5rem]"
      >
        <div
          class="tw-m-auto lg:tw-ml-auto xl:tw-m-auto tw-w-[80%] sm:tw-w-[60%] md:tw-w-[50%] lg:tw-w-full tw-py-[40px] lg:tw-py-0"
        >
          <h1
            class="tw-text-black-core/[0.87] tw-text-[1.5rem] tw-leading-[1.813rem] md:tw-text-[2.188rem] md:tw-leading-[2.625rem] xl:tw-text-[3.125rem] xl:tw-leading-[3.813rem] tw-font-inter-bold"
          >
            Whether you need...
          </h1>
          <ul
            class="tw-pt-6 tw-text-[1rem] tw-text-black-core/[0.6] tw-leading-[1.25rem] md:tw-text-[1.188rem] md:tw-leading-[1.5rem] xl:tw-text-[1.375rem] xl:tw-leading-[1.8125rem] tw-font-inter-semibold"
          >
            <li v-for="reason in reasons" :key="reason" class="tw-flex">
              <span class="tw-pr-1">*</span>
              <div class="sm:tw-w-[90%]">{{ reason }}</div>
            </li>
          </ul>
          <div
            class="tw-pt-6 tw-text-black-core/[0.6] tw-font-inter-regular tw-text-[0.875rem] tw-leading-[1.125rem] md:tw-text-[1.125rem] md:tw-leading-[1.375rem] xl:tw-text-[1.375rem] xl:tw-leading-[1.813rem]"
          >
            Bring us your toughest challenge and we'll show you the path to a
            sleek solution.
          </div>
          <router-link
            class="tw-flex tw-items-center tw-mx-auto lg:tw-mx-0 tw-w-max tw-rounded-full tw-p-3 tw-text-center tw-font-normal tw-text-[1rem] tw-leading-[1.1875rem] md:tw-text-[1.09375rem] md:tw-leading-[1.3125rem] xl:tw-text-[1.1875rem] xl:tw-leading-[1.4375rem] tw-font-inter-medium !tw-tracking-[0] gradient-btn hover:tw-text-white active:tw-scale-[0.98]"
            to="/contact"
            @click.native="mixpanel.track('tap_home_portfolio_cta')"
          >
            <span class="tw-text-white tw-font-inter-semibold"
              >Talk to our experts</span
            >
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script type="module">
import cta400 from "@/assets/images/portfolio/new-portfolio/new-cta-400w.webp";
import cta800 from "@/assets/images/portfolio/new-portfolio/new-cta-800w.webp";
import cta1600 from "@/assets/images/portfolio/new-portfolio/new-cta-1600w.webp";
import cta2400 from "@/assets/images/portfolio/new-portfolio/new-cta-2400w.webp";

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
