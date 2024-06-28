<template>
  <div class="mt-16 bg-white-smoke 3xl:outer-container">
    <div class="container flex flex-col lg:flex-row">
      <div class="hidden lg:block relative lg:w-1/2">
        <img
          :src="CTAImage[0]"
          :srcset="`${CTAImage[0]} 400w, ${CTAImage[1]} 800w`"
          alt="portfolio-cta-image"
          class="w-fit h-fit object-contain absolute right-10 -top-[0.3rem] xl:right-6 xl:-top-[5.5rem] 2xl:right-14 2xl:-top-[9.5rem]"
          loading="lazy"
        />
      </div>

      <div
        class="flex lg:w-1/2 py-2.5 lg:py-10 2xl:py-[3.125rem] 2xl:pl-[8.5rem]"
      >
        <div
          class="m-auto lg:ml-auto xl:m-auto w-4/5 sm:w-[63%] md:w-1/2 lg:w-full py-10 lg:py-0"
        >
          <p class="mobile-header-1 lg:desk-header-2 text-black-87">
            Whether you need...
          </p>
          <ul
            class="pt-4 lg:pt-6 sub-h3-semibold lg:mobile-header-2-semibold text-black-60 pl-0"
          >
            <li v-for="reason in reasons" :key="reason" class="flex">
              <span class="pr-1">*</span>
              <div class="sm:w-[90%]">{{ reason }}</div>
            </li>
          </ul>
          <div
            class="pt-4 sub-h4-regular lg:mobile-header-3-regular text-black-60"
          >
            Bring us your toughest challenge and we'll show you the path to a
            sleek solution.
          </div>
          <nuxt-link
            class="mt-6 lg:mt-8 md:mx-auto md:ml-0 gradient-btn primary-btn"
            to="/contact"
            @click.native="$mixpanel.track('tap_landing_cta')"
          >
            <span class="sub-h3-semibold lg:sub-h1-semibold"
              >Talk To Our Experts
            </span>
          </nuxt-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script type="module">
export default {
  data() {
    return {
      CTAImage: [
        "/images/portfolio/new-portfolio/new-cta-400w.webp",
        "/images/portfolio/new-portfolio/new-cta-800w.webp",
      ],
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
