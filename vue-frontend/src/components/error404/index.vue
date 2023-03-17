<template>
  <div>
    <metainfo>
      <template v-slot:title="{ content }">
        {{ content }}
      </template>
    </metainfo>
    <Header />
    <div
      class="tw-container tw-flex tw-items-center tw-justify-center tw-flex-col tw-my-0 tw-mx-auto tw-py-[150px] tw-px-[5%] tw-min-h-[50vh]"
    >
      <div class="tw-flex tw-flex-row tw-my-0 tw-mx-auto tw-text-center">
        <img :src="firstErrorLetter" alt="404" />
        <img :src="middleErrorLetter" class="tw-m-4 tw-w-[6.5rem]" alt="404" />
        <img :src="lastErrorLetter" alt="404" />
      </div>
      <h1 class="tw-mt-[2rem] tw-text-center tw-text-[1.4rem] tw-leading-8">
        The page you’re looking for was moved, renamed or might never existed.
      </h1>
      <router-link
        class="tw-my-16 gradient-border-btn tw-flex tw-flex-row tw-items-center"
        to="/"
      >
        <font-awesome-icon
          class="arrow tw-w-6 tw-h-6 fa tw-text-[#f2709c] tw-text-[2.1rem] tw-mr-[5px]"
          icon="arrow-left"
          id="leftArrow"
        />
        <span class="tw-text-center tw-font-bold tw-text-[1.1rem]"
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
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import Config from "@/config.js";
import { useMeta } from "vue-meta";

export default {
  setup() {
    var seoData = Config.SEO_META_DATA;
    useMeta({
      title: seoData.title,
      description: seoData.description,
      og: {
        type: seoData.type,
        title: seoData.title,
        url: seoData.url,
        image: seoData.image,
      },
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
    FontAwesomeIcon,
  },
  inject: ["mixpanel"],
  mounted() {
    this.mixpanel.track("view_404_page");
  },
};
</script>
