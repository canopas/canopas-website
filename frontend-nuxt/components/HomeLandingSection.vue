<template>
  <div class="container-fluid tw-relative tw-pt-8 lg:tw-mt-16 lg:tw-pb-24">
    <div class="tw-overflow-hidden background">
      <p
        class="tw-rounded-full tw-absolute tw-z-0 tw-w-2.5 tw-h-2.5 tw-top-4 tw-left-1/4 tw-bg-orange-300 tw-animate-zoomOut"
      ></p>
      <p
        class="tw-rounded-full tw-absolute tw-z-0 tw-hidden tw-w-5 tw-h-5 tw-top-32 tw-left-14 tw-bg-black-900 lg:tw-block tw-animate-zoomIn"
      ></p>
      <p
        class="tw-rounded-full tw-absolute tw-z-0 tw-w-3 tw-h-3 tw-top-96 tw-left-4 tw-bg-gradient-to-r tw-from-pink-300 tw-to-orange-300 tw-animate-zoomOut"
      ></p>
      <p
        class="tw-rounded-full tw-absolute tw-z-0 tw-w-3.5 tw-h-3.5 tw-top-3/4 tw-left-2/4 tw-bg-black-900 tw-animate-zoomIn"
      ></p>
      <p
        class="tw-rounded-full tw-absolute tw-z-0 tw-w-5 tw-h-5 tw-top-0 tw-right-1/4 tw-bg-gradient-to-r tw-from-pink-300 tw-to-orange-300 tw-animate-zoomIn"
      ></p>
      <p
        class="tw-rounded-full tw-absolute tw-z-0 tw-hidden tw-w-2.5 tw-h-2.5 tw-top-32 tw-right-36 tw-bg-black-900 lg:tw-block tw-animate-zoomOut"
      ></p>
      <p
        class="tw-rounded-full tw-absolute tw-z-0 tw-w-2 tw-h-2 tw-top-64 tw-right-2/4 tw-bg-orange-300 tw-animate-zoomOut"
      ></p>
      <p
        class="tw-rounded-full tw-absolute tw-z-0 tw-w-5 tw-h-5 tw-bottom-60 tw-right-20 tw-bg-pink-300 tw-animate-zoomIn"
      ></p>
      <p
        class="tw-rounded-full tw-absolute tw-z-0 tw-w-3 tw-h-3 tw-top-80 tw-right-5 tw-bg-orange-300 tw-animate-zoomOut"
      ></p>
    </div>

    <div class="container tw-mx-auto tw-z-10">
      <div
        class="tw-flex tw-flex-col tw-items-center lg:tw-flex-row-reverse lg:tw-items-start"
      >
        <picture class="tw-flex-1 tw-text-right tw-pb-8">
          <source
            srcset="data:image/gif;base64,R0lGODlhAQABAAD/ACwAAAAAAQABAAACADs="
            media="(max-width: 992px)"
          />
          <img
            class="tw-hidden tw-object-cover lg:tw-block tw-align-top tw-w-full tw-h-full"
            :src="landingImages[2]"
            :srcset="`${landingImages[0]} 400w, ${landingImages[1]} 800w, ${landingImages[2]} 1200w`"
            sizes="33vw"
            alt="landing-view-image"
          />
        </picture>
        <div class="tw-flex-2">
          <div class="description">
            <div
              class="tw-pb-10 tw-w-11/12 v2-header-3-text tw-font-normal lg:tw-pb-16 lg:tw-pt-16 xl:tw-pb-20"
            >
              We develop amazing products to help
              <div class="typewriter tw-flex tw-items-center tw-w-fit">
                <span class="wrapper">
                  <div
                    v-for="animationWord in animationWords"
                    :key="animationWord"
                  >
                    <p
                      v-if="animationWord.isActive"
                      class="v2-canopas-gradient-text tw-font-bold tw-animate-typingErase tw-overflow-hidden tw-border-r-2 tw-border-orange-300 tw-mr-2.5 tw-mt-2.5 tw-mb-0"
                    >
                      {{ animationWord.name }}
                    </p>
                  </div>
                </span>
              </div>
              bring their ideas to life.
            </div>
            <div class="tw-flex tw-flex-col tw-items-center lg:tw-items-start">
              <div class="tw-text-center">
                <router-link
                  class="group v2-button tw-flex tw-items-center tw-py-4 tw-px-20 tw-bg-black-900 tw-text-white hover:tw-bg-white v2-normal-3-text"
                  :to="contactURL"
                >
                  <span class="group-hover:tw-text-black-900">Let's talk</span>
                  <font-awesome-icon
                    class="arrow fa tw-h-4 tw-w-4 group-hover:tw-text-black-900"
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
import landing400w from "@/assets/images/Landing/landing-image-400w.webp";
import landing800w from "@/assets/images/Landing/landing-image-800w.webp";
import landing1200w from "@/assets/images/Landing/landing-image-1200w.webp";

import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
export default {
  data() {
    return {
      landingImages: [landing400w, landing800w, landing1200w],
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
.v2-button:hover > span,
.v2-button:hover > .fa {
  color: #3d3d3d;
}

.background {
  z-index: -1;
}

@include media-breakpoint-up(lg) {
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

@supports (-webkit-touch-callout: none) {
  .v2-header-3-text {
    letter-spacing: -1px;
  }
}
</style>
