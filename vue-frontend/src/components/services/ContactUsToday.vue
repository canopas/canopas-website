<template>
  <section
    class="lg:tw-relative tw-mt-10 lg:tw-mb-[10%] xll:tw-mb-[5%] tw-bg-gradient-to-t tw-from-[#FF835B]/[0.08] tw-to-[#F2709C]/[0.08] tw-pt-8 tw-pb-8 lg:tw-bg-none"
  >
    <div class="tw-w-full tw-h-full">
      <img
        v-if="width > 991"
        :src="contactbgImages[0]"
        :srcset="`${contactbgImages[0]} 1200w, ${contactbgImages[1]} 1600w`"
        loading="lazy"
        class="tw-hidden lg:tw-block tw-absolute tw-top-[217px] xll:tw-top-[50px] tw-left-0 tw--z-[1] tw-object-cover xll:tw-object-scale-down"
        alt="background-image"
      />
    </div>
    <div
      class="tw-container tw-flex tw-flex-col lg:tw-flex-row-reverse tw-items-center lg:tw-gap-6"
    >
      <div class="tw-flex tw-flex-col lg:tw-w-[59%] xl:tw-w-[50%]">
        <h2
          class="tw-font-inter-regular tw-text-[1.875rem] md:tw-text-[2.5rem] lg:tw-text-[3rem] tw-leading-[2.813rem] md:tw-leading-[3.75rem] lg:tw-leading-[4.5rem] tw-text-center lg:tw-text-left tw-text-black-core/[0.87]"
        >
          Do you want to
          <span
            class="tw-font-inter-bold lg:tw-text-[3.438rem] lg:tw-leading-[5.156rem]"
            >provide your customers</span
          >
          with value at
          <span
            class="tw-font-inter-bold lg:tw-text-[3.438rem] lg:tw-leading-[5.156rem]"
            >their fingertips?</span
          >
        </h2>
        <div class="tw-hidden lg:tw-block tw-mt-6">
          <router-link
            class="tw-relative tw-inline-block animation-underline tw-cursor-pointer tw-bg-gradient-[1deg] tw-pb-3 tw-text-[2.5rem] tw-leading-[3.75rem] v2-canopas-gradient-text tw-font-inter-bold gradient-underline tw-no-underline tw-cursor-pointer after:tw-absolute after:tw-bottom-0 after:tw-left-0 after:tw-w-full after:tw-h-[5px] after:tw-bg-gradient-to-l after:tw-to-[#F2709C] after:tw-from-[#FF835B] after:tw-origin-bottom-left after:tw-scale-x-100 after:tw-duration-500 after:tw-delay-150 hover:after:tw-animate-underlineOut"
            to="/contact"
            @click.native="mixpanel.track('tap_services_contact')"
            ><span> Contact Us Today</span></router-link
          >
        </div>
      </div>
      <div
        class="tw-flex tw-flex-col tw-items-center lg:tw-w-[41%] xl:tw-w-[50%]"
      >
        <img
          class="tw-h-full tw-w-full tw-object-cover"
          :src="contactImages[0]"
          :srcset="`${contactImages[0]} 400w, ${contactImages[1]} 800w`"
          alt="contact-view-image"
          loading="lazy"
        />
        <div class="tw-mt-[-35px] lg:tw-hidden">
          <router-link
            class="tw-flex tw-items-center tw-w-max tw-rounded-full tw-py-3 tw-px-3 gradient-btn consultation-btn"
            to="/contact"
            @click.native="mixpanel.track('tap_services_contact')"
          >
            <span
              class="tw-mr-2.5 tw-font-normal tw-text-[1rem] tw-leading-[1.1875rem] md:tw-text-[1.5rem] md:tw-leading-8 tw-font-inter-semibold"
              >Contact Us Today</span
            >
          </router-link>
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
