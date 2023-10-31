<template>
  <div>
    <Header />
    <div
      class="container flex items-center justify-center flex-col my-0 mx-auto py-[150px] px-[5%] min-h-[50vh] xll:min-h-[79vh]"
    >
      <div class="flex flex-row my-0 mx-auto text-center">
        <img :src="firstErrorLetter" alt="404" />
        <img :src="middleErrorLetter" class="m-4 w-[6.5rem]" alt="404" />
        <img :src="lastErrorLetter" alt="404" />
      </div>
      <h1 class="mt-8 text-center text-[1.4rem] leading-8">
        The page you’re looking for was moved, renamed or might never existed.
      </h1>
      <router-link
        class="my-16 gradient-border-btn flex flex-row items-center"
        to="/"
      >
        <Icon
          class="arrow w-6 h-6 fa text-[#f2709c] text-[2.1rem] mr-[5px]"
          name="fa6-solid:arrow-left"
          id="leftArrow"
        />
        <span class="text-center font-bold text-[1.1rem]"
          >Back to Home page</span
        >
      </router-link>
    </div>
    <NewFooter />
  </div>
</template>

<script>
import Header from "@/components/partials/NewHeader.vue";
import NewFooter from "@/components/partials/NewFooter.vue";
import firstErrorLetter from "@/assets/images/logo/404page_4_1.svg";
import middleErrorLetter from "@/assets/images/logo/canopas-icon.svg";
import lastErrorLetter from "@/assets/images/logo/404page_4_2.svg";
import config from "@/config.js";

export default {
  setup() {
    const seoData = config.SEO_META_DATA;
    useSeoMeta({
      title: seoData.title,
      description: seoData.description,
      ogTitle: seoData.title,
      ogType: seoData.type,
      ogUrl: seoData.url,
      ogImage: seoData.image,
    });
  },
  data() {
    return {
      firstErrorLetter: firstErrorLetter,
      middleErrorLetter: middleErrorLetter,
      lastErrorLetter: lastErrorLetter,
    };
  },
  components: {
    Header,
    NewFooter,
  },
  inject: ["mixpanel"],
  mounted() {
    this.$mixpanel.track("view_404_page");
  },
};
</script>
