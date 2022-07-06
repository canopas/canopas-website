<template>
  <section class="background-image tw-relative tw-bg-white">
    <aspect-ratio height="76.26% ">
      <img
        :src="response.backgroundImage"
        class="tw-w-full tw-h-full tw-object-cover"
      />
    </aspect-ratio>
  </section>

  <div
    class="v2-header-2-text tw-font-bold tw-text-center ... tw-bg-white tw-py-32 sm:tw-py-40 md:tw-py-44 2xl:tw-py-60"
  >
    <a href="/" class="animation-underline hover:tw-text-[#3d3d3d]">
      {{ response.title }}
    </a>
  </div>
</template>
<script>
import AspectRatio from "@/components/utils/AspectRatio.vue";
export default {
  props: ["response"],
  components: {
    AspectRatio,
  },
};
</script>
<style lang="postcss" scoped>
section.background-image {
  transform: translateZ(-1px) scale(1.5);
  z-index: -1;
}
.animation-underline {
  @apply tw-inline-block tw-relative;
}

.animation-underline:after {
  content: "";
}

.animation-underline:after {
  @apply tw-absolute tw-w-full tw-scale-x-0 tw-h-0.5 tw-bottom-0 tw-left-0 tw-bg-black-900 tw-origin-bottom-left tw-duration-300;
}
.animation-underline:hover:after {
  @apply tw-scale-x-100 tw-origin-bottom-left;
}
</style>
