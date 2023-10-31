<template>
  <div>
    <HomeNewIndex />
  </div>
</template>

<script>
import config from "@/config.js";
import HomeNewIndex from "@/components/home-new/HomeIndex.vue";

export default {
  setup() {
    const seoData = config.SEO_META_DATA;
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
    HomeNewIndex,
  },
  inject: ["mixpanel"],
  mounted() {
    this.$mixpanel.track("view_new_home_page");
  },
};
</script>
