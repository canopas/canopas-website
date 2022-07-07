<template>
  <div>
    <HomeIndex />
  </div>
</template>

<script setup>

import config from "@/config.js";

const seoData = config.SEO_META_DATA;

useHead({
  title: seoData.title,
  meta: [
    { name: 'description', content: seoData.description },
    { property: 'og:type', content: seoData.type },
    { property: 'og:title', content: seoData.title },
    { property: 'og:url', content: seoData.url },
  ],
});

</script>

<style scoped></style>
