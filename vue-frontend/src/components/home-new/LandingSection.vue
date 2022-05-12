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
                  <p class="v2-canopas-gradient-text fw-bold">Entrepreneurs</p>
                </span>
              </div>
              bring their ideas to life.
            </div>
            <div class="mt-5">
              <a :href="contactURL" class="v2-normal-3-text v2-button">
                Let's talk
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
import landingbackgroundImage from "@/assets/images/Landing/landingbackground.webp";
import landingViewImage from "@/assets/images/Landing/landingImage.webp";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

export default {
  data() {
    return {
      landingbackgroundImage: landingbackgroundImage,
      landingViewImage: landingViewImage,
      contactURL: "/contact",
    };
  },
  components: {
    FontAwesomeIcon,
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
}

.typewriter p {
  overflow: hidden;
  border-right: 0.05em solid #ff9472;
  white-space: nowrap;
  margin-right: 10px;
  margin-bottom: 0;
}

.wrapper {
  margin-top: 24px;
  margin-bottom: -17px;
  animation: slide-text 20s steps(1, end) infinite;
}

.wrapper p {
  animation: typing-erase 8s steps(80, end) infinite,
    blink-caret 0.5s step-end infinite;
}

@keyframes typing-erase {
  0% {
    width: 0;
  }
  50%,
  70% {
    width: 100%;
  }
  90%,
  100% {
    width: 0;
  }
}

@keyframes blink-caret {
  from,
  to {
    border-color: transparent;
  }
  50% {
    border-color: #ff9472;
  }
}

@keyframes slide-text {
  0%,
  25% {
    transform: translateY(-1rem);
  }
  25%,
  50% {
    transform: translateY(-1rem);
  }
  50%,
  75% {
    transform: translateY(-1rem);
  }
  75%,
  100% {
    transform: translateY(-1rem);
  }
}

.v2-button {
  background-color: #3d3d3d;
  color: #fff !important;
  box-shadow: none;
  padding: 20px 40px;
}
.arrow {
  color: white;
}

@media (hover: hover) and (pointer: fine) {
  .v2-button:hover {
    background-color: #fff;
    color: #3d3d3d !important;
    border: 1px solid #3d3d3d;
  }
  .v2-button:hover > .arrow {
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
