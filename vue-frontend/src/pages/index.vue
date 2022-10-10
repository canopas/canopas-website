<template>
  <div>
    <metainfo>
      <template v-slot:title="{ content }">
        {{ content }}
      </template>
    </metainfo>
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
};
</script>

<style scoped></style>
