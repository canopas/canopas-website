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
  -webkit-user-select: none;
  -khtml-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  -o-user-select: none;
  user-select: none;
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
  font-size: 1.5rem;
  line-height: 1.813rem;
}

.header-2-text {
  color: #3d3d3d;
  font-weight: bold;
  font-size: 1.25rem;
  line-height: 1.6rem;
  letter-spacing: 0.1rem;
}

.underline-text {
  border-bottom: 3px solid rgba(61, 61, 61, 0.8);
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

.gradient-btn {
  border-radius: 0.6rem;
  padding: 1rem;
  margin: 1.25rem auto 0 auto;
  text-align: center;
  color: white;
  background: linear-gradient(91.53deg, #f2709c 3.91%, #ff9472 100%);
  border: 1px solid transparent;
}

.gradient-btn > span {
  margin: 0 6px;
  font-weight: 700;
}

@media (hover: hover) and (pointer: fine) {
  .gradient-btn:hover {
    background: linear-gradient(rgba(255, 255, 255, 0), rgba(255, 255, 255, 0)),
      linear-gradient(to top, #ff9472, #f2709c);
    box-shadow: 2px 1000px 1px #fff inset;
  }

  .gradient-btn:hover > span {
    background: linear-gradient(89.88deg, #f2709c -1.25%, #ff9472 108.15%);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
  }

  .gradient-btn:hover > .fa,
  .gradient-btn:hover > .fas {
    color: #f2709c;
  }
}

@include media-breakpoint-up(md) {
  .header-text {
    line-height: 2rem;
    font-size: 1.875rem;
  }

  .header-2-text {
    font-size: 1.75rem;
    line-height: 2rem;
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

  .header-2-text {
    font-size: 2.25rem;
    line-height: 2.729rem;
  }

  .normal-text {
    font-size: 1.5rem;
    line-height: 1.8rem;
  }

  .gradient-btn {
    padding: 1rem 1rem;
    margin: 1.25rem 0.5rem 0 0.5rem;
  }
}
</style>