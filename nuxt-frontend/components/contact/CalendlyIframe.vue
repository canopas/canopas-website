<template>
  <div class="w-full h-[550px] border-0 pt-10 md:h-screen">
    <div v-if="isLoading" class="iframe-loader">
      <img
        :src="loader"
        class="fixed left-[50%] top-[50%] -translate-y-1/2 -translate-x-1/2 z-[100]"
        alt="loader-image"
      />
    </div>
    <div
      id="calendly-embed"
      class="calendly-inline-widget h-screen w-full"
      :data-url="calendlyUrl"
      style="min-width: 320px; height: 700px"
      v-on:load="isLoading = false"
    ></div>
  </div>
</template>

<script>
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

      window.addEventListener("message", function (e) {
        if (
          e.data.event &&
          e.data.event.indexOf("calendly") === 0 &&
          e.data.event === "calendly.event_scheduled"
        ) {
          this.$router.push({
            path: "/thank-you",
          });
        }
      });
    },
  },
};
</script>
