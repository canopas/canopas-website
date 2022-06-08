<template>
  <div>
    <Header />
    <VideoSection v-bind:response="details.video" />
  </div>
</template>

<script>
import Header from "@/components/partials/ScreenHeaderV2.vue";
import VideoSection from "@/components/portfolio/VideoSection.vue";
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
  },
};
</script>
