<template>
  <div class="do-you-know-div">
    <div class="flex-row">
      <div class="que-image">
        <img :src="doYouKnowImage" class="image" alt="do-you-know" />
      </div>
      <div class="text-div do-you-know-text">
        <h4 class="header-text pb-30">Do you know ?</h4>

        <span class="normal-text pink-text mt-20">
          We’ve worked with just <b>10 clients</b> in the last <b>5 years</b>.
          WHY?
        </span>
        <br /><br />
        <span class="normal-text mt-20">
          Because 8 of them have been offering(committed to) us repeat business
          since we started working for them. Think for a moment, Isn't it the
          main reason why you're here.</span
        >
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
};
</script>

<style lang="scss" scoped>
.do-you-know-div {
  width: 80%;
  margin: auto;
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

.do-you-know-text {
  flex: 75%;
  margin: 10% 0;
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

  .que-image {
    display: block;
    width: 25%;
  }

  .do-you-know-text {
    padding-left: 10%;
    margin: 10% 0 28%;
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

  .do-you-know-div {
    width: 70%;
  }
}

@include media-breakpoint-up(xl) {
  .que-image {
    width: 24%;
  }

  .do-you-know-text {
    margin: 0 0 15% 14%;
    padding-left: 0%;
  }
}

@include media-breakpoint-up(xxl) {
  .que-image {
    width: 20%;
  }

  .do-you-know-text {
    margin: 0 0 6% 14%;
  }
}
</style>
