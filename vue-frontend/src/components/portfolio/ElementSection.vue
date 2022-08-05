<template>
  <section class="tw-bg-white tw-relative">
    <div class="tw-relative container">
      <div>
        <div class="tw-py-20 sm:tw-py-40 lg:tw-py-80">
          <div
            :class="
              response.class ? response.class : 'v2-title-2-text tw-font-bold'
            "
            v-html="response.title"
          ></div>
        </div>
      </div>
    </div>
  </section>

  <section
    v-if="response.detail"
    class="tw-relative tw-bg-white tw-pb-20 sm:tw-pb-40 lg:tw-pb-80"
  >
    <div class="container tw-flex tw-flex-col sm:tw-flex-row tw-relative">
      <div class="tw-basis-1/2">
        <div v-for="data in flex1" :key="data" class="tw-p-3">
          <aspect-ratio
            :height="data.aspectRatio"
            :style="[data.background ? { background: data.background } : {}]"
          >
            <img
              v-if="data.image"
              class="tw-w-full tw-h-full tw-object-cover"
              :src="data.image[3]"
              :srcset="`${data.image[0]} 400w, ${data.image[1]} 800w, ${data.image[2]} 1200w, ${data.image[3]} 1600w`"
              :alt="data.alt"
            />
            <img
              v-if="data.animation"
              :src="data.animation"
              class="tw-absolute tw-inset-0 tw-m-auto tw-object-cover"
              :alt="data.alt"
            />
            <video
              v-else
              class="tw-h-full tw-w-full tw-object-cover tw-px-20 tw-pt-16 sm:tw-px-8 sm:tw-pt-10 lg:tw-pt-20 lg:tw-px-24"
              preload="auto"
              loop
              muted
              autoplay
              playsinline
            >
              <source :src="data.video" type="video/mp4" />
            </video>
          </aspect-ratio>
          <div
            v-if="data.title"
            class="v2-normal-text tw-px-6 tw-py-6 lg:tw-px-12 lg:tw-py-12 xl:tw-p-20"
          >
            {{ data.title }}
          </div>
        </div>
      </div>
      <div class="sm:tw-mt-36 lg:tw-mt-60 tw-basis-1/2">
        <div v-for="data in flex2" :key="data" class="tw-p-3">
          <aspect-ratio
            :height="data.aspectRatio"
            :style="[data.background ? { background: data.background } : {}]"
          >
            <img
              v-if="data.image"
              class="tw-w-full tw-h-full tw-object-cover"
              :src="data.image[3]"
              :srcset="`${data.image[0]} 400w, ${data.image[1]} 800w, ${data.image[2]} 1200w, ${data.image[3]} 1600w`"
              :alt="data.alt"
            />
            <img
              v-if="data.animation"
              :src="data.animation"
              class="tw-absolute tw-inset-0 tw-m-auto tw-object-cover tw-rounded-3xl tw-h-6/6"
              :alt="data.alt"
            />
            <video
              v-else
              class="tw-h-full tw-w-full tw-object-cover tw-px-20 tw-pt-16 sm:tw-px-8 sm:tw-pt-10 lg:tw-pt-20 lg:tw-px-24"
              preload="auto"
              loop
              muted
              autoplay
              playsinline
            >
              <source :src="data.video" type="video/mp4" />
            </video>
          </aspect-ratio>
          <div
            v-if="data.title"
            class="v2-normal-text tw-px-10 tw-py-6 sm:tw-px-7 lg:tw-px-12 lg:tw-py-12 xl:tw-px-20 xl:tw-py-12"
          >
            {{ data.title }}
          </div>
        </div>
      </div>
    </div>
  </section>

  <section v-if="response.subTitle" class="tw-bg-white">
    <div
      class="container tw-flex tw-flex-col tw-justify-between tw-py-20 sm:tw-py-40 lg:tw-flex-row"
    >
      <div class="v2-normal-text tw-font-bold">
        {{ response.subTitle }}
      </div>
      <div class="tw-pt-5 lg:tw-pl-16 lg:tw-w-4/5 lg:tw-pt-0">
        <div class="v2-normal-text tw-font-light">
          {{ response.description }}
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import AspectRatio from "@/components/utils/AspectRatio.vue";

export default {
  props: ["response"],
  data() {
    return {
      flex1: this.response.detail.firstDetail,
      flex2: this.response.detail.secondDetail,
    };
  },
  components: {
    AspectRatio,
  },
};
</script>
