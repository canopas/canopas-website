<template>
  <div ref="contactScreen">
    <ScreenHeader />
    <ContactLanding />
    <ContactForm />
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
    };
  },
  mounted() {
    window.gtag("event", "view_canopas_client_contact");
  },
};
</script>

<style scoped></style>
