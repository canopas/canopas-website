<template>
  <section class="pb-[50px] pt-[100px] md:pb-[150px] md:pt-[200px]">
    <div class="container">
      <div class="hidden md:block">
        <span
          class="v2-canopas-gradient-text text-left font-inter-medium text-[1.5rem] leading-[1.816rem]"
        >
          Testimonials
        </span>
      </div>
      <p
        class="text-center font-inter-bold text-[1.875rem] leading-[2.269rem] text-white md:text-left md:font-inter-semibold md:text-[3.125rem] md:leading-[3.625rem] xl:text-[4.125rem] xl:leading-[4.993rem]"
      >
        What our <br />
        Clients Say
      </p>
      <div class="flex items-center md:flex-row lg:gap-12 xl:items-start">
        <swiper
          :spaceBetween="30"
          :loop="true"
          :slidesPerView="1"
          :autoplay="{
            delay: 5000,
          }"
          :modules="modules"
          class="md:w-[60%] lg:w-[55%]"
        >
          <swiper-slide
            v-for="(review, index) in reviews"
            :key="index"
            class="mt-[2.57rem] cursor-pointer md:mt-16"
          >
            <div class="flex flex-col items-center md:items-start">
              <div
                class="flex flex-col items-center md:flex-row md:items-start md:gap-1 xl:gap-2"
              >
                <img
                  :src="quote"
                  alt="quote"
                  class="h-6 w-6 md:h-8 md:w-8 xl:h-10 xl:w-10"
                />
                <div
                  class="mx-6 text-center font-inter-regular text-[1.25rem] leading-[1.875rem] text-white md:mx-0 md:text-left md:text-[1.5rem] md:leading-[2.188rem] xl:text-[1.75rem] xl:leading-[2.625rem]"
                >
                  {{ review.review }}
                </div>
              </div>

              <span
                class="v2-canopas-gradient-text mt-8 font-inter-bold text-[1.25rem] leading-[1.875rem] md:ml-9 md:text-[1.5rem] md:leading-[2.188rem] xl:ml-[3.25rem] xl:text-[1.75rem] xl:leading-[2.625rem]"
                >━ {{ review.name }}</span
              >
              <span
                class="mt-1 font-inter-regular text-[0.875rem] leading-[1.313rem] text-white md:ml-9 md:text-[0.938rem] md:leading-[1.375rem] xl:ml-[3.25rem] xl:text-[1rem] xl:leading-[1.5rem]"
                >{{ review.designation }}</span
              >
            </div>
          </swiper-slide>
        </swiper>

        <div
          class="relative hidden md:mr-[-30px] md:block md:w-[40%] lg:mr-auto lg:w-[45%]"
        >
          <img
            :src="circle"
            alt="circle"
            class="h-[25rem] w-80 object-contain lg:w-[25rem] xl:w-[35rem] 2xl:h-[28rem]"
          />
          <img
            :src="rating"
            alt="rating"
            class="absolute top-[-38px] h-40 w-40 lg:top-[-110px] lg:h-48 lg:w-48 xl:top-[-140px] xl:h-60 xl:w-60 2xl:top-[-135px]"
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
