<template>
  <section class="background-image tw-relative tw-z-[-1] tw-bg-white">
    <aspect-ratio height="56.26% ">
      <img
        :src="response.backgroundImage[3]"
        :srcset="`${response.backgroundImage[0]} 400w, ${response.backgroundImage[1]} 800w, ${response.backgroundImage[2]} 1400w, ${response.backgroundImage[3]} 2400w`"
        class="tw-w-full tw-h-full tw-object-cover"
        :alt="response.alt"
      />
    </aspect-ratio>
  </section>

  <div
    class="v2-header-2-text tw-font-bold tw-text-center ... tw-bg-white tw-py-24 sm:tw-py-32 lg:tw-py-40"
  >
    <a
      href="/"
      class="animation-underline tw-inline-block tw-relative hover:tw-text-[#3d3d3d] after:tw-absolute after:tw-w-full after:tw-scale-x-0 after:tw-h-0.5 after:tw-bottom-0 after:tw-left-0 after:tw-bg-black-900 after:tw-origin-bottom-left after:tw-duration-300 hover:after:tw-scale-x-100 hover:after:tw-origin-bottom-left"
    >
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
}

.animation-underline:after {
  content: "";
}
</style>
