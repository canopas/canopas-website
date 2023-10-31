<template>
  <div class="vue-aspect-ratio relative" :style="outerStyle">
    <div class="absolute inset-0">
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
