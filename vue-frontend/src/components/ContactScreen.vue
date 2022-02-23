<template>
  <div ref="contactScreen">
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
    ContactForm,
  },
  data() {
    return {
      showJobs: Config.IS_SHOW_JOBS,
      isLoading: false,
      loader: loader,
    };
  },
  mounted() {
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
