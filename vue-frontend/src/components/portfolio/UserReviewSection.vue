<template>
  <section
    class="tw-bg-white md:tw-pt-36 md:tw-pb-0 lg:tw-pt-50"
    :ref="response.review ? response.review.ref : ''"
  >
    <div class="header-text tw-text-center">
      Listen what {{ name.charAt(0).toUpperCase() + name.slice(1) }} users
      <br />
      say about us...
    </div>
    <div
      class="tw-mt-[32px] tw-text-left tw-overflow-x-hidden tw-overflow-y-auto tw-scrollbar-0"
    >
      <div
        class="tw-animate-gridAnimationReverse"
        :class="[pausedId == 0 ? gridAnimation : 'animation-paused']"
      >
        <div
          class="tw-animate-gridAnimation"
          :class="[pausedId == 0 ? 'animation-running' : 'animation-paused']"
        >
          <div class="tw-flex tw-ml-[160px]" v-if="grid1.length > 0">
            <div
              v-for="element in grid1"
              :key="element.id"
              class="tw-flex tw-flex-col tw-justify-center tw-flex-[0_0_320px] sm:tw-flex-[0_0_480px] lg:tw-flex-[0_0_528px] tw-mt-[32px] tw-ml-[32px] tw-rounded-2xl tw-border-[1px] tw-border-solid tw-border-transparent tw-shadow-[inset_2px_1000px_1px_#fff] tw-from-[#ff9472] tw-via-[#ff909c] tw-to-[#f2709c] tw-bg-gradient-to-r tw-p-[16px] lg:tw-p-[20px] tw-text-[1rem] tw-leading-[1.125rem] md:tw-text-[1.0625rem] md:tw-leading-[1.5rem] lg:tw-text-[1.1875rem] lg:tw-leading-[1.875rem] tw-text-left"
              :class="
                pausedId == element.id
                  ? 'tw-scale-[0.97] tw-cursor-pointer'
                  : ''
              "
              @mouseover="pausedId = element.id"
              @mouseleave="pausedId = 0"
              @touchstart.passive="pausedId = element.id"
              @touchend="pausedId = 0"
              :ref="'card-1-' + element.id"
            >
              {{ element.review }}
              <div class="tw-flex tw-flex-row">
                <font-awesome-icon
                  v-for="(item, index) in 5"
                  :key="index"
                  icon="star"
                  class="tw-mt-3 tw-w-5 tw-h-5 tw-text-yellow-500"
                />
              </div>
            </div>
          </div>
          <div class="tw-flex tw-mb-[10px]" v-if="grid2.length > 0">
            <div
              v-for="element in grid2"
              :key="element.id"
              class="tw-flex tw-flex-col tw-justify-center tw-flex-[0_0_320px] sm:tw-flex-[0_0_480px] lg:tw-flex-[0_0_528px] tw-mt-[32px] tw-ml-[32px] tw-rounded-2xl tw-border-[1px] tw-border-solid tw-border-transparent tw-shadow-[inset_2px_1000px_1px_#fff] tw-from-[#ff9472] tw-via-[#ff909c] tw-to-[#f2709c] tw-bg-gradient-to-r tw-p-[16px] lg:tw-p-[20px] tw-text-[1rem] tw-leading-[1.125rem] md:tw-text-[1.0625rem] md:tw-leading-[1.5rem] lg:tw-text-[1.1875rem] lg:tw-leading-[1.875rem] tw-text-left"
              :class="
                pausedId == element.id
                  ? 'tw-scale-[0.97] tw-cursor-pointer'
                  : ''
              "
              @mouseover="pausedId = element.id"
              @mouseleave="pausedId = 0"
              @touchstart.passive="pausedId = element.id"
              @touchend="pausedId = 0"
              :ref="'card-1-' + element.id"
            >
              {{ element.review }}
              <div class="tw-flex tw-flex-row">
                <font-awesome-icon
                  v-for="(item, index) in 5"
                  :key="index"
                  icon="star"
                  class="tw-mt-3 tw-w-5 tw-h-5 tw-text-yellow-500"
                />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

export default {
  props: ["json"],
  data() {
    return {
      name: this.$route.params.id,
      response: this.json,
      grid1: [],
      grid2: [],
      portfolioRef: null,
      gridAnimation: "animation-paused",
      pausedId: 0,
      lastWidth: null,
      lastScrollY: 0,
      event: "",
    };
  },
  watch: {
    json: function (newVal, oldVal) {
      this.grid1 = [];
      this.grid2 = [];
      this.name = this.$route.params.id;
      this.response = newVal;
      this.appendKeyframes();
      setTimeout(() => {
        this.setReviewGrid();
      }, 100);
    },
  },
  components: {
    FontAwesomeIcon,
  },
  created() {
    this.setReviewGrid();
  },
  mounted() {
    this.appendKeyframes();
    this.portfolioRef = document.getElementById("response");
    this.portfolioRef.addEventListener("scroll", this.handleScroll);
    window.addEventListener("scroll", this.handleScroll);
  },
  unmounted() {
    this.portfolioRef.removeEventListener("scroll", this.handleScroll);
    window.removeEventListener("scroll", this.handleScroll);
  },
  methods: {
    handleScroll() {
      let scrollTop =
        window.innerWidth < 992 ? window.scrollY : this.portfolioRef.scrollTop;
      if (scrollTop < this.lastScrollY) {
        this.gridAnimation = "animation-running";
      } else {
        this.gridAnimation = "animation-paused";
      }
      this.lastScrollY = scrollTop;
    },
    setReviewGrid() {
      var length = this.response.usersReviews.length;
      this.grid1 = new Array(3)
        .fill(this.response.usersReviews.slice(0, length / 2))
        .flat();
      this.grid2 = new Array(3)
        .fill(this.response.usersReviews.slice(length / 2, length))
        .flat();
    },
    appendKeyframes() {
      let width =
        ((this.$refs["card-1-1"][0].offsetWidth + 32) * // 32 is left margin of each cards
          this.response.usersReviews.length) /
        2;

      const head = document.getElementsByTagName("head")[0];
      const styles = document.getElementsByTagName("style");
      var found = false;

      // find style titled with portfolio, if found replace width according to data
      for (var i = 0; i < styles.length; i++) {
        if (styles[i].title == "portfolio") {
          found = true;
          styles[i].innerText = styles[i].innerText.replaceAll(
            this.lastWidth,
            width
          );
        }
      }

      // else append keyframes
      if (!found) {
        let keyframes =
          `@keyframes scroll { 0% {transform: translateX(0px)} 100% {transform:translateX(-` +
          width +
          `px)}}\n` +
          `@-webkit-keyframes scroll { 0% {transform: translateX(0px)} 100% {transform:translateX(-` +
          width +
          `px)}}\n` +
          `@-webkit-keyframes scroll-reverse { 0% {transform:translateX(-` +
          width +
          `px)} 100% {transform: translateX(0px)}}\n` +
          `@keyframes scroll-reverse { 0% {transform:translateX(-` +
          width +
          `px)} 100% {transform: translateX(0px)}}`;

        let style = document.createElement("style");
        style.title = "portfolio";
        style.innerHTML = keyframes;

        head.appendChild(style);
      }

      this.lastWidth = width;
    },
  },
};
</script>
