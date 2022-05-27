<template>
  <div class="container-fluid overlay">
    <div class="background">
      <p class="zoom-out-animation dot1"></p>
      <p class="zoom-in-animation dot2"></p>
      <p class="zoom-out-animation dot3"></p>
      <p class="zoom-out-animation dot4"></p>
      <p class="zoom-in-animation dot5"></p>
      <p class="zoom-in-animation dot6"></p>
      <p class="zoom-out-animation dot7"></p>
      <p class="zoom-in-animation dot8"></p>
      <p class="zoom-out-animation dot9"></p>
      <p class="zoom-in-animation dot10"></p>
    </div>
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
            <a
              class="v2-button v2-normal-3-text"
              :href="contactURL"
              target="_blank"
            >
              <span>Start your free trial</span>
              <font-awesome-icon
                class="arrow fa"
                icon="arrow-right"
                id="leftArrow"
              />
            </a>
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

.background {
  overflow: hidden;
}

.zoom-in-animation {
  border-radius: 50%;
  position: absolute;
  animation: zoom-in 4s ease-in infinite;
  z-index: -1;
}

@keyframes zoom-in {
  0% {
    transform: scale(0, 0);
  }
  50% {
    transform: scale(1, 1);
  }
  100% {
    transform: scale(0, 0);
  }
}

.zoom-out-animation {
  border-radius: 50%;
  position: absolute;
  animation: zoom-out 4s ease-in infinite;
  z-index: -1;
}

@keyframes zoom-out {
  0% {
    transform: scale(1, 1);
  }
  50% {
    transform: scale(0, 0);
  }
  100% {
    transform: scale(1, 1);
  }
}

.dot1 {
  width: 10px;
  height: 10px;
  top: 7%;
  left: 20%;
  background: linear-gradient(270.11deg, #ff835b -24.42%, #f2709c 101.76%);
}

.dot2 {
  width: 20px;
  height: 20px;
  top: 20%;
  left: 5%;
  background: #3d3d3d;
}

.dot3 {
  width: 15px;
  height: 15px;
  top: 45%;
  left: 1%;
  background: #ff9472;
}

.dot4 {
  width: 10px;
  height: 10px;
  top: 38%;
  left: 50%;
  background: #3d3d3d;
}

.dot5 {
  width: 15px;
  height: 15px;
  top: 75%;
  left: 43%;
  background: #ff9472;
}

.dot6 {
  width: 25px;
  height: 25px;
  top: 0;
  right: 18%;
  background: linear-gradient(270.11deg, #ff835b -24.42%, #f2709c 101.76%);
}

.dot7 {
  width: 15px;
  height: 15px;
  top: 8%;
  right: 13%;
  background: #3d3d3d;
}

.dot8 {
  width: 10px;
  height: 10px;
  top: 18%;
  right: 2%;
  background: #ff9472;
}

.dot9 {
  width: 30px;
  height: 30px;
  bottom: 14%;
  left: 9%;
  background: #f2709c;
}

.dot10 {
  width: 20px;
  height: 20px;
  top: 60%;
  right: 9%;
  background: linear-gradient(270.11deg, #ff835b -24.42%, #f2709c 101.76%);
}

.overlay {
  padding: 7% 0;
}

.animated-down-arrow {
  height: 192px;
  margin: 5% auto;
}

.description {
  margin: auto;
  width: 72%;
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
  padding: 1rem 2rem;
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
