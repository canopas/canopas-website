<template>
  <div
    class="sm:tw-container tw-relative tw-pt-[50px] tw-pb-20 md:tw-pt-[100px] md:tw-pb-[150px] lg:tw-pt-[150px] lg:tw-pb-[250px]"
  >
    <div
      class="v2-header-3-text tw-m-auto tw-text-center tw-w-[80%] sm:tw-w-full lg:tw-w-[80%]"
    >
      We say our products are
      <span class="v2-canopas-gradient-text">amazing</span>
      but
      <span class="v2-canopas-gradient-text">user reviews </span>confirm it.
    </div>

    <div v-if="backgroundImage !== null">
      <LottieAnimation
        :jsonData="backgroundImage"
        class="tw-hidden tw md:tw-block md:tw-w-full md:tw-h-[70%] md:tw-absolute md:tw-object-cover md:-tw-z-[5] tw-mx-auto md:tw-opacity-40"
      />
    </div>

    <div class="tw-pt-6">
      <div
        class="tw-flex tw-m-auto tw-flex-wrap md:tw-items-center flex-div md:tw-w-[90%] lg:tw-w-4/5 xl:tw-w-[70%]"
      >
        <div
          v-for="userReview in usersReviews.slice(0, 3)"
          :key="userReview.id"
          class="tw-flex-[1_1_100%] md:tw-flex-[1_1_20%] tw-border-[1.5px] tw-border-solid tw-border-[#fa7b78] tw-rounded-[16px] tw-p-5 tw-mt-[15px] tw-mx-[15px] tw-mb-0 tw-flex tw-flex-col tw-justify-center md:tw-text-center md:tw-m-2.5 md:tw-py-5 md:tw-px-[15px] hover:tw-cursor-pointer hover:tw-shadow-[0_1px_6px_2px_rgba(249,129,136,1)] v2-normal-3-text"
          :class="[userReview.ClassName]"
        >
          {{ userReview.review }}
          <div class="v2-title-3-text tw-mt-4">
            <font-awesome-icon icon="star" class="tw-text-yellow-400" />
            <font-awesome-icon icon="star" class="tw-text-yellow-400" />
            <font-awesome-icon icon="star" class="tw-text-yellow-400" />
            <font-awesome-icon icon="star" class="tw-text-yellow-400" />
            <font-awesome-icon icon="star" class="tw-text-yellow-400" />
          </div>
        </div>
      </div>
      <div
        v-if="slpitReview"
        class="tw-flex tw-m-auto tw-flex-wrap md:tw-items-center flex-div md:tw-w-full lg:tw-w-full xl:tw-w-[90%]"
      >
        <div
          v-for="userReview in usersReviews.slice(3, 5)"
          :key="userReview.id"
          class="tw-flex-[1_1_100%] md:tw-flex-[1_1_20%] tw-border-[1.5px] tw-border-solid tw-border-[#fa7b78] tw-rounded-[16px] tw-p-5 tw-mt-[15px] tw-mx-[15px] tw-mb-0 tw-flex tw-flex-col tw-justify-center md:tw-text-center md:tw-m-2.5 md:tw-py-5 md:tw-px-[15px] hover:tw-cursor-pointer hover:tw-shadow-[0_1px_6px_2px_rgba(249,129,136,1)] v2-normal-3-text"
          :class="[userReview.ClassName]"
        >
          {{ userReview.review }}
          <div class="v2-title-3-text tw-mt-4">
            <font-awesome-icon icon="star" class="tw-text-yellow-400" />
            <font-awesome-icon icon="star" class="tw-text-yellow-400" />
            <font-awesome-icon icon="star" class="tw-text-yellow-400" />
            <font-awesome-icon icon="star" class="tw-text-yellow-400" />
            <font-awesome-icon icon="star" class="tw-text-yellow-400" />
          </div>
        </div>
      </div>
      <div
        v-if="slpitReview"
        class="tw-flex tw-m-auto tw-flex-wrap md:tw-items-center flex-div md:tw-w-[80%] lg:tw-w-[75%] lg:tw-ml-[20%] xl:tw-w-[70%]"
      >
        <div
          v-for="userReview in usersReviews.slice(5, 7)"
          :key="userReview.id"
          class="tw-flex-[1_1_100%] md:tw-flex-[1_1_20%] tw-border-[1.5px] tw-border-solid tw-border-[#fa7b78] tw-rounded-[16px] tw-p-5 tw-mt-[15px] tw-mx-[15px] tw-mb-0 tw-flex tw-flex-col tw-justify-center md:tw-text-center md:tw-m-2.5 md:tw-py-5 md:tw-px-[15px] hover:tw-cursor-pointer hover:tw-shadow-[0_1px_6px_2px_rgba(249,129,136,1)] v2-normal-3-text"
          :class="[userReview.ClassName]"
        >
          {{ userReview.review }}
          <div class="v2-title-3-text tw-mt-4">
            <font-awesome-icon icon="star" class="tw-text-yellow-400" />
            <font-awesome-icon icon="star" class="tw-text-yellow-400" />
            <font-awesome-icon icon="star" class="tw-text-yellow-400" />
            <font-awesome-icon icon="star" class="tw-text-yellow-400" />
            <font-awesome-icon icon="star" class="tw-text-yellow-400" />
          </div>
        </div>
      </div>
      <div
        v-else
        class="tw-flex tw-m-auto tw-flex-wrap md:tw-items-center flex-div md:tw-w-[95%] xl:tw-w-[90%] md:tw-mt-[15px] xl:tw-mt-0"
      >
        <div
          v-for="userReview in usersReviews.slice(3, 7)"
          :key="userReview.id"
          class="tw-flex-[1_1_100%] md:tw-flex-[1_1_20%] tw-border-[1.5px] tw-border-solid tw-border-[#fa7b78] tw-rounded-[16px] tw-p-5 tw-mt-[15px] tw-mx-[15px] tw-mb-0 tw-flex tw-flex-col tw-justify-center md:tw-text-center md:tw-m-2.5 md:tw-py-5 md:tw-px-[15px] hover:tw-cursor-pointer hover:tw-shadow-[0_1px_6px_2px_rgba(249,129,136,1)] v2-normal-3-text"
          :class="[userReview.ClassName]"
        >
          {{ userReview.review }}
          <div class="v2-title-3-text tw-mt-4">
            <font-awesome-icon icon="star" class="tw-text-yellow-400" />
            <font-awesome-icon icon="star" class="tw-text-yellow-400" />
            <font-awesome-icon icon="star" class="tw-text-yellow-400" />
            <font-awesome-icon icon="star" class="tw-text-yellow-400" />
            <font-awesome-icon icon="star" class="tw-text-yellow-400" />
          </div>
        </div>
      </div>
      <div
        class="tw-flex tw-m-auto tw-flex-wrap md:tw-items-center flex-div lg:tw-w-[75%] lg:tw-ml-[20%] md:tw-mt-[15px] xl:tw-mt-0 xl:tw-w-[70%]"
      >
        <div
          v-for="userReview in usersReviews.slice(7, 10)"
          :key="userReview.id"
          class="tw-flex-[1_1_100%] md:tw-flex-[1_1_20%] tw-border-[1.5px] tw-border-solid tw-border-[#fa7b78] tw-rounded-[16px] tw-p-5 tw-mt-[15px] tw-mx-[15px] tw-mb-0 tw-flex tw-flex-col tw-justify-center md:tw-text-center md:tw-m-2.5 md:tw-py-5 md:tw-px-[15px] hover:tw-cursor-pointer hover:tw-shadow-[0_1px_6px_2px_rgba(249,129,136,1)] v2-normal-3-text"
          :class="[userReview.ClassName]"
        >
          {{ userReview.review }}
          <div class="v2-title-3-text tw-mt-4">
            <font-awesome-icon icon="star" class="tw-text-yellow-400" />
            <font-awesome-icon icon="star" class="tw-text-yellow-400" />
            <font-awesome-icon icon="star" class="tw-text-yellow-400" />
            <font-awesome-icon icon="star" class="tw-text-yellow-400" />
            <font-awesome-icon icon="star" class="tw-text-yellow-400" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script type="module">
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { defineAsyncComponent } from "vue";

