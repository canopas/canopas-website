<template>
  <section class="tw-bg-white tw-relative">
    <div class="tw-relative container">
      <div>
        <div
          class="tw-py-40 lg:tw-py-80 tw-flex tw-flex-col tw-justify-between xl:tw-w-3/4"
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

  <section class="image tw-relative tw-z-[-1]">
    <aspect-ratio
      :height="isMobile ? '100%' : '56.26%'"
      class="tw-overflow-hidden"
    >
      <img
        :src="response[0].backgroundImage[3]"
        :srcset="`${response[0].backgroundImage[0]} 400w, ${response[0].backgroundImage[1]} 800w, ${response[0].backgroundImage[2]} 1400w, ${response[0].backgroundImage[3]} 2400w`"
        class="tw-w-full tw-h-full tw-object-cover"
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
      class="container tw-flex tw-flex-col tw-justify-between tw-py-20 sm:tw-py-40 lg:tw-flex-row lg:tw-py-80"
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
    <div class="tw-relative">
      <div
        class="container tw-py-40 lg:tw-py-80 tw-flex tw-flex-col tw-justify-between xl:tw-flex-row"
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
        class="tw-w-full tw-h-full tw-object-cover"
        :alt="response[1].alt"
      />
    </aspect-ratio>
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
  },
};
</script>

<style lang="scss" scoped>
@include media-breakpoint-up(lg) {
  section.image {
    transform: translateZ(-1px) scale(1.5);
  }
}
@media (min-width: 3840px) {
  section.image {
    transform: translateZ(-1px) scale(1.3);
  }
}
</style>
