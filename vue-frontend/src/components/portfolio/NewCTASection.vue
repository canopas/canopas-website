<template>
  <div
    class="tw-relative tw-top-10 sm:tw-top-0.5 tw-pt-12 tw-font-inter-medium"
  >
    <div class="tw-flex tw-flex-col tw-items-center">
      <img
        :src="Cta2400"
        :srcset="`${Cta400} 400w, ${Cta800} 800w, ${Cta1200} 1200w, ${Cta2400} 2400w`"
        class="tw-absolute tw-bottom-8 sm:tw-bottom-0.5 tw-max-w-[1700px] tw-w-[93%] tw-h-[68%] sm:tw-h-full tw-rounded-[8px] sm:tw-rounded-[12px] md:tw-rounded-[14px] lg:tw-rounded-[16px] xl2:tw-rounded-[20px] tw-bg-gradient-to-r tw-from-[#F2709C] tw-to-[#FF835B]"
        alt="canopas-CTA"
      />
      <div class="tw-container tw-mx-6 md:tw-mx-10 tw-z-[1] tw-text-center">
        <div
          class="md:tw-m-2 lg:tw-m-8 tw-text-[1.125rem] sm:tw-text-[1.625rem] md:tw-text-[2.125rem] lg:tw-text-[2.625rem] 2xl:tw-text-[3.125rem] tw-leading-[1.3rem] sm:tw-leading-[2.1rem] lg:tw-leading-[2.9rem] tw-text-white tw-font-bold"
        >
          <div class="tw-mt-6 sm:tw-mt-2 md:tw-mt-5 2xl:tw-mt-20">
            Got a Project in Mind? Get Started Now.
          </div>
          <router-link
            to="/contact"
            class="tw-relative tw-inline-block tw-mt-6 sm:tw-mt-9 md:tw-mt-12 lg:tw-mt-16 2xl:tw-mt-28 tw-pb-2 sm:tw-pb-3 md:tw-pb-4 2xl:tw-pb-8 animation-underline tw-cursor-pointer after:tw-absolute after:tw-w-full after:tw-h-[2px] after:sm:tw-h-[3px] after:2xl:tw-h-[5px] after:tw-bottom-0 after:tw-left-0 after:tw-bg-white after:tw-origin-bottom-left after:tw-scale-x-100 after:tw-duration-500 after:tw-delay-150 hover:after:tw-animate-underlineOut"
          >
            Let’s Work Together
          </router-link>
        </div>
      </div>
    </div>
    <div class="tw-top-1 tw-bg-black-900 tw-px-full tw-py-8 2xl:tw-py-10"></div>
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