export default {
  data() {
    return {
      backgroundImage: null,
      usersReviews: [
        {
          id: 1,
          review: "It's smooth, I love, cheers to the team behind it.",
          ClassName: "md:tw-mt-[15%] lg:tw-mt-[15%]",
        },
        {
          id: 2,
          review:
            "The best Moroccan app I have ever seen, clean work.. cheers to developers 🤙",
          ClassName: "",
        },
        {
          id: 3,
          review: "Elegant , splendid well designed app.",
          ClassName: "md:tw-mt-[12%] lg:tw-mt-[12%]",
        },
        {
          id: 4,
          review: "I am surprised by how good this app is !! I love it. ",
          ClassName: "md:tw--mt-[8%] lg:-tw-mt-[4%]",
        },
        {
          id: 5,
          review:
            "The app is well designed, giving a very pleasing and smooth User Experience. Bravo!",
          ClassName: "md:tw-mt-[6%] lg:-tw-mt-0",
        },
        {
          id: 6,
          review: "Perfect design, love how classy they made it",
          ClassName: "lg:tw--mt-[8%] md:tw--mt-[17%]",
        },
        {
          id: 7,
          review: "Slick and elegant.Nice user interface and great content.",
          ClassName: "md:tw-mt-[2%] xl:tw-mt-[5%]",
        },
        {
          id: 8,
          review: "Thank you, developers, for this amazing app experience.",
          ClassName: "md:tw--mt-[17%] lg:tw--mt-[5%]",
        },
        {
          id: 9,
          review: "Impeccable app. Great user experience. ",
          ClassName: "md:tw--mt-[5%] xl:tw--mt-[5%]",
        },
        {
          id: 10,
          review:
            "The app has made incredible changes, I had it since the beginning, always improving !!! ",
          ClassName: "",
        },
      ],
      slpitReview: false,
    };
  },
  components: {
    FontAwesomeIcon,
    LottieAnimation: defineAsyncComponent(() =>
      import("@/components/utils/LottieAnimation.vue")
    ),
  },
  mounted() {
    if (window.innerWidth >= 768 && window.innerWidth <= 992) {
      this.slpitReview = true;
    }
    if (window.innerWidth < 768) {
      this.usersReviews = this.usersReviews.slice(0, 6);
    } else {
      import("@/assets/lottie/globe.json").then((data) => {
        this.backgroundImage = data.default;
      });
    }
  },
};
</script>
