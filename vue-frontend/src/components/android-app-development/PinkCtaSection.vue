<template>
  <div
    class="tw-relative md:tw-h-[376px] xl:tw-h-[476px] 3xl:tw-outer-container"
  >
    <img
      :src="background"
      alt="Background-Image"
      class="tw-h-full tw-w-full tw-object-cover"
      loading="lazy"
    />
    <div
      class="tw-container tw-absolute tw-left-1/2 tw-flex tw-w-full tw--translate-x-1/2 tw--translate-y-[116%] xl:tw--translate-y-[120%] tw-transform tw-flex-row tw-items-center lg:tw-mt-4"
    >
      <div class="tw-flex tw-w-[66%]">
        <div
          class="tw-w-[80%] tw-py-5 tw-text-white lg:tw-w-[90%] xl:tw-m-auto 2xl:tw-w-9/12"
        >
          <p class="header-2">Whether you need...</p>
          <ul class="tw-pt-6 sub-h1-semibold">
            <li v-for="reason in reasons" :key="reason" class="tw-flex">
              <span class="tw-pr-1">*</span>
              <div class="sm:tw-w-[90%]">{{ reason }}</div>
            </li>
          </ul>
          <div
            class="tw-text-white/[0.8] tw-pb-8 tw-pt-4 sub-h3-regular xl:tw-w-[85%]"
          >
            Bring us your toughest challenge and we'll show you the path to a
            sleek solution.
          </div>
        </div>
      </div>
      <div class="tw-pt--[8rem] tw-ml-auto xl:tw-mx-auto">
        <router-link
          class="cta-button primary-btn tw-border tw-border-solid tw-border-white tw-bg-white hover:tw-bg-transparent active:tw-scale-[0.98]"
          to="/contact"
          @click.native="mixpanel.track('tap_android_development_cta')"
        >
          <span class="sub-h4-semibold v2-canopas-gradient-text"
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
