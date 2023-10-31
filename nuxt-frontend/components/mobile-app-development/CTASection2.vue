<template>
  <section class="relative pb-[50px] md:pt-[50px] 2xl:-mt-px">
    <div
      class="absolute right-0 top-1 hidden h-[140px] border-l-[20px] border-[#F2709C] md:block lg:top-9 xl:h-[185px] 2xl:top-[4.25rem]"
    ></div>
    <div
      class="absolute bottom-[6.25rem] left-0 hidden h-[115px] border-l-[20px] border-[#FF9472] md:block"
    ></div>

    <div
      class="container xll:max-w-[1640px] flex flex-col items-center md:flex-row md:flex-row-reverse"
    >
      <div class="flex flex-col md:relative md:items-end lg:flex-[55%]">
        <p
          class="text-center font-inter-bold text-[2.625rem] leading-[3.125rem] text-white md:text-right md:text-[3.275rem] md:leading-[3.875rem] xl:text-[5.625rem] xl:leading-[6.125rem]"
        >
          <span class="bg-gradient-L bg-clip-text text-transparent">L</span>et's
          <span class="mr-[-0.2rem] bg-gradient-W bg-clip-text text-transparent"
            >W</span
          >ork
          <span class="mr-[-0.2rem] bg-gradient-T bg-clip-text text-transparent"
            >T</span
          >ogether
        </p>
        <p
          class="mt-6 w-[385px] text-center font-inter-regular text-base text-white md:text-right md:text-[1.4rem] md:leading-8 md:text-white/[0.75] xl:w-[82%] xl:text-[1.75rem] xl:leading-[2.625rem] 2xl:w-[68%] xll:w-[55%]"
        >
          Not sure where to start? We also offer code and architecture reviews,
          strategic planning, and more.
        </p>
        <nuxt-link
          class="gradient-btn consultation-btn mr-0 mt-10 hidden w-max items-end rounded-full py-5 text-right md:flex"
          to="/contact"
        >
          <span
            class="mr-2.5 font-inter-semibold text-[1.09375rem] leading-[1.3125rem] xl:text-[1.188rem] xl:leading-[1.188rem]"
            >Get Free Consultation
          </span>
        </nuxt-link>
      </div>
      <div class="relative lg:flex-[45%]">
        <img
          :src="ctaImage400w"
          :srcset="`${ctaImage400w} 400w, ${ctaImage800w} 800w`"
          class="mt-8 h-full w-full md:mt-0"
          alt="cta-image"
          loading="lazy"
        />
        <div
          class="gradient-border absolute bottom-4 left-[0.5rem] h-[55px] border-l-[13px] md:hidden"
        ></div>
      </div>

      <nuxt-link
        class="gradient-btn consultation-btn mt-10 flex w-max items-center rounded-full py-5 text-center md:hidden"
        to="/contact"
      >
        <span
          class="mr-2.5 font-inter-semibold text-[1.125rem] leading-[1.125rem]"
          >Get Free Consultation
        </span>
      </nuxt-link>
    </div>
  </section>
</template>
<script>
import ctaImage400w from "@/assets/images/mobile-app-development/cta-2/cta-image-400w.webp";
import ctaImage800w from "@/assets/images/mobile-app-development/cta-2/cta-image-800w.webp";

export default {
  data() {
    return {
      ctaImage400w,
      ctaImage800w,
    };
  },
};
</script>
<style scoped>
.gradient-border {
  border-image: linear-gradient(to bottom, #f2709c, #ff835b) 1;
}
</style>
