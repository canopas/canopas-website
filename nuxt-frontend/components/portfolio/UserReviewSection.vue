<template>
  <section
    class="bg-white md:pt-36 md:pb-0 lg:pt-50"
    :ref="response.review ? response.review.ref : ''"
  >
    <div class="header-text text-center">
      Listen what {{ name.charAt(0).toUpperCase() + name.slice(1) }} users
      <br />
      say about us...
    </div>
    <div class="mt-8 text-left overflow-x-hidden overflow-y-auto scrollbar-0">
      <div
        class="animate-gridAnimationReverse"
        :style="{
          animationPlayState: `${pausedId == 0 ? gridAnimation : 'paused'}`,
        }"
      >
        <div
          class="animate-gridAnimation"
          :style="{
            animationPlayState: `${pausedId == 0 ? 'running' : 'paused'}`,
          }"
        >
          <div class="flex ml-40" v-if="grid1.length > 0">
            <div
              v-for="element in grid1"
              :key="element.id"
              class="flex flex-col justify-center flex-[0_0_320px] sm:flex-[0_0_480px] lg:flex-[0_0_528px] mt-8 ml-8 rounded-2xl border border-solid border-transparent shadow-[inset_2px_1000px_1px_#fff] from-[#ff9472] via-[#ff909c] to-[#f2709c] bg-gradient-to-r p-[16px] lg:p-[20px] text-[1rem] leading-[1.125rem] md:text-[1.0625rem] md:leading-6 lg:text-[1.1875rem] lg:leading-[1.875rem] text-left"
              :class="
                pausedId == element.id ? 'scale-[0.97] cursor-pointer' : ''
              "
              @mouseover="pausedId = element.id"
              @mouseleave="pausedId = 0"
              @touchstart.passive="pausedId = element.id"
              @touchend="pausedId = 0"
              :ref="'card-1-' + element.id"
            >
              {{ element.review }}
              <div class="flex flex-row">
                <Icon
                  v-for="(item, index) in 5"
                  :key="index"
                  name="fa6-solid:star"
                  class="mt-3 w-5 h-5 text-yellow-500"
                />
              </div>
            </div>
          </div>
          <div class="flex mb-2.5" v-if="grid2.length > 0">
            <div
              v-for="element in grid2"
              :key="element.id"
              class="flex flex-col justify-center flex-[0_0_320px] sm:flex-[0_0_480px] lg:flex-[0_0_528px] mt-8 ml-8 rounded-2xl border border-solid border-transparent shadow-[inset_2px_1000px_1px_#fff] from-[#ff9472] via-[#ff909c] to-[#f2709c] bg-gradient-to-r p-4 lg:p-5 text-[1rem] leading-[1.125rem] md:text-[1.0625rem] md:leading-[1.5rem] lg:text-[1.1875rem] lg:leading-[1.875rem] text-left"
              :class="
                pausedId == element.id ? 'scale-[0.97] cursor-pointer' : ''
              "
              @mouseover="pausedId = element.id"
              @mouseleave="pausedId = 0"
              @touchstart.passive="pausedId = element.id"
              @touchend="pausedId = 0"
              :ref="'card-1-' + element.id"
            >
              {{ element.review }}
              <div class="flex flex-row">
                <Icon
                  v-for="(item, index) in 5"
                  :key="index"
                  name="fa6-solid:star"
                  class="mt-3 w-5 h-5 text-yellow-500"
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
export default {
  props: ["json"],
  data() {
    return {
      name: this.$route.params.id,
      response: this.json,
      grid1: [],
      grid2: [],
      portfolioRef: null,
      gridAnimation: "paused",
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
        this.gridAnimation = "running";
      } else {
        this.gridAnimation = "paused";
      }
      this.lastScrollY = scrollTop;
    },
    setReviewGrid() {
      const length = this.response.usersReviews.length;
      this.grid1 = new Array(3)
        .fill(this.response.usersReviews.slice(0, length / 2))
        .flat();
      this.grid2 = new Array(3)
        .fill(this.response.usersReviews.slice(length / 2, length))
        .flat();
    },
    appendKeyframes() {
      const width =
        ((this.$refs["card-1-1"][0].offsetWidth + 32) * // 32 is left margin of each cards
          this.response.usersReviews.length) /
        2;

      const head = document.getElementsByTagName("head")[0];
      const styles = document.getElementsByTagName("style");
      let found = false;

      // find style titled with portfolio, if found replace width according to data
      for (const element of styles) {
        if (element.title == "portfolio") {
          found = true;
          element.innerText = element.innerText.replaceAll(
            this.lastWidth,
            width,
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
