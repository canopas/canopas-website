<template>
  <section class="tw-bg-white tw-relative">
    <div class="tw-relative container">
      <img :src="response.backgroundImage" class="background-image" />
      <div class="flex-div">
        <div class="v2-normal-text tw-font-bold">{{ response.title }}</div>
        <div class="description">
          <div class="v2-normal-text tw-font-light">
            {{ response.description }}
          </div>
          <div class="tw-pt-16">
            <a
              v-for="button in response.buttons"
              :key="button"
              class="tw-pr-10 is-animation-tab v2-normal-2-text hover:tw-text-black-900"
              :href="button.link"
            >
              {{ button.name }}
            </a>
          </div>
        </div>
      </div>
    </div>
  </section>

  <section class="container video tw-relative">
    <aspect-ratio height="56.25%" class="tw-overflow-hidden">
      <video class="" preload="auto" loop muted autoplay playsinline>
        <source
          src="https://cdn.cuberto.com/cb/video/sleepiest/1.mp4?2"
          type="video/mp4"
        />
      </video>
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
.background-image {
  @apply tw-hidden tw-absolute tw-object-cover tw-py-40 lg:tw-block;
}

.flex-div {
  @apply tw-flex tw-flex-col tw-justify-between tw-py-20 sm:tw-py-40 lg:tw-flex-row lg:tw-py-80;
}

.description {
  @apply tw-pt-5 lg:tw-pl-16 lg:tw-w-4/5 lg:tw-pt-0;
}

.is-animation-tab {
  @apply tw-inline-block tw-relative;
}

.is-animation-tab:after {
  content: "";
}

.is-animation-tab:after {
  @apply tw-absolute tw-w-1/4 tw-scale-x-0 tw-h-0.5 tw-bottom-0 tw-left-0 tw-bg-black-900 tw-origin-bottom-left tw-duration-300;
}

.is-animation-tab:hover:after {
  @apply tw-scale-x-100 tw-origin-bottom-left;
}

section.video {
  transform: translateZ(-1px) scale(1.5);
  z-index: -1;
}
</style>
