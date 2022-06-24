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
    </div>

    <div class="container">
      <div class="flex-div">
        <picture class="image">
          <source
            srcset="data:image/gif;base64,R0lGODlhAQABAAD/ACwAAAAAAQABAAACADs="
            media="(max-width: 992px)"
          />
          <img
            class="landing-image"
            :src="landingImages[2]"
            :srcset="`${landingImages[0]} 400w, ${landingImages[1]} 800w, ${landingImages[2]} 1200w`"
            sizes="33vw"
            alt="landing-view-image"
          />
        </picture>
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
            <div class="flex-button">
              <div class="text-center">
                <router-link
                  class="v2-button v2-normal-3-text"
                  :to="contactURL"
                >
                  <span>Let's talk</span>
                  <font-awesome-icon
                    class="arrow fa"
                    icon="arrow-right"
                    id="leftArrow"
                  />
                </router-link>
                <small class="v2-canopas-gradient-text pt-1">
                  100% MONEY BACK GUARANTEE
                </small>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script type="module">
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import landingImage1x from "@/assets/images/Landing/landing-image-400w.webp";
import landingImage2x from "@/assets/images/Landing/landing-image-800w.webp";
import landingImage3x from "@/assets/images/Landing/landing-image-1200w.webp";

export default {
  data() {
    return {
      landingImages: [landingImage1x, landingImage2x, landingImage3x],
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
  padding-bottom: 20px;
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

.landing-image {
  display: none;
}

.flex-div > .flex-detail {
  flex: 2 0 0%;
}

.v2-header-3-text {
  font-weight: 400;
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
  padding: 1rem 5rem;
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

.background {
  overflow: hidden;
}

.zoom-in-animation {
  border-radius: 50%;
  position: absolute;
  animation: zoom-in 4s ease-in infinite;
  z-index: 0;
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
  z-index: 0;
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
  top: 12%;
  left: 20%;
  background: linear-gradient(270.11deg, #ff835b -24.42%, #f2709c 101.76%);
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
  top: 75%;
  left: 43%;
  background: #3d3d3d;
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
  width: 15px;
  height: 15px;
  top: -7%;
  right: 60%;
  background: #ff9472;
}

.dot10 {
  width: 20px;
  height: 20px;
  top: 60%;
  right: 9%;
  background: linear-gradient(270.11deg, #ff835b -24.42%, #f2709c 101.76%);
}

small {
  font-size: 90%;
}

@include media-breakpoint-up(sm) {
  .description {
    width: 90%;
  }
}

@include media-breakpoint-up(lg) {
  .flex-div {
    flex-direction: row-reverse;
    align-items: flex-start;
  }

  .description {
    padding-top: 10%;
    margin: 0;
  }

  .v2-header-3-text {
    font-size: 3.125rem;
    line-height: 4.375rem;
  }

  .overlay {
    padding-top: 30px;
    padding-bottom: 100px;
  }

  .landing-image {
    display: block;
  }

  .dot2 {
    width: 20px;
    height: 20px;
    top: 20%;
    left: 5%;
    background: #3d3d3d;
  }

  .dot6 {
    width: 25px;
    height: 25px;
    top: 0;
    right: 18%;
    background: linear-gradient(270.11deg, #ff835b -24.42%, #f2709c 101.76%);
  }

  .flex-button {
    align-items: flex-start;
  }
}

@include media-breakpoint-up(xl) {
  .v2-header-3-text {
    font-size: 4.0625rem;
    line-height: 5.875rem;
  }
}

@supports (-webkit-touch-callout: none) {
  .v2-header-3-text {
    letter-spacing: -1px;
  }
}
</style>
