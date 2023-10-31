<template>
  <div>
    <Header />
    <BlogList :mixpanel="$mixpanel" :showDrafts="config.SHOW_DRAFT_POSTS" />
    <BlogFooter :mixpanel="$mixpanel" />
  </div>
</template>

<script setup>
import Header from "@/components/partials/NewHeader.vue";
import config from "./../../config";

const { $mixpanel } = useNuxtApp();
</script>
