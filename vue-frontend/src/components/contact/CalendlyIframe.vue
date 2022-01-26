<template>
  <div class="calendly-iframe-div">
    <div v-if="isLoading" class="iframe-loader">
      <img :src="loader" class="loader-div" />
    </div>
    <iframe
      class="calendly-iframe"
      :src="calendlyUrl"
      title="calendly"
      v-on:load="isLoading = false"
      v-show="!isLoading"
    ></iframe>
  </div>
</template>

<script type="module">
import config from "@/config.js";
import loader from "@/assets/images/theme/loader.svg";

export default {
  data() {
    return {
      calendlyUrl: config.CALENDLY_IFRAME_URL,
      loader: loader,
      isLoading: true,
    };
  },
};
</script>

<style scoped>
.calendly-iframe-div {
  height: 500px;
  border: none;
}

.iframe-loader {
  display: flex;
  justify-content: center;
}

.iframe-div {
  height: 100%;
  width: 100%;
  border: none !important;
}

.calendly-iframe {
  height: 100%;
  width: 100%;
}

.calendly-iframe::-webkit-scrollbar {
  display: none;
}
</style>
