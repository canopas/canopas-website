<template>
  <div
    :class="
      loader
        ? 'tw-fixed tw-z-[1] tw-top-0 tw-left-0 tw-w-full tw-h-full tw-bg-[#00000080]'
        : ''
    "
    v-if="isLoad"
  >
    <img
      class="tw-fixed tw-left-[50%] tw-top-[50%] tw-z-[100] -tw-translate-x-1/2 -tw-translate-y-1/2"
      :src="loaderImage"
      alt="loader-image"
    />
  </div>
</template>

<script>
import loaderImage from "@/assets/images/theme/loader.svg";

export default {
  props: ["loader"],
  data() {
    return {
      isLoad: this.loader,
      loaderImage: loaderImage,
    };
  },
  mounted() {
    setTimeout(() => (this.isLoad = true), 700);
  },
};
</script>
