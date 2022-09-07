<template>
  <div>
    <metainfo>
      <template v-slot:title="{ content }">
        {{ content }}
      </template>
    </metainfo>
    <ScreenHeaderV2 />
    <div
      class="tw-container tw-flex tw-items-center tw-justify-center tw-flex-col tw-my-0 tw-mx-auto tw-py-[150px] tw-px-[5%] tw-min-h-[50vh]"
    >
      <div class="tw-flex tw-flex-row tw-my-0 tw-mx-auto tw-text-center">
        <img :src="firstErrorLetter" alt="404" />
        <img :src="middleErrorLetter" class="tw-m-4 tw-w-[6.5rem]" alt="404" />
        <img :src="lastErrorLetter" alt="404" />
      </div>
      <div class="tw-mt-[2rem] tw-text-center tw-text-[1.4rem] tw-leading-8">
        The page you’re looking for was moved, renamed or might never existed.
      </div>
      <router-link
        class="tw-w-auto tw-py-[1rem] tw-px-[1.5rem] tw-my-[7%] tw-mx-auto tw-border-[1rem] tw-border-solid tw-border-transparent tw-rounded-[10px] tw-flex tw-justify-around tw-items-center gradient-border-btn"
        to="/"
      >
        <font-awesome-icon
          class="fa tw-text-[#f2709c] tw-text-[1.5rem] tw-mr-[16px]"
          icon="arrow-left"
          id="leftArrow"
        />
        <span
          class="canopas-gradient-text tw-text-center tw-font-bold tw-text-[1.1rem]"
          >Back to Home page</span
        >
      </router-link>
    </div>
    <ScreenFooter2 />
  </div>
</template>

<script>
import ScreenHeaderV2 from "@/components/partials/ScreenHeaderV2.vue";
import ScreenFooter2 from "@/components/partials/ScreenFooter2.vue";
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
    ScreenHeaderV2,
    ScreenFooter2,
    FontAwesomeIcon,
  },
  mounted() {
    this.$gtag.event("view_page_404");
  },
};
</script>
