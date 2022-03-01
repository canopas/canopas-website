<template>
  <div class="loader-div" v-if="isLoad">
    <img :src="loaderImage" />
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

<style scoped>
.loader-div {
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  z-index: 100;
}
</style>
