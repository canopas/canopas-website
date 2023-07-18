<template>
  <section class="md:tw-pb-[50px] lg:tw-pb-[150px] 3xl:tw-pb-[300px]">
    <!-- MOBILE UI START -->
    <div class="tw-container tw-flex tw-flex-col tw-pt-10 md:tw-hidden">
      <div class="tw-flex tw-flex-row tw-gap-4">
        <p class="gradient-border tw-h-auto tw-border-l-[1px]"></p>
        <div>
          <h1
            class="tw-w-[75%] tw-font-inter-bold tw-text-[1.875rem] tw-leading-[2.625rem] tw-text-white sm:tw-w-[80%]"
          >
            YOUR PREMIER
            <span class="v2-canopas-gradient-text">MOBILE APP DEVELOPMENT</span>
            COMPANY
          </h1>
          <p
            class="tw-mt-4 tw-font-inter-medium tw-text-[1.125rem] tw-leading-[1.688rem] tw-text-white/[0.85]"
          >
            Canopas - A mobile app development company revolutionizing
            businesses with bespoke Android, iOS, and Flutter applications.
          </p>
        </div>
      </div>
      <aspect-ratio height="93%" class="tw-mt-10">
        <img
          class="tw-h-full tw-w-full tw-object-contain"
          :src="mobile_mockup_400w"
          :srcset="`${mobile_mockup_400w} 400w, ${mobile_mockup_800w} 800w`"
          alt="landing-view-image"
        />
      </aspect-ratio>
    </div>
    <!-- MOBILE UI END -->

    <!-- DESKTOP UI START -->
    <div class="tw-container tw-hidden md:tw-flex">
      <div
        class="tw-relative tw-mx-auto tw-flex tw-flex-row tw-gap-[6rem] xl:tw-gap-[10rem] 2xl:tw-gap-[12rem]"
        v-for="index in 6"
        :key="index"
      >
        <p
          class="tw-h-[35rem] tw-border-l-[1px] tw-border-[#434343]/[0.4] lg:tw-h-[45rem] xl:tw-h-[50rem]"
        ></p>
      </div>
      <div
        class="tw-absolute tw-top-[12.5rem] tw-flex tw-flex-row lg:tw-top-[17.5rem]"
      >
        <div
          class="tw-relative tw-flex tw-flex-[55%] tw-flex-row tw-gap-4 lg:tw-flex-[55%] xll:tw-max-w-[38%]"
        >
          <p
            class="gradient-border tw-h-[58%] tw-border-l-[1px] xll:tw-h-[48%]"
          ></p>
          <div>
            <h1
              class="tw-w-[98%] tw-font-inter-bold tw-text-[2rem] tw-leading-[3rem] tw-text-white lg:tw-w-[77%] xl:tw-w-[100%] xl:tw-text-[3.125rem] xl:tw-leading-[4.375rem] 2xl:tw-w-[80%] xll:tw-w-[100%] 3xl:tw-w-[72%]"
            >
              YOUR PREMIER
              <span class="v2-canopas-gradient-text"
                >MOBILE APP DEVELOPMENT</span
              >
              COMPANY
            </h1>
            <p
              class="tw-mt-4 tw-w-[98%] tw-font-inter-medium tw-text-[1.225rem] tw-leading-[2rem] tw-text-white/[0.85] lg:tw-w-[76%] xl:tw-w-[100%] xl:tw-text-[2.125rem] xl:tw-leading-[3.188rem] 2xl:tw-w-[85%] xll:tw-w-[91%] 3xl:tw-w-[76%]"
            >
              Canopas - A mobile app development company revolutionizing
              businesses with bespoke Android, iOS, and Flutter applications.
            </p>
            <div class="tw-mt-0 tw-flex tw-justify-end">
              <img
                :src="card_400w"
                :srcset="`${card_400w} 400w, ${card_800w} 800w`"
                alt="landing"
                class="tw-h-[40%] tw-w-[40%] tw-object-contain xl:tw-w-[33%] 2xl:tw-w-[26%] xll:tw-w-[45%]"
              />
            </div>
          </div>
        </div>
        <aspect-ratio
          class="tw-flex-[45%] lg:tw-flex-[45%] lg:!tw-pt-[62] xl:!tw-pt-[61%] 2xl:!tw-pt-[51%] xll:tw-max-w-[49%]"
          height="67%"
        >
          <img
            class="tw-absolute tw-top-[-55px] tw-h-full tw-object-contain lg:tw-top-[-95px] xl:tw-top-[-66px] 2xl:tw-top-[-105px] xll:tw-h-[90%] xll:tw-w-[100%] 3xl:tw-top-[-132px] 3xl:tw-h-[77%] 3xl:tw-w-[72%]"
            :src="mockup_400w"
            :srcset="`${mockup_400w} 400w, ${mockup_800w} 800w`"
            alt="landing-view-image"
            loading="lazy"
          />
        </aspect-ratio>
      </div>
    </div>
    <!-- DESKTOP UI END -->
  </section>
</template>
<script>
import card_400w from "@/assets/images/mobile-app-development/landing/card-400w.webp";
import card_800w from "@/assets/images/mobile-app-development/landing/card-800w.webp";
import mobile_mockup_400w from "@/assets/images/mobile-app-development/landing/mobile-mockup-400w.webp";
import mobile_mockup_800w from "@/assets/images/mobile-app-development/landing/mobile-mockup-800w.webp";
import mockup_400w from "@/assets/images/mobile-app-development/landing/mockup-400w.webp";
import mockup_800w from "@/assets/images/mobile-app-development/landing/mockup-800w.webp";
import AspectRatio from "../utils/AspectRatio.vue";

export default {
  components: { AspectRatio },
  data() {
    return {
      card_400w,
      card_800w,
      mobile_mockup_400w,
      mobile_mockup_800w,
      mockup_400w,
      mockup_800w,
    };
  },
};
</script>

<style scoped>
.gradient-border {
  border-image: linear-gradient(
      91deg,
      #ff835b 0%,
      #f2709c 39.44%,
      rgba(242, 112, 156, 0) 92.29%
    )
    1;
}
</style>
