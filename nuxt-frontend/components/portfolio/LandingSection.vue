<template>
  <section class="bg-white" :ref="response.landingref1">
    <div class="container z pb-16 lg:pb-60">
      <h1
        class="sub-h1-semibold lg:sub-h2-medium text-black-60 animate__animated animate__fadeInUp"
      >
        {{ response.title }}
      </h1>
      <div
        class="lg:w-4/5 mobile-header-1 lg:desk-header-1 text-black-87 mt-4 lg:mt-6 animate__animated animate__slideInUp"
      >
        {{ response.subTitle }}
      </div>
    </div>
  </section>

  <section class="background-image relative z-[-1]" :ref="response.landingref2">
    <aspect-ratio :height="isMobile ? '100%' : '56.26%'">
      <img
        :src="response.backgroundImage[3]"
        :srcset="`${response.backgroundImage[0]} 400w, ${response.backgroundImage[1]} 800w, ${response.backgroundImage[2]} 1400w, ${response.backgroundImage[3]} 2400w`"
        class="w-full h-full object-cover"
        :alt="response.alt"
      />
    </aspect-ratio>
  </section>
</template>

<script>
import AspectRatio from "@/components/utils/AspectRatio.vue";

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
  },
  mounted() {
    if (window.innerWidth < 768) {
      this.isMobile = true;
    }
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
