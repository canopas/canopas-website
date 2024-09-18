<template>
  <div
    class="relative bg-gradient-to-t from-pink-5 to-white landing-bg-mobile-clip-path md:landing-bg-clip-path overflow-hidden px-0 pb-[3.125rem] md:pb-[6.25rem] 3xl:outer-container"
  >
    <img
      :src="background400"
      :srcset="`${background400} 400w, ${background800} 800w, ${background1200} 1200w, ${background2100} 2100w`"
      class="absolute top-0 left-0 w-full h-full -z-[1] object-contain"
      alt="canopas-landing"
      preload
    />
    <div
      class="relative container flex flex-col items-center lg:flex-row-reverse lg:items-start pb-[1.563rem] lg:pb-[3.125rem]"
    >
      <aspect-ratio
        class="flex-[45%] lg:block hidden mt-10 2xl:mt-6"
        height="66%"
      >
        <img
          class="w-full h-full object-contain"
          :src="landing400"
          :srcset="`${landing400} 400w, ${landing800} 800w`"
          alt="landing-view-image"
          preload
        />
      </aspect-ratio>
      <div
        class="flex-[55%] flex flex-col items-center lg:items-start mt-6 lg:mt-[6.25rem]"
      >
        <h1
          class="text-center lg:text-left w-full md:w-[95%] lg:w-[90%] xl:w-[94%] 2xl:w-[81%] mobile-header-1 lg:desk-header-2 xl:desk-header-1"
        >
          For over a decade we have helped businesses grow exponentially. <br />
          -
          <span class="text-primary-1"> Is Yours Next?</span>
        </h1>
        <div
          class="sub-h1-regular lg:mobile-header-2-regular text-black-60 w-full md:w-[75%] lg:w-full 2xl:w-[89%] pt-4 lg:pt-6 text-center lg:text-left"
        >
          We develop Mobile and Web apps that help you
          <span class="sub-h1-semibold lg:mobile-header-2-semibold"
            >generate more revenue, acquire new users,</span
          >
          and
          <span class="sub-h1-semibold lg:mobile-header-2-semibold"
            >deliver an exceptional user experience.</span
          >
        </div>

        <div
          class="flex items-center flex-col-reverse lg:flex-row xl:pt-[1.5rem]"
        >
          <div class="text-center">
            <nuxt-link
              class="gradient-btn primary-btn"
              to="/contact"
              @click.native="$mixpanel.track('tap_landing_cta')"
            >
              <span class="sub-h3-semibold lg:sub-h1-semibold"
                >Get Free Consultation
              </span>
            </nuxt-link>
            <nuxt-link
              to="/contact#guarantee"
              class="lg:hidden sub-h4-regular text-black-60 border-b border-black-60"
              >100% MONEY BACK GUARANTEE FOR <br />
              TWO WEEKS</nuxt-link
            >
          </div>

          <div
            class="flex items-center pt-4 md:pl-7 lg:pl-3 xl:pl-[4.62rem] lg:pb-0"
          >
            <div v-for="client in clients" :key="client">
              <div
                v-if="client.image"
                class="rounded-full border-[2px] border-[#EBEBEB]"
                :class="client.className"
              >
                <img
                  :src="client.image"
                  class="w-10 h-10 object-cover"
                  loading="lazy"
                  alt="client-image"
                />
              </div>
              <div v-if="client.plusIcon">
                <div
                  class="w-10 h-10 border-[2px] border-[#EBEBEB] rounded-full bg-white"
                  :class="client.className"
                >
                  <Icon
                    class="fa w-[18px] lg:w-[20px] h-[18px] lg:h-[20px] pt-[7px] lg:pt-[0.525rem] text-[#999999] box-content pl-2"
                    name="fa6-solid:plus"
                  />
                </div>
              </div>
            </div>
            <div class="pl-2.5 md:pl-1.5">
              <div class="flex mt-[3px]">
                <Icon
                  v-for="i in 5"
                  :key="i"
                  class="fas w-3.5 h-3.5 p-px text-[#FFC700] box-content"
                  name="fa6-solid:star"
                />
              </div>
              <div>
                <span
                  class="sub-h4-regular lg:sub-h4-medium uppercase text-black-60"
                >
                  89+ Reviews
                </span>
              </div>
            </div>
          </div>
        </div>
        <small
          class="hidden lg:block w-[74%] xl:w-[72%] pt-3 lg:pt-4 sub-h4-medium text-black-60"
        >
          We offer a
          <nuxt-link
            to="/contact#guarantee"
            class="secondary-color border-b border-[#F2709C]"
            >100% MONEY BACK GUARANTEE FOR TWO WEEKS
          </nuxt-link>
          if you don't like our work. No questions asked.
        </small>
      </div>
    </div>
  </div>
</template>
<script>
import client1 from "@/assets/images/landing/client1-100w.webp";
import client2 from "@/assets/images/landing/client2-100w.webp";
import client3 from "@/assets/images/landing/client3-100w.webp";
import client4 from "@/assets/images/landing/client4-100w.webp";
import AspectRatio from "../utils/AspectRatio.vue";
export default {
  data() {
    return {
      landing400: "/images/landing/landing-400w.webp",
      landing800: "/images/landing/landing-800w.webp",
      background400: "/images/landing/bg/background-400w.webp",
      background800: "/images/landing/bg/background-800w.webp",
      background1200: "/images/landing/bg/background-1200w.webp",
      background2100: "/images/landing/bg/background-2100w.webp",
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
};
</script>
