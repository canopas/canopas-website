<template>
  <section class="mb-40 lg:mb-60">
    <div class="container relative pt-16 lg:pt-60">
      <div class="mobile-header-2 lg:desk-header-2 text-black-87 text-center">
        We say our products are
        <span class="text-primary-1">amazing</span>
        but
        <span class="text-primary-1">user reviews </span>confirm it.
      </div>

      <div v-if="backgroundImage !== null && width > 991">
        <LottieAnimation
          :jsonData="backgroundImage"
          class="w-full h-[70%] absolute -z-[5] mx-auto opacity-30"
        />
      </div>
      <div v-else>
        <img
          :src="background_800w"
          :srcset="`${background_400w} 400w, ${background_800w} 800w,`"
          class="absolute top-[6.75rem] w-full h-full object-contain -z-[1]"
          alt="canopas-cta"
        />
      </div>
      <div class="mt-32 block lg:hidden">
        <swiper
          :slidesPerView="1.5"
          :loop="true"
          :spaceBetween="10"
          :modules="modules"
          :autoplay="{
            delay: 3000,
            disableOnInteraction: false,
          }"
          :centeredSlides="false"
          class="swiper-container"
        >
          <swiper-slide
            v-for="userReview in usersReviews"
            :key="userReview.id"
            class="rounded-[0.813rem] !h-56 flex flex-col items-center justify-center text-center m-0.5 py-6 px-5 hover:cursor-pointer bg-white bg-opacity-80 drop-shadow-md sub-h3-medium text-black-87"
            :class="[userReview.ClassName]"
          >
            {{ userReview.review }}
            <div class="flex justify-center gap-1 mt-4">
              <Icon
                v-for="i in 5"
                :key="i"
                class="fa-star w-6 h-6 text-yellow"
                name="fa6-solid:star"
              />
            </div>
          </swiper-slide>
        </swiper>
      </div>
      <div class="hidden lg:block pt-6">
        <div class="flex m-auto flex-wrap lg:items-center lg:w-4/5 xl:w-[70%]">
          <div
            v-for="userReview in usersReviews.slice(0, 3)"
            :key="userReview.id"
            class="flex-[1_1_100%] lg:flex-[1_1_20%] rounded-[0.813rem] flex flex-col justify-center text-center m-2.5 py-11 px-5 hover:cursor-pointer bg-white bg-opacity-80 drop-shadow-md sub-h1-regular text-black-87"
            :class="[userReview.ClassName]"
          >
            {{ userReview.review }}
            <div class="flex justify-center gap-1 mt-4">
              <Icon
                v-for="i in 5"
                :key="i"
                class="fa-star w-6 h-6 text-yellow"
                name="fa6-solid:star"
              />
            </div>
          </div>
        </div>
        <div
          class="flex m-auto flex-wrap lg:items-center lg:w-[95%] xl:w-[90%] lg:mt-[0.938rem] xl:mt-0"
        >
          <div
            v-for="userReview in usersReviews.slice(3, 7)"
            :key="userReview.id"
            class="flex-[1_1_100%] lg:flex-[1_1_20%] rounded-[0.813rem] flex flex-col justify-center text-center m-2.5 py-11 px-5 hover:cursor-pointer bg-white bg-opacity-80 drop-shadow-md sub-h1-regular text-black-87"
            :class="[userReview.ClassName]"
          >
            {{ userReview.review }}
            <div class="flex justify-center gap-1 mt-4">
              <Icon
                v-for="i in 5"
                :key="i"
                class="fa-star w-6 h-6 text-yellow"
                name="fa6-solid:star"
              />
            </div>
          </div>
        </div>
        <div
          class="flex m-auto flex-wrap lg:items-center lg:w-3/4 lg:ml-[20%] lg:mt-[0.938rem] xl:mt-0 xl:w-[70%]"
        >
          <div
            v-for="userReview in usersReviews.slice(7, 10)"
            :key="userReview.id"
            class="flex-[1_1_100%] lg:flex-[1_1_20%] rounded-[0.813rem] flex flex-col justify-center text-center m-2.5 py-11 px-5 hover:cursor-pointer bg-white bg-opacity-80 drop-shadow-md sub-h1-regular text-black-87"
            :class="[userReview.ClassName]"
          >
            {{ userReview.review }}
            <div class="flex justify-center gap-1 mt-4">
              <Icon
                v-for="i in 5"
                :key="i"
                class="fa-star w-6 h-6 text-yellow"
                name="fa6-solid:star"
              />
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script type="module">
import { defineAsyncComponent } from "vue";
import { Swiper, SwiperSlide } from "swiper/vue";
import { Autoplay } from "swiper/modules";
import background_400w from "@/assets/images/portfolio/background-400w.webp";
import background_800w from "@/assets/images/portfolio/background-800w.webp";

export default {
  data() {
    return {
      width: 0,
      backgroundImage: null,
      background_400w,
      background_800w,
      modules: [Autoplay],
      usersReviews: [
        {
          id: 1,
          review:
            "The best Moroccan app I have ever seen, clean work .. cheers to developers 🤙",
          ClassName: "lg:mt-[150px] xl:mt-[200px]",
        },
        {
          id: 2,
          review:
            "Honestly, 5 stars is an underrating for this app! It's too good. The UI is so much attractive and easy to use. Keep up the Good Work🎯️",
          ClassName: "",
        },
        {
          id: 3,
          review: "Thank you, developers, for this amazing app experience",
          ClassName: " lg:mt-[100px] xl:mt-[150px]",
        },
        {
          id: 4,
          review: "Great UI. Easy and fun to use.",

          ClassName: "",
        },
        {
          id: 5,
          review:
            "The app has made incredible changes, i had it since the beginning, always improving !!! ",
          ClassName: "",
        },
        {
          id: 6,
          review:
            "Kudos to the designers who designed the UI and the developers who implemented it.",
          ClassName: "lg:mt-[-70px] xl:mt-[-50px]",
        },
        {
          id: 7,
          review:
            "The app is well designed, giving a very pleasing and smooth User Experience. Bravo!",
          ClassName: "",
        },
        {
          id: 8,
          review:
            "Good easy to use apps, And can make slideshow look very nice",

          ClassName: "lg:mt-[-70px] xl:mt-[-40px]",
        },
        {
          id: 9,
          review:
            "I’ve always struggled to make something special because it is always difficult. I found this very easy to produce just what I wanted.",
          ClassName: "",
        },
        {
          id: 10,
          review: "Great app, easy to use and amazing design",
          ClassName: "lg:mt-[-170px] xl:mt-[-100px]",
        },
      ],
      splitReview: false,
    };
  },
  components: {
    LottieAnimation: defineAsyncComponent(
      () => import("@/components/utils/LottieAnimation.vue"),
    ),
    Swiper,
    SwiperSlide,
  },
  mounted() {
    import("@/assets/lottie/globe.json").then((data) => {
      this.backgroundImage = data.default;
    });
    this.width = window.innerWidth;
  },
};
</script>
<style lang="postcss">
@import "swiper/css";
</style>
