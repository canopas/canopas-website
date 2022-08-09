<template>
  <section class="tw-bg-white">
    <div
      v-if="response.solution"
      class="container tw-py-20 sm:tw-py-40 lg:tw-py-80 tw-relative"
    >
      <div class="tw-flex tw-flex-col tw-justify-between lg:tw-flex-row">
        <div class="v2-normal-text tw-font-bold">
          {{ response.solution.title }}
        </div>
        <div class="tw-pt-5 lg:tw-pl-16 lg:tw-w-4/5 lg:tw-pt-0">
          <div
            class="v2-normal-text tw-font-light"
            v-html="response.solution.description"
          ></div>
          <div class="tw-pt-16">
            <a
              v-for="button in response.solution.buttons"
              :key="button"
              class="tw-pr-10 is-animation-tab tw-inline-block tw-relative v2-normal-2-text hover:tw-text-black-900 after:tw-absolute after:tw-w-1/4 after:tw-scale-x-0 after:tw-h-0.5 after:tw-bottom-0 after:tw-left-0 after:tw-bg-black-900 after:tw-origin-bottom-left after:tw-duration-300 hover:after:tw-scale-x-100 hover:after:tw-origin-bottom-left"
              :href="button.link"
              target="_blank"
            >
              {{ button.name }}
            </a>
          </div>
        </div>
      </div>
    </div>
    <div
      v-if="response.title"
      class="container tw-py-20 sm:tw-py-40 lg:tw-py-80 tw-relative"
    >
      <div class="v2-header-3-text tw-pt-20" v-html="response.title"></div>
    </div>
  </section>

  <section
    v-if="response.backgroundImage[3]"
    class="background-image tw-relative tw-z-[-1]"
  >
    <aspect-ratio height="56.26%">
      <img
        :src="response.backgroundImage[3]"
        :srcset="`${response.backgroundImage[0]} 400w, ${response.backgroundImage[1]} 800w, ${response.backgroundImage[2]} 1400, ${response.backgroundImage[3]} 2400w`"
        class="tw-w-full tw-h-full tw-object-cover"
        :alt="response.alt"
      />
    </aspect-ratio>
  </section>

  <section class="tw-bg-white">
    <div
      class="container tw-pt-24 md:tw-pt-48 xl:tw-pt-80 tw-flex tw-flex-col ... md:tw-flex-row ... md:tw-gap-x-16"
    >
      <div class="tw-basis-1/2">
        <div v-for="data in flex1" :key="data">
          <div class="tw-relative">
            <aspect-ratio
              :height="data.aspectRatio"
              :style="[data.background ? { background: data.background } : {}]"
            >
              <div
                :class="[
                  data.video
                    ? 'tw-px-4 tw-pb-4 md:tw-px-2 lg:tw-px-8 lg:tw-pb-8'
                    : '',
                ]"
              >
                <img
                  v-if="data.image"
                  :src="data.image[3]"
                  :srcset="`${data.image[0]} 400w, ${data.image[1]} 800w, ${data.image[2]} 1200w, ${data.image[3]} 1600w`"
                  class="tw-w-full tw-h-full tw-object-cover"
                  :alt="data.alt"
                />
                <img
                  v-if="data.gif"
                  :src="data.gif"
                  class="tw-absolute tw-inset-0 tw-m-auto tw-rounded-3xl tw-object-cover tw-h-3/3"
                  :alt="data.alt"
                />
                <video
                  v-if="data.video"
                  id="video-preview"
                  controls
                  playsinline
                  loop
                  muted
                  autoplay
                  :class="[
                    data.video
                      ? 'tw-rounded-b-[90px] tw-border-b-8 tw-border-x-8 tw-border-solid tw-border-white tw-px-8 tw-pb-8'
                      : '',
                  ]"
                >
                  <source :src="data.video" type="video/mp4" />
                </video>
              </div>
            </aspect-ratio>
          </div>
          <div
            class="v2-normal-text tw-bg-white tw-font-light ..."
            :class="[
              data.title
                ? 'tw-py-8  sm:tw-px-32 sm:tw-py-8 md:tw-px-14 xl:tw-p-16  2xl:tw-px-32 tw-bg-white tw-text-center'
                : 'tw-mt-12',
            ]"
          >
            {{ data.title }}
          </div>
        </div>
      </div>
      <div class="tw-basis-1/2 md:tw-mt-36 lg:tw-mt-64">
        <div v-for="data in flex2" :key="data">
          <div class="tw-relative">
            <aspect-ratio
              :height="data.aspectRatio"
              :style="[data.background ? { background: data.background } : {}]"
            >
              <div
                :class="[
                  data.video
                    ? 'tw-px-4 tw-pb-4 md:tw-px-2 lg:tw-px-8 lg:tw-pb-8'
                    : '',
                ]"
              >
                <img
                  v-if="data.image"
                  class="tw-w-full tw-h-full tw-object-cover"
                  :src="data.image[3]"
                  :srcset="`${data.image[0]} 400w, ${data.image[1]} 800w, ${data.image[2]} 1200w, ${data.image[3]} 1600w`"
                  :alt="data.alt"
                />
                <img
                  v-if="data.gif"
                  :src="data.gif"
                  class="tw-absolute tw-inset-0 tw-m-auto tw-rounded-3xl tw-object-cover tw-h-3/3"
                  :alt="data.alt"
                />
                <LottieAnimation
                  v-else-if="data.animation"
                  :jsonData="data.animation"
                  class="tw-mt-14 lg:tw-mt-20"
                />
                <video
                  v-if="data.video"
                  preload="auto"
                  loop
                  muted
                  autoplay
                  playsinline
                  id="video-preview"
                  :class="[
                    data.video
                      ? 'tw-rounded-b-[90px] tw-border-b-8 tw-border-x-8 tw-border-solid tw-border-white tw-px-8 tw-pb-8'
                      : '',
                  ]"
                >
                  <source :src="data.video" type="video/mp4" />
                </video>
              </div>
            </aspect-ratio>
          </div>
          <div
            class="v2-normal-text tw-bg-white tw-font-light ..."
            :class="[
              data.title
                ? 'tw-py-8  sm:tw-px-32 sm:tw-py-8 md:tw-px-14 xl:tw-p-16  2xl:tw-px-32 tw-bg-white tw-text-center'
                : 'tw-mt-12',
            ]"
          >
            {{ data.title }}
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
<script>
import AspectRatio from "@/components/utils/AspectRatio.vue";
import LottieAnimation from "@/components/utils/LottieAnimation.vue";

export default {
  props: ["response"],
  data() {
    return {
      flex1: this.response.details.firstDetail,
      flex2: this.response.details.secondDetail,
    };
  },
  components: {
    AspectRatio,
    LottieAnimation,
  },
};
</script>
<style lang="scss" scoped>
section.background-image {
  transform: translateZ(-1px) scale(1.5);
}

@media (min-width: 3840px) {
  section.background-image {
    transform: translateZ(-1px) scale(1.3);
  }
}
</style>
