<template>
  <div class="do-you-know-div">
    <div class="flex-row que-section">
      <img :src="doYouKnowImage" class="image que-image" alt="do-you-know" />
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
.que-section {
  margin: 15% 10%;
}

.flex-row {
  display: flex;
  align-items: center;
}

.flex-row > .image {
  display: none;
}

.pink-text {
  color: #f2709c;
}

.do-you-know-text {
  width: 100%;
  margin: 0;
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

  .flex-row {
    margin: 0 10%;
  }

  .flex-row > .image {
    display: block;
    width: 190px;
  }

  .do-you-know-text {
    width: 50%;
    margin: -160px auto 0;
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

  .flex-row {
    margin: 0 12%;
  }

  .flex-row > .image {
    display: block;
    width: 250px;
  }
}
</style>
