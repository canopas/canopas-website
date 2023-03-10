<template>
  <div class="tw-w-full tw-h-[550px] tw-border-0 tw-pt-[20px] md:tw-h-[700px]">
    <div v-if="isLoading" class="iframe-loader">
      <img
        :src="loader"
        class="tw-fixed tw-left-[50%] tw-top-[50%] -tw-translate-y-1/2 -tw-translate-x-1/2 tw-z-[100]"
        alt="loader-image"
      />
    </div>
    <iframe
      class="tw-h-full tw-w-full"
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
