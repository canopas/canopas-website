<template>
  <div v-if="!assets">
    <Header :animateOnScroll="false" />
    <div :class="post.is_resource ? '' : 'container pl-0'">
      <BlogDetail
        :mixpanel="$mixpanel"
        :recaptcha-key="config.VITE_RECAPTCHA_SITE_KEY"
        :post="post"
        :website-url="config.BASE_URL"
        :contact-api-url="config.API_BASE"
      />
    </div>
    <div class="cta-section">
      <CTA1 v-if="CTACompName == 'CTA1'" />
      <CTA2
        v-if="CTACompName == 'CTA2'"
        :recaptcha-key="config.VITE_RECAPTCHA_SITE_KEY"
        :contact-api-url="config.API_BASE"
      />
      <CTA3
        v-if="CTACompName == 'CTA3'"
        :recaptcha-key="config.VITE_RECAPTCHA_SITE_KEY"
        :contact-api-url="config.API_BASE"
      />
      <div
        v-if="CTACompName == 'CTA4'"
        class="mt-5 lg:mt-[150px] xl:mt-[170px] 2xl:mt-[250px] 2x:mt-[280px]"
      >
        <CTA4 />
      </div>
      <CTA5 v-if="CTACompName == 'CTA5'" />
    </div>
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
import CTA1 from "@/components/CTA/CTA1.vue";
import CTA2 from "@/components/CTA/CTA2.vue";
import CTA3 from "@/components/CTA/CTA3.vue";
import CTA4 from "@/components/CTA/CTA4.vue";
import CTA5 from "@/components/CTA/CTA5.vue";
import { useRoute } from "vue-router";
import config from "@/config";
import { useBlogDetailStore } from "@/stores/resources";

const { $mixpanel } = useNuxtApp();
const post = ref({});
const route = useRoute();
const slug = ref(route.params.slug);
const CTACompName = ref("");

const store = useBlogDetailStore();
const postData = computed(() => store.item);
const status = computed(() => store.status);

let published_time;

const assets =
  slug.value === "favicon.ico" || slug.value === "flutter_service_worker.js";

if (!assets) {
  await useAsyncData("blog", () =>
    store.loadResource(slug.value, config.SHOW_DRAFT_POSTS),
  );
}

if (status.value !== config.SUCCESS) {
  navigateTo({
    name: "Error404Page",
    params: { pathMatch: [slug.value] },
  });
} else {
  post.value = postData?.value;
  published_time = new Date(post.value?.published_on).toLocaleTimeString();
}

const CTAData = post.value?.cta;
CTACompName.value = CTAData?.component_name;

useHead({
  script: [
    {
      innerHTML: JSON.stringify(getJsonLdSchema()),
      type: "application/ld+json",
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
    ogUrl: config.BASE_URL + "/" + post.value.slug,
    ogDescription: post.value.meta_description,
    twitterTitle: post.value.title,
    twitterDescription: post.value.meta_description,
    twitterSite: "https://canopas.com/",
    twitterCard: "summary_large_image",
    twitterLabel1: "Reading time",
    twitterData1: post.value.reading_time + ` min read`,
    twitterCta: "Read on Canopas",
    keywords: post.value.keywords,
    twitterTileInfo1Icon: "Person",
    twitterTileInfo1Text: post.value.author.name,
    twitterTileInfo2Icon: "Calendar",
    twitterTileInfo2Text: post.value.published_on,
    twitterImageSrc: post.value.image_url,
    twitterTileImage: post.value.image_url,
    articleAuthor: post.value.author.name,
    articlePublished_time: published_time,
    ogSite_name: "Canopas blogs",
    author: post.value.author.name,
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
      url: config.BASE_URL,
    },
    url: config.BASE_URL + "/" + post.value.slug,
    datePublished: post.value.published_on,
    dateModified: post.value.published_on,
    description: post.value.meta_description,
    author: {
      "@type": "Person",
      name: post.value.author.name,
    },
    mainEntityOfPage: {
      "@type": "Blog Website",
      "@id":
        "https://canopas.com/" + post.value.is_resource ? "blog" : "resources",
    },
  };
}
//for copy to clipboard of code blocks
async function copyCode(bash) {
  const code = bash.querySelector("code");
  const text = code.innerText;
  await navigator.clipboard.writeText(text);
}

onMounted(() => {
  $mixpanel.track("view_blog_page", { Title: post.value.title });

  const bashes = document.querySelectorAll("pre");
  bashes.forEach((bash) => {
    if (navigator.clipboard) {
      const wrapperDiv = document.createElement("div");
      bash.parentNode.insertBefore(wrapperDiv, bash);
      wrapperDiv.appendChild(bash);
      wrapperDiv.classList.add("relative");

      const button = document.createElement("button");
      button.classList.add("copy-btn");
      button.innerText = "copy";

      bash.appendChild(button);

      const isSingleLine = bash.textContent.trim().split("\n").length === 1;
      if (isSingleLine) {
        bash.classList.add("!pr-20");
      }

      button.addEventListener("click", async () => {
        await copyCode(bash);
        button.innerText = "copied!";
        button.classList.add("!text-green-600");

        setTimeout(() => {
          button.innerText = "copy";
          button.classList.remove("!text-green-600");
        }, 2000);
      });
    }
  });
});
</script>
<style lang="postcss">
.copy-btn {
  @apply absolute top-3 right-3 border-1 px-2 bg-white text-[#1f2937] rounded-lg font-semibold;
}
</style>
