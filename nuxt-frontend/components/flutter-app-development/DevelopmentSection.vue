<template>
  <div
    class="py-14 lg:py-[7.5rem] bg-[#121212] text-white text-center 3xl:outer-container"
  >
    <div class="container flex flex-col gap-4 lg:gap-6 lg:pb-[4.5rem]">
      <span class="mobile-header-2 lg:desk-header-2">
        Flutter app development services
      </span>
      <span
        class="sub-h1-regular lg:mobile-header-2-regular text-white-core-80"
      >
        We leverage Flutter's rich, customizable widgets and fast development
        capabilities to create beautiful and user-friendly apps that save you
        time and cost of platform-specific development.
      </span>
    </div>
    <!-- Flutter Mobile Swiper UI-->
    <div class="swiper-content mt-6 lg:hidden text-left">
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
            class="flex flex-col gap-4 py-8 px-4 bg-white-core-90 rounded-xl"
          >
            <div class="mobile-header-3-semibold">
              <span v-html="item.title"></span>
            </div>
            <div class="sub-h3-medium text-white-core-80">
              <span
                v-html="item.description"
                class="flex gap-4 flex-col text-white-core-80"
              ></span>
            </div>
          </div>
        </swiper-slide>
        <div class="swiper-pagination"></div>
      </swiper>
    </div>
    <!-- Flutter Desktop UI -->
    <div class="hidden lg:block relative h-full overflow-hidden">
      <hr class="absolute h-[300vh] left-[50%] border-[0.1px] opacity-[0.5]" />
      <div
        v-for="(item, index) in items"
        :key="index"
        class="mb-[4.5rem] container item opacity-50"
        :class="{ '!opacity-100': index === activeIndex }"
      >
        <div
          :class="[
            ' flex gap-16',
            {
              ' flex-row text-right ': index % 2 === 0,
              ' flex-row-reverse text-left': index % 2 !== 0,
            },
          ]"
        >
          <div class="flex flex-col flex-1 gap-6">
            <div class="desk-header-3">
              <span v-html="item.title"></span>
            </div>
            <div class="mobile-header-3-medium">
              <span v-html="item.description"></span>
            </div>
          </div>
          <div class="flex-1" :class="{ relative: index % 2 !== 0 }">
            <img
              :src="item.image"
              alt="development-service-image"
              class="w-24 h-24"
              :class="{ 'absolute right-0': index % 2 !== 0 }"
              loading="lazy"
            />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup type="module">
import { Autoplay, Pagination } from "swiper/modules";
import { Swiper, SwiperSlide } from "swiper/vue";
import consultation from "@/assets/images/flutter-app-development/development/consultation.webp";
import design from "@/assets/images/flutter-app-development/development/design.webp";
import development from "@/assets/images/flutter-app-development/development/development.webp";
import deployment from "@/assets/images/flutter-app-development/development/deployment.webp";
import maintainance from "@/assets/images/flutter-app-development/development/maintainance.webp";

const modules = [Autoplay, Pagination];
var activeIndex = ref(0);
const items = [
  {
    title: `Flutter App Development <br/>Consultation`,
    image: consultation,
    description: `<span>We begin every project by understanding your vision and business requirements.</span>
            <span>Our expert consultants guide you through the entire process, providing insights on market trends,competition, technical feasibility, and ideal app monetization strategies.</span>
            <span>We dive deep into your project's specifics, brainstorming innovative ways to bring your idea to life using Flutter's powerful features.</span>`,
  },
  {
    title: "Flutter UI/UX Design",
    image: design,
    description: `<span>We believe in designing applications that are not just aestheticaly pleasing, but also remarkably intuitive and user-friendly.</span>
            <span>With Flutter's extensive widget library and flexible UI capabilities, we craft customized, interactive,and smooth interfaces that captivate your users and keep them engaged.</span>
            <span>We place a strong emphasis on user-centric design, incorporating the latest design principles to create a seamless,intuitive flow. From the initial wireframe to the final design mockup, we ensure every pixel adds value to your user's journey, making your app stand out in the bustling app marketplace.</span>
          `,
  },
  {
    title: `Custom Flutter Application <br /> Development`,
    image: development,
    description: `<span>We transform your unique business needs into robust, efficient, and beautifully designed mobile apps.</span>
            <span>We leverage Flutter's single codebase feature, allowing us to expedite the development process without compromising on the app's performance, feel, or look. Our team employs Flutter's hot-reload functionality to build dynamic, high-quality applications that offer a native experience on both iOS and Android platforms.</span>
            <span>From feature-rich eCommerce apps to interactive social networking apps, we deliver custom Flutter solutions that align with your business goals and deliver a seamless user experience across all devices.</span>`,
  },
  {
    title: "Flutter App Deployment",
    image: deployment,
    description: `
            <span>We understand that app development doesn't end with coding. Ensuring that your application gets successfully deployed on the App Store and Google Play Store is equally important.</span>
            <span>Our team navigates the complexities of app submission guidelines, takes care of the application's configuration, manages the app store metadata, and executes a successful launch.</span>
            <span>We also monitor the app's performance post-launch, ensuring smooth functioning and immediate troubleshooting. Experience a worry-free deployment and watch your Flutter application take flight with our diligent deployment services.</span>
          `,
  },
  {
    title: "Maintenance and Support",
    image: maintainance,
    description: `
            <span>We believe that maintaining an application is as crucial as developing it. After your app's successful deployment, we provide ongoing support to ensure its optimal performance, stability, and user satisfaction.</span>
            <span>Our services include regular updates, bug fixes, performance monitoring, code optimization, and adding new features as per market trends or user feedback.</span>
            <span>With our round-the-clock support, you can rest assured that your application stays up-to-date,secure, and relevant in the ever-evolving app market, ensuring its long-term success and user engagement.</span>
          `,
  },
];
const pagination = {
  el: ".swiper-pagination",
  clickable: true,
};
onMounted(() => {
  window.addEventListener("scroll", handleScroll);
});

onUnmounted(() => {
  window.removeEventListener("scroll", handleScroll);
});
function handleScroll() {
  const child = document.querySelectorAll(".item");
  child.forEach((item, index) => {
    const itemRect = item.getBoundingClientRect();
    const itemTop = itemRect.top;
    if (itemTop < 350) {
      activeIndex.value = index;
    }
  });
}
</script>
<style lang="postcss" scoped>
@import "swiper/css";
@import "swiper/css/pagination";

.swiper-pagination {
  @apply static pt-8;
}
.swiper-pagination-bullet {
  @apply !w-3 !h-3 !rounded-full !bg-pink-90 !border !border-solid !border-white-core-80;
}
.swiper-pagination-bullet-active {
  @apply !border-none !opacity-100;
}
</style>
