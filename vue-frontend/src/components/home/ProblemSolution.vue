<template>
  <div class="parallax">
    <div class="rellax rellax-image-parent">
      <img class="rellax-img" :src="developerBackgroundImage" alt="Avatar" />
    </div>
    <div class="rellax" data-rellax-speed="7">
      <div class="resource-content">
        <div class="resource-item" id="item1">
          <div
            class="header-text canopas-gradient-text"
            data-scroll-class="item1"
          >
            What's more painful than hiring a developer or team who
            misunderstands every little thing about your project.
          </div>
          <div class="normal-text resources-text" data-scroll-class="item2">
            <div class="text-white">
              <b> THREE REASONS WHY </b> You can’t afford to outsource your
              project to someone who struggles to understand you and your
              business.
            </div>
            <ul class="text-white resources-list mt-20">
              <li>
                <span class="bullet"></span>
                <div>
                  You might end up wasting your precious resources like Time,
                  Money, and Energy.
                </div>
              </li>
              <li>
                <span class="bullet"></span>
                <div>
                  Your final product might have issues like bad quality, poor
                  performance, awful design, and sometimes you won’t receive
                  anything.
                </div>
              </li>
              <li>
                <span class="bullet"></span>
                <div>
                  Most IMPORTANTLY, unsuccessful business relationships might
                  affect your mood and behaviour. You might experience stress,
                  anxiety, negativity, insecurity, uncertainty, etc. It might
                  influence your behaviour, and you may end up doubting everyone
                  around you including YOURSELF.
                </div>
              </li>
            </ul>
          </div>
        </div>
        <div class="resource-item" id="item2">
          <div class="header-text canopas-gradient-text">
            You need a solution that will reward you with money, peace, and
            satisfaction, not headache and nightmares.
          </div>
          <div class="normal-text resources-text">
            <div class="normal-text text-white pb-30">
              You probably need a team that helps you to develop a digital
              product that is…
            </div>

            <ul class="resources-list normal-text text-white">
              <li>
                <span class="bullet"></span>
                <div>
                  Simple<br />
                  Easy to use<br />
                  Beautifully Crafted<br />
                  And Users love it.<br />
                </div>
              </li>
            </ul>
            <div class="normal-text text-white mt-20">
              But, I guess that’s not enough
              <br />
              <br />
              You want somebody who asks questions, suggests ways to improve
              your apps and tries to understand YOU, your vision(or helps you
              achieve your goals or dreams), and the value of your app or
              product.
              <br />
              <br />
              You’re looking for a team who knows that your time is limited and
              knows how to use it wisely to build a high-quality product.
              <br />
              <br />
              Canopas understands these issues and knows how to handle them. So,
              you don't have to.
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script type="module">
import Rellax from "rellax";
import developerBackgroundImage from "@/assets/images/theme/developer.webp";

export default {
  data() {
    return {
      developerBackgroundImage: developerBackgroundImage,
    };
  },
  mounted() {
    new Rellax(".rellax", {
      vertical: true,
      horizontal: false,
      center: true,
    });
  },
};
</script>

<style scoped>
.parallax {
  overflow: hidden;
  margin: 5% auto;
  background-image: l;
}

.rellax-image-parent {
  background: linear-gradient(
    to bottom,
    rgba(61, 61, 61, 0.85) 0%,
    rgba(61, 61, 61, 0.85) 100%
  );
}

.rellax-img {
  width: 100%;
  position: absolute;
  left: 0;
  right: 0;
  top: 0;
  bottom: 0;
  display: block;
}

.resource-content {
  position: relative;
  color: #fff;
  width: 71%;
  margin-left: 15%;
  padding-top: 5%;
}

ul,
li {
  list-style-type: none;
}

li > div {
  width: 100%;
}

#item1 {
  padding-top: 10%;
}

.resource-item {
  padding-bottom: 10%;
}

.resource-item > div,
.resource-item > ul {
  margin: 32px 0;
}

.resources-list > li {
  padding: 0.6rem 0;
  display: flex;
  opacity: 75%;
}

.resources-list > li > div {
  padding: 10px;
}

.resources-list > li > .bullet {
  height: auto;
  width: 5px;
  border-radius: 0;
}

.bullet {
  content: "";
  background: linear-gradient(180deg, #f2709c 0%, #ff9472 100%);
  width: 1rem;
  height: 1rem;
  border-radius: 30%;
  display: inline-block;
  line-height: 1;
  margin: 0.3rem 1.5rem 0 0;
}

@media (min-width: 768px) {
  .header-text {
    line-height: 2rem;
    font-size: 1.875rem;
  }

  .normal-text {
    font-size: 1.125rem;
    line-height: 1.375rem;
  }

  .resource-content {
    width: 49%;
  }

  .carousel-indicator {
    display: flex;
    flex-direction: column;
    position: fixed;
    top: 50%;
    right: 15%;
    transform: translate(-50%, -50%);
  }

  .carousel-indicator > a {
    width: 20px;
    height: 20px;
    margin: 10px;
    cursor: pointer;
    box-sizing: border-box;
    border-radius: 20px;
    transform: rotate(90deg);
    background: #7e7e7e;
  }

  .carousel-indicator > .active {
    width: 20px;
    height: 28px;
    background: linear-gradient(180deg, #ff9472 0%, #f2709c 100%);
    border: 1px solid #ffffff;
    transform: scale(1.6);
  }
}

@media (min-width: 992px) {
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
