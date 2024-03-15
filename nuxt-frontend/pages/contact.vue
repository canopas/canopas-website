<template>
  <div ref="contactScreen">
    <Header />
    <ScreenLoader v-if="isLoading" v-bind:loader="true" />
    <div class="container lg:flex lg:flex-row mt-6 md:mt-0 lg:mt-20">
      <NewContactLanding class="basis-3/6 2xl:basis-[40%]" />
      <NewContactForm
        v-on:isLoading="setLoader"
        class="basis-3/6 lg:basis-[70%] 2xl:basis-[60%]"
      />
    </div>
    <Guaranty id="guaranty" />
    <NewFooter />
  </div>
</template>

<script>
import Header from "@/components/partials/NewHeader.vue";
import NewFooter from "@/components/partials/NewFooter.vue";
import ScreenLoader from "@/components/utils/ScreenLoader.vue";
import NewContactLanding from "@/components/contact/NewContactLanding.vue";
import NewContactForm from "@/components/contact/NewContactForm.vue";
import Guaranty from "@/components/contact/Guaranty.vue";
import config from "@/config.js";

export default {
  setup() {
    const seoData = config.CONTACT_SEO_META_DATA;
    useSeoMeta({
      title: seoData.title,
      description: seoData.description,
      ogTitle: seoData.title,
      ogType: seoData.type,
      ogUrl: seoData.url,
      ogImage: seoData.image,
    });
  },
  components: {
    Header,
    ScreenLoader,
    NewContactLanding,
    NewContactForm,
    Guaranty,
    NewFooter,
  },
  data() {
    return {
      isLoading: false,
    };
  },
  inject: ["mixpanel"],
  mounted() {
    let recaptchaScript = document.createElement("script");
    recaptchaScript.setAttribute(
      "src",
      "https://www.google.com/recaptcha/enterprise.js?render=" +
        config.VITE_RECAPTCHA_SITE_KEY,
    );
    recaptchaScript.setAttribute("async", "true");
    recaptchaScript.setAttribute("defer", "true");
    document.head.appendChild(recaptchaScript);
    this.$mixpanel.track("view_contact_page");
  },
  methods: {
    setLoader(loading) {
      this.isLoading = loading;
    },
  },
};
</script>
