<template>
  <div class="mt-16 lg:mt-60 text-center 3xl:outer-container">
    <div
      class="container flex flex-col gap-4 lg:gap-6 pb-6 lg:pb-[4.5rem] lg:w-[51.625rem]"
    >
      <h2 class="mobile-header-2 lg:desk-header-2 text-black-87">
        Case studies
      </h2>
      <span class="sub-h1-regular lg:mobile-header-2-regular lg: text-black-60">
        Explore our case studies to witness how we've transformed our client's
        ideas into successful Flutter apps.
      </span>
    </div>
    <!-- Mobile UI -->
    <div class="lg:hidden text-left bg-[#F8F8F8]">
      <div v-for="(portfolio, index) in portfolios" :key="index">
        <img
          :src="portfolio.image[0]"
          :srcset="`${portfolio.image[0]} 400w,${portfolio.image[1]} 800w`"
          class="w-[414px] h-[412px] object-cover mx-auto"
          :alt="portfolio.title"
          loading="lazy"
        />
        <div class="pt-6 pb-14 relative text-white bg-[#EF233C]">
          <div
            class="absolute -top-1 left-1/2 -translate-x-1/2 -translate-y-1/2 w-0 h-0 border-transparent border-t-0 border-l-[13px] border-r-[13px] border-b-[14px] border-b-[#EF233C] border-b-solid"
          ></div>
          <div class="container flex flex-col gap-4">
            <h3 class="mobile-header-3-semibold">{{ portfolio.title }}</h3>
            <p class="sub-h1-regular">{{ portfolio.description }}</p>
          </div>
        </div>
      </div>
    </div>
    <!-- Mobile UI ends -->
    <!-- Desktop UI -->
    <div id="stickyParent" class="hidden lg:block h-[150vh] xl:h-[110vh]">
      <div class="sticky top-0 max-h-full main">
        <div v-for="(portfolio, index) in portfolios" :key="index">
          <div class="flex bg-[#F8F8F8] h-screen">
            <div class="flex-1">
              <img
                :src="portfolio.image[2]"
                :srcset="`${portfolio.image[2]} 800w,${portfolio.image[3]} 1200w`"
                :alt="portfolio.title"
                class="w-full h-full object-contain opacity-0"
                loading="lazy"
                :class="{ 'animate-fadeInLeft opacity-100': isAnimated }"
              />
            </div>
            <div
              class="flex-1 relative text-left text-white bg-[#EF233C] opacity-0"
              :class="{ 'animate-moveIn opacity-100': isAnimated }"
            >
              <div
                class="absolute -left-4 top-1/2 -translate-y-1/2 w-0 h-0 border-transparent border-t-[18px] border-b-[18px] border-r-[18px] border-r-[#EF223C] border-l-solid"
              ></div>
              <div
                class="absolute flex flex-col gap-6 w-[27.25rem] top-1/2 left-14 -translate-y-1/2"
              >
                <h3 class="desk-header-3">{{ portfolio.title }}</h3>
                <span class="mobile-header-2-regular">{{
                  portfolio.description
                }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <!-- Desktop UI ends -->
  </div>
</template>
<script setup type="module">
import mobile_reelnews1 from "@/assets/images/flutter-app-development/casestudy/mobile/p1-400w.webp";
import mobile_reelnews2 from "@/assets/images/flutter-app-development/casestudy/mobile/p1-800w.webp";
import desktop_reelnews1 from "@/assets/images/flutter-app-development/casestudy/p1-800w.webp";
import desktop_reelnews2 from "@/assets/images/flutter-app-development/casestudy/p1-1200w.webp";

const isAnimated = ref(false);
const portfolios = [
  {
    title: "Reelnews",
    image: [
      mobile_reelnews1,
      mobile_reelnews2,
      desktop_reelnews1,
      desktop_reelnews2,
    ],
    description:
      "Ditch Old News, Get reelnews: The News App for the Upcoming Wave of News Creators and Consumers.",
  },
];

onMounted(() => {
  window.addEventListener("scroll", handleScroll);
});
onUnmounted(() => {
  window.removeEventListener("scroll", handleScroll);
});

const handleScroll = () => {
  const sections = document.querySelectorAll(".main");
  const windowHeight = window.innerHeight;

  sections.forEach((section) => {
    const react = section.getBoundingClientRect();
    if (react.top <= windowHeight / 4 && react.bottom >= 0) {
      isAnimated.value = true;
    }
  });
};
</script>
