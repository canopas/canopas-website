<template>
  <div
    class="tw-relative tw-mt-16 md:tw-mt-40 xl:tw-mt-16 tw-h-[217px] md:tw-h-[376px] xl:tw-h-[476px]"
  >
    <img
      :src="BackgroundImage[0]"
      :srcset="`${BackgroundImage[0]} 400w, ${BackgroundImage[1]} 800w, ${BackgroundImage[2]} 1200w, ${BackgroundImage[3]} 2400w`"
      alt="Background-Image"
      class="tw-h-full tw-w-full tw-object-cover"
      loading="lazy"
    />
    <div
      class="tw-container tw-absolute tw-top-4 md:tw-top-20 xl:tw-top-32 tw-w-full tw-p-4 tw-text-center tw-text-[#FFF] xs:tw-left-[13%] xll:tw-left-[24%] 3xl:tw-left-[33%] xs:tw-w-[75%]"
    >
      <div class="tw-m-auto">
        <p class="tw-text-white header-1">Have an app idea?</p>

        <div class="tw-mt-2 tw-text-white/[0.8] sub-h1-regular md:tw-my-4">
          See if we can make it a reality.
        </div>
      </div>
      <div class="tw-flex tw-justify-center tw-mt-2">
        <router-link
          class="gradient-btn primary-btn"
          to="/contact"
          @click.native="mixpanel.track('tap_android_development_cta')"
        >
          <span class="sub-h3-semibold xl:sub-h4-semibold">Let's Chat </span>
        </router-link>
      </div>
    </div>
  </div>
</template>

<script type="module">
import background400 from "@/assets/images/andriod-app-development/cta/cta2background-400w.webp";
import background800 from "@/assets/images/andriod-app-development/cta/cta2background-800w.webp";
import background1200 from "@/assets/images/andriod-app-development/cta/cta2background-1200w.webp";
import background2400 from "@/assets/images/andriod-app-development/cta/cta2background-2400w.webp";

export default {
  data() {
    return {
      BackgroundImage: [
        background400,
        background800,
        background1200,
        background2400,
      ],
    };
  },
  inject: ["mixpanel"],
};
</script>
