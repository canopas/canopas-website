<template>
  <div
    class="3xl:outer-container relative h-[15.625rem] md:h-[17.188rem] lg:mt-60 lg:h-[29.75rem] bg-white-smoke"
  >
    <img
      :src="BackgroundImage[0]"
      :srcset="`${BackgroundImage[0]} 400w, ${BackgroundImage[1]} 800w, ${BackgroundImage[2]} 1200w, ${BackgroundImage[3]} 2400w`"
      alt="Background-Image"
      class="h-full w-full object-cover"
      loading="lazy"
    />
    <div
      class="container absolute top-[1.2rem] w-full p-4 text-center md:left-0 md:top-1/2 md:top-16 md:-mt-6 md:text-center lg:!left-1/2 lg:!top-[57%] xll:!top-[60%] lg:-translate-x-1/2 lg:-translate-y-1/2 lg:transform xl:w-[60%] 2xl:w-[49%] xs:left-[13%] xs:w-[75%] xl2:w-[55%] 2xll:w-[45%] xll:w-[61%]"
    >
      <div class="flex w-full lg:mb-10">
        <div class="m-auto">
          <p
            class="mobile-header-2 lg:desk-header-1 mb-4 lg:mb-6 text-black-87"
          >
            Your vision. Our expertise.
          </p>

          <div class="sub-h1-regular lg:mobile-header-2-regular text-black-60">
            Together, Let's redefine what's possible on the web.
          </div>
        </div>
      </div>
      <div class="mx-auto mt-6">
        <nuxt-link
          class="gradient-btn primary-btn m-0 mx-auto flex w-max items-center p-3 px-7 text-center"
          to="/contact"
          @click.native="$mixpanel.track('tap_ios_development_cta')"
        >
          <span class="mr-2.5 sub-h3-semibold lg:sub-h1-semibold"
            >Unlock Excellence
          </span>
        </nuxt-link>
      </div>
    </div>
  </div>
</template>

<script type="module">
import background400 from "@/assets/images/ios-app-development/cta/cta2background-400w.webp";
import background800 from "@/assets/images/ios-app-development/cta/cta2background-800w.webp";
import background1200 from "@/assets/images/ios-app-development/cta/cta2background-1200w.webp";
import background2400 from "@/assets/images/ios-app-development/cta/cta2background-2400w.webp";

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
