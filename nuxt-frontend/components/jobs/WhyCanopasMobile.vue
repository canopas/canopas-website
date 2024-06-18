<template>
  <div class="text-center container m-0 p-0 max-w-full relative">
    <div class="header-text main-text text-[2.5625rem] leading-[2.125rem]">
      <span class="underline-text">Why Ca</span>nopas?
    </div>
    <div></div>
    <div class="reasons-container my-[50px] mx-auto w-full">
      <div
        v-for="(reason, index) in reasons"
        :key="index"
        class="reason flex justify-end relative py-5 px-0 even:flex-row-reverse animate__animated"
        ref="whyCanopasMobile"
      >
        <div
          :class="reason.class"
          class="content my-auto mx-5 flex-[70%] flex flex-col"
        >
          <img
            :src="reason.bullet"
            alt="why-canopas-image"
            loading="lazy"
            class="bullets w-[50px] h-[50px] -mr-5"
          />
          <div class="header-text py-[15px] px-0">{{ reason.title }}</div>
          <div class="normal-text description py-[15px] px-0 sm:w-[70%]">
            {{ reason.description }}
          </div>
        </div>
        <img
          :src="reason.image"
          alt="why-canopas-image"
          loading="lazy"
          class="reason-image w-[126px] h-full"
        />
        <div
          class="absolute bottom-10 h-3 w-3 bg-[#F98386] rounded-full"
          :class="index % 2 == 0 ? 'left-5' : 'right-5'"
        ></div>
        <div
          class="absolute top-10 h-1.5 w-1.5 bg-[#F37299] rounded-full"
          :class="index % 2 == 0 ? 'left-40' : 'right-40'"
        ></div>
      </div>
    </div>
  </div>
</template>

<script>
import bullet1 from "@/assets/images/why/why-canopas-dot-1.svg";
import bullet2 from "@/assets/images/why/why-canopas-dot-2.svg";
import bullet3 from "@/assets/images/why/why-canopas-dot-3.svg";

import circle1 from "@/assets/images/why/why-canopas-mobile-1.webp";
import circle2 from "@/assets/images/why/why-canopas-mobile-2.webp";
import circle3 from "@/assets/images/why/why-canopas-mobile-3.webp";

export default {
  data() {
    return {
      reasons: [
        {
          title: "Our Vision",
          class: "text-right items-end",
          description:
            "To put humanity and the world into a better position than it was prior to our existence.",
          image: circle1,
          bullet: bullet1,
        },
        {
          title: "What's in it for me?",
          class: "text-left items-start",
          description:
            "We have some of the most brilliant, diverse, and humble people you'll ever meet. At Canopas, we have created an environment where everyone asks questions, shares expertise and knowledge, learns something new that helps everyone grow.",
          image: circle2,
          bullet: bullet2,
        },
        {
          title: "How is it relevant to me?",
          class: "text-right items-end",
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
