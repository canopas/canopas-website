<template>
  <section class="bg-white relative">
    <div class="relative container">
      <div v-if="response.title">
        <div class="lg:pt-40 pb-12 lg:pb-80 pt-0 md:pt-28">
          <div
            :class="
              response.class ? response.class : 'v2-header-3-text font-medium'
            "
            v-html="response.title"
          ></div>
        </div>
      </div>
    </div>
  </section>

  <section v-if="response.detail" class="relative bg-white lg:pb-80">
    <div class="container flex flex-col sm:flex-row relative">
      <div v-if="gridData1" class="basis-1/2">
        <div v-for="data in gridData1" :key="data" class="px-3">
          <aspect-ratio
            :height="data.aspectRatio"
            :style="[data.background ? { background: data.background } : {}]"
          >
            <img
              v-if="data.image"
              class="w-full h-full object-cover"
              :src="data.image[3]"
              :srcset="`${data.image[0]} 400w, ${data.image[1]} 800w, ${data.image[2]} 1200w, ${data.image[3]} 1600w`"
              :alt="data.alt"
            />
            <img
              v-if="data.animation"
              :src="data.animation"
              class="absolute inset-0 m-auto object-cover"
              :alt="data.alt"
            />
            <video
              v-else
              class="h-full w-2/3 object-cover m-auto rounded-t-lg pt-16 sm:pt-10 lg:pt-20"
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
                ? 'v2-normal-text px-6 py-12 lg:px-12 lg:py-12 xl:p-20 text-center'
                : 'pt-6',
            ]"
          >
            {{ data.title }}
          </div>
        </div>
      </div>
      <div v-if="gridData2" class="sm:mt-36 lg:mt-60 basis-1/2">
        <div v-for="data in gridData2" :key="data" class="px-3">
          <aspect-ratio
            :height="data.aspectRatio"
            :style="[data.background ? { background: data.background } : {}]"
          >
            <img
              v-if="data.image"
              class="w-full h-full object-cover"
              :src="data.image[3]"
              :srcset="`${data.image[0]} 400w, ${data.image[1]} 800w, ${data.image[2]} 1200w, ${data.image[3]} 1600w`"
              :alt="data.alt"
            />
            <img
              v-if="data.animation"
              :src="data.animation"
              class="absolute inset-0 m-auto object-cover rounded-3xl h-6/6"
              :alt="data.alt"
            />
            <video
              v-else
              class="h-full w-full object-cover px-20 pt-16 sm:px-8 sm:pt-10 lg:pt-20 lg:px-24"
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
                ? 'v2-normal-text px-6 py-12 lg:px-12 lg:py-12 xl:p-20 text-center'
                : 'pt-6',
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
