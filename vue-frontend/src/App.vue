<template>
  <Home />
</template>

<script>
import Home from "./components/Home.vue";

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
  font-display: swap;
  src: local("Open Sans"), local("OpenSans"),
    url(https://fonts.gstatic.com/s/productsans/v5/HYvgU2fE2nRJvZ5JFAumwegdm0LZdjqr5-oayXSOefg.woff2)
      format("woff2");
}

body {
  font-family: "Product Sans" !important;
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
</style>
