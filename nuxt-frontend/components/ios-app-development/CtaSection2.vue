<template>
  <div
    class="relative mb-16 h-48 md:h-72 xll:h-[30rem] bg-gradient-to-l from-[#FF835B]/[10%] to-[#F2709C]/[10%]"
  >
    <img
      :src="BackgroundImage[0]"
      :srcset="`${BackgroundImage[0]} 400w, ${BackgroundImage[1]} 800w, ${BackgroundImage[2]} 1200w, ${BackgroundImage[3]} 2400w`"
      alt="Background-Image"
      class="h-full xll:h-full w-full object-cover"
      loading="lazy"
    />
    <div
      class="container absolute top-[1.2rem] w-full p-4 text-center text-black-core/[0.87] md:left-0 md:top-1/2 md:top-16 md:-mt-6 md:text-center lg:!left-[50%] lg:!top-[57%] xll:!top-[60%] lg:-translate-x-1/2 lg:-translate-y-1/2 lg:transform xl:w-[60%] 2xl:w-[49%] xs:left-[13%] xs:w-[75%] xl2:w-[55%] 2xll:w-[45%] xll:w-[61%]"
    >
      <div class="flex w-full">
        <div class="m-auto">
          <p
            class="font-inter-semibold text-[1.375rem] leading-[2.0625rem] md:font-inter-bold md:text-[2.40625rem] lg:text-[3.4375rem] md:leading-[3.11125rem] lg:leading-[4.16rem]"
          >
            Have an app idea?
          </p>

          <div
            class="font-inter-regular text-base leading-6 md:my-4 md:font-inter-medium md:text-xl md:leading-[1.875rem] lg:text-2xl lg:leading-9"
          >
            See if we can make it a reality.
          </div>
        </div>
      </div>
      <div class="mx-auto mt-4">
        <nuxt-link
          class="gradient-btn consultation-btn m-0 mx-auto flex w-max items-center rounded-full p-3 px-7 text-center"
          to="/contact"
          @click.native="$mixpanel.track('tap_ios_development_cta')"
        >
          <span
            class="mr-2.5 font-inter-medium text-base leading-[1.21rem] md:block md:text-[1.09375rem] md:leading-[1.3125rem] lg:text-[1.1875rem] lg:leading-[1.4375rem]"
            >Let's Chat
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
