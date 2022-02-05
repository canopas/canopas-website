<template>
  <router-view></router-view>
</template>

<script>
import { library } from "@fortawesome/fontawesome-svg-core";
import {
  faArrowRight,
  faArrowLeft,
  faStar,
} from "@fortawesome/free-solid-svg-icons";

library.add(faArrowRight, faStar, faArrowLeft);

export default {
  name: "App",
  components: {},
};
</script>

<style lang="scss">
@import "~bootstrap/dist/css/bootstrap.min.css";

@font-face {
  font-family: "Product Sans";
  font-style: normal;
  font-weight: 400;
  font-display: swap;
  src: local("Open Sans"), local("OpenSans"),
    url(https://fonts.gstatic.com/s/productsans/v5/HYvgU2fE2nRJvZ5JFAumwegdm0LZdjqr5-oayXSOefg.woff2)
      format("woff2");
}

body {
  font-family: "Product Sans" !important;
}

body::-webkit-scrollbar {
  display: none;
}

.canopas-gradient-text {
  background: linear-gradient(89.88deg, #f2709c -1.25%, #ff9472 108.15%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.abs-fit-parent {
  position: absolute;
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
}

.normal-text {
  font-style: normal;
  font-weight: 400;
  font-size: 1.1rem;
  line-height: 1.5rem;
  color: rgba(0, 0, 0, 0.5);
}

.header-text {
  font-weight: 700;
  line-height: 1.813rem;
  font-size: 1.5rem;
}

.swiper-wrapper {
  height: 90% !important;
}

.swiper-pagination-bullet {
  width: 15px !important;
  height: 15px !important;
  border-radius: 3px !important;
  background: transparent !important;
  border: 1px solid !important;
}

.swiper-pagination-bullet-active {
  border: none !important;
  background: linear-gradient(
    45deg,
    rgb(246, 146, 89),
    rgb(241, 105, 117)
  ) !important;
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

  .swiper-wrapper {
    height: 65% !important;
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

  .swiper-wrapper {
    height: 60% !important;
  }
}
</style>
