<template>
  <div class="background mt-5 mb-5">
    <img :src="backgroundImage" class="render-background" />
    <div class="container cta-section text-center">
      <div class="v2-header-3-text canopas-gradient-text">
        How do you know we’re a good fit for you?
      </div>
      <div class="v2-title-3-text mt-4">Let’s have a short chat.</div>
      <div class="communication">
        <div class="item px-4 py-3">
          <div class="v2-normal-3-text">I love talking.</div>
          <a class="gradient-btn cta-button" :href="contactURL">
            <font-awesome-icon
              class="fa v2-title-3-text"
              icon="phone"
              aria-hidden="true"
            />
            <span class="v2-normal-3-text">Schedule a call</span>
          </a>
        </div>
        <div class="item px-4 py-3">
          <div class="v2-normal-3-text">I prefer mail.</div>
          <a class="gradient-btn cta-button" :href="contactURL">
            <font-awesome-icon
              class="fa v2-title-3-text"
              icon="paper-plane"
              aria-hidden="true"
            />
            <span class="v2-normal-3-text">Drop a mail</span>
          </a>
        </div>
      </div>
      <div
        class="v2-title-3-text text-left v2-canopas-gradient-text reason-text mt-5"
      >
        Four reasons why do we do this call?
      </div>
      <ul class="v2-normal-2-text chat-list text-left mt-4">
        <li>
          <span class="bullet"></span>
          <div>It gives an opportunity to learn more about each other.</div>
        </li>
        <li>
          <span class="bullet"></span>
          <div>
            We want to give you a chance to put your idea into action. You might
            have faced this situation before where you wish you had someone to
            help you.
          </div>
        </li>
        <li>
          <span class="bullet"></span>
          <div>
            Sometimes, you want to check the feasibility of your idea. We will
            give you honest feedback based on our industry experience.
          </div>
        </li>
        <li>
          <span class="bullet"></span>
          <div>
            We offer 100% Money-back guarantee. so, there’s no risk to you.
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>

<script type="module">
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import backgroundImage from "@/assets/images/theme/cta-background.webp";

export default {
  data() {
    return {
      backgroundImage: backgroundImage,
      contactURL: "/contact",
    };
  },
  components: {
    FontAwesomeIcon,
  },
};
</script>

<style lang="scss" scoped>
.background {
  background: #3d3d3d;
  padding: 100px 0;
  overflow: hidden;
  border-radius: 100px 0;
  position: relative;
  z-index: 0;
}

.render-background {
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
  position: absolute;
  object-fit: cover;
  z-index: -1;
}

.container {
  padding: 0 25px;
}

.v2-header-3-text {
  margin: auto;
}

.v2-title-3-text,
.v2-normal-3-text,
.v2-normal-2-text {
  color: #ffffff;
}

.communication {
  display: flex;
  flex-flow: wrap;
  justify-content: center;
  margin: auto;
}

.item {
  margin: 20px;
  border: 1px solid #e2e2e2;
  border-radius: 15px;
  text-align: -webkit-center;
}

.cta-button {
  display: flex;
  align-items: center;
}

.v2-title-3-text {
  font-size: 1.35rem;
  line-height: 1.5rem;
}

.gradient-btn {
  border-radius: 2rem;
}

.gradient-btn > span {
  margin: 0 0 0 10px;
  font-weight: 400;
}

@media (hover: hover) and (pointer: fine) {
  .gradient-btn:hover {
    box-shadow: 2px 1000px 1px #3d3d3d inset;
  }
}

.text-left {
  text-align: left;
}

.reason-text {
  margin: auto;
}

.bullet {
  content: "";
  background: linear-gradient(180deg, #f2709c 0%, #ff9472 100%);
  width: 16px;
  height: 16px;
  border-radius: 30%;
  display: inline-block;
  line-height: 1;
  margin: 5px 24px 0 0;
}

.chat-list {
  list-style-type: none;
  margin: auto;
  padding: 0;
}

.chat-list > li {
  padding: 10px 0;
  display: flex;
}

.chat-list > li > div {
  width: 90%;
}

@include media-breakpoint-up(sm) {
  .background {
    border-radius: 150px 0;
  }

  .cta-section {
    width: 80%;
  }
}

@include media-breakpoint-up(md) {
  .background {
    border-radius: 200px 0;
  }
}

@include media-breakpoint-up(lg) {
  .background {
    border-radius: 250px 0;
  }

  .chat-list,
  .reason-text,
  .v2-header-3-text {
    width: 80%;
  }

  .cta-section {
    width: 70%;
  }
}
@supports (-webkit-touch-callout: none) {
  .v2-header-3-text {
    letter-spacing: -2px;
  }
}
</style>
