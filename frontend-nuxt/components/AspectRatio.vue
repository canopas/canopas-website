<template>
  <div class="vue-aspect-ratio" :style="outerStyle">
    <div class="vue-aspect-ratio-content">
      <slot />
    </div>
  </div>
</template>

<script>
export default {
  props: ["height"],
  computed: {
    outerStyle() {
      return {
        paddingTop: this.height,
      };
    },
  },
};
</script>

<style scoped>
.vue-aspect-ratio {
  position: relative;
}

.vue-aspect-ratio-content {
  position: absolute;
  top: 0;
  left: 0;
  bottom: 0;
  right: 0;
}
</style>
