<template>
  <section
    v-if="response[1] || response[0].designref2 == 'tognessfeedback1'"
    class="tw-bg-white tw-relative"
  >
    <div class="tw-relative tw-container">
      <div :ref="response[0].designref ? response[0].designref : ''">
        <div
          class="tw-pt-20 lg:tw-py-80 tw-flex tw-flex-col tw-justify-between xl:tw-w-3/4"
        >
          <div
            class="v2-title-text tw-font-bold"
            v-html="response[0].title"
          ></div>
          <div
            v-if="response[0].description"
            class="description tw-pt-5 tw-w-4/5 lg:tw-pt-10 xl:tw-pt-20 xl:tw-w-full"
          >
            <div class="v2-normal-text tw-font-light">
              {{ response[0].description }}
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>

  <section
    v-if="response[1] || response[0].designref2 == 'tognessfeedback1'"
    class="image tw-relative tw-z-[-1]"
    :ref="response[0].designref1 ? response[0].designref1 : ''"
  >
    <aspect-ratio
      :height="isMobile ? '100%' : '56.26%'"
      class="tw-overflow-hidden"
    >
      <img
        :src="response[0].backgroundImage[3]"
        :srcset="`${response[0].backgroundImage[0]} 400w, ${response[0].backgroundImage[1]} 800w, ${response[0].backgroundImage[2]} 1400w, ${response[0].backgroundImage[3]} 2400w`"
        class="tw-w-full tw-h-full"
        :class="isMobile ? 'tw-object-contain' : 'tw-object-cover'"
        :alt="response[0].alt"
      />
    </aspect-ratio>
    <LottieAnimation
      v-if="response[0].animation"
      :jsonData="response[0].animation"
      class="tw-absolute tw-inset-0 tw-left-auto -tw-top-1/6 tw-m-auto tw-h-2/4 tw-w-[31.333333%] tw-object-cover -tw-scale-x-100"
    />
  </section>
  <section v-if="response[0].subTitle" class="tw-bg-white">
    <div
      class="tw-container tw-flex tw-flex-col tw-justify-between tw-pb-20 sm:tw-py-40 lg:tw-flex-row lg:tw-py-80"
      :ref="response[0].subTitle ? response[0].designref2 : ''"
    >
      <div class="v2-normal-text tw-font-bold">
        {{ response[0].subTitle }}
      </div>
      <div class="tw-pt-5 lg:tw-pl-16 lg:tw-w-4/5 lg:tw-pt-0">
        <div
          class="v2-normal-text tw-font-light"
          v-html="response[0].subDescription"
        ></div>
      </div>
    </div>
  </section>

  <section class="tw-bg-white tw-relative" v-if="response[1]">
    <div class="tw-relative" :ref="response[1].designref">
      <div
        class="tw-container tw-py-20 lg:tw-py-80 tw-flex tw-flex-col tw-justify-between xl:tw-flex-row"
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

  <section class="image tw-relative tw-z-[-1]" v-if="response[1]">
    <aspect-ratio
      :height="isMobile ? '96.26%' : '56.26%'"
      class="tw-overflow-hidden"
    >
      <img
        :src="response[1].backgroundImage[3]"
        :srcset="`${response[1].backgroundImage[0]} 400w, ${response[1].backgroundImage[1]} 800w,${response[1].backgroundImage[2]} 1400w,${response[1].backgroundImage[3]} 2400w`"
        class="tw-w-full tw-h-full"
        :class="isMobile ? 'tw-object-contain' : 'tw-object-cover'"
        :alt="response[1].alt"
      />
    </aspect-ratio>
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
      event: "",
      events: {
        luxefeedback1: "view_luxe_feedback1",
        tognessfeedback1: "view_togness_feedback1",
        justlyfeedback1: "view_justly_feedback1",
        tognessparallax1: "view_togness_parallax1",
      },
    };
  },
  watch: {
    json: function (newVal, oldVal) {
      this.response = newVal;
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
  section.image {
    transform: translateZ(-1px) scale(1.5);
  }
}
@screen 3xl {
  section.image {
    transform: translateZ(-1px) scale(1.3);
  }
}
</style>
