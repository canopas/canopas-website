<template>
  <div>
    <div class="text-center m-0 p-0 relative">
      <div class="header-text">
        <span class="underline-text">Why </span>Canopas?
      </div>
      <img
        :src="WhyCanopasImage"
        alt="why-canopas-image"
        loading="lazy"
        class="w-full h-full left-0 absolute object-contain -z-[5]"
      />
      <div class="container flex my-[150px] mx-auto m-0 p-0 relative">
        <div
          v-for="reason in reasons"
          :key="reason"
          class="reason flex flex-col flex-[33%] even:mt-[-70px]"
        >
          <img
            :src="reason.image"
            alt="why-canopas-image "
            loading="lazy"
            v-if="reason.image"
            class="w-full h-full object-contain"
          />
          <div
            :class="reason.class"
            class="content flex flex-col even:mt-[5%] xl:even:mt-[10%] first-of-type:w-full last-of-type:w-full animate__animated"
            ref="whyCanopas"
          >
            <img
              :src="reason.bullet"
              alt="why-canopas-image"
              loading="lazy"
              class="bullets w-[50px] h-[50px] -mr-5"
            />
            <div
              class="normal-text even:w-[70%] py-[15px] px-0 w-full text-black-900 font-bold 2xl:w-[65%]"
            >
              {{ reason.title }}
            </div>
            <div
              class="normal-2-text even:w-[70%] description py-[15px] px-0 w-full 2xl:w-[65%]"
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
          class: "items-end text-right pr-[10%]",
          description:
            "To put humanity and the world into a better position than it was prior to our existence.",
          bullet: bullet1,
        },
        {
          title: "How is it relevant to me?",
          class: "items-center",
          description:
            "At Canopas, you will thrive if you want to grow, learn, and stretch yourself. Canopas is built on teamwork and overcoming obstacles, helping each other to achieve shared goals.",
          bullet: bullet3,
          image: circle,
        },
        {
          title: "What's in it for me?",
          class: "items-start text-left pl-[10%]",
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
