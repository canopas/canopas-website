<template>
  <div class="vue-aspect-ratio tw-relative" :style="outerStyle">
    <div class="tw-absolute tw-inset-0">
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
