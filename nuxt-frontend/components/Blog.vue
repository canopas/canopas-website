<template>
  <div :class="count == 0 || status == config.NOT_FOUND ? 'h-screen' : ''">
    <Header />
    <div
      v-if="count == 0 || status == config.NOT_FOUND"
      class="h-1/2 flex text-[1.4rem] text-black-900 items-center justify-center"
    >
      {{ config.POST_NOT_FOUND_MESSAGE }}
    </div>
    <BlogList
      v-else
      :mixpanel="$mixpanel"
      :posts="posts"
      :featurePosts="featurePosts"
      :count="count"
      :status="status"
    />
    <div
      :class="
        count == 0 || status == config.NOT_FOUND
          ? 'absolute bottom-0 w-full'
          : ''
      "
    >
      <BlogFooter
        :mixpanel="$mixpanel"
        :socialMediaData="config.SOCIAL_MEDIA_DATA"
        :apiUrl="config.STRAPI_URL"
        :companyName="config.COMPANY_NAME"
      />
    </div>
  </div>
</template>

<script setup>
import Header from "@/components/partials/NewHeader.vue";
import config from "@/config";
import { useBlogListStore } from "@/stores/resources";
import { toRefs } from "vue";

const props = defineProps({
  isResource: Boolean,
});

const { isResource } = toRefs(props);
const { $mixpanel } = useNuxtApp();
const posts = ref([]);
const store = useBlogListStore();
const resources = computed(() => store.items);
const count = computed(() => store.totalPosts);
const status = computed(() => store.status);
let postLimit = config.POST_PAGINATION_LIMIT;
let featurePosts = [];

await useAsyncData("blogs", () =>
  store.loadResources(config.SHOW_DRAFT_POSTS, isResource.value, 0, postLimit),
);

if (status.value === config.SUCCESS) {
  posts.value.push(...resources.value);
  featurePosts = resources.value?.filter((post) => post.is_featured);
}

const handleScroll = () => {
  if (
    window.innerHeight + document.documentElement.scrollTop >=
    document.documentElement.offsetHeight - 100
  ) {
    if (posts.value.length < count.value) {
      useAsyncData("blogs", () =>
        store.loadResources(
          config.SHOW_DRAFT_POSTS,
          false,
          posts.value.length,
          5,
        ),
      ).then(() => {
        posts.value.push(...resources.value);
      });
    }
  }
};

onMounted(() => {
  window.addEventListener("scroll", handleScroll);
});

onUnmounted(() => {
  window.removeEventListener("scroll", handleScroll);
});

const seoData = isResource.value
  ? config.RESOURCES_SEO_META_DATA
  : config.BLOG_SEO_META_DATA;
useSeoMeta({
  title: seoData.title,
  description: seoData.description,
  ogTitle: seoData.title,
  ogType: seoData.type,
  ogUrl: seoData.url,
  ogImage: seoData.image,
});
</script>
