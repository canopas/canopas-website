<template>
  <section
    v-if="response[1] || response[0].designref2 == 'tognessfeedback1'"
    class="bg-white relative"
  >
    <div class="relative container">
      <div :ref="response[0].designref ? response[0].designref : ''">
        <div
          class="py-16 lg:pt-60 lg:pb-[4.5rem] flex flex-col justify-between xl:w-3/4"
        >
          <div
            class="mobile-header-2 lg:desk-header-2 text-black-87"
            v-html="response[0].title"
          ></div>
          <div
            v-if="response[0].description"
            class="description pt-5 lg:pt-6 xl:w-full"
          >
            <div
              class="sub-h1-regular lg:mobile-header-2-regular text-black-60"
            >
              {{ response[0].description }}
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>

  <section
    v-if="response[1] || response[0].designref2 == 'tognessfeedback1'"
    class="image relative z-[-1]"
    :ref="response[0].designref1 ? response[0].designref1 : ''"
  >
    <aspect-ratio :height="isMobile ? '70%' : '56.26%'" class="overflow-hidden">
      <img
        :src="response[0].backgroundImage[3]"
        :srcset="`${response[0].backgroundImage[0]} 400w, ${response[0].backgroundImage[1]} 800w, ${response[0].backgroundImage[2]} 1400w, ${response[0].backgroundImage[3]} 2400w`"
        class="w-full h-full"
        :class="isMobile ? 'object-contain' : 'object-cover'"
        :alt="response[0].alt"
      />
    </aspect-ratio>
    <LottieAnimation
      v-if="response[0].animation"
      :jsonData="response[0].animation"
      class="absolute inset-0 left-auto -top-1/6 m-auto h-2/4 w-[31.333333%] object-cover -scale-x-100"
    />
  </section>
  <section v-if="response[0].subTitle" class="bg-white">
    <div
      class="container flex flex-col justify-between py-16 lg:flex-row lg:py-60"
      :ref="response[0].subTitle ? response[0].designref2 : ''"
    >
      <div class="mobile-header-3-semibold lg:desk-header-3 text-black-87">
        {{ response[0].subTitle }}
      </div>
      <div class="pt-5 lg:pl-16 lg:w-4/5 lg:pt-0">
        <div
          class="sub-h1-regular lg:mobile-header-2-regular text-black-60"
          v-html="response[0].subDescription"
        ></div>
      </div>
    </div>
  </section>

  <section class="bg-white relative" v-if="response[1]">
    <div class="relative" :ref="response[1].designref">
      <div
        class="container py-16 lg:py-60 flex flex-col justify-between xl:flex-row"
      >
        <div class="mobile-header-3-semibold lg:desk-header-3 text-black-87">
          {{ response[1].title }}
        </div>
        <div class="description lg:w-11/12 xl:pl-16">
          <div
            class="mt-4 sub-h1-regular lg:mobile-header-2-regular text-black-60"
          >
            {{ response[1].description }}
          </div>
        </div>
      </div>
    </div>
  </section>

  <section class="image relative z-[-1]" v-if="response[1]">
    <aspect-ratio :height="isMobile ? '70%' : '56.26%'" class="overflow-hidden">
      <img
        :src="response[1].backgroundImage[3]"
        :srcset="`${response[1].backgroundImage[0]} 400w, ${response[1].backgroundImage[1]} 800w,${response[1].backgroundImage[2]} 1400w,${response[1].backgroundImage[3]} 2400w`"
        class="w-full h-full"
        :class="isMobile ? 'object-contain' : 'object-cover'"
        :alt="response[1].alt"
      />
    </aspect-ratio>
  </section>
</template>

<script>
import AspectRatio from "@/components/utils/AspectRatio.vue";
import { elementInViewPort } from "@/utils.js";

const LottieAnimation = defineAsyncComponent(
  () => import("@/components/utils/LottieAnimation.vue"),
);

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
        this.$mixpanel.track(event);
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
