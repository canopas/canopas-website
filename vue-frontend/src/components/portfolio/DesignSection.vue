<template>
  <section class="tw-bg-white tw-relative">
    <div class="tw-relative container">
      <div>
        <div
          class="flex-div tw-flex tw-flex-col tw-justify-between xl:tw-w-3/4"
        >
          <div
            class="v2-header-2-text tw-font-bold"
            v-html="response[0].title"
          ></div>
          <div
            class="description tw-pt-5 tw-w-4/5 lg:tw-pt-10 xl:tw-pt-20 xl:tw-w-full"
          >
            <div class="v2-normal-text tw-font-light">
              {{ response[0].description }}
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>

  <section class="image tw-relative">
    <aspect-ratio height="56.25%" class="tw-overflow-hidden">
      <img
        :src="response.backgroundImage[3]"
        :srcset="`${response.backgroundImage[0]} 400w, ${response.backgroundImage[1]} 800w, ${response.backgroundImage[2]} 1200w, ${response.backgroundImage[3]} 1600w`"
        sizes="(min-width: 992px) 45vw, 100vw"
        loading="lazy"
      />
    </aspect-ratio>
  </section>

  <section class="tw-bg-white tw-relative">
    <div class="tw-relative">
      <div
        class="container flex-div tw-flex tw-flex-col tw-justify-between lg:tw-flex-row"
      >
        <div class="v2-title-2-text tw-font-bold">{{ response[1].title }}</div>
        <div class="description tw-w-4/5">
          <div class="v2-normal-text tw-font-light">
            {{ response[1].description }}
          </div>
        </div>
      </div>
    </div>
  </section>

  <section class="image tw-relative tw-px-4 lg:tw-px-12">
    <aspect-ratio height="56.25%" class="tw-overflow-hidden">
      <img
        :src="response.backgroundImage[3]"
        :srcset="`${response.backgroundImage[0]} 400w, ${response.backgroundImage[1]} 800w, ${response.backgroundImage[2]} 1200w, ${response.backgroundImage[3]} 1600w`"
        sizes="(min-width: 992px) 45vw, 100vw"
        loading="lazy"
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
.flex-div {
  @apply tw-py-20 sm:tw-py-40 lg:tw-py-80;
}

section.image {
  transform: translateZ(-1px) scale(1.5);
  z-index: -1;
}

img {
  @apply tw-w-full tw-h-full tw-object-cover;
}
</style>
