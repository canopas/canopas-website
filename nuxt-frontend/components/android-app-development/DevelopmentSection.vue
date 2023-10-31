<template>
  <div class="mt-16 md:mt-40 xl:mt-60 bg-white h-[70%]">
    <div class="sticky-parent">
      <div class="lg:sticky lg:!top-[128px] 3xl:!top-[500px]">
        <div class="container flex flex-col text-center xll:-mb-72">
          <span
            class="header-2 text-black-core/[0.87] lg:!mx-0 lg:!w-full xs:mx-auto xs:w-[71%]"
          >
            Android app development services</span
          >
          <span class="mt-4 md:mt-6 sub-h1-regular text-black-core/[0.60]"
            >With over a decade of expertise in Android app development, our
            team can help you craft high- quality, user-friendly apps that drive
            business growth and customer engagement.</span
          >
        </div>
        <!-- ---------------------MOBILE UI START------------------ -->
        <div class="swiper-content !mt-6 md:!mt-12 lg:hidden">
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
              class="cursor-pointer"
            >
              <div
                class="mr-4 rounded-[1.25rem] px-4 py-6 md:w-[34.7rem] lg:w-[40.75rem] xl:!w-[45.75rem]"
                :class="[item.bgColor]"
              >
                <div
                  class="justify-left mb-5 flex flex-row items-center space-x-4"
                >
                  <img
                    alt="development-icon"
                    :src="item.icon"
                    class="h-10 w-10 object-contain lg:hidden"
                    loading="lazy"
                  />
                  <span class="header-3 text-black-core/[0.87]">{{
                    item.title
                  }}</span>
                </div>
                <div class="sub-h3-regular text-black-core/[0.60] lg:p-4">
                  <ul class="pl-8 pr-4 lg:hidden">
                    <li
                      v-for="list in item.description"
                      :key="list"
                      class="mb-3.5 list-disc"
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

        <div class="sticky main overflow-hidden hidden lg:block">
          <div class="flex">
            <div
              v-for="(item, index) in items"
              :key="index"
              class="cursor-pointer -mt-12 xll:mt-[43rem] min-w-[100vw] h-[900px] xl:h-full"
            >
              <div
                class="relative !-mr-2 mb-12 mt-24 xl:mt-[6.5rem] xll:-mt-60"
              >
                <div
                  class="border-b-2 border-dotted border-[#000000]/[0.6] text-center"
                ></div>
              </div>
              <div
                class="absolute top-[2%] xl:top-[15px] xll:top-[37.5%] ml-[45%] xll:ml-[47%] hidden h-20 !bg-white text-center lg:block 2xll:ml-[47%] xll:ml-[48%]"
              >
                <span class="text-center header-1 text-black-core/[0.87]">
                  0{{ index + 1 }}</span
                >
              </div>
              <div
                class="ml-auto xll:ml-[12%] xll:ml-auto rounded-[1.25rem] p-[1.88rem] !w-[39.75rem] xl:!w-[45rem] mx-auto"
                :class="[item.bgColor]"
              >
                <div class="justify-left flex flex-row items-center">
                  <img
                    alt="development-icon"
                    :src="item.icon"
                    class="h-10 w-10 object-contain lg:hidden"
                    loading="lazy"
                  />
                  <span
                    class="font-opensans-semibold text-black-core/[0.87] header-3"
                    >{{ item.title }}</span
                  >
                </div>
                <div class="text-black-core/[0.6] sub-h3-medium">
                  <ul class="pl-8 pr-4 lg:hidden">
                    <li
                      v-for="list in item.description"
                      :key="list"
                      class="mb-3.5 list-disc"
                    >
                      {{ list }}
                    </li>
                  </ul>
                  <p
                    v-for="list in item.description"
                    :key="list"
                    class="mt-4 hidden lg:block"
                  >
                    {{ list }}
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div
      class="hidden lg:block h-[25vh] xl:h-[70vh] xll:h-[60vh] w-screen"
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
      items: [
        {
          title: "Android App Development Consultation",
          bgColor: "bg-[#F2709C]/[0.04]",
          icon: consultation,
          description: [
            "At Canopas, we believe that a successful project begins with a thorough understanding of your vision, goals, and market dynamics.",

            "Our team of seasoned professionals collabortively works with you during the consultation phase,using their deep expertise to advise on optimal strategies, potential challenges, and innovative solutions.",

            "We meticulously assess your project requirements, suggest suitable tech stacks, provide a realistic timeline, and help you understand the project's potential.",
          ],
        },
        {
          title: "Android UI/UX Design",
          bgColor: "bg-[#F8F8F8]",
          icon: design,
          description: [
            "At Canopas, our design philosophy is centered on creating intuitive, user-friendly, and visually stunning interfaces that resonate with your target audience.",

            "We prioritize the user journey, ensuring seamless navigation, quick loading times, and attractive design elements that captivate and retain users. Our team stays updated with the latest design trends, including its varied device screen sizes and resolutions.",

            "From wireframing to prototyping to the final design, we work closely with you at each step,translating your vision into a compelling app design that looks great and delivers an unmatched user experience.",
          ],
        },
        {
          title: "Custom Android Application Development",
          bgColor: "bg-[#F2709C]/[0.04]",
          icon: custom,
          description: [
            "Our approach to development is rooted in agility and focused collaboration.  We build in two-week sprints where our team devotes their complete energy and attention to sculpting the foundation of your app.",

            "With check-ins every two weeks, we ensure you're part of the development journey, giving you the opportunity to review our progress and provide valuable feedback. This leads to a development process that's streamlined and dynamic.",

            "Our aim is to create apps that leave a lasting impression, spur conversations, and consistently draw users back. With our team of professionals, we bring your ideas to life, crafting elegant, all-encompassing designs that transform your vision into a winning app in the competitive marketplace.",
          ],
        },
        {
          title: "Android App Deployment",
          bgColor: "bg-[#F8F8F8]",
          icon: app,
          description: [
            "Launching your Android app into the marketplace involves more than just hitting 'publish'. At Canopas, we'll guide you through this crucial deployment phase with careful planning, meticulous testing, and comprehensive support.",

            "Our team ensures your app complies with all of Google Play Store's guidelines, optimizing everything from the app description and graphics to backend configurations for a smooth launch.",

            "We also devise and implement a robust pre and post-launch marketing strategy to build hype and drive initial downloads. Your success is our end goal, and we stand by you as your Android app takes its first steps into the digital world.",
          ],
        },
        {
          title: "Maintenance and Support",
          bgColor: "bg-[#F2709C]/[0.04]",
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
    let sticky = document.querySelector(".main");
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
          (scrollWidth / verticalScrollHeight) * -scrolled * 0.9;
      }
    }
    document.addEventListener("scroll", horizontalScroll);
  },
  beforeDestroy() {
    document.removeEventListener("scroll", this.horizontalScroll);
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
  @apply lg:h-[700vh] xll:!h-[400vh];
}
</style>
