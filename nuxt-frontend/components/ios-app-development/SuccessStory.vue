<template>
  <section class="my-20 lg:my-60 3xl:mt-[315px]">
    <div class="container md:mt-36 lg:mb-8 mb-7 mx-auto text-center">
      <span
        class="mb-2.5 mobile-header-2 md:hidden text-black-87 md:!mx-0 md:!w-full xs:mx-auto xs:w-[71%]"
      >
        Success Stories</span
      >
    </div>
    <!-- Mobile UI start -->
    <div
      class="ml-auto lg:hidden mx-auto h-[32.188rem] relative xll:container overflow-hidden bg-[#1E1E1D]"
    >
      <img
        :src="reviewBackgroundImage"
        :srcset="`${reviewBackgroundImage} 400w`"
        class="h-full w-full object-scale-down ease-in-out duration-1000"
        alt="ios-review-image"
      />
      <div
        class="container swiper-content absolute p-4 w-full text-center text-[#FFF] transform left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2 xl:w-[80%] md:top-[49%] lg:top-[48%]"
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
            class="cursor-pointer my-auto"
          >
            <div
              class="mx-auto sm:px-10 md:px-6 lg:px-4 xl:px-24 text-white flex flex-col justify-between items-center object-center"
            >
              <div class="transition-all ease duration-500 sub-h1-semibold">
                {{ client.review }}
              </div>
              <div class="flex flex-col mt-8 mb-2.5">
                <span class="flex flex-row">
                  <Icon
                    v-for="i in 5"
                    :key="i"
                    class="fa-star w-5 h-5 text-[#FFC700]"
                    name="fa6-solid:star"
                /></span>
              </div>
              <div class="flex flex-col">
                <span class="sub-h3-medium text-white-core-80">{{
                  client.name
                }}</span>
                <span class="sub-h4-regular text-white-core-60 mt-[0.3rem]">{{
                  client.designation
                }}</span>
              </div>
            </div>
          </swiper-slide>
        </swiper>
      </div>
    </div>
    <!-- Mobile UI END -->
    <!-- Desktop UI -->
    <div
      class="hidden lg:flex item-center rounded-[20px] w-[95%] lg:w-11/12 xl:w-5/6 2xl:w-[75%] h-[660px] mx-auto drop-shadow-md xll:container"
    >
      <div
        class="w-2/5 flex justify-center items-center bg-[#F4F4F4] rounded-tl-xl rounded-bl-xl"
      >
        <h2 class="desk-header-2 w-3/4 2xl:w-[65%] mx-auto text-black-87">
          What our clients says!
        </h2>
      </div>

      <div
        class="w-3/5 bg-[#121212] rounded-tr-xl rounded-br-xl flex flex-col justify-center"
      >
        <div class="pl-7 lg:pl-12">
          <img
            :src="imageleft"
            alt="quote-left-image"
            loading="lazy"
            class="w-[1.1rem] h-4 lg:w-8 lg:h-7 text-white-core-60"
          />
        </div>
        <div class="flex items-center">
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
            class="swiper-container cursor-pointer"
          >
            <swiper-slide v-for="(client, index) in clients" :key="index">
              <div class="my-auto px-7 lg:px-12 py-7 lg:py-12 text-white">
                <span class="flex mb-2.5">
                  <Icon
                    v-for="i in 5"
                    :key="i"
                    class="fa-star w-[1.375rem] h-[1.375rem] lg:w-6 lg:h-6 text-[#FFC700]"
                    name="fa6-solid:star"
                /></span>
                <div class="mobile-header-2-semibold">
                  {{ client.review }}
                </div>
                <div class="flex flex-col text-left mt-6">
                  <span class="text-white-core-80 sub-h2-medium">{{
                    client.name
                  }}</span>
                  <span class="text-white-core-60 sub-h1-regular">{{
                    client.designation
                  }}</span>
                </div>
              </div>
            </swiper-slide>
            <div class="flex justify-between flex-row-reverse w-[90%]">
              <div class="swiper-pagination"></div>
              <div>
                <img
                  :src="imageright"
                  alt="quote-image-right"
                  loading="lazy"
                  class="w-[1.1rem] h-4 lg:w-8 lg:h-7 text-white-core-60"
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
  },
};
</script>
<style lang="postcss" scoped>
@import "swiper/css";
@import "swiper/css/pagination";
.swiper-wrapper {
  @apply items-center;
}

.swiper-pagination {
  @apply pl-7 lg:pl-12 text-left;
}

.swiper-pagination-bullet {
  @apply !w-[10px] !h-[10px] !rounded-full !bg-transparent !border !border-solid !border-[#FFFFFF]/[0.80];
}

.swiper-pagination-bullet-active {
  @apply !border-none !bg-white;
}
</style>
