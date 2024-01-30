<template>
  <section class="xll:pb-48">
    <!-- MOBILE UI START -->
    <div class="container flex flex-col pt-10 lg:hidden">
      <div class="text-center">
        <h1 class="mobile-header-1 text-white md:w-3/5 md:mx-auto">
          YOUR PREMIER
          <span class="text-primary-1">MOBILE APP DEVELOPMENT</span>
          COMPANY
        </h1>
        <p class="mt-4 sub-h1-regular text-white-core-80 md:w-[72%] md:mx-auto">
          Canopas - A mobile app development company revolutionizing businesses
          with bespoke Android, iOS, and Flutter applications.
        </p>
        <nuxt-link
          class="mt-8 gradient-btn primary-btn"
          to="/contact"
          @click.native="$mixpanel.track('tap_landing_cta')"
        >
          <span class="sub-h3-semibold lg:sub-h1-semibold"
            >Get Free Consultation
          </span>
        </nuxt-link>
      </div>
      <aspect-ratio height="93%" class="mt-10 md:!pt-[61%]">
        <img
          class="h-full w-full object-contain"
          :src="mobile_mockup_400w"
          :srcset="`${mobile_mockup_400w} 400w, ${mobile_mockup_800w} 800w`"
          alt="landing-view-image"
          loading="lazy"
        />
      </aspect-ratio>
    </div>
    <!-- MOBILE UI END -->

    <!-- DESKTOP UI START -->
    <div class="container hidden lg:flex">
      <div
        class="relative mx-auto flex gap-24 xl:gap-40 2xl:gap-48"
        v-for="index in 6"
        :key="index"
      >
        <p
          class="h-[35rem] border-l border-black-suede lg:h-[45rem] xl:h-[50rem]"
        ></p>
      </div>
      <div class="absolute top-[12.5rem] flex lg:top-[17.5rem]">
        <div class="relative flex basis-[55%] gap-4 xll:max-w-[38%]">
          <p class="gradient-border h-[58%] border-l xll:h-[48%]"></p>
          <div>
            <h1
              class="text-white mobile-header-1 xl:desk-header-1 2xl:w-[93%] xll:w-[94%] 3xl:w-4/5"
            >
              YOUR PREMIER
              <span class="text-primary-1">MOBILE APP DEVELOPMENT</span>
              COMPANY
            </h1>
            <p
              class="mt-6 mobile-header-2-regular text-white-core-80 xl:w-[91%] 2xl:w-[84%] xll:w-[91%] 3xl:w-[76%]"
            >
              Canopas - A mobile app development company revolutionizing
              businesses with bespoke Android, iOS, and Flutter applications.
            </p>
            <nuxt-link
              class="mt-10 ml-0 gradient-btn primary-btn"
              to="/contact"
              @click.native="$mixpanel.track('tap_landing_cta')"
            >
              <span class="sub-h3-semibold lg:sub-h1-semibold"
                >Get Free Consultation
              </span>
            </nuxt-link>
            <div
              class="mt-0 flex justify-end xll:mr-[-9.5rem] 3xl:mr-[-5.5rem]"
            >
              <img
                :src="card_400w"
                :srcset="`${card_400w} 400w, ${card_800w} 800w`"
                alt="landing"
                class="h-2/5 w-2/5 object-contain xl:w-[33%] 2xl:w-[26%] xll:w-[37%]"
              />
            </div>
          </div>
        </div>
        <aspect-ratio
          class="basis-[45%] lg:!pt-[62] xl:!pt-[61%] 2xl:!pt-[51%] xll:max-w-[49%]"
          height="67%"
        >
          <img
            class="absolute top-[-55px] h-full object-contain lg:top-[-95px] xl:top-[-66px] 2xl:top-[-105px] xll:h-[90%] xll:w-full 3xl:top-[-132px] 3xl:h-[77%] 3xl:w-[72%]"
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
