<template>
  <div>
    <HomeNewIndex />
  </div>
</template>

<script>
import config from "@/config.js";
import HomeNewIndex from "@/components/home-new/HomeIndex.vue";
import { useMeta } from "vue-meta";

export default {
  setup() {
    var seoData = config.SEO_META_DATA;
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
    HomeNewIndex,
  },
  mounted() {
    this.$gtag.event("view_page_home");
  },
};
</script>

<style scoped></style>
