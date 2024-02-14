<template>
  <section class="bg-white relative">
    <div class="relative container">
      <div v-if="response.title">
        <div class="pt-10 pb-16 md:pt-16 lg:pt-28 lg:pb-60 pt-0">
          <div
            :class="
              response.class
                ? response.class
                : 'mobile-header-3-semibold lg:desk-header-2 text-black-60'
            "
            v-html="response.title"
          ></div>
        </div>
      </div>
    </div>
  </section>

  <section v-if="response.detail" class="relative bg-white lg:pb-60">
    <div class="container flex flex-col sm:flex-row sm:gap-x-8 relative">
      <div v-if="gridData1" class="basis-1/2">
        <div v-for="data in gridData1" :key="data">
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
              class="h-full w-2/3 object-cover m-auto rounded-t-md pt-16 sm:pt-10 lg:pt-20"
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
                ? 'mobile-header-3-regular lg:mobile-header-2-regular text-black-87 px-6 pt-6 pb-14 text-center'
                : 'pt-8',
            ]"
          >
            {{ data.title }}
          </div>
        </div>
      </div>
      <div v-if="gridData2" class="sm:mt-36 lg:mt-60 basis-1/2">
        <div v-for="data in gridData2" :key="data">
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
                ? 'mobile-header-3-regular lg:mobile-header-2-regular px-6 py-12 lg:px-12 lg:py-12 xl:p-20 text-center'
                : 'pt-8',
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
