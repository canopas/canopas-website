<template>
  <section class="tw-bg-white">
    <div class="container tw-py-20 sm:tw-pb-40 lg:tw-pb-80 lg:tw-pt-60">
      <div class="v2-normal-2-text animate__animated animate__fadeInUp">
        {{ response.title }}
      </div>
      <div class="v2-title-text tw-mt-10 animate__animated animate__slideInUp">
        {{ response.subTitle }}
      </div>
    </div>
  </section>
  <section class="background-image tw-relative tw-z-[-1]">
    <aspect-ratio height="56.26%">
      <img
        v-if="response.responsiveImages"
        :src="response.backgroundImage[3]"
        :srcset="`${response.backgroundImage[0]} 400w, ${response.backgroundImage[1]} 800w, ${response.backgroundImage[2]} 1200w, ${response.backgroundImage[3]} 1600w`"
        sizes="(min-width: 992px) 45vw, 100vw"
        class="tw-w-full tw-h-full tw-object-cover"
        :alt="response.alt"
      />
      <img
        v-else
        :src="response.backgroundImage"
        class="tw-w-full tw-h-full tw-object-cover"
        :alt="response.alt"
      />
    </aspect-ratio>
  </section>
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
</style>
