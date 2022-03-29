<template>
  <div class="container-fluid">
    <div class="container-xl">
      <div class="flex-row" ref="doYouKnow">
        <div ref="doYouKnowImage" class="que-image animate__animated">
          <img :src="doYouKnowImage" alt="do-you-know" />
        </div>
        <div ref="doYouKnowDetails" class="container animate__animated">
          <h4 class="header-text pb-30">Do you know ?</h4>
          <span class="normal-text pink-text mt-20">
            We’ve worked with just <b>10 clients</b> in the last <b>5 years</b>.
            WHY?
          </span>
          <br /><br />
          <span class="normal-text mt-20">
            Because 8 of them have been offering(committed to) us repeat
            business since we started working for them. Think for a moment,
            Isn't it the main reason why you're here.</span
          >
        </div>
      </div>
    </div>
  </div>
</template>

<script type="module">
import doYouKnowImage from "@/assets/images/theme/question.svg";

export default {
  data() {
    return {
      doYouKnowImage: doYouKnowImage,
    };
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
        name: this.$refs.doYouKnow,
        animation: "",
        childRef: [
          {
            name: this.$refs.doYouKnowImage,
            animation: "animate__fadeInLeft",
          },
          {
            name: this.$refs.doYouKnowDetails,
            animation: "animate__fadeInRight",
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
  z-index: 10000;
  background: white;
}

.container {
  flex: 75%;
  margin: 10% auto;
}

.flex-row {
  display: flex;
  align-items: center;
}

.que-image {
  display: none;
}

.pink-text {
  color: #f2709c;
}

@include media-breakpoint-up(md) {
  .que-image {
    display: block;
    width: 25%;
  }

  .container {
    padding-left: 10%;
    margin-bottom: 28%;
  }
}

@include media-breakpoint-up(lg) {
  .do-you-know-div {
    width: 70%;
  }
}

@include media-breakpoint-up(xl) {
  .que-image {
    width: 24%;
  }

  .container {
    margin-bottom: 15%;
    padding-left: 14%;
  }
}

@include media-breakpoint-up(xxl) {
  .que-image {
    width: 20%;
  }

  .do-you-know-text {
    margin-bottom: 6%;
  }
}
</style>
