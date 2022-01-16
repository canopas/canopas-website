<template>
  <div>
    <Header />
    <Landing />
    <Form />
    <Footer />
  </div>
</template>

<script>
import Header from "./partials/Header.vue";
import Footer from "./partials/Footer.vue";
import Landing from "./contact/Landing.vue";
import Form from "./contact/Form.vue";

export default {
  components: {
    Header,
    Landing,
    Footer,
    Form,
  },
};
</script>

<style></style>
