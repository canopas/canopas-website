<template>
  <section>
    <div class="tw-container tw-my-20 md:tw-my-32">
      <div
        class="tw-flex tw-flex-col lg:tw-flex-row tw-space-x-10 tw-gap-y-20 xl:tw-justify-between"
      >
        <div
          class="tw-flex tw-flex-col tw-items-center tw-justify-center tw-basis-[40%]"
        >
          <span
            class="tw-text-[1.875rem] md:tw-text-[3.4375rem] tw-leading-[2.4375rem] md:tw-leading-[5.15625rem] tw-text-center lg:tw-text-left tw-font-inter-bold md:tw-font-inter-semibold tw-text-black-core/[0.87]"
            >Join our happy Clients!</span
          >
          <span
            class="tw-mt-5 tw-text-[1rem] md:tw-text-[1.5rem] tw-leading-[1.5rem] md:tw-leading-[2.25rem] tw-text-center lg:tw-text-left tw-font-inter-medium tw-text-black-core/[0.60]"
          >
            The best products start with a foundation of great planning and are
            built through collaboration.
          </span>
        </div>
        <div class="tw-basis-[60%] swiper-content">
          <swiper
            :slidesPerView="2"
            :centeredSlides="true"
            :spaceBetween="30"
            :direction="'vertical'"
            :autoplay="{
              delay: 5000,
              disableOnInteraction: false,
            }"
            :loop="true"
            :loopedSlides="50"
            class="swiper-container tw-h-[500px] md:tw-h-[700px]"
          >
            <swiper-slide
              v-for="(client, index) in clients"
              :key="index"
              class="tw-cursor-pointer"
            >
              <div
                class="tw-flex tw-flex-col tw-w-fit tw-p-10 tw-gap-y-10 md:tw-gap-y-20 tw-bg-white tw-shadow-[0px_0px_50px_rgba(0,0,0,0.08)]"
              >
                <div class="tw-flex tw-flex-row tw-justify-between">
                  <div class="tw-flex tw-flex-col">
                    <span
                      class="v2-canopas-gradient-text tw-font-inter-semibold tw-text-[1rem] md:tw-text-[1.625rem]"
                      >{{ client.name }}</span
                    >
                    <span
                      class="tw-text-[0.875rem] md:tw-text-[1.25rem] tw-leading-[1.375rem] tw-text-black-core/[0.87]"
                      >{{ client.designation }}</span
                    >
                  </div>
                  <a
                    :href="clutchLink"
                    target="_blank"
                    @click.native="mixpanel.track('tap_clutch_ratings')"
                  >
                    <div
                      class="tw-flex tw-flex-row tw--mt-[0.15rem] sm:tw--mt-3 md:tw--mt-2 lg:tw-mt-[0.1rem] tw-w-[132px] sm:tw-w-fit"
                    >
                      <div
                        class="tw-flex tw-flex-col sm:tw-justify-center tw-mt-[4px] sm:tw-mt-[7px]"
                      >
                        <span
                          class="sm:tw-justify-between tw-mr-1 md:tw-mr-2 tw-text-[0.625rem] md:tw-text-[0.625rem] tw-leading-[1.375rem] tw-text-black-core/[0.87]"
                        >
                          REVIEWED ON
                        </span>

                        <img
                          :src="reviewImage"
                          class="tw-mr-2 tw-w-[3.313rem] md:tw-w-[4.25rem] tw-h-[0.938rem] md:tw-h-[1.188rem]"
                          alt="reviewClutchImage"
                        />
                      </div>
                      <div
                        class="tw-flex tw-flex-col tw-mt-2.5 sm:tw-mt-5 sm:tw-justify-end"
                      >
                        <span class="tw-flex tw-flex-row">
                          <font-awesome-icon
                            v-for="i in 5"
                            :key="i"
                            class="fa-star tw-w-[11px] md:tw-w-[16px] tw-h-[11px] md:tw-h-[16px] tw-text-[#FF3D2E]"
                            icon="star"
                        /></span>

                        <span
                          class="md:tw-font-inter-regular tw-text-[0.5rem] md:tw-text-[0.75rem] tw-leading-[1.375rem] tw-text-black-core/[0.87]"
                        >
                          5.0 RATINGS
                        </span>
                      </div>
                    </div>
                  </a>
                </div>
                <div
                  class="tw-mx-auto sm:tw-px-10 md:tw-px-6 sm:tw-text-black-core/[0.87] tw-text-black-900"
                >
                  <div
                    class="tw-font-roboto-medium tw-text-[1.375rem] md:tw-text-[1.625rem] tw-leading-[1.875rem] md:tw-leading-[2.4375rem] tw-transition-all tw-ease tw-duration-500"
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
import SwiperCore, { Pagination, Autoplay } from "swiper";
import { Swiper, SwiperSlide } from "swiper/vue";
import reviewImage from "@/assets/images/clients/v3-client/ReviewImage.webp";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
SwiperCore.use([Pagination, Autoplay]);
import Config from "@/config.js";
export default {
  data() {
    return {
      reviewImage: reviewImage,
      clutchLink: Config.CLUTCH_URL,
      clients: [
        {
          id: 1,
          name: "Rebecca Kimura",
          designation: "Founder at Togness, Australia",
          review: `“There was rarely ever a second explanation needed. Even if we struggled to explain technically what we wanted, they 
            understood the first time.”`,
          class: "tw-pb-12 md:tw-pb-36",
        },
        {
          id: 2,
          name: "Elyass Bouchater",
          designation: "Product Manager at Luxe, Morocco",
          review: `“The Play Store is the hardest app store to get good reviews on, and we’ve just reached a five-star rating, which has been one of our biggest achievements, partly 
           thanks to Canopas’ work.”`,
          class: "tw-pb-12 md:tw-pb-36",
        },
        {
          id: 3,
          name: "Rob Eberhard",
          designation: "Founder at ActivScout, Canada",
          review: `“I was especially impressed with the skills of their backend developer and how well the project manager and she worked with one another to create a
            high performing iOS app.”`,
          class: "tw-pb-12 md:tw-pb-36",
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
          class: "tw-pb-12 md:tw-pb-36",
        },
      ],
    };
  },

  components: {
    Swiper,
    SwiperSlide,
    FontAwesomeIcon,
  },
};
</script>
<style lang="postcss">
@import "swiper/css";
.swiper-wrapper {
  @apply !tw-items-center;
}
.swiper-slide {
  @apply tw-relative tw-flex tw-items-center tw-justify-center tw-z-10 tw-transition-transform tw-duration-300;
}
.swiper-slide-active {
  @apply tw-scale-[0.9]  md:tw-w-[550px]  xl:tw-w-[610px] tw-blur-none tw-z-20 tw-opacity-100;
}
.swiper-slide:not(.swiper-slide-active) {
  @apply md:tw-w-[494px] tw-opacity-40 tw-shadow-[0px_0px_50px_rgba(0,0,0,0.08)];
}
</style>
