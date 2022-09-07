<template>
  <section class="tw-bg-white">
    <div
      class="container tw-pt-20 tw-pb-40 sm:tw-pt-32 lg:tw-pb-80 lg:tw-pt-60"
    >
      <div class="v2-normal-2-text animate__animated animate__fadeInUp">
        {{ response.title }}
      </div>
      <div class="v2-title-text tw-mt-10 animate__animated animate__slideInUp">
        {{ response.subTitle }}
      </div>
    </div>
  </section>

  <section class="background-image tw-relative tw-z-[-1]">
    <aspect-ratio :height="isMobile ? '100%' : '56.26%'">
      <img
        :src="response.backgroundImage[3]"
        :srcset="`${response.backgroundImage[0]} 400w, ${response.backgroundImage[1]} 800w, ${response.backgroundImage[2]} 1400w, ${response.backgroundImage[3]} 2400w`"
        class="tw-w-full tw-h-full tw-object-cover"
        :alt="response.alt"
      />
    </aspect-ratio>
  </section>
</template>

<script>
import AspectRatio from "@/components/utils/AspectRatio.vue";

export default {
  props: ["json"],
  data() {
    return {
      isMobile: false,
      response: this.json,
    };
  },
  watch: {
    json: function (newVal, oldVal) {
      this.response = newVal;
    },
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
