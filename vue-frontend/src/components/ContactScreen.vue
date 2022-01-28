<template>
  <div>
    <ScreenHeader />
    <div v-on:formSubmitted="!formSubmit" v-if="!formSubmitted">
    <ContactLanding />
    <ContactForm />
    </div>
    <ThankYou v-if="formSubmitted" />
    <ScreenFooter :class="formSubmitted ? 'footer-fixed' : ''" />
  </div>
</template>

<script>
import ScreenHeader from "./partials/ScreenHeader.vue";
import ScreenFooter from "./partials/ScreenFooter.vue";
import ContactLanding from "./contact/ContactLanding.vue";
import ContactForm from "./contact/ContactForm.vue";
import ThankYou from "./contact/ThankYou.vue";

export default {
  data() {
    return {
      formSubmitted: true,
    };
  },
  components: {
    ScreenHeader,
    ContactLanding,
    ScreenFooter,
    ContactForm,
    ThankYou,
  },
  methods: {
    formSubmit(data) {
      this.formSubmitted = data;
      setTimeout(() => {
        this.$router.push("/");
      }, 5000);
    },
  },
};
</script>

<style scoped>
.footer-fixed {
  position: fixed;
  bottom: 0;
}
</style>
