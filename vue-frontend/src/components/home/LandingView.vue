<template>
  <div class="container-fluid overlay">
    <img :src="backgroundImage" class="render-background" />
    <div class="container text-center">
      <div class="description">
        <div class="header-text canopas-gradient-text pb-3">
          Do you have a massive ad budget for your business?
        </div>

        <div class="mt-5 normal-text">
          Unless you’re a Multimillion or a Billion-dollar company, you probably
          don’t have a multimillion-dollar ad budget or professional Spinners.
          <br /><br /><br />
          Your product needs to stand out on its own merits like App Quality,
          Performance, UI design, and User Experience.
        </div>
        <div class="text-center down-arrow">
          <img class="animated-down-arrow" :src="arrowImage" alt="" />
        </div>
        <div class="normal-text second-div">
          Most companies don't care about you, your product, and your vision or
          dreams. They don't give a damn about whether their work helped you to
          get more business, revenue, users, or solving a problem.
          <br /><br /><br />
          That's where CANOPAS comes into the picture.
          <br /><br /><br />
          We help Entrepreneurs, startups, and small companies to bring their
          IDEA to LIFE by developing digital products for their business.
          <br /><br /><br />
          We offer a 100% MONEY BACK GUARANTEE if you don't like our work. No
          questions asked.
        </div>
      </div>
    </div>
  </div>
</template>
 
<script type="module">
import backgroundImage from "@/assets/images/theme/bg.webp";
import arrowImage from "@/assets/images/theme/arrow.svg";

export default {
  data() {
    return {
      backgroundImage: backgroundImage,
      arrowImage: arrowImage,
    };
  },
  components: {},
  method: {},
};
</script>
 
<style lang="scss" scoped>
.container-fluid {
  position: relative;
  padding: 0;
  margin: 0;
}

.render-background {
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
  position: absolute;
  object-fit: cover;
  z-index: -5;
}

.overlay {
  background: rgba(255, 255, 255, 0.95);
  padding: 7% 0;
}

.animated-down-arrow {
  height: 12rem;
  margin: 5% 0;
}

.description {
  margin: auto;
  width: 72%;
}

@include media-breakpoint-up(sm) {
  .description {
    width: 49%;
  }
}

@include media-breakpoint-up(md) {
  .header-text {
    line-height: 2rem;
    font-size: 1.875rem;
  }

  .normal-text {
    font-size: 1.125rem;
    line-height: 1.375rem;
  }
}

@include media-breakpoint-up(lg) {
  .header-text {
    font-size: 3rem;
    line-height: 3.6rem;
  }

  .normal-text {
    font-size: 1.5rem;
    line-height: 1.8rem;
  }
}
</style>
