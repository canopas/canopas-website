<template>
  <div>
    <Header />
    <div class="parallax">
      <VideoSection v-bind:response="details.video" />
      <DesignSection v-bind:response="details.design" />
    </div>
  </div>
</template>

<script>
import Header from "@/components/partials/ScreenHeaderV2.vue";
import VideoSection from "@/components/portfolio/VideoSection.vue";
import DesignSection from "@/components/portfolio/DesignSection.vue";
import luxeradioResponse from "@/luxeradio-data.js";

export default {
  data() {
    return {
      id: this.$route.params.id,
      details: "",
    };
  },
  created() {
    this.getData();
  },
  methods: {
    getData() {
      if (this.id == luxeradioResponse.name) {
        this.details = luxeradioResponse.detail;
      }
    },
  },

  components: {
    Header,
    VideoSection,
    DesignSection,
  },
};
</script>

<style scoped>
.parallax {
  height: 100vh;
  overflow: hidden;
  overflow-y: auto;
  perspective: 2px;
}
.parallax::-webkit-scrollbar {
  display: none; /* Chrome, Safari and Opera */
}
.parallax {
  -ms-overflow-style: none; /* IE and Edge */
  scrollbar-width: none; /* Firefox */
}
</style>
