<template>
  <div
    class="my-16 xs:my-24 md:mt-24 xl:!mt-60 md:!mb-0 flex flex-col 3xl:container"
  >
    <div class="container mb-2.5 flex flex-col text-center md:mb-20">
      <h2
        class="tracking-[-0.04rem] mb-2.5 font-opensans-bold text-2xl leading-9 text-black-core/[0.87] md:!mx-0 md:!w-full md:text-[2.46875rem] lg:text-[3.4375rem] md:leading-[3.359375rem] lg:leading-[4.46875rem] xs:mx-auto xs:w-[71%]"
      >
        Case studies
      </h2>
      <span
        class="font-inter-regular text-lg md:text-xl leading-[1.6875rem] md:leading-[1.875rem] text-black-core/[0.60] md:font-inter-medium lg:text-2xl lg:leading-9 w-11/12 2xl:w-9/12 mx-auto"
        >Explore our case studies to witness how we've transformed our client's
        ideas into successful iOS apps.</span
      >
    </div>
    <!-- Mobile UI -->
    <div class="md:hidden flex flex-col">
      <div
        v-for="(portfolio, index) in portfolios"
        :key="index"
        class="flex flex-row mt-14"
        :class="portfolio.background"
      >
        <div
          class="basis-[50%] flex"
          :class="
            index == 0 || index == 2 ? 'items-end justify-end xs:-mr-6' : ''
          "
        >
          <img
            :src="portfolio.image[0]"
            :srcset="`${portfolio.image[0]} 400w`"
            :alt="portfolio.title"
            class="w-[137px] xs:w-[200px] xs:h-[350px] h-[297px] object-cover"
            loading="lazy"
            :class="index == 1 ? 'mt-32 xs:-ml-6' : ''"
          />
        </div>
        <div
          class="flex flex-col basis-[50%]"
          :class="index == 1 ? 'items-end' : ''"
        >
          <img
            :src="portfolio.image[1]"
            :srcset="`${portfolio.image[1]} 400w`"
            :alt="portfolio.title"
            class="w-[137px] xs:w-[200px] xs:h-[350px] h-[297px] object-cover"
            loading="lazy"
          />
          <div
            class="w-11/12 sm:w-9/12 flex flex-col ml-[10%] pr-3.5 xs:p-3"
            :class="index == 1 ? 'pl-3.5 w-full' : ''"
          >
            <h2
              class="tracking-[0.04rem] mb-2.5 font-opensans-bold text-2xl leading-9 text-black-core/[0.87]"
            >
              {{ portfolio.title }}
            </h2>
            <span
              class="font-inter-regular text-lg leading-[1.6875rem] text-black-core/[0.60]"
              >{{ portfolio.description }}</span
            >
          </div>
        </div>
      </div>
    </div>
    <!-- Mobile UI END-->
    <!-- Desktop UI -->

    <div
      id="stickyParent"
      class="sticky-parent h-[300vh] xll:h-[250vh] 3xl:!h-[220vh] hidden md:block"
    >
      <div
        class="sticky sticky top-0 max-h-full main overflow-hidden"
        :class="{ '!top-[75px] xl:!top-[40px]': isScrollingUp }"
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
          class="swiper-container h-screen"
          @swiper="setSwiperRef"
          @slideChange="onSlideChange"
        >
          <swiper-slide v-for="(portfolio, index) in portfolios" :key="index">
            <div class="hidden md:flex flex-row justify-center h-screen">
              <div
                :class="portfolio.row1Background"
                class="flex flex-col justify-center !basis-[40%] justify-around px-8"
              >
                <div class="flex justify-end relative">
                  <div
                    :class="portfolio.row1Background"
                    class="absolute top-8 lg:top-4 xl:top-8 right-[-30%] lg:right-[-25%] xl:right-[-17%] 2xll:right-[-18%] xll:right-[-12%] 3xl:right-[-18%] font-opensans-bold text-2xl leading-[1.95rem] lg:text-5xl lg:leading-[3.9rem] text-[#FFFFFF] tracking-[-0.04rem] w-40 h-40 lg:w-[250px] lg:h-[250px] p-8 xl:w-[300px] xl:h-[300px] rounded-full flex justify-center items-center"
                  >
                    <h2 class="title">
                      {{ portfolio.title }}
                    </h2>
                  </div>
                </div>
                <div
                  class="w-[17rem] lg:w-[19rem] xl:w-[33rem] mx-auto mt-32 lg:mt-44 xl:mt-72 description"
                  :class="{ '!mt-12 xl:!mt-72': isScrollingUp }"
                >
                  <h2
                    class="font-opensans-bold text-2xl leading-[1.95rem] lg:text-5xl lg:leading-[3.9rem] text-[#FFFFFF] tracking-[-0.04rem]"
                  >
                    {{ portfolio.desktopDescription }}
                  </h2>
                  <div class="mt-6">
                    <span
                      class="font-inter-medium text-[#FFFFFF]/[0.80] text-base lg:text-xl lg:leading-[1.875rem]"
                      >{{ portfolio.subDescription }}</span
                    >
                  </div>
                </div>
                <div
                  class="flex justify-end pr-[8%] 2xll:pr-[14%] cursor-pointer"
                  :class="{ 'mt-[-45%] xl:mt-0': isScrollingUp }"
                >
                  <span
                    @click="scrollToNext()"
                    class="font-inter-medium text-[#FFFFFF]/[0.80] text-base lg:text-xl lg:leading-[1.875rem] mt-6"
                    >SKIP</span
                  >
                </div>
              </div>

              <div
                class="flex flex-row !basis-[60%] items-center p-8 lg:p-12 2xll:p-20 bg-white"
                :class="portfolio.row2Background"
              >
                <div class="basis-[50%] mt-48 flex items-end justify-end">
                  <img
                    :src="portfolio.image[2]"
                    :srcset="`${portfolio.image[2]} 800w`"
                    :alt="portfolio.title"
                    class="image w-[175px] h-[350px] lg:w-[284px] lg:h-[490px] xl:h-[590px] object-cover"
                    loading="lazy"
                  />
                </div>
                <div class="hidden md:block relative">
                  <div class="flex flex-col -mt-32 basis-[50%]">
                    <img
                      :src="portfolio.image[3]"
                      :srcset="`${portfolio.image[3]} 800w`"
                      :alt="portfolio.title"
                      class="image w-[175px]main h-[350px] lg:w-[284px] lg:h-[490px] xl:h-[590px] object-cover"
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
          background: "bg-justly-pink-gradient-background",
          row1Background: "bg-[#ED722D]",
          row2Background: "bg-justly-radical-gradient",
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
            "flex-row-reverse items-center bg-luxeradio-yellow-gradient-background",
          row1Background: "bg-[#101010]",
          row2Background: "bg-luxeradio-radical-gradient",
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
          background: "bg-togness-blue-gradient-background",
          row1Background: "bg-[#125EE3]",
          row2Background: "bg-togness-radical-gradient",
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
  @apply mx-0;
}
.swiper-slide-active .image,
.swiper-slide-active .title,
.swiper-slide-active .description {
  @apply animate-moveUp;
}
</style>
