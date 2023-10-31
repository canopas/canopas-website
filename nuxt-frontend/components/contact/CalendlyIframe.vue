<template>
  <div class="w-full h-[550px] border-0 pt-5 md:h-screen">
    <div v-if="isLoading" class="iframe-loader">
      <img
        :src="loader"
        class="fixed left-[50%] top-[50%] -translate-y-1/2 -translate-x-1/2 z-[100]"
        alt="loader-image"
      />
    </div>
    <iframe
      class="h-screen w-full"
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
