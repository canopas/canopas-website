<template>
  <section
    class="lg:bg-gradient-to-b from-[#FFFFFF]/[0.05] via-[#F98584]/[0.05] to-[#F47695]/[0.05] mt-[100px] pb-10"
  >
    <div class="relative 2xl:container">
      <img
        :src="ctaimage400w"
        :srcset="`${ctaimage400w} 400w`"
        class="absolute top-[20.5rem] xl:top-[25.8rem] 2xl:top-[24.5rem] xll:top-[22rem] left-0 max-w-[178px] xl:max-w-[190px] 2xl:max-w-[215px] xll:max-w-[253px] hidden lg:block object-cover"
        alt="thankyou-cta"
      />
      <img
        :src="ctaimage1400w"
        :srcset="`${ctaimage1400w} 400w`"
        class="absolute top-[20.5rem] xl:top-[25rem] 2xl:top-[23rem] xll:top-[21rem] right-0 max-w-[165px] xl:max-w-[188px] 2xl:max-w-[220px] xll:max-w-[250px] hidden lg:block object-cover"
        alt="thankyou-cta-1"
      />
      <hr
        class="absolute top-[30.8rem] xl:top-[36.8rem] w-full h-[3px] hidden lg:block border-[1px] bg-gray-500"
      />
      <div class="container md:mx-auto lg:max-w-[652px] xl:max-w-[797px]">
        <p
          class="text-center font-inter-bold text-[1.125rem] md:text-[1.563rem] xl:text-[1.875rem] leading-[1.688rem] md:leading-[2.375rem] xl:leading-[2.813rem] text-black-core/[0.87]"
        >
          We just want you to make an informed decision and we hope this
          information might help you get a clear understanding of the
          exceptional customer experiences we offer at
          <span class="font-iinter-medium v2-canopas-gradient-text"
            >Canopas.</span
          >
        </p>
        <p
          class="mt-8 text-center font-inter-regular text-[1rem] md:text-[1.375rem] xl:text-[1.875rem] leading-[1.5rem] md:leading-[2.188rem] xl:leading-[2.813rem] text-black-core/[0.87]"
        >
          We would love the opportunity to discuss how we can help you achieve
          your business goals through our exceptional customer experience
          services.
          <br />
          <br />
          If you're interested in learning more about how we can collaborate,
          <span class="v2-canopas-gradient-text"
            >please schedule a call with us</span
          >
          to learn more about your specific needs.
        </p>

        <img
          :src="ctaimage1400w"
          :srcset="`${ctaimage1400w} 400w`"
          class="mt-10 mx-auto w-[30%] h-[30%] lg:hidden"
          alt="thankyou-cta-image"
        />

        <div
          class="sticky mx-auto mt-0 lg:mt-12 rounded-full p-3 text-center gradient-btn consultation-btn z-[3] cursor-pointer"
          @click.native="openCalendlyIframe()"
        >
          <span
            class="mr-2.5 font-inter-bold text-[1.875rem] md:text-[2.75rem] leading-[3rem] md:leading-[4rem]"
            >Schedule a Meeting
          </span>
        </div>
        <p
          class="mt-10 text-center font-inter-bold text-[1.125rem] md:text-[1.375rem] xl:text-[1.875rem] leading-[1.688rem] md:leading-[2.188rem] xl:leading-[2.813rem] text-black-core/[0.87]"
        >
          If you'd like to propose an alternate time, just let me know by
          replying to the email you've got from us.
        </p>

        <ul
          class="mt-8 mx-2 pl-8 list-disc list-inside -indent-[2em] xl:-indent-[1.2em] space-y-4 text-[1rem] md:text-[1.125rem] xl:text-[1.75rem] leading-6 md:leading-[1.688rem] xl:leading-[2.625rem] text-black-core/[0.87] font-inter-regular"
        >
          <li class="">
            <span class="font-inter-bold">
              We offer Two Week of Money Back Guarantee.</span
            >
          </li>
          <li>
            <span
              >If you have any questions or concerns, please don't hesitate to
              reach out to me directly. I'm here to help and support you
              throughout the process.</span
            >
          </li>
          <li>
            Thank you again for considering our company for your needs. We look
            forward to the opportunity to work with you and to build a long and
            successful partnership.
          </li>
        </ul>

        <div class="flex items-end flex-col mt-10 pr-2">
          <p
            class="text-left v2-canopas-gradient-text font-inter-bold text-[1.5rem] md:text-[1.688rem] xl:text-[1.875rem] leading-[1.8rem] md:leading-[1.9rem] xl:leading-[2.3rem]"
          >
            Thanks, <br /><span
              class="mt-0 font-inter-medium text-[1.125rem] md:text-[1.25rem] xl:text-[1.5rem] leading-[1.688rem] md:leading-8 xl:leading-[2.25rem] text-black-core/[0.87]"
            >
              Jimmy</span
            >
          </p>
        </div>
      </div>
    </div>
    <!-- Show Calendly Iframe -->
    <div v-if="openCalendlyIframeModal">
      <transition name="modal">
        <div
          class="modal-mask fixed top-0 left-0 w-full h-full table mask z-[1] bg-[#00000080] z-[5]"
        >
          <div
            class="flex mx-auto left-auto sm:mx-auto h-full login-modal modal-xl"
            role="document"
          >
            <div
              class="relative flex flex-col w-full border-1 border-gray border-solid rounded-md rounded-3xl bg-white bg-clip-padding outline-0"
            >
              <div class="relative flex-auto">
                <CalendlyIframe class="w-full h-screen overflow-hidden" />

                <button
                  type="button"
                  class="absolute right-2.5 top-0 close modal-close-btn border-none text-pink-300 text-4xl font-light bg-transparent focus:outline-none"
                  data-dismiss="modal"
                  aria-label="Close"
                >
                  <span aria-hidden="true" @click="closeCalendlyIframeModal"
                    >&times;</span
                  >
                </button>
              </div>
            </div>
          </div>
        </div>
      </transition>
    </div>
  </section>
</template>

<script>
import ctaimage400w from "@/assets/images/client-thankyou/cta-image-400w.webp";
import ctaimage1400w from "@/assets/images/client-thankyou/cta-image-1-400w.webp";
import CalendlyIframe from "../CalendlyIframe.vue";

export default {
  data() {
    return {
      ctaimage400w,
      ctaimage1400w,
      openCalendlyIframeModal: false,
    };
  },
  components: {
    CalendlyIframe,
  },
  inject: ["mixpanel"],
  methods: {
    openCalendlyIframe() {
      this.openCalendlyIframeModal = true;
      this.$mixpanel.track("tap_thankyou_schedule_meeting");
    },
    closeCalendlyIframeModal() {
      this.openCalendlyIframeModal = false;
    },
  },
};
</script>

<style scoped>
li::marker {
  font-size: 1.5rem;
  padding-top: 1rem;
}
</style>
