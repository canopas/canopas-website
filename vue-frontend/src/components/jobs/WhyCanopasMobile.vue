<template>
  <div class="text-center container">
    <div class="header-text main-text">
      <span class="underline-text">Why Ca</span>nopas?
    </div>
    <img
      :src="WhyCanopasImage"
      alt="why-canopas-image"
      loading="lazy"
      class="image"
    />
    <div class="reasons-container">
      <div
        v-for="reason in reasons"
        :key="reason"
        class="reason animate__animated"
        ref="whyCanopasMobile"
      >
        <div class="content">
          <img
            :src="reason.bullet"
            alt="why-canopas-image"
            loading="lazy"
            class="bullets"
          />
          <div class="header-text">{{ reason.title }}</div>
          <div class="normal-text description">
            {{ reason.description }}
          </div>
        </div>
        <img
          :src="reason.image"
          alt="why-canopas-image"
          loading="lazy"
          class="reason-image tw-w-auto tw-h-auto"
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
          description:
            "To put humanity and the world into a better position than it was prior to our existence.",
          image: circle1,
          bullet: bullet1,
        },
        {
          title: "What's in it for me?",
          description:
            "We have some of the most brilliant, diverse, and humble people you'll ever meet. At Canopas, we have created an environment where everyone asks questions, shares expertise and knowledge, learns something new that helps everyone grow.",
          image: circle2,
          bullet: bullet2,
        },
        {
          title: "How is it relevant to me?",
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
          childRef: [],
        };
        if (data) {
          this.$emit("add-animation", data);
        }
      }
    },
  },
};
</script>

<style lang="scss" scoped>
.container {
  margin: 0;
  padding: 0;
  max-width: 100%;
  position: relative;
}

.image {
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
  position: absolute;
  object-fit: contain;
}

.reasons-container {
  margin: 50px auto;
  width: 100%;
}

.reason {
  display: flex;
  justify-content: flex-end;
  position: relative;
  padding: 20px 0;
}

.reason:nth-of-type(even) {
  flex-direction: row-reverse;
  .content {
    text-align: left;
    align-items: start;
  }
}

.main-text {
  line-height: 2.125rem;
  font-size: 2.5625rem;
}

.content {
  text-align: right;
  margin: auto 20px;
  flex: 70%;
  display: flex;
  flex-direction: column;
  align-items: flex-end;

  div {
    padding: 15px 0;
  }
}

.bullets {
  width: 50px;
  height: 50px;
  margin-right: -20px;
}

@include media-breakpoint-up(sm) {
  .description {
    width: 70%;
  }

  .image {
    object-fit: cover;
  }
}
</style>
