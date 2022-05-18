<template>
  <div class="container-fluid overlay">
    <div class="background">
      <p class="zoom-in-animation dot1"></p>
      <p class="zoom-out-animation dot2"></p>
      <p class="zoom-in-animation dot3"></p>
      <p class="zoom-out-animation dot4"></p>
      <p class="zoom-in-animation dot5"></p>
      <p class="zoom-out-animation dot6"></p>
      <p class="zoom-in-animation dot7"></p>
      <p class="zoom-out-animation dot8"></p>
      <p class="zoom-in-animation dot9"></p>
      <p class="zoom-out-animation dot10"></p>
      <p class="zoom-in-animation dot11"></p>
    </div>

    <div class="container">
      <div class="flex-div">
        <div class="image">
          <img :src="landingViewImage" alt="landing-view-image" />
        </div>
        <div class="flex-detail">
          <div class="description">
            <div class="title-div v2-header-3-text">
              We develop amazing products to help
              <div class="typewriter">
                <span class="wrapper">
                  <div
                    v-for="animationWord in animationWords"
                    :key="animationWord"
                  >
                    <p
                      v-if="animationWord.isActive"
                      class="v2-canopas-gradient-text fw-bold"
                    >
                      {{ animationWord.name }}
                    </p>
                  </div>
                </span>
              </div>
              bring their ideas to life.
            </div>
            <div class="mt-3">
              <a :href="contactURL" class="v2-normal-3-text v2-button">
                <span>Let's talk</span>
                <font-awesome-icon
                  class="arrow fa v2-normal-3-text"
                  icon="arrow-right"
                  id="leftArrow"
                />
              </a>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script type="module">
import landingViewImage from "@/assets/images/Landing/landingImage.webp";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

export default {
  data() {
    return {
      landingViewImage: landingViewImage,
      contactURL: "/contact",
      animationWords: [
        { name: "Entrepreneurs ", isActive: true },
        { name: "Startups ", isActive: false },
        { name: "Businesses ", isActive: false },
      ],
    };
  },
  components: {
    FontAwesomeIcon,
  },
  mounted() {
    var animateText = document.querySelector(".wrapper");

    animateText.addEventListener(
      "animationiteration",
      () => {
        for (var i = 0; i < this.animationWords.length; i++) {
          if (this.animationWords[i].isActive) {
            this.animationWords[i].isActive = false;
            if (i == this.animationWords.length - 1) {
              this.animationWords[0].isActive = true;
            } else {
              this.animationWords[i + 1].isActive = true;
            }
            break;
          }
        }
      },
      false
    );
  },
};
</script>

<style lang="scss" scoped>
.container-fluid {
  position: relative;
  margin-top: 70px;
}

.overlay {
  padding: 30px 0 100px;
}

.flex-div {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.flex-div > .image {
  flex: 1 0 0%;
  text-align: right;
}

.image > img {
  vertical-align: top;
  width: 100%;
  height: 100%;
}

.flex-div > .flex-detail {
  flex: 2 0 0%;
}

.v2-header-3-text {
  font-weight: 400;
}

.description {
  padding-top: 10%;
  width: 90%;
  margin: auto;
}

.title-div {
  padding-bottom: 10%;
}

.landing-image {
  padding-bottom: 5%;
  width: 80%;
  height: 80%;
  object-fit: cover;
}

.typewriter {
  display: flex;
  align-items: center;
  width: fit-content;
}

.typewriter p {
  overflow: hidden;
  border-right: 2px solid #ff9472;
  white-space: nowrap;
  margin-right: 10px;
  margin-top: 10px;
  margin-bottom: 0;
}

.wrapper p {
  animation: typing-erase 4s steps(40, end) infinite;
}

@keyframes typing-erase {
  0% {
    width: 0;
  }
  50%,
  60% {
    width: 100%;
  }
  95%,
  100% {
    width: 0;
  }
}

.v2-button {
  background-color: #3d3d3d;
  color: #fff !important;
  padding: 20px 40px;
}

.v2-button > .fa {
  color: #fff;
}

.v2-button > span {
  margin-right: 10px;
}

@media (hover: hover) and (pointer: fine) {
  .v2-button:hover {
    background-color: #fff;
    color: #3d3d3d;
  }

  .v2-button:hover > span,
  .v2-button:hover > .fa {
    color: #3d3d3d;
  }
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
  width: 17px;
  height: 17px;
  top: 7%;
  left: 20%;
  background: linear-gradient(270.11deg, #ff835b -24.42%, #f2709c 101.76%);
}

.dot2 {
  width: 25px;
  height: 25px;
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
  background: #f2709c;
}

.dot5 {
  width: 15px;
  height: 15px;
  top: 80%;
  left: 43%;
  background: #3d3d3d;
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
  top: 0%;
  right: 50%;
  background: #ff9472;
}

.dot10 {
  width: 20px;
  height: 20px;
  top: 60%;
  right: 9%;
  background: linear-gradient(270.11deg, #ff835b -24.42%, #f2709c 101.76%);
}

.dot11 {
  width: 10px;
  height: 10px;
  bottom: 7%;
  right: 15%;
  background: #ff9472;
}

@include media-breakpoint-up(lg) {
  .flex-div {
    flex-direction: row-reverse;
    align-items: flex-start;
  }

  .description {
    margin: 0;
  }

  .v2-header-3-text {
    font-size: 3.125rem;
    line-height: 4.375rem;
  }

  .overlay {
    padding: 30px 0 250px;
  }
}

@include media-breakpoint-up(xl) {
  .v2-header-3-text {
    font-size: 4.0625rem;
    line-height: 5.875rem;
  }
}
</style>
