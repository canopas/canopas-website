<template>
  <div>
    <Header />
    <section class="container min-h-[50vh]">
      <div
        v-if="status == config.NOT_FOUND"
        class="h-1/2 flex text-[1.4rem] text-black-900 items-center justify-center"
      >
        {{ config.POST_NOT_FOUND_MESSAGE }}
      </div>
      <div v-else class="md:mx-8 xl:mx-20">
        <nuxt-link :to="'/author/' + slug" class="flex space-x-4 items-center">
          <div class="w-8 h-8 md:w-9 md:h-9">
            <Icon
              name="fa6-solid:user"
              class="w-full h-full"
              v-if="posts[0].author.image == null"
            />
            <img
              v-else
              :src="posts[0].author.image"
              :alt="slug"
              class="w-full h-full rounded-full"
            />
          </div>

          <h1
            class="my-6 md:my-10 text-2xl md:text-3xl xl:text-4xl leading-7 tracking-none capitalize font-inter-semibold"
          >
            {{ posts[0].author.name }}
          </h1>
        </nuxt-link>
        <div class="mb-6 md:mb-10">
          <AuthorPosts :posts="posts" :mixpanel="mixpanel" />
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
import { useAuthorListStore } from "@/stores/author";

const { $mixpanel } = useNuxtApp();
const route = useRoute();
const slug = ref(route.params.slug);
const posts = ref([]);
const store = useAuthorListStore();
const resources = computed(() => store.items);
const status = computed(() => store.status);
let postLimit = 10;

await useAsyncData("authors", () =>
  store.loadAuthorBlogs(config.SHOW_DRAFT_POSTS, slug.value, 0, postLimit),
);

if (status.value === config.SUCCESS) {
  posts.value = resources.value?.slice(0, postLimit);
}

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
  title: "Stories by " + slug.value + " | Canopas",
  ogTitle: "Stories by " + slug.value + " | Canopas",
  ogImage: config.BASE_URL + "/apple-touch-icon.png",
  ogUrl: config.BASE_URL + "/author/" + slug.value,
  twitterTitle: "Stories by " + slug.value + " | Canopas",
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
