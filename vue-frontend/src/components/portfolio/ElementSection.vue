<template>
  <section class="tw-bg-white tw-relative">
    <div class="tw-relative tw-container">
      <div v-if="response.title">
        <div
          class="lg:tw-pt-40 tw-pb-[3rem] lg:tw-pb-80 tw-pt-0 md:tw-pt-[7rem]"
        >
          <div
            :class="
              response.class
                ? response.class
                : 'v2-header-3-text tw-font-medium'
            "
            v-html="response.title"
          ></div>
        </div>
      </div>
    </div>
  </section>

  <section v-if="response.detail" class="tw-relative tw-bg-white lg:tw-pb-80">
    <div class="tw-container tw-flex tw-flex-col sm:tw-flex-row tw-relative">
      <div v-if="gridData1" class="tw-basis-1/2">
        <div v-for="data in gridData1" :key="data" class="tw-px-3">
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
              class="tw-h-full tw-w-2/3 tw-object-cover tw-m-auto tw-rounded-t-lg tw-pt-16 sm:tw-pt-10 lg:tw-pt-20"
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
            :class="[
              data.title
                ? 'v2-normal-text tw-px-6 tw-py-12 lg:tw-px-12 lg:tw-py-12 xl:tw-p-20 tw-text-center'
                : 'tw-pt-6',
            ]"
          >
            {{ data.title }}
          </div>
        </div>
      </div>
      <div v-if="gridData2" class="sm:tw-mt-36 lg:tw-mt-60 tw-basis-1/2">
        <div v-for="data in gridData2" :key="data" class="tw-px-3">
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
            :class="[
              data.title
                ? 'v2-normal-text tw-px-6 tw-py-12 lg:tw-px-12 lg:tw-py-12 xl:tw-p-20 tw-text-center'
                : 'tw-pt-6',
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

export default {
  props: ["json"],
  data() {
    return {
      response: this.json,
      gridData1: this.json.detail.gridData1,
      gridData2: this.json.detail.gridData2,
    };
  },
  watch: {
    json: function (newVal, oldVal) {
      this.response = newVal;
      this.gridData1 = this.response.detail.gridData1;
      this.gridData2 = this.response.detail.gridData2;
    },
  },
  components: {
    AspectRatio,
  },
};
</script>
