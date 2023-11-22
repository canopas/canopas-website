<template>
  <div
    class="sm:container relative pt-[3.125rem] pb-20 sm:pb-20 md:pt-[6.25rem] md:pb-[9.375rem] lg:pt-[9.375rem]"
  >
    <div
      class="v2-header-3-text m-auto text-center w-[80%] sm:w-full lg:w-[80%]"
    >
      We say our products are
      <span class="v2-canopas-gradient-text">amazing</span>
      but
      <span class="v2-canopas-gradient-text">user reviews </span>confirm it.
    </div>

    <div v-if="backgroundImage !== null">
      <LottieAnimation
        :jsonData="backgroundImage"
        class="hidden lg:block lg:w-full lg:h-[70%] lg:absolute lg:-z-[5] mx-auto lg:opacity-30"
      />
    </div>

    <div class="pt-6">
      <div
        class="flex m-auto flex-wrap lg:items-center flex-div lg:w-4/5 xl:w-[70%]"
      >
        <div
          v-for="userReview in usersReviews.slice(0, 3)"
          :key="userReview.id"
          class="flex-[1_1_100%] lg:flex-[1_1_20%] border-[1.5px] border-solid border-[#fa7b78] rounded-[16px] p-5 mt-[15px] mx-[15px] mb-0 flex flex-col justify-center lg:text-center lg:m-2.5 lg:py-5 :px-[15px] hover:cursor-pointer hover:shadow-[0_1px_6px_2px_rgba(249,129,136,1)] v2-normal-3-text"
          :class="[userReview.ClassName]"
        >
          {{ userReview.review }}
          <div class="v2-title-3-text mt-4">
            <Icon name="fa6-solid:star" class="text-yellow-400" />
            <Icon name="fa6-solid:star" class="text-yellow-400" />
            <Icon name="fa6-solid:star" class="text-yellow-400" />
            <Icon name="fa6-solid:star" class="text-yellow-400" />
            <Icon name="fa6-solid:star" class="text-yellow-400" />
          </div>
        </div>
      </div>
      <div
        v-if="splitReview"
        class="flex m-auto flex-wrap lg:items-center flex-div md:w-full xl:w-[90%]"
      >
        <div
          v-for="userReview in usersReviews.slice(3, 5)"
          :key="userReview.id"
          class="flex-[1_1_100%] lg:flex-[1_1_20%] border-[1.5px] border-solid border-[#fa7b78] rounded-[16px] p-5 mt-[15px] mx-[15px] mb-0 flex flex-col justify-center lg:text-center lg:m-2.5 lg:py-5 lg:px-[15px] hover:cursor-pointer hover:shadow-[0_1px_6px_2px_rgba(249,129,136,1)] v2-normal-3-text"
          :class="[userReview.ClassName]"
        >
          {{ userReview.review }}
          <div class="v2-title-3-text mt-4">
            <Icon name="fa6-solid:star" class="text-yellow-400" />
            <Icon name="fa6-solid:star" class="text-yellow-400" />
            <Icon name="fa6-solid:star" class="text-yellow-400" />
            <Icon name="fa6-solid:star" class="text-yellow-400" />
            <Icon name="fa6-solid:star" class="text-yellow-400" />
          </div>
        </div>
      </div>
      <div
        v-if="splitReview"
        class="flex m-auto flex-wrap lg:items-center flex-div lg:w-[75%] lg:ml-[20%] xl:w-[70%]"
      >
        <div
          v-for="userReview in usersReviews.slice(5, 7)"
          :key="userReview.id"
          class="flex-[1_1_100%] lg:flex-[1_1_20%] border-[1.5px] border-solid border-[#fa7b78] rounded-[16px] p-5 mt-[15px] mx-[15px] mb-0 flex flex-col justify-center lg:text-center lg:m-2.5 lg:py-5 lg:px-[15px] hover:cursor-pointer hover:shadow-[0_1px_6px_2px_rgba(249,129,136,1)] v2-normal-3-text"
          :class="[userReview.ClassName]"
        >
          {{ userReview.review }}
          <div class="v2-title-3-text mt-4">
            <Icon name="fa6-solid:star" class="text-yellow-400" />
            <Icon name="fa6-solid:star" class="text-yellow-400" />
            <Icon name="fa6-solid:star" class="text-yellow-400" />
            <Icon name="fa6-solid:star" class="text-yellow-400" />
            <Icon name="fa6-solid:star" class="text-yellow-400" />
          </div>
        </div>
      </div>
      <div
        v-else
        class="flex m-auto flex-wrap lg:items-center flex-div lg:w-[95%] xl:w-[90%] lg:mt-[15px] xl:mt-0"
      >
        <div
          v-for="userReview in usersReviews.slice(3, 7)"
          :key="userReview.id"
          class="flex-[1_1_100%] lg:flex-[1_1_20%] border-[1.5px] border-solid border-[#fa7b78] rounded-[16px] p-5 mt-[15px] mx-[15px] mb-0 flex flex-col justify-center lg:text-center lg:m-2.5 lg:py-5 lg:px-[15px] hover:cursor-pointer hover:shadow-[0_1px_6px_2px_rgba(249,129,136,1)] v2-normal-3-text"
          :class="[userReview.ClassName]"
        >
          {{ userReview.review }}
          <div class="v2-title-3-text mt-4">
            <Icon name="fa6-solid:star" class="text-yellow-400" />
            <Icon name="fa6-solid:star" class="text-yellow-400" />
            <Icon name="fa6-solid:star" class="text-yellow-400" />
            <Icon name="fa6-solid:star" class="text-yellow-400" />
            <Icon name="fa6-solid:star" class="text-yellow-400" />
          </div>
        </div>
      </div>
      <div
        class="flex m-auto flex-wrap lg:items-center flex-div lg:w-[75%] lg:ml-[20%] lg:mt-[15px] xl:mt-0 xl:w-[70%]"
      >
        <div
          v-for="userReview in usersReviews.slice(7, 10)"
          :key="userReview.id"
          class="flex-[1_1_100%] lg:flex-[1_1_20%] border-[1.5px] border-solid border-[#fa7b78] rounded-[16px] p-5 mt-[15px] mx-[15px] mb-0 flex flex-col justify-center lg:text-center lg:m-2.5 lg:py-5 lg:px-[15px] hover:cursor-pointer hover:shadow-[0_1px_6px_2px_rgba(249,129,136,1)] v2-normal-3-text"
          :class="[userReview.ClassName]"
        >
          {{ userReview.review }}
          <div class="v2-title-3-text mt-4">
            <Icon name="fa6-solid:star" class="text-yellow-400" />
            <Icon name="fa6-solid:star" class="text-yellow-400" />
            <Icon name="fa6-solid:star" class="text-yellow-400" />
            <Icon name="fa6-solid:star" class="text-yellow-400" />
            <Icon name="fa6-solid:star" class="text-yellow-400" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script type="module">
import { defineAsyncComponent } from "vue";

export default {
  data() {
    return {
      backgroundImage: null,
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
  },
  mounted() {
    if (window.innerWidth >= 768 && window.innerWidth <= 992) {
      this.splitReview = true;
    }
    if (window.innerWidth < 992) {
      this.usersReviews = this.usersReviews.slice(0, 6);
    } else {
      import("@/assets/lottie/globe.json").then((data) => {
        this.backgroundImage = data.default;
      });
    }
  },
};
</script>
