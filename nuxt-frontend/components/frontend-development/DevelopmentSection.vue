<template>
  <section
    class="my-16 lg:mt-60 lg:mb-0 relative 3xl:h-[280vh] 3xl:outer-container"
  >
    <div
      class="flex flex-col gap-4 lg:gap-6 container text-center lg:sticky top-28 3xl:!pb-0"
      id="servicesContainer"
    >
      <h2 class="text-black-87 mobile-header-2 lg:desk-header-2">
        Frontend <br class="md:hidden" />development services
      </h2>
      <p
        class="sub-h1-regular lg:mobile-header-2-regular text-black-60 xs:px-2 2xl:!px-32"
      >
        At Canopas, we understand the power of a stellar frontend interface. We
        believe that a potent blend of aesthetics, responsiveness, and intuitive
        design is paramount to captivate your audience and drive engagement. We
        aim to create digital platforms that don't just tell your brand's story
        but also offer an unparalleled user experience.
      </p>
    </div>
    <!-- Mobile Swiper UI starts-->
    <div class="swiper-content mt-6 lg:hidden">
      <swiper
        :slidesPerView="1"
        :centeredSlides="true"
        :spaceBetween="10"
        :modules="modules"
        :autoplay="{
          delay: 3000,
          disableOnInteraction: false,
        }"
        :speed="1500"
        :pagination="pagination"
        :breakpoints="{
          '540': {
            slidesPerView: 1.2,
            spaceBetween: 0,
          },
          '620': {
            slidesPerView: 1.4,
            spaceBetween: 0,
          },
          '700': {
            slidesPerView: 1.5,
          },
        }"
        class="swiper-container"
      >
        <swiper-slide
          v-for="(item, index) in items"
          :key="index"
          class="cursor-pointer px-4"
        >
          <div
            :class="{
              'bg-pink-40': index % 2 === 0,
              'bg-white-smoke': index % 2 !== 0,
            }"
            class="container py-6 rounded-xl"
          >
            <h3 class="mobile-header-3-semibold text-black-87">
              {{ item.title }}
            </h3>
            <div
              class="sub-h3-regular text-black-60 mt-4 flex gap-4 flex-col"
              v-html="item.description"
            ></div>
          </div>
        </swiper-slide>
        <div class="swiper-pagination"></div>
      </swiper>
    </div>
    <!-- Mobile Swiper UI ends-->
    <!-- Desktop Swiper UI starts-->
    <div class="pt-24 xl:pt-20 main hidden lg:flex flex-col gap-32">
      <div
        v-for="(item, index) in items"
        :key="index"
        class="services bg-white sticky top-96 mb-52"
        :id="`${'item-' + index}`"
      >
        <div
          :class="{
            'bg-pink-40': index % 2 === 0,
            'bg-white-smoke': index % 2 !== 0,
          }"
          class="p-10 container rounded-xl w-[45rem] shadow-lg h-[33rem]"
        >
          <h3 class="desk-header-3 text-black-87">
            {{ index + 1 }}.
            {{ item.title }}
          </h3>
          <div
            class="mobile-header-3-medium text-black-60 mt-4 flex gap-4 flex-col"
            v-html="item.description"
          ></div>
        </div>
      </div>
    </div>
    <!-- Desktop Swiper UI ends-->
  </section>
</template>
<script setup type="module">
import { Autoplay, Pagination } from "swiper/modules";
import { Swiper, SwiperSlide } from "swiper/vue";

