<template>
  <section class="bg-white">
    <div
      v-if="response.solution"
      class="container py-20 lg:py-80 relative"
      :ref="response.solution.ref"
    >
      <div class="flex flex-col justify-between lg:flex-row">
        <div class="v2-normal-text font-bold">
          {{ response.solution.title }}
        </div>
        <div class="pt-5 lg:pl-16 lg:w-4/5 lg:pt-0">
          <div
            class="v2-normal-text font-light"
            v-html="response.solution.description"
          ></div>

          <div class="flex w-fit flex-row pt-12 md:pt-16">
            <div
              v-for="button in response.solution.buttons"
              :key="button"
              class="flex items-center pr-4 py-2"
            >
              <a
                target="_blank"
                :href="button.link"
                @click.native="$mixpanel.track(button.event)"
              >
                <img :src="button.image" class="w-48" :alt="button.name"
              /></a>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div
      v-if="response.title"
      class="container pt-0 pb-20 md:pt-16 lg:pb-80 relative"
    >
      <div class="v2-header-3-text pt-20" v-html="response.title"></div>
    </div>
  </section>

  <section
    v-if="response.backgroundImage[3]"
    class="background-image relative z-[-1]"
    :ref="response.brandingref ? response.brandingref : ''"
  >
    <aspect-ratio height="100%">
      <img
        :src="response.backgroundImage[3]"
        :srcset="`${response.backgroundImage[0]} 400w, ${response.backgroundImage[1]} 800w, ${response.backgroundImage[2]} 1400, ${response.backgroundImage[3]} 2400w`"
        class="w-full h-full object-cover"
        :alt="response.alt"
      />
    </aspect-ratio>
  </section>

  <section class="bg-white" :ref="response.details ? response.details.ref : ''">
    <div
      class="container flex flex-col md:flex-row md:gap-x-16 pt-24 md:pt-48 xl:pt-80"
    >
      <div v-if="gridData1" class="basis-1/2">
        <div v-for="data in gridData1" :key="data">
          <div class="relative">
            <aspect-ratio
              :height="data.aspectRatio"
              :style="[data.background ? { background: data.background } : {}]"
            >
              <div
                :class="[data.video ? 'px-4 pb-4 md:px-2 lg:px-8 lg:pb-8' : '']"
              >
                <img
                  v-if="data.image"
                  :src="data.image[2]"
                  :srcset="`${data.image[0]} 400w, ${data.image[1]} 800w, ${data.image[2]} 1200w`"
                  class="w-full h-full object-cover"
                  :alt="data.alt"
                />
                <img
                  v-if="data.gif"
                  :src="data.gif"
                  class="absolute inset-0 m-auto lg:rounded-xl rounded-3xl object-cover h-3/3"
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
                      ? 'rounded-b-[90px] border-b-8 border-x-8 border-solid border-white px-8 pb-8'
                      : '',
                  ]"
                >
                  <source :src="data.video" type="video/mp4" />
                </video>
              </div>
            </aspect-ratio>
          </div>
          <div
            class="v2-normal-text bg-white font-light ... px-6 sm:px-8 lg:px-12 xl:px-20 py-12 xl:py-20 text-center"
            :class="[
              response.details.ref == 'justlyui1'
                ? ' pb-8 sm:pb-12 xl:pb-8  2xl:pb-12 text-[1.188rem] leading-[1.781rem] md:text-[1.281rem] md:leading-[1.875rem] xl:text-[1.375rem] xl:leading-8 font-inter-regular text-black-core/[0.87] '
                : '',
              response.details.ref == 'justlyui1' && data.id == 2
                ? 'pt-32 sm:pt-28 md:pt-20 lg:pt-28 xl:pt-28 2xl:pt-36 md:!pb-0'
                : '',
              response.details.ref == 'justlyui1' && data.id == 1
                ? 'pt-40 sm:pt-40 md:pt-28 lg:pt-36 xl:pt-40 2xl:pt-48'
                : '',
            ]"
          >
            {{ data.title }}
          </div>
        </div>
      </div>
      <div
        v-if="gridData2"
        class="basis-1/2"
        :class="[
          response.details.ref !== 'justlyui1' ? 'md:mt-36 lg:mt-64' : '',
        ]"
      >
        <div v-for="data in gridData2" :key="data">
          <div class="relative">
            <aspect-ratio
              :height="data.aspectRatio"
              :style="[data.background ? { background: data.background } : {}]"
            >
              <div
                :class="[data.video ? 'px-4 pb-4 md:px-2 lg:px-8 lg:pb-8' : '']"
              >
                <img
                  v-if="data.image"
                  class="w-full h-full object-cover"
                  :src="data.image[2]"
                  :srcset="`${data.image[0]} 400w, ${data.image[1]} 800w, ${data.image[2]} 1200w`"
                  :alt="data.alt"
                />
                <video
                  v-if="data.video"
                  preload="auto"
                  loop
                  muted
                  autoplay
                  playsinline
                  class="absolute inset-0 m-auto object-cover"
                  :class="
                    data.title !== 'togness'
                      ? 'rounded-2xl md:rounded-xl xl:rounded-3xl h-3/3'
                      : ''
                  "
                >
                  <source :src="data.video" type="video/mp4" />
                </video>
                <LottieAnimation
                  v-else-if="data.animation"
                  :jsonData="data.animation"
                  class="mt-14 lg:mt-20"
                />
              </div>
            </aspect-ratio>
          </div>
          <div
            class="v2-normal-text bg-white font-light ... px-6 sm:px-8 lg:px-12 xl:px-20 py-12 xl:py-20 text-center"
            :class="[
              response.details.ref == 'justlyui1'
                ? 'px-6  sm:px-8 lg:px-12  2xl:px-20 pb-8 sm:pb-12 md:pb-44 lg:pb-52 xl:pb-[13.5rem]  2xl:pb-[14.8rem] text-[1.188rem] leading-[1.781rem] md:text-[1.281rem] md:leading-[1.875rem] xl:text-[1.375rem] xl:leading-8 font-inter-regular text-center text-black-core/[0.87] '
                : '',
              response.details.ref == 'justlyui1' && data.id == 4
                ? 'pt-4 md:pt-4 xl:pt-0  md:!pb-0'
                : '',
              response.details.ref == 'justlyui1' && data.id == 3
                ? 'pt-56 sm:pt-52 md:pt-36 lg:pt-48 xl:pt-56 2xl:pt-64'
                : '',
              response.details.ref == 'tognessui1' && data.title == 'togness'
                ? 'hidden'
                : '',
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
  props: ["json"],
  data() {
    return {
      isMobile: false,
      response: this.json,
      gridData1: this.json.details.gridData1,
      gridData2: this.json.details.gridData2,
      event: "",
      events: {
        luxesolution: "view_luxe_radio_solution",
        luxeparallax1: "view_luxe_first_parallax_section",
        luxeui1: "view_luxe_ui_section1",
        tognesssolution: "view_togness_solution",
        tognessui1: "view_togness_ui_section1",
        justlyparallax1: "view_justly_parallax",
        justlyui1: "view_justly_ui_section1",
      },
    };
  },
  watch: {
    json: function (newVal, oldVal) {
      this.response = newVal;
      this.gridData1 = this.response.details.gridData1;
      this.gridData2 = this.response.details.gridData2;
    },
  },
  components: {
    AspectRatio,
    LottieAnimation,
  },
  mounted() {
    if (window.innerWidth < 768) {
      this.isMobile = true;
    }
  },
  inject: ["mixpanel"],
};
</script>
<style lang="postcss" scoped>
@screen lg {
  section.background-image {
    transform: translateZ(-1px) scale(1.5);
  }
}
@screen 3xl {
  section.background-image {
    transform: translateZ(-1px) scale(1.3);
  }
}
</style>
