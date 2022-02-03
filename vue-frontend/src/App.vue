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

.red-text{
  font-style: normal;
  font-weight: 400;
  font-size: 1.1rem;
  line-height: 1.5rem;
  color: #E8505B;
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

@include media-breakpoint-up(md) {
  .header-text {
    line-height: 2rem;
    font-size: 1.875rem;
  }

  .normal-text {
    font-size: 1.125rem;
    line-height: 1.375rem;
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
}
</style>
