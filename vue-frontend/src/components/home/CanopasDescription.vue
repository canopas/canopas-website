<template>
  <div class="container-fluid">
    <div class="container" ref="canopasDescription">
      <div
        ref="description"
        class="mt-4 canopas-text-div normal-text animate__animated"
      >
        <h4 class="header-text canopas-gradient-text mb-4">What is Canopas?</h4>
        Canopas is a company helping Entrepreneurs, startups, and small
        companies to bring their IDEA to LIFE by developing digital products(or
        improving existing products) for their business.
        <br /><br />
        Our vision is
        <b
          ><i
            ><q
              >To put humanity in better condition compared to what it was
              before our presence in the world.</q
            ></i
          ></b
        >
        <br /><br />
        We're quality-focused people who want to work on a few handful apps that
        can be used by many people worldwide.
      </div>
      <div class="bongo-cat-wrapper">
        <div class="bongo-cat-div text-right animate__animated" ref="bongoCat">
          <BongoCat class="bongo-cat-image" />
        </div>
      </div>
    </div>
  </div>
</template>

<script type="module">
import BongoCat from "./BongoCat.vue";

export default {
  data() {
    return {};
  },
  components: {
    BongoCat,
  },
  mounted() {
    window.addEventListener("scroll", this.handleScroll);
  },
  unmounted() {
    window.removeEventListener("scroll", this.handleScroll);
  },
  methods: {
    handleScroll() {
      let data = {
        name: this.$refs.canopasDescription,
        animation: "",
        childRef: [
          {
            name: this.$refs.description,
            animation: "animate__fadeInLeft",
          },
          {
            name: this.$refs.bongoCat,
            animation: "animate__zoomIn",
          },
        ],
      };
      if (data) {
        this.$emit("add-animation", data);
      }
    },
  },
};
</script>

<style lang="scss" scoped>
.container-fluid {
  padding: 5% 10%;
}

.container {
  display: flex;
  flex-direction: column-reverse;
  margin: 5% 0 5% auto;
  padding: 0;
  overflow: hidden;
  align-items: center;
}

.bongo-cat-wrapper {
  width: 90%;
}

.bongo-cat-div {
  position: relative;
  width: 100%;
  padding-top: 63%;
}

.bongo-cat-image {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
}

@include media-breakpoint-up(md) {
  .container {
    flex-direction: row;
    margin: 5% auto;
  }

  .canopas-text-div {
    flex: 1 1 0;
  }

  .bongo-cat-wrapper {
    flex: 1 1 0;
    margin-right: -10%;
    margin-left: 5%;
  }
}

@include media-breakpoint-up(xl) {
  .bongo-cat-wrapper {
    flex: 1.3 1 0;
  }
}
</style>
