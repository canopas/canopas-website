<template>
  <section class="tw-bg-white tw-relative">
    <div class="tw-relative container">
      <div>
        <div
          class="tw-pt-20 tw-pb-14 sm:tw-pt-40 sm:tw-pb-24 lg:tw-pt-80 lg:tw-pb-44"
        >
          <div
            class="v2-header-2-text tw-font-bold"
            v-html="response.title"
          ></div>
        </div>
      </div>
    </div>
  </section>

  <section
    class="tw-relative tw-flex tw-flex-col sm:tw-flex-row tw-bg-white tw-pb-8 sm:tw-pb-20 xl:tw-pb-28"
  >
    <div class="container tw-flex tw-flex-col sm:tw-flex-row">
      <div>
        <div v-for="data in flex1" :key="data" class="tw-p-3">
          <aspect-ratio
            :height="data.aspectRatio"
            :style="[data.background ? { background: data.background } : {}]"
          >
            <img
              class="tw-h-full tw-w-full tw-object-cover"
              :src="data.image"
              v-if="data.image"
            />
            <video
              class="tw-h-full tw-w-full tw-object-cover tw-px-20 tw-pt-16 sm:tw-px-8 sm:tw-pt-10 lg:tw-pt-20 lg:tw-px-24"
              preload="auto"
              loop
              muted
              autoplay
              playsinline
              v-else
            >
              <source :src="data.video" type="video/mp4" />
            </video>
          </aspect-ratio>
          <div
            v-if="data.title"
            class="v2-normal-3-text tw-bg-white tw-relative tw-px-6 tw-pt-3 tw-pb-6 lg:tw-px-12 lg:tw-pt-6 lg:tw-pb-12 xl:tw-px-20 xl:tw-py-12"
          >
            {{ data.title }}
          </div>
        </div>
      </div>
      <div class="sm:tw-mt-36 lg:tw-mt-60">
        <div v-for="data in flex2" :key="data" class="tw-p-3">
          <aspect-ratio
            :height="data.aspectRatio"
            :style="[data.background ? { background: data.background } : {}]"
          >
            <img
              class="tw-h-full tw-w-full tw-object-cover"
              :src="data.image"
              v-if="data.image"
            />
            <video
              class="tw-h-full tw-w-full tw-object-cover tw-px-20 tw-pt-16 sm:tw-px-8 sm:tw-pt-10 lg:tw-pt-20 lg:tw-px-24"
              preload="auto"
              loop
              muted
              autoplay
              playsinline
              v-else
            >
              <source :src="data.video" type="video/mp4" />
            </video>
          </aspect-ratio>
          <div
            v-if="data.title"
            class="v2-normal-3-text tw-px-6 tw-pt-3 tw-pb-6 lg:tw-px-12 lg:tw-pt-6 lg:tw-pb-12 xl:tw-px-20 xl:tw-py-12"
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
