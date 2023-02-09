<template>
  <section class="tw-bg-white" :ref="response.landingref1">
    <div
      class="tw-container tw-pt-20 tw-pb-40 sm:tw-pt-32 lg:tw-pb-80 lg:tw-pt-60"
    >
      <div class="v2-normal-2-text animate__animated animate__fadeInUp">
        {{ response.title }}
      </div>
      <div class="v2-title-text tw-mt-10 animate__animated animate__slideInUp">
        {{ response.subTitle }}
      </div>
    </div>
  </section>

  <section
    class="background-image tw-relative tw-z-[-1]"
    :ref="response.landingref2"
  >
    <aspect-ratio :height="isMobile ? '100%' : '56.26%'">
      <img
        :src="response.backgroundImage[3]"
        :srcset="`${response.backgroundImage[0]} 400w, ${response.backgroundImage[1]} 800w, ${response.backgroundImage[2]} 1400w, ${response.backgroundImage[3]} 2400w`"
        class="tw-w-full tw-h-full tw-object-cover"
        :alt="response.alt"
      />
    </aspect-ratio>
  </section>
</template>

<script>
import AspectRatio from "@/components/utils/AspectRatio.vue";
import { elementInViewPort } from "@/utils.js";

export default {
  props: ["json"],
  data() {
    return {
      isMobile: false,
      response: this.json,
      event: "",
      events: {
        luxelanding: "view_luxe_radio_landing",
        luxebanner: "view_luxe_first_banner_image",
        tognesslanding: "view_togness_landing",
        tognessbanner: "view_togness_first_banner_image",
        justlylanding: "view_justly_landing",
        justlybanner: "view_justly_first_banner_image",
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
  },
  mounted() {
    if (window.innerWidth < 768) {
      this.isMobile = true;
    }
    window.addEventListener("scroll", this.sendEvent);
  },
  unmounted() {
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
