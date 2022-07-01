<template>
  <div ref="lavContainer" />
</template>

<script>
import lottie from "lottie-web/build/player/lottie_light.js";

export default {
  props: {
    jsonData: {
      required: true,
    },
    speed: {
      type: Number,
      required: false,
      default: 1,
    },
    loop: {
      type: Boolean,
      required: false,
      default: true,
    },
    autoPlay: {
      type: Boolean,
      required: false,
      default: true,
    },
    loopDelayMin: {
      type: Number,
      required: false,
      default: 0,
    },
    loopDelayMax: {
      type: Number,
      required: false,
      default: 0,
    },
  },
  data: () => ({
    name: "lottie-animation",
    rendererSettings: {
      progressiveLoad: true,
    },
    anim: null,
  }),
  mounted() {
    this.init();
  },
  methods: {
    async init() {
      if (this.anim) {
        this.anim.destroy(); // Releases resources. The DOM element will be emptied.
      }
      this.anim = lottie.loadAnimation({
        container: this.$refs.lavContainer,
        renderer: "svg",
        loop: this.loop,
        autoplay: this.autoPlay,
        animationData: this.jsonData,
        rendererSettings: this.rendererSettings,
      });
      this.$emit("AnimControl", this.anim);
      this.anim.setSpeed(this.speed);
      if (this.loopDelayMin > 0) {
        this.anim.loop = false;
        this.anim.autoplay = false;
        this.executeLoop();
      }
    },
    getRandomInt(min, max) {
      min = Math.ceil(min);
      max = Math.floor(max);
      return Math.floor(Math.random() * (max - min)) + min; //The maximum is exclusive and the minimum is inclusive
    },
    executeLoop() {
      this.anim.play();
      setTimeout(() => {
        this.anim.stop();
        this.executeLoop();
      }, this.getRandomInt(this.loopDelayMin, this.loopDelayMax == 0 ? this.loopDelayMin : this.loopDelayMax));
    },
  },
  watch: {
    path: function () {
      this.init();
    },
  },
};
</script>
