<template>
  <section
    class="lg:relative mt-10 lg:mb-[10%] xll:mb-[5%] bg-gradient-to-t from-[#FF835B]/[0.08] to-[#F2709C]/[0.08] pt-8 pb-8 lg:bg-none"
  >
    <div class="w-full h-full">
      <img
        v-if="width > 991"
        :src="contactbgImages[0]"
        :srcset="`${contactbgImages[0]} 1200w, ${contactbgImages[1]} 1600w`"
        loading="lazy"
        class="hidden lg:block absolute top-[217px] xll:top-[50px] left-0 -z-[1] object-cover xll:object-scale-down"
        alt="background-image"
      />
    </div>
    <div
      class="container flex flex-col lg:flex-row-reverse items-center lg:gap-6"
    >
      <div class="flex flex-col lg:w-[59%] xl:w-[50%]">
        <h2
          class="font-inter-regular text-[1.875rem] md:text-[2.5rem] lg:text-[3rem] leading-[2.813rem] md:leading-[3.75rem] lg:leading-[4.5rem] text-center lg:text-left text-black-core/[0.87]"
        >
          Do you want to
          <span class="font-inter-bold lg:text-[3.438rem] lg:leading-[5.156rem]"
            >provide your customers</span
          >
          with value at
          <span class="font-inter-bold lg:text-[3.438rem] lg:leading-[5.156rem]"
            >their fingertips?</span
          >
        </h2>
        <div class="hidden lg:block mt-6">
          <nuxt-link
            class="relative inline-block animation-underline cursor-pointer bg-gradient-[1deg] pb-3 text-[2.5rem] leading-[3.75rem] v2-canopas-gradient-text font-inter-bold gradient-underline no-underline cursor-pointer after:absolute after:bottom-0 after:left-0 after:w-full after:h-[5px] after:bg-gradient-to-l after:to-[#F2709C] after:from-[#FF835B] after:origin-bottom-left after:scale-x-100 after:duration-500 after:delay-150 hover:after:animate-underlineOut"
            to="/contact"
            @click.native="$mixpanel.track('tap_services_contact')"
            ><span> Contact Us Today</span></nuxt-link
          >
        </div>
      </div>
      <div class="flex flex-col items-center lg:w-[41%] xl:w-[50%]">
        <img
          class="h-full w-full object-cover"
          :src="contactImages[0]"
          :srcset="`${contactImages[0]} 400w, ${contactImages[1]} 800w`"
          alt="contact-view-image"
          loading="lazy"
        />
        <div class="mt-[-35px] lg:hidden">
          <nuxt-link
            class="flex items-center w-max rounded-full py-3 px-3 gradient-btn consultation-btn"
            to="/contact"
            @click.native="$mixpanel.track('tap_services_contact')"
          >
            <span
              class="mr-2.5 font-normal text-[1rem] leading-[1.1875rem] md:text-[1.5rem] md:leading-8 font-inter-semibold"
              >Contact Us Today</span
            >
          </nuxt-link>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import contact400w from "@/assets/images/services/contact/contact-400w.webp";
import contact800w from "@/assets/images/services/contact/contact-800w.webp";
import contactbg1200w from "@/assets/images/services/contact/bg/contact-bg-1200w.webp";
import contactbg1600w from "@/assets/images/services/contact/bg/contact-bg-1600w.webp";

export default {
  data() {
    return {
      width: 991,
      contactImages: [contact400w, contact800w],
      contactbgImages: [contactbg1200w, contactbg1600w],
    };
  },
  inject: ["mixpanel"],
  mounted() {
    this.width = window.innerWidth;
  },
};
</script>
