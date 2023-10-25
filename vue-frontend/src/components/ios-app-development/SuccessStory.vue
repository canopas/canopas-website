<template>
  <section class="tw-mb-20">
    <div
      class="tw-container md:tw-mt-36 lg:tw-mb-8 tw-mb-7 tw-mx-auto tw-text-center"
    >
      <span
        class="tw-tracking-[-0.04rem] tw-mb-2.5 tw-font-inter-bold tw-text-3xl tw-leading-[2.4375rem] md:tw-hidden tw-text-black-core/[0.87] md:!tw-mx-0 md:!tw-w-full md:tw-text-[2.65625rem] lg:tw-text-[3.4375rem] md:tw-leading-[3.796875rem] lg:tw-leading-[5.15625rem] xs:tw-mx-auto xs:tw-w-[71%]"
      >
        Success Stories</span
      >
    </div>
    <!-- Mobile UI start -->
    <div
      class="tw-ml-auto md:tw-hidden tw-mx-auto tw-h-[30rem] tw-relative xll:tw-container tw-overflow-hidden tw-bg-[#1E1E1D]"
    >
      <img
        :src="reviewBackgroundImage"
        :srcset="`${reviewBackgroundImage} 400w`"
        class="tw-h-full tw-w-full tw-object-scale-down tw-ease-in-out tw-duration-1000"
        alt="ios-review-image"
      />
      <div
        class="tw-container swiper-content tw-absolute tw-p-4 tw-w-full tw-text-center tw-text-[#FFF] tw-transform tw-left-1/2 tw-top-1/2 tw--translate-x-1/2 tw--translate-y-1/2 xl:tw-w-[80%] md:tw-top-[49%] lg:tw-top-[48%]"
      >
        <swiper
          :slidesPerView="1"
          :centeredSlides="true"
          :autoplay="{
            delay: 3000,
            disableOnInteraction: false,
          }"
          :speed="3000"
          :loop="true"
          :spaceBetween="20"
          :modules="modules"
          class="swiper-container"
        >
          <swiper-slide
            v-for="(client, index) in clients"
            :key="index"
            class="tw-cursor-pointer tw-my-auto"
          >
            <div
              class="tw-mx-auto sm:tw-px-10 md:tw-px-6 lg:tw-px-4 xl:tw-px-24 tw-text-white tw-flex tw-flex-col tw-justify-between tw-items-center tw-object-center"
            >
              <div
                class="tw-font-inter-semibold tw-text-lg tw-leading-[1.6875rem] tw-transition-all tw-ease tw-duration-500 tw-tracking-[-0.02rem]"
              >
                {{ client.review }}
              </div>
              <div class="tw-flex tw-flex-col tw-mt-8 tw-mb-2.5">
                <span class="tw-flex tw-flex-row">
                  <font-awesome-icon
                    v-for="i in 5"
                    :key="i"
                    class="fa-star tw-w-5 tw-h-5 tw-text-[#FFC700]"
                    icon="star"
                /></span>
              </div>
              <div class="tw-flex tw-flex-col">
                <span
                  class="tw-font-inter-semibold tw-text-base tw-leading-[1.2rem]"
                  >{{ client.name }}</span
                >
                <span
                  class="tw-font-inter-regular tw-text-xs tw-leading-[0.9rem] tw-mt-[0.3rem]"
                  >{{ client.designation }}</span
                >
              </div>
            </div>
          </swiper-slide>
        </swiper>
      </div>
    </div>
    <!-- Mobile UI END -->
    <!-- Desktop UI -->
    <div
      class="tw-hidden md:tw-flex tw-item-center tw-rounded-[20px] tw-w-[95%] lg:tw-w-11/12 xl:tw-w-5/6 2xl:tw-w-[75%] tw-h-[660px] tw-mx-auto tw-drop-shadow-md xll:tw-container"
    >
      <div
        class="tw-w-2/5 tw-flex tw-justify-center tw-items-center tw-bg-[#F4F4F4] tw-rounded-tl-xl tw-rounded-bl-xl"
      >
        <h2
          class="tw-text-[2.4375rem] tw-font-opensans-bold tw-w-3/4 2xl:tw-w-[65%] tw-mx-auto lg:tw-text-5xl tw-leading-[2.79375rem] lg:tw-leading-[3.9rem] !tw-tracking-[-0.04rem] tw-text-black-core/[0.87]"
        >
          What our clients say!
        </h2>
      </div>

      <div
        class="tw-w-3/5 tw-bg-[#3D3D3D] tw-rounded-tr-xl tw-rounded-br-xl tw-flex tw-flex-col tw-justify-center"
      >
        <div class="tw-pl-7 lg:tw-pl-12">
          <img
            :src="imageleft"
            alt="quote-left-image"
            loading="lazy"
            class="tw-w-[1.1rem] tw-h-4 lg:tw-w-8 lg:tw-h-7 tw-text-[#FFFFFF]/[0.60]"
          />
        </div>
        <div class="tw-flex tw-items-center">
          <swiper
            :slidesPerView="1"
            :speed="1000"
            :loop="true"
            :spaceBetween="20"
            :autoplay="{
              delay: 3000,
              disableOnInteraction: false,
            }"
            :pagination="pagination"
            :modules="modules"
            class="swiper-container tw-cursor-pointer"
          >
            <swiper-slide v-for="(client, index) in clients" :key="index">
              <div class="tw-my-auto tw-px-7 lg:tw-px-12 tw-py-7 lg:tw-py-12">
                <span class="tw-flex tw-mb-2.5">
                  <font-awesome-icon
                    v-for="i in 5"
                    :key="i"
                    class="fa-star tw-w-[1.375rem] tw-h-[1.375rem] lg:tw-w-6 lg:tw-h-6 tw-text-[#FFC700]"
                    icon="star"
                /></span>
                <div
                  class="tw-font-inter-semibold tw-text-[#FFFFFF] tw-text-[1.3125rem] lg:tw-text-2xl tw-leading-[2.203125rem] lg:tw-leading-9"
                >
                  {{ client.review }}
                </div>
                <div class="tw-flex tw-flex-col tw-text-left tw-mt-6">
                  <span
                    class="tw-font-inter-medium tw-text-[#FFFFFF]/[0.80] tw-text-text-3xl lg:tw-text-[1.375rem] tw-leading-[1.63125rem] lg:tw-leading-[2.0625rem]"
                    >{{ client.name }}</span
                  >
                  <span
                    class="tw-font-inter-regular tw-text-[#FFFFFF]/[0.60] tw-text-[0.9375rem] lg:tw-text-lg tw-leading-[2.5875rem] lg:tw-leading-[1.6875rem]"
                    >{{ client.designation }}</span
                  >
                </div>
              </div>
            </swiper-slide>
            <div
              class="tw-flex tw-justify-between tw-flex-row-reverse tw-w-[90%]"
            >
              <div class="swiper-pagination"></div>
              <div>
                <img
                  :src="imageright"
                  alt="quote-image-right"
                  loading="lazy"
                  class="tw-w-[1.1rem] tw-h-4 lg:tw-w-8 lg:tw-h-7 tw-text-[#FFFFFF]/[0.60]"
                />
              </div>
            </div>
          </swiper>
        </div>
      </div>
    </div>

    <!-- Desktop UI END-->
  </section>
