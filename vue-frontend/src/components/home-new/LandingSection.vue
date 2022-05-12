<template>
  <div class="container-fluid overlay">
    <img :src="landingbackgroundImage" class="background" />
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
            <div class="mt-5">
              <a
                :href="contactURL"
                class="v2-normal-3-text v2-button lets-button"
              >
                <span>Let's talk</span>
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
  </div>
</template>
<script type="module">
import landingbackgroundImage from "@/assets/images/Landing/landingbackground.png";
import landingViewImage from "@/assets/images/Landing/landingImage.webp";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

export default {
  data() {
    return {
      landingbackgroundImage: landingbackgroundImage,
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
  padding: 30px 0 250px;
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
  max-width: 100%;
  max-height: 100%;
}
.flex-div > .flex-detail {
  flex: 2 0 0%;
}

.background {
  width: 100%;
  height: 100%;
  top: 0;
  object-fit: cover;
  position: absolute;
  left: 0;
  z-index: -5;
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

.lets-button {
  display: flex;
  align-items: center;
  width: fit-content;
}

.v2-button {
  background-color: #3d3d3d;
  color: #fff !important;
  padding: 15px 40px;
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
}

@include media-breakpoint-up(xl) {
  .v2-header-3-text {
    font-size: 4.0625rem;
    line-height: 5.875rem;
  }
}
</style>
