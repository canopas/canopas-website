<template>
  <section class="relative pb-[46px] md:pt-[50px]">
    <div
      class="absolute right-0 top-1 hidden h-[140px] border-l-[20px] border-[#F2709C] lg:block lg:top-9 xl:h-[185px] 2xl:top-[4.25rem]"
    ></div>
    <div
      class="absolute bottom-[6.25rem] left-0 hidden h-[115px] border-l-[20px] border-[#FF9472] lg:block"
    ></div>

    <div
      class="container xll:max-w-[1640px] flex flex-col items-center lg:flex-row lg:flex-row-reverse"
    >
      <div
        class="flex flex-col lg:relative items-center lg:items-end lg:flex-[55%]"
      >
        <p
          class="text-center mobile-header-1 lg:desk-header-1 text-white lg:text-right xl:w-[89%] 2xl:w-[77%] xll:w-[61%]"
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
          class="mt-4 lg:mt-[2.5rem] w-[98%] md:w-[70%] lg:w-[80%] xl:w-[75%] 2xl:w-[65%] xll:w-[52%] text-center sub-h1-regular lg:mobile-header-2-regular text-white-core-80 lg:text-right"
        >
          Not sure where to start? We also offer code and architecture reviews,
          strategic planning, and more.
        </p>
        <nuxt-link
          class="gradient-btn primary-btn mr-0 hidden lg:flex lg:mt-10"
          to="/contact"
        >
          <span class="sub-h1-semibold">Get Free Consultation </span>
        </nuxt-link>
      </div>
      <div class="relative lg:flex-[45%]">
        <img
          :src="ctaImage400w"
          :srcset="`${ctaImage400w} 400w, ${ctaImage800w} 800w`"
          class="mt-8 h-full w-full md:w-[50%] md:h-[50%] lg:h-full lg:w-full lg:mt-0 md:mx-auto"
          alt="cta-image"
          loading="lazy"
        />
      </div>

      <nuxt-link class="gradient-btn primary-btn lg:hidden" to="/contact">
        <span class="sub-h3-semibold">Get Free Consultation </span>
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
