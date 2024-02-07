<template>
  <section
    class="lg:bg-gradient-to-b from-white-core-5 via-orange-peel to-light-crimson mt-16 pb-10"
  >
    <div class="relative 2xl:container">
      <img
        :src="ctaimage400w"
        :srcset="`${ctaimage400w} 400w`"
        class="absolute top-[27.7rem] xl:top-96 2xl:top-[18.6rem] left-0 max-w-[11.125rem] xl:max-w-[11.875rem] 2xl:max-w-[13.438rem] hidden lg:block object-cover"
        alt="thankyou-cta"
      />
      <img
        :src="ctaimage1400w"
        :srcset="`${ctaimage1400w} 400w`"
        class="absolute top-[27.7rem] xl:top-[23.3rem] 2xl:top-[17.2rem] right-0 max-w-[10.313rem] xl:max-w-[11.75rem] 2xl:max-w-[13.75rem] hidden lg:block object-cover"
        alt="thankyou-cta-1"
      />
      <hr
        class="absolute top-[38rem] xl:top-[35rem] 2xl:top-[31rem] w-full h-[3px] hidden lg:block border bg-gray-500"
      />
      <div
        class="container md:mx-auto lg:max-w-[40.75rem] xl:max-w-[49.813rem] 2xl:max-w-[55.75rem]"
      >
        <p
          class="text-center mobile-header-3-semibold lg:desk-header-3 text-black-87"
        >
          We just want you to make an informed decision and we hope this
          information might help you get a clear understanding of the
          exceptional customer experiences we offer at
          <span class="text-primary-1">Canopas.</span>
        </p>
        <div
          class="mt-4 lg:mt-10 text-center sub-h1-regular lg:mobile-header-2-regular text-black-60"
        >
          <p>
            We would love the opportunity to discuss how we can help you achieve
            your business goals through our exceptional customer experience
            services.
          </p>
          <p class="mt-4 lg:mt-10">
            If you're interested in learning more about how we can collaborate,
            <span
              class="v2-canopas-gradient-text sub-h1-semibold lg:lg:mobile-header-2-semibold"
              >please schedule a call with us</span
            >
            to learn more about your specific needs.
          </p>
        </div>

        <img
          :src="ctaimage1400w"
          :srcset="`${ctaimage1400w} 400w`"
          class="mt-8 mx-auto w-[30%] h-[30%] lg:hidden"
          alt="thankyou-cta-image"
        />

        <div
          class="sticky mx-auto mt-0 lg:mt-16 rounded-xl p-3 text-center gradient-btn consultation-btn z-[3] cursor-pointer"
          @click.native="openCalendlyIframe()"
        >
          <span class="mr-2.5 mobile-header-1 lg:desk-header-2 xl:desk-header-1"
            >Schedule a Meeting
          </span>
        </div>
        <p
          class="mt-4 lg:mt-6 text-center text-black-60 sub-h3-regular lg:mobile-header-3-medium"
        >
          If you'd like to propose an alternate time, just let me know by
          replying to the email you've got from us.
        </p>

        <ul
          class="mt-8 lg:mt-10 mx-2 pl-8 list-disc list-inside -indent-[2em] xl:-indent-[1.2em] space-y-4 sub-h3-medium lg:mobile-header-2-regular"
        >
          <li>
            <span class="sub-h3-semibold lg:mobile-header-2-semibold">
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

        <div class="flex items-end flex-col mt-8 pr-2">
          <p class="text-left text-primary-1 sub-h1-semibold lg:desk-header-3">
            Thanks, <br /><span
              class="mt-0 sub-h1-regular lg:mobile-header-2-regular text-black-87"
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
