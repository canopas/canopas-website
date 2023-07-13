<template>
  <section
    class="tw-pb-[115px] tw-pt-[100px] md:tw-pb-[150px] md:tw-pt-[200px]"
  >
    <div class="tw-container">
      <div class="tw-hidden md:tw-block">
        <span
          class="v2-canopas-gradient-text tw-text-left tw-font-inter-medium tw-text-[1.5rem] tw-leading-[1.816rem]"
        >
          Testimonials
        </span>
      </div>
      <p
        class="tw-text-center tw-font-inter-bold tw-text-[1.875rem] tw-leading-[2.269rem] tw-text-white md:tw-text-left md:tw-font-inter-semibold md:tw-text-[3.125rem] md:tw-leading-[3.625rem] xl:tw-text-[4.125rem] xl:tw-leading-[4.993rem]"
      >
        What our <br />
        Clients Says
      </p>
      <div
        class="tw-mt-8 tw-flex tw-flex-col tw-items-center tw-px-4 sm:tw-px-0 md:tw-hidden"
      >
        <div class="tw-flex tw-flex-col tw-items-center">
          <img :src="quote" alt="quote" class="tw-h-6 tw-w-6" />
          <div
            class="tw-text-center tw-font-inter-regular tw-text-[1.25rem] tw-leading-[1.875rem] tw-text-white"
          >
            {{ filteredReview.review }}
          </div>
        </div>

        <span
          class="v2-canopas-gradient-text tw-mt-8 tw-font-inter-bold tw-text-[1.25rem] tw-leading-[1.875rem]"
          >{{ filteredReview.name }}</span
        >
        <span
          class="tw-mt-1 tw-font-inter-regular tw-text-[0.875rem] tw-leading-[1.313rem] tw-text-white"
          >{{ filteredReview.designation }}</span
        >
      </div>
      <div
        class="tw-hidden tw-items-center md:tw-flex md:tw-flex-row lg:tw-gap-[3rem] xl:tw-items-start"
      >
        <swiper
          :spaceBetween="30"
          :loop="true"
          :slidesPerView="1"
          :autoplay="{
            delay: 3000,
          }"
          class="md:tw-w-[60%] lg:tw-w-[55%]"
        >
          <swiper-slide
            v-for="(review, index) in reviews"
            :key="index"
            class="tw-mt-16 tw-cursor-pointer"
          >
            <div class="tw-flex tw-flex-col tw-items-start">
              <div class="tw-flex tw-flex-row tw-gap-1 xl:tw-gap-2">
                <img
                  :src="quote"
                  alt="quote"
                  class="tw-h-8 tw-w-8 xl:tw-h-10 xl:tw-w-10"
                />
                <div
                  class="tw-text-center tw-font-inter-regular tw-text-white md:tw-text-left md:tw-text-[1.5rem] md:tw-leading-[2.188rem] xl:tw-text-[1.75rem] xl:tw-leading-[2.625rem]"
                >
                  {{ review.review }}
                </div>
              </div>

              <span
                class="v2-canopas-gradient-text tw-ml-9 tw-mt-8 tw-font-inter-bold md:tw-text-[1.5rem] md:tw-leading-[2.188rem] xl:tw-ml-[3.25rem] xl:tw-text-[1.75rem] xl:tw-leading-[2.625rem]"
                >━ {{ review.name }}</span
              >
              <span
                class="tw-ml-9 tw-mt-1 tw-font-inter-regular tw-text-white md:tw-text-[0.938rem] md:tw-leading-[1.375rem] xl:tw-ml-[3.25rem] xl:tw-text-[1rem] xl:tw-leading-[1.5rem]"
                >{{ review.designation }}</span
              >
            </div>
          </swiper-slide>
        </swiper>

        <div
          class="tw-relative tw-hidden md:tw-mr-[-30px] md:tw-block md:tw-w-[40%] lg:tw-mr-auto lg:tw-w-[45%]"
        >
          <img
            :src="circle"
            alt="circle"
            class="tw-h-[25rem] tw-w-[20rem] tw-object-contain lg:tw-w-[25rem] xl:tw-w-[35rem] 2xl:tw-h-[28rem]"
          />
          <img
            :src="rating"
            alt="rating"
            class="tw-absolute tw-top-[-38px] tw-h-[10rem] tw-w-[10rem] lg:tw-top-[-110px] lg:tw-h-[12rem] lg:tw-w-[12rem] xl:tw-top-[-140px] xl:tw-h-[15rem] xl:tw-w-[15rem] 2xl:tw-top-[-135px]"
          />
        </div>
      </div>
    </div>
  </section>
</template>
<script>
import SwiperCore, { Autoplay } from "swiper";
import { Swiper, SwiperSlide } from "swiper/vue";
import quote from "@/assets/images/mobile-app-development/client-review/quote.webp";
import circle from "@/assets/images/mobile-app-development/client-review/circle.gif";
import rating from "@/assets/images/mobile-app-development/client-review/rating-star.gif";

SwiperCore.use([Autoplay]);
export default {
  data() {
    return {
      currentSlide: 0,
      quote,
      circle,
      rating,
      reviews: [
        {
          id: 1,
          name: "Elyass Bouchater",
          designation: "Product Manager at Luxe, Morocco",
          review:
            "The Play Store is the hardest app store to get good reviews on, and we’ve just reached a five-star rating, which has been one of our biggest achievements, partly thanks to Canopas’ work.",
        },
        {
          id: 2,
          name: "Cyril Trosset",
          designation: "CTO at Udini,France",
          review:
            "Multiple versions of this Android app have been successfully delivered over time. They are always very responsive on bug resolution. They are very efficient at producing complex interfaces and high quality apps.",
        },
        {
          id: 3,
          name: "Rob Eberhard",
          designation: "Founder at ActivScout, Canada",
          review:
            "I was especially impressed with the skills of their backend developer and how well the project manager and she worked with one another to create a high performing iOS app.",
        },
        {
          id: 4,
          name: "Lisa Weinstein",
          designation: "Founder at Brickandbatten, USA",
          review:
            "There is not enough space to say all the wonderful things I would want to share about Canopas. The team is incredibly helpful, stays calm even when we had to deal with tough issues on our app and always found a way to help us fix whatever was needed or roll out any new features for our app in both the iOS and Android stores.",
        },
      ],
    };
  },
  components: {
    Swiper,
    SwiperSlide,
  },
  computed: {
    filteredReview() {
      return this.reviews[this.currentSlide];
    },
  },

  methods: {
    startSlideTransition() {
      setInterval(() => {
        this.currentSlide = (this.currentSlide + 1) % this.reviews.length;
      }, 5000);
    },
  },
  mounted() {
    this.startSlideTransition();
  },
};
</script>
<style lang="postcss">
@import "swiper/css";
</style>
