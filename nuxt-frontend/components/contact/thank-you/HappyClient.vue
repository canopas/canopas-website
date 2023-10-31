<template>
  <section>
    <div class="container my-20 md:my-32">
      <div class="flex flex-col lg:flex-row gap-y-20 xl:justify-between">
        <div class="flex flex-col items-center justify-center basis-[40%]">
          <span
            class="text-[1.875rem] md:text-[3.4375rem] leading-[2.4375rem] md:leading-[5.15625rem] text-center lg:text-left font-inter-bold md:font-inter-semibold text-black-core/[0.87]"
            >Join our happy Clients!</span
          >
          <span
            class="mt-5 text-[1rem] md:text-[1.5rem] leading-6 md:leading-9 text-center lg:text-left font-inter-medium text-black-core/[0.60]"
          >
            The best products start with a foundation of great planning and are
            built through collaboration.
          </span>
        </div>
        <div class="basis-[60%] swiper-content">
          <swiper
            :slidesPerView="2"
            :centeredSlides="true"
            :spaceBetween="30"
            :direction="'vertical'"
            :autoplay="{
              delay: 4000,
              disableOnInteraction: false,
            }"
            :loop="true"
            :modules="modules"
            class="swiper-container h-[500px] md:h-[700px]"
          >
            <swiper-slide
              v-for="(client, index) in clients"
              :key="index"
              class="cursor-pointer"
            >
              <div
                class="flex flex-col w-fit p-6 sm:p-10 gap-y-10 md:gap-y-12 bg-white shadow-[0px_0px_50px_rgba(0,0,0,0.08)]"
              >
                <div class="flex flex-row justify-between">
                  <div class="flex flex-col">
                    <span
                      class="v2-canopas-gradient-text font-inter-semibold text-[1rem] md:text-[1.625rem]"
                      >{{ client.name }}</span
                    >
                    <span
                      class="text-[0.875rem] md:text-[1.25rem] leading-[1.375rem] text-black-core/[0.87]"
                      >{{ client.designation }}</span
                    >
                  </div>
                  <a
                    :href="clutchLink"
                    target="_blank"
                    @click.native="$mixpanel.track('tap_clutch_ratings')"
                  >
                    <div
                      class="flex flex-row -mt-[0.15rem] sm:-mt-3 md:-mt-2 lg:mt-[0.1rem] w-[132px] sm:w-fit"
                    >
                      <div
                        class="flex flex-col sm:justify-center mt-1 sm:mt-[7px]"
                      >
                        <span
                          class="sm:justify-between mr-1 md:mr-2 text-[0.625rem] md:text-[0.625rem] leading-[1.375rem] text-black-core/[0.87]"
                        >
                          REVIEWED ON
                        </span>

                        <img
                          :src="reviewImage"
                          class="mr-2 w-[3.313rem] md:w-[4.25rem] h-[0.938rem] md:h-[1.188rem]"
                          alt="reviewClutchImage"
                        />
                      </div>
                      <div class="flex flex-col mt-2.5 sm:mt-5 sm:justify-end">
                        <span class="flex flex-row">
                          <Icon
                            v-for="i in 5"
                            :key="i"
                            class="fa-star w-[11px] md:w-4 h-[11px] md:h-4 text-[#FF3D2E]"
                            name="fa6-solid:star"
                        /></span>

                        <span
                          class="md:font-inter-regular text-[0.5rem] md:text-[0.75rem] leading-[1.375rem] text-black-core/[0.87]"
                        >
                          5.0 RATINGS
                        </span>
                      </div>
                    </div>
                  </a>
                </div>
                <div
                  class="mx-auto md:px-6 sm:text-black-core/[0.87] text-black-900"
                >
                  <div
                    class="font-roboto-medium text-[1.375rem] md:text-[1.625rem] leading-[1.875rem] md:leading-[2.4375rem] transition-all ease duration-500"
                    :class="client.class"
                  >
                    {{ client.review }}
                  </div>
                </div>
              </div>
            </swiper-slide>
          </swiper>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import { Autoplay, Pagination } from "swiper/modules";
import { Swiper, SwiperSlide } from "swiper/vue";
import reviewImage from "@/assets/images/clients/v3-client/ReviewImage.webp";
import Config from "@/config.js";
export default {
  data() {
    return {
      modules: [Pagination, Autoplay],
      reviewImage: reviewImage,
      clutchLink: Config.CLUTCH_URL,
      clients: [
        {
          id: 1,
          name: "Rebecca Kimura",
          designation: "Founder at Togness, Australia",
          review: `“There was rarely ever a second explanation needed. Even if we struggled to explain technically what we wanted, they 
            understood the first time.”`,
          class: "pb-12 md:pb-36",
        },
        {
          id: 2,
          name: "Elyass Bouchater",
          designation: "Product Manager at Luxe, Morocco",
          review: `“The Play Store is the hardest app store to get good reviews on, and we’ve just reached a five-star rating, which has been one of our biggest achievements, partly 
           thanks to Canopas’ work.”`,
          class: "pb-12 md:pb-36",
        },
        {
          id: 3,
          name: "Rob Eberhard",
          designation: "Founder at ActivScout, Canada",
          review: `“I was especially impressed with the skills of their backend developer and how well the project manager and she worked with one another to create a
            high performing iOS app.”`,
          class: "pb-12 md:pb-36",
        },
        {
          id: 4,
          name: "Lisa Weinstein",
          designation: "Founder at Brickandbatten, USA",
          review: `“There is not enough space to say all the wonderful things I would want to share about Canopas. The team is incredibly helpful, stays calm even when we had to deal with tough issues on our app and always found a way to help us fix whatever was needed or roll out any new features for our app in both the
            iOS and Android stores.”`,
        },
        {
          id: 5,
          name: "Cyril Trosset",
          designation: "CTO at Udini, France",
          review: `“Multiple versions of this Android app have been successfully delivered over time. They are always very responsive on bug resolution. They are very efficient at producing complex interfaces and 
            high quality apps.”`,
          class: "pb-12 md:pb-36",
        },
      ],
    };
  },
  inject: ["mixpanel"],
  components: {
    Swiper,
    SwiperSlide,
  },
};
</script>
<style lang="postcss">
@import "swiper/css";
.swiper-wrapper {
  @apply !items-center;
}
.swiper-slide {
  @apply relative flex items-center justify-center z-10 transition-transform duration-300;
}
.swiper-slide-active {
  @apply scale-[0.9]  md:w-[550px]  xl:w-[610px] blur-none z-20 opacity-100;
}
.swiper-slide:not(.swiper-slide-active) {
  @apply md:w-[494px] opacity-40 shadow-[0px_0px_50px_rgba(0,0,0,0.08)];
}
</style>
