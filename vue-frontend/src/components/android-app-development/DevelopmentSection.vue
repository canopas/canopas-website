<template>
  <div class="tw-mt-16 md:tw-mt-40 xl:tw-mt-60 tw-bg-white tw-h-[70%]">
    <div
      class="xl:tw-h-[0vh] xl:!tw-sticky"
      :class="{
        'xl:!tw-top-[128px] 3xl:!tw-top-[500px]':
          lastSlideReached && !isScrollingUp,
        'xl:!tw-top-[8rem] 3xl:!tw-top-[500px]':
          lastSlideReached && isScrollingUp,
      }"
    >
      <div
        class="tw-container tw-flex tw-flex-col tw-text-center xll:tw--mb-72"
      >
        <span
          class="header-2 tw-text-black-core/[0.87] lg:!tw-mx-0 lg:!tw-w-full xs:tw-mx-auto xs:tw-w-[71%]"
        >
          Android app development services</span
        >
        <span
          class="tw-mt-4 md:tw-mt-6 sub-h1-regular tw-text-black-core/[0.60]"
          >With over a decade of expertise in Android app development, our team
          can help you craft high- quality, user-friendly apps that drive
          business growth and customer engagement.</span
        >
      </div>
    </div>
    <!-- ---------------------MOBILE UI START------------------ -->
    <div class="swiper-content !tw-mt-6 md:!tw-mt-12 lg:tw-hidden">
      <swiper
        :slidesPerView="1.1"
        :centeredSlides="true"
        :spaceBetween="10"
        :modules="modules"
        class="swiper-container"
        @swiper="setSwiperRef"
        :breakpoints="{
          '768': {
            slidesPerView: 1.3,
          },
          '992': {
            slidesPerView: 1.5,
          },
          '1200': {
            slidesPerView: 1.6,
          },
          '1400': {
            slidesPerView: 1.9,
          },
          '1600': {
            slidesPerView: 2.1,
          },
          '1800': {
            slidesPerView: 2.4,
          },
          '2000': {
            slidesPerView: 2.6,
          },
          '2200': {
            slidesPerView: 2.8,
          },
          '2500': {
            slidesPerView: 3.2,
          },
          '2800': {
            slidesPerView: 3.8,
          },
          '3500': {
            slidesPerView: 4.5,
          },
        }"
      >
        <swiper-slide
          v-for="(item, index) in items"
          :key="index"
          class="tw-cursor-pointer"
        >
          <div
            class="tw-mr-4 tw-rounded-[1.25rem] tw-px-4 tw-py-6 md:tw-w-[34.7rem] lg:tw-w-[40.75rem] xl:!tw-w-[45.75rem]"
            :class="[item.bgColor]"
          >
            <div
              class="tw-justify-left tw-mb-5 tw-flex tw-flex-row tw-items-center tw-space-x-4"
            >
              <img
                alt="development-icon"
                :src="item.icon"
                class="tw-h-10 tw-w-10 tw-object-contain lg:tw-hidden"
                loading="lazy"
              />
              <span class="header-3 tw-text-black-core/[0.87]">{{
                item.title
              }}</span>
            </div>
            <div class="sub-h3-regular tw-text-black-core/[0.60] lg:tw-p-4">
              <ul class="tw-pl-8 tw-pr-4 lg:tw-hidden">
                <li
                  v-for="list in item.description"
                  :key="list"
                  class="tw-mb-3.5 tw-list-disc"
                >
                  {{ list }}
                </li>
              </ul>
            </div>
          </div>
        </swiper-slide>
      </swiper>
    </div>
    <!-- ---------------------MOBILE UI END------------------ -->

    <!-- ---------------------DESKTOP UI START------------------ -->
    <div class="sticky-parent tw-hidden lg:tw-block">
      <div
        class="sticky tw-sticky tw-top-[92px] xl:!tw-top-[100px] 3xl:!tw-top-[470px] xl:tw-max-h-full main tw-overflow-hidden"
        :class="{ 'xl:!tw-top-[100px]': isScrollingUp }"
      >
        <div class="tw-flex">
          <div
            v-for="(item, index) in items"
            :key="index"
            class="tw-cursor-pointer tw-mt--48 xl:tw-mt-32 xll:tw-mt-[36rem] tw-min-w-[100vw] tw-h-[900px] xl:tw-h-full"
          >
            <div
              class="tw-relative !-tw-mr-2 tw-mb-12 tw-mt-[6rem] xl:tw-mt-[6.5rem] xll:tw-mt-[-15rem]"
            >
              <div
                class="tw-border-b-2 tw-border-dotted tw-border-[#000000]/[0.6] tw-text-center"
              ></div>
            </div>
            <div
              class="tw-absolute tw-top-[7%] xl:tw-top-[23%] xll:tw-top-[31.5%] tw-ml-[45%] xll:tw-ml-[47%] tw-hidden tw-h-20 !tw-bg-white tw-text-center lg:tw-block 2xll:tw-ml-[47%] xll:tw-ml-[48%]"
            >
              <span class="tw-text-center header-1 tw-text-black-core/[0.87]">
                0{{ index + 1 }}</span
              >
            </div>
            <div
              class="tw-ml-auto xll:tw-ml-[12%] xll:tw-ml-auto tw-rounded-[1.25rem] tw-p-[1.88rem] !tw-w-[39.75rem] xl:!tw-w-[45rem] tw-mx-auto"
              :class="[item.bgColor]"
            >
              <div class="tw-justify-left tw-flex tw-flex-row tw-items-center">
                <img
                  alt="development-icon"
                  :src="item.icon"
                  class="tw-h-10 tw-w-10 tw-object-contain lg:tw-hidden"
                  loading="lazy"
                />
                <span
                  class="tw-font-opensans-semibold tw-text-black-core/[0.87] header-3"
                  >{{ item.title }}</span
                >
              </div>
              <div class="tw-text-black-core/[0.6] sub-h3-medium">
                <ul class="tw-pl-8 tw-pr-4 lg:tw-hidden">
                  <li
                    v-for="list in item.description"
                    :key="list"
                    class="tw-mb-3.5 tw-list-disc"
                  >
                    {{ list }}
                  </li>
                </ul>
                <p
                  v-for="list in item.description"
                  :key="list"
                  class="tw-mt-4 tw-hidden lg:tw-block"
                >
                  {{ list }}
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div
      class="tw-hidden lg:tw-block tw-h-[25vh] xl:tw-h-[70vh] xll:tw-h-[60vh] tw-w-screen"
    ></div>
    <!-- ---------------------DESKTOP UI END------------------ -->
  </div>