const modules = [Autoplay, Pagination];
const pagination = {
  el: ".swiper-pagination",
  clickable: true,
};
const items = [
  {
    title: "Responsive Web Design",
    description: `<span>As you know, mobile usage has been growing every day for more than a decade now.</span>
            <span>In such a time, a responsive design isn't a luxury — it's a necessity.</span>
            <span>We ensure your website looks and functions flawlessly, be it on a desktop, tablet, or mobile. Our team will help you design pixel perfect designs for all types of devices.</span>`,
    className: "rotate-6",
  },
  {
    title: "Single Page Applications (SPAs)",
    description: `<span>By dynamically loading content and eliminating traditional page reloads, SPAs offer users a fluid, app-like experience within their browsers.</span>
            <span>At Canopas, we specialize in leveraging modern frameworks such as React, Angular, and Vue.js to craft these responsive and interactive platforms.</span>
            <span>Coupled with seamless backend integrations, our SPAs are not only user-friendly but also robust, ensuring fast performance and a smooth browsing experience.</span>`,
    className: "-rotate-6",
  },
  {
    title: "Progressive Web Apps (PWAs)",
    description: `<span>Progressive Web Apps (PWAs) Offer the best of both worlds, delivering web accessibility with the immersive experience of a native app.</span>
            <span>At Canopas, we recognize the transformative potential of PWAs. By harnessing their offline capabilities, push notifications, and device hardware access, we create digital platforms that engage users seamlessly, regardless of network conditions or device type.</span>
            <span>We ensure your website looks and functions flawlessly, be it on a desktop, tablet, or mobile. Our team will help you design pixel perfect designs for all types of devices.</span>`,
    className: "rotate-6",
  },
  {
    title: "UI/UX Design",
    description: `<span>As you know, exceptional digital experiences begin with intuitive UI/UX Design, the art and science of crafting user interfaces that resonate and user experiences that captivate.</span>
            <span>We delve deep into user behaviors, needs, and motivations to create designs that are not only visually appealing but also functionally streamlined.</span>
            <span>Every pixel we craft and every interaction we design is aimed at providing users with a seamless journey, turning visitors into loyal customers. </span>
            <span>By blending aesthetics with ergonomics, we ensure that our digital solutions are both beautiful and user-centric.</span>`,
    className: "!mb-44 -rotate-6",
  },
  {
    title: "Frontend Audits and Consultation",
    description: `<span>Not sure where your current frontend stands? Our expert audits dive deep into your platform, identifying areas of improvement, and providing actionable insights for enhancement.</span>
            <span>Leveraging a combination of cutting-edge tools and seasoned expertise, we dive deep into your web architecture, pinpointing areas that can benefit from refinement.</span>
            <span>Our aim? To elevate your online platform's functionality, accessibility, and user engagement, turning potential weak points into strengths.</span>`,
    className: `!mb-60 rotate-6`,
  },
  {
    title: "Performance Optimization",
    description: `<span>User patience is finite. We fine-tune your frontend assets, ensuring lightningfast load times, reduced bounce rates, and an overall smoother user experience.</span>
            <span>Through a meticulous analysis of your site's architecture, load times, and asset delivery, we pinpoint bottlenecks and implement strategic enhancements.</span>
            <span>The result? A digital experience that not only loads faster but retains users, boosts engagement, and improves overall site metrics.</span>`,
  },
];
onMounted(() => {
  window.addEventListener("scroll", handleScroll);
});
onUnmounted(() => {
  window.removeEventListener("scroll", handleScroll);
});
function handleScroll() {
  const service = document.querySelectorAll(".services");
  const parent = document.getElementById("servicesContainer");
  service.forEach((item, index) => {
    const itemRect = item.getBoundingClientRect();
    const itemTop = itemRect.top;
    if (itemTop < 380 && index == 5) {
      parent.classList.add("lg:!pb-[48rem]");
    } else if (itemTop > 420) {
      parent.classList.remove("lg:!pb-[48rem]");
    }
  });
}
</script>
<style lang="postcss">
@import "swiper/css";
@import "swiper/css/pagination";

.swiper-pagination {
  @apply static !mt-8;
}
.swiper-pagination-bullet {
  @apply !w-3 !h-3 !rounded-full !bg-pink-90 !border !border-solid !border-white-core-80;
}
.swiper-pagination-bullet-active {
  @apply !border-none !opacity-100;
}
.child-content:last-child {
  margin-bottom: 13rem;
}
</style>
