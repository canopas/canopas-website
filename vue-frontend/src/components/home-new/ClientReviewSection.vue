<template>
  <div class="tw-container tw-mt-[8rem] tw-ml-auto tw-mx-auto">
    <div class="swiper-content">
      <swiper
        :slidesPerView="1"
        :centeredSlides="true"
        :autoplay="{
          delay: 3000,
          disableOnInteraction: false,
        }"
        :loop="true"
        :loopedSlides="50"
        :spaceBetween="20"
        :pagination="pagination"
        class="swiper-container"
      >
        <swiper-slide v-for="(client, index) in clients" :key="index">
          <div
            class="tw-mx-auto sm:tw-px-10 md:tw-px-6 lg:tw-px-4 xl:tw-px-24 sm:tw-text-black-core/[0.87] tw-text-black-900"
            :class="[client.className]"
          >
            <div
              class="tw-font-roboto-medium tw-text-[1.375rem] md:tw-text-[2.25rem] lg:tw-text-[3.125rem] tw-leading-[1.875rem] md:tw-leading-[2.75rem] lg:tw-leading-[3.75rem]"
              v-html="client.review"
            ></div>
            <div
              class="tw-flex tw-flex-row tw-justify-between tw-mt-11 md:tw-mt-20"
            >
              <div class="tw-flex tw-flex-col">
                <span
                  class="v2-canopas-gradient-text tw-font-inter-semibold tw-text-[1rem] md:tw-text-[1.625rem]"
                  >{{ client.name }}</span
                >
                <span
                  class="tw-text-[0.75rem] md:tw-text-[1.25rem] tw-leading-[1.375rem] tw-text-black-core/[0.87]"
                  >{{ client.designation }}</span
                >
              </div>
              <a
                :href="clutchLink"
                target="_blank"
                @click.native="mixpanel.track('tap_clutch_ratings')"
              >
                <div
                  class="tw-flex tw-flex-row sm:tw-justify-end tw-w-[132px] sm:tw-w-fit"
                >
                  <div
                    class="tw-flex tw-flex-col sm:tw-justify-center tw-mt-[4px] sm:tw-mt-[7px]"
                  >
                    <span
                      class="sm:tw-justify-between tw-mr-1 md:tw-mr-2 tw-text-[0.5rem] md:tw-text-[0.625rem] tw-leading-[1.375rem] tw-text-black-core/[0.87]"
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
          </div>
        </swiper-slide>
      </swiper>
    </div>
  </div>
</template>

<script>
import SwiperCore, { Pagination, Autoplay } from "swiper";
import { Swiper, SwiperSlide } from "swiper/vue";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import reviewImage from "@/assets/images/clients/v3-client/ReviewImage.webp";
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
          className: "2xl:tw-px-48",
          designation: "Founder at Togness, Australia",
          review: ` “There was rarely ever a second explanation needed. Even if we struggled to explain technically what we wanted, they <span class="tw-bg-[length:114%_26px] md:tw-bg-[length:50%_auto] tw-bg-single-color-underline
            ">understood the first time.</span>”`,
        },
        {
          id: 2,
          name: "Elyass Bouchater",
          className: "2xl:tw-px-48",
          designation: "Product Manager at Luxe, Morocco",
          review: `“The Play Store is the hardest app store to get good reviews on, and we’ve just reached a five-star rating, which has been one of our biggest achievements, partly <span class="tw-bg-[length:114%_26px] md:tw-bg-[length:50%_auto] tw-bg-single-color-underline
            ">thanks to Canopas’ work.</span>”`,
        },
        {
          id: 3,
          name: "Rob Eberhard",
          className: "2xl:tw-px-52",
          designation: "Founder at ActivScout, Canada",
          review: `“I was especially impressed with the skills of their backend developer and how well the project manager and she worked with one another to create a <span class="tw-bg-[length:114%_26px] md:tw-bg-[length:50%_auto] tw-bg-single-color-underline
            ">high performing iOS app.</span>”`,
        },
        {
          id: 4,
          name: "Lisa Weinstein",
          className: "2xl:tw-px-52",
          designation: "Founder at Brickandbatten, USA",
          review: `“There is not enough space to say all the wonderful things I would want to share about Canopas. The team is incredibly helpful, stays calm even when we had to deal with tough issues on our app and always found a way to help us fix whatever was needed or roll out any new features for our app in both the
            <span class="tw-bg-[length:114%_26px] md:tw-bg-[length:50%_auto]
               tw-bg-single-color-underline
            "> iOS and Android stores.</span>”`,
        },
      ],

      pagination: {
        el: ".swiper-pagination",
        clickable: true,
      },
    };
  },
  components: {
    Swiper,
    SwiperSlide,
    FontAwesomeIcon,
  },
  inject: ["mixpanel"],
};
</script>

<style lang="postcss">
@import "swiper/css";
@import "swiper/css/pagination";

.swiper-wrapper {
  @apply !tw-items-center;
}
</style>
