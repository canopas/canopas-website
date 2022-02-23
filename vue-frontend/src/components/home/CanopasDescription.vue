<template>
  <div class="container-fluid">
    <div class="container">
      <div class="mt-4 canopas-text-div normal-text">
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
        <div class="bongo-cat-div text-right">
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
  method: {},
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
  .header-text {
    line-height: 2rem;
    font-size: 1.875rem;
  }

  .normal-text {
    font-size: 1.125rem;
    line-height: 1.375rem;
  }

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

@include media-breakpoint-up(lg) {
  .header-text {
    font-size: 3rem;
    line-height: 3.6rem;
  }

  .normal-text {
    font-size: 1.5rem;
    line-height: 1.8rem;
  }
}

@include media-breakpoint-up(xl) {
  .bongo-cat-wrapper {
    flex: 1.3 1 0;
  }
}
</style>
