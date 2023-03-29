<template>
  <div ref="contactScreen">
    <metainfo>
      <template v-slot:title="{ content }">
        {{ content }}
      </template>
    </metainfo>
    <Header />
    <SuccessMessage v-if="message" />
    <ScreenLoader v-if="isLoading" v-bind:loader="true" />

    <div class="tw-container lg:tw-flex lg:tw-flex-row lg:tw-mt-20">
      <NewContactLanding class="tw-basis-3/6 2xl:tw-basis-[40%]" />
      <NewContactForm
        @showMessage="showSuccessMessage"
        v-on:isLoading="setLoader"
        class="tw-basis-3/6 lg:tw-basis-[70%] 2xl:tw-basis-[60%] lg:tw--mr-14 xl:tw--mr-24"
      />
    </div>
    <NewFooter />
  </div>
</template>

<script>
import Header from "@/components/partials/NewHeader.vue";
import NewFooter from "@/components/partials/NewFooter.vue";
import ScreenLoader from "@/components/utils/ScreenLoader.vue";
import NewContactLanding from "@/components/contact/NewContactLanding.vue";
import NewContactForm from "@/components/contact/NewContactForm.vue";
import SuccessMessage from "@/components/contact/SuccessMessage.vue";
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
        image: seoData.image,
      },
    });
  },
  components: {
    Header,
    ScreenLoader,
    NewFooter,
    NewContactLanding,
    NewContactForm,
    SuccessMessage,
  },
  data() {
    return {
      isLoading: false,
      message: false,
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
    showSuccessMessage(data) {
      this.message = data;
    },
  },
};
</script>

<style scoped></style>