</template>

<script>
import { Autoplay, Pagination } from "swiper/modules";
import { Swiper, SwiperSlide } from "swiper/vue";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import imageleft from "@/assets/images/ios-app-development/quote-left.webp";
import imageright from "@/assets/images/ios-app-development/quote-right.webp";
import reviewBackgroundImage from "@/assets/images/ios-app-development/review-background-400w.webp";
export default {
  data() {
    return {
      modules: [Pagination, Autoplay],
      reviewBackgroundImage,
      imageleft,
      imageright,
      clients: [
        {
          id: 1,
          name: "Rebecca Kimura",
          designation: "Founder at Togness, Australia",
          review:
            "There was rarely ever a second explanation needed. Even if we struggled to explain technically what we wanted, they understood the first time.",
        },
        {
          id: 2,
          name: "Elyass Bouchater",
          designation: "Product Manager at Luxe, Morocco",
          review:
            "The Play Store is the hardest app store to get good reviews on, and we've just reached a five-star rating, which has been one of our biggest achievements, partly thanks to Canopas' work.",
        },
        {
          id: 3,
          name: "Rob Eberhard",
          designation: "Founder at ActivScout, Canada",
          review:
            "I was especially impressed with the skills of their backend developer and how well the project manager and she worked with one another to create a  high performing iOS app.",
        },
        {
          id: 4,
          name: "Lisa Weinstein",
          designation: "Founder at Brickandbatten, USA",
          review:
            "There is not enough space to say all the wonderful things I would want to share about Canopas. The team is incredibly helpful, stays calm even when we had to deal with tough issues on our app and always found a way to help us fix whatever was needed or roll out any new features for our app in both the iOS and Android stores.",
        },
        {
          id: 5,
          name: "Cyril Trosset",
          designation: "CTO at Udini, France",
          review:
            "Multiple versions of this Android app have been successfully delivered over time. They are always very responsive on bug resolution. They are very efficient at producing complex interfaces and  high quality apps.",
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
  @apply tw-items-center;
}

.swiper-pagination {
  @apply tw-pl-7 lg:tw-pl-12 tw-text-left;
}

.swiper-pagination-bullet {
  @apply !tw-w-[10px] !tw-h-[10px] !tw-rounded-full !tw-bg-transparent !tw-border !tw-border-solid !tw-border-[#FFFFFF]/[0.80];
}

.swiper-pagination-bullet-active {
  @apply !tw-border-none !tw-bg-white;
}
</style>
