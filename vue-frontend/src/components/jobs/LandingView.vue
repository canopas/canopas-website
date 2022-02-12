<template>
  <div class="text-center container">
    <div class="header-2-text">
      Are you looking for a job where you can leverage your knowledge and
      experience to build digital products that will be used by millions of
      customers.
    </div>
    <div class="normal-text desc-text">
      You've come to the right place. We're looking for an AVENGER who can join
      our existing team of AVENGERS.
    </div>
    <button class="gradient-btn normal-text">
      <span>SEE ALL POSITIONS</span>
    </button>
  </div>
</template>


<style lang="scss" scoped>
.container {
  max-width: 687px;
  margin: 30px auto 75px;
  padding: 0 20px;
}

.desc-text {
  margin: 30px 0;
}

@include media-breakpoint-up(sm) {
  .gradient-btn {
    margin: auto;
  }

  .container {
    padding: 0 100px;
  }
}

@include media-breakpoint-up(lg) {
  .header-text {
    font-size: 2.5rem;
    line-height: 3rem;
  }

  .desc-text {
    margin: 50px 0;
  }

  .container {
    padding: 0;
  }
}
</style>