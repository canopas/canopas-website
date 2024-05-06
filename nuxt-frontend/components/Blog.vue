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
      :feature-posts="featurePosts"
      :count="count"
      :status="status"
    />
    <div v-if="showLoader" class="py-10">
      <img :src="loaderImage" class="h-20 w-full" />
    </div>
    <div
      :class="
        count == 0 || status == config.NOT_FOUND
          ? 'absolute bottom-0 w-full'
          : ''
      "
    >
      <BlogFooter
        :mixpanel="$mixpanel"
        :social-media-data="config.SOCIAL_MEDIA_DATA"
        :api-url="config.STRAPI_URL"
        :company-name="config.COMPANY_NAME"
      />
    </div>
  </div>
</template>

<script setup>
import loaderImage from "@/assets/images/theme/loader.svg";
import Header from "@/components/partials/NewHeader.vue";
import config from "@/config";
import { useBlogListStore } from "@/stores/resources";
import { toRefs } from "vue";
const props = defineProps({
  isResource: Boolean,
});

const showLoader = ref(false);
const { isResource } = toRefs(props);
const { $mixpanel } = useNuxtApp();
const posts = ref([]);
const store = useBlogListStore();
const resources = computed(() => store.items);
const featurePosts = computed(() => store.featuredItems);
const count = computed(() => store.totalPosts);
const status = computed(() => store.status);
let postLimit = config.POST_PAGINATION_LIMIT;

await useAsyncData("blogs", () =>
  store.loadResources(config.SHOW_DRAFT_POSTS, isResource.value, 0, postLimit),
);

if (status.value === config.SUCCESS) {
  posts.value.push(...resources.value);
}

const handleScroll = () => {
  if (
    window.innerHeight + document.documentElement.scrollTop >=
    document.documentElement.offsetHeight - 100
  ) {
    if (posts.value.length >= count.value) {
      return;
    }

    showLoader.value = true;

    useAsyncData("paginate", () =>
      store.loadPaginateResources(
        config.SHOW_DRAFT_POSTS,
        isResource.value,
        posts.value.length,
        postLimit,
      ),
    ).then(() => {
      showLoader.value = false;
      posts.value.push(...resources.value);
      posts.value = Array.from(new Set(posts.value.map(JSON.stringify))).map(
        JSON.parse,
      );
    });
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
