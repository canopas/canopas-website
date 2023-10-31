<template>
  <section>
    <div
      class="container flex flex-col lg:flex-row items-center mt-[100px] 2xl:gap-[5.5rem]"
    >
      <div
        class="md:w-[80%] lg:w-[48%] xl:w-[55%] 2xl:w-[50%] lg:pr-10 2xl:pr-0"
      >
        <img
          :src="CTAImage[0]"
          :srcset="`${CTAImage[0]} 400w, ${CTAImage[1]} 800w, ${CTAImage[2]} 1600w, ${CTAImage[3]} 2400w`"
          alt="aboutus-cta-image"
          class="w-full h-full object-contain"
          loading="lazy"
        />
      </div>
      <div
        class="flex flex-col items-center lg:items-start lg:mt-[5.5rem] xl:mt-10 lg:w-[52%] xl:w-[45%] 2xl:w-[50%]"
      >
        <div
          class="md:w-[70%] lg:w-full text-center lg:text-left text-[1.875rem] md:text-[2.5rem] 2xl:text-[3.125rem] leading-[2.813rem] md:leading-[3.438rem] 2xl:leading-[4.688rem] font-inter-bold text-black-core/[0.87]"
        >
          Ready to take the next step and work together?
        </div>
        <div
          class="md:w-[70%] lg:w-full pt-4 lg:text-left text-center text-[1rem] md:text-[1.125rem] 2xl:text-[1.188rem] leading-6 md:leading-[1.656rem] 2xl:leading-[1.781rem] font-inter-medium text-black-core/[0.87]"
        >
          We'd love to ask some questions, learn a little more about what you do
          and who you serve and kick around some ideas for you.
        </div>
        <div class="py-8">
          <router-link
            class="flex items-center w-max m-0 rounded-full p-3 text-center gradient-btn consultation-btn"
            to="/contact"
            @click.native="mixpanel.track('tap_about_cta_contact')"
          >
            <span
              class="mr-2.5 font-normal text-[1rem] leading-[1.1875rem] md:text-[1.09375rem] md:leading-[1.3125rem] lg:text-[1.1875rem] lg:leading-[1.4375rem] font-inter-medium"
              >Talk to our experts
            </span>
          </router-link>
        </div>
      </div>
    </div>
  </section>
</template>
<script type="module">
import cta400 from "@/assets/images/about/cta-400w.webp";
import cta800 from "@/assets/images/about/cta-800w.webp";
import cta1600 from "@/assets/images/about/cta-1600w.webp";
import cta2400 from "@/assets/images/about/cta-2400w.webp";
export default {
  data() {
    return {
      CTAImage: [cta400, cta800, cta1600, cta2400],
    };
  },
  inject: ["mixpanel"],
};
</script>
