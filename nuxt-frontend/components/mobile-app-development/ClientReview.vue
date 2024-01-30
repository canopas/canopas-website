<template>
  <section class="pb-16 pt-14 lg:pb-28 3xl:pb-[9.188rem] lg:pt-32">
    <div class="container">
      <div class="hidden lg:block">
        <span class="mobile-header-2-semibold text-primary-1">
          Testimonials
        </span>
      </div>
      <p class="text-center desk-header-2 text-white lg:text-left">
        What our <br />
        clients says
      </p>
      <div class="flex items-center lg:gap-12 xl:items-start">
        <swiper
          :spaceBetween="30"
          :loop="true"
          :slidesPerView="1"
          :autoplay="{
            delay: 5000,
          }"
          :modules="modules"
          class="lg:w-[55%]"
        >
          <swiper-slide
            v-for="(review, index) in reviews"
            :key="index"
            class="mt-8 cursor-pointer lg:mt-16"
          >
            <div class="flex flex-col items-center lg:items-start">
              <div
                class="flex flex-col items-center lg:flex-row lg:items-start lg:gap-1 xl:gap-2"
              >
                <img
                  :src="quote"
                  alt="quote"
                  class="h-6 w-6 lg:h-8 lg:w-8 xl:h-10 xl:w-10"
                />
                <div
                  class="mt-2 mx-6 text-center sub-h1-regular lg:mobile-header-2-semibold text-white lg:mx-0 lg:text-left"
                >
                  {{ review.review }}
                </div>
              </div>

              <span
                class="mt-6 text-white sub-h3-regular lg:sub-h2-regular lg:ml-9 xl:ml-[3.25rem]"
                >━ {{ review.name }}</span
              >
              <span
                class="mt-1 text-white-core-80 lg:ml-[3.25rem] sub-h4-regular lg:sub-h1-regular"
                >{{ review.designation }}</span
              >
            </div>
          </swiper-slide>
        </swiper>

        <div
          class="relative hidden lg:mr-[1.875rem] lg:block lg:w-2/5 lg:mr-auto lg:w-[45%]"
        >
          <img
            :src="circle"
            alt="circle"
            class="h-[25rem] w-80 object-contain lg:w-[25rem] xl:w-[35rem] 2xl:h-[28rem]"
          />
          <img
            :src="rating"
            alt="rating"
            class="absolute top-[-2.375rem] h-40 w-40 lg:top-[-6.875rem] lg:h-48 lg:w-48 xl:top-[-8.75rem] xl:h-60 xl:w-60 2xl:top-[-8.438rem]"
          />
        </div>
      </div>
    </div>
  </section>
</template>
<script>
import { Autoplay } from "swiper/modules";
import { Swiper, SwiperSlide } from "swiper/vue";
import quote from "@/assets/images/mobile-app-development/client-review/quote.webp";
import circle from "@/assets/images/mobile-app-development/client-review/circle.gif";
import rating from "@/assets/images/mobile-app-development/client-review/rating-star.gif";

export default {
  data() {
    return {
      quote,
      circle,
      rating,
      modules: [Autoplay],
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
};
</script>
<style lang="postcss">
@import "swiper/css";
</style>
