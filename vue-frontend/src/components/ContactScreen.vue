<template>
  <div ref="contactScreen">
    <ScreenMeta v-bind:seoData="seoData" />
    <ScreenHeader />
    <ScreenLoader v-if="isLoading" v-bind:loader="true" />
    <div>
      <ContactLanding />
      <ContactForm v-on:isLoading="setLoader" />
    </div>
    <ScreenFooter2 />
  </div>
</template>

<script>
import ScreenHeader from "./partials/ScreenHeader.vue";
import ScreenFooter2 from "./partials/ScreenFooter2.vue";
import ScreenMeta from "./partials/ScreenMeta.vue";
import ScreenLoader from "./utils/ScreenLoader.vue";
import ContactLanding from "./contact/ContactLanding.vue";
import ContactForm from "./contact/ContactForm.vue";
import config from "@/config.js";

export default {
  components: {
    ScreenHeader,
    ContactLanding,
    ScreenFooter2,
    ScreenMeta,
    ContactForm,
    ScreenLoader,
  },
  data() {
    return {
      isLoading: false,
      seoData: config.CONTACT_SEO_META_DATA,
    };
  },
  mounted() {
    this.$gtag.event("view_page_contact");
  },
  methods: {
    setLoader(loading) {
      this.isLoading = loading;
    },
  },
};
</script>

<style scoped></style>
