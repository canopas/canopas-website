<template>
  <div>
    <Header />
  </div>
</template>

<script>
import Header from "@/components/partials/ScreenHeaderV2.vue";
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
  },
};
</script>
