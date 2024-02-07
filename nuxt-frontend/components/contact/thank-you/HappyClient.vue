<template>
  <section>
    <div class="container mt-16 lg:my-60">
      <div class="flex gap-y-20 xl:justify-between">
        <div class="flex flex-col justify-center lg:basis-2/5">
          <span
            class="mobile-header-2 lg:desk-header-2 text-center lg:text-left text-black-87"
            >Join our happy Clients!</span
          >
          <span
            class="mt-4 text-center lg:text-left sub-h1-regular lg:mobile-header-2-regular text-black-60"
          >
            The best products start with a foundation of great planning and are
            built through collaboration.
          </span>
        </div>
        <!-- Desktop UI -->
        <div class="hidden lg:block basis-3/5 swiper-content">
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
            class="swiper-container h-[31.25rem] md:h-[43.75rem]"
          >
            <swiper-slide
              v-for="(client, index) in clients"
              :key="index"
              class="cursor-pointer"
            >
              <div
                class="flex flex-col w-fit p-6 sm:p-10 gap-y-10 md:gap-y-12 bg-white shadow-[0px_0px_50px_rgba(0,0,0,0.08)]"
              >
                <div class="flex justify-between">
                  <div class="flex flex-col">
                    <span class="sub-h2-medium text-black-87">{{
                      client.name
                    }}</span>
                    <span class="sub-h1-regular text-black-60">{{
                      client.designation
                    }}</span>
                  </div>
                  <a
                    :href="clutchLink"
                    target="_blank"
                    @click.native="$mixpanel.track('tap_clutch_ratings')"
                  >
                    <div class="flex sm:justify-end w-full sm:w-fit">
                      <div
                        class="flex flex-col sm:justify-center mt-1 sm:mt-[0.438rem]"
                      >
                        <span class="sm:justify-between sub-h4-regular">
                          REVIEWED ON
                        </span>

                        <img
                          :src="reviewImage"
                          class="mr-2 w-[5.0625rem] h-[1.4375rem]"
                          alt="reviewClutchImage"
                        />
                      </div>
                      <div class="flex flex-col mt-2.5 sm:mt-5 sm:justify-end">
                        <span class="flex">
                          <Icon
                            v-for="i in 5"
                            :key="i"
                            class="fa-star w-3 h-3 text-yellow"
                            name="fa6-solid:star"
                        /></span>

                        <span class="sub-h4-regular"> 5.0 RATINGS </span>
                      </div>
                    </div>
                  </a>
                </div>
                <div class="">
                  <div
                    class="mobile-header-2-semibold text-black-87 transition-all ease duration-500"
                    :class="client.class"
                  >
                    {{ client.review }}
                  </div>
                </div>
              </div>
            </swiper-slide>
          </swiper>
        </div>
        <!-- Desktop UI End-->
      </div>
    </div>
    <!-- Mobile UI -->
    <div class="swiper-content -mt-40 lg:hidden">
      <swiper
        :slidesPerView="1.2"
        :effect="'coverflow'"
        :coverflowEffect="{
          rotate: 0,
          stretch: 0,
          depth: 100,
          modifier: 3,
          slideShadows: false,
        }"
        :autoplay="{
          delay: 4000,
          disableOnInteraction: false,
        }"
        :centeredSlides="true"
        :loop="true"
        :modules="modules"
        class="swiper-container"
      >
        <swiper-slide
          v-for="(client, index) in clients"
          :key="index"
          class="cursor-pointer"
        >
          <div
            class="flex flex-col h-[35.25rem] sm:h-[31.25rem] md:h-[29.25rem] w-[32.25rem] w-fit p-6 sm:p-10 gap-y-10 md:gap-y-12 bg-white shadow-[0px_0px_50px_rgba(0,0,0,0.08)]"
          >
            <a
              :href="clutchLink"
              target="_blank"
              @click.native="$mixpanel.track('tap_clutch_ratings')"
            >
              <div class="flex sm:justify-end w-full sm:w-fit">
                <div class="flex flex-col sm:justify-center mt-1 sm:mt-[7px]">
                  <span class="sm:justify-between sub-h4-regular">
                    REVIEWED ON
                  </span>

                  <img
                    :src="reviewImage"
                    class="mr-2 w-[5.0625rem] h-[1.4375rem]"
                    alt="reviewClutchImage"
                  />
                </div>
                <div class="flex flex-col mt-2.5 sm:mt-5 sm:justify-end">
                  <span class="flex">
                    <Icon
                      v-for="i in 5"
                      :key="i"
                      class="fa-star w-3 h-3 text-yellow"
                      name="fa6-solid:star"
                  /></span>

                  <span class="sub-h4-regular"> 5.0 RATINGS </span>
                </div>
              </div>
            </a>

            <div class="mx-auto md:px-6">
              <div
                class="mobile-header-3-semibold !tracking-normal text-black-87 transition-all ease duration-500"
                :class="client.class"
              >
                {{ client.review }}
              </div>
            </div>
            <div class="flex flex-col">
              <span class="sub-h3-medium text-black-87">{{ client.name }}</span>
              <span class="sub-h4-regular text-black-60">{{
                client.designation
              }}</span>
            </div>
          </div>
        </swiper-slide>
      </swiper>
    </div>
    <!-- Mobile UI End -->
  </section>
</template>

<script>
import { Autoplay, Pagination, EffectCoverflow } from "swiper/modules";
import { Swiper, SwiperSlide } from "swiper/vue";
import reviewImage from "@/assets/images/clients/v3-client/ReviewImage.webp";
import Config from "@/config.js";
export default {
  data() {
    return {
      modules: [Pagination, Autoplay, EffectCoverflow],
      reviewImage: reviewImage,
      clutchLink: Config.CLUTCH_URL,
      clients: [
        {
          id: 1,
          name: "Rebecca Kimura",
          designation: "Founder at Togness, Australia",
          review: `“There was rarely ever a second explanation needed. Even if we struggled to explain technically what we wanted, they 
            understood the first time.”`,
          class: "lg:pb-36",
        },
        {
          id: 2,
          name: "Elyass Bouchater",
          designation: "Product Manager at Luxe, Morocco",
          review: `“The Play Store is the hardest app store to get good reviews on, and we’ve just reached a five-star rating, which has been one of our biggest achievements, partly 
           thanks to Canopas’ work.”`,
          class: "lg:pb-36",
        },
        {
          id: 3,
          name: "Rob Eberhard",
          designation: "Founder at ActivScout, Canada",
          review: `“I was especially impressed with the skills of their backend developer and how well the project manager and she worked with one another to create a
            high performing iOS app.”`,
          class: "lg:pb-36",
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
          class: "lg:pb-36",
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
  @apply !items-center mt-0;
}
.swiper-slide {
  @apply relative flex items-center justify-center z-10 transition-transform duration-300;
}
.swiper-slide-active {
  @apply mt-[255px] lg:mt-0 mb-4 lg:mb-0 md:w-[550px]  xl:w-[610px] blur-none z-20 opacity-100;
}
.swiper-slide:not(.swiper-slide-active) {
  @apply lg:w-[494px] opacity-40 shadow-[0px_0px_50px_rgba(0,0,0,0.08)];
}
</style>
