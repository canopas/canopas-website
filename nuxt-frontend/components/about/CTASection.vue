<template>
  <section>
    <div
      class="tw-container tw-flex tw-flex-col lg:tw-flex-row tw-items-center tw-mt-[100px] 2xl:tw-gap-[5.5rem]"
    >
      <div
        class="md:tw-w-[80%] lg:tw-w-[48%] xl:tw-w-[55%] 2xl:tw-w-[50%] lg:tw-pr-10 2xl:tw-pr-0"
      >
        <img
          :src="CTAImage[0]"
          :srcset="`${CTAImage[0]} 400w, ${CTAImage[1]} 800w, ${CTAImage[2]} 1600w, ${CTAImage[3]} 2400w`"
          alt="aboutus-cta-image"
          class="tw-w-full tw-h-full tw-object-contain"
          loading="lazy"
        />
      </div>
      <div
        class="tw-flex tw-flex-col tw-items-center lg:tw-items-start lg:tw-mt-[5.5rem] xl:tw-mt-10 lg:tw-w-[52%] xl:tw-w-[45%] 2xl:tw-w-[50%]"
      >
        <div
          class="md:tw-w-[70%] lg:tw-w-full tw-text-center lg:tw-text-left tw-text-[1.875rem] md:tw-text-[2.5rem] 2xl:tw-text-[3.125rem] tw-leading-[2.813rem] md:tw-leading-[3.438rem] 2xl:tw-leading-[4.688rem] tw-font-inter-bold tw-text-black-core/[0.87]"
        >
          Ready to take the next step and work together?
        </div>
        <div
          class="md:tw-w-[70%] lg:tw-w-full tw-pt-4 lg:tw-text-left tw-text-center tw-text-[1rem] md:tw-text-[1.125rem] 2xl:tw-text-[1.188rem] tw-leading-6 md:tw-leading-[1.656rem] 2xl:tw-leading-[1.781rem] tw-font-inter-medium tw-text-black-core/[0.87]"
        >
          We'd love to ask some questions, learn a little more about what you do
          and who you serve and kick around some ideas for you.
        </div>
        <div class="tw-py-8">
          <nuxt-link
            class="tw-flex tw-items-center tw-w-max tw-m-0 tw-rounded-full tw-p-3 tw-text-center gradient-btn consultation-btn"
            to="/contact"
            @click.native="mixpanel.track('tap_about_cta_contact')"
          >
            <span
              class="tw-mr-2.5 tw-font-normal tw-text-[1rem] tw-leading-[1.1875rem] md:tw-text-[1.09375rem] md:tw-leading-[1.3125rem] lg:tw-text-[1.1875rem] lg:tw-leading-[1.4375rem] tw-font-inter-medium"
              >Talk to our experts
            </span>
          </nuxt-link>
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
  // // inject: ["mixpanel"],
};
</script>
