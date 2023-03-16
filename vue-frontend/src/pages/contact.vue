<template>
  <div ref="contactScreen">
    <metainfo>
      <template v-slot:title="{ content }">
        {{ content }}
      </template>
    </metainfo>
    <Header />
    <div class="tw-relative">
      <div
        v-if="message && type == 1"
        class="tw-absolute md:tw-mt-[1.5rem] tw-w-full tw-from-[#ff835b] tw-to-[#f2709c] tw-bg-gradient-[270.11deg] tw-text-center xl:tw-text-right tw-p-[1rem] md:tw-p-[0.5rem] tw-bg-pink-500 tw-font-inter-semibold tw-text-white tw-text-[1rem] md:tw-text-[1.5rem] xl:tw-text-[1.875rem] xl:tw-leading-[2.813rem] tw-right-[-122px] md:"
      >
        <span class="lg:tw-mr-16">
          Thank you for choosing us to make a difference in your business.</span
        >
      </div>
    </div>
    <ScreenLoader v-if="isLoading" v-bind:loader="true" />

    <div
      class="tw-container lg:tw-flex lg:tw-flex-row tw-mt-[.5rem] md:tw-mt-16 lg:tw-mt-32"
    >
      <NewContactLanding class="tw-basis-3/6 2xl:tw-basis-[40%]" />
      <NewContactForm
        @passData="GetData"
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
    NewFooter,
    NewContactLanding,
    NewContactForm,
  },
  data() {
    return {
      isLoading: false,
      message: false,
      type: 0,
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
    GetData(data) {
      this.message = data.message;
      this.type = data.type;
    },
  },
};
</script>

<style scoped></style>
