<template>
  <div class="relative sm:pt-12">
    <div class="flex flex-col items-center">
      <img
        :src="Cta2400"
        :srcset="`${Cta400} 400w, ${Cta800} 800w, ${Cta1200} 1200w, ${Cta2400} 2400w`"
        class="absolute h-[81%] w-[94%] object-cover bottom-8 sm:bottom-0.5 max-w-[1700px] sm:h-full rounded-lg sm:rounded-xl md:rounded-[14px] lg:rounded-2xl xl2:rounded-[20px] bg-gradient-to-r from-orange-300 to-pink-300"
        alt="canopas-CTA"
      />
      <div class="container mx-6 lg:mx-10 z-[1] text-center">
        <div class="md:m-2 lg:m-8 text-white">
          <div
            class="mt-6 sm:mt-2 md:mt-5 2xl:mt-16 sub-h3-semibold lg:desk-header-3"
          >
            Got a Project in Mind? Get Started Now.
          </div>
          <nuxt-link
            to="/contact"
            class="sub-h1-semibold lg:desk-header-2 relative inline-block mt-4 lg:mt-8 pb-2 animation-underline cursor-pointer after:absolute after:w-full after:h-0.5 after:sm:h-[3px] after:2xl:h-[5px] after:bottom-0 after:left-0 after:bg-white after:origin-bottom-left after:scale-x-100 after:duration-500 after:delay-150 hover:after:animate-underlineOut"
          >
            Let’s Work Together
          </nuxt-link>
        </div>
      </div>
    </div>
    <div class="top-1 bg-deep-charcoal px-full py-8 2xl:py-10"></div>
  </div>
</template>

<script type="module">
import Cta400 from "@/assets/images/CTA/Cta-400w.webp";
import Cta800 from "@/assets/images/CTA/Cta-800w.webp";
import Cta1200 from "@/assets/images/CTA/Cta-1200w.webp";
import Cta2400 from "@/assets/images/CTA/Cta-2400w.webp";

export default {
  data() {
    return {
      Cta400,
      Cta800,
      Cta1200,
      Cta2400,
    };
  },
};
</script>
