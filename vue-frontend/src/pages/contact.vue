<template>
  <div ref="contactScreen">
    <ScreenHeaderV2 />
    <ScreenLoader v-if="isLoading" v-bind:loader="true" />
    <div>
      <ContactLanding />
      <ContactForm v-on:isLoading="setLoader" />
    </div>
    <ScreenFooterV2 />
  </div>
</template>

<script>
import ScreenHeaderV2 from "@/components/partials/ScreenHeaderV2.vue";
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
    ScreenHeaderV2,
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
  mounted() {
    this.$gtag.event("view_page_contact");
  },
  methods: {
    setLoader(loading) {
      this.isLoading = loading;
    },
  },
};
</script>

<style scoped></style>
