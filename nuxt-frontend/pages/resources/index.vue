<template>
  <div>
    <Header />
    <BlogList
      :mixpanel="$mixpanel"
      :posts="posts"
      :featurePosts="featurePosts"
      :count="count"
      :status="status"
    />
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
import config from "@/config";
import { useBlogListStore } from "@/stores/resources";

const { $mixpanel } = useNuxtApp();
const posts = ref([]);
const postLimit = ref(config.POST_PAGINATION_LIMIT);
const store = useBlogListStore();
const resources = computed(() => store.items);
const status = computed(() => store.status);

await useAsyncData("blogs", () =>
  store.loadResources(config.SHOW_DRAFT_POSTS, true),
);

posts.value = resources.value?.slice(0, postLimit.value);
const count = resources.value?.length || 0;
const featurePosts = resources.value?.filter((post) => post.is_featured);

const handleScroll = () => {
  if (
    window.innerHeight + document.documentElement.scrollTop >=
    document.documentElement.offsetHeight - 100
  ) {
    const oldCount = postLimit.value;
    postLimit.value += 5;
    posts.value.push(...resources.value.slice(oldCount, postLimit.value));
  }
};
onMounted(() => {
  window.addEventListener("scroll", handleScroll);
});

onUnmounted(() => {
  window.removeEventListener("scroll", handleScroll);
});

const seoData = config.RESOURCES_SEO_META_DATA;
useSeoMeta({
  title: seoData.title,
  description: seoData.description,
  ogTitle: seoData.title,
  ogType: seoData.type,
  ogUrl: seoData.url,
  ogImage: seoData.image,
});
</script>
