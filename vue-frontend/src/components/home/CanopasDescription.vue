<template>
  <div class="container-canopas-description">
    <div class="mt-4 canopas-text-div normal-text">
      <h4 class="header-text canopas-gradient-text mb-4">What is Canopas?</h4>

      Canopas is a company helping Entrepreneurs, startups, and small companies
      to bring their IDEA to LIFE by developing digital products(or improving
      existing products) for their business.
      <br /><br />
      Our vision is
      <b
        ><i
          ><q
            >To put humanity in better condition compared to what it was before
            our presence in the world.</q
          ></i
        ></b
      >
      <br /><br />
      We're quality-focused people who want to work on a few handful apps that
      can be used by many people worldwide.
    </div>
    <BongoCat class="px-5 bongo-cat-div" />
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

<style scoped>
.container-canopas-description {
  display: flex;
  flex-direction: column-reverse;
  margin: 5% 10%;
  overflow: hidden;
  align-items: center;
}

.bongo-cat-div {
  width: 100%;
  height: 100%;
}

@media (min-width: 768px) {
  .header-text {
    line-height: 2rem;
    font-size: 1.875rem;
  }

  .normal-text {
    font-size: 1.125rem;
    line-height: 1.375rem;
  }

  .container-canopas-description {
    flex-direction: row;
    margin: 5% 0 5% 15%;
  }

  .canopas-text-div {
    flex: 1 1 0;
  }

  .bongo-cat-div {
    flex: 1 1 0;
    margin-right: -10%;
  }
}

@media (min-width: 992px) {
  .header-text {
    font-size: 3rem;
    line-height: 3.6rem;
  }

  .normal-text {
    font-size: 1.5rem;
    line-height: 1.8rem;
  }
}

@media (min-width: 1280px) {
  .bongo-cat-div {
    flex: 1.4 1 0;
  }
}
</style>
