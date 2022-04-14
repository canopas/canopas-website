<template>
  <div>
    <router-view></router-view>
  </div>
</template>

<script>
import { library } from "@fortawesome/fontawesome-svg-core";
import {
  faArrowRight,
  faArrowLeft,
  faStar,
} from "@fortawesome/free-solid-svg-icons";
import Config from "@/config.js";
import { useMeta } from "vue-meta";

library.add(faArrowRight, faStar, faArrowLeft);

export default {
  name: "App",
  components: {},
  setup() {
    useMeta({
      meta: [
        {
          name: "robots",
          content: Config.IS_PROD ? "" : "noindex",
          vmid: "robots",
        },
      ],
    });
  },
};
</script>

<style lang="scss">
@import "bootstrap/dist/css/bootstrap.min.css";

@font-face {
  font-family: "Product Sans";
  font-style: normal;
  font-weight: 400;
  font-display: swap;
  src: local("Open Sans"), local("OpenSans"),
    url(https://fonts.gstatic.com/s/productsans/v5/HYvgU2fE2nRJvZ5JFAumwegdm0LZdjqr5-oayXSOefg.woff2)
      format("woff2");
}

@font-face {
  font-family: "FuturaLT-Bold";
  font-style: bold;
  font-weight: 700;
  src: url(./assets/fonts/FuturaLT-Bold.woff2);
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

a {
  text-decoration: none !important;
}

.canopas-gradient-text {
  background: linear-gradient(89.88deg, #f2709c -1.25%, #ff9472 108.15%);
  background-clip: text;
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
  font-size: 1.15rem;
  line-height: 1.5rem;
  color: rgba(61, 61, 61, 0.8);
  letter-spacing: 0.3px;
}

.normal-2-text {
  font-style: normal;
  font-weight: 400;
  font-size: 1.1rem;
  line-height: 1.3rem;
  color: rgba(61, 61, 61, 0.8);
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
  letter-spacing: 0.5px;
}

// 120px/119px - 90px/107px - 65px/94px
.v2-header-text {
  font-family: "FuturaLT-Bold";
  font-weight: 700;
  font-size: 4.0625rem;
  line-height: 5.875rem;
  letter-spacing: 1px;
  color: #3d3d3d;
}

// 90px/107px - 65px/94px - 50px/80px
.v2-header-2-text {
  font-weight: 600;
  font-size: 3.125rem;
  line-height: 5rem;
  letter-spacing: 1px;
  color: #3d3d3d;
}

// 65px/94px - 50px/80px - 35px/60px
.v2-header-3-text {
  font-weight: 600;
  font-size: 2.1875rem;
  line-height: 3.75rem;
  letter-spacing: 1px;
  color: #3d3d3d;
}

// 80px/95px - 60px/75px - 40px/55px
.v2-title-text {
  font-weight: 500;
  font-size: 2.5rem;
  line-height: 3.4375rem;
  letter-spacing: 1px;
  color: #3d3d3d;
}

// 35px/68px - 30px/60px - 25px/45px
.v2-title-2-text {
  font-weight: 500;
  font-size: 1.5625rem;
  line-height: 2.8125rem;
  letter-spacing: 1px;
  color: #3d3d3d !important;
}

// 24px/29px - 20px/24px - 16px/20px
.v2-title-3-text {
  font-weight: 500;
  font-size: 1rem;
  line-height: 1.25rem;
  letter-spacing: 1px;
  color: #3d3d3d;
}

// 27px/43px - 22px/32px - 19px/30px
.v2-normal-text {
  font-weight: 400;
  font-size: 1.1875rem;
  line-height: 1.875rem;
  letter-spacing: 1px;
  color: #3d3d3d;
}

// 22px/32px - 19px/30px - 17px/19px
.v2-normal-2-text {
  font-weight: 400;
  font-size: 1.0625rem;
  line-height: 1.1875rem;
  letter-spacing: 1px;
  color: #3d3d3d;
}

// 18px/30px - 17px/24px - 16px/18px
.v2-normal-3-text {
  font-weight: 400;
  font-size: 1rem;
  line-height: 1.125rem;
  letter-spacing: 1px;
  color: #3d3d3d;
}

.underline-text {
  text-decoration: underline solid #3d3d3d;
  text-underline-offset: 1rem;
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

.v2-white-button {
  border-radius: 3rem;
  padding: 0.7rem 1.5rem;
  text-align: center;
  border: 1px solid #3d3d3d;
  background-color: #fff;
  color: #3d3d3d !important;
}

.v2-button > span {
  margin-right: 6px;
}
.v2-black-btn {
  background-color: #000;
  border-radius: 25px;
  padding: 10px 25px;
  color: #fff;
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
  text-decoration: none;
}

@media (hover: hover) and (pointer: fine) {
  .v2-button:hover {
    background-color: #3d3d3d;
    color: #fff;
  }

  .v2-button:hover > span,
  .v2-button:hover > .fa {
    color: #fff;
  }
  .v2-black-btn:hover > a {
    color: black;
  }
  .v2-black-btn:hover {
    background-color: white;
    color: black;
  }
}

.gradient-btn {
  border-radius: 0.6rem;
  padding: 1rem;
  margin: 1.25rem auto 0 auto;
  text-align: center;
  color: white !important;
  background: linear-gradient(270.11deg, #ff835b -24.42%, #f2709c 101.76%);
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

  .gradient-btn:hover * {
    background: linear-gradient(89.88deg, #f2709c -1.25%, #ff9472 108.15%);
    background-clip: text;
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
  }

  .gradient-btn:hover > .fa,
  .gradient-btn:hover > .fas,
  .gradient-btn:hover > a > .fab {
    color: #f2709c;
  }
}

.gradient-border-btn {
  border-radius: 0.6rem;
  padding: 1rem;
  margin: 5px;
  text-align: center;
  border: 1px solid transparent !important;
  background: linear-gradient(rgba(255, 255, 255, 0), rgba(255, 255, 255, 0)),
    linear-gradient(to top, #ff9472, #f2709c);
  box-shadow: 2px 1000px 1px #fff inset;
}

.gradient-btn:active,
.gradient-border-btn:active,
.v2-button:active {
  transform: scale(0.98);
}

.gradient-border-btn > span {
  margin: 0 6px;
  color: rgba(61, 61, 61, 1);
}

@media (hover: hover) and (pointer: fine) {
  .gradient-border-btn:hover {
    background: linear-gradient(91.53deg, #f2709c 3.91%, #ff9472 100%);
    border: 2px solid transparent;
    box-shadow: none;
  }

  .gradient-border-btn:hover > span,
  .gradient-border-btn:hover > span > span {
    -webkit-text-fill-color: white;
  }

  .gradient-border-btn:hover > .fa,
  .gradient-border-btn:hover > .arrow {
    color: white;
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
    font-size: 1.2rem;
    line-height: 1.375rem;
  }

  // 120px/119px - 90px/107px - 65px/94px
  .v2-header-text {
    font-size: 5.625rem;
    line-height: 6.6875rem;
  }

  // 90px/107px - 65px/94px - 50px/80px
  .v2-header-2-text {
    font-size: 4.0625rem;
    line-height: 5.875rem;
  }

  // 65px/94px - 50px/80px - 35px/60px
  .v2-header-3-text {
    font-size: 3.125rem;
    line-height: 5rem;
  }

  // 80px/95px - 60px/75px - 40px/55px
  .v2-title-text {
    font-size: 3.75rem;
    line-height: 4.6875rem;
  }

  // 35px/68px - 30px/60px - 25px/45px
  .v2-title-2-text {
    font-size: 1.875rem;
    line-height: 3.75rem;
  }

  // 24px/29px - 20px/24px - 16px/20px
  .v2-title-3-text {
    font-size: 1.25rem;
    line-height: 1.5rem;
  }

  // 27px/43px - 22px/32px - 19px/30px
  .v2-normal-text {
    font-size: 1.375rem;
    line-height: 2rem;
  }

  // 22px/32px - 19px/30px - 17px/19px
  .v2-normal-2-text {
    font-size: 1.1875rem;
    line-height: 1.875rem;
  }

  // 18px/30px - 17px/24px - 16px/18px
  .v2-normal-3-text {
    font-size: 1.0625rem;
    line-height: 1.5rem;
  }

  .underline-text {
    text-decoration: underline solid rgba(61, 61, 61, 0.8) 3px;
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

  .normal-2-text {
    font-size: 1.1875rem;
    line-height: 1.4rem;
  }

  // 120px/119px - 90px/107px - 65px/94px
  .v2-header-text {
    font-size: 7.5rem;
    line-height: 7.4375rem;
  }

  // 90px/107px - 65px/94px - 50px/80px
  .v2-header-2-text {
    font-size: 5.625rem;
    line-height: 6.6875rem;
  }

  // 65px/94px - 50px/80px - 35px/60px
  .v2-header-3-text {
    font-size: 4.0625rem;
    line-height: 5.875rem;
  }

  // 80px/95px - 60px/75px - 40px/55px
  .v2-title-text {
    font-size: 5rem;
    line-height: 5.9375rem;
  }

  // 35px/68px - 30px/60px - 25px/45px
  .v2-title-2-text {
    font-size: 2.1875rem;
    line-height: 4.25rem;
  }

  // 24px/29px - 20px/24px - 16px/20px
  .v2-title-3-text {
    font-size: 1.5rem;
    line-height: 1.8125rem;
  }

  // 27px/43px - 22px/32px - 19px/30px
  .v2-normal-text {
    font-size: 1.6875rem;
    line-height: 2.6875rem;
  }

  // 22px/32px - 19px/30px - 17px/19px
  .v2-normal-2-text {
    font-size: 1.375rem;
    line-height: 2rem;
  }

  // 18px/30px - 17px/24px - 16px/18px
  .v2-normal-3-text {
    font-size: 1.125rem;
    line-height: 1.875rem;
  }

  .gradient-btn {
    padding: 1rem 1rem;
    margin: 1.25rem 0.5rem 0 0.5rem;
  }
}
</style>
