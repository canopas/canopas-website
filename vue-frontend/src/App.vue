<template>
  <Home />
</template>

<script>
import { library } from "@fortawesome/fontawesome-svg-core";
import {
  faArrowRight,
  faArrowLeft,
  faStar,
} from "@fortawesome/free-solid-svg-icons";
import Home from "./components/Home.vue";

library.add(faArrowRight, faStar, faArrowLeft);

export default {
  name: "App",
  components: {
    Home,
  },
};
</script>

<style>
@font-face {
  font-family: "Product Sans";
  font-style: normal;
  font-weight: 400;
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
</style>
