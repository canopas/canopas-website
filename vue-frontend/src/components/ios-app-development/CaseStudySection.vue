<template>
  <div
    class="tw-mt-16 xs:tw-mt-24 md:tw-mt-24 xl:!tw-mt-60 md:!tw-mb-0 tw-flex tw-flex-col 3xl:tw-outer-container"
  >
    <div
      class="tw-container tw-mb-2.5 tw-flex tw-flex-col tw-text-center md:tw-mb-20"
    >
      <h2
        class="header-2 tw-mb-2.5 tw-text-black-core/[0.87] md:!tw-mx-0 md:!tw-w-full xs:tw-mx-auto xs:tw-w-[71%]"
      >
        Case studies
      </h2>
      <span
        class="sub-h1-regular tw-text-black-core/[0.60] tw-w-11/12 2xl:tw-w-9/12 tw-mx-auto"
        >Explore our case studies to witness how we've transformed our client's
        ideas into successful iOS apps.</span
      >
    </div>
    <!-- Mobile UI -->
    <div class="lg:tw-hidden tw-flex tw-flex-col tw-mb-24 md:tw-mb-36">
      <div
        v-for="(portfolio, index) in portfolios"
        :key="index"
        class="tw-flex tw-flex-row tw-mt-14"
        :class="portfolio.background"
      >
        <div
          class="tw-basis-[50%] tw-flex"
          :class="
            index == 0 || index == 2
              ? 'tw-items-end tw-justify-end xs:-tw-mr-6'
              : ''
          "
        >
          <img
            :src="portfolio.image[0]"
            :srcset="`${portfolio.image[0]} 400w`"
            :alt="portfolio.title"
            class="tw-w-[137px] xs:tw-w-[200px] xs:tw-h-[350px] tw-h-[297px] tw-object-cover"
            loading="lazy"
            :class="index == 1 ? 'tw-mt-32 xs:-tw-ml-6' : ''"
          />
        </div>
        <div
          class="tw-flex tw-flex-col tw-basis-[50%]"
          :class="index == 1 ? 'tw-items-end' : ''"
        >
          <img
            :src="portfolio.image[1]"
            :srcset="`${portfolio.image[1]} 400w`"
            :alt="portfolio.title"
            class="tw-w-[137px] xs:tw-w-[200px] xs:tw-h-[350px] tw-h-[297px] tw-object-cover"
            loading="lazy"
          />
          <div
            class="tw-w-11/12 sm:tw-w-9/12 tw-flex tw-flex-col tw-ml-[10%] tw-pr-3.5 xs:tw-p-3"
            :class="index == 1 ? 'tw-pl-3.5 tw-w-full' : ''"
          >
            <h2 class="tw-mb-2.5 header-2">
              {{ portfolio.title }}
            </h2>
            <span class="sub-h1-regular tw-text-black-core/[0.60]">{{
              portfolio.description
            }}</span>
          </div>
        </div>
      </div>
    </div>
    <!-- Mobile UI END-->
    <!-- Desktop UI -->

    <div
      id="stickyParent"
      class="sticky-parent tw-h-[300vh] xll:tw-h-[240vh] 3xl:!tw-h-[220vh] tw-hidden lg:tw-block"
    >
      <div
        class="sticky tw-sticky tw-top-0 tw-max-h-full main tw-overflow-hidden"
        :class="{ '!tw-top-[75px] xl:!tw-top-[40px]': isScrollingUp }"
      >
        <swiper
          :direction="'vertical'"
          :slidesPerView="1"
          :centeredSlides="true"
          :spaceBetween="0"
          :speed="500"
          :mousewheel="{
            enabled: false,
            releaseOnEdges: true,
            sensitivity: 1,
            thresholdDelta: 1,
          }"
          :touchReleaseOnEdges="true"
          :modules="modules"
          :allowTouchMove="false"
          class="swiper-container tw-h-screen"
          @swiper="setSwiperRef"
          @slideChange="onSlideChange"
        >
          <swiper-slide v-for="(portfolio, index) in portfolios" :key="index">
            <div
              class="tw-hidden md:tw-flex tw-flex-row tw-justify-center tw-h-screen"
            >
              <div
                :class="portfolio.row1Background"
                class="tw-flex tw-flex-col tw-justify-center !tw-basis-[40%] tw-justify-around tw-px-8"
              >
                <div class="tw-flex tw-justify-end tw-relative">
                  <div
                    :class="portfolio.row1Background"
                    class="tw-absolute tw-top-8 lg:tw-top-4 xl:tw-top-8 tw-right-[-30%] lg:tw-right-[-25%] xl:tw-right-[-17%] 2xll:tw-right-[-18%] xll:tw-right-[-12%] 3xl:tw-right-[-13%] tw-font-opensans-bold tw-text-2xl tw-leading-[1.95rem] lg:tw-text-5xl lg:tw-leading-[3.9rem] tw-text-[#FFFFFF] tw-tracking-[-0.04rem] tw-w-40 tw-h-40 lg:tw-w-[250px] lg:tw-h-[250px] tw-p-8 xl:tw-w-[300px] xl:tw-h-[300px] tw-rounded-full tw-flex tw-justify-center tw-items-center"
                  >
                    <h2 class="header-1 title">
                      {{ portfolio.title }}
                    </h2>
                  </div>
                </div>
                <div
                  class="tw-w-[17rem] lg:tw-w-[19rem] xl:tw-w-[33rem] tw-mx-auto tw-mt-32 lg:tw-mt-44 xl:tw-mt-72 description"
                  :class="{ '!tw-mt-12 xl:!tw-mt-72': isScrollingUp }"
                >
                  <h2 class="header-2 tw-text-[#FFFFFF]">
                    {{ portfolio.desktopDescription }}
                  </h2>
                  <div class="tw-mt-6">
                    <span class="sub-h3-semibold tw-text-[#FFFFFF]/[0.80]">{{
                      portfolio.subDescription
                    }}</span>
                  </div>
                </div>
                <div
                  class="tw-flex tw-justify-end tw-pr-[8%] 2xll:tw-pr-[14%] tw-cursor-pointer"
                  :class="{ 'tw-mt-[-45%] xl:tw-mt-0': isScrollingUp }"
                >
                  <span
                    @click="scrollToNext()"
                    class="tw-text-[#FFFFFF]/[0.80] tw-mt-6 sub-h3-semibold"
                    >SKIP</span
                  >
                </div>
              </div>

              <div
                class="tw-flex tw-flex-row !tw-basis-[60%] tw-items-center tw-p-8 lg:tw-p-12 2xll:tw-p-20 tw-bg-white"
                :class="portfolio.row2Background"
              >
                <div
                  class="tw-basis-[50%] tw-mt-48 tw-flex tw-items-end tw-justify-end"
                >
                  <img
                    :src="portfolio.image[2]"
                    :srcset="`${portfolio.image[2]} 800w`"
                    :alt="portfolio.title"
                    class="image tw-w-[175px] tw-h-[350px] lg:tw-w-[284px] lg:tw-h-[490px] xl:tw-h-[590px] tw-object-cover"
                    loading="lazy"
                  />
                </div>
                <div class="tw-hidden md:tw-block tw-relative">
                  <div class="tw-flex tw-flex-col tw--mt-32 tw-basis-[50%]">
                    <img
                      :src="portfolio.image[3]"
                      :srcset="`${portfolio.image[3]} 800w`"
                      :alt="portfolio.title"
                      class="image tw-w-[175px]main tw-h-[350px] lg:tw-w-[284px] lg:tw-h-[490px] xl:tw-h-[590px] tw-object-cover"
                      loading="lazy"
                    />
                  </div>
                </div>
              </div>
            </div>
          </swiper-slide>
        </swiper>
      </div>
    </div>
  </div>
