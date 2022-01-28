<template>
  <div>
    <ScreenHeader />
    <ContactLanding v-on:formSubmitted="formSubmit" v-if="formSubmitted" />
    <ContactForm v-on:formSubmitted="formSubmit" v-if="formSubmitted" />
    <div class="finally" :style="!formSubmitted ? 'position: absolute;' : ''">
      <ThankYou v-if="!formSubmitted" />
      <ScreenFooter />
    </div>
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
.finally {
  height: 60%;
  width: 100%;
}
</style>
