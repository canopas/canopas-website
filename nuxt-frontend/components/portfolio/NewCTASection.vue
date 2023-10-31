<template>
  <div class="relative top-10 sm:top-0.5 pt-12 font-inter-medium">
    <div class="flex flex-col items-center">
      <img
        :src="Cta2400"
        :srcset="`${Cta400} 400w, ${Cta800} 800w, ${Cta1200} 1200w, ${Cta2400} 2400w`"
        class="absolute h-full w-full object-cover bottom-8 sm:bottom-0.5 max-w-[1700px] w-[93%] h-[68%] sm:h-full rounded-lg sm:rounded-xl md:rounded-[14px] lg:rounded-2xl xl2:rounded-[20px] bg-gradient-to-r from-[#F2709C] to-[#FF835B]"
        alt="canopas-CTA"
      />
      <div class="container mx-6 md:mx-10 z-[1] text-center">
        <div
          class="md:m-2 lg:m-8 text-[1.125rem] sm:text-[1.625rem] md:text-[2.125rem] lg:text-[2.625rem] 2xl:text-[3.125rem] leading-[1.3rem] sm:leading-[2.1rem] lg:leading-[2.9rem] text-white font-bold"
        >
          <div class="mt-6 sm:mt-2 md:mt-5 2xl:mt-20">
            Got a Project in Mind? Get Started Now.
          </div>
          <nuxt-link
            to="/contact"
            class="relative inline-block mt-6 sm:mt-9 md:mt-12 lg:mt-16 2xl:mt-28 pb-2 sm:pb-3 md:pb-4 2xl:pb-8 animation-underline cursor-pointer after:absolute after:w-full after:h-[2px] after:sm:h-[3px] after:2xl:h-[5px] after:bottom-0 after:left-0 after:bg-white after:origin-bottom-left after:scale-x-100 after:duration-500 after:delay-150 hover:after:animate-underlineOut"
          >
            Let’s Work Together
          </nuxt-link>
        </div>
      </div>
    </div>
    <div class="top-1 bg-black-900 px-full py-8 2xl:py-10"></div>
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
