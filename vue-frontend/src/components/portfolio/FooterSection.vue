<template>
  <section class="background-image tw-relative tw-z-[-1] tw-bg-white">
    <aspect-ratio :height="isMobile ? '100%' : '56.26%'">
      <img
        :src="response.backgroundImage[3]"
        :srcset="`${response.backgroundImage[0]} 400w, ${response.backgroundImage[1]} 800w, ${response.backgroundImage[2]} 1400w, ${response.backgroundImage[3]} 2400w`"
        class="tw-w-full tw-h-full tw-object-cover"
        :alt="response.alt"
      />
    </aspect-ratio>
  </section>

  <section
    class="v2-header-2-text tw-font-bold tw-text-center ... tw-bg-white tw-py-20 sm:tw-py-40 lg:tw-py-80"
  >
    <a
      :href="response.url"
      class="animation-underline tw-inline-block tw-relative hover:tw-text-[#3d3d3d] after:tw-absolute after:tw-w-full after:tw-scale-x-0 after:tw-h-0.5 after:tw-bottom-0 after:tw-left-0 after:tw-bg-black-900 after:tw-origin-bottom-left after:tw-duration-300 hover:after:tw-scale-x-100 hover:after:tw-origin-bottom-left"
    >
      {{ response.title }}
    </a>
  </section>
</template>

<script>
import AspectRatio from "@/components/utils/AspectRatio.vue";
export default {
  props: ["response"],
  data() {
    return {
      isMobile: false,
    };
  },
  components: {
    AspectRatio,
  },
  mounted() {
    if (window.innerWidth < 768) {
      this.isMobile = true;
    }
  },
};
</script>

<style lang="scss" scoped>
.animation-underline:after {
  content: "";
}

@include media-breakpoint-up(lg) {
  section.background-image {
    transform: translateZ(-1px) scale(1.5);
  }
}

@media (min-width: 3840px) {
  section.background-image {
    transform: translateZ(-1px) scale(1.3);
  }
}
</style>
