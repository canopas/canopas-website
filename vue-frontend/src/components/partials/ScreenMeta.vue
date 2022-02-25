<template>
  <Teleport to="head">
    <title>{{ seoData.title }}</title>
    <meta name="description" :content="seoData.description" />
    <meta property="og:title" :content="seoData.title" />
    <meta property="og:type" :content="seoData.type" />
    <meta property="og:url" :content="seoData.url" />
  </Teleport>
</template>

<script>
export default {
  props : ["seoData"],
};
</script>
