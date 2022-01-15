<template>
  <div>
    <Header />
    <Landing />
    <Footer />
  </div>
</template>

<script>
import Header from "./partials/Header.vue";
import Footer from "./partials/Footer.vue";
import Landing from "./contact/Landing.vue";

export default {
  components: {
    Header,
    Landing,
    Footer,
  },
};
</script>

<style></style>
