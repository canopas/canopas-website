<template>
  <div>
    <Header />
    <TagList
      :slug="slug"
      :mixpanel="$mixpanel"
      :showDrafts="config.SHOW_DRAFT_POSTS"
      @notfound="showNotFound"
    />
    <BlogFooter :mixpanel="$mixpanel" />
  </div>
</template>

<script setup>
import Header from "@/components/partials/NewHeader.vue";
import { useRoute } from "vue-router";
import config from "./../../config";

const { $mixpanel } = useNuxtApp();
const route = useRoute();
const slug = ref(route.params.slug);

function showNotFound() {
  showError({
    statusCode: 404,
    fatal: true,
    name: "Error404Page",
  });
}
</script>
