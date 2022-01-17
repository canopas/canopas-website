<template>
  <div>
    <ScreenHeader />
    <ContactLanding />
    <ContactForm />
    <ScreenFooter />
  </div>
</template>

<script>
import ScreenHeader from "./partials/ScreenHeader.vue";
import ScreenFooter from "./partials/ScreenFooter.vue";
import ContactLanding from "./contact/ContactLanding.vue";
import ContactForm from "./contact/ContactForm.vue";

export default {
  components: {
    ScreenHeader,
    ContactLanding,
    ScreenFooter,
    ContactForm,
  },
};
</script>

<style scoped></style>
