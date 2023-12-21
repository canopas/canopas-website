<template>
  <div class="relative md:h-[376px] xl:h-[476px] 3xl:outer-container">
    <img
      :src="background"
      alt="Background-Image"
      class="h-full w-full object-cover"
      loading="lazy"
    />
    <div
      class="container absolute left-1/2 flex w-full -translate-x-1/2 -translate-y-[116%] xl:-translate-y-[120%] transform flex-row items-center lg:mt-4"
    >
      <div class="flex w-[66%]">
        <div class="w-[80%] py-5 text-white lg:w-[90%] xl:m-auto 2xl:w-9/12">
          <p class="mobile-header-2 lg:desk-header-2">Whether you need...</p>
          <ul class="pt-6 sub-h1-semibold lg:mobile-header-2-semibold">
            <li v-for="reason in reasons" :key="reason" class="flex">
              <span class="pr-1">*</span>
              <div class="sm:w-[90%]">{{ reason }}</div>
            </li>
          </ul>
          <div
            class="text-white-core-80 pb-8 pt-4 sub-h3-regular lg:mobile-header-3-regular xl:w-[85%]"
          >
            Bring us your toughest challenge and we'll show you the path to a
            sleek solution.
          </div>
        </div>
      </div>
      <div class="pt--[8rem] ml-auto xl:mx-auto">
        <router-link
          class="cta-button primary-btn border border-solid border-white bg-white hover:bg-transparent active:scale-[0.98]"
          to="/contact"
          @click.native="$mixpanel.track('tap_android_development_cta')"
        >
          <span class="sub-h1-semibold v2-canopas-gradient-text"
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
