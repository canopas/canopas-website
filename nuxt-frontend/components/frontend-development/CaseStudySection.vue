<template>
  <section class="mt-16 lg:!mt-60">
    <div
      class="container mb-2 lg:-mb-4 2xl:-mb-12 xll:-mb-52 3xl:-mb-[35rem] text-center"
    >
      <h2 class="mobile-header-2 lg:desk-header-2 text-black-87">
        Case studies
      </h2>
      <p
        class="sub-h1-regular lg:mobile-header-2-regular mt-4 lg:mt-6 text-black-60 xs:px-4"
      >
        Explore our case studies to witness how we've transformed our
        <br class="hidden md:inline-block" />client's ideas into successful
        websites.
      </p>
    </div>
    <!-- Mobile View starts -->
    <div class="lg:hidden mb-16">
      <div
        v-for="(portfolio, index) in portfolios"
        :key="index"
        class="mt-4 py-12"
        :class="portfolio.bgcolor"
      >
        <div class="container">
          <img
            :src="portfolio.image[0]"
            :srcset="`${portfolio.image[0]} 400w,${portfolio.image[1]} 800w`"
            :alt="portfolio.title"
            class="w-full h-full object-cover"
          />
          <h2 class="mobile-header-2 text-black-87 mt-2">
            {{ portfolio.title }}
          </h2>
          <p class="sub-h1-regular text-black-60 mt-2">
            {{ portfolio.description }}
          </p>
        </div>
      </div>
    </div>
    <!-- Mobile View ends -->
    <!-- Desktop View starts -->
    <div class="h-[200vh] mb-32 hidden lg:block" id="parentContainer">
      <div
        class="sticky top-0 h-screen main overflow-hidden snap-y snap-mandatory scroll-smooth"
      >
        <div
          v-for="(portfolio, index) in portfolios"
          :key="index"
          :id="`portfolio-${index}`"
          class="section text-center h-screen flex justify-evenly snap-start opacity-0"
        >
          <div
            class="flex flex-col w-[75rem] py-28 container rounded-xl"
            :class="portfolio.bgdesk"
          >
            <img
              :src="portfolio.image[0]"
              :srcset="`${portfolio.image[0]} 400w,${portfolio.image[1]} 800w`"
              :alt="portfolio.title"
              class="w-[36rem] h-96 object-cover m-auto"
              loading="lazy"
            />
            <h2 class="mobile-header-2 text-black-87 mt-2">
              {{ portfolio.title }}
            </h2>
            <p class="sub-h1-regular text-black-60 mt-2">
              {{ portfolio.description }}
            </p>
          </div>
        </div>
      </div>
    </div>
    <!-- Desktop View Ends -->
  </section>
</template>
<script setup>
import web400 from "@/assets/images/frontend-development/casestudy/web-400w.webp";
import web800 from "@/assets/images/frontend-development/casestudy/web-800w.webp";
import luxe400 from "@/assets/images/frontend-development/casestudy/luxe-400w.webp";
import luxe800 from "@/assets/images/frontend-development/casestudy/luxe-800w.webp";

const portfolios = [
  {
    title: "Canopas",
    image: [web400, web800],
    description: "Web and mobile app development company",
    bgcolor: "bg-web-gradient-background",
    bgdesk: "bg-web-desk-gradient-background",
  },
  {
    title: "Luxeradio",
    image: [luxe400, luxe800],
    description: "New personalized radio and modern art gallery.",
    bgcolor: "bg-luxeradio-gradient-background",
    bgdesk: "bg-luxeradio-desk-gradient-background",
  },
];
onMounted(() => {
  window.addEventListener("scroll", handleScroll);
});
onUnmounted(() => {
  window.removeEventListener("scroll", handleScroll);
});
function handleScroll() {
  const parentContainer = document.getElementById("parentContainer");
  const parentTop = parentContainer.offsetTop;
  const childContainer = document.querySelector(".main");
  const scrollPosition = window.scrollY || window.pageYOffset;
  const sections = document.querySelectorAll(".section");
  if (scrollPosition >= parentTop) {
    const scrollAmount = scrollPosition - parentTop;
    childContainer.scrollTop = scrollAmount;
  }
  sections.forEach((section) => {
    const itemReact = section.getBoundingClientRect();
    if (itemReact.top < 350 || childContainer.offsetTop - parentTop > 500) {
      section.classList.add("animate-popUp");
    }
  });
}
</script>
