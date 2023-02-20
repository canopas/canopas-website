<template>
  <section class="tw-bg-white">
    <div
      v-if="response.solution"
      class="tw-container tw-py-20 lg:tw-py-80 tw-relative"
      :ref="response.solution.ref"
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

          <div class="tw-flex tw-w-fit tw-flex-row tw-pt-12 md:tw-pt-16">
            <div
              v-for="button in response.solution.buttons"
              :key="button"
              class="tw-flex tw-items-center tw-pr-4 tw-py-2"
            >
              <a
                target="_blank"
                :href="button.link"
                @click.native="mixpanel.track(button.event)"
              >
                <img :src="button.image" class="tw-w-48"
              /></a>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div
      v-if="response.title"
      class="tw-container tw-py-20 lg:tw-py-80 tw-relative"
    >
      <div class="v2-header-3-text tw-pt-20" v-html="response.title"></div>
    </div>
  </section>

  <section
    v-if="response.backgroundImage[3]"
    class="background-image tw-relative tw-z-[-1]"
    :ref="response.brandingref ? response.brandingref : ''"
  >
    <aspect-ratio :height="isMobile ? '100%' : '56.26%'">
      <img
        :src="response.backgroundImage[3]"
        :srcset="`${response.backgroundImage[0]} 400w, ${response.backgroundImage[1]} 800w, ${response.backgroundImage[2]} 1400, ${response.backgroundImage[3]} 2400w`"
        class="tw-w-full tw-h-full tw-object-cover"
        :alt="response.alt"
      />
    </aspect-ratio>
  </section>

  <section
    class="tw-bg-white"
    :ref="response.details ? response.details.ref : ''"
  >
    <div
      class="tw-container tw-pt-48 md:tw-pt-48 xl:tw-pt-80 tw-flex tw-flex-col ... md:tw-flex-row ... md:tw-gap-x-16"
    >
      <div v-if="gridData1" class="tw-basis-1/2">
        <div v-for="data in gridData1" :key="data">
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
                  class="tw-absolute tw-inset-0 tw-m-auto lg:tw-rounded-xl tw-rounded-3xl tw-object-cover tw-h-3/3"
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
                ? 'tw-px-6 tw-py-12 sm:tw-px-8 lg:tw-px-12 lg:tw-py-12 xl:tw-px-20 xl:tw-py-20 tw-text-center'
                : 'tw-mt-12',
            ]"
          >
            {{ data.title }}
          </div>
        </div>
      </div>
      <div v-if="gridData2" class="tw-basis-1/2 md:tw-mt-36 lg:tw-mt-64">
        <div v-for="data in gridData2" :key="data">
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
                  class="tw-absolute tw-inset-0 tw-m-auto tw-rounded-2xl md:tw-rounded-xl xl:tw-rounded-3xl tw-object-cover tw-h-3/3"
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
                ? 'tw-px-6 tw-py-12 sm:tw-px-8 lg:tw-px-12 lg:tw-py-12 xl:tw-px-20 xl:tw-py-20 tw-text-center'
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
import { elementInViewPort } from "@/utils.js";

export default {
  props: ["json"],
  data() {
    return {
      portfolioRef: null,
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
    this.portfolioRef = document.getElementById("response");
    this.portfolioRef.addEventListener("scroll", this.sendEvent);
    window.addEventListener("scroll", this.sendEvent);
  },
  unmounted() {
    this.portfolioRef.removeEventListener("scroll", this.sendEvent);
    window.removeEventListener("scroll", this.sendEvent);
  },
  inject: ["mixpanel"],
  methods: {
    sendEvent() {
      const event = this.events[elementInViewPort(this.$refs)];
      if (event && this.event !== event) {
        this.event = event;
        this.mixpanel.track(event);
      }
    },
  },
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
