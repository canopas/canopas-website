<template>
  <section class="tw-relative tw-pb-[50px] md:tw-pt-[50px] 2xl:tw-mt-[-1px]">
    <div
      class="tw-absolute tw-right-0 tw-top-1 tw-hidden tw-h-[140px] tw-border-l-[20px] tw-border-[#F2709C] md:tw-block lg:tw-top-[2.25rem] xl:tw-h-[185px] 2xl:tw-top-[4.25rem]"
    ></div>
    <div
      class="tw-absolute tw-bottom-[6.25rem] tw-left-0 tw-hidden tw-h-[115px] tw-border-l-[20px] tw-border-[#FF9472] md:tw-block"
    ></div>

    <div
      class="tw-container tw-flex tw-flex-col tw-items-center md:tw-flex-row md:tw-flex-row-reverse"
    >
      <div
        class="tw-flex tw-flex-col md:tw-relative md:tw-items-end lg:tw-flex-[55%]"
      >
        <p
          class="tw-text-center tw-font-inter-bold tw-text-[2.625rem] tw-leading-[3.125rem] tw-text-white md:tw-text-right md:tw-text-[3.275rem] md:tw-leading-[3.875rem] xl:tw-text-[5.625rem] xl:tw-leading-[6.125rem]"
        >
          <span class="tw-bg-gradient-L tw-bg-clip-text tw-text-transparent"
            >L</span
          >et's
          <span
            class="tw-mr-[-0.2rem] tw-bg-gradient-W tw-bg-clip-text tw-text-transparent"
            >W</span
          >ork
          <span
            class="tw-mr-[-0.2rem] tw-bg-gradient-T tw-bg-clip-text tw-text-transparent"
            >T</span
          >ogether
        </p>
        <p
          class="tw-mt-6 tw-w-[385px] tw-text-center tw-font-inter-regular tw-text-[1rem] tw-leading-[1.5rem] tw-text-white md:tw-text-right md:tw-text-[1.4rem] md:tw-leading-[2rem] md:tw-text-white/[0.75] xl:tw-w-[82%] xl:tw-text-[1.75rem] xl:tw-leading-[2.625rem] 2xl:tw-w-[68%]"
        >
          Not sure where to start? We also offer code and architecture reviews,
          strategic planning, and more.
        </p>
        <router-link
          class="gradient-btn consultation-btn tw-mr-0 tw-mt-10 tw-hidden tw-w-max tw-items-end tw-rounded-full tw-py-[1.25rem] tw-text-right md:tw-flex"
          :to="contactURL"
        >
          <span
            class="tw-mr-2.5 tw-font-inter-semibold tw-text-[1.09375rem] tw-leading-[1.3125rem] xl:tw-text-[1.188rem] xl:tw-leading-[1.188rem]"
            >Get Free Consultation
          </span>
        </router-link>
      </div>
      <div class="tw-relative lg:tw-flex-[45%]">
        <img
          :src="ctaImage400w"
          :srcset="`${ctaImage400w} 400w, ${ctaImage800w} 800w`"
          class="tw-mt-8 tw-h-full tw-w-full md:tw-mt-0"
          alt="cta-image"
        />
        <div
          class="gradient-border tw-absolute tw-bottom-4 tw-left-[0.5rem] tw-h-[55px] tw-border-l-[13px] md:tw-hidden"
        ></div>
      </div>

      <router-link
        class="gradient-btn consultation-btn tw-mt-10 tw-flex tw-w-max tw-items-center tw-rounded-full tw-py-[1.25rem] tw-text-center md:tw-hidden"
        :to="contactURL"
      >
        <span
          class="tw-mr-2.5 tw-font-inter-semibold tw-text-[1.125rem] tw-leading-[1.125rem]"
          >Get Free Consultation
        </span>
      </router-link>
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
      contactURL: "/contact",
    };
  },
};
</script>
<style scoped>
.gradient-border {
  border-image: linear-gradient(to bottom, #f2709c, #ff835b) 1;
}
</style>
