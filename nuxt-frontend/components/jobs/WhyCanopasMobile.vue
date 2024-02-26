<template>
  <div class="mt-16 relative">
    <div class="text-center mobile-header-2 text-black-87">Why Canopas?</div>
    <img
      :src="WhyCanopasImage"
      alt="why-canopas-image"
      loading="lazy"
      class="w-full h-full top-0 left-0 absolute object-contain sm:object-cover"
    />
    <div class="my-12 mx-auto w-full">
      <div
        v-for="reason in reasons"
        :key="reason"
        class="flex justify-end relative py-5 px-0 even:flex-row-reverse animate__animated"
        ref="whyCanopasMobile"
      >
        <div
          :class="reason.class"
          class="my-auto mx-5 basis-[70%] flex flex-col"
        >
          <img
            :src="reason.bullet"
            alt="why-canopas-image"
            loading="lazy"
            class="w-[3.125rem] h-[3.125rem] -mr-5"
          />
          <div class="mobile-header-3-semibold text-black-87">
            {{ reason.title }}
          </div>
          <div class="sub-h3-regular text-black-60 py-4 px-0 sm:w-[70%]">
            {{ reason.description }}
          </div>
        </div>
        <img
          :src="reason.image"
          alt="why-canopas-image"
          loading="lazy"
          class="w-[7.875rem] h-full"
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

import circle1 from "@/assets/images/why/why-canopas-mobile-1.webp";
import circle2 from "@/assets/images/why/why-canopas-mobile-2.webp";
import circle3 from "@/assets/images/why/why-canopas-mobile-3.webp";

export default {
  data() {
    return {
      WhyCanopasImage,
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
