<template>
  <div class="tw-relative tw-mb-16 tw-h-80 xll:tw-h-[30rem]">
    <img
      :src="background"
      alt="Background-Image"
      class="tw-h-full tw-w-full tw-object-cover"
      loading="lazy"
    />
    <div
      class="tw-container tw-absolute tw-left-1/2 tw-flex tw-w-full tw--translate-x-1/2 tw--translate-y-[100%] tw-transform tw-flex-row tw-items-center lg:tw-mt-4 xll:!tw--mt-[4rem]"
    >
      <div class="tw-flex tw-w-[66%]">
        <div
          class="tw-w-[80%] tw-py-[20px] tw-text-white lg:tw-w-[90%] xl:tw-m-auto 2xl:tw-w-9/12"
        >
          <p
            class="tw-font-inter-bold tw-text-[2.188rem] tw-leading-[2.625rem] lg:tw-text-[3.4375rem] lg:tw-leading-[3.4375rem]"
          >
            Whether you need...
          </p>
          <ul
            class="tw-pt-6 tw-font-inter-semibold tw-text-[1rem] tw-leading-[1.25rem] md:tw-text-[1.188rem] md:tw-leading-[1.5rem] lg:tw-text-[1.375rem] lg:tw-leading-[1.8125rem]"
          >
            <li v-for="reason in reasons" :key="reason" class="tw-flex">
              <span class="tw-pr-1">*</span>
              <div class="sm:tw-w-[90%]">{{ reason }}</div>
            </li>
          </ul>
          <div
            class="tw-pb-8 tw-pt-6 tw-font-inter-regular tw-text-[0.875rem] tw-leading-[1.125rem] md:tw-text-[1.125rem] md:tw-leading-[1.375rem] lg:tw-text-[1.375rem] lg:tw-leading-[1.813rem] xl:tw-w-[85%]"
          >
            Bring us your toughest challenge and we'll show you the path to a
            sleek solution.
          </div>
        </div>
      </div>
      <div class="tw-ml-auto xl:tw-mx-auto">
        <router-link
          class="cta-button tw-p-y-[0.8125rem] tw-mx-auto tw-flex tw-w-max tw-items-center tw-rounded-full tw-border-[1px] tw-border-solid tw-border-white tw-bg-white tw-p-3 tw-px-[1.875rem] tw-text-center tw-font-inter-medium tw-font-inter-medium tw-text-[1rem] tw-font-normal tw-leading-[1.1875rem] !tw-tracking-[0] hover:tw-bg-transparent hover:tw-text-white active:tw-scale-[0.98] md:tw-text-[1.09375rem] md:tw-leading-[1.3125rem] lg:tw-mx-0 lg:tw-text-[1.1875rem] lg:tw-leading-[1.436875rem]"
          to="/contact"
          @click.native="mixpanel.track('tap_android_development_cta')"
        >
          <span
            class="tw-from-[#ff9472] tw-to-[#f2709c] tw-bg-clip-text tw-font-inter-medium tw-text-transparent tw-bg-gradient-[270.11deg]"
            >Let's Chat</span
          >
        </router-link>
      </div>
    </div>
  </div>
</template>

<script type="module">
import background from "@/assets/images/andriod-app-development/cta/pinkBackground.webp";
export default {
  data() {
    return {
      background,
      reasons: [
        "High-performing mobile apps",
        "Bulletproof cloud solutions",
        "Custom solutions for your business.",
      ],
    };
  },
  inject: ["mixpanel"],
};
</script>
