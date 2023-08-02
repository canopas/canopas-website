<template>
  <div class="tw-relative tw-mb-16 tw-h-48 md:tw-h-72 xll:tw-h-[30rem]">
    <img
      :src="BackgroundImage[0]"
      :srcset="`${BackgroundImage[0]} 400w, ${BackgroundImage[1]} 800w, ${BackgroundImage[2]} 1200w, ${BackgroundImage[3]} 2400w`"
      alt="Background-Image"
      class="tw-h-full lg:tw-h-[22.5rem] xll:tw-h-full tw-w-full tw-object-cover"
      loading="lazy"
    />
    <div
      class="tw-container tw-absolute tw-top-[1rem] tw-w-full tw-flex-col tw-items-center tw-justify-center tw-p-4 tw-text-center tw-text-[#FFF] md:tw-left-0 md:tw-top-1/2 md:tw-top-[4rem] md:tw--mt-10 md:tw-text-center lg:!tw-left-[50%] lg:!tw-top-[72%] xll:!tw-top-[60%] lg:tw--translate-x-1/2 lg:tw--translate-y-1/2 lg:tw-transform xl:tw-w-[60%] 2xl:tw-w-[49%] xs:tw-left-[13%] xs:tw-w-[75%] xl2:tw-w-[55%] 2xll:tw-w-[45%] xll:tw-w-[61%]"
    >
      <div class="tw-flex tw-w-full">
        <div class="tw-m-auto tw-text-white">
          <p
            class="tw-font-inter-semibold tw-text-[1.375rem] tw-leading-[2.625rem] md:tw-font-inter-bold md:tw-text-[3.4375rem] md:tw-leading-[4.16rem]"
          >
            Have an app idea?
          </p>

          <div
            class="tw-font-inter-regular tw-text-[1rem] tw-leading-[1.5rem] md:tw-my-4 md:tw-font-inter-medium md:tw-text-[1.5rem] md:tw-leading-[1.815625rem]"
          >
            See if we can make it a reality.
          </div>
        </div>
      </div>
      <div class="tw-mx-auto tw-mt-4">
        <router-link
          class="gradient-btn consultation-btn tw-m-0 tw-mx-auto tw-flex tw-w-max tw-items-center tw-rounded-full tw-p-3 tw-px-7 tw-text-center"
          :to="contactURL"
          @click.native="mixpanel.track('tap_android_development_cta')"
        >
          <span
            class="tw-mr-2.5 tw-font-inter-medium tw-text-[1rem] tw-font-normal tw-leading-[1.1875rem] !tw-tracking-[0] md:tw-block md:tw-text-[1.09375rem] md:tw-leading-[1.3125rem] lg:tw-text-[1.1875rem] lg:tw-leading-[1.4375rem]"
            >Let's Chat
          </span>
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
      contactURL: "/contact",
    };
  },
  inject: ["mixpanel"],
};
</script>
