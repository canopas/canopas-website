<template>
  <div ref="contactScreen">
    <HeaderV2 v-if="!isNewHeader" />
    <ScreenHeader v-else />
    <ScreenLoader v-if="isLoading" v-bind:loader="true" />
    <div>
      <ContactLanding />
      <ContactForm v-on:isLoading="setLoader" />
    </div>
    <ScreenFooter2 />
  </div>
</template>

<script>
import ScreenHeader from "@/components/partials/ScreenHeader.vue";
import ScreenFooter2 from "@/components/partials/ScreenFooter2.vue";
import ScreenLoader from "@/components/utils/ScreenLoader.vue";
import ContactLanding from "@/components/contact/ContactLanding.vue";
import ContactForm from "@/components/contact/ContactForm.vue";
import HeaderV2 from "@/components/partials/HeaderV2.vue";
import config from "@/config.js";
import { useMeta } from "vue-meta";

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
    ScreenHeader,
    HeaderV2,
    ContactLanding,
    ScreenFooter2,
    ContactForm,
    ScreenLoader,
  },
  data() {
    return {
      isLoading: false,
      isnewheader,
      isNewHeader: config.IS_NEW_HEADER,
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
