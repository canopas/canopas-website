<template>
  <div>
    <ScreenHeaderV2 />
    <div class="container">
      <div class="image-404">
        <img :src="firstErrorLetter" alt="404" />
        <img :src="middleErrorLetter" class="middle-logo" alt="404" />
        <img :src="lastErrorLetter" alt="404" />
      </div>
      <div class="middle-text">
        The page you’re looking for was moved, renamed or might never existed.
      </div>
      <a
        type="button"
        class="back-btn gradient-border-btn"
        onclick="window.location.replace('/')"
      >
        <font-awesome-icon class="arrow" icon="arrow-left" id="leftArrow" />
        <span class="canopas-gradient-text">Back to Home page</span>
      </a>
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

<style lang="scss" scoped>
.container {
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  margin: 0 auto;
  padding: 150px 5%;
  min-height: 50vh;
}

.image-404 {
  margin: 0 auto;
  text-align: center;
  display: flex;
  flex-direction: row;
}

.middle-text {
  margin-top: 2rem;
  text-align: center;
  font-size: 1.4rem !important;
  line-height: 2rem !important;
}

.middle-logo {
  margin: 1rem;
  width: 6.5rem;
}

.back-btn,
.back-btn:hover,
.back-btn:focus {
  width: auto;
  padding: 1rem 1.5rem !important;
  margin: 7% auto;
  border: 1px solid transparent;
  border-radius: 10px;
  display: flex;
  justify-content: space-around !important;
  align-items: center;
  text-decoration: none;
}

.back-btn > span {
  text-align: center !important;
  font-weight: 700;
  font-size: 1.1rem !important;
}

.black {
  color: rgba(61, 61, 61, 1);
}

.back-btn > i {
  margin-right: 1rem;
}

.arrow {
  color: #f2709c;
  font-size: 1.5rem;
  margin-right: 16px;
}
</style>
