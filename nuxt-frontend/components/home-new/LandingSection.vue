<template>
  <div
    class="relative bg-gradient-to-t from-pink-300/[0.05] to-white landing-bg-mobile-clip-path md:landing-bg-clip-path overflow-hidden px-0 pb-[3.125rem] md:pb-[6.25rem] font-inter-regular"
  >
    <img
      :src="background2100"
      :srcset="`${background400} 400w, ${background800} 800w, ${background1200} 1200w, ${background2100} 2100w`"
      class="absolute top-0 left-0 w-full h-full -z-[1] object-contain"
      alt="canopas-landing"
      loading="eager"
    />

    <div
      class="container flex flex-col items-center lg:flex-row-reverse lg:items-start pb-[1.563rem] md:pb-[3.125rem]"
    >
      <aspect-ratio class="flex-[45%] lg:block hidden mt-10" height="66%">
        <img
          class="w-full h-full object-contain"
          :src="landing800"
          :srcset="`${landing400} 400w, ${landing800} 800w`"
          alt="landing-view-image"
        />
      </aspect-ratio>
      <div class="flex-[55%] flex flex-col mt-10 md:mt-5 lg:mt-[6.25rem]">
        <h1
          class="w-[100%] sm:w-[85%] pb-5 lg:pb-8 xl:pb-10 tracking-[0] text-[2.1875rem] leading-[2.8125rem] md:text-[2.65625rem] md:leading-[3.28125rem] lg:text-[3.28125rem] lg:leading-[3.70625rem] 2xl:text-[4.0625rem] 2xl:leading-[4.6875rem] text-black-core/[.87] font-roboto-bold v2-header-3-text"
        >
          For Over A Decade We Have Helped Businesses Grow Exponentially.
          <div>
            — <span class="canopas-gradient-text"> Is Yours Next?</span>
          </div>
        </h1>
        <div
          class="w-[85%] md:w-[65%] pb-[29.5px] md:pb-[47px] v2-normal-text tracking-[0] text-[1rem] leading-[1.375rem] md:text-[1.09375rem] md:leading-[1.53125rem] lg:text-[1.1875rem] lg:leading-7 text-black-core/[0.6] font-inter-regular"
        >
          We develop Mobile and Web apps that help you
          <span
            class="bg-no-repeat bg-bottom bg-[length:100%_2.5px] bg-gradient-underline pb-[3px] font-inter-bold"
            >generate more revenue,</span
          >&nbsp;
          <span
            class="bg-no-repeat bg-bottom bg-[length:100%_2.5px] bg-gradient-underline pb-[3px] font-inter-bold"
            >acquire new users,</span
          >
          and
          <span
            class="bg-no-repeat bg-bottom bg-[length:100%_2.5px] bg-gradient-underline pb-[3px] font-inter-bold"
            >deliver an exceptional user experience.</span
          >
        </div>

        <div class="flex items-center flex-col-reverse lg:flex-row">
          <div class="text-center">
            <nuxt-link
              class="flex items-center w-max !m-0 rounded-full p-3 text-center gradient-btn consultation-btn"
              to="/contact"
              @click.native="$mixpanel.track('tap_landing_cta')"
            >
              <span
                class="mr-2.5 font-normal text-[1rem] leading-[1.1875rem] md:text-[1.09375rem] md:leading-[1.3125rem] lg:text-[1.1875rem] lg:leading-[1.4375rem] font-inter-medium !tracking-[0]"
                >Get Free Consultation
              </span> </nuxt-link
            ><small
              class="lg:hidden font-inter-regular text-black-core/[0.6] text-[0.75rem] leading-[1.125rem]"
              >100% MONEY BACK GUARANTEE <br />
              FOR 14 DAYS</small
            >
          </div>

          <div class="flex items-center pb-3.5 md:pl-7 lg:pl-3 xl:pl-7 lg:pb-0">
            <div v-for="client in clients" :key="client">
              <div
                v-if="client.image"
                class="rounded-full"
                :class="client.className"
              >
                <img
                  :src="client.image"
                  class="w-10 md:w-[50px] h-10 md:h-[50px] object-cover"
                  loading="lazy"
                  alt="client-image"
                />
              </div>
              <a v-if="client.plusIcon" href="/contact" aria-label="contact">
                <div
                  class="w-10 md:w-[50px] h-10 md:h-[50px] border-[3px] border-white rounded-full bg-[#ffc4c0]"
                  :class="client.className"
                >
                  <Icon
                    class="fa w-[35px] md:w-[45px] h-[23px] pt-[5px] md:pt-2.5 text-white box-content"
                    name="fa6-solid:plus"
                  />
                </div>
              </a>
            </div>
            <div class="pl-2.5 md:pl-1.5">
              <div class="flex justify-between mt-[3px]">
                <Icon
                  v-for="i in 5"
                  :key="i"
                  class="fas w-3.5 h-3.5 p-px text-[#FF3D2E] box-content"
                  name="fa6-solid:star"
                />
              </div>
              <div class="">
                <span
                  class="font-medium text-[12px] leading-[22px] uppercase text-black-core/[0.6] tracking-[0]"
                >
                  40+ Reviews
                </span>
              </div>
            </div>
          </div>
        </div>
        <small
          class="hidden lg:block w-[74%] xl:w-[56%] pt-3 lg:pt-4 text-[0.875rem] leading-[1.125rem] text-black-core/[0.6]"
        >
          We offer a
          <span class="font-inter-bold"
            >100% MONEY BACK GUARANTEE FOR 14 DAYS
          </span>
          if you don't like our work. No questions asked.
        </small>
      </div>
    </div>
  </div>
</template>
<script>
import landing400 from "@/assets/images/landing/landing-400w.webp";
import landing800 from "@/assets/images/landing/landing-800w.webp";
import client1 from "@/assets/images/landing/client1-100w.webp";
import client2 from "@/assets/images/landing/client2-100w.webp";
import client3 from "@/assets/images/landing/client3-100w.webp";
import client4 from "@/assets/images/landing/client4-100w.webp";
import background400 from "@/assets/images/landing/bg/background-400w.webp";
import background800 from "@/assets/images/landing/bg/background-800w.webp";
import background1200 from "@/assets/images/landing/bg/background-1200w.webp";
import background2100 from "@/assets/images/landing/bg/background-2100w.webp";
import AspectRatio from "../utils/AspectRatio.vue";
export default {
  data() {
    return {
      width: 680,
      landing400,
      landing800,
      background400,
      background800,
      background1200,
      background2100,
      clients: [
        {
          className: "z-0",
          image: client4,
        },
        {
          className: "z-10 -ml-[21px]",
          image: client3,
        },
        {
          className: "z-20 -ml-[21px]",
          image: client2,
        },
        {
          className: "z-30 -ml-[21px]",
          image: client1,
        },
        {
          className: "z-40 -ml-[21px]",
          plusIcon: true,
        },
      ],
    };
  },
  components: {
    AspectRatio,
  },
  inject: ["mixpanel"],
  mounted() {
    this.width = window.innerWidth;
  },
};
</script>
