<template>
  <div class="w-full h-[550px] border-0 md:h-screen">
    <div v-if="isLoading" class="iframe-loader">
      <img
        :src="loader"
        class="fixed left-[50%] top-[50%] -translate-y-1/2 -translate-x-1/2 z-[100]"
        alt="loader-image"
      />
    </div>
    <div v-if="!isLoading" class="h-28" id="calendly-embed-before"></div>
    <div
      id="calendly-embed"
      class="h-[calc(100vh-7rem)] w-full"
      :data-url="calendlyUrl"
      style="min-width: 320px"
    ></div>
  </div>
</template>

<script type="module">
import config from "@/config.js";
import loader from "@/assets/images/theme/loader.svg";

export default {
  data() {
    return {
      calendlyUrl: config.CALENDLY_IFRAME_URL,
      loader: loader,
      isLoading: true,
    };
  },
  mounted() {
    this.loadCalendlyWidget();
    this.initializeCalendlyEventListeners();
  },
  beforeDestroy() {
    window.removeEventListener("message", this.handleCalendlyEvent);
  },
  methods: {
    loadCalendlyWidget() {
      if (window.Calendly) {
        this.initCalendlyWidget();
      } else {
        const script = document.createElement("script");
        script.src = "https://assets.calendly.com/assets/external/widget.js";
        script.async = true;
        script.onload = this.initCalendlyWidget;
        document.head.appendChild(script);
      }
    },
    initCalendlyWidget() {
      Calendly.initInlineWidget({
        url: this.calendlyUrl,
        parentElement: document.getElementById("calendly-embed"),
      });
    },
    initializeCalendlyEventListeners() {
      setTimeout(() => {
        this.isLoading = false;
      }, 1000);
      window.addEventListener("message", this.handleCalendlyEvent);
    },
    handleCalendlyEvent(e) {
      if (
        e.data.event &&
        e.data.event.indexOf("calendly") === 0 &&
        e.data.event === "calendly.event_scheduled"
      ) {
        setTimeout(() => {
          this.$emit("close");
        }, 2000);
      }
    },
  },
};
</script>

<style scroped>
@media (min-width: 650px) and (max-width: 1000px) {
  #calendly-embed-before {
    height: 4.5rem;
  }

  #calendly-embed {
    height: calc(100vh - 4.5rem);
  }
}
</style>
