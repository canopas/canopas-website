<template>
  <div>
    <HomeIndex v-if="!isShowNewHomePage" />
    <HomeNewIndex v-else />
  </div>
</template>

<script>
import config from "@/config.js";
import HomeIndex from "@/components/home/HomeIndex.vue";
import HomeNewIndex from "@/components/home-new/HomeIndex.vue";

export default {
  data() {
    return {
      isShowNewHomePage: config.IS_SHOW_NEW_HOME_PAGE,
    };
  },
  components: {
    HomeIndex,
    HomeNewIndex,
  },
  mounted() {
    this.$gtag.event("view_page_home");
  },
};
</script>

<style scoped></style>
