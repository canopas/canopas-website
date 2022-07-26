<template>
  <section class="tw-bg-white tw-relative">
    <div class="tw-relative container">
      <div>
        <div
          :class="[
            id == 'togness'
              ? 'tw-pb-20 md:tw-pt-40 '
              : 'tw-py-20 sm:tw-py-40 lg:tw-py-80 tw-flex tw-flex-col tw-justify-between xl:tw-w-3/4',
          ]"
        >
          <div
            class="v2-title-text tw-font-bold"
            v-html="response[0].title"
          ></div>
          <div
            class="description tw-pt-5 tw-w-4/5 lg:tw-pt-10 xl:tw-pt-20 xl:tw-w-full"
            :class="[id == 'togness' ? 'tw-hidden' : '']"
          >
            <div class="v2-normal-text tw-font-light">
              {{ response[0].description }}
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>

  <section class="image tw-relative tw-z-[-1]">
    <aspect-ratio height="56.25%" class="tw-overflow-hidden">
      <img
        :src="response[0].backgroundImage[3]"
        :srcset="`${response[0].backgroundImage[0]} 400w, ${response[0].backgroundImage[1]} 800w,${response[0].backgroundImage[2]} 1400w,${response[0].backgroundImage[3]} 2400w`"
        :alt="response.alt"
        class="tw-w-full tw-h-full tw-object-cover"
      />
    </aspect-ratio>
    <img
      v-if="response[0].gif"
      :src="response[0].gif"
      class="tw-absolute tw-inset-0 tw-left-auto tw-h-2/4 tw-w-2/6 tw-object-cover cycling-animation"
      alt="cycling-animation"
    />
  </section>

  <section
    class="tw-bg-white tw-relative"
    :class="[id == 'togness' ? '' : 'tw-hidden']"
  >
    <div class="tw-relative container">
      <div class="tw-pt-40 tw-pt-20 tw-flex tw-flex-col tw-justify-between">
        <div class="v2-title-2-text xl:tw-w-1/3">
          {{ response[0].technology.title }}
        </div>
      </div>
      <div class="tw-flex tw-justify-between tw-pt-16">
        <div
          class="description"
          v-for="technology in response[0].technology.details"
          :key="technology"
        >
          <div
            class="v2-title-2-text tw-font-bold tw-border tw-border-inherit tw-divide-x tw-pr-5 tw-pl-5"
          >
            {{ technology.title }}
          </div>
        </div>
      </div>
    </div>
  </section>

  <section
    class="tw-bg-white tw-relative"
    v-if="response[1]"
    :class="[id == 'togness' ? 'tw-hidden' : '']"
  >
    <div class="tw-relative">
      <div
        class="container tw-py-20 sm:tw-py-40 lg:tw-py-80 tw-flex tw-flex-col tw-justify-between xl:tw-flex-row"
      >
        <div class="v2-title-2-text tw-font-bold">{{ response[1].title }}</div>
        <div class="description lg:tw-w-11/12 xl:tw-pl-16">
          <div class="v2-normal-text tw-font-light">
            {{ response[1].description }}
          </div>
        </div>
      </div>
    </div>
  </section>

  <section
    class="image tw-relative tw-px-4 lg:tw-px-12 tw-z-[-1]"
    v-if="response[1]"
  >
    <aspect-ratio height="56.25%" class="tw-overflow-hidden">
      <img
        :src="response[1].backgroundImage[3]"
        :srcset="`${response[1].backgroundImage[0]} 400w, ${response[1].backgroundImage[1]} 800w,${response[1].backgroundImage[2]} 1400w,${response[1].backgroundImage[3]} 2400w`"
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
  data() {
    return {
      id: this.$route.params.id,
    };
  },
  components: {
    AspectRatio,
  },
};
</script>

<style lang="postcss" scoped>
section.image {
  transform: translateZ(-1px) scale(1.5);
}

.cycling-animation {
  transform: matrix(-1, 0, 0, 1, 0, 0);
}
</style>