</template>

<script type="module">
import { Pagination, Autoplay, Mousewheel } from "swiper/modules";
import { Swiper, SwiperSlide } from "swiper/vue";
import mobile_justly1 from "@/assets/images/ios-app-development/casestudy/mobile/justly1.webp";
import mobile_justly2 from "@/assets/images/ios-app-development/casestudy/mobile/justly2.webp";
import mobile_luxeradio1 from "@/assets/images/ios-app-development/casestudy/mobile/luxeradio1.webp";
import mobile_luxeradio2 from "@/assets/images/ios-app-development/casestudy/mobile/luxeradio2.webp";
import mobile_togness1 from "@/assets/images/ios-app-development/casestudy/mobile/togness1.webp";
import mobile_togness2 from "@/assets/images/ios-app-development/casestudy/mobile/togness2.webp";

import desktop_justly1 from "@/assets/images/ios-app-development/casestudy/justly1-800w.webp";
import desktop_justly2 from "@/assets/images/ios-app-development/casestudy/justly2-800w.webp";
import desktop_luxeradio1 from "@/assets/images/ios-app-development/casestudy/luxe1-800w.webp";
import desktop_luxeradio2 from "@/assets/images/ios-app-development/casestudy/luxe2-800w.webp";
import desktop_togness1 from "@/assets/images/ios-app-development/casestudy/togness1-800w.webp";
import desktop_togness2 from "@/assets/images/ios-app-development/casestudy/togness2-800w.webp";
export default {
  data() {
    return {
      modules: [Pagination, Autoplay, Mousewheel],
      swiper: null,
      swiperRef: 0,
      activeIndex: 0,
      skipNext: true,
      lastTouchY: null,
      lastScrollY: 0,
      isScrollingUp: false,
      scrollPosition: 0,
      lastSlideReached: false,
      portfolios: [
        {
          title: "Justly",
          image: [
            mobile_justly1,
            mobile_justly2,
            desktop_justly1,
            desktop_justly2,
          ],
          description: "Create systems to build your dream life",
          background: "tw-bg-justly-pink-gradient-background",
          row1Background: "tw-bg-[#ED722D]",
          row2Background: "tw-bg-justly-radical-gradient",
          desktopDescription: "Overcome depression by building habits",
          subDescription:
            "Justly aims to tackle loneliness, depression, and mental health through innovative solutions.",
        },
        {
          title: "Luxeradio",
          image: [
            mobile_luxeradio1,
            mobile_luxeradio2,
            desktop_luxeradio1,
            desktop_luxeradio2,
          ],
          description: "New personalized radio and modern art gallery.",
          background:
            "tw-flex-row-reverse tw-items-center tw-bg-luxeradio-yellow-gradient-background",
          row1Background: "tw-bg-[#101010]",
          row2Background: "tw-bg-luxeradio-radical-gradient",
          desktopDescription: "A radio, music and podcast platform",
          subDescription:
            "Luxe Radio displays the best of Moroccan and international creation emphasizing taste , elegance.",
        },
        {
          title: "Togness",
          image: [
            mobile_togness1,
            mobile_togness2,
            desktop_togness1,
            desktop_togness2,
          ],
          description: "Easiest video editor & slideshow maker",
          background: "tw-bg-togness-blue-gradient-background",
          row1Background: "tw-bg-[#125EE3]",
          row2Background: "tw-bg-togness-radical-gradient",
          desktopDescription: "Photo editor and video maker app",
          subDescription:
            "Togness are create slideshows and add quotes and your own music for every special occassion of your life!",
        },
      ],
    };
  },
  mounted() {
    document.addEventListener("scroll", this.handleScroll);
    document.addEventListener("wheel", this.handleWheel);
    window.addEventListener("touchstart", this.handleTouchStart);
    window.addEventListener("touchmove", this.handleTouchMove);
    window.addEventListener("touchend", this.handleTouchEnd);
    window.addEventListener("scroll", this.handleScrollTop);
  },
  unmounted() {
    document.removeEventListener("wheel", this.handleWheel);
    window.removeEventListener("touchstart", this.handleTouchStart);
    window.removeEventListener("touchmove", this.handleTouchMove);
    window.removeEventListener("touchend", this.handleTouchEnd);
    window.removeEventListener("scroll", this.handleScrollTop);
    document.removeEventListener("scroll", this.handleScroll);
  },
  methods: {
    setSwiperRef(swiper) {
      this.swiperRef = swiper;
    },
    // handle mouseScroll event
    scrollDirectionIsUp(event) {
      if (event.wheelDelta) {
        return event.wheelDelta > 0;
      }
      return event.deltaY < 0;
    },

    handleWheel(event) {
      if (this.scrollDirectionIsUp(event)) {
        this.skipNext = false;
      } else {
        this.skipNext = true;
      }
    },
    //handleTouch event
    handleTouchStart(event) {
      this.lastTouchY = event.touches[0].clientY;
    },
    handleTouchMove(event) {
      if (this.lastTouchY !== null) {
        const currentTouchY = event.touches[0].clientY;
        if (currentTouchY > this.lastTouchY) {
          this.skipNext = false;
        } else if (currentTouchY < this.lastTouchY) {
          this.skipNext = true;
        }
        this.lastTouchY = currentTouchY;
      }
    },
    handleTouchEnd() {
      this.lastTouchY = null;
    },

    scrollToNext() {
      if (this.skipNext) {
        this.$emit("scroll-to-next");
      } else {
        this.$emit("scroll-to-previous");
      }
    },
    handleScrollTop() {
      const stickyParent = document.getElementById("stickyParent");
      const stickyTop = stickyParent.offsetTop;
      let scrollTop = window.scrollY;
      var position = stickyParent.getBoundingClientRect();

      if (
        window.pageYOffset >= stickyTop &&
        position.bottom >=
          (window.innerHeight || document.documentElement.clientHeight)
      ) {
        this.swiperRef.mousewheel.enable();
        this.swiperRef.allowTouchMove = true;
      } else {
        this.swiperRef.mousewheel.disable();
        this.swiperRef.allowTouchMove = false;
      }
      this.lastScrollY = scrollTop;
    },
    handleScroll() {
      const windowHeight = window.innerHeight;
      const lastSlide = document.querySelector(".main");
      const lastSlidePosition = lastSlide.getBoundingClientRect().top;
      if (lastSlidePosition <= windowHeight && lastSlidePosition >= 0) {
        this.lastSlideReached = true;
        this.isScrollingUp = window.scrollY < this.scrollPosition;
      } else {
        this.lastSlideReached = false;
      }
      this.scrollPosition = window.scrollY;
    },
  },
  components: {
    Swiper,
    SwiperSlide,
  },
};
</script>
<style lang="postcss">
@import "swiper/css";
@import "swiper/css/pagination";
.swiper-container {
  @apply tw-mx-0;
}
.swiper-slide-active .image,
.swiper-slide-active .title,
.swiper-slide-active .description {
  @apply tw-animate-moveUp;
}
</style>
