<template>
  <div>
    <Header />
    <BlogDetail
      :slug="slug"
      :mixpanel="$mixpanel"
      :recaptchaKey="config.VITE_RECAPTCHA_SITE_KEY"
      :post="post"
      :recommandedPosts="recommandedPosts"
      :websiteUrl="config.BASE_URL"
      :contactApiUrl="config.API_BASE"
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
import { useRoute } from "vue-router";
import config from "@/config";
import {
  useBlogDetailStore,
  useRecommandedBlogStore,
} from "@/stores/resources";
import { filterPostsByCategoryAndTag } from "@/utils";

const { $mixpanel } = useNuxtApp();
const post = ref([]);
const route = useRoute();
const slug = ref(route.params.slug);

const store = useBlogDetailStore();
const postData = computed(() => store.item);
const status = computed(() => store.status);

const recStore = useRecommandedBlogStore();
const recommandedBlog = computed(() => recStore.items);

let published_time, recommandedPosts;
await useAsyncData("blog", () => store.loadResource(slug.value));

if (status.value !== config.SUCCESS) {
  navigateTo({
    name: "Error404Page",
    params: { pathMatch: [slug.value] },
  });
} else {
  post.value = postData?.value;
}

published_time = new Date(post?.value?.published_on).toLocaleTimeString();

await useAsyncData("recommandedBlog", () =>
  recStore.loadRecommandedBlog(slug.value, config.SHOW_DRAFT_POSTS),
);

recommandedPosts = filterPostsByCategoryAndTag(
  post.value,
  recommandedBlog.value,
);

useHead({
  script: [
    {
      innerHTML: JSON.stringify(getJsonLdSchema()),
      type: "application/ld+json",
    },
    {
      src:
        "//cdn.iframe.ly/embed.js?card=small&key=" + config.VITE_IFRAMELY_KEY,
      onLoad: () => {
        // Load media preview
        document.querySelectorAll("oembed[url]").forEach((element) => {
          iframely.load(element, element.attributes.url.value);
        });
      },
    },
  ],
});

useSeoMeta(seoData());
function seoData() {
  return {
    title: post.value.title,
    description: post.value.meta_description,
    ogTitle: post.value.title,
    ogType: "article",
    ogImage: post.value.image_url,
    ogUrl: config.BASE_URL + post.value.slug,
    ogDescription: post.value.meta_description,
    twitterTitle: post.value.title,
    twitterDescription: post.value.meta_description,
    twitterSite: "https://canopas.com/",
    twitterCard: "summary_large_image",
    twitterLabel1: "Reading time",
    twitterData1: post.value.readingTime + ` min read`,
    twitterCta: "Read on Canopas",
    keywords: post.value.keywords,
    twitterTileInfo1Icon: "Person",
    twitterTileInfo1Text: post.value.authorName,
    twitterTileInfo2Icon: "Calendar",
    twitterTileInfo2Text: post.value.published_on,
    twitterImageSrc: post.value.image_url,
    twitterTileImage: post.value.image_url,
    articleAuthor: post.value.authorName,
    articlePublished_time: published_time,
    ogSite_name: "Canopas blogs",
    author: post.value.authorName,
  };
}

function getJsonLdSchema() {
  return {
    "@context": "https://schema.org",
    "@type": "Article",
    headline: post.value.title,
    image: post.value.image_url,
    publisher: {
      "@type": "Organization",
      name: "Canopas",
      url: "https://canopas.com/",
    },
    url: config.BASE_URL + post.value.slug,
    datePublished: post.value.published_on,
    dateModified: post.value.published_on,
    description: post.value.meta_description,
    author: {
      "@type": "Person",
      name: post.value.authorName,
    },
    mainEntityOfPage: {
      "@type": "Blog Website",
      "@id": "https://canopas.com/resources",
    },
  };
}
//for copy to clipboard of code blocks
function copyCode(bash) {
  const code = bash.querySelector("code");
  const text = code.innerText;
  navigator.clipboard.writeText(text);
}
onMounted(() => {
  const bashes = document.querySelectorAll("pre");
  bashes.forEach((bash) => {
    if (navigator.clipboard) {
      const wrapperDiv = document.createElement("div");
      bash.parentNode.insertBefore(wrapperDiv, bash);
      wrapperDiv.appendChild(bash);
      wrapperDiv.classList.add("relative");

      const isSingleLine = bash.textContent.trim().split("\n").length === 1;
      if (isSingleLine) {
        bash.classList.add("!pr-20");
      }
      const button = document.createElement("button");
      button.classList.add("copy-btn");
      button.innerText = "copy";
      bash.appendChild(button);
      button.addEventListener("click", async () => {
        await copyCode(bash, button);
        button.innerText = "copied!";
        button.classList.add("text-green-500");
      });
    }
  });
});
</script>
<style lang="postcss">
.copy-btn {
  @apply absolute top-3 right-3 border-1 px-2 bg-gray-700 rounded-lg;
}
</style>
