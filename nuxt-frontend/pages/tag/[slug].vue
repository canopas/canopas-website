<template>
  <div>
    <Header />
    <TagList :slug="slug" :posts="posts" :mixpanel="$mixpanel" />
    <BlogFooter
      :mixpanel="$mixpanel"
      :socialMediaData="config.SOCIAL_MEDIA_DATA"
      :apiUrl="config.STRAPI_URL"
      :companyName="config.COMPANY_NAME"
    />
  </div>
</template>

<script setup>
import Header from "@/components/partials/NewHeader.vue";
import { useRoute } from "vue-router";
import config from "@/config";
import { useTagListStore } from "@/stores/tags";

const { $mixpanel } = useNuxtApp();
const route = useRoute();
const slug = ref(route.params.slug);

const store = useTagListStore();
const posts = computed(() => store.items);
const status = computed(() => store.status);

await useAsyncData("tags", () => store.loadTagBlogs(slug.value));

if (status.value !== config.SUCCESS) {
  navigateTo({
    name: "Error404Page",
    params: { pathMatch: ["tag", slug.value] },
  });
}
</script>
