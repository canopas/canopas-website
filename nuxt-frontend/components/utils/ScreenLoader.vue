<template>
  <div
    :class="
      loader ? 'fixed z-[1] top-0 left-0 w-full h-full bg-[#00000080]' : ''
    "
    v-if="isLoad"
  >
    <img
      class="fixed left-[50%] top-[50%] z-[100] -translate-x-1/2 -translate-y-1/2"
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
