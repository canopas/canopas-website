<template>
  <section class="md:pb-[50px] lg:pb-[150px] 3xl:pb-[300px]">
    <!-- MOBILE UI START -->
    <div class="container flex flex-col pt-10 md:hidden">
      <div class="flex flex-row gap-4">
        <p class="gradient-border h-auto border-l"></p>
        <div>
          <h1
            class="w-[75%] font-inter-bold text-[1.875rem] leading-[2.625rem] text-white sm:w-[80%]"
          >
            YOUR PREMIER
            <span class="v2-canopas-gradient-text">MOBILE APP DEVELOPMENT</span>
            COMPANY
          </h1>
          <p
            class="mt-4 font-inter-medium text-[1.125rem] leading-[1.688rem] text-white/[0.85]"
          >
            Canopas - A mobile app development company revolutionizing
            businesses with bespoke Android, iOS, and Flutter applications.
          </p>
        </div>
      </div>
      <aspect-ratio height="93%" class="mt-10">
        <img
          class="h-full w-full object-contain"
          :src="mobile_mockup_400w"
          :srcset="`${mobile_mockup_400w} 400w, ${mobile_mockup_800w} 800w`"
          alt="landing-view-image"
        />
      </aspect-ratio>
    </div>
    <!-- MOBILE UI END -->

    <!-- DESKTOP UI START -->
    <div class="container hidden md:flex">
      <div
        class="relative mx-auto flex flex-row gap-24 xl:gap-40 2xl:gap-48"
        v-for="index in 6"
        :key="index"
      >
        <p
          class="h-[35rem] border-l border-[#434343]/[0.4] lg:h-[45rem] xl:h-[50rem]"
        ></p>
      </div>
      <div class="absolute top-[12.5rem] flex flex-row lg:top-[17.5rem]">
        <div
          class="relative flex flex-[55%] flex-row gap-4 lg:flex-[55%] xll:max-w-[38%]"
        >
          <p class="gradient-border h-[58%] border-l xll:h-[48%]"></p>
          <div>
            <h1
              class="w-[98%] font-inter-bold text-[2rem] leading-[3rem] text-white lg:w-[77%] xl:w-[100%] xl:text-[3.125rem] xl:leading-[4.375rem] 2xl:w-[80%] xll:w-[100%] 3xl:w-[72%]"
            >
              YOUR PREMIER
              <span class="v2-canopas-gradient-text"
                >MOBILE APP DEVELOPMENT</span
              >
              COMPANY
            </h1>
            <p
              class="mt-4 w-[98%] font-inter-medium text-[1.225rem] leading-[2rem] text-white/[0.85] lg:w-[76%] xl:w-[100%] xl:text-[2.125rem] xl:leading-[3.188rem] 2xl:w-[85%] xll:w-[91%] 3xl:w-[76%]"
            >
              Canopas - A mobile app development company revolutionizing
              businesses with bespoke Android, iOS, and Flutter applications.
            </p>
            <div
              class="mt-0 flex justify-end xll:mr-[-9.5rem] 3xl:mr-[-5.5rem]"
            >
              <img
                :src="card_400w"
                :srcset="`${card_400w} 400w, ${card_800w} 800w`"
                alt="landing"
                class="h-[40%] w-[40%] object-contain xl:w-[33%] 2xl:w-[26%] xll:w-[37%]"
              />
            </div>
          </div>
        </div>
        <aspect-ratio
          class="flex-[45%] lg:flex-[45%] lg:!pt-[62] xl:!pt-[61%] 2xl:!pt-[51%] xll:max-w-[49%]"
          height="67%"
        >
          <img
            class="absolute top-[-55px] h-full object-contain lg:top-[-95px] xl:top-[-66px] 2xl:top-[-105px] xll:h-[90%] xll:w-[100%] 3xl:top-[-132px] 3xl:h-[77%] 3xl:w-[72%]"
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
