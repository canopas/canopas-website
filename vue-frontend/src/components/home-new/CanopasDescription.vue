<template>
  <div class="container-fluid overlay">
    <div class="container text-center">
      <div class="description" ref="description">
        <div
          ref="firstHeaderText"
          class="header-text canopas-gradient-text pb-3"
        >
          Do you have a massive ad budget for your business?
        </div>

        <div ref="secondHeaderText" class="mt-5 normal-text">
          Unless you’re a Multimillion or a Billion-dollar company, you probably
          don’t have a multimillion-dollar ad budget or professional Spinners.
          <br /><br /><br />
          Your product needs to stand out on its own merits like App Quality,
          Performance, UI design, and User Experience.
        </div>
        <div ref="arrowImage" class="text-center down-arrow">
          <img class="animated-down-arrow" :src="arrowImage" alt="" />
        </div>
        <div ref="thirdHeaderText" class="normal-text second-div">
          Most companies don't care about you, your product, and your vision or
          dreams. They don't give a damn about whether their work helped you to
          get more business, revenue, users, or solving a problem.
          <br /><br /><br />
          That's where CANOPAS comes into the picture.
          <br /><br /><br />
          We help Entrepreneurs, startups, and small companies to bring their
          IDEA to LIFE by developing digital products for their business.
          <br /><br /><br />
          We offer a 100% MONEY BACK GUARANTEE FOR 14 DAYS if you don't like our
          work. No questions asked.
        </div>

        <div class="flex-button pt-5">
          <div class="text-center">
            <router-link class="v2-button v2-normal-3-text" :to="contactURL">
              <span>Start your free trial</span>
              <font-awesome-icon
                class="arrow fa"
                icon="arrow-right"
                id="leftArrow"
              />
            </router-link>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script type="module">
import arrowImage from "@/assets/images/theme/arrow.svg";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

export default {
  data() {
    return {
      arrowImage: arrowImage,
      contactURL: "/contact",
    };
  },
  components: { FontAwesomeIcon },
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
  padding: 60px 0;
}

.animated-down-arrow {
  height: 192px;
  margin: 5% auto;
}

.description {
  margin: auto;
  width: 74%;
}

.flex-button {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.v2-button {
  display: flex;
  align-items: center;
  background-color: #3d3d3d;
  color: #fff !important;
  padding: 1rem 1.1rem;
}

.v2-button > .fa {
  color: #fff;
}

.v2-button > span {
  margin-right: 10px;
}

.arrow {
  height: 18px;
  width: 18px;
}

@media (hover: hover) and (pointer: fine) {
  .v2-button:hover {
    background-color: #fff;
    color: #3d3d3d !important;
  }

  .v2-button:hover > span,
  .v2-button:hover > .fa {
    color: #3d3d3d;
  }
}

@include media-breakpoint-up(sm) {
  .description {
    width: 50%;
  }
}
</style>
