<template>
  <section
    class="background-image relative z-[-1] bg-white"
    :ref="response.ref"
  >
    <aspect-ratio :height="isMobile ? '100%' : '56.26%'">
      <img
        :src="response.backgroundImage[3]"
        :srcset="`${response.backgroundImage[0]} 400w, ${response.backgroundImage[1]} 800w, ${response.backgroundImage[2]} 1400w, ${response.backgroundImage[3]} 2400w`"
        class="w-full h-full"
        :class="isMobile ? 'object-contain' : 'object-cover'"
        :alt="response.alt"
      />
    </aspect-ratio>
  </section>
  <section>
    <UserReviewSection
      v-bind:json="response.review"
      :ref="response.review.ref"
    />
  </section>
  <section
    class="v2-header-2-text font-bold text-center bg-white pt-20 sm:py-40"
  >
    <nuxt-link
      class="animation-underline inline-block relative hover:text-[#3d3d3d] after:content-[''] after:absolute after:w-full after:scale-x-0 after:h-0.5 after:bottom-0 after:left-0 after:bg-black-900 after:origin-bottom-left after:duration-300 hover:after:scale-x-100 hover:after:origin-bottom-left"
      :to="response.url"
      @click.native="$mixpanel.track(response.event)"
      >{{ response.title }}</nuxt-link
    >
  </section>
</template>

<script>
import AspectRatio from "@/components/utils/AspectRatio.vue";
import UserReviewSection from "@/components/portfolio/UserReviewSection.vue";
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
        luxeparallax2: "view_luxe_last_parallax_section",
        tognessparallax2: "view_togness_last_parallax",
        justlyparallax2: "view_justly_last_parallax",
        luxeradioreview: "view_luxeradio_user_review",
        tognessreview: "view_togness_user_review",
        justlyreview: "view_justly_user_review",
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
    UserReviewSection,
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
