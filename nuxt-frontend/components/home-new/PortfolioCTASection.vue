<template>
  <div
    class="from-[#F98584] to-[#F47695] bg-gradient-[180deg] font-inter-regular mt-[2.813rem] lg:mt-[7.813rem]"
  >
    <div class="container flex flex-col lg:flex-row">
      <div class="hidden lg:block lg:w-[45%] xl:w-[50%]">
        <img
          :src="CTAImage[3]"
          :srcset="`${CTAImage[0]} 400w, ${CTAImage[1]} 800w, ${CTAImage[2]} 1600w, ${CTAImage[3]} 2400w`"
          alt="portfolio-cta-image"
          class="w-full h-full object-fill xl:object-contain"
          loading="lazy"
        />
      </div>

      <div class="flex lg:w-[55%] xl:w-[50%]">
        <div
          class="m-auto lg:ml-auto xl:m-auto w-[80%] sm:w-[85%] md:w-[69%] lg:w-[100%] xl:w-[90%] 2xl:w-[92%] py-5 text-white"
        >
          <p
            class="text-[1.5rem] leading-[1.813rem] md:text-[2.188rem] md:leading-[2.625rem] lg:text-[3.125rem] lg:leading-[3.813rem] font-inter-bold"
          >
            Whether you need...
          </p>
          <ul
            class="pt-6 text-[1rem] leading-5 md:text-[1.188rem] md:leading-6 lg:text-[1.375rem] lg:leading-[1.8125rem] font-inter-semibold"
          >
            <li v-for="reason in reasons" :key="reason" class="flex">
              <span class="pr-1">*</span>
              <div class="sm:w-[90%]">{{ reason }}</div>
            </li>
          </ul>
          <div
            class="xl:w-[100%] pt-6 pb-8 text-[0.875rem] leading-[1.125rem] md:text-[1.125rem] md:leading-[1.375rem] lg:text-[1.375rem] lg:leading-[1.813rem]"
          >
            Bring us your toughest challenge and we'll show you the path to a
            sleek solution.
          </div>
          <nuxt-link
            class="flex items-center mx-auto lg:mx-0 w-max rounded-full p-3 text-center font-normal text-[1rem] leading-[1.1875rem] md:text-[1.09375rem] md:leading-[1.3125rem] lg:text-[1.1875rem] lg:leading-[1.4375rem] font-inter-medium !tracking-[0] border border-solid border-white bg-white hover:bg-transparent hover:text-white active:scale-[0.98] cta-button"
            to="/contact"
            @click.native="$mixpanel.track('tap_home_portfolio_cta')"
          >
            <span
              class="bg-gradient-[270.11deg] bg-clip-text text-transparent from-[#ff9472] to-[#f2709c] font-inter-semibold"
              >Talk to our experts</span
            >
          </nuxt-link>
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
