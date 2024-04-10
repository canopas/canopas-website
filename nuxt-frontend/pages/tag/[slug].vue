<template>
  <div>
    <Header />
    <div
      v-if="slug == '' || status == config.NOT_FOUND"
      class="h-1/2 flex text-[1.4rem] text-black-900 items-center justify-center"
    >
      {{ config.POST_NOT_FOUND_MESSAGE }}
    </div>
    <section
      v-else
      class="container min-h-[50vh] my-10 md:my-16 sm:mx-auto 3xl:px-24"
    >
      <div class="md:mx-8 xl:mx-20">
        <div class="flex space-x-4 items-center">
          <div class="w-6 h-6 md:w-7 md:h-7 mt-1">
            <Icon name="fa6-solid:tags" class="w-full h-full" />
          </div>
          <h1
            class="my-6 md:my-10 text-2xl md:text-3xl xl:text-4xl leading-7 tracking-none capitalize font-inter-semibold"
          >
            {{ tagName }}
          </h1>
        </div>
        <div class="mt-4 md:mt-6 xl:mt-8">
          <PostList :posts="posts" :slug="slug" :mixpanel="mixpanel" />
        </div>
      </div>
    </section>
    <BlogFooter
      :mixpanel="$mixpanel"
      :social-media-data="config.SOCIAL_MEDIA_DATA"
      :api-url="config.STRAPI_URL"
      :company-name="config.COMPANY_NAME"
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
const posts = ref([]);
const store = useTagListStore();
const resources = computed(() => store.items);
const status = computed(() => store.status);
let postLimit = 2;

await useAsyncData("tags", () => store.loadTagBlogs(slug.value));

if (status.value === config.SUCCESS) {
  posts.value = resources.value?.slice(0, postLimit);
}

const tagName = posts.value[0].tags.filter((tag) => tag.slug === slug.value)[0]
  .name;

const handleScroll = () => {
  if (
    window.innerHeight + document.documentElement.scrollTop >=
    document.documentElement.offsetHeight - 100
  ) {
    const oldCount = postLimit;
    postLimit += 2;
    posts.value.push(...resources.value.slice(oldCount, postLimit));
  }
};

onMounted(() => {
  window.addEventListener("scroll", handleScroll);
});

onUnmounted(() => {
  window.removeEventListener("scroll", handleScroll);
});

useSeoMeta({
  title: "Stories on " + slug.value + " | Canopas",
  ogTitle: "Stories on " + slug.value + " | Canopas",
  ogImage: config.BASE_URL + "/apple-touch-icon.png",
  ogUrl: config.BASE_URL + "/tag/" + slug.value,
  twitterTitle: "Stories on " + slug.value + " | Canopas",
  twitterSite: "https://canopas.com/",
  twitterCard: "summary_large_image",
  twitterCta: "Read on Canopas",
  keywords: slug,
  twitterImageSrc: config.BASE_URL + "/apple-touch-icon.png",
  twitterTileImage: config.BASE_URL + "/apple-touch-icon.png",
  ogSite_name: "Canopas",
  author: "canopas",
});
</script>
