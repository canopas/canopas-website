<template>
  <div ref="contactScreen">
    <ScreenMeta v-bind:seoData="seoData" />
    <ScreenHeader />
    <div v-if="isLoading" class="loader-div">
      <img :src="loader" />
    </div>
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
import ContactLanding from "./contact/ContactLanding.vue";
import ContactForm from "./contact/ContactForm.vue";
import Config from "@/config.js";
import loader from "@/assets/images/theme/loader.svg";

export default {
  components: {
    ScreenHeader,
    ContactLanding,
    ScreenFooter,
    ScreenFooter2,
    ScreenMeta,
    ContactForm,
  },
  data() {
    return {
      showJobs: Config.IS_SHOW_JOBS,
      isLoading: false,
      loader: loader,
      seoData: Config.SEO_META_DATA,
    };
  },
  mounted() {
    this.seoData.title =
      "Let's discuss your idea or problem you're facing in your business.";
    this.seoData.description =
      "Thank you for choosing Canopas for your business. Fill out all the required information about you and your business. We will get back to you soon.";
    window.gtag("event", "view_canopas_client_contact");
  },
  methods: {
    setLoader(loading) {
      this.isLoading = loading;
    },
  },
};
</script>
 
<style scoped>
.loader-div {
  position: fixed;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  z-index: 100;
}
</style>
