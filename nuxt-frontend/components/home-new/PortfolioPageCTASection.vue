<template>
  <div
    class="mt-5 lg:mt-[100px] xl:mt-[200px] 2xl:mt-[250px] bg-[#F8F8F8] font-inter-regular"
  >
    <div class="container flex flex-col lg:flex-row">
      <div class="lg:relative order-2 lg:order-first lg:w-[50%]">
        <img
          :src="CTAImage[0]"
          :srcset="`${CTAImage[0]} 400w, ${CTAImage[1]} 800w, ${CTAImage[2]} 1600w, ${CTAImage[3]} 2400w`"
          alt="portfolio-cta-image"
          class="-mb-3 w-fit h-fit object-contain sm:-mb-6 md:-mb-7 lg:absolute lg:right-10 lg:-top-20 xl:right-6 xl:-top-[7.5rem] 2xl:right-14 2xl:-top-[11.5rem]"
          loading="lazy"
        />
      </div>

      <div
        class="flex lg:w-[50%] py-2.5 lg:py-10 2xl:py-[50px] 2xl:pl-[8.5rem]"
      >
        <div
          class="m-auto lg:ml-auto xl:m-auto w-[80%] sm:w-[60%] md:w-[50%] lg:w-full py-10 lg:py-0"
        >
          <p
            class="text-black-core/[0.87] text-[1.5rem] leading-[1.813rem] md:text-[2.188rem] md:leading-[2.625rem] xl:text-[3.125rem] xl:leading-[3.813rem] font-inter-bold"
          >
            Whether you need...
          </p>
          <ul
            class="pt-6 text-[1rem] text-black-core/[0.6] leading-5 md:text-[1.188rem] md:leading-6 xl:text-[1.375rem] xl:leading-[1.8125rem] font-inter-semibold"
          >
            <li v-for="reason in reasons" :key="reason" class="flex">
              <span class="pr-1">*</span>
              <div class="sm:w-[90%]">{{ reason }}</div>
            </li>
          </ul>
          <div
            class="pt-6 text-black-core/[0.6] font-inter-regular text-[0.875rem] leading-[1.125rem] md:text-[1.125rem] md:leading-[1.375rem] xl:text-[1.375rem] xl:leading-[1.813rem]"
          >
            Bring us your toughest challenge and we'll show you the path to a
            sleek solution.
          </div>
          <nuxt-link
            class="flex items-center mx-auto lg:mx-0 w-max rounded-full p-3 text-center font-normal text-[1rem] leading-[1.1875rem] md:text-[1.09375rem] md:leading-[1.3125rem] xl:text-[1.1875rem] xl:leading-[1.4375rem] font-inter-medium !tracking-[0] gradient-btn hover:text-white active:scale-[0.98]"
            to="/contact"
            @click.native="$mixpanel.track('tap_portfolio_page_cta')"
          >
            <span class="text-white font-inter-semibold"
              >Talk to our experts</span
            >
          </nuxt-link>
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
