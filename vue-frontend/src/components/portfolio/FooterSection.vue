<template>
  <section class="background-image tw-relative tw-z-[-1] tw-bg-white">
    <aspect-ratio height="66.26% ">
      <img
        v-if="response.responsiveImages"
        :src="response.backgroundImage[3]"
        :srcset="`${response.backgroundImage[0]} 400w, ${response.backgroundImage[1]} 800w, ${response.backgroundImage[2]} 1200w, ${response.backgroundImage[3]} 1600w`"
        sizes="(min-width: 992px) 45vw, 100vw"
        class="tw-w-full tw-h-full tw-object-cover"
      />
      <img
        v-else
        :src="response.backgroundImage"
        class="tw-w-full tw-h-full tw-object-cover"
      />
    </aspect-ratio>
  </section>

  <div
    class="v2-header-2-text tw-font-bold tw-text-center ... tw-bg-white tw-py-32 sm:tw-py-40 md:tw-py-44 2xl:tw-py-60"
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
