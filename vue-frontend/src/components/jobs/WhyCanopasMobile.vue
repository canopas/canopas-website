<template>
  <div
    class="tw-text-center tw-container tw-m-0 tw-p-0 tw-max-w-full tw-relative"
  >
    <div
      class="header-text main-text tw-text-[2.5625rem] tw-leading-[2.125rem]"
    >
      <span class="underline-text">Why Ca</span>nopas?
    </div>
    <img
      :src="WhyCanopasImage"
      alt="why-canopas-image"
      loading="lazy"
      class="image tw-w-full tw-h-full tw-top-0 tw-left-0 tw-absolute tw-object-contain sm:tw-object-cover"
    />
    <div class="reasons-container tw-my-[50px] tw-mx-auto tw-w-full">
      <div
        v-for="reason in reasons"
        :key="reason"
        class="reason tw-flex tw-justify-end tw-relative tw-py-[20px] tw-px-0 even:tw-flex-row-reverse animate__animated"
        ref="whyCanopasMobile"
      >
        <div
          :class="reason.class"
          class="content tw-my-auto tw-mx-[20px] tw-flex-[70%] tw-flex tw-flex-col"
        >
          <img
            :src="reason.bullet"
            alt="why-canopas-image"
            loading="lazy"
            class="bullets tw-w-[50px] tw-h-[50px] -tw-mr-[20px]"
          />
          <div class="header-text tw-py-[15px] tw-px-0">{{ reason.title }}</div>
          <div
            class="normal-text description tw-py-[15px] tw-px-0 sm:tw-w-[70%]"
          >
            {{ reason.description }}
          </div>
        </div>
        <img
          :src="reason.image"
          alt="why-canopas-image"
          loading="lazy"
          class="reason-image tw-w-[126px] tw-h-full"
        />
      </div>
    </div>
  </div>
</template>

<script>
import WhyCanopasImage from "@/assets/images/why/why-canopas-mobile-bg.svg";

import bullet1 from "@/assets/images/why/why-canopas-dot-1.svg";
import bullet2 from "@/assets/images/why/why-canopas-dot-2.svg";
import bullet3 from "@/assets/images/why/why-canopas-dot-3.svg";

import circle1 from "@/assets/images/why/why-canopas-mobile-1.svg";
import circle2 from "@/assets/images/why/why-canopas-mobile-2.svg";
import circle3 from "@/assets/images/why/why-canopas-mobile-3.svg";

export default {
  data() {
    return {
      WhyCanopasImage,
      reasons: [
        {
          title: "Our Vision",
          class: "tw-text-right tw-items-end",
          description:
            "To put humanity and the world into a better position than it was prior to our existence.",
          image: circle1,
          bullet: bullet1,
        },
        {
          title: "What's in it for me?",
          class: "tw-text-left tw-items-start",
          description:
            "We have some of the most brilliant, diverse, and humble people you'll ever meet. At Canopas, we have created an environment where everyone asks questions, shares expertise and knowledge, learns something new that helps everyone grow.",
          image: circle2,
          bullet: bullet2,
        },
        {
          title: "How is it relevant to me?",
          class: "tw-text-right tw-items-end",
          description:
            "At Canopas, you will thrive if you want to grow, learn, and stretch yourself. Canopas is built on teamwork and overcoming obstacles, helping each other to achieve shared goals.",
          image: circle3,
          bullet: bullet3,
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
          name: this.$refs.whyCanopasMobile[i],
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
