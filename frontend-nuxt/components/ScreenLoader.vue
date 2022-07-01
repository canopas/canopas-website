<template>
  <div :class="loader ? 'loader-div' : ''" v-if="isLoad">
    <img class="loader-image" :src="loaderImage" />
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

<style scoped lang="scss">
.loader-div {
  position: fixed;
  z-index: 1;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: table;
  transition: opacity 0.3s ease;
}

.loader-image {
  position: fixed;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  z-index: 100;
}
</style>
