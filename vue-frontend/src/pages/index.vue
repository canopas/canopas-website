<template>
  <div>
    <HomeIndex v-if="!isShowNewHomePage" />
    <HomeNewIndex v-else />
  </div>
</template>

<script>
import config from "@/config.js";
import HomeIndex from "@/components/home/HomeIndex.vue";
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
  data() {
    return {
      isShowNewHomePage: config.IS_SHOW_NEW_HOME_PAGE,
    };
  },
  components: {
    HomeIndex,
    HomeNewIndex,
  },
  mounted() {
    this.$gtag.event("view_page_home");
  },
};
</script>

<style scoped></style>
