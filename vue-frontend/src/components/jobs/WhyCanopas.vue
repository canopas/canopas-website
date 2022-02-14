<template>
  <div>
    <div class="text-center container-fluid">
      <div class="header-text">
        <span class="underline-text">Why </span>Canopas?
      </div>
      <img
        :src="WhyCanopasImage"
        alt="why-canopas-image"
        loading="lazy"
        class="image"
      />
      <div class="container">
        <div v-for="reason in reasons" :key="reason" class="reason">
          <img
            :src="reason.image"
            alt="why-canopas-image"
            loading="lazy"
            v-if="reason.image"
          />
          <div class="content">
            <img
              :src="reason.bullet"
              alt="why-canopas-image"
              loading="lazy"
              class="bullets"
            />
            <div class="jobs-normal-text">{{ reason.title }}</div>
            <div class="normal-2-text description">
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

export default {
  data() {
    return {
      WhyCanopasImage,
      reasons: [
        {
          title: "Our Vision",
          description:
            "To put humanity and the world into a better position than it was prior to our existence.",
          bullet: require("@/assets/images/why/why-canopas-dot-1.svg"),
        },
        {
          title: "How is it relevant to me?",
          description:
            "At Canopas, you will thrive if you want to grow, learn, and stretch yourself. Canopas is built on teamwork and overcoming obstacles, helping each other to achieve shared goals.",
          bullet: require("@/assets/images/why/why-canopas-dot-3.svg"),
          image: require("@/assets/images/why/why-canopas-circle.svg"),
        },
        {
          title: "What's in it for me?",
          description:
            "We have some of the most brilliant, diverse, and humble people you'll ever meet. At Canopas, we have created an environment where everyone asks questions, shares expertise and knowledge, learns something new that helps everyone grow.",
          bullet: require("@/assets/images/why/why-canopas-dot-2.svg"),
        },
      ],
    };
  },
};
</script>

<style lang="scss" scoped>
.container,
.container-fluid {
  margin: 0;
  padding: 0;
  position: relative;
}

.image {
  width: 100%;
  height: 100%;
  left: 0;
  position: absolute;
  object-fit: contain;
  z-index: -5;
}

.jobs-normal-text {
  color: #3d3d3d;
  font-weight: 700;
}

.container {
  display: flex;
  margin: 150px auto;
}

.reason {
  display: flex;
  flex-direction: column;
  flex: 33%;
}

.content {
  display: flex;
  flex-direction: column;
  align-items: center;

  div {
    padding: 15px 0;
    width: 100%;
  }
}

.reason:nth-of-type(2) {
  margin-top: -70px;

  .content > div {
    width: 70%;
  }
}

.reason:first-of-type > .content,
.reason:last-of-type > .content {
  width: 100%;
}

.reason:first-of-type > .content {
  align-items: flex-end;
  text-align: right;
  padding-right: 10%;
}

.reason:last-of-type > .content {
  align-items: flex-start;
  text-align: left;
  padding-left: 10%;
}

.bullets {
  width: 50px;
  height: 50px;
  margin-right: -20px;
}

@include media-breakpoint-up(xl) {
  .reason:nth-of-type(2) > .content {
    margin-top: 10%;
  }
}

@include media-breakpoint-up(xxl) {
  .content > div {
    width: 65%;
  }
}
</style>