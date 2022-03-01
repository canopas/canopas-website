<template>
  <div ref="contactScreen">
    <ScreenMeta v-bind:seoData="seoData" />
    <ScreenHeader />
    <ScreenLoader v-if="isLoading" v-bind:loader="true" />
    <div
      :style="{
        'pointer-events': isLoading ? 'none' : '',
        filter: isLoading ? 'blur(1px)' : '',
      }"
    >
      <ContactLanding />
      <ContactForm v-on:isLoading="setLoader" />
    </div>
    <ScreenFooter v-if="!showJobs" />
    <ScreenFooter2 v-if="showJobs" />
  </div>
</template>

<script>
import ScreenHeader from "./partials/ScreenHeader.vue";
import ScreenFooter from "./partials/ScreenFooter.vue";
import ScreenFooter2 from "./partials/ScreenFooter2.vue";
import ScreenMeta from "./partials/ScreenMeta.vue";
import ScreenLoader from "./utils/ScreenLoader.vue";
import ContactLanding from "./contact/ContactLanding.vue";
import ContactForm from "./contact/ContactForm.vue";
import Config from "@/config.js";

export default {
  components: {
    ScreenHeader,
    ContactLanding,
    ScreenFooter,
    ScreenFooter2,
    ScreenMeta,
    ContactForm,
    ScreenLoader,
  },
  data() {
    return {
      showJobs: Config.IS_SHOW_JOBS,
      isLoading: false,
      seoData: Config.SEO_META_DATA,
    };
  },
  mounted() {
    this.seoData.title =
      "Let's discuss your idea or problem you're facing in your business.";
    this.seoData.description =
      "Thank you for choosing Canopas for your business. Fill out all the required information about you and your business. We will get back to you soon.";
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
