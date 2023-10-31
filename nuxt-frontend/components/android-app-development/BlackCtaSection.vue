<template>
  <div
    class="relative mt-16 md:mt-40 xl:mt-16 h-[217px] md:h-[376px] xl:h-[476px] 3xl:outer-container"
  >
    <img
      :src="BackgroundImage[0]"
      :srcset="`${BackgroundImage[0]} 400w, ${BackgroundImage[1]} 800w, ${BackgroundImage[2]} 1200w, ${BackgroundImage[3]} 2400w`"
      alt="Background-Image"
      class="h-full w-full object-cover"
      loading="lazy"
    />
    <div
      class="container absolute top-4 md:top-20 xl:top-32 w-full p-4 text-center text-[#FFF] xs:left-[13%] xll:left-[24%] xs:w-[75%]"
    >
      <div class="m-auto">
        <p class="text-white header-1">Have an app idea?</p>

        <div class="mt-2 text-white/[0.8] sub-h1-regular md:my-4">
          See if we can make it a reality.
        </div>
      </div>
      <div class="flex justify-center mt-2">
        <router-link
          class="gradient-btn primary-btn"
          to="/contact"
          @click.native="$mixpanel.track('tap_android_development_cta')"
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
