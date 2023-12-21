<template>
  <div class="md:!mb-0 flex flex-col 3xl:outer-container">
    <div class="container mb-2.5 flex flex-col text-center md:mb-20">
      <h2
        class="mobile-header-2 lg:desk-header-2 mb-2.5 text-black-87 md:!mx-0 md:!w-full xs:mx-auto xs:w-[71%]"
      >
        Case studies
      </h2>
      <span
        class="sub-h1-regular lg:mobile-header-2-regular text-black-60 w-11/12 2xl:w-9/12 mx-auto"
        >Explore our case studies to witness how we've transformed our client's
        ideas into successful iOS apps.</span
      >
    </div>
    <!-- Mobile UI -->
    <div class="lg:hidden flex flex-col mb-24 md:mb-36">
      <div
        v-for="(portfolio, index) in portfolios"
        :key="index"
        class="flex flex-row mt-14"
        :class="portfolio.background"
      >
        <div
          class="basis-[50%] flex"
          :class="index == 0 || index == 2 ? 'items-end justify-end -mr-4' : ''"
        >
          <img
            :src="portfolio.image[0]"
            :srcset="`${portfolio.image[0]} 400w`"
            :alt="portfolio.title"
            class="w-[200px] h-[350px] object-cover"
            loading="lazy"
            :class="index == 1 ? 'mt-32 -ml-4' : ''"
          />
        </div>
        <div
          class="flex flex-col basis-[50%]"
          :class="index == 1 ? 'items-end basis-[70%]' : ''"
        >
          <img
            :src="portfolio.image[1]"
            :srcset="`${portfolio.image[1]} 400w`"
            :alt="portfolio.title"
            class="w-[200px] h-[350px] object-cover"
            loading="lazy"
          />
          <div
            class="w-[94%] sm:w-9/12 flex flex-col ml-[10%] pr-3.5 xs:p-3"
            :class="index == 1 ? 'pl-3.5 !w-[91%] sm:!w-[66%]' : ''"
          >
            <h2 class="mb-2.5 mobile-header-2">
              {{ portfolio.title }}
            </h2>
            <span class="sub-h1-regular text-black-60">{{
              portfolio.description
            }}</span>
          </div>
        </div>
      </div>
    </div>
    <!-- Mobile UI END-->
    <!-- Desktop UI -->

    <div id="stickyParent" class="relative sticky-parent hidden lg:block">
      <div class="sticky top-0 w-full h-screen">
        <swiper
          :direction="'vertical'"
          :slidesPerView="1"
          :centeredSlides="true"
          :spaceBetween="0"
          :speed="500"
          :mousewheel="{
            enabled: false,
            releaseOnEdges: true,
          }"
          :modules="modules"
          :allowTouchMove="false"
          class="swiper-container h-screen"
          @swiper="setSwiperRef"
        >
          <swiper-slide v-for="(portfolio, index) in portfolios" :key="index">
            <div class="flex flex-row justify-center h-screen">
              <div
                :class="portfolio.row1Background"
                class="flex flex-col justify-center !basis-[40%] justify-around px-8 text-white"
              >
                <div class="flex justify-end relative">
                  <div
                    :class="portfolio.row1Background"
                    class="absolute top-4 xl:top-8 right-[-25%] xl:right-[-17%] 2xll:right-[-18%] xll:right-[-12%] 3xl:right-[-13%] w-[250px] h-[250px] p-8 xl:w-[300px] xl:h-[300px] rounded-full flex items-center"
                  ></div>
                  <h2 class="absolute top-32 desk-header-1 title right-[-30px]">
                    {{ portfolio.title }}
                  </h2>
                </div>
                <div
                  class="w-[19rem] xl:w-[33rem] mx-auto mt-44 xl:mt-72 description"
                >
                  <h2 class="desk-header-2">
                    {{ portfolio.desktopDescription }}
                  </h2>
                  <div class="mt-6">
                    <span class="mobile-header-3-medium text-white-core-80">{{
                      portfolio.subDescription
                    }}</span>
                  </div>
                </div>
              </div>

              <div
                class="flex flex-row !basis-[60%] items-center p-12 2xll:p-20 bg-white"
                :class="portfolio.row2Background"
              >
                <div class="basis-[50%] mt-48 flex items-end justify-end">
                  <img
                    :src="portfolio.image[2]"
                    :srcset="`${portfolio.image[2]} 800w`"
                    :alt="portfolio.title"
                    class="image w-[284px] h-[490px] xl:h-[590px] object-cover"
                    loading="lazy"
                  />
                </div>
                <div class="hidden md:block relative">
                  <div class="flex flex-col -mt-32 basis-[50%]">
                    <img
                      :src="portfolio.image[3]"
                      :srcset="`${portfolio.image[3]} 800w`"
                      :alt="portfolio.title"
                      class="image main w-[284px] h-[490px] xl:h-[590px] object-cover"
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
      lastScrollY: 0,
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
  },
  unmounted() {
    document.removeEventListener("scroll", this.handleScroll);
  },
  methods: {
    setSwiperRef(swiper) {
      this.swiperRef = swiper;
    },

    handleScroll() {
      const stickyParent = document.getElementById("stickyParent");
      const stickyTop = stickyParent.offsetTop;
      var position = stickyParent.getBoundingClientRect();

      if (
        window.pageYOffset >= stickyTop &&
        position.bottom >=
          (window.innerHeight || document.documentElement.clientHeight)
      ) {
        this.swiperRef.mousewheel.enable();
        this.$emit("sectionFullscreen", false);
      } else {
        this.swiperRef.mousewheel.disable();
        this.$emit("sectionFullscreen", true);
      }
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
.sticky-parent {
  @apply h-[300vh] 3xl:!h-[220vh];
}
.swiper-container {
  @apply mx-0;
}
.swiper-slide-active .image,
.swiper-slide-active .title,
.swiper-slide-active .description {
  @apply animate-moveUp;
}
</style>