</template>

<script type="module">
import consultation from "@/assets/images/andriod-app-development/development/1.webp";
import design from "@/assets/images/andriod-app-development/development/2.webp";
import custom from "@/assets/images/andriod-app-development/development/3.webp";
import app from "@/assets/images/andriod-app-development/development/4.webp";
import maintenance from "@/assets/images/andriod-app-development/development/5.webp";
import { Swiper, SwiperSlide } from "swiper/vue";
export default {
  data() {
    return {
      isScrollingUp: false,
      currentCardIndex: 0,
      scrollPosition: 0,
      lastSlideReached: false,
      items: [
        {
          title: "Android App Development Consultation",
          bgColor: "tw-bg-[#F2709C]/[0.04]",
          icon: consultation,
          description: [
            "At Canopas, we believe that a successful project begins with a thorough understanding of your vision, goals, and market dynamics.",

            "Our team of seasoned professionals collabortively works with you during the consultation phase,using their deep expertise to advise on optimal strategies, potential challenges, and innovative solutions.",

            "We meticulously assess your project requirements, suggest suitable tech stacks, provide a realistic timeline, and help you understand the project's potential.",
          ],
        },
        {
          title: "Android UI/UX Design",
          bgColor: "tw-bg-[#F8F8F8]",
          icon: design,
          description: [
            "At Canopas, our design philosophy is centered on creating intuitive, user-friendly, and visually stunning interfaces that resonate with your target audience.",

            "We prioritize the user journey, ensuring seamless navigation, quick loading times, and attractive design elements that captivate and retain users. Our team stays updated with the latest design trends, including its varied device screen sizes and resolutions.",

            "From wireframing to prototyping to the final design, we work closely with you at each step,translating your vision into a compelling app design that looks great and delivers an unmatched user experience.",
          ],
        },
        {
          title: "Custom Android Application Development",
          bgColor: "tw-bg-[#F2709C]/[0.04]",
          icon: custom,
          description: [
            "Our approach to development is rooted in agility and focused collaboration.  We build in two-week sprints where our team devotes their complete energy and attention to sculpting the foundation of your app.",

            "With check-ins every two weeks, we ensure you're part of the development journey, giving you the opportunity to review our progress and provide valuable feedback. This leads to a development process that's streamlined and dynamic.",

            "Our aim is to create apps that leave a lasting impression, spur conversations, and consistently draw users back. With our team of professionals, we bring your ideas to life, crafting elegant, all-encompassing designs that transform your vision into a winning app in the competitive marketplace.",
          ],
        },
        {
          title: "Android App Deployment",
          bgColor: "tw-bg-[#F8F8F8]",
          icon: app,
          description: [
            "Launching your Android app into the marketplace involves more than just hitting 'publish'. At Canopas, we'll guide you through this crucial deployment phase with careful planning, meticulous testing, and comprehensive support.",

            "Our team ensures your app complies with all of Google Play Store's guidelines, optimizing everything from the app description and graphics to backend configurations for a smooth launch.",

            "We also devise and implement a robust pre and post-launch marketing strategy to build hype and drive initial downloads. Your success is our end goal, and we stand by you as your Android app takes its first steps into the digital world.",
          ],
        },
        {
          title: "Maintenance and Support",
          bgColor: "tw-bg-[#F2709C]/[0.04]",
          icon: maintenance,
          description: [
            "Post-deployment, we monitor the app's performance and user feedback closely, ready to make quick updates or fixes if needed. Our maintenance and support services ensure your app stays relevant, competitive, and effective in the face of new devices, OS updates, and emerging user expectations.",

            "We swiftly handle any emerging issues, from minor bug fixes to major updates, all aimed at enhancing user experience and engagement.",

            "Moreover, we believe in proactive communication, providing you with regular reports and insights into your app's performance and user feedback, while being readily available to answer any queries or concerns. With our ongoing maintenance and support, you can focus on your business, knowing your app's functionality and relevance are in expert hands.",
          ],
        },
      ],
    };
  },
  mounted() {
    this.width = window.innerWidth;
    document.addEventListener("scroll", this.handleScroll);

    let sticky = document.querySelector(".sticky");
    let stickyParent = document.querySelector(".sticky-parent");

    let scrollWidth = sticky.scrollWidth;
    let verticalScrollHeight =
      stickyParent.getBoundingClientRect().height -
      sticky.getBoundingClientRect().height;

    function horizontalScroll() {
      let scrolled = stickyParent.getBoundingClientRect().top;

      if (scrolled > 0) {
        sticky.scrollLeft = 0;
      } else {
        sticky.scrollLeft =
          (scrollWidth / verticalScrollHeight) * -scrolled * 0.85;
      }
      requestAnimationFrame(horizontalScroll);
    }

    // Attach the existing horizontalScroll function separately
    document.addEventListener("scroll", horizontalScroll);
  },
  beforeDestroy() {
    document.removeEventListener("scroll", this.handleScroll);

    document.removeEventListener("scroll", this.horizontalScroll);
  },
  methods: {
    handleScroll() {
      const windowHeight = window.innerHeight;
      const lastSlide = document.querySelector(".main:last-child");
      const lastSlidePosition = lastSlide.getBoundingClientRect().top;
      let threshold = windowHeight * 0.2;
      if (this.width < 3840) {
        threshold = 56.84375;
      }
      if (lastSlidePosition <= windowHeight && lastSlidePosition >= threshold) {
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

<style lang="postcss" scoped>
@import "swiper/css";
.sticky-parent {
  @apply tw-h-[700vh] xll:!tw-h-[400vh];
}
</style>
