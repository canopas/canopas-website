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

<style lang="scss" scoped>
.calendly-iframe-div {
  width: 100%;
  height: 550px;
  border: none;
  padding-top: 20px;
}

.loader-div {
  position: fixed;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  z-index: 100;
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

@include media-breakpoint-up(md) {
  .calendly-iframe-div {
    height: 700px;
  }
}
</style>
