<template>
  <section
    class="lg:tw-bg-gradient-to-b tw-from-[#FFFFFF]/[0.05] tw-via-[#F98584]/[0.05] tw-to-[#F47695]/[0.05] tw-mt-[100px] tw-pb-10"
  >
    <div class="tw-relative 2xl:tw-container">
      <img
        :src="ctaimage400w"
        :srcset="`${ctaimage400w} 400w`"
        class="tw-absolute tw-top-[20.5rem] xl:tw-top-[25.8rem] 2xl:tw-top-[24.5rem] xll:tw-top-[22rem] tw-left-0 tw-max-w-[178px] xl:tw-max-w-[190px] 2xl:tw-max-w-[215px] xll:tw-max-w-[253px] tw-hidden lg:tw-block tw-object-cover"
        alt="thankyou-cta"
      />
      <img
        :src="ctaimage1400w"
        :srcset="`${ctaimage1400w} 400w`"
        class="tw-absolute tw-top-[20.5rem] xl:tw-top-[25rem] 2xl:tw-top-[23rem] xll:tw-top-[21rem] tw-right-0 tw-max-w-[165px] xl:tw-max-w-[188px] 2xl:tw-max-w-[220px] xll:tw-max-w-[250px] tw-hidden lg:tw-block tw-object-cover"
        alt="thankyou-cta-1"
      />
      <hr
        class="tw-absolute tw-top-[30.8rem] xl:tw-top-[36.8rem] tw-w-full tw-h-[3px] tw-hidden lg:tw-block tw-border-[1px] tw-bg-gray-500"
      />
      <div
        class="tw-container md:tw-mx-auto lg:tw-max-w-[652px] xl:tw-max-w-[797px]"
      >
        <p
          class="tw-text-center tw-font-inter-bold tw-text-[1.125rem] md:tw-text-[1.563rem] xl:tw-text-[1.875rem] tw-leading-[1.688rem] md:tw-leading-[2.375rem] xl:tw-leading-[2.813rem] tw-text-black-core/[0.87]"
        >
          We just want you to make an informed decision and we hope this
          information might help you get a clear understanding of the
          exceptional customer experiences we offer at
          <span class="tw-font-iinter-medium v2-canopas-gradient-text"
            >Canopas.</span
          >
        </p>
        <p
          class="tw-mt-8 tw-text-center tw-font-inter-regular tw-text-[1rem] md:tw-text-[1.375rem] xl:tw-text-[1.875rem] tw-leading-[1.5rem] md:tw-leading-[2.188rem] xl:tw-leading-[2.813rem] tw-text-black-core/[0.87]"
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
          class="tw-mt-10 tw-mx-auto tw-w-[30%] tw-h-[30%] lg:tw-hidden"
          alt="thankyou-cta-image"
        />

        <div
          class="tw-sticky tw-mx-auto tw-mt-0 lg:tw-mt-12 tw-rounded-full tw-p-3 tw-text-center gradient-btn consultation-btn tw-z-[3] tw-cursor-pointer"
          @click.native="openCalendlyIframe()"
        >
          <span
            class="tw-mr-2.5 tw-font-inter-bold tw-text-[1.875rem] md:tw-text-[2.75rem] tw-leading-[3rem] md:tw-leading-[4rem]"
            >Schedule a Meeting
          </span>
        </div>
        <p
          class="tw-mt-10 tw-text-center tw-font-inter-bold tw-text-[1.125rem] md:tw-text-[1.375rem] xl:tw-text-[1.875rem] tw-leading-[1.688rem] md:tw-leading-[2.188rem] xl:tw-leading-[2.813rem] tw-text-black-core/[0.87]"
        >
          If you'd like to propose an alternate time, just let me know.
        </p>

        <ul
          class="tw-mt-8 tw-mx-2 tw-pl-[2rem] tw-list-disc tw-list-inside tw--indent-[2em] xl:tw--indent-[1.2em] tw-space-y-4 tw-text-[1rem] md:tw-text-[1.125rem] xl:tw-text-[1.75rem] tw-leading-[1.5rem] md:tw-leading-[1.688rem] xl:tw-leading-[2.625rem] tw-text-black-core/[0.87] tw-font-inter-regular"
        >
          <li class="">
            <span class="tw-font-inter-bold">
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

        <div class="tw-flex tw-items-end tw-flex-col tw-mt-10 tw-pr-2">
          <p
            class="tw-text-left v2-canopas-gradient-text tw-font-inter-bold tw-text-[1.5rem] md:tw-text-[1.688rem] xl:tw-text-[1.875rem] tw-leading-[1.8rem] md:tw-leading-[1.9rem] xl:tw-leading-[2.3rem]"
          >
            Thanks, <br /><span
              class="tw-mt-0 tw-font-inter-medium tw-text-[1.125rem] md:tw-text-[1.25rem] xl:tw-text-[1.5rem] tw-leading-[1.688rem] md:tw-leading-[2rem] xl:tw-leading-[2.25rem] tw-text-black-core/[0.87]"
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
          class="modal-mask tw-fixed tw-top-0 tw-left-0 tw-w-full tw-h-full tw-table mask tw-z-[1] tw-bg-[#00000080] tw-z-[5]"
        >
          <div
            class="tw-flex tw-mx-auto tw-left-auto sm:tw-mx-auto tw-h-full login-modal modal-xl"
            role="document"
          >
            <div
              class="tw-relative tw-flex tw-flex-col tw-w-full tw-border-1 tw-border-gray tw-border-solid tw-rounded-md tw-rounded-3xl tw-bg-white tw-bg-clip-padding tw-outline-0"
            >
              <div class="tw-relative tw-flex-auto">
                <CalendlyIframe
                  class="tw-w-full tw-h-screen tw-overflow-hidden"
                />

                <button
                  type="button"
                  class="tw-absolute tw-right-2.5 tw-top-0 close modal-close-btn tw-border-none tw-text-pink-300 tw-text-4xl tw-font-light tw-bg-transparent focus:tw-outline-none"
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
      this.mixpanel.track("tap_thankyou_schedule_meeting");
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
