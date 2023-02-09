<template>
  <div ref="contactScreen">
    <metainfo>
      <template v-slot:title="{ content }">
        {{ content }}
      </template>
    </metainfo>
    <Header />
    <ScreenLoader v-if="isLoading" v-bind:loader="true" />
    <div>
      <ContactLanding />
      <ContactForm v-on:isLoading="setLoader" />
    </div>
    <ScreenFooterV2 />
  </div>
</template>

<script>
import Header from "@/components/partials/NewHeader.vue";
import ScreenFooterV2 from "@/components/partials/ScreenFooter2.vue";
import ScreenLoader from "@/components/utils/ScreenLoader.vue";
import ContactLanding from "@/components/contact/ContactLanding.vue";
import ContactForm from "@/components/contact/ContactForm.vue";
import config from "@/config.js";
import { useMeta } from "vue-meta";

import { library } from "@fortawesome/fontawesome-svg-core";
import { faChevronRight } from "@fortawesome/free-solid-svg-icons";

library.add(faChevronRight);

export default {
  setup() {
    var seoData = config.CONTACT_SEO_META_DATA;
    useMeta({
      title: seoData.title,
      description: seoData.description,
      og: {
        type: seoData.type,
        title: seoData.title,
        url: seoData.url,
      },
    });
  },
  components: {
    Header,
    ScreenLoader,
    ScreenFooterV2,
    ContactLanding,
    ContactForm,
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
        import.meta.env.VITE_RECAPTCHA_SITE_KEY
    );
    recaptchaScript.setAttribute("async", "true");
    recaptchaScript.setAttribute("defer", "true");
    document.head.appendChild(recaptchaScript);
    this.mixpanel.track("view_contact_page");
  },
  methods: {
    setLoader(loading) {
      this.isLoading = loading;
    },
  },
};
</script>

<style scoped></style>
