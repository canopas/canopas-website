<template>
  <div>
    <div class="tw-text-center tw-m-0 tw-p-0 tw-relative">
      <div class="header-text">
        <span class="underline-text">Why </span>Canopas?
      </div>
      <img
        :src="WhyCanopasImage"
        alt="why-canopas-image"
        loading="lazy"
        class="tw-w-full tw-h-full tw-left-0 tw-absolute tw-object-contain -tw-z-[5]"
      />
      <div
        class="tw-container tw-flex tw-my-[150px] tw-mx-auto tw-m-0 tw-p-0 tw-relative"
      >
        <div
          v-for="reason in reasons"
          :key="reason"
          class="reason tw-flex tw-flex-col tw-flex-[33%] even:tw-mt-[-70px]"
        >
          <img
            :src="reason.image"
            alt="why-canopas-image"
            loading="lazy"
            v-if="reason.image"
          />
          <div
            :class="reason.class"
            class="content tw-flex tw-flex-col even:tw-mt-[5%] xl:even:tw-mt-[10%] first-of-type:tw-w-full last-of-type:tw-w-full animate__animated"
            ref="whyCanopas"
          >
            <img
              :src="reason.bullet"
              alt="why-canopas-image"
              loading="lazy"
              class="bullets tw-w-[50px] tw-h-[50px] -tw-mr-[20px]"
            />
            <div
              class="normal-text even:tw-w-[70%] tw-py-[15px] tw-px-0 tw-w-full tw-text-black-900 tw-font-bold 2xl:tw-w-[65%]"
            >
              {{ reason.title }}
            </div>
            <div
              class="normal-2-text even:tw-w-[70%] description tw-py-[15px] tw-px-0 tw-w-full 2xl:tw-w-[65%]"
            >
              {{ reason.description }}
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import WhyCanopasImage from "@/assets/images/why/why-canopas-bg.svg";

import bullet1 from "@/assets/images/why/why-canopas-dot-1.svg";
import bullet2 from "@/assets/images/why/why-canopas-dot-2.svg";
import bullet3 from "@/assets/images/why/why-canopas-dot-3.svg";
import circle from "@/assets/images/why/why-canopas-circle.webp";

export default {
  data() {
    return {
      WhyCanopasImage,
      reasons: [
        {
          title: "Our Vision",
          class: "tw-items-end tw-text-right tw-pr-[10%]",
          description:
            "To put humanity and the world into a better position than it was prior to our existence.",
          bullet: bullet1,
        },
        {
          title: "How is it relevant to me?",
          class: "tw-items-center",
          description:
            "At Canopas, you will thrive if you want to grow, learn, and stretch yourself. Canopas is built on teamwork and overcoming obstacles, helping each other to achieve shared goals.",
          bullet: bullet3,
          image: circle,
        },
        {
          title: "What's in it for me?",
          class: "tw-items-start tw-text-left tw-pl-[10%]",
          description:
            "We have some of the most brilliant, diverse, and humble people you'll ever meet. At Canopas, we have created an environment where everyone asks questions, shares expertise and knowledge, learns something new that helps everyone grow.",
          bullet: bullet2,
        },
      ],
    };
  },
  mounted() {
    window.addEventListener("scroll", this.handleScroll);
  },
  unmounted() {
    window.removeEventListener("scroll", this.handleScroll);
  },
  methods: {
    handleScroll() {
      for (var i = 0; i < 3; i++) {
        let data = {
          name: this.$refs.whyCanopas[i],
          animation: "animate__fadeIn",
        };
        if (data) {
          this.$emit("add-animation", data);
        }
      }
    },
  },
};
</script>
